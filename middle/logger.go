package middle

import (
	"go-web/context"
	"log"
	"time"
)

func Logger() context.HandlerFunc {
	return func(c *context.Context) {
		// Start timer
		t := time.Now()
		// Process request
		c.Next()
		// Calculate resolution time
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
