package controller

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sajid414/Go-react-auth/database"
	"github.com/sajid414/Go-react-auth/helper"
	"github.com/sajid414/Go-react-auth/model"
	"golang.org/x/crypto/bcrypt"
)

type formData struct {
	FullName        string `json:"fullName"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	Username        string `json:"userName" `
	Phone           string `json:"phone" `
	ConfirmPassword string `json:"confirmPassword"`
	Gender          string `json:"gender" `
}

func Login(c *gin.Context) {
	returnObj := gin.H{
		"status": "",
		"msg":    "Something went wrong",
	}
	//c.JSON(200, returnObj)

	var formData formData
	if err := c.ShouldBind(&formData); err != nil {
		log.Println("Form Data binding Error")
		c.JSON(400, returnObj)
		return
	}
	var user model.User
	database.DBConn.First(&user, "email=?", formData.Email)
	if user.UserID == "" {
		returnObj["msg"] = "User Not Found."
		returnObj["status"] = "Create Account First"
		c.JSON(400, returnObj)
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formData.Password))
	if err != nil {
		log.Println("Invalid Password")
		returnObj["msg"] = "Password doesn't match"
		c.JSON(401, returnObj)
		return
	}

	token, err := helper.GenerateToken(user)
	if err != nil {
		returnObj["msg"] = "token Creation Error"
		c.JSON(401, returnObj)
		return
	}

	returnObj["token"] = token
	returnObj["status"] = "OK"
	returnObj["msg"] = "User authenticated."
	returnObj["user"] = gin.H{
		"userId":   user.UserID,
		"email":    user.Email,
		"fullName": user.FullName,
	}
	c.JSON(201, returnObj)
}

func Register(c *gin.Context) {
	returnObj := gin.H{
		"status": "OK",
		"msg":    "Register route",
	}
	var formData formData
	if err := c.ShouldBind(&formData); err != nil {
		log.Println("Error in JSON Binding")
		returnObj["msg"] = "Error occurs"
		c.JSON(400, returnObj)
		return
	}

	var user model.User
	user.UserID = uuid.New().String()
	user.FullName = formData.FullName
	user.Email = formData.Email
	user.Password = helper.HashPassword(formData.Password)
	user.Username = formData.Username
	user.Phone = formData.Phone
	user.Gender = formData.Gender
	result := database.DBConn.Create(&user)

	if result.Error != nil {
		log.Println(result.Error)
		returnObj["msg"] = "User already exist"
		c.JSON(400, returnObj)
		return
	}

	returnObj["status"] = "OK"
	returnObj["msg"] = "User added successfully."
	returnObj["userId"] = user.UserID
	c.JSON(201, returnObj)

}

func Logout(c *gin.Context) {

}

func RefreshToken(c *gin.Context) {
	returnObj := gin.H{
		"status": "OK",
		"msg":    "Refresh token route",
	}
	c.JSON(200, returnObj)
}
