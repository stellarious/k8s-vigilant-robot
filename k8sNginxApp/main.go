package main

import (
	"fmt"
	"io"
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
		ctx.JSON(http.StatusOK, fmt.Sprintf("🅝 Hello from %s! 🅝", name))
	})
	r.GET("/nginx", func(ctx *gin.Context) {
		resp, err := http.Get("https://nginx")
		if err != nil {
			ctx.AbortWithStatus(500)
		}
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ctx.AbortWithStatus(500)
		}
		ctx.JSON(http.StatusOK, string(body))
	})
	r.Run()
}
