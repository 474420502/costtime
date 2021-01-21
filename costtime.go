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

// CostTime 消费时间基本结构
type CostTime struct {
	logdeep  int64
	logfirst int64
	skip     int

	eventCost EventFunc
	condition ConditionFunc

	logprefix string

	costlog       *log.Logger
	costlogNoDate *log.Logger
}

// New 创建一个新Cost. 每个协程里都要创建独立一个
func New(fname string) *CostTime {
	c := &CostTime{}
	c.logdeep = -1
	c.skip = 2

	file := func() *os.File {
		f, err := os.OpenFile(fmt.Sprintf("%s/%s.log", logDirectory, fname), os.O_CREATE|os.O_RDWR|os.O_SYNC, 0666)
		if err != nil {
			log.Panic(err)
		}
		return f
	}()

	c.costlog = func() *log.Logger {
		l := log.New(file, "", log.Ldate|log.Ltime)
		return l
	}()

	c.costlogNoDate = func() *log.Logger {
		l := log.New(file, "", log.Ltime)
		return l
	}()
	return c
}

// ConditionFunc 日志输出条件判断函数
type ConditionFunc func(cost time.Duration) bool

// EventFunc 日志输出条件判断函数
type EventFunc func(name string, cost time.Duration)

// var c.loglevel int64 = -1
// var colors = []string{"\033[32m%v\033[0m", "\033[34m%v\033[0m", "\033[31m%v\033[0m", "\033[31m\033[05m%v\033[0m"}

type color struct {
	level    int64
	value    time.Duration
	colorstr string
}

var colorlevels []*color = []*color{
	{
		level:    0,
		colorstr: "\033[32m%v\033[0m", //绿色
		value:    time.Millisecond * 100,
	},
	{
		level:    1,
		colorstr: "\033[36m%v\033[0m", //天蓝色
		value:    time.Millisecond * 500,
	},
	{
		level:    2,
		colorstr: "\033[34m%v\033[0m", //蓝色
		value:    time.Millisecond * 1000,
	},
	{
		level:    3,
		colorstr: "\033[33m%v\033[0m", //黄色
		value:    time.Millisecond * 2000,
	},
	{
		level:    4,
		colorstr: "\033[33m\033[05m%v\033[0m", //黄色闪烁
		value:    time.Millisecond * 4000,
	},

	{
		level:    5,
		colorstr: "\033[31m%v\033[0m", //红色
		value:    time.Millisecond * 8000,
	},

	{
		level:    6,
		colorstr: "\033[31m\033[05m%v\033[0m", //红色闪烁
		value:    time.Millisecond * 16000,
	},
}

func checkLevel(t time.Duration) *color {
	for _, c := range colorlevels {
		if t <= c.value {
			return c
		}
	}
	return colorlevels[len(colorlevels)-1]
}

// SetEeventCost 设置输出cost事件. 可以做邮件通知. 钉釘等办公通知. 只触发CostLog(). Cost()不触发
func (c *CostTime) SetEeventCost(event EventFunc) {
	c.eventCost = event
}

// SetLogCondition 设置输出cost条件
func (c *CostTime) SetLogCondition(cond ConditionFunc) {
	c.condition = cond
}

// Cost 里面计算消耗时间
func (c *CostTime) Cost(run func()) {

	atomic.AddInt64(&c.logdeep, 1)
	defer func() {
		atomic.AddInt64(&c.logdeep, -1)
	}()

	file, line, funcName := c.getRuntimeInfo()
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
	if atomic.LoadInt64(&c.logdeep) > 0 {
		lf := atomic.AddInt64(&c.logfirst, 1)
		var logprefix string
		if lf == 1 {
			logprefix = "┌─"
		} else {
			logprefix = "├─"
		}

		for i := int64(0); i < c.logdeep-1; i++ {
			logprefix += "──"
		}
		logprefix += " "

		c.costlogNoDate.SetPrefix(fmt.Sprintf(selcolor, logprefix))
		prefix += fmt.Sprintf(selcolor, "● ")
		c.costlogNoDate.Printf("%s%s:%d(%s %s ms)", prefix, file, line, funcName, coststr)
	} else {
		c.costlog.Printf("%s%s:%d(%s %s ms)", prefix, file, line, funcName, coststr)
		atomic.StoreInt64(&c.logfirst, 0)
	}
}

// CostLog 计算消耗的时间
func (c *CostTime) CostLog(name string, run func()) {
	atomic.AddInt64(&c.logdeep, 1)
	defer func() {
		atomic.AddInt64(&c.logdeep, -1)
	}()

	file, line, funcName := c.getRuntimeInfo()

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
	if atomic.LoadInt64(&c.logdeep) > 0 {
		lf := atomic.AddInt64(&c.logfirst, 1)
		var logprefix string
		if lf == 1 {
			logprefix = "┌─"
		} else {
			logprefix = "├─"
		}

		for i := int64(0); i < c.logdeep-1; i++ {
			logprefix += "──"
		}
		logprefix += " "

		c.costlogNoDate.SetPrefix(fmt.Sprintf(selcolor, logprefix))
		prefix += fmt.Sprintf(selcolor, "● ")
		c.costlogNoDate.Printf("%s%s:%d(%s %s ms): %s", prefix, file, line, funcName, coststr, name)
	} else {
		c.costlog.Printf("%s%s:%d(%s %s ms): %s", prefix, file, line, funcName, coststr, name)
		atomic.StoreInt64(&c.logfirst, 0)
	}

	if c.eventCost != nil {
		c.eventCost(name, cost)
	}
}

func countCostColor(now time.Time) (time.Duration, string) {
	cost := time.Now().Sub(now)
	return cost, checkLevel(cost).colorstr // fmt.Sprintf(selcolor, cost.Milliseconds())
}

func (c *CostTime) getRuntimeInfo() (file string, line int, funcName string) {

	pc, file, line, _ := runtime.Caller(int(c.skip))
	funcName = runtime.FuncForPC(pc).Name()

	var i int
	i = strings.LastIndexByte(funcName, '.')
	funcName = funcName[i+1:]

	i = strings.LastIndexByte(file, '/')
	file = file[i+1:]

	return file, line, funcName
}
