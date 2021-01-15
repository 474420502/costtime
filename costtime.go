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

	var i int
	i = strings.LastIndexByte(fname, '/')
	fname = fname[i+1:]

	i = strings.LastIndexByte(file, '/')
	file = file[i+1:]

	now := time.Now()
	run()
	end := time.Now()

	var coststr string
	cost := end.Sub(now)
	switch {
	case cost < time.Millisecond*100:
		coststr = fmt.Sprintf(colors[0], cost)
	case cost < time.Second:
		coststr = fmt.Sprintf(colors[1], cost)
	case cost < time.Second*10:
		coststr = fmt.Sprintf(colors[2], cost)
	default:
		coststr = fmt.Sprintf(colors[3], cost)
	}
	costlog.Printf("%s:%d(%s) cost(%s ms): %s", file, line, fname, coststr, name)
}
