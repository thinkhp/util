package thinkJson

import (
	"bytes"
	"encoding/json"
)

func MustMarshal(obj interface{}) []byte {
	paramsJson, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return paramsJson
}

func MustMarshalWithoutEscapeHTML(obj interface{}) []byte {
	bf := bytes.NewBuffer([]byte{})
	json.NewEncoder(bf)
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(obj)
	if err != nil {
		panic(err)
	}
	return bf.Bytes()
}


