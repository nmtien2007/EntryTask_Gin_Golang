package login

type UserLoginSchema struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
