package gaes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

/*
   @File: aes.go
   @Author: khaosles
   @Time: 2023/6/3 22:00
   @Desc:
*/

func Encrypt(text, salt string) (data string, err error) {
	println(salt)
	// 将密钥转换为字节数组
	block, err := aes.NewCipher([]byte(salt))

	src := []byte(text)
	if err != nil {
		return "", err
	} else if len(src) == 0 {
		return "", errors.New("src is empty")
	}

	plaintext, err := pkcs7Pad(src, block.BlockSize())

	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	bm := cipher.NewCBCEncrypter(block, iv)
	bm.CryptBlocks(ciphertext[aes.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

func Decrypt(text, salt string) (data string, err error) {

	src, _ := hex.DecodeString(text)

	if len(src) < aes.BlockSize {
		return "", errors.New("data length error")
	}

	iv := src[:aes.BlockSize]
	ciphertext := src[aes.BlockSize:]

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext is not a multiple of the block size")
	}

	block, err := aes.NewCipher([]byte(salt))
	if err != nil {
		return "", err
	}

	bm := cipher.NewCBCDecrypter(block, iv)
	bm.CryptBlocks(ciphertext, ciphertext)
	ciphertext, err = pkcs7Unpad(ciphertext, aes.BlockSize)

	if err != nil {
		return "", err
	}

	return string(ciphertext), nil
}

// PKCS5Padding 填充函数
func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS5UnPadding 去除填充函数
func PKCS5UnPadding(plaintext []byte) []byte {
	length := len(plaintext)
	unpadding := int(plaintext[length-1])
	return plaintext[:(length - unpadding)]
}
