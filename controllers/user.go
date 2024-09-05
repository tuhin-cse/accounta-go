package controllers

import (
	"accounta-go/config"
	"accounta-go/middleware"
	"accounta-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Role = "user"
	var db = config.GetDB()
	var d = db.Create(&user)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Failed to register user",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Register user",
	})
}

func UserLogin(c *gin.Context) {
	var auth models.LoginUser
	var user models.User
	if err := c.ShouldBindJSON(&auth); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var db = config.GetDB()
	var d = db.Where("email = ?", auth.Email).First(&user)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User not found",
		})
		return
	}
	err := user.ComparePassword(auth.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Invalid password",
		})
		return
	}
	token, err := user.GenerateJWT()
	if err != nil {
		println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Failed to generate token",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Login user",
		"data": gin.H{
			"token": token,
		},
	})
}

func UserProfile(c *gin.Context) {
	cUser := c.MustGet("user").(middleware.AuthUser)
	var db = config.GetDB()
	var user models.User
	var d = db.Where("id = ?", cUser.ID).First(&user)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "User not found",
		})
		return
	}
	user.Mask()
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Successfully gets profile",
		"data":  user,
	})
}
