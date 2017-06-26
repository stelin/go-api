package data

import (
	"fmt"

	"api/models/dao"
	"api/pool/redis"
)

type UserData struct {
}

func (this *UserData) GetUserInfos(uids ...int) {
	redis.Get("myTest")
	redis.Set("myTest2", "this my test redis info", 60*1000)
	fmt.Println("this is user data,set=")
	userDao := &dao.User{}
	userDao.Add("stelin", 18)
	userDao.GetById(790)
	fmt.Println(userDao)

	userExtDao := &dao.UserExt{}
	userExtDao.AddExtInfo(8, "i like football")
	userExtDao.GetInfoLevel(18)
	userExtDao.GetListByLevel(18)
}
