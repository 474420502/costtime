package main

import (
	"time"

	"github.com/474420502/costtime"
)

func costtest() {
	costtime.CostLog("deep1", func() {
		time.Sleep(time.Millisecond * 10)
		costtime.CostLog("deep2", func() {

			time.Sleep(time.Millisecond * 100)
			costtime.Cost(func() {
				time.Sleep(time.Millisecond * 100)
			})
			costtime.CostLog("deep3", func() {
				time.Sleep(time.Millisecond * 100)
			})

			costtime.CostLog("deep3-2", func() {
				time.Sleep(time.Millisecond * 100)
				costtime.CostLog("deep4", func() {
					time.Sleep(time.Millisecond * 100)
					In()
				})
			})

		})
	})

	// costtime.Cost(func() {
	// 	time.Sleep(time.Millisecond * 10)
	// 	costtime.CostLog("child string", func() {
	// 		time.Sleep(time.Millisecond * 100)
	// 		costtime.Cost(func() {
	// 			time.Sleep(time.Millisecond * 100)
	// 		})
	// 		costtime.CostLog("child string", func() {
	// 			time.Sleep(time.Millisecond * 100)
	// 		})
	// 		costtime.Cost(func() {
	// 			time.Sleep(time.Millisecond * 100)
	// 			costtime.CostLog("new child string", func() {
	// 				time.Sleep(time.Millisecond * 100)
	// 			})
	// 		})

	// 	})
	// })
}
