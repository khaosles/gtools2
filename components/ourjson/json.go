//

package ourjson

import (
	"github.com/bytedance/sonic"
)

func ParseObject(jsonStr string) (*JsonObject, error) {
	value := new(Value)

	err := sonic.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}

	return value.JsonObject(), nil
}

func ParseArray(jsonStr string) (*JsonArray, error) {
	value := new(Value)

	err := sonic.Unmarshal([]byte(jsonStr), &value.data)
	if err != nil {
		return nil, err
	}
	return value.JsonArray(), nil
}
