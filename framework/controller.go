package framework

import (
	"../orm"
	"fmt"
)

type ControllerHandler func(c *Context) error


func SubjectListController(c *Context) error {
	c.Json("ok, SubjectNameController")
	return nil
}


func UserLoginController(c *Context) error {
	// 打印控制器名字
	c.Json("ok, UserLoginController")
	return nil
}

func GetPassWordAPI(c *Context) error{

	ret := orm.SearchData(orm.DB, 1)
	c.Json(ret)
	//orm.UpdateData(orm.DB, obj.(uint))
	return nil
}

func ChangePassWordAPI(c *Context) error{

	var obj interface{}
	c.BindJson(obj)
	fmt.Println(obj.(string))
	//orm.UpdateData(orm.DB, obj.(uint))
	return nil
}