package ourjson

import (
	"encoding/json"

	glog "github.com/khaosles/gtools2/core/log"
)

type JsonObject struct {
	m map[string]*Value
}

func NewJsonObject() *JsonObject {
	return &JsonObject{m: map[string]*Value{}}
}

// Check if the key is existed
func (j *JsonObject) HasKey(key string) bool {
	if j.m == nil {
		return false
	}
	_, ok := j.m[key]
	return ok
}

func (j *JsonObject) Get(key string) (*Value, error) {
	if !j.HasKey(key) {
		return nil, KeyNotFoundError{key}
	}
	return j.m[key], nil
}

// Get a child node of JsonObject from this parent node
func (j *JsonObject) GetJsonObject(key string) *JsonObject {
	val, err := j.Get(key)
	if err != nil {
		panic(err)
	}
	return val.JsonObject()
}

// Get a child node of JsonArray from this parent node
func (j *JsonObject) GetJsonArray(key string) *JsonArray {
	val, err := j.Get(key)
	if err != nil {
		panic(err)
	}
	return val.JsonArray()
}

func (j *JsonObject) GetString(key string) (string, error) {
	val, err := j.Get(key)
	if err != nil {
		return "", err
	}
	return val.String()
}

func (j *JsonObject) GetInt(key string) (int, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int()
}

func (j *JsonObject) GetNullInt(key string) (*Integer, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullInt()
}

func (j *JsonObject) GetInt64(key string) (int64, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int64()
}

func (j *JsonObject) GetNullLong(key string) (*Long, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullLong()
}

func (j *JsonObject) GetFloat64(key string) (float64, error) {
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Float64()
}

func (j *JsonObject) GetNullFloat(key string) (*Float, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullFloat()
}

func (j *JsonObject) GetBoolean(key string) (bool, error) {
	val, err := j.Get(key)
	if err != nil {
		return false, err
	}
	return val.Boolean()
}

func (j *JsonObject) GetNullBoolean(key string) (*Boolean, error) {
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullBoolean()
}

func (j *JsonObject) Put(key string, val interface{}) {
	j.m[key] = &Value{val}
}

func (j *JsonObject) String() string {
	if j.m == nil {
		return ""
	}
	data, err := json.Marshal(j.m)
	if err != nil {
		glog.Error(err)
		return ""
	}
	return string(data)
}

func (j *JsonObject) Value() map[string]*Value {
	return j.m
}
