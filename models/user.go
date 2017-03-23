package models

import (
	//	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

//Model Struct
type User struct {
	Id       int
	Name     string `orm:"size(20)"`
	PassWord string `orm:"size(32)"`
}

func init() {
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)
}
func (this *User) Add(name, password string) error {
	o := orm.NewOrm()
	user := User{Name: name, PassWord: password}
	_, err := o.Insert(&user)
	return err
}

func (this *User) Get(name, password string) (User, error) {
	o := orm.NewOrm()
	var users []User
	_, err := o.Raw("select id,name,pass_word as password from user where name='" + name + "' and pass_word = '" + password + "'").QueryRows(&users)
	user := User{}
	if len(users) > 0 {
		user = users[0]
	}
	return user, err
}
