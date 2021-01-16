package costtime

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync/atomic"

	"time"
)

type CostTime struct {
	loglevel  int64
	condition ConditionFunc

	costlog       *log.Logger
	costlogNoDate *log.Logger
}

// New 创建一个新Cost. 每个协程里都要创建独立一个
func New() *CostTime {
	c := &CostTime{}
	c.loglevel = -1
	c.costlog = func() *log.Logger {
		l := log.New(os.Stderr, "", log.Ldate|log.Ltime)
		return l
	}()

	c.costlogNoDate = func() *log.Logger {
		l := log.New(os.Stderr, "", log.Ltime)
		return l
	}()
	return c
}

var defaultCost *CostTime = New()

// ConditionFunc 日志输出条件判断函数
type ConditionFunc func(cost time.Duration) bool

// var c.loglevel int64 = -1
var colors = []string{"\033[32m%#v\033[0m", "\033[34m%#v\033[0m", "\033[31m%#v\033[0m", "\033[31m\033[05m%#v\033[0m"}

// SetLogCondition 设置输出cost条件
func SetLogCondition(cond ConditionFunc) {
	defaultCost.SetLogCondition(cond)
}

// SetLogCondition 设置输出cost条件
func (c *CostTime) SetLogCondition(cond ConditionFunc) {
	c.condition = cond
}

// Cost 里面计算消耗时间
func Cost(run func()) {
	defaultCost.Cost(run)
}

// Cost 里面计算消耗时间
func (c *CostTime) Cost(run func()) {

	atomic.AddInt64(&c.loglevel, 1)
	defer func() {
		atomic.AddInt64(&c.loglevel, -1)
	}()

	file, line, funcName := getRuntimeInfo()
	now := time.Now()
	run()
	cost, selcolor := countCostColor(now)
	if c.condition != nil {
		if !c.condition(cost) {
			return
		}
	}
	coststr := fmt.Sprintf(selcolor, cost.Milliseconds())
	var prefix string
	if atomic.LoadInt64(&c.loglevel) > 0 {
		for i := int64(0); i < c.loglevel; i++ {
			for i := int64(0); i < c.loglevel-1; i++ {
				prefix += "  "
			}
			prefix += "┌─ "
		}
		c.costlogNoDate.Printf("%s%s:%d(%s) cost(%s ms)", prefix, file, line, funcName, coststr)
	} else {
		c.costlog.Printf("%s%s:%d(%s) cost(%s ms)", prefix, file, line, funcName, coststr)
	}
}

// CostLog 计算消耗的时间
func CostLog(name string, run func()) {
	defaultCost.CostLog(name, run)
}

// CostLog 计算消耗的时间
func (c *CostTime) CostLog(name string, run func()) {
	atomic.AddInt64(&c.loglevel, 1)
	defer func() {
		atomic.AddInt64(&c.loglevel, -1)
	}()

	file, line, funcName := getRuntimeInfo()

	now := time.Now()
	run()
	cost, selcolor := countCostColor(now)
	if c.condition != nil {
		if !c.condition(cost) {
			return
		}
	}
	coststr := fmt.Sprintf(selcolor, cost.Milliseconds())
	var prefix string
	if atomic.LoadInt64(&c.loglevel) > 0 {
		for i := int64(0); i < c.loglevel; i++ {
			for i := int64(0); i < c.loglevel-1; i++ {
				prefix += "  "
			}
			prefix += "┌─ "
		}
		c.costlogNoDate.Printf("%s%s:%d(%s) cost(%s ms):%s", prefix, file, line, funcName, coststr, name)
	} else {
		c.costlog.Printf("%s%s:%d(%s) cost(%s ms):%s", prefix, file, line, funcName, coststr, name)
	}
}

func countCostColor(now time.Time) (time.Duration, string) {
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
	return cost, selcolor // fmt.Sprintf(selcolor, cost.Milliseconds())
}

func getRuntimeInfo() (file string, line int, funcName string) {
	pc, file, line, _ := runtime.Caller(3)
	funcName = runtime.FuncForPC(pc).Name()

	var i int
	i = strings.LastIndexByte(funcName, '.')
	funcName = funcName[i+1:]

	i = strings.LastIndexByte(file, '/')
	file = file[i+1:]

	return file, line, funcName
}
