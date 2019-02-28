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
	think.IsNil(err)

	_, rows := SelectList(nil, "SHOW GLOBAL VARIABLES")
	var waitTimeout string
	// max_allowed_packet
	for i := 0; i < len(rows); i++ {
		//fmt.Println(rows[i])
		if strings.Compare(rows[i][0], "wait_timeout") == 0 {
			waitTimeout = rows[i][1]
			break
		}
	}
	//fmt.Println("waitTimeout",waitTimeout)
	timeout, err := strconv.Atoi(waitTimeout)
	think.IsNil(err)
	//fmt.Println("timeout",time.Duration(timeout) * time.Second)
	// 设置连接的最长存活时间 = mysql.cnf.wait_timeout
	//int(time.Second) = 1000000000
	// 连接存活最长存活时间,(若idle为0,此项无意义)
	Idb.SetConnMaxLifetime(time.Duration(timeout) * time.Second)
	// 最多打开的连接数
	Idb.SetMaxOpenConns(20)
	// 空闲连接
	Idb.SetMaxIdleConns(0)
	//source = sourceName

	//go func() {
	//	for true {
	//		time.Sleep(time.Millisecond * 1000)
	//		fmt.Println("database: ",Idb.Stats().OpenConnections)
	//	}
	//}()
	return Idb
}

// 63206,66492

// transaction
func BeginTrans() *sql.Tx {
	tx, err := Idb.Begin()
	think.IsNil(err)
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
