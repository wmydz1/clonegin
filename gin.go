package clonegin
import (

)
const Version = "v1.0rc1"
var default404Body = []byte("404 page not found")
var default405Body = []byte("405 method not allowed")

type(
HandlerFunc func(*Context)
HandlersChain  []HandlerFunc

Engine struct {
    
}

)