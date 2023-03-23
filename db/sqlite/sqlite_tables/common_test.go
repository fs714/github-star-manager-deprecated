package sqlite_tables

import (
	"fmt"
	"strings"
	"testing"
)

type User struct {
	Id     string `db:"id"`
	Name   string `db:"name"`
	Email  string `db:"email"`
	Age    int    `db:"age"`
	Role   string `db:"role"`
	Points int    `db:"points"`
	TableBaseColumn
}

var UserFields []string = GetStructFields(User{}, []string{})

func TestGetStructFields(t *testing.T) {
	fields := GetStructFields(User{}, []string{})
	fmt.Println(strings.Join(fields, ", "))
}

func TestSqlGetAllColumn(t *testing.T) {
	sql := "SELECT " + GetSelectColumnsSql(UserFields) + " FROM user WHERE name = ?"
	fmt.Println(sql)
}

func TestSqlGetWithSkip(t *testing.T) {
	fields := GetStructFields(User{}, []string{"id", "created_at", "updated_at"})
	sql := "SELECT " + GetSelectColumnsSql(fields) + " FROM user WHERE name = ?"
	fmt.Println(sql)
}

func TestSqlInsert(t *testing.T) {
	sql := "INSERT INTO user " + GetInsertColumnsSql(UserFields) + " VALUES " + GetInsertNamedValuesSql(UserFields)
	fmt.Println(sql)
}

func TestSqlUpdate(t *testing.T) {
	sql := "UPDATE user SET " + GetUpdateSetSql(UserFields) + " WHERE id = :id"
	fmt.Println(sql)
}
