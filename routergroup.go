package clonegin
import (

    "strings"
    "net/http"
)
type RouterGroup struct {
    Handlers HandlersChain
    BasePath string
    engine *Engine
}
func (group *RouterGroup) Use(middlewares ...HandlerFunc) {
    group.Handlers =append(group.Handlers, middlewares...)
}
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
    return &RouterGroup{

    }
}
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) {

}
func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers  ...HandlerFunc) {
    group.handle(httpMethod, relativePath, handlers)
}
func (group *RouterGroup) POST(relativePath string, handlers ...HandlerFunc) {
    group.handle("POST", relativePath, handlers)
}
func (group *RouterGroup)GET(relativePath string, handlers ...HandlerFunc) {
    group.handle("GET", relativePath, handlers)
}
func (group *RouterGroup) DELETE(relativePath string, handlers ...HandlerFunc) {
    group.handle("DELETE", relativePath, handlers)
}
func (group *RouterGroup) PATCH(relativePath string, handlers ...HandlerFunc) {
    group.handle("PATCH", relativePath, handlers)
}
func (group *RouterGroup) PUT(relativePath string, handlers ...HandlerFunc) {
    group.handle("PUT", relativePath, handlers)
}
func (group *RouterGroup) OPTIONS(relativePath string, handlers ...HandlerFunc) {
    group.handle("OPTIONS", relativePath, handlers)
}
func (group *RouterGroup) HEAD(relativePath string, handlers ...HandlerFunc) {
    group.handle("HEAD", relativePath, handlers)
}
func (group *RouterGroup) Any(relativePath string, handlers ...HandlerFunc) {
    group.handle("GET", relativePath, handlers)
    group.handle("POST", relativePath, handlers)
    group.handle("PUT", relativePath, handlers)
    group.handle("PATCH", relativePath, handlers)
    group.handle("HEAD", relativePath, handlers)
    group.handle("OPTIONS", relativePath, handlers)
    group.handle("DELETE", relativePath, handlers)
    group.handle("CONNECT", relativePath, handlers)
    group.handle("TRACE", relativePath, handlers)
}
func (group *RouterGroup) StaticFile(relativePath, filepath string) {
    if strings.Contains(relativePath, ":")||strings.Contains(relativePath, "*") {
        panic("URL parameters can not be used when serving a static file")
    }
    handler := func(c *Context) {
        c.File(filepath)
    }
    group.GET(relativePath, handler)
    group.HEAD(relativePath, handler)
}

func (group *RouterGroup) Static(relativePath, root string) {
    group.StaticFile(relativePath, Dir(root, false))
}
func (group *RouterGroup) StaticFS(relativePath string, fs http.FileSystem){
    if strings.Contains(relativePath,":")||strings.Contains(relativePath,"*"){
        panic("URL parameters can not be used when serving a static floder")
    }
//    handler :=group.
}

func (group *RouterGroup) createStaticHandler(relativePath string,fs http.FileSystem) HandlerFunc{

}

func (group *RouterGroup) combineHandlers(handlers HandlersChain) HandlersChain{

}
func (group *RouterGroup) calculateAbsolutePath (relativePath  string) string{
//    return
}