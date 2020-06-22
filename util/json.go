package util

import "encoding/json"

func ToJsonString(v interface{}) string {
	var bytes []byte
	var err error
	if bytes, err = json.Marshal(bytes); err != nil {
		panic(err)
	}
	return string(bytes)
}
