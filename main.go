package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Static("/assets", "./dist/assets")

	r.LoadHTMLGlob("dist/index.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Go + Vite + React + TS",
		})
	})

	r.Run()
}
