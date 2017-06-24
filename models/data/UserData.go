package data

import (
	"fmt"

	"api/pool/redis"
)

type UserData struct {
}

func (this *UserData) GetUserInfos(uids ...int) {
	redis.Get("myTest")
	redis.Set("myTest2", "this my test redis info", 60*1000)
	fmt.Println("this is user data,set=")
}
