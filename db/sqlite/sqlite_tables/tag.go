package sqlite_tables

import (
	"fmt"

	"github.com/fs714/github-star-manager/db/sqlite"
	"github.com/pkg/errors"
)

type Tag struct {
	Id          string `db:"id"`
	User        string `db:"user"`
	Name        string `db:"name"`
	Description string `db:"description"`
	IsRoot      bool   `json:"is_root"`
	Level       int    `json:"level"`
	FullPath    string `json:"full_path"`
	TableBaseColumn
}

const TagTableName = "tag"
const SqlCreateTagTable = `
CREATE TABLE IF NOT EXISTS tag (
	id				TEXT		PRIMARY KEY,
	user			TEXT		NOT NULL,
	name			TEXT    	NOT NULL,
	description		TEXT		NOT NULL,
	is_root			BOOLEAN		NOT NULL,
	level			INTEGER		NOT NULL,
	full_path		TEXT		NOT NULL,
	created_at		INTEGER		NOT NULL,
	updated_at		INTEGER		NOT NULL
)
`

var SqlIndexesTagTable = []string{
	fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s_user_name_idx ON %s (user, name)", TagTableName, TagTableName),
}

var TagFields []string = GetStructFields(Tag{}, []string{})

func CreateTagTable() (err error) {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s", TagTableName)
	_, err = sqlite.DBSqlx.Exec(sql)
	if err != nil {
		return errors.Wrapf(err, "failed to drop table %s", TagTableName)
	}

	_, err = sqlite.DBSqlx.Exec(SqlCreateTagTable)
	if err != nil {
		return errors.Wrapf(err, "failed to create table %s", TagTableName)
	}

	if len(SqlIndexesTagTable) > 0 {
		for _, sql := range SqlIndexesTagTable {
			_, err = sqlite.DBSqlx.Exec(sql)
			if err != nil {
				return errors.Wrapf(err, "failed to create index for table %s", TagTableName)
			}
		}
	}

	return
}

func GetAllTags() (tags []Tag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s", GetSelectColumnsSql(TagFields), TagTableName)
	err = sqlite.DBSqlx.Select(&tags, sql)
	if err != nil {
		err = errors.Wrap(err, "failed to get all tags")
		return
	}

	return
}

func GetTagById(id string) (tag Tag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", GetSelectColumnsSql(TagFields), TagTableName)
	err = sqlite.DBSqlx.Get(&tag, sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to get tag by id")
		return
	}

	return
}

func GetTagsByName(name string) (tags []Tag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE name = ?", GetSelectColumnsSql(TagFields), TagTableName)
	err = sqlite.DBSqlx.Select(&tags, sql, name)
	if err != nil {
		err = errors.Wrap(err, "failed to get tags by name")
		return
	}

	return
}

func GetTagsByUser(user string) (tags []Tag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE user = ?", GetSelectColumnsSql(TagFields), TagTableName)
	err = sqlite.DBSqlx.Select(&tags, sql, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get tags by user")
		return
	}

	return
}

func GetTagByNameAndUser(name string, user string) (tag Tag, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE name = ? and user = ?", GetSelectColumnsSql(TagFields), TagTableName)
	err = sqlite.DBSqlx.Get(&tag, sql, name, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get tag by name and user")
		return
	}

	return
}

func InsertTag(tag Tag) (err error) {
	sql := fmt.Sprintf("INSERT INTO %s %s VALUES %s",
		TagTableName, GetInsertColumnsSql(TagFields), GetInsertNamedValuesSql(TagFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, tag)
	if err != nil {
		err = errors.Wrap(err, "failed to insert tag")
		return
	}

	return
}

func UpdateTag(tag Tag) (err error) {
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", TagTableName, GetUpdateSetSql(TagFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, tag)
	if err != nil {
		err = errors.Wrap(err, "failed to update tag")
		return
	}

	return
}

func DeleteTagById(id string) (err error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TagTableName)
	_, err = sqlite.DBSqlx.Exec(sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to delete tag by id")
		return
	}

	return
}
