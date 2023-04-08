package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func hello_world(e *gin.Context) {

		e.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})

}

func HelloWorld(e *gin.Engine) {

	e.GET("/hello_world", hello_world)
}

