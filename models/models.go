package models

import (
	"time"
	
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Category struct {
	Id				int64
	Title			string
	Created			time.Time `orm:"index"`//创建时间
	Views			int64	  `orm:"index"`//查看次数
	TopicTime		time.Time `orm:"index"`
	TopicCount  		int64	
	TopicLastUserId  int64
}

type Topic struct {
	Id				int64
	Uid				int64
	Title			string
	Content			string 	 `orm:"size(5000)"`
	Attachment		string
	Created			time.Time `orm:"index"`
	Updated			time.Time `orm:"index"`
	Views			int64	  `orm:"index"`
	Author			string
	ReplyTime		time.Time  `orm:"index"`
	ReplyCount		int64
	ReplyLastUserId  int64
}

func Registry() {
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDataBase("default","mysql","root:mysql@/test?charset=utf8",10)
	orm.RunSyncdb("default",false,true)
}