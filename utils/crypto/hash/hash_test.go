package ghash

import "testing"

/*
   @File: hash_test.go
   @Author: khaosles
   @Time: 2023/6/4 00:22
   @Desc:
*/

func TestMd5(t *testing.T) {
	r := BcryptHash("123")
	println(r)
	println(BcryptCheck("123", r))
	//println(Sha1("111", "123"))
}
