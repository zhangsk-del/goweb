package context

import "net/http"

type Context struct {
	resp http.ResponseWriter
	req  *http.Request
}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		resp: resp,
		req:  req,
	}
}
