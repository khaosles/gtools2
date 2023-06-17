package gmachine

import (
	"errors"

	"github.com/denisbrodbeck/machineid"
)

// GetId 获得机器唯一id
func GetId(name string) (string, error) {
	id, err := machineid.ProtectedID(name)
	if err != nil {
		return "", errors.New("get machine id failed")
	}
	return id, nil
}
