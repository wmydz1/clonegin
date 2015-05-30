package clonegin
import (

    "io"
    "github.com/mattn/go-colorable"
    "os"
)
const ENV_GIN_MODE = "GIN_MODE"
const (
    DebugMode   string = "debug"
    ReleaseMode string = "release"
    TestMode    string = "test"
)
const (
    debugCode = iota
    releaseCode = iota
    testCode = iota
)
var DefaultWriter io.Writer = colorable.NewColorableStdout()
var ginMode int = debugCode
var modeName string = DebugMode

func init() {
    mode := os.Getenv(ENV_GIN_MODE)
    if len(mode) ==0 {
        SetMode(DebugMode)
    }else {
        SetMode(mode)
    }
}
func SetMode(value string) {
    switch value{
        case DebugMode:
        ginMode =debugCode
        case ReleaseMode:
        ginMode=releaseCode
        case TestMode:
        ginMode=testCode
        default:
        panic("gin mode unknown: "+value)
    }
    modeName=value
}
func Mode() string {
    return modeName
}
