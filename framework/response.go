package framework


// IResponse 代表返回方法
type IResponse interface {
	// Json 输出
	Json(obj interface{}) IResponse

	// Jsonp 输出
	Jsonp(obj interface{}) IResponse


	// header
	SetHeader(key string, val string) IResponse

	// Cookie
	SetCookie(key string, val string, maxAge int, path, domain string, secure, httpOnly bool) IResponse

	// 设置状态码
	SetStatus(code int) IResponse

	// 设置 200 状态
	SetOkStatus() IResponse
}