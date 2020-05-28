package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_group(t *testing.T) {
	_, e := json.Marshal(nil)
	fmt.Print(e)
}
