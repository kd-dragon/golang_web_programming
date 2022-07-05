package membership

type CreateRequest struct {
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type" example:"toss"`
}

type CreateResponse struct {
	Code           int    `json:"-"`
	Message        string `json:"message" example:"Created"`
	ID             string `json:"id" example:"354660dc-f798-11ec-b939-0242ac120002"`
	MembershipType string `json:"membership_type" example:"toss"`
}

type UpdateRequest struct {
	ID             string `param:"id" example:"354660dc-f798-11ec-b939-0242ac120002"`
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type" example:"toss"`
}

type UpdateResponse struct {
	Code       int        `json:"-"`
	Message    string     `json:"message" example:"Updated"`
	Membership Membership `json:membership,omitempty`
}

type ReadResponse struct {
	Code       int        `json:"-"`
	Message    string     `json:"message" example:"Read"`
	Membership Membership `json:membership,omitempty`
}

type ReadAllResponse struct {
	Code        int          `json:"-"`
	Message     string       `json:"message" example:"ReadAll"`
	Memberships []Membership `json:"memberships,omitempty"`
}

type DeleteResponse struct {
	Code    int    `json:"-"`
	Message string `json:"message" example:"Deleted"`
}

type GetResponse struct {
	Code           int    `json:"-"`
	Message        string `json:"message" example:"OK"`
	ID             string `json:"id,omitempty" example:"354660dc-f798-11ec-b939-0242ac120002"`
	UserName       string `json:"user_name" example:"andy"`
	MembershipType string `json:"membership_type" example:"toss"`
}

type Fail400GetResponse struct {
	Message string `json:"message" example:"Bad Request"`
}
