package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)
func main() {
    fmt.Println("Hello, World!")

    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "OK!!",
        })
    })

    r.Run(":8000")

}
