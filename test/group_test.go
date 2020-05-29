package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_group(t *testing.T) {
	_, e := json.Marshal(nil)
	fmt.Print(e)
}

func Test_Time(t *testing.T) {
	fmt.Println(time.Now().Add(0))
}
