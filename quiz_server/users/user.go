package users

type CreateUserArgs struct {
	Name           string `json:"name"`
	Password       string `json:"password"`
	PasswordVerify string `json:"verify"`
}
