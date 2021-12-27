package framework

import (
	"log"
	"net/http"
	"strings"
)


type Core struct {
	router map[string]*Tree // all routers
	middlewares []ControllerHandler // 从core这边设置的中间件
}

func (c *Core) CreateGroup(prefix string) IGroup {
	return NewGroup(c, prefix)
}

// 初始化core结构
func NewCore() *Core {
	// 初始化路由
	router := map[string]*Tree{}
	router["GET"] = NewTree()
	router["POST"] = NewTree()
	return &Core{router: router}
}



// 注册中间件
func (c *Core) Use(middlewares ...ControllerHandler) {
	c.middlewares = append(c.middlewares, middlewares...)
}



// 匹配GET 方法, 增加路由规则
func (c *Core) Get(url string, handlers ...ControllerHandler) {
	// 将core的middleware 和 handlers结合起来
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["GET"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}


// 匹配POST 方法, 增加路由规则
func (c *Core) Post(url string, handlers ...ControllerHandler) {
	allHandlers := append(c.middlewares, handlers...)
	if err := c.router["POST"].AddRouter(url, allHandlers); err != nil {
		log.Fatal("add router error: ", err)
	}
}


// 匹配路由，如果没有匹配到，返回nil
func (c *Core) FindRouteByRequest(request *http.Request) []ControllerHandler {
	// uri 和 method 全部转换为大写，保证大小写不敏感
	uri := request.URL.Path
	method := request.Method
	upperMethod := strings.ToUpper(method)

	// 查找第一层map
	if methodHandlers, ok := c.router[upperMethod]; ok {
		return methodHandlers.FindHandler(uri)
	}
	return nil
}


func (c *Core) ServeHTTP(response  http.ResponseWriter, request  *http.Request) {
	log.Println("core.serveHTTP star")
	ctx := NewContext(request, response)

	handlers := c.FindRouteByRequest(request)

	if handlers == nil {
		ctx.Json("not found,http.Status: 404")
		return
	}
	ctx.SetHandlers(handlers)
	if err := ctx.Next(); err != nil {
		ctx.Json("inner error,http.Status: 500")
		return
	}
	log.Println("core.serveHTTP end")


}





