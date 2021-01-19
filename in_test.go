package costtime

import "time"

func In() {
	defaultCost.CostLog("In string", func() {
		time.Sleep(time.Millisecond * 200)
		defaultCost.CostLog("In string deep", func() {
			time.Sleep(time.Millisecond * 200)
		})
	})
}
