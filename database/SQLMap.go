package database

import (
	"database/sql"
	"fmt"
	"util/think"
)

func Update(tx *sql.Tx, sqlString string, args ...interface{}) int64 {
	var result sql.Result
	var err error
	var affect int64
	if tx == nil {
		result, err = Idb.Exec(sqlString, args...)
	} else {
		result, err = tx.Exec(sqlString, args...)
	}
	think.Check(err)
	affect, err = result.RowsAffected()
	think.Check(err)

	return affect
}

func Insert(tx *sql.Tx, sqlString string, args ...interface{}) int64 {
	var result sql.Result
	var err error
	var last int64
	if tx == nil {
		result, err = Idb.Exec(sqlString, args...)
	} else {
		result, err = tx.Exec(sqlString, args...)
	}
	think.Check(err)
	last, err = result.LastInsertId()
	think.Check(err)

	return last
}

func SelectMap(tx *sql.Tx, sqlString string, args ...interface{}) ([]string, []map[string]string) {
	resultMapSlice := make([]map[string]string, 0)
	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = Idb.Query(sqlString, args...)
	} else {
		rows, err = tx.Query(sqlString, args...)
	}
	defer rows.Close()
	think.Check(err)
	columns, err := rows.Columns()
	//fmt.Println("columns",len(columns),columns[0])
	think.Check(err)
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
	think.Check(err)
	for rows.Next() {
		columnsType, err := rows.ColumnTypes()
		think.Check(err)
		for j := 0; j < len(columnsType); j++ {
			fmt.Println(columnsType[j].ScanType(), "")
		}
		fmt.Println()
	}

}
func SelectList(tx *sql.Tx, sqlString string, args ...interface{}) ([]string, [][]string) {
	var rows *sql.Rows
	var err error
	if tx == nil {
		rows, err = Idb.Query(sqlString, args...)
	} else {
		rows, err = tx.Query(sqlString, args...)
	}
	think.Check(err)
	//thinkLog.DebugLog.PrintSQL(sqlString, args)
	columns, err := rows.Columns()
	think.Check(err)

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
