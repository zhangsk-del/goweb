package go_web

import (
	"go-web/context"
	"go-web/gweb"
	"net/http"
	"testing"
)

func TestWeb(t *testing.T) {
	engine := gweb.New()

	engine.Get("/", func(c *context.Context) {
		c.String(200, "URL.Path = %q\n", c.Path)

		// fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
	})

	engine.Get("/login", func(c *context.Context) {
		res := make(map[string]string)
		res["username"] = "gweb"
		res["password"] = "123"
		c.JSON(http.StatusOK, res)
	})

	engine.Post("/login", func(c *context.Context) {
		res := make(map[string]string)
		res["username"] = "gweb"
		res["password"] = "123"
		c.JSON(http.StatusOK, res)
	})

	engine.Run(":8080")
}
