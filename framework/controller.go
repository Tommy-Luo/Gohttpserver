package framework

import (
	"../orm"
	"encoding/json"
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

func GetInfomationAPI(c *Context) error{

	var  ret interface{}
	val, _ := c.QueryInt("id",0)
	byt := orm.SearchData(orm.DB, uint(val))

	_ = json.Unmarshal(byt, &ret)
	c.Json(ret)

	return nil
}
