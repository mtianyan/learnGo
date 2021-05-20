package apiprac

import (
	"fmt"
	"goSnippets/logger"
	"log"
	"reflect"
	"runtime"
	"time"
)

// 计算函数的运行时间
func calExecTime(fn interface{}) {
	if reflect.Func == reflect.TypeOf(fn).Kind() {
		start := time.Now()
		reflect.ValueOf(fn).Call(nil) // call function
		duration := time.Since(start)
		log.Printf("Time duration of %v: %v\n",
			runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), duration)
		return
	}
	panic(fmt.Sprintf("%v is not callable", fn))
}

func init() {
	logger.DefaultLogger.Log()
	calExecTime(test)
}

func test() {
	fmt.Println("test func starts...")
	time.Sleep(time.Second)
	fmt.Println("test func ends...")
}
