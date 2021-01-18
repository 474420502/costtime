package main

import (
	"time"

	"github.com/474420502/costtime"
)

func In() {
	costtime.CostLog("In string", func() {
		time.Sleep(time.Millisecond * 200)
		costtime.CostLog("In string deep", func() {
			time.Sleep(time.Millisecond * 200)
		})
	})
}
