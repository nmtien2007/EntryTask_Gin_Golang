package main

import (
	"fmt"
	"os"
	"source_apis/db_connection"
	"source_apis/login"
	"source_apis/middleware"
	"source_apis/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"message": "pong",
	//	})
	//})
	//routers.HelloWorld(r)
	db, err := db_connection.ConnectDB()
	if err != nil {
		fmt.Println("Connect DB Fail")
		fmt.Println(err)
	} else {
		fmt.Println("Connect DB Successfully")
	}
	fmt.Println(db)

	//Run middlewares
	r.Use(cors.New(middleware.GetCorsConfigs()))

	login.RouterLogin(r, db)

	r.Use(middleware.ValidateToken(db))

	// do not set table name is plural
	db.SingularTable(true)

	user.RouterUser(r, db)
	r.Run(os.Getenv("PORT"))
}
