package main

import (
	_ "api/pool/mgo"
	_ "api/pool/mysql"
	_ "api/pool/redis"
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
