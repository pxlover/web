package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webProject/restful"
)

func HandleNotFound(c *gin.Context) {
	restful.Fail(restful.NotFound, c)
}

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			recoverText(r, c)
			c.Abort()
		}
	}()
	c.Next()
}

func recoverText(r interface{}, c *gin.Context) {
	switch v := r.(type) {
	case error:
		restful.Fail(v, c)
	default:
		restful.RenderJson(restful.ServerErr, nil, fmt.Sprintf("%v", r), c)
	}
}

func Cors(c *gin.Context) {
	r := c.Request
	defer func() {
		c.Next()
	}()

	origin := r.Header.Get("Origin")
	if origin == "" {
		return
	}

	c.Header("Vary", "Origin")
	c.Header("Vary", "Access-Control-Request-Method")
	c.Header("Vary", "Access-Control-Request-Headers")
	c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	c.Header("Access-Control-Allow-Origin", origin)
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Access-Control-Allow-Credentials", strconv.FormatBool(restful.DefaultOptions.AllowCredentials))
	c.Header("Access-Control-Max-Age", strconv.Itoa(restful.DefaultOptions.MaxAge))

	if r.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
	}
}