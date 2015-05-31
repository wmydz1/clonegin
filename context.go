package clonegin
import (


    "math"
    "net/http"
)
const (
    MIMEJSON = "application/json"
    MIMEHTML = "text/html"
    MIMEXML = "application/xml"
    MIMEXML2 = "text/xml"
    MIMEPlain = "text/plain"
    MIMEPOSTForm = "application/x-www-form-urlencoded"
    MIMEMultipartPOSTForm = "multipart/form-data"
)
const AbortIndex = math.MaxInt8/2
type Context struct {
    writermem responseWriter
    Request *http.Request
    Writer ResponseWriter

    Params Params
    handlers HandlersChain
    index int8

    engine *Engine
    Keys map[string]interface{}
    Errors   errorMsgs
    Accepted []string
}

func (c *Context) File(filepath string) {
    http.ServeFile(c.Writer, c.Request, filepath)
}