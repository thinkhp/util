package database

import (
	"database/sql"
	"fmt"
	"util/think"
	"util/thinkLog"
)

// 文件描述:
// 数据库 mysql 的增删改查
// 使用 (SQL语句, 参数...) 来提交语句

func Delete(tx *sql.Tx, sqlString string, args ...interface{}) int64 {
	thinkLog.DebugLog.Println(SprintSQL(sqlString, args...))
	var result sql.Result
	var err error
	var affect int64
	if tx == nil {
		result, err = Idb.Exec(sqlString, args...)
	} else {
		result, err = tx.Exec(sqlString, args...)
	}
	think.IsNil(err)
	affect, err = result.RowsAffected()
	think.IsNil(err)

	return affect
}

func Update(tx *sql.Tx, sqlString string, args ...interface{}) (affect int64, err error) {
	thinkLog.DebugLog.Println(SprintSQL(sqlString, args...))
	var result sql.Result
	if tx == nil {
		result, err = Idb.Exec(sqlString, args...)
	} else {
		result, err = tx.Exec(sqlString, args...)
	}
	if err != nil {
		return 0, err
	}
	affect, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affect, nil
}

func Insert(tx *sql.Tx, sqlString string, args ...interface{}) (last int64, err error) {
	thinkLog.DebugLog.Println(SprintSQL(sqlString, args...))
	var result sql.Result
	if tx == nil {
		result, err = Idb.Exec(sqlString, args...)
	} else {
		result, err = tx.Exec(sqlString, args...)
	}
	if err != nil {
		return 0, err
	}
	last, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return last, nil
}

func InsertIgnoreBatch(tx *sql.Tx ,tableName string, cols []string, values [][]string) int64 {
	sqlString := "INSERT IGNORE INTO " + tableName + " ("
	for i := 0; i < len(cols); i++ {
		sqlString += cols[i] + ","
	}
	sqlString = sqlString[:len(sqlString)-1] + ")"

	var valueString = " VALUES "
	for i := 0; i < len(values); i++ {
		value := values[i]
		valueString += "("
		for j := 0; j < len(value); j++ {
			valueString += "'" + value[j] + "'" + ","
		}
		valueString = valueString[:len(valueString)-1] + ")"
		//thinkString.ReplaceLastRune(&valueString, ')')
		valueString += ","
	}
	valueString = valueString[:len(valueString)-1]

	sqlString += valueString
	thinkLog.DebugLog.Println(sqlString)

	var result sql.Result
	var err error
	var affect int64
	if tx == nil {
		result, err = Idb.Exec(sqlString)
	} else {
		result, err = tx.Exec(sqlString)
	}
	think.IsNil(err)
	affect, err = result.RowsAffected()
	think.IsNil(err)

	return affect
}

// INSERT INTO table_name (col1,col2,col3...) VALUES (v1,v2,v3...),(v1,v2,v3...),(v1,v2,v3...)
func InsertBatch(tx *sql.Tx ,tableName string, cols []string, values [][]string) int64 {
	//sqlString := insertKind + " " + tableName + " ("
	sqlString := "INSERT INTO " + tableName + " ("
	for i := 0; i < len(cols); i++ {
		sqlString += cols[i] + ","
	}
	sqlString = sqlString[:len(sqlString)-1] + ")"

	var valueString = " VALUES "
	for i := 0; i < len(values); i++ {
		value := values[i]
		valueString += "("
		for j := 0; j < len(value); j++ {
			valueString += "'" + value[j] + "'" + ","
		}
		valueString = valueString[:len(valueString)-1] + ")"
		//thinkString.ReplaceLastRune(&valueString, ')')
		valueString += ","
	}
	valueString = valueString[:len(valueString)-1]

	sqlString += valueString
	thinkLog.DebugLog.Println(sqlString)

	var result sql.Result
	var err error
	var affect int64
	if tx == nil {
		result, err = Idb.Exec(sqlString)
	} else {
		result, err = tx.Exec(sqlString)
	}
	think.IsNil(err)
	affect, err = result.RowsAffected()
	think.IsNil(err)

	return affect
}

func SelectMap(tx *sql.Tx, sqlString string, args ...interface{}) ([]string, []map[string]string) {
	thinkLog.DebugLog.Println(SprintSQL(sqlString, args...))
	resultMapSlice := make([]map[string]string, 0)
	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = Idb.Query(sqlString, args...)
	} else {
		rows, err = tx.Query(sqlString, args...)
	}
	defer rows.Close()
	think.IsNil(err)
	columns, err := rows.Columns()
	//fmt.Println("columns",len(columns),columns[0])
	think.IsNil(err)
	for rows.Next() {
		resultMap := make(map[string]string)
		tempByteSLice := make([]sql.RawBytes, len(columns))
		tempSlice := make([]interface{}, len(columns))
		for i := 0; i < len(columns); i++ {
			tempSlice[i] = &tempByteSLice[i]
		}
		rows.Scan(tempSlice...)

		for i := 0; i < len(columns); i++ {
			temp := tempByteSLice[i]
			if temp == nil {
				resultMap[columns[i]] = ""
			} else {
				resultMap[columns[i]] = string(temp)
			}
			//fmt.Println(resultMap)
		}

		resultMapSlice = append(resultMapSlice, resultMap)
	}

	return columns, resultMapSlice
}
func GetColumnsType(sqlString string) {
	var rows *sql.Rows
	rows, err := Idb.Query(sqlString)
	think.IsNil(err)
	defer rows.Close()
	for rows.Next() {
		columnsType, err := rows.ColumnTypes()
		think.IsNil(err)
		for j := 0; j < len(columnsType); j++ {
			fmt.Println(columnsType[j].ScanType(), "")
		}
		fmt.Println()
	}

}
func SelectList(tx *sql.Tx, sqlString string, args ...interface{}) ([]string, [][]string) {
	thinkLog.DebugLog.Println(SprintSQL(sqlString, args...))
	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = Idb.Query(sqlString, args...)
	} else {
		rows, err = tx.Query(sqlString, args...)
	}
	think.IsNil(err)
	defer rows.Close()
	columns, err := rows.Columns()
	think.IsNil(err)

	// 因为 go不会自动把 slice 转换成 interface{} 类型的 slice
	// 所以必须手动转换
	// 有值才有指针
	results := make([][]string, 0)
	for rows.Next() {
		value := make([]interface{}, 0)
		tempResult := make([]sql.RawBytes, len(columns))
		for i := 0; i < len(tempResult); i++ {
			value = append(value, &tempResult[i])
		}
		rows.Scan(value...)
		result := make([]string, len(columns))
		for i := 0; i < len(tempResult); i++ {
			temp := tempResult[i]
			if temp == nil {
				result[i] = ""
			} else {
				result[i] = string(temp)
			}
		}
		//fmt.Println(colsName,result)
		results = append(results, result)
		//fmt.Println("results:",results)
	}

	return columns, results
}

// 单行
func SelectGet(tx *sql.Tx, sqlString string, args ...interface{}) ([]string, []string) {
	columns, results := SelectList(tx, sqlString, args...)
	if len(results) == 0 {
		return columns, make([]string, len(columns))
	} else {
		return columns, results[0]
	}
}
