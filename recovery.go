package clonegin
import (

    "io"
    "log"
)
var (
    dunno = []byte("???")
    centerDot = []byte("·")
    dot = []byte(".")
    slash = []byte("/")
)
func Recovery() HandlerFunc {

}
func RecoveryWithWriter(out io.Writer) HandlerFunc {
    var logger *log.Logger
    if out !=nil {
        logger =log.New(out, "", log.LstdFlags)
    }
    return func(c *Context){
        defer func(){

        }()

    }
}
