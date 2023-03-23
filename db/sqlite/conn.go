package sqlite

import (
	"strings"
	"time"

	"github.com/fs714/github-star-manager/pkg/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pkg/errors"
)

var DBSqlx *sqlx.DB

func InitSqlxFromConfig() (err error) {
	err = InitSqlx(config.Config.Database.Path)
	if err != nil {
		errors.WithMessage(err, "failed to init db from config")
		return
	}

	return
}

func InitSqlx(path string) (err error) {
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	DBSqlx, err = sqlx.Open("sqlite3", "file:"+path+"github_star_manager.db?_busy_timeout=10000&_journal_mode=WAL&_synchronous=NORMAL&cache=shared")
	if err != nil {
		err = errors.Wrap(err, "failed to connect to DB")
		return
	}

	DBSqlx.DB.SetMaxOpenConns(2)
	DBSqlx.DB.SetMaxIdleConns(1)
	DBSqlx.DB.SetConnMaxLifetime(time.Duration(21600) * time.Second)

	err = DBSqlx.Ping()
	if err != nil {
		err = errors.Wrap(err, "failed to ping DB")
		return
	}

	return
}
