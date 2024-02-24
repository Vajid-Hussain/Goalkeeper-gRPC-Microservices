package resposemodel

type UserData struct {
	ID       string `json:"id"`
	Email    string `json:"name"`
	Jwt      string `json:"jwt_token"`
	Password string
}
