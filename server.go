package main 

import (
	"github.com/git-gonic/gin"
)

func main(){
	r := gin.Default();

	r.GET("/", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message" : "Hello server",
		})
	})

	r.Run(":3000")
}
