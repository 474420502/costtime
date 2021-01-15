package costtime

import (
	"testing"
	"time"
)

func TestBase(t *testing.T) {

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

	Cost(func() {
		time.Sleep(time.Nanosecond)
	})
}

func TestCond(t *testing.T) {
	SetLogCondition(CondGT(time.Millisecond * 10))

	CostLog("name string", func() {
		time.Sleep(1)
	})

	SetLogCondition(CondGT(time.Millisecond * 101))

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 100)
	})

	SetLogCondition(CondLTE(time.Millisecond * 1001))

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 1000)
	})

	SetLogCondition(CondLT(time.Millisecond * 10))

	CostLog("name string", func() {
		time.Sleep(time.Millisecond)
	})

	SetLogCondition(CondGTE(time.Millisecond * 101))

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 10)
	})

	SetLogCondition(CondRange(time.Millisecond*50, time.Millisecond*101))

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 10)
	})

	CostLog("name string", func() {
		time.Sleep(time.Millisecond * 50)
	})
}
