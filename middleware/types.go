package middleware

import "github.com/gin-gonic/gin"

// Blacklister is a function which will be used to decide which
// requests should not be included in the metrics.
type Blacklister = func(c *gin.Context) bool
