//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/mysql"
)

var User = newUserTable("membership", "user", "")

type userTable struct {
	mysql.Table

	// Columns
	ID       mysql.ColumnInteger
	UserId   mysql.ColumnString
	Password mysql.ColumnString

	AllColumns     mysql.ColumnList
	MutableColumns mysql.ColumnList
}

type UserTable struct {
	userTable

	NEW userTable
}

// AS creates new UserTable with assigned alias
func (a UserTable) AS(alias string) *UserTable {
	return newUserTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new UserTable with assigned schema name
func (a UserTable) FromSchema(schemaName string) *UserTable {
	return newUserTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new UserTable with assigned table prefix
func (a UserTable) WithPrefix(prefix string) *UserTable {
	return newUserTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new UserTable with assigned table suffix
func (a UserTable) WithSuffix(suffix string) *UserTable {
	return newUserTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newUserTable(schemaName, tableName, alias string) *UserTable {
	return &UserTable{
		userTable: newUserTableImpl(schemaName, tableName, alias),
		NEW:       newUserTableImpl("", "new", ""),
	}
}

func newUserTableImpl(schemaName, tableName, alias string) userTable {
	var (
		IDColumn       = mysql.IntegerColumn("ID")
		UserIdColumn   = mysql.StringColumn("UserId")
		PasswordColumn = mysql.StringColumn("Password")
		allColumns     = mysql.ColumnList{IDColumn, UserIdColumn, PasswordColumn}
		mutableColumns = mysql.ColumnList{UserIdColumn, PasswordColumn}
	)

	return userTable{
		Table: mysql.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:       IDColumn,
		UserId:   UserIdColumn,
		Password: PasswordColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
