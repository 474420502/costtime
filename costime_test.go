package costtime

import (
	"log"
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
		time.Sleep(time.Millisecond * 10000)
	})

	defaultCost.Cost(func() {
		time.Sleep(time.Nanosecond)
	})
}

func TestTreeLog(t *testing.T) {
	defaultCost.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 10)
		defaultCost.CostLog("child string", func() {
			time.Sleep(time.Millisecond * 100)
			defaultCost.Cost(func() {
				time.Sleep(time.Millisecond * 100)
			})
			defaultCost.CostLog("child string", func() {
				time.Sleep(time.Millisecond * 100)
			})
			defaultCost.CostLog("child string", func() {
				time.Sleep(time.Millisecond * 100)
				defaultCost.CostLog("new child string", func() {
					time.Sleep(time.Millisecond * 100)
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

func TestLevel(t *testing.T) {
	c := New()
	c.SetEeventCost(func(cost time.Duration) {
		if cost >= time.Millisecond*100 {
			log.Println("show")
		}
	})
	c.SetLevel(0, time.Millisecond*500)
	c.CostLog("name string", func() {
		time.Sleep(time.Millisecond * 101)
	})
}

func TestEvent(t *testing.T) {
	c := New()
	c.SetEeventCost(func(cost time.Duration) {
		if cost >= time.Millisecond*5 {
			log.Println("show")
		}
	})

	c.Cost(func() {
		time.Sleep(time.Millisecond * 10)
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
