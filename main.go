package lib

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/yudirada/golongan-lib/p"
)

func run() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	r := gin.New()
	r.Use(gin.Logger())
	//router.LoadHTMLGlob("templates/*.tmpl.html")
	//router.Static("/static", "static")

	//router.GET("/", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.tmpl.html", nil)
	//})

	p.RenderIndex(r)

	r.Run(":" + port)
}
