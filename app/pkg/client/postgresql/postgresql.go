package postgresql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"testAssignment/internal/config"
)

func NewPostgresDB(config *config.Config) (*sqlx.DB, error) {
	cfg := config.PostgreSQL
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		cfg.Host, cfg.Username, cfg.Password, cfg.Database, cfg.Port, cfg.SSLMode))
	if err != nil {
		logrus.Errorf("failed to connect to db, err: %s", err.Error())
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logrus.Errorf("failed to ping db, err: %v", err)
		return nil, err
	}

	return db, nil
}
