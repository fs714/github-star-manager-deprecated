package sqlite_tables

import (
	"fmt"

	"github.com/fs714/github-star-manager/db/sqlite"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

type RepoTag struct {
	Id     string `db:"id"`
	User   string `db:"user"`
	RepoId string `db:"repo_id"`
	TagId  string `db:"tag_id"`
	TableBaseColumn
}

const RepoTagTableName = "repo_tag"
const SqlCreateRepoTagTable = `
CREATE TABLE IF NOT EXISTS repo_tag (
	id				TEXT		PRIMARY KEY,
	user			TEXT		NOT NULL,
	repo_id			TEXT    	NOT NULL,
	tag_id			TEXT    	NOT NULL,
	created_at		INTEGER		NOT NULL,
	updated_at		INTEGER		NOT NULL
)
`

var SqlIndexesRepoTagTable = []string{
	fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s_user_repo_idx ON %s (user, repo_id)", RepoTagTableName, RepoTagTableName),
	fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s_user_tag_idx ON %s (user, tag_id)", RepoTagTableName, RepoTagTableName),
}

var RepoTagFields []string = GetStructFields(RepoTag{}, []string{})

func CreateRepoTagTable() (err error) {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s", RepoTagTableName)
	_, err = sqlite.DBSqlx.Exec(sql)
	if err != nil {
		return errors.Wrapf(err, "failed to drop table %s", RepoTagTableName)
	}

	_, err = sqlite.DBSqlx.Exec(SqlCreateRepoTagTable)
	if err != nil {
		return errors.Wrapf(err, "failed to create table %s", RepoTagTableName)
	}

	if len(SqlIndexesRepoTagTable) > 0 {
		for _, sql := range SqlIndexesRepoTagTable {
			_, err = sqlite.DBSqlx.Exec(sql)
			if err != nil {
				return errors.Wrapf(err, "failed to create index for table %s", RepoTagTableName)
			}
		}
	}

	return
}

func GetAllRepoTags() (repoTags []RepoTag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s", GetSelectColumnsSql(RepoTagFields), RepoTagTableName)
	err = sqlite.DBSqlx.Select(&repoTags, sql)
	if err != nil {
		err = errors.Wrap(err, "failed to get all repo_tags")
		return
	}

	return
}

func GetRepoTagById(id string) (repoTag RepoTag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", GetSelectColumnsSql(RepoTagFields), RepoTagTableName)
	err = sqlite.DBSqlx.Get(&repoTag, sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to get repo_tag by id")
		return
	}

	return
}

func GetRepoTagsByUser(user string) (repoTags []RepoTag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE user = ?", GetSelectColumnsSql(RepoTagFields), RepoTagTableName)
	err = sqlite.DBSqlx.Select(&repoTags, sql, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get repo_tags by user")
		return
	}

	return
}

func InsertRepoTag(repoTag RepoTag) (err error) {
	if repoTag.Id == "" {
		repoTag.Id = ulid.Make().String()
	}

	sql := fmt.Sprintf("INSERT INTO %s %s VALUES %s",
		RepoTagTableName, GetInsertColumnsSql(RepoTagFields), GetInsertNamedValuesSql(RepoTagFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, repoTag)
	if err != nil {
		err = errors.Wrap(err, "failed to insert repo_tag")
		return
	}

	return
}

func UpdateRepoTag(repoTag RepoTag) (err error) {
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", RepoTagTableName, GetUpdateSetSql(RepoTagFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, repoTag)
	if err != nil {
		err = errors.Wrap(err, "failed to update repo_tag")
		return
	}

	return
}

func DeleteRepoTagById(id string) (err error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?", RepoTagTableName)
	_, err = sqlite.DBSqlx.Exec(sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to delete repo_tag by id")
		return
	}

	return
}
