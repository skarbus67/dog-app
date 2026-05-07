package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"dog-app/clients/dogfact"
	"dog-app/clients/dogpic"
)

type response struct{
		res string
		err error
}

func main(){
	router := gin.Default()

	router.LoadHTMLGlob("../frontend/*")

	router.GET("/api/dog", func(c *gin.Context){

		ctx := c.Request.Context()

		client := &http.Client{}

		picClient := dogpic.NewClient(client)
		factClient := dogfact.NewClient(client)
		

		chPic := make (chan response, 1)
		chFact := make (chan response, 1)

		go func(){
			pictureUrl, err := picClient.GetPicture(ctx)
			chPic <- response{res: pictureUrl, err: err}
		}()

		go func(){
			fact, err := factClient.GetFact(ctx)
			chFact <- response{res: fact, err: err}
		}()

		var fact string
        var picURL string
		for i := 0; i < 2; i++ {
    		select {
				case f := <-chFact:
        			if f.err != nil {
            			log.Printf("Fact error: %v", f.err)
            			c.JSON(http.StatusInternalServerError, gin.H{"error": "fact failed"})
            			return
        			}
					fact = f.res
    			case p := <-chPic:
        			if p.err != nil {
            			log.Printf("Pic error: %v", p.err)
            			c.JSON(http.StatusInternalServerError, gin.H{"error": "pic failed"})
            			return
        			}
					picURL = p.res
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"url": picURL,
			"fact": fact,
		})
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "main.html", nil)})

router.Run(":8080")
}