package costtime

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

// ConditionFunc 日志输出条件判断函数
type ConditionFunc func(cost time.Duration) bool

var costlog = func() *log.Logger {
	l := log.New(os.Stderr, "", log.Ldate|log.Ltime)
	return l
}()

var colors = []string{"\033[32m%d\033[0m", "\033[34m%d\033[0m", "\033[31m%d\033[0m", "\033[31m\033[05m%d\033[0m"}

var condition ConditionFunc

// SetLogCondition 设置输出cost条件
func SetLogCondition(cond ConditionFunc) {
	condition = cond
}

// Cost 里面计算消耗时间
func Cost(run func()) {

	file, line, funcName := getRuntimeInfo()

	now := time.Now()
	run()
	cost, coststr := countCostString(now)
	if condition != nil {
		if !condition(cost) {
			return
		}
	}
	costlog.Printf("%s:%d(%s) cost(%s ms)", file, line, funcName, coststr)
}

// CostLog 计算消耗的时间
func CostLog(name string, run func()) {
	file, line, funcName := getRuntimeInfo()

	now := time.Now()
	run()
	cost, coststr := countCostString(now)
	if condition != nil {
		if !condition(cost) {
			return
		}
	}
	costlog.Printf("%s:%d(%s) cost(%s ms): %s", file, line, funcName, coststr, name)
}

func countCostString(now time.Time) (time.Duration, string) {
	end := time.Now()

	var selcolor string
	cost := end.Sub(now)
	switch {
	case cost < time.Millisecond*100:
		selcolor = colors[0]
	case cost < time.Second:
		selcolor = colors[1]
	case cost < time.Second*10:
		selcolor = colors[2]
	default:
		selcolor = colors[3]
	}
	return cost, fmt.Sprintf(selcolor, cost.Milliseconds())
}

func getRuntimeInfo() (file string, line int, funcName string) {
	pc, file, line, _ := runtime.Caller(2)
	funcName = runtime.FuncForPC(pc).Name()

	var i int
	i = strings.LastIndexByte(funcName, '.')
	funcName = funcName[i+1:]

	i = strings.LastIndexByte(file, '/')
	file = file[i+1:]

	return file, line, funcName
}
