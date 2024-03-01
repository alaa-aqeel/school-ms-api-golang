package database

import (
	"database/sql"
	"fmt"

	"github.com/alaa-aqeel/school-ms-api-golang/config"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"
)

var Sql *goqu.Database

func dbConnect() *sql.DB {
	db, err := sql.Open(config.DB_DRIVER, fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s port=%s host=%s",
		config.Env("DB_USERNAME", config.DB_USERNAME),
		config.Env("DB_PASSWORD", config.DB_PASSWORD),
		config.Env("DB_NAME", config.DB_NAME),
		config.Env("DB_SSL", config.DB_SSLMODE),
		config.Env("DB_PORT", config.DB_PORT),
		config.Env("DB_HOST", config.DB_HOST),
	))
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InitSql() *goqu.Database {
	Sql = goqu.New(config.DB_DRIVER, dbConnect())
	return Sql
}
