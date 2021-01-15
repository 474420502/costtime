package costtime

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	CostLog("name string", func() {
		time.Sleep(1)
	})

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 100)
	})

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 1000)
	})

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 10000)
	})

}
