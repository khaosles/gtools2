package ghash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

/*
   @File: hash.go
   @Author: khaosles
   @Time: 2023/6/4 00:22
   @Desc:
*/

func Sha1(orig string, salt string) string {
	hn := sha1.New()
	hn.Write([]byte(orig))
	hn.Write([]byte(salt))
	data := hn.Sum([]byte(""))
	return hex.EncodeToString(data)
}

func Sha256(orig string, salt string) string {
	hn := sha256.New()
	hn.Write([]byte(orig))
	hn.Write([]byte(salt))
	data := hn.Sum([]byte(""))
	return hex.EncodeToString(data)
}

func Md5(orig string, salt string) string {
	hn := md5.New()
	hn.Write([]byte(orig))
	hn.Write([]byte(salt))
	data := hn.Sum([]byte(""))
	return hex.EncodeToString(data)
}

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// 第一次加密
	encrypt1, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// aes加密
	return string(encrypt1)
}

// BcryptCheck 对比明文密码和数据库的哈希值
func BcryptCheck(password, hash string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
