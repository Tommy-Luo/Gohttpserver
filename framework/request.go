package framework

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
)

// 代表请求包含的方法
type IRequest interface {

	// json body
	BindJson(obj interface{}) error


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
