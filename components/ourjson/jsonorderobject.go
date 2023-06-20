package ourjson

import (
	"sync"

	"github.com/bytedance/sonic"
	"github.com/elliotchance/orderedmap/v2"
	glog "github.com/khaosles/gtools2/core/log"
)

/*
   @File: jsonorderobject.go
   @Author: khaosles
   @Time: 2023/6/20 09:49
   @Desc:
*/

type JsonOrderObject struct {
	m     *orderedmap.OrderedMap[string, *Value]
	mutex sync.Mutex
}

func NewJsonOrderObject() *JsonOrderObject {
	return &JsonOrderObject{m: orderedmap.NewOrderedMap[string, *Value]()}
}

func (j *JsonOrderObject) HasKey(key string) bool {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	if j.m == nil {
		return false
	}
	_, ok := j.m.Get(key)
	return ok
}

func (j *JsonOrderObject) Get(key string) (*Value, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	if !j.HasKey(key) {
		return nil, KeyNotFoundError{key}
	}
	value, _ := j.m.Get(key)
	return value, nil
}

func (j *JsonOrderObject) GetJsonOrderObject(key string) *JsonOrderObject {
	val, err := j.Get(key)
	if err != nil {
		panic(err)
	}
	return val.JsonOrderObject()
}

func (j *JsonOrderObject) GetJsonArray(key string) *JsonArray {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		panic(err)
	}
	return val.JsonArray()
}

func (j *JsonOrderObject) GetString(key string) (string, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return "", err
	}
	return val.String()
}

func (j *JsonOrderObject) GetInt(key string) (int, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int()
}

func (j *JsonOrderObject) GetNullInt(key string) (*Integer, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullInt()
}

func (j *JsonOrderObject) GetInt64(key string) (int64, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Int64()
}

func (j *JsonOrderObject) GetNullLong(key string) (*Long, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullLong()
}

func (j *JsonOrderObject) GetFloat64(key string) (float64, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return 0, err
	}
	return val.Float64()
}

func (j *JsonOrderObject) GetNullFloat(key string) (*Float, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullFloat()
}

func (j *JsonOrderObject) GetBoolean(key string) (bool, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return false, err
	}
	return val.Boolean()
}

func (j *JsonOrderObject) GetNullBoolean(key string) (*Boolean, error) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	val, err := j.Get(key)
	if err != nil {
		return nil, err
	}
	return val.NullBoolean()
}

func (j *JsonOrderObject) Put(key string, val interface{}) {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	j.m.Set(key, &Value{val})
}

func (j *JsonOrderObject) String() string {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	if j.m == nil {
		return ""
	}
	data, err := sonic.Marshal(j.m)
	if err != nil {
		glog.Error(err)
		return ""
	}
	return string(data)
}

func (j *JsonOrderObject) Value() orderedmap.OrderedMap[string, *Value] {
	defer j.mutex.Unlock()
	j.mutex.Lock()
	return *j.m
}
