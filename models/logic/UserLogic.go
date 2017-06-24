package logic

import (
	"api/models/data"

	"fmt"
)

type UserLogic struct {
}

func (this *UserLogic) GetUserInfos(uids ...int) {
	userData := data.UserData{}
	userData.GetUserInfos(3, 4)
	fmt.Println("this user infos logic uids", uids)
}
