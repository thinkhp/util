package database

import (
	"database/sql"
	"util/think"
	"util/thinkLog"
)

// 表字段的详细信息
type TableFiled struct {
	// 字段名
	Field string
	// 字段sql类型
	Type string
	// 字符集
	Collation string
	// 是否允许为空
	Null string
	// 索引类型
	Key string
	// 默认值
	Default string
	// 自增
	Extra string
	// CRUD
	Privileges string
	// 备注
	Comment string
}

// 获取全部数据库
func GetDatabases() []string {
	sqlString := "SHOW DATABASES"
	_, rows := SelectList(nil, sqlString)

	databases := make([]string, 0)
	for i := 0; i < len(rows); i++ {
		databases = append(databases, rows[i][0])
	}
	return databases
}

// 获取数据库下全部数据表
func GetTables(databaseName string) []string {
	sqlString := "SELECT distinct TABLE_NAME FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_SCHEMA=?"
	_, rows := SelectList(nil, sqlString, databaseName)

	tables := make([]string, 0)
	for i := 0; i < len(rows); i++ {
		tables = append(tables, rows[i][0])
	}
	return tables
}

// 获取数据表下全部字段
func GetFields(databaseName, tableName string) []TableFiled {
	sqlString := ""
	if databaseName == "" {
		sqlString = "SHOW FULL COLUMNS FROM " + tableName
	} else {
		sqlString = "SHOW FULL COLUMNS FROM " + databaseName + "." + tableName
	}
	rows := Select(nil, sqlString)
	defer rows.Close()

	fields := make([]TableFiled, 0)
	var Field sql.NullString
	var Type sql.NullString
	var Collation sql.NullString
	var Null sql.NullString
	var Key sql.NullString
	var Default sql.NullString
	var Extra sql.NullString
	var Privileges sql.NullString
	var Comment sql.NullString
	for rows.Next() {
		err := rows.Scan(&Field, &Type, &Collation, &Null, &Key, &Default, &Extra, &Privileges, &Comment)
		think.IsNil(err)
		filed := TableFiled{Field.String, Type.String, Collation.String, Null.String, Key.String, Default.String,
			Extra.String, Privileges.String, Comment.String}
		fields = append(fields, filed)
	}

	return fields
}

func Exec(sqlString string) (sql.Result, error) {
	thinkLog.DebugLog.Println(sqlString)
	result, err := Idb.Exec(sqlString)

	return result, err
}

func TruncateTable(tableName string) {
	sqlString := "truncate table " + tableName

	result, err := Exec(sqlString)
	think.IsNil(err)
	_, err = result.RowsAffected()
	think.IsNil(err)
}

func GetDDL(tableName string) string {
	sqlString := "SHOW CREATE TABLE " + tableName
	_, rows := SelectMap(nil, sqlString)
	//fmt.Println(rows)
	return rows[0]["Create Table"]
}

func sqlStringRemark() {
	str := make([]string, 0)
	str = append(str, "SHOW GLOBAL VARIABLES")
	str = append(str, "SET GLOBAL max_allowed_packet = 4*1024*1024")
	str = append(str, "")
}
