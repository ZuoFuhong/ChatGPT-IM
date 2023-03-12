package service

import (
	"fmt"
	"testing"
)

func Test_ProxyRobotPost(t *testing.T) {
	answer, err := ProxyRobotPost("hello")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(answer)
}
