package router

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// SetupRouter sets up the Gin engine with routes
func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Set up a route for the /hello endpoint
    r.GET("/hello", func(c *gin.Context) {
        c.String(http.StatusOK, "Hello World")
    })

    return r
}