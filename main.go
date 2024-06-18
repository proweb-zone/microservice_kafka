package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main(){
	router := gin.Default()
	router.POST("/send/:topic_name", response)
	router.Run("localhost:8082")
}

// getAlbums responds with the list of all albums as JSON.
func response(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, "success")
}
