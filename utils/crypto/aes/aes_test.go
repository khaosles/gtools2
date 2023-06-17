package gaes

/*
   @File: aes_test.go
   @Author: khaosles
   @Time: 2023/6/3 22:02
   @Desc:
*/

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAes(t *testing.T) {
	s := "0123456789"
	key := "01234567890123456789012345678901"
	fmt.Println(len(key))
	encrypt, err := Encrypt(s, key)
	fmt.Println(encrypt)
	if err != nil {
		assert.NoError(t, err)
	}
	decrypt, err := Decrypt(encrypt, key)
	if err != nil {
		assert.NoError(t, err)
	}
	fmt.Println(decrypt)
	assert.Equal(t, s, decrypt)
}
