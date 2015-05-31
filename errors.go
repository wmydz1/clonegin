package clonegin
import (

)
type ErrorType uint64
const (
    ErrorTypeBind    ErrorType = 1 << 63 // used when c.Bind() fails
    ErrorTypeRender  ErrorType = 1 << 62 // used when c.Render() fails
    ErrorTypePrivate ErrorType = 1 << 0
    ErrorTypePublic  ErrorType = 1 << 1

    ErrorTypeAny ErrorType = 1<<64 - 1
    ErrorTypeNu = 2
)
type (
Error struct {
    Err error
    Type ErrorType
    Meta interface{}
}
errorMsgs []*Error

)
//判断 &Error{} 是不是  error 接口的实现
var _ error=&Error{}
