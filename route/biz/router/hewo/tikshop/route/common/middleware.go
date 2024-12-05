package common

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Cors() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		method := c.Request.Method()

		c.Header("Access-Control-Allow-Origin")
	}
}
