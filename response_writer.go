package clonegin
import (

    "net/http"
    "net"
    "bufio"
)
const (
    noWritten = -1
    defaultStatus = 200
)
type (
ResponseWriter interface {
    http.ResponseWriter
    http.Hijacker
    http.Flusher
    http.CloseNotifier

    Status() int
    Size() int
    Written() bool
    WriteHeaderNow()
}
responseWriter struct {
    http.ResponseWriter
    size int
    status int
}
)
func (w *responseWriter) reset(writer http.ResponseWriter) {
    w.ResponseWriter=writer
    w.size=noWritten
    w.status=defaultStatus
}
func (w *responseWriter) WriteHeader(code int) {
    if code>0&&w.status!=code {
         if w.Written(){
             debugP
         }
    }
}
func (w *responseWriter) WriteHeaderNow() {

}
func (w *responseWriter) Write(data []byte) (n int, err error) {
    w.WriteHeaderNow()
    n, err =w.ResponseWriter.Write(data)
    w.size+=n
    return
}
func (w *responseWriter) Status() int {
    return w.status
}
func (w responseWriter) Size() int {
    return w.size
}
func (w *responseWriter)  Written() bool {
    return w.size!=noWritten
}
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
    if w.size<0 {
        w.size=0
    }
    return w.ResponseWriter.(http.Hijacker).Hijack()
}
func (w *responseWriter) CloseNotify() <-chan bool {
    return w.ResponseWriter.(http.CloseNotifier).CloseNotify()
}
func (w *responseWriter) Flush() {
    w.ResponseWriter.(http.Flusher).Flush()
}
