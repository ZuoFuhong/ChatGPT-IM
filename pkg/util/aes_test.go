package util

import (
	"fmt"
	"testing"
	"time"
)

// appId:user_id:device_id:expire token格式
func TestRsaEncrypt(t *testing.T) {
	token, err := GetToken(1000000000000000, 100000000000000000, 100000000000000000, 1000000000000000, PublicKey)
	fmt.Println(err)
	fmt.Println(token)
	fmt.Println(len(token))

	fmt.Println(DecryptToken(token, PrivateKey))
}

func Test_GetToken(t *testing.T) {
	token, _ := GetToken(1, 54146910402904064, 54151683218866176, time.Now().Add(1*time.Hour).Unix(), PublicKey)
	fmt.Println(token)

	token = "BkeIxFOJEmwJQer99pFg0reKM/zAxXyqsIlpuKlfkfQ98GsMNULPBzmQndDF/G0+bmVu0QMkgZ8tzgeR9RZq6el3Of6XLF4QcphxlGV0ODJxv6I3f0ot0vU4OZIzLxqo1jZQfqLneBXjo+vGwX2y4vLyEk1NUaEzH0udPPXD2eA="
	info, _ := DecryptToken(token, PrivateKey)
	fmt.Println(info)
}
