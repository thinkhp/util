package database

import (
	"database/sql"
	"reflect"
	"strconv"
	"util/think"
	"util/thinkLog"
	"util/thinkString"
)

// 文件描述:
// 数据库 mysql 的增删改查
// 根据参数创建完整的 SQL 语句, 并提交

func UpdateStruct(tx *sql.Tx, tableName string, notNilMap map[string]string, primaryKeyName string, primaryKey int) int64 {
	sqlString := "UPDATE " + tableName + " SET "
	for k, v := range notNilMap {
		sqlString += k + "=" + v + ","
	}
	thinkString.ReplaceLastRune(&sqlString, ' ')
	sqlString += "WHERE " + primaryKeyName + "=" + strconv.Itoa(primaryKey)
	if primaryKey == 0 {
		panic("primaryKey not nil")
	}
	thinkLog.DebugLog.Println(sqlString)

	var result sql.Result
	var err error
	if tx == nil {
		result, err = Idb.Exec(sqlString)
	} else {
		result, err = tx.Exec(sqlString)
	}
	think.IsNil(err)
	affect, err := result.RowsAffected()
	think.IsNil(err)

	return affect
}

func DeleteStruct(tx *sql.Tx, tableName string, primaryKeyName string, primaryKey int) int64 {
	sqlString := "DELETE FROM " + tableName + " WHERE " + primaryKeyName + "=" + strconv.Itoa(primaryKey)
	thinkLog.DebugLog.Println(sqlString)

	var result sql.Result
	var err error
	if tx == nil {
		result, err = Idb.Exec(sqlString)
	} else {
		result, err = tx.Exec(sqlString)
	}
	think.IsNil(err)
	affect, err := result.RowsAffected()
	think.IsNil(err)

	return affect
}

// 存在唯一索引的插入,不会panic
// INSERT IGNORE INTO table_name (col1,col2,col3...) VALUES (v1,v2,v3...),(v1,v2,v3...),(v1,v2,v3...)
func InsertStructUnique(tx *sql.Tx, tableName string, notNilMap map[string]string) int64 {
	sqlString := "INSERT IGNORE INTO " + tableName + " ("
	var value = " VALUES ("
	for k, v := range notNilMap {
		sqlString += k + ","
		value += v + ","
	}
	thinkString.ReplaceLastRune(&sqlString, ')')
	thinkString.ReplaceLastRune(&value, ')')
	thinkLog.DebugLog.Println(sqlString + value)

	var result sql.Result
	var err error
	if tx == nil {
		result, err = Idb.Exec(sqlString + value)
	} else {
		result, err = tx.Exec(sqlString + value)
	}
	think.IsNil(err)
	last, err := result.LastInsertId()
	think.IsNil(err)

	return last
}

func InsertStruct(tx *sql.Tx, tableName string, notNilMap map[string]string) int64 {
	sqlString := "INSERT INTO " + tableName + " ("
	var value = " VALUES ("
	for k, v := range notNilMap {
		sqlString += k + ","
		value += v + ","
	}
	thinkString.ReplaceLastRune(&sqlString, ')')
	thinkString.ReplaceLastRune(&value, ')')
	thinkLog.DebugLog.Println(sqlString + value)

	var result sql.Result
	var err error
	if tx == nil {
		result, err = Idb.Exec(sqlString + value)
	} else {
		result, err = tx.Exec(sqlString + value)
	}
	think.IsNil(err)
	last, err := result.LastInsertId()
	think.IsNil(err)

	return last
}

// 批量插入,insert ignore into
// 因为插入语句中的 insert (col...) 中 col 是统一的,所以在插入结构体时,必须保证属性中值的存在具有一致性
func InsertBatchStruct(tx *sql.Tx, tableName string, notNilMapList *[]map[string]string) int64 {
	keyList := make([]string, 0)
	sqlString := "INSERT IGNORE INTO " + tableName + " ("
	for k, _ := range (*notNilMapList)[0] {
		sqlString += k + ","
		keyList = append(keyList, k)
	}
	thinkString.ReplaceLastRune(&sqlString, ')')
	sqlString += " VALUES"
	var value = ""
	for _, notNilMap := range *notNilMapList {
		value += " ("
		// 实现map有序取出
		for i := 0; i < len(keyList); i++ {
			value += notNilMap[keyList[i]] + ","
		}
		thinkString.ReplaceLastRune(&value, ')')
		value += ","
	}
	thinkString.ReplaceLastRune(&value, ' ')
	thinkLog.DebugLog.Println(sqlString + value)

	var result sql.Result
	var err error
	if tx == nil {
		result, err = Idb.Exec(sqlString + value)
	} else {
		result, err = tx.Exec(sqlString + value)
	}
	think.IsNil(err)
	affect, err := result.RowsAffected()
	think.IsNil(err)

	return affect
}

// !!!重要提示:
// *sql.Rows 未关闭
func Select(tx *sql.Tx, sqlString string, args ...interface{}) *sql.Rows {
	thinkLog.DebugLog.PrintSQL(sqlString, args)
	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = Idb.Query(sqlString, args...)
	} else {
		rows, err = tx.Query(sqlString, args...)
	}
	think.IsNil(err)

	return rows
}

func SelectStruct(tx *sql.Tx, ptr interface{}, sqlString string, args ...interface{}) {
	value := reflect.ValueOf(ptr).Elem()
	sliceType := value.Type() // []pkg.type
	if sliceType.Kind() != reflect.Slice {
		thinkLog.ErrorLog.Println("not a slice")
		return
	}
	//fmt.Println(sliceType)
	IType := sliceType.Elem()
	//fmt.Println(IType)
	//thinkLog.DebugLog.PrintSQL(sqlString, args)
	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = Idb.Query(sqlString, args...)
	} else {
		rows, err = tx.Query(sqlString, args...)
	}
	think.IsNil(err)
	defer rows.Close()

	length := 0
	s := reflect.MakeSlice(sliceType, length, length)
	//
	//s := reflect.Zero(reflect.SliceOf(IType))
	for rows.Next() {
		length++
		temp := reflect.New(IType).Elem()
		tempSlice := make([]interface{}, temp.NumField())
		for i := 0; i < temp.NumField(); i++ {
			tempSlice[i] = temp.Field(i).Addr().Interface()
		}
		//fmt.Println(reflect.TypeOf(tempSlice[0]))
		rows.Scan(tempSlice...)
		s = reflect.Append(s, temp)
	}

	value.Set(reflect.ValueOf(s.Interface()))
}

// fields,表字段
//
//func SelectStruct(tableName string, fields []string, queryMap map[string]string) string {
//	sqlString := "SELECT "
//	if fields == nil || len(fields) == 0 {
//		sqlString += "* "
//	} else {
//		for _, field := range fields {
//			sqlString += field + ","
//		}
//		thinkString.ReplaceLastRune(&sqlString, ' ')
//	}
//	sqlString += "FROM " + tableName + " WHERE "
//	if queryMap == nil || len(queryMap) == 0 {
//		sqlString += "1=1"
//	} else {
//		for k, v := range queryMap {
//			sqlString += k + "=" + v + ","
//		}
//		thinkString.ReplaceLastRune(&sqlString, ' ')
//	}
//	thinkLog.DebugLog.Println(sqlString)
//
//	return sqlString
//}
