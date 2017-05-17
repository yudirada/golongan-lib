package p

import "github.com/gin-gonic/gin"

// RenderIndex contains router for index
func RenderIndex(r *gin.Engine) {
	r.GET("/", PageIndex)
}

// PageIndex render Index page
func PageIndex(c *gin.Context) {
	c.Writer.Write([]byte("TEST3"))
}
