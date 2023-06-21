package grsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"sync"
)

/*
   @File: rsa_with_java.go
   @Author: khaosles
   @Time: 2023/6/21 11:40
   @Desc:
*/

var rsaWithJavaInstance RsaWithJava
var once sync.Once

type RsaWithJava interface {
	// Encrypt 与java匹配的加密
	Encrypt(data, publicKey string) (string, error)
	// Decrypt 与java匹配的解密
	Decrypt(data, privateKey string) (string, error)
	// Sign 与java匹配的签名
	Sign(data, privateKey string) (string, error)
	// Verify 与java匹配的验证签名
	Verify(data, sign, publicKey string) error
	// GenKey 生成密钥
	GenKey(bits int) (string, string)
}

func GetRsaWithJavaInstance() RsaWithJava {
	once.Do(func() {
		if rsaWithJavaInstance == nil {
			rsaWithJavaInstance = &rsaWithJava{}
		}
	})
	return rsaWithJavaInstance
}

type rsaWithJava struct{}

// Encrypt 与java匹配的加密
func (r rsaWithJava) Encrypt(data, publicKey string) (string, error) {
	publicKey = fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKey)
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return "", errors.New("public key error")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", err
	}
	pub := pubInterface.(*rsa.PublicKey)
	encryptPKCS1v15, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(data))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encryptPKCS1v15), nil
}

// Decrypt 与java匹配的解密
func (r rsaWithJava) Decrypt(data, privateKey string) (string, error) {
	privateKey = fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error!")
	}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	priv := privInterface.(*rsa.PrivateKey)
	decodeString, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decryptPKCS1v15, err := rsa.DecryptPKCS1v15(rand.Reader, priv, decodeString)
	if err != nil {
		return "", err
	}
	return string(decryptPKCS1v15), nil
}

// Sign 与java匹配的签名
func (r rsaWithJava) Sign(data, privateKey string) (string, error) {
	privateKey = fmt.Sprintf("-----BEGIN PRIVATE KEY-----\n%s\n-----END PRIVATE KEY-----", privateKey)
	block, _ := pem.Decode([]byte(privateKey))
	if block == nil {
		return "", errors.New("private key error!")
	}
	privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}
	priv := privInterface.(*rsa.PrivateKey)
	h := crypto.SHA256
	hn := h.New()
	hn.Write([]byte(data))
	sum := hn.Sum(nil)
	signPKCS1v15, err := rsa.SignPKCS1v15(rand.Reader, priv, h, sum)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signPKCS1v15), nil
}

// Verify 与java匹配的验证签名
func (r rsaWithJava) Verify(data, sign, publicKey string) error {
	publicKey = fmt.Sprintf("-----BEGIN PUBLIC KEY-----\n%s\n-----END PUBLIC KEY-----", publicKey)
	block, _ := pem.Decode([]byte(publicKey))
	if block == nil {
		return errors.New("public key error!")
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)
	h := crypto.SHA256
	hn := h.New()
	hn.Write([]byte(data))
	sum := hn.Sum(nil)
	signData, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return err
	}
	err = rsa.VerifyPKCS1v15(pub, h, sum, signData)
	return err
}

// GenKey 生成密钥
func (r rsaWithJava) GenKey(bits int) (string, string) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	publicKey := &privateKey.PublicKey
	bytePri, _ := x509.MarshalPKCS8PrivateKey(privateKey)
	pri := base64.StdEncoding.EncodeToString(bytePri)
	bytePub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pub := base64.StdEncoding.EncodeToString(bytePub)
	fmt.Println("私钥： ", pri)
	fmt.Println("公钥： ", pub)
	return pub, pri
}
