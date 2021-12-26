package error_code

import (
	"fmt"
	"net/http"
)

/**
内部错误的数据结构
*/
type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

var codes = map[int]string{}

func NewError(code int, msg string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = msg
	return &Error{code: code, msg: msg}
}

/**
下面的几个方法有点像 Java 的习惯, 在 Go 里面可能更加习惯于直接访问结构体的成员变量的值;
*/
func (e *Error) Error() string {
	return fmt.Sprintf("错误码：%d, 错误信息:：%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}

func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) Msgf(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := *e
	newError.details = []string{}
	newError.details = append(newError.details, details...)
	return &newError
}

/**
针对一些特定错误码进行状态码的转换
不同的内部错误码对应于不同的 HTTP 状态码
需要将其区分开来，便于客户端以及监控/报警等系统的识别和监听;

也可以用 map 来做映射, 不过那样能实现的逻辑会比较死, 用一个 func 实现的效果相对更加灵活
*/
func (e *Error) StatusCode() int {
	switch e.Code() {
	case Success.Code():
		return http.StatusOK
	case ServerError.Code():
		return http.StatusInternalServerError
	case InvalidParams.Code():
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.Code():
		// go 语言的 case 自带 break 语句
		// 如果不用 fallthrough 这个关键字, 到这里没有语句可以执行, 就会跳出 switch
		fallthrough
	case UnauthorizedTokenError.Code():
		fallthrough
	case UnauthorizedTokenGenerate.Code():
		fallthrough
	case UnauthorizedTokenTimeout.Code():
		return http.StatusUnauthorized
	case TooManyRequests.Code():
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
