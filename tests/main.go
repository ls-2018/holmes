package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

func main() {
	go func() {
		for {
			var m = make(map[string]string, 1073741824)
			for i := 0; i < 1000; i++ {
				m[fmt.Sprint(i)] = fmt.Sprint(i)
			}
			_ = m
		}
	}()
	time.Sleep(time.Second)
	//pprof.Lookup("heap").WriteTo(os.Stdout, 1)
	//pprof.Lookup("heap").WriteTo(os.Stdout, 0)
	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)
	fmt.Println("------------")
	pprof.Lookup("threadcreate").WriteTo(os.Stdout, 1)
	runtime.LockOSThread()
	//pprof.Lookup("goroutine").WriteTo(os.Stdout, 0)

	//file, _ := os.OpenFile("/tmp/threadcreate.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	//pprof.Lookup("threadcreate").WriteTo(file, 1)
	//pprof.Lookup("threadcreate").WriteTo(file, 0)
	//
	//file, _ = os.OpenFile("/tmp/cpu.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	//pprof.StartCPUProfile(file)
	//time.Sleep(time.Second)
	//pprof.StopCPUProfile()
	// go tool pprof -http=:8081 goroutine.20230225172551.190.log
}
