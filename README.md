# Gin Access Limit Middleware

[![Go Report Card](https://goreportcard.com/badge/github.com/bu/gin-access-limit)](https://goreportcard.com/report/github.com/bu/gin-access-limit)
![Build Status](https://github.com/bu/gin-access-limit/workflows/build/badge.svg)
[![Documentation](https://godoc.org/github.com/bu/gin-access-limit?status.svg)](http://godoc.org/github.com/bu/gin-access-limit)

A [Gin web framework](https://github.com/gin-gonic/gin) middleware for IP restriction by specifying CIDR notations.

## Usage

```go

package main

import (
    gin "github.com/gin-gonic/gin"
    limit "github.com/bu/gin-access-limit"
)

func main() {
    // create a Gin engine
    r := gin.Default()

    // this API is only accessible from Docker containers
    r.Use(limit.CIDR("172.18.0.0/16"))

    // if need to specify serveral range of allowed sources, use comma to concatenate them
    // r.Use(limit.CIDR("172.18.0.0/16, 127.0.0.1/32"))

    // routes
    r.GET("/", func (c *gin.Context) {
        c.String(200, "pong")
    })

    // listen to request
    r.Run(":8080")
}

```
