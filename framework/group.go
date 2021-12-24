package framework


// IGroup 代表前缀分组
type IGroup interface {
	Get(string, ...ControllerHandler)
	Post(string, ...ControllerHandler)

	Group(string) IGroup
	Use(middlewares ...ControllerHandler)
}



// Group struct 实现了IGroup
type Group struct {
	core   *Core
	prefix string
	parent *Group
	middlewares []ControllerHandler // 存放中间件
}

// 初始化Group
func NewGroup(core *Core, prefix string) *Group {
	return &Group{
		core:   core,
		parent: nil,
		prefix: prefix,
	}
}


// 注册中间件
func (g *Group) Use(middlewares ...ControllerHandler) {
	g.middlewares = append(g.middlewares, middlewares...)
}

func (g *Group) getAbsolutePrefix() string {
	if g.parent == nil {
		return g.prefix
	}
	return g.parent.getAbsolutePrefix() + g.prefix
}
func (g *Group) getMiddlewares() []ControllerHandler {
	if g.parent == nil {
		return g.middlewares
	}
	return append(g.parent.getMiddlewares(), g.middlewares...)
}

// 实现Get方法
func (g *Group) Get(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Get(uri, allHandlers...)
}
// 实现Post方法
func (g *Group) Post(uri string, handlers ...ControllerHandler) {
	uri = g.getAbsolutePrefix() + uri
	allHandlers := append(g.getMiddlewares(), handlers...)
	g.core.Post(uri, allHandlers...)
}


// 从core中初始化这个Group
func (g *Group) Group(uri string) IGroup {
	cgroup := NewGroup(g.core, uri)
	cgroup.parent = g
	return cgroup
}