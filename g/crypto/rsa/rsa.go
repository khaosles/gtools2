package grsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
)

// Encrypt encrypt data
func Encrypt(src string, key *rsa.PublicKey) (data string, err error) {
	h := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(h, rand.Reader, key, []byte(src), nil)

	if err != nil {
		return "", err
	}
	return hex.EncodeToString(ciphertext), nil
}

// Decrypt decrypt data
func Decrypt(text string, key *rsa.PrivateKey) (data string, err error) {
	src, _ := hex.DecodeString(text)
	h := sha256.New()
	oaep, err := rsa.DecryptOAEP(h, rand.Reader, key, src, nil)
	if err != nil {
		return "", err
	}
	return string(oaep), nil
}

// Sign sign data
func Sign(key *rsa.PrivateKey, src []byte) (sign []byte, err error) {
	h := crypto.SHA256
	hn := h.New()
	hn.Write(src)
	sum := hn.Sum(nil)
	return rsa.SignPSS(rand.Reader, key, h, sum, nil)
}

// Verify verify data
func Verify(key *rsa.PublicKey, sign, src []byte) (err error) {
	h := crypto.SHA256
	hn := h.New()
	hn.Write(src)
	sum := hn.Sum(nil)
	return rsa.VerifyPSS(key, h, sum, sign, nil)
}

// CreateKeyX509PKCS1 create rsa keys
func CreateKeyX509PKCS1(bits int) (pub string, pri string) {
	privateKey, _ := rsa.GenerateKey(rand.Reader, bits)
	publicKey := &privateKey.PublicKey
	bytePri := x509.MarshalPKCS1PrivateKey(privateKey)
	pri = base64.StdEncoding.EncodeToString(bytePri)
	bytePub := x509.MarshalPKCS1PublicKey(publicKey)
	pub = base64.StdEncoding.EncodeToString(bytePub)
	return pub, pri
}

// PrivateKeyFromX509PKCS1 string private key to rsa.PrivateKey
func PrivateKeyFromX509PKCS1(pri string) (*rsa.PrivateKey, error) {
	data, err := base64.StdEncoding.DecodeString(pri)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PrivateKey(data)
}

// PrivateKeyToPKCS1 convert private key to a string
func PrivateKeyToPKCS1(pri *rsa.PrivateKey) string {
	bytePri := x509.MarshalPKCS1PrivateKey(pri)
	return base64.StdEncoding.EncodeToString(bytePri)
}

// PublicKeyToPKCS1 convert public key to a string
func PublicKeyToPKCS1(pub *rsa.PublicKey) string {
	bytePub := x509.MarshalPKCS1PublicKey(pub)
	return base64.StdEncoding.EncodeToString(bytePub)
}

// PublicKeyFromX509PKCS1 convert public key to a string
func PublicKeyFromX509PKCS1(pub string) (*rsa.PublicKey, error) {
	data, err := base64.StdEncoding.DecodeString(pub)
	if err != nil {
		return nil, err
	}
	return x509.ParsePKCS1PublicKey(data)
}
