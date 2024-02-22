package requestmodel

type RegisterReqestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
