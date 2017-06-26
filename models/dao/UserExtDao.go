package dao

import (
	"fmt"

	"api/pool/mgo"

	"gopkg.in/mgo.v2/bson"
)

const (
	UserExtCol string = "userExt"
)

type UserExt struct {
	Level int8
	Desc  string
}

func (this *UserExt) AddExtInfo(level int8, desc string) {
	fmt.Println("add infos=", &UserExt{18, "I like football"})
	mgo.Adds(UserExtCol, &UserExt{Level: 18, Desc: "I like football"})
}

func (this *UserExt) GetInfoLevel(level int) {
	result := &UserExt{}
	err := mgo.FindOne(UserExtCol, bson.M{"level": level}, result)
	fmt.Println(" query data=", result.Desc, "err", err)
}

func (this *UserExt) GetListByLevel(level int) {
	//	users := []UserExt{}
	var users []UserExt
	err := mgo.Query(UserExtCol, bson.M{"level": level}, &users)

	for key, value := range users {
		fmt.Println(key, value)
	}
	fmt.Println(" query data=", users, "err", err)
}
