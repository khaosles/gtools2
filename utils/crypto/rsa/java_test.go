package grsa

/*
   @File: java_test.go
   @Author: khaosles
   @Time: 2023/6/9 21:00
   @Desc:
*/

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"
)

// go 语言支持原生的密钥  java 需要转成pkcs8 才能用
var privateKey = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICdwIBADANBgkqhkiG9w0BAQEFAASCAmEwggJdAgEAAoGBAOAsGajbJJqD1QFY/FlnlXXIQOVIZm7Pt/UMEp9JW38bCV7v0qRvCZCf85/JGCpdYQ+ql6ykank8fs00HmRBM5uOlhboNhI86PyBqGU850WQjmOtCCgbEOCOWM4/D1favEu/p8mVqxJ7gYC4gEcgJICyMxjDBU9B4D6begtLaUpZAgMBAAECgYBVoGK4veQwZRTSq/PQDqHnWHN5YPtHbm5c2pyuXS3m0iP1MHPsPUGRDZfYO87QN9TgUBAZcL/+yR3CMhs9vi4AkOMahgvirviXDtYBrT3nIHRQpZxqEw5EYak8OBXHoIfvSaz90iMCgquMbaZ675g/XoPv32u2/w3lyRrq4G8oUQJBAPzJZ0gIsw6iFCy4+1MzPEqH5xEmx+3q4gG7tp/Y3cTVdrDa+YqOtJA/9T5bUT2KUAYXXb2Fez4xs1pdq/gsNI0CQQDjBZUDyhXg8P6R74VeVz9WX3ypKfoR9n98WOH7C/p/Hc0ylwbDm91AnbR+W883zsE0s1g9c+ZaVQCaeRiiz4f9AkEAitihJxrIJwh1Zl8whHGG8zUUgQI5HIBAJU2SsNfwb7YEHH4aRLW/jd/jd5220MOQ0tewwHF50R6BcegzlfvJ3QJAdJh6Vw7kO7oqVNNagQB4VCkIgm0/tRgPk9KmhWQ6jCzHJbNxUudrM/OLLtaCT5xNmH5/1FgBN+WuQKfvIjdKFQJBAKY06Sc4IGZErUQJKFdVAz/NTPgL6Ed4cNzIpTJfJgbX1PCkiVKL2o+aVPFgojyRVglK/t8ZisNlhr/obJEWIRo=
-----END RSA PRIVATE KEY-----`)

var publicKey = []byte(`-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgLBmo2ySag9UBWPxZZ5V1yEDlSGZuz7f1DBKfSVt/Gwle79KkbwmQn/OfyRgqXWEPqpespGp5PH7NNB5kQTObjpYW6DYSPOj8gahlPOdFkI5jrQgoGxDgjljOPw9X2rxLv6fJlasSe4GAuIBHICSAsjMYwwVPQeA+m3oLS2lKWQIDAQAB
-----END PUBLIC KEY-----`)

func TestSin(t *testing.T) {
	cipherText := RSAEncrypt("123456234567890-09876543")
	fmt.Println("密文：", cipherText)
	content := RSADecrypt(cipherText)
	fmt.Println("解密：" + content)
}

// 加密
func RSAEncrypt(origData string) string {
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "public key error"
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "public key error"
	}
	pub := pubInterface.(*rsa.PublicKey)

	aimByte, _ := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(origData))
	str := base64.StdEncoding.EncodeToString(aimByte)
	return str
}

// 解密
func RSADecrypt(ciphertext string) string {

	block, _ := pem.Decode(privateKey)
	if block == nil {
		return "private key error!"
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "private key error!"
	}
	// base 64 解密
	sDec, _ := base64.StdEncoding.DecodeString(ciphertext)
	bb, _ := rsa.DecryptPKCS1v15(rand.Reader, priv, sDec)

	return string(bb)
}
