package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

func ControllerTest1(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
        "data":   "111",
    })
}

func ControllerTest2(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "status": "ok",
        "data":   "222",
    })
}