package context

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerFunc 提供处理函数模板
type HandlerFunc func(*Context)

type Context struct {
	resp http.ResponseWriter
	req  *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		resp:   resp,
		req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// req
func (c *Context) Query(key string) string {
	return c.req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.req.FormValue(key)
}

// resp
func (c *Context) SetHeader(key, value string) {
	c.resp.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.resp.WriteHeader(code)
}

func (c *Context) String(code int, message string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.resp.Write([]byte(fmt.Sprintf(message, values...)))
}

func (c *Context) JSON(code int, obj ...interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.resp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.resp, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.resp.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.resp.Write([]byte(html))
}
