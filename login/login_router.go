package login

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RouterLogin(e *gin.Engine, db *gorm.DB) {
	handlerDB := New(db)
	e.POST("/login", handlerDB.LoginUser)
}
