package middleware

import (
	"net/http"
	"source_apis/core"

	"source_apis/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func ValidateToken(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {

		userId, err := core.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,
				gin.H{
					"error":   "error_unauthorized",
					"message": "Chứng thực user không thành công. Vui lòng đăng nhập lại"})

			c.Abort()
			return
		}

		var userModel user.User
		if err := db.Where("id = ?", userId).Preload("Groups.Perms").First(&userModel).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest,
				gin.H{
					"error":   "user_not_existed",
					"message": "User không tồn tại"})

			c.Abort()
			return
		}
		c.Set("UserInfo", userModel)
		c.Next()
	}
}
