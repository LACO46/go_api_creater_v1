package api

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func ConnectDb(s *DbSettingModel) (*sqlx.DB, error) {
	return sqlx.Connect(
		"postgres",
		fmt.Sprintf(
			"user=%s password=%s dbname=%s sslmode=%s host=%s",
			s.User,
			s.Password,
			s.Dbname,
			s.Sslmode,
			s.Host,
		),
	)
}
