package main

import (
	_ "api/routers"
	"fmt"
	"runtime"

	"github.com/astaxie/beego"
)

func main() {
	cpus := runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	fmt.Println("this is cpu num=", cpus)
	beego.Run()
}
