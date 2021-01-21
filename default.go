package costtime

import (
	"log"
	"os"
)

var logDirectory = "costtimelog"
var logDefaultName = "default"

func init() {
	info, err := os.Stat(logDirectory)
	if err != nil {
		log.Println(err)
	}

	if info == nil {
		err = os.Mkdir(logDirectory, os.ModeDir|os.ModeTemporary|os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	} else {
		if !info.IsDir() {
			log.Panicf("%s is exists. and not directory.", logDirectory)
		}
	}

	defaultCost = New(logDefaultName)
	defaultCost.skip = 3
}

var defaultCost *CostTime

// CostLog 计算消耗的时间
func CostLog(name string, run func()) {
	defaultCost.CostLog(name, run)
}

// Cost 里面计算消耗时间
func Cost(run func()) {
	defaultCost.Cost(run)
}

// SetLogCondition 设置输出cost条件
func SetLogCondition(cond ConditionFunc) {
	defaultCost.SetLogCondition(cond)
}

// SetEventCost 设置输出cost事件. 可以做邮件通知. 钉釘等办公通知
func SetEventCost(event EventFunc) {
	defaultCost.SetEeventCost(event)
}
