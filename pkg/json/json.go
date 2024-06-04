// Package jsonPkg 对象序列化方法
package jsonPkg

import "encoding/json"

// Marshal 序列化
func Marshal(value interface{}) (string, error) {
	var b []byte
	b, err := json.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// UnMarshal 反序列化
func UnMarshal(data string, value interface{}) error {
	err := json.Unmarshal([]byte(data), value)
	return err
}
