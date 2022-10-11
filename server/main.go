package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("tpls/*")
	r.Static("/static/", "static")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "lix7",
		})
	})
	r.GET("/version", func(c *gin.Context) {
		cmd := exec.Command("git", "rev-parse", "--short", "HEAD")
		version, err := cmd.Output()
		if err != nil {
			log.Printf("version: %s, err: %s, path: %s", version, err, cmd.Path)
			version = []byte("error")
		}
		c.JSON(http.StatusOK, gin.H{
			"version": string(version),
		})
	})
	r.Run("0.0.0.0:80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
