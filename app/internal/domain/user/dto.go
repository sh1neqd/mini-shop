package user

type CreateUserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthorizeUserDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
