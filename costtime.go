package costtime

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"
)

var costlog = func() *log.Logger {
	l := log.New(os.Stderr, "", log.Ldate|log.Ltime)
	return l
}()

var colors = []string{"\033[32m%d\033[0m", "\033[36m%d\033[0m", "\033[34m%d\033[0m", "\033[31m%d\033[0m"}

// CostTime 计算消耗的时间
func CostTime(name string, run func()) {

	pc, file, line, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()

	var i, count int
	i = strings.LastIndexFunc(fname, func(c rune) bool {
		if c == '.' || c == '/' {
			count++
		}
		if count >= 2 {
			return true
		}
		return false
	})
	fname = fname[i+1:]

	i = strings.LastIndexByte(file, '/')
	file = file[i+1:]

	now := time.Now()
	run()
	end := time.Now()

	var coststr string
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
	coststr = fmt.Sprintf(selcolor, cost.Milliseconds())
	costlog.Printf("%s:%d(%s) cost(%s ms): %s", file, line, fname, coststr, name)
}
