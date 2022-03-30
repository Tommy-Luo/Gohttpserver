package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/cast"
	"io/ioutil"
)

// 代表请求包含的方法
type IRequest interface {
	QueryInt(key string, def int) (int, bool)
	// json body
	BindJson(obj interface{}) IRequest


}


type recordConfig struct {
	Id uint `json:"int"`
}

func (ctx *Context) QueryAll() map[string][]string {
	if ctx.request != nil {
		return map[string][]string(ctx.request.URL.Query())
	}
	return map[string][]string{}
}

func (ctx *Context) QueryInt(key string, def int) (int, bool) {
	params := ctx.QueryAll()
	if vals, ok := params[key]; ok {
		if len(vals) > 0 {
			// 使用cast库将string转换为Int
			return cast.ToInt(vals[0]), true
		}
	}
	return def, false
}


// 将 body 文本解析到 obj 结构体中
func (ctx *Context) BindJson(obj interface{}) error {
	if ctx.request != nil {
		// 读取文本
		body, err := ioutil.ReadAll(ctx.request.Body)
		if err != nil {
			return err
		}
		// 重新填充 request.Body，为后续的逻辑二次读取做准备
		ctx.request.Body = ioutil.NopCloser(bytes.NewBuffer(body))

		// 解析到 obj 结构体中
		err = json.Unmarshal(body, obj)
		if err != nil {
			return err
		}
	} else {
		return errors.New("ctx.request empty")
	}
	return nil
}