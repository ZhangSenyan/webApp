package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
const(
_DB_NAME="root:zhsy4631@tcp(127.0.0.1:3306)/beeblog?charset=utf8"
_Mysql_DRIVER="mysql"
)
// 格式化代码 ctrl + alt + L
type Category struct {
	// GO 根据首字母大小写判断是否public
	Id              int64
	Title           string    //默认255
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 表示一篇文章
type Topic struct {
	Id              int64
	Uid             int64  //用户ID
	Title           string
	Content         string  `orm:"size(5000)"` //文章正文
	Attachment      string					   //附件
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`    //最后更新时间
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

func RegisterDB(){

	orm.RegisterModel(new(Category),new(Topic))
	orm.RegisterDriver(_Mysql_DRIVER,orm.DRMySQL)
	orm.RegisterDataBase("default",_Mysql_DRIVER,_DB_NAME,10)

}