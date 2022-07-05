package routes

import (
	"GolangLivePT01/golang_web_programming/app/logo"
	membership "GolangLivePT01/golang_web_programming/app/membership"
	user "GolangLivePT01/golang_web_programming/app/user"
	docs "GolangLivePT01/golang_web_programming/cmd/docs"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"os"
	"strings"
)

type Controllers struct {
	MembershipController membership.Controller
	LogoController       logo.Controller
	UserController       user.Controller
}

type Middlewares struct {
	Middleware user.Middleware
}

func InitializeRoutes(e *echo.Group) {

	c := initController()
	userMiddleware := initUserMiddleware().Middleware

	// Host 정보 동적으로 수정
	if os.Getenv("PHASE") == "prod" {
		docs.SwaggerInfo.Host = "petstore.swagger.io"
	}

	memberships := e.Group("/memberships")
	logo := e.Group("/logo")

	memberships.Use(middleware.BodyDump(func(c echo.Context, reqBody []byte, resBody []byte) {
		uri := c.Request().RequestURI
		method := c.Request().Method
		status := http.StatusText(c.Response().Status)
		reqStr := strings.Trim(string(reqBody), "\n")
		resStr := strings.Trim(string(resBody), "\n")
		c.Logger().Output().Write([]byte(fmt.Sprintf("URI:[%s], Method:[%s], Status:[%s]\n"+
			"RequestBody:[%s]\nResponseBody:[%s]\n\n", uri, method, status, reqStr, resStr)))
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "URI:[${uri}], Method:[${method}], StatusCode:[${status}]\n",
	}))

	jwtMiddleware := middleware.JWTWithConfig(middleware.JWTConfig{Claims: &user.Claims{}, SigningKey: user.DefaultSecret})

	membershipCtl := c.MembershipController
	logoCtl := c.LogoController
	userCtl := c.UserController

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Static("/", "assets")
	memberships.POST("", membershipCtl.Create)
	memberships.GET("", membershipCtl.ReadAll, jwtMiddleware, userMiddleware.ValidateAdmin)
	memberships.GET("/:id", membershipCtl.Read, jwtMiddleware, userMiddleware.ValidateMemberOrAdmin)
	memberships.PUT("/:id", membershipCtl.Update, jwtMiddleware, userMiddleware.ValidateMember)
	memberships.DELETE("/:id", membershipCtl.Delete, jwtMiddleware, userMiddleware.ValidateMember)

	logo.GET("", logoCtl.Get)

	e.POST("/login", userCtl.Login)
}

func initController() *Controllers {
	data := map[string]membership.Membership{}
	service := membership.NewService(*membership.NewRepository(data))
	controller := membership.NewController(*service)

	return &Controllers{
		MembershipController: *controller,
		LogoController:       *logo.NewController(),
		UserController:       *user.NewController(*user.NewService(user.DefaultSecret, *membership.NewRepository(data))),
	}
}

func initUserMiddleware() *Middlewares {
	data := map[string]membership.Membership{}
	return &Middlewares{
		Middleware: *user.NewMiddleware(*membership.NewRepository(data)),
	}
}
