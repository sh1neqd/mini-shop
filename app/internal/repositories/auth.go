package repositories

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"testAssignment/internal/domain/user"
)

type Auth struct {
	db *sqlx.DB
}

func NewAuth(db *sqlx.DB) *Auth {
	return &Auth{db: db}
}

func (r *Auth) CreateUser(dto user.CreateUserDTO) (int, error) {
	var id int
	q := `INSERT INTO public.user (username, password, email) VALUES ($1, $2, $3) RETURNING id`

	row := r.db.QueryRow(q, dto.Username, dto.Password, dto.Email)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	logrus.Println(q)
	return id, nil
}

func (r *Auth) GetById(id int) (user.User, error) {
	var u user.User
	q := `SELECT id, username, password, email FROM public.user WHERE id = $1`
	row := r.db.QueryRow(q, id)
	err := row.Scan(&u)
	if err != nil {
		return user.User{}, err
	}
	return u, err
}

func (r *Auth) GetUser(username, password string) (user.User, error) {
	var u user.User
	q := `SELECT id FROM public.user WHERE username=$1`
	err := r.db.Get(&u, q, username)
	logrus.Println(u)
	if err != nil {
		fmt.Printf("failed to get user, err: %v", err)
		return user.User{}, err
	}
	logrus.Println(q)
	return u, err
}

func (r *Auth) PasswordsPass(username, password string) bool {
	q := `SELECT password FROM public.user WHERE username=$1`
	var qpass string
	row := r.db.QueryRow(q, username)
	err := row.Scan(&qpass)
	if err != nil {
		logrus.Errorf("error:%v", err)
	}
	if qpass != password {
		return false
	}
	return true
}

func (r *Auth) UserExist(username string) bool {
	var usern string
	q := `SELECT username FROM public.user WHERE username=$1`
	row := r.db.QueryRow(q, username)
	err := row.Scan(&usern)
	if err != nil {
		logrus.Errorf("error:%v", err)
	}
	if usern == "" {
		return false
	}

	return true

}
