package mesh_v1

import (
	"encoding/json"
	"fmt"
)

// func ToJsonStr(obj *interface{}) (string, *error) {
// 	jsonBytes, marshalErr := json.Marshal(obj)
// 	if marshalErr != nil {
// 		return "", &marshalErr
// 	}
// 	return string(jsonBytes), nil
// }

func ToJsonStr(data any) (string, error) {
	// Convert the data to JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	// Convert JSON bytes to a string and return
	return string(jsonBytes), nil
}

func Log(message string, params ...any) {
	msg := fmt.Sprintf(message, params...)
	fmt.Printf("[mesh_v1]: "+msg+"\n", params...)
	// fmt.Print("[mesh_v1]: %v", msg)
}

func LogWithJson[T any](msg string, obj T) error {
	jsonStr, err := ToJsonStr(obj)
	if err != nil {
		return err
	}
	Log(msg, jsonStr)
	return nil
}
