package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Static("assets/css", "./assets/css")
	router.LoadHTMLGlob("templates/index.html")

	router.GET("/hash", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})

	router.POST("/hash/hashvalue", func(ctx *gin.Context) {
		from_html := string(ctx.PostForm("for_gin_num"))
		res := hash(from_html)
		ctx.HTML(200, "index.html", gin.H{"value": from_html, "data": res})
	})

	router.Run(":3000")
}

func hash(s string) []string {
	from_html_list := strings.Split(s, "\r\n")
	hashed := make([]string, 0)
	for _, ss := range from_html_list {
		b := []byte(ss)
		sha256 := sha256.Sum256(b)
		hashed = append(hashed, hex.EncodeToString(sha256[:]))
	}
	return hashed
}
