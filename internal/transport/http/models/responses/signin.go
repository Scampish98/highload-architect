package responses

type SignInResponse struct {
	Token string `json:"token" binding:"required,jwt"`
}
