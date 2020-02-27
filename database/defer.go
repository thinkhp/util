package database

import (
	"database/sql"
	"runtime/debug"
	"util/think"
	"util/thinkLog"
)

func ClearTransaction(tx *sql.Tx) {
	if tx == nil {
		return
	}
	err := tx.Rollback()
	if err == nil {
		// 回滚无异常,说明发生回滚
		thinkLog.ErrorLog.Println("[rollback] 事务回滚")
		// 发生回滚,说明有异常发生,输出之前发生的异常
		thinkLog.ErrorLog.Println("事务未提交,或请自行查找错误原因")
		thinkLog.ErrorLog.Println(string(debug.Stack()))
	} else {
		// 存在err
		if err == sql.ErrTxDone { // 已经commit或者rollback
			thinkLog.DebugLog.Println("[commit|rollback] 事务已提交")
		} else { // 存在未知异常
			// 输出之前的异常
			// TODO 未知异常是否应该输出日志并 panic
			thinkLog.ErrorLog.Println("请自行查找错误原因")
			thinkLog.ErrorLog.Println(string(debug.Stack()))
			think.IsNil(err)
		}
	}
}
