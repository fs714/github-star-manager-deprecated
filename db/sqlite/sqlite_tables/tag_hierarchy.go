package sqlite_tables

import (
	"fmt"

	"github.com/fs714/github-star-manager/db/sqlite"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
)

type TagHierarchy struct {
	Id       string  `db:"id"`
	User     string  `db:"user"`
	ParentID string  `db:"parent_id"`
	ChildID  string  `db:"child_id"`
	Order    float32 `db:"order"`
	TableBaseColumn
}

const TagHierarchyTableName = "tag_hierarchy"
const SqlCreateTagHierarchyTable = `
CREATE TABLE IF NOT EXISTS tag_hierarchy (
	id				TEXT		PRIMARY KEY,
	user			TEXT		NOT NULL,
	parent_id		TEXT		NOT NULL,
	child_id		TEXT		NOT NULL,
	order			REAL		NOT NULL,
	created_at		INTEGER		NOT NULL,
	updated_at		INTEGER		NOT NULL
)
`

var SqlIndexesTagHierarchyTable = []string{
	fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s_user_parent_idx ON %s (user, parent_id)", TagHierarchyTableName, TagHierarchyTableName),
	fmt.Sprintf("CREATE INDEX IF NOT EXISTS %s_user_child_idx ON %s (user, child_id)", TagHierarchyTableName, TagHierarchyTableName),
}

var TagHierarchyFields []string = GetStructFields(Tag{}, []string{})

func CreateTagHierarchyTable() (err error) {
	sql := fmt.Sprintf("DROP TABLE IF EXISTS %s", TagHierarchyTableName)
	_, err = sqlite.DBSqlx.Exec(sql)
	if err != nil {
		return errors.Wrapf(err, "failed to drop table %s", TagHierarchyTableName)
	}

	_, err = sqlite.DBSqlx.Exec(SqlCreateTagTable)
	if err != nil {
		return errors.Wrapf(err, "failed to create table %s", TagHierarchyTableName)
	}

	if len(SqlIndexesTagTable) > 0 {
		for _, sql := range SqlIndexesTagTable {
			_, err = sqlite.DBSqlx.Exec(sql)
			if err != nil {
				return errors.Wrapf(err, "failed to create index for table %s", TagHierarchyTableName)
			}
		}
	}

	return
}

func GetAllTagHierarchys() (tagHierarchys []TagHierarchy, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s", GetSelectColumnsSql(TagHierarchyFields), TagHierarchyTableName)
	err = sqlite.DBSqlx.Select(&tagHierarchys, sql)
	if err != nil {
		err = errors.Wrap(err, "failed to get all tag_hierarchys")
		return
	}

	return
}

func GetTagHierarchyById(id string) (tagHierarchy TagHierarchy, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", GetSelectColumnsSql(TagHierarchyFields), TagHierarchyTableName)
	err = sqlite.DBSqlx.Get(&tagHierarchy, sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to get tag_hierarchy by id")
		return
	}

	return
}

func GetTagHierarchysByUser(user string) (tagHierarchys []TagHierarchy, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE user = ?", GetSelectColumnsSql(TagHierarchyFields), TagHierarchyTableName)
	err = sqlite.DBSqlx.Select(&tagHierarchys, sql, user)
	if err != nil {
		err = errors.Wrap(err, "failed to get tag_hierarchys by user")
		return
	}

	return
}

func GetTagHierarchysByUserAndParentID(user string, parentID string) (tagHierarchys []TagHierarchy, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE user = ? and parent_id = ?",
		GetSelectColumnsSql(TagHierarchyFields), TagHierarchyTableName)
	err = sqlite.DBSqlx.Select(&tagHierarchys, sql, user, parentID)
	if err != nil {
		err = errors.Wrap(err, "failed to get tag_hierarchys by user and parent_id")
		return
	}

	return
}

func GetTagHierarchysByUserAndChildID(user string, childID string) (tagHierarchys []TagHierarchy, err error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE user = ? and child_id = ?",
		GetSelectColumnsSql(TagHierarchyFields), TagHierarchyTableName)
	err = sqlite.DBSqlx.Select(&tagHierarchys, sql, user, childID)
	if err != nil {
		err = errors.Wrap(err, "failed to get tag_hierarchys by user and child_id")
		return
	}

	return
}

func InsertTagHierarchy(tagHierarchy TagHierarchy) (err error) {
	if tagHierarchy.Id == "" {
		tagHierarchy.Id = ulid.Make().String()
	}

	sql := fmt.Sprintf("INSERT INTO %s %s VALUES %s",
		TagHierarchyTableName, GetInsertColumnsSql(TagHierarchyFields), GetInsertNamedValuesSql(TagHierarchyFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, tagHierarchy)
	if err != nil {
		err = errors.Wrap(err, "failed to insert tag_hierarchy")
		return
	}

	return
}

func UpdateTagHierarchy(tagHierarchy TagHierarchy) (err error) {
	sql := fmt.Sprintf("UPDATE %s SET %s WHERE id = :id", TagHierarchyTableName, GetUpdateSetSql(TagHierarchyFields))
	_, err = sqlite.DBSqlx.NamedExec(sql, tagHierarchy)
	if err != nil {
		err = errors.Wrap(err, "failed to update tag_hierarchy")
		return
	}

	return
}

func DeleteTagHierarchyById(id string) (err error) {
	sql := fmt.Sprintf("DELETE FROM %s WHERE id = ?", TagHierarchyTableName)
	_, err = sqlite.DBSqlx.Exec(sql, id)
	if err != nil {
		err = errors.Wrap(err, "failed to delete tag_hierarchy by id")
		return
	}

	return
}
