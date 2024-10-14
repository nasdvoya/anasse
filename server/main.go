package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	IntializeDatabase()
	defer DB.Close()

	r := gin.Default()
	// register templates directory into the engine
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		partners := ReadPartnerList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"partners": partners,
		})
	})

	r.POST("/partners", func(c *gin.Context) {
		nome := c.PostForm("nome")
		morada := c.PostForm("morada")

		id, _ := CreatePartner(Partner{
			Nome:   nome,
			Morada: morada,
		})

		c.HTML(http.StatusOK, "partner.html", gin.H{
			"id":     id,
			"nome":   nome,
			"morada": morada,
		})
	})

	r.Run(":8080")
}
