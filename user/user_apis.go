package user

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type handlerDB struct {
	DB *gorm.DB
}

func New(db *gorm.DB) handlerDB {
	return handlerDB{db}
}

func ConvertToUser(T any) User {
	s, _ := T.(User)
	return s
}

func GetCurrentUser(e *gin.Context) User {
	userInfo, _ := e.Get("UserInfo")
	data := ConvertToUser(userInfo)
	return data
}

func (h handlerDB) GetUserInfo(e *gin.Context) {
	var user []User

	// fmt.Println(GetCurrentUser(e))

	currentUser := GetCurrentUser(e)
	fmt.Println(currentUser)

	if err := h.DB.Preload("Company").Preload("Groups.Perms").Find(&user).Error; err != nil {
		fmt.Println(err)
		e.JSON(http.StatusOK, gin.H{
			"user_infos": nil,
		})
	} else {
		fmt.Println(user)
		e.JSON(http.StatusOK, gin.H{
			"user_infos": user,
		})
	}

}

func (h handlerDB) CreateUser(e *gin.Context) {
	var userSchema UserSchema
	if err := e.ShouldBindJSON(&userSchema); err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Validate Error",
				"message": err.Error()})

		return
	}
	user := User{
		Username:  userSchema.Username,
		Password:  userSchema.Password,
		FirstName: userSchema.FirstName,
		LastName:  userSchema.LastName,
	}
	result := h.DB.Create(&user)
	e.JSON(http.StatusOK, gin.H{
		"success": result.RowsAffected,
	})
}

func (h handlerDB) UpdateUser(e *gin.Context) {
	username := e.Param("username")
	var user User

	if err := h.DB.Where("username = ?", username).First(&user).Error; err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "user_not_existed",
				"message": "User không tồn tại"})

		return
	}

	var updateSchema UserUpdateSchema
	if err := e.ShouldBindJSON(&updateSchema); err != nil {
		e.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Validate Error",
				"message": err.Error()})

		return
	}

	h.DB.Model(&user).Update(updateSchema)

	// user.FirstName = "Tiến Minh"
	// h.DB.Save(&user)

	e.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h handlerDB) DeleteUser(e *gin.Context) {
	id := e.Param("id")
	h.DB.Delete(&User{}, id)

	e.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (h handlerDB) GetCompanyInfos(e *gin.Context) {
	var companies []Company

	if err := h.DB.Preload("Users").Find(&companies).Error; err != nil {
		fmt.Println(err)
		e.JSON(http.StatusOK, gin.H{
			"company_infos": nil,
		})
	} else {

		e.JSON(http.StatusOK, gin.H{
			"company_infos": companies,
		})
	}
}
