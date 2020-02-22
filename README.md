# go-gin-prometheus-middleware
Golang middleware for the gin framework

[![GoDoc](https://godoc.org/github.com/TsvetanMilanov/go-gin-prometheus-middleware/middleware?status.svg)](https://godoc.org/github.com/TsvetanMilanov/go-gin-prometheus-middleware/middleware)
![Go](https://github.com/TsvetanMilanov/go-gin-prometheus-middleware/workflows/Go/badge.svg?branch=master)

## Quick Start
```Go
middleware := middleware.New()

gin.Default().Use(middleware)
```

### Example output with default options
```
# HELP http_request_duration_seconds Duration summary of http responses labeled with: method,path,status_code
# TYPE http_request_duration_seconds histogram
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="0.1"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="0.25"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="0.5"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="1"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="2"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="5"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="10"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="15"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="20"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="30"} 1
http_request_duration_seconds_bucket{method="GET",path="/error",status_code="500",le="+Inf"} 1
http_request_duration_seconds_sum{method="GET",path="/error",status_code="500"} 2.0226e-05
http_request_duration_seconds_count{method="GET",path="/error",status_code="500"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="0.1"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="0.25"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="0.5"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="1"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="2"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="5"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="10"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="15"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="20"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="30"} 1
http_request_duration_seconds_bucket{method="GET",path="/ok",status_code="200",le="+Inf"} 1
http_request_duration_seconds_sum{method="GET",path="/ok",status_code="200"} 8.6441e-05
http_request_duration_seconds_count{method="GET",path="/ok",status_code="200"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="0.1"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="0.25"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="0.5"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="1"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="2"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="5"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="10"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="15"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="20"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="30"} 1
http_request_duration_seconds_bucket{method="GET",path="/param/:value",status_code="200",le="+Inf"} 1
http_request_duration_seconds_sum{method="GET",path="/param/:value",status_code="200"} 2.3551e-05
http_request_duration_seconds_count{method="GET",path="/param/:value",status_code="200"} 1
```
