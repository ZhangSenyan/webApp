package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id          int
	Age         int16
	Name        string
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterModel(new(User))

	orm.RegisterDataBase("default", "mysql", "root:zhsy4631@tcp(127.0.0.1:3306)/beeblog?charset=utf8")
}

func main() {
	orm.RunSyncdb()
	o := orm.NewOrm()
	o.Using("default")

	user := new(User)
	user.Age = 30
	user.Name = "slene"

	fmt.Println(o.Insert(&user))
}
