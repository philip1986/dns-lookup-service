package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.Static("/ui", "./frontend/dist")
    router.Run("localhost:8080")
}

