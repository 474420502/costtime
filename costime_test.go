package costtime

import (
	"testing"
	"time"
)

func TestDefault(t *testing.T) {
	CostLog("name string", func() {
		time.Sleep(1)
	})

	Cost(func() {
		time.Sleep(time.Millisecond * 100)
	})
}

func TestEvent(t *testing.T) {

	e := New("test")
	e.SetLogCondition(CondGT(100 * time.Millisecond))
	e.SetEeventCost(func(name string, cost time.Duration) {
		if name != "event" {
			t.Error("event")
		}
	})

	e.Cost(func() {

	})
}

func TestBase(t *testing.T) {

	defaultCost.CostLog("name string", func() {
		time.Sleep(1)
	})

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 100)
	})

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 200)
	})

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 1000)
	})

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 8000)
	})

	defaultCost.Cost(func() {
		time.Sleep(time.Nanosecond)
	})
}

func TestTreeLog(t *testing.T) {
	defaultCost.CostLog("deep1", func() {
		time.Sleep(time.Millisecond * 10)
		defaultCost.CostLog("deep2", func() {

			time.Sleep(time.Millisecond * 100)
			defaultCost.Cost(func() {
				time.Sleep(time.Millisecond * 100)
			})
			defaultCost.CostLog("deep3", func() {
				time.Sleep(time.Millisecond * 100)
			})

			defaultCost.CostLog("deep3-2", func() {
				time.Sleep(time.Millisecond * 100)
				defaultCost.CostLog("deep4", func() {
					time.Sleep(time.Millisecond * 100)
					In()
				})
			})

		})
	})

	defaultCost.Cost(func() {
		time.Sleep(time.Millisecond * 10)
		defaultCost.CostLog("child string", func() {
			time.Sleep(time.Millisecond * 100)
			defaultCost.Cost(func() {
				time.Sleep(time.Millisecond * 100)
			})
			defaultCost.CostLog("child string", func() {
				time.Sleep(time.Millisecond * 100)
			})
			defaultCost.Cost(func() {
				time.Sleep(time.Millisecond * 100)
				defaultCost.CostLog("new child string", func() {
					time.Sleep(time.Millisecond * 100)
				})
			})

		})
	})

}

func TestCond(t *testing.T) {
	defaultCost.SetLogCondition(CondGT(time.Millisecond * 10))

	defaultCost.CostLog("name string", func() {
		time.Sleep(1)
	})

	defaultCost.SetLogCondition(CondGT(time.Millisecond * 101))

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 100)
	})

	defaultCost.SetLogCondition(CondLTE(time.Millisecond * 1001))

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 1000)
	})

	defaultCost.SetLogCondition(CondLT(time.Millisecond * 10))

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond)
	})

	defaultCost.SetLogCondition(CondGTE(time.Millisecond * 101))

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 10)
	})

	defaultCost.SetLogCondition(CondRange(time.Millisecond*50, time.Millisecond*101))

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 10)
	})

	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 50)
	})
}
