package costtime

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	CostTime("name string", func() {
		time.Sleep(1)
	})

	CostTime("name string", func() {
		time.Sleep(time.Millisecond * 100)
	})

	CostTime("name string", func() {
		time.Sleep(time.Millisecond * 1000)
	})

}
