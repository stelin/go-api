package dao

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(User))
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (this *User) Add(name string, age int) {
	o := orm.NewOrm()
	var user User
	user.Name = name
	user.Age = age

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	}
	fmt.Println("insert id=", id)
}

func (this *User) GetById(id int) {
	o := orm.NewOrm()
	user := User{Id: id}

	err := o.Read(&user)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(user.Id, user.Name, user.Age)
	}
}
