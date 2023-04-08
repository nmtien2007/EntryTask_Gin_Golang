package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RouterUser(e *gin.Engine, db *gorm.DB) {
	handlerDB := New(db)
	user := e.Group("/user")
	{
		user.GET("/get_user_info", handlerDB.GetUserInfo)
		user.POST("/create_user", handlerDB.CreateUser)
		user.PUT("/update_user/:username", handlerDB.UpdateUser)
		user.DELETE("/delete_user/:id", handlerDB.DeleteUser)

		user.GET("/get_company_infos", handlerDB.GetCompanyInfos)
	}
}
