package go_web

import (
	"go-web/context"
	"go-web/gweb"
	"go-web/middle"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestWeb(t *testing.T) {
	engine := gweb.New()

	engine.Use(middle.Logger())

	engine.Get("/", func(c *context.Context) {
		c.String(200, "URL.Path = %q\n", c.Path)

		// fmt.Fprintf(resp, "URL.Path = %q\n", req.URL.Path)
	})
	v1 := engine.Group("/v1")
	v1.Use(onlyForV1())
	v1.Get("/login", func(c *context.Context) {
		res := make(map[string]string)
		res["username"] = "gweb"
		res["password"] = "123"
		c.JSON(http.StatusOK, res)
	})

	v2 := engine.Group("/v2")
	v2.Get("/user", func(c *context.Context) {
		res := make(map[string]string)
		res["username"] = "gweb --user"
		res["password"] = "123 --user"
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

func onlyForV1() context.HandlerFunc {
	return func(c *context.Context) {
		// Start timer
		t := time.Now()

		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v1", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
