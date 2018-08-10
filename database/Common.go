package database

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
	"util/think"
	"util/thinkString"
	"util/thinkTimer"
)

var Idb *sql.DB
var source string
var CacheDatabase map[string]interface{}


func SetConn(sourceName string) *sql.DB {
	var err error
	Idb, err = sql.Open("mysql", sourceName)
	think.Check(err)

	_, rows := SelectList(nil, "SHOW GLOBAL VARIABLES")
	var waitTimeout string
	var maxAllowedPacket string
	for i := 0; i < len(rows); i++ {
		if strings.Contains(rows[i][0], "wait_timeout") {
			waitTimeout = rows[i][1]
			break
		}
		if strings.Contains(rows[i][0], "max_allowed_packet") {
			maxAllowedPacket = rows[i][1]
		}
	}
	//fmt.Println("waitTimeout",waitTimeout)
	_ = maxAllowedPacket
	timeout, err := strconv.Atoi(waitTimeout)
	think.Check(err)
	//fmt.Println("timeout",time.Duration(timeout) * time.Second)
	// 设置连接的最长存活时间 = mysql.cnf.wait_timeout
	//int(time.Second) = 1000000000
	Idb.SetConnMaxLifetime(time.Duration(timeout) * time.Second)
	Idb.SetMaxOpenConns(5)
	source = sourceName
	return Idb
}

// 63206,66492

// transaction
func BeginTrans() *sql.Tx {
	tx, err := Idb.Begin()
	think.Check(err)
	return tx
}

//KeepAlive
func KeepAliveTask() {
	go thinkTimer.Task("* * * * 0 0", keepAlive)
}
func keepAlive() {
	if Idb != nil {
		SelectList(nil, "SELECT user FROM mysql.user")
	}
}

func Cache(f func(args ...interface{}) interface{}, args ...interface{}) interface{} {
	k := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	for i := 0; i < len(args); i++ {
		k += thinkString.AsString(args[i])
	}
	v, ok := CacheDatabase[k]
	if ok {
		return v
	}
	// 执行
	v = f(args)
	// 保存
	CacheDatabase[k] = v

	return v
}
