package clonegin
import (

    "net/http"
    "encoding/xml"
)
func WrapF(f http.HandlerFunc) HandlerFunc {
    return func(c *Context) {
        f(c.Writer, c.Request)
    }
}
func WrapH(h http.Handler) HandlerFunc {
    return func(c *Context) {
        h.ServeHTTP(c.Writer, c.Request)
    }
}
type H map[string]interface{}

func (h H) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

    return nil
}
