package database

import (
	"fmt"
	"util/thinkString"
)

func SprintSQL(sqlString string, args ...interface{}) string {
	log := ""
	log += "\n"
	log += "query:  " + sqlString + "\n"
	if len(args) != 0 {
		log += "params: "
		for _, str := range args {
			log += fmt.Sprint(str)
			log += ","
		}
		thinkString.ReplaceLastRune(&log, '\n')
	}
	return log
}