package dbrepo

import (
	"database/sql"

	"github.com/raymondchr/go-hotel-app/internal/config"
	"github.com/raymondchr/go-hotel-app/internal/repository"
)

type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
