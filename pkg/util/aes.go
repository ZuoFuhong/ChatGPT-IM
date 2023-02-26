package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"errors"
	"log"
)

type TokenInfo struct {
	UserId   int64 `json:"user_id"`   // 用户id
	DeviceId int64 `json:"device_id"` // 设备id
	Expire   int64 `json:"expire"`    // 过期时间
}

// GetToken 获取token
func GetToken(userId, deviceId int64, expire int64, publicKey string) (string, error) {
	info := TokenInfo{
		UserId:   userId,
		DeviceId: deviceId,
		Expire:   expire,
	}
	bytes, err := json.Marshal(info)
	if err != nil {
		log.Print(err)
		return "", err
	}

	token, err := RsaEncrypt(bytes, []byte(publicKey))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(token), nil
}

// DecryptToken 对加密的token进行解码
func DecryptToken(token string, privateKey string) (*TokenInfo, error) {
	bytes, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	result, err := RsaDecrypt(bytes, Str2bytes(privateKey))
	if err != nil {
		log.Print(err)
		return nil, err
	}

	var info TokenInfo
	err = json.Unmarshal(result, &info)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return &info, nil
}

// RsaEncrypt 加密
func RsaEncrypt(origData []byte, publicKey []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// RsaDecrypt 解密
func RsaDecrypt(ciphertext []byte, privateKey []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
