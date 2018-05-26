# Gin Access Limit Middleware

[![Go Report Card](https://goreportcard.com/badge/github.com/bu/gin-access-limit)](https://goreportcard.com/report/github.com/bu/gin-access-limit)
[![Build Status](https://travis-ci.org/bu/gin-access-limit.svg?branch=master)](https://travis-ci.org/bu/gin-access-limit)

A [Gin web framework](https://github.com/gin-gonic/gin) middleware for IP access control by specifying CIDR notations.

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

    // routes
    r.GET("/", func (c *gin.Context) {
        c.String(200, "pong")
    })

    // listen to request
    r.Run(":8080")
}

```
