package go_web

import (
	"go-web/context"
	"go-web/gweb"
	"net/http"
	"testing"
)

func TestRecover(t *testing.T) {
	r := gweb.NewDefault()

	r.Get("/", func(c *context.Context) {
		c.String(http.StatusOK, "Hello golang web \n")
	})

	r.Get("/panic", func(c *context.Context) {
		names := []string{"golang web"}
		c.String(http.StatusOK, names[100])
	})
	r.Run(":8080")
}
