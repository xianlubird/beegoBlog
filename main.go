package main

import (
	_ "github.com/xianlubird/beegoBlog/routers"
	"github.com/astaxie/beego"
	"github.com/xianlubird/beegoBlog/models"
)

func init() {
	models.Registry()
}

func main() {
	beego.Run()
}

