package util

import (
	"ChatGPT-IM/backend/consts"
	"fmt"
	"testing"
)

func Test_GetToken(t *testing.T) {
	token, err := GetToken(1629756677962076160, 1629756789316653056, 1000000000000000, consts.PublicKey)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(token)
	fmt.Println(DecryptToken(token, consts.PrivateKey))
}
