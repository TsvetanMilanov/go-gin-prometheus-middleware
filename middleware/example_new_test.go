package middleware_test

import (
	"github.com/TsvetanMilanov/go-gin-prometheus-middleware/middleware"
	"github.com/gin-gonic/gin"
)

func ExampleNew() {
	middleware := middleware.New()

	gin.Default().Use(middleware)
}
