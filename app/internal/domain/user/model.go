package user

type User struct {
	ID       int32  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func SetUser(
	username string,
	password string,
	email string,
) User {
	return User{
		Username: username,
		Password: password,
		Email:    email,
	}
}
