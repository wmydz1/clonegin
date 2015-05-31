package clonegin
import (

    "net/http"
    "os"
)
type (
onlyfileFS struct {
    fs http.FileSystem
}
neuteredReaddirFile struct {
    http.File
}
)
func Dir(root string, listDirectoy bool) http.FileSystem {
    fs := http.Dir(root)
    if listDirectoy {
        return fs
    }else {
        return &onlyfileFS{fs}
    }
}
func (fs onlyfileFS) Open(name string) (http.File, error) {
    f, err := fs.fs.Open(name)
    if err !=nil {
        return nil, err
    }
    return neuteredReaddirFile{f}, nil
}
func (f neuteredReaddirFile) Readdir(count int) ([]os.FileInfo, error) {
    return nil, nil
}
