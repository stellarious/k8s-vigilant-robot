package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	name, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, fmt.Sprintf("⭐️ VER3: Hello from %s! ⭐️", name))
	})
	r.Run()
}
