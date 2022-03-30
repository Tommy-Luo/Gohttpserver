package framework

// 注册路由规则
func RegisterRouter(core *Core) {
	// 需求1+2:HTTP方法+静态路由匹配
	core.Post("/user/GetPassWordAPI", Test3(), GetInfomationAPI)

	// 批量通用前缀
	subjectApi := core.CreateGroup("/subject")
	{
		subjectApi.Use(Test3())
		// 动态路由

		subjectApi.Get("/:id", Test3(), SubjectListController)
		subjectApi.Get("/list/all", SubjectListController)

		subjectInnerApi := subjectApi.Group("/info")
		{
			subjectInnerApi.Get("/name", SubjectListController)
		}
	}
}
