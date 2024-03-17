package user

//import (
//	"context"
//	"fmt"
//	"github.com/jackc/pgconn"
//	"golang.org/x/crypto/bcrypt"
//	"testAssignment/pkg/client/postgresql"
//	"testAssignment/pkg/logging"
//)
//
//type UserRepository struct {
//	client postgresql.Client
//	logger *logging.Logger
//}
//
//func (r UserRepository) Create(ctx context.Context, dto CreateUserDTO) error {
//	q := `INSERT INTO public.user (username, password, email) VALUES ($1, $2, $3) RETURNING id`
//	r.logger.Tracef("SQL query: %s", q)
//	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
//	if err != nil {
//		r.logger.Errorf("password hashing error: %v", err)
//	}
//	u := User{
//		Username: dto.Username,
//		Password: string(hashedPassword),
//		Email:    dto.Email,
//	}
//	if err := r.client.QueryRow(ctx, q, u.Username, u.Password, u.Email).Scan(&u.ID); err != nil {
//		if pgErr, ok := err.(*pgconn.PgError); ok {
//			er := fmt.Errorf(fmt.Sprintf("sql error %s, details: %s, where: %s",
//				pgErr.Message, pgErr.Detail, pgErr.Where))
//			r.logger.Error(er)
//			return er
//		}
//		return err
//	}
//	return nil
//}
//
//func (r UserRepository) FindAll(ctx context.Context) (u []User, err error) {
//	return nil, nil
//}
//
//func (r UserRepository) FindOneById(ctx context.Context, id int) (User, error) {
//	var u User
//	q := `SELECT id, username, email FROM public.user WHERE id = $1`
//	if err := r.client.QueryRow(ctx, q, id).Scan(&u.ID, &u.Username, &u.Email); err != nil {
//		if pgErr, ok := err.(*pgconn.PgError); ok {
//			er := fmt.Errorf(fmt.Sprintf("sql error %s, details: %s, where: %s",
//				pgErr.Message, pgErr.Detail, pgErr.Where))
//			r.logger.Error(er)
//			return u, err
//		}
//	}
//	return u, nil
//}
//
//func (r UserRepository) FindOneByUsername(ctx context.Context, username string) (User, error) {
//	var u User
//	q := `SELECT id, username, email FROM public.user WHERE username = $1`
//	if err := r.client.QueryRow(ctx, q, username).Scan(&u.ID, &u.Username, &u.Email); err != nil {
//		if pgErr, ok := err.(*pgconn.PgError); ok {
//			er := fmt.Errorf(fmt.Sprintf("sql error %s, details: %s, where: %s",
//				pgErr.Message, pgErr.Detail, pgErr.Where))
//			r.logger.Error(er)
//			return u, err
//		}
//	}
//	r.logger.Println(u) // for test
//	return u, nil
//}
//
//func (r UserRepository) CheckPassword(ctx context.Context, dto AuthorizeUserDto) (bool, error) {
//	var hashedPassword string
//	err := r.client.QueryRow(ctx, "SELECT password FROM public.user WHERE username = $1", dto.Username).Scan(&hashedPassword)
//	if err != nil {
//		return false, err
//	}
//
//	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(dto.Password))
//	if err != nil {
//		r.logger.Errorf("compare passwords error:%v", err)
//		return false, err
//	}
//
//	return true, nil
//}
//
//func NewRepository(client postgresql.Client, logger *logging.Logger) UserRepository {
//	return UserRepository{
//		client: client,
//		logger: logger,
//	}
//}
