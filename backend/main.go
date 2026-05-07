package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()

	router.LoadHTMLGlob("../frontend/*")

	router.GET("/api/dog", func(c *gin.Context){

		pictureUrl, err := GetPicture()
		if(err != nil){
			log.Printf("critical error : %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		fact, err := GetFact()
		if(err != nil){
			log.Printf("critical error : %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "internal server error",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"url": pictureUrl,
			"fact": fact,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", nil)})

router.Run(":8080")
}