# go-gin-prometheus-middleware
Golang middleware for the gin framework

[![GoDoc](https://godoc.org/github.com/TsvetanMilanov/go-gin-prometheus-middleware/middleware?status.svg)](https://godoc.org/github.com/TsvetanMilanov/go-gin-prometheus-middleware/middleware)

## Quick Start
```Go
middleware := middleware.New()

gin.Default().Use(middleware)
```
