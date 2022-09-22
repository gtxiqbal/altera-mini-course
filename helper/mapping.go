package helper

import (
	"github.com/goccy/go-json"
)

func MappingStruct(dst any, src any) error {
	jsonByte, err := json.Marshal(src)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonByte, dst)
}
