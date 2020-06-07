package api

import (
	"github.com/jmoiron/sqlx"
)

type TestApiModel struct {
	Id   string `db:"id"`
	Name string `db:"name"`
}

func Test(db *sqlx.DB) []TestApiModel {
	var testModel []TestApiModel
	var query string = "select id, name from test_table"
	err := db.Select(&testModel, query)

	if err != nil {
		return nil
	}

	return testModel
}
