package context

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HandlerFunc 提供处理函数模板
type HandlerFunc func(*Context)

type Context struct {
	Resp http.ResponseWriter
	Req  *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int

	// middleware
	Handlers []HandlerFunc
	index    int
}

func NewContext(resp http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Resp:   resp,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
		index:  -1,
	}
}

// req
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// resp
func (c *Context) SetHeader(key, value string) {
	c.Resp.Header().Set(key, value)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Resp.WriteHeader(code)
}

func (c *Context) String(code int, message string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Resp.Write([]byte(fmt.Sprintf(message, values...)))
}

func (c *Context) JSON(code int, obj ...interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Resp)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Resp, err.Error(), 500)
	}
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Resp.Write(data)
}

func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Resp.Write([]byte(html))
}

func (c *Context) Next() {
	c.index++
	s := len(c.Handlers)
	for ; c.index < s; c.index++ {
		c.Handlers[c.index](c)
	}
}
func (c *Context) Fail(code int, err string) {
	c.index = len(c.Handlers)
	c.JSON(code, "message:", err)
}
