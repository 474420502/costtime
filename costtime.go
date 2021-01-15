package costtime

import (
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

var flag = struct {
	IsShowLongFileName bool
}{
	IsShowLongFileName: false,
}

// CostTime 计算消耗的时间
func CostTime(name string, run func()) {

	pc, file, line, _ := runtime.Caller(1)
	fname := runtime.FuncForPC(pc).Name()

	if !flag.IsShowLongFileName {
		i := strings.LastIndexByte(file, '/')
		file = file[i+1:]
	}

	now := time.Now()
	costlog.Printf("%s:%d(%s %s): start %s", file, line, fname, name, now.Local().Format("2006-01-02 15:04:05"))
	run()
	end := time.Now()
	costlog.Printf("%s:%d(%s %s): end %s, cost: %d ms", file, line, fname, name, end.Local().Format("2006-01-02 15:04:05"), end.Sub(now).Milliseconds())
}
