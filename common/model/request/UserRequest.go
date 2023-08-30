package request

type UserLoginRequest struct {
	UserAccount  string `json:"userAccount"`
	UserPassword string `json:"userPassword"`
}
