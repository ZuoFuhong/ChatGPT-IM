package tinyid

import (
	"fmt"
	"testing"
)

func Test_NextId(t *testing.T) {
	t.SkipNow()

	for i := 0; i < 5; i++ {
		id := NextId()
		fmt.Println(id)
	}
}
