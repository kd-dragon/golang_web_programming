package membership

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Controller struct {
	service Service
}

func NewController(service Service) *Controller {
	return &Controller{service: service}
}

// Read godoc
// @Summary      멤버십 정보 단건 조회
// @Description  멤버십 정보를 조회합니다. (상세 설명)
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        id             path      string  true  "Membership uuid"
// @Success      200            {object}  ReadResponse
// @Router       /api/v2/memberships/{id} [get]
func (c Controller) Read(ctx echo.Context) error {
	id := ctx.Param("id")
	res, _ := c.service.Read(id)

	return ctx.JSON(res.Code, res)
}

// ReadAll godoc
// @Summary      멤버십 정보 목록 조회
// @Description  멤버십 정보 전체 목록 조회합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        id             query      string  false  "offset"
// @Param        id             query      string  false  "limit"
// @Success      200            {object}  ReadAllResponse
// @Router       /api/v2/memberships [get]
func (c Controller) ReadAll(ctx echo.Context) error {
	offset := ctx.QueryParam("offset")
	limit := ctx.QueryParam("limit")

	res, _ := c.service.ReadAll(offset, limit)

	return ctx.JSON(res.Code, res)
}

// Create godoc
// @Summary      멤버십 생성
// @Description  멤버십을 생성합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @Param        requestBody  body      CreateRequest  true  "UserName:사용자의 이름, MembershipType:naver,toss,pacyco 중 하나"
// @Success      200          {object}  CreateResponse
// @Router       /api/v2/memberships [post]
func (c Controller) Create(ctx echo.Context) error {
	var req CreateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, CreateResponse{Message: "invalid_request_format"})
	}
	res, _ := c.service.Create(req)

	return ctx.JSON(res.Code, res)
}

// Update godoc
// @Summary      멤버십 수정
// @Description  멤버십을 수정합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        requestBody  body      UpdateRequest  true  "ID: 아이디, UserName:사용자의 이름, MembershipType:naver,toss,pacyco 중 하나"
// @Success      200          {object}  UpdateResponse
// @Router       /api/v2/memberships [put]
func (c Controller) Update(ctx echo.Context) error {
	var req UpdateRequest
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, UpdateResponse{Message: "invalid_request_format"})
	}
	res, _ := c.service.Update(req)
	return ctx.JSON(res.Code, res)
}

// Delete godoc
// @Summary      멤버십 삭제
// @Description  멤버십을 삭제합니다.
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        id  		   path      string true  "Membership uuid"
// @Success      200          {object}  DeleteResponse
// @Router       /api/v2/memberships [delete]
func (c Controller) Delete(ctx echo.Context) error {
	id := ctx.Param("id")
	res, _ := c.service.Delete(id)
	return ctx.JSON(res.Code, res)
}

// GetByID godoc
// @Summary      멤버십 정보 단건 조회
// @Description  멤버십 정보를 조회합니다. (상세 설명)
// @Accept       json
// @Tags         Memberships
// @Produce      json
// @param        Authorization  header    string  true  "Authorization"  default(Bearer <Add access token here>)
// @Param        id             path      string  true  "Membership uuid"
// @Success      200            {object}  GetResponse
// @Failure      400            {object}  Fail400GetResponse
// @Router       /api/v2/memberships/{id} [get]
func (controller Controller) GetByID(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}

func (controller Controller) GetAll(c echo.Context) error {
	return c.JSON(http.StatusOK, "hello world")
}
