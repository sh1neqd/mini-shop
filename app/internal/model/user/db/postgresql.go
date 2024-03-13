package user

import (
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"testAssignment/internal/model/user"
	"testAssignment/pkg/client/postgresql"
	"testAssignment/pkg/logging"
)

type repository struct {
	client postgresql.Client
	logger *logging.Logger
}

func (r repository) Create(ctx context.Context, user *user.User) error {
	q := `INSERT INTO public.user (username, password, email) VALUES ($1, $2, $3) RETURNING id`
	r.logger.Tracef("SQL query:%s", q)
	if err := r.client.QueryRow(ctx, q, user.Username, user.Password, user.Email).Scan(&user.ID); err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			er := fmt.Errorf(fmt.Sprintf("sql error %s, details: %s, where: %s",
				pgErr.Message, pgErr.Detail, pgErr.Where))
			r.logger.Error(er)
			return er
		}
		return err
	}

	return nil

}

func (r repository) FindAll(ctx context.Context) (u []user.User, err error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) FindOne(ctx context.Context, id string) (user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r repository) Update(ctx context.Context, user user.User) error {
	//TODO implement me
	panic("implement me")
}

func (r repository) Delete(ctx context.Context, id string) error {
	//TODO implement me
	panic("implement me")
}

func NewRepository(client postgresql.Client, logger *logging.Logger) user.Repository {
	return &repository{
		client: client,
		logger: logger,
	}
}
