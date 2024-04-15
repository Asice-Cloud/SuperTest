package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"Chat/model"
	"Chat/service"

	"github.com/gin-gonic/gin"
)

func Login(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Login successfully",
	})
}

// Index
// User Home Page
// @Tags User Home
// @Success	200	{string} welcome,user
// @router /user/index [get]
func Index(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Welcome!",
	})
}

// GetUserList
// @Summary Find all user
// @Tags UserModule
// @Success	 200	{string} json{"code","message"}
// @router /user/userlist [get]
func GetUserList(context *gin.Context) {
	var data []model.UserBasic
	data = service.GetUserList()
	if len(data) != 0 {
		context.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	}
}

// CreateUser
// @Summary	Add user
// @Tags UserModule
// @param name query string false "Name"
// @param password query string false "Password"
// @param repassword query string false "Twice Password"
// @Success	200	{string} json{"code","message"}
// @router /user/adduser [get]
func CreateUser(context *gin.Context) {
	user := model.UserBasic{}
	user.Name = context.Query("name")
	password := context.Query("password")
	repassword := context.Query("repassword")
	if password != repassword {
		context.JSON(-1, gin.H{
			"message": "twice password is not matched",
		})
		return
	}
	user.Password = password
	rep, err := service.CreatUser(user)
	if rep == -1 {
		context.JSON(-1, gin.H{
			"message": "User already exist",
		})
	}
	if err != nil {
		context.JSON(-1, gin.H{
			"message": "Failed to add user",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully create user: %d", rep),
	})
}

// DeleteUser
// @Summary	Delete user
// @Tags UserModule
// @param id query string false "id"
// @Success	200	{string} json{"code","message"}
// @router /user/deluser [delete]
func DeleteUser(context *gin.Context) {
	var user model.UserBasic
	id, err := strconv.Atoi(context.Query("id"))
	if err != nil {
		context.JSON(-1, gin.H{
			"message": "Please Input a valid number",
		})
		return
	}
	user.ID = uint(id)
	err = service.DeleteUser(user)
	if err != nil {
		context.JSON(-1, gin.H{
			"message": "Failed to delete user",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Successfully delete user",
	})
}

// UpdateUser
// @Summary	Update user
// @Tags UserModule
// @Accept x-www-form-urlencoded
// @Param id formData string false "id"
// @Param name formData string false "name"
// @Param password formData string false "password"
// @Success	200	{string} json{"code","message"}
// @router /user/updateuser [post]
func UpdateUser(context *gin.Context) {
	var user model.UserBasic
	id, _ := strconv.Atoi(context.PostForm("id"))
	user.ID = uint(id)
	user.Name = context.PostForm("name")
	user.Password = context.PostForm("password")
	rep, err := service.UpdateUser(user)
	if err != nil {
		context.JSON(-1, gin.H{
			"message": "Failed to update user information",
		})
		return
	}
	if rep == -1 {
		context.JSON(-1, gin.H{
			"message": "User not exists",
		})
	}
	context.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Successfully update user: %d", rep),
	})
}
