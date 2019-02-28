package thinkJson

import (
	"encoding/json"
)

func MustMarshal(obj interface{}) []byte {
	paramsJson, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return paramsJson
}
