package login

import (
	"net/http"

	"source_apis/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type handlerDB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handlerDB {
	return handlerDB{db}
}

func (h handlerDB) LoginUser(e *gin.Context) {
	var userLoginSchame UserLoginSchema
	if err := e.ShouldBindJSON(&userLoginSchame); err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Validate Error",
				"message": err.Error()})

		return
	}

	token, err := user.CheckLogin(userLoginSchame.Username, userLoginSchame.Password, h.DB)

	if err != nil {
		e.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	e.JSON(http.StatusOK, gin.H{"token": token})
}
