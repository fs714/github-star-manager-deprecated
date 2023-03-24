package sqlite_tables

import (
	"fmt"

	"github.com/fs714/github-star-manager/db/sqlite"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

type Repo struct {
	Id          string `db:"id"`
	User        string `db:"user"`
	Name        string `db:"name"`
	Url         string `db:"url"`
	Language    string `db:"language"`
	CreatedAt   int64  `db:"created_at"`
	UpdatedAt   int64  `db:"updated_at"`
	PushedAt    int64  `db:"pushed_at"`
	StarsCount  int    `db:"stars_count"`
	ForksCount  int    `db:"forks_count"`
	Description string `db:"description"`
}

const RepoTableName = "repo"
const SqlCreateRepoTable = `
CREATE TABLE IF NOT EXISTS repo (
	id				TEXT		PRIMARY KEY,
	user			TEXT		NOT NULL,
	name			TEXT    	NOT NULL,
	url				TEXT    	NOT NULL,
	language		TEXT		NOT NULL,
	created_at		INTEGER		NOT NULL,
	updated_at		INTEGER		NOT NULL,
	pushed_at		INTEGER		NOT NULL,
	stars_count		INTEGER		NOT NULL,
	forks_count		INTEGER		NOT NULL,
	description		TEXT		NOT NULL
)
`

var SqlIndexesRepoTable = []string{
	fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s_user_name_idx ON %s (user, name)", RepoTableName, RepoTableName),
}

var RepoFields []string = GetStructFields(Repo{}, []string{})

func CreateRepoTable() (err error) {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s", RepoTableName)
	_, err = sqlite.DBSqlx.Exec(sql)
	if err != nil {
		return errors.Wrapf(err, "failed to drop table %s", RepoTableName)
	}

	_, err = sqlite.DBSqlx.Exec(SqlCreateRepoTable)
	if err != nil {
		return errors.Wrapf(err, "failed to create table %s", RepoTableName)
	}

	if len(SqlIndexesRepoTable) > 0 {
		for _, sql := range SqlIndexesRepoTable {
			_, err = sqlite.DBSqlx.Exec(sql)
			if err != nil {
				return errors.Wrapf(err, "failed to create index for table %s", RepoTableName)
			}
		}
	}

	return
}

func GetAllRepos() (repos []Repo, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s", GetSelectColumnsSql(RepoFields), RepoTableName)
	err = sqlite.DBSqlx.Select(&repos, sql)
	if err != nil {
		err = errors.Wrap(err, "failed to get all repos")
		return
	}

	return
}

func GetRepoById(id string) (repo Repo, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", GetSelectColumnsSql(RepoFields), RepoTableName)
	err = sqlite.DBSqlx.Get(&repo, sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to get repo by id")
		return
	}

	return
}

func GetReposByName(name string) (repos []Repo, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE name = ?", GetSelectColumnsSql(RepoFields), RepoTableName)
	err = sqlite.DBSqlx.Select(&repos, sql, name)
	if err != nil {
		err = errors.Wrap(err, "failed to get repos by name")
		return
	}

	return
}

func GetReposByUser(user string) (repos []Repo, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE user = ?", GetSelectColumnsSql(RepoFields), RepoTableName)
	err = sqlite.DBSqlx.Select(&repos, sql, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get repos by user")
		return
	}

	return
}

func GetRepoByNameAndUser(name string, user string) (repo Repo, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE name = ? and user = ?", GetSelectColumnsSql(RepoFields), RepoTableName)
	err = sqlite.DBSqlx.Get(&repo, sql, name, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get repo by name and user")
		return
	}

	return
}

func InsertRepo(repo Repo) (err error) {
	if repo.Id == "" {
		repo.Id = ulid.Make().String()
	}

	sql := fmt.Sprintf("INSERT INTO %s %s VALUES %s",
		RepoTableName, GetInsertColumnsSql(RepoFields), GetInsertNamedValuesSql(RepoFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, repo)
	if err != nil {
		err = errors.Wrap(err, "failed to insert repo")
		return
	}

	return
}

func UpdateRepo(repo Repo) (err error) {
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", RepoTableName, GetUpdateSetSql(RepoFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, repo)
	if err != nil {
		err = errors.Wrap(err, "failed to update repo")
		return
	}

	return
}

func DeleteRepoById(id string) (err error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?", RepoTableName)
	_, err = sqlite.DBSqlx.Exec(sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to delete repo by id")
		return
	}

	return
}
