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
	run()
	end := time.Now()
	costlog.Printf("%s:%d(%s) costtime(%d ms): %s", file, line, fname, end.Sub(now).Milliseconds(), name)
}
