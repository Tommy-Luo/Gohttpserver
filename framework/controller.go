package framework

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


