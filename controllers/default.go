/**
 * Copyright (c) 2019 - xushuhui
 * Author: xushuhui
 * 微信公众号: 互联网工程师
 * Email: xushuhui@qq.com
 * 博客: https://www.phpst.cn
 */
package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["json"] = "123@qq.com"
	c.ServeJSON()
}

//func (c *MainController) Post() {
//	var u *user
//	utils.BindJson(c.Ctx.Input.RequestBody, &u)
//	userModel, _ := models.GetUserById(1)
//
//	fmt.Println(userModel)
//	fmt.Println(beego.AppConfig)
//	c.Data["json"] = core.SetData(userModel)
//	c.ServeJSON()
//}
