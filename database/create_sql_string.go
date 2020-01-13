package database

import "time"

// 拼接 SQL 语句

//
func CreateUpdate(tableName, primaryKeyName string, updateTime time.Time, params map[string]interface{}) (string, []interface{}) {
	s := "UPDATE " + tableName
	l := make([]interface{}, 0)
	if len(params) == 0 { //无参数更新
		return "", l
	}
	id := params[primaryKeyName]
	if id == nil { // 无主键更新,不安全
		return "", l
	}
	s += " SET "
	if !updateTime.IsZero() {
		s += "update_time=?,"
		l = append(l, updateTime.String()[:19])
	}

	for k, v := range params {
		if k == primaryKeyName { // 主键
			continue
		}
		l = append(l, v)
		s += " " + k + "=?,"
	}
	s = s[:len(s)-1] //去除最后一个逗号

	s += " WHERE " + primaryKeyName + "=?"
	l = append(l, id)

	return s, l
}

func CreateInsert(tableName, primaryKeyName string, createTime time.Time, params map[string]interface{}) (string, []interface{}) {
	l := make([]interface{}, 0)
	s := "INSERT INTO " + tableName + " ("
	if !createTime.IsZero() {
		s += "create_time,"
		l = append(l, createTime.String()[:19])
	}

	for k, v := range params {
		if k == primaryKeyName { // 主键自增,不添加进入参数
			continue
		}
		l = append(l, v)
		s += k + ","
	}
	s = s[:len(s)-1]
	s += ") VALUES ("
	for i := 0; i < len(l); i++ {
		s += "?,"
	}
	s = s[:len(s)-1]
	s += ")"

	return s, l
}

func CreateDelete(tableName, primaryKeyName string, primaryKey int) (string, []interface{}) {
	l := make([]interface{}, 0)
	l = append(l, primaryKey)
	s := "DELETE FROM " + tableName + " WHERE " + primaryKeyName + "=?"

	return s, l
}

// 只适合单表
// s : 示例: SELECT * FROM table
// 从参数 params, s 拼接条件(WHERE condation=?,.....)
func CreateSelect(s string, params map[string]interface{}) (string, []interface{}) {
	l := make([]interface{}, 0)
	if len(params) == 0 {
		return s, l
	}
	s += " WHERE"
	for k, v := range params {
		if k == "limit" || k == "offset" { //非 where 条件
			continue
		}
		if v == nil { //空对象
			continue
		}
		if s, ok := v.(string); ok == true && s == "" { //空串
			continue
		}
		l = append(l, v)
		s += " " + k + "=?"
		s += " AND"
	}
	s = s[:len(s)-4]
	limit := params["limit"]
	offset := params["offset"]
	if limit != nil && offset != nil {
		s += " LIMIT ?,?"
		l = append(l, offset, limit)
	}

	return s, l
}
