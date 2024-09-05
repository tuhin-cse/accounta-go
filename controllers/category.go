package controllers

import (
	"accounta-go/config"
	"accounta-go/middleware"
	"accounta-go/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCategories(c *gin.Context) {
	cUser := c.MustGet("user").(middleware.AuthUser)
	var categories []models.Category
	db := config.GetDB()
	d := db.Where("user_id = ?", cUser.ID).Find(&categories)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Failed to get categories",
		})
		return
	}
	for i := range categories {
		categories[i].Mask()
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Successfully fetched categories",
		"data":  categories,
	})
}

func GetCategory(c *gin.Context) {
	cUser := c.MustGet("user").(middleware.AuthUser)
	id, found := c.Params.Get("id")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": false,
			"msg":   "Invalid Request",
		})
		return
	}
	var category models.Category
	db := config.GetDB()
	d := db.Where("id = ? AND user_id = ?", id, cUser.ID).First(&category)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Failed to get category",
		})
		return
	}
	category.Mask()
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Successfully fetched category",
		"data":  category,
	})
}

func PostCategory(c *gin.Context) {
	cUser := c.MustGet("user").(middleware.AuthUser)
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": false,
			"msg":   err.Error(),
		})
		return
	}
	category.UserID = cUser.ID
	db := config.GetDB()
	d := db.Create(&category)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Failed to create category",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Category created",
	})
}

func PatchCategory(c *gin.Context) {
	cUser := c.MustGet("user").(middleware.AuthUser)
	id, found := c.Params.Get("id")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": false,
			"msg":   "Invalid Request",
		})
		return
	}
	var category models.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return
	}
	update := make(map[string]interface{})
	if category.Name != "" {
		update["name"] = category.Name
	}
	if category.Type != "" {
		update["type"] = category.Type
	}
	if category.Description != "" {
		update["description"] = category.Description
	}
	db := config.GetDB()
	d := db.Model(&category).Where("id = ? AND user_id = ?", id, cUser.ID).Updates(update)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Failed to update category",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Category updated",
	})
}

func DeleteCategory(c *gin.Context) {
	cUser := c.MustGet("user").(middleware.AuthUser)
	id, found := c.Params.Get("id")
	if !found {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": false,
			"msg":   "Invalid Request",
		})
		return
	}
	var category models.Category
	db := config.GetDB()
	d := db.Where("id = ? AND user_id = ?", id, cUser.ID).Delete(&category)
	if d.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": true,
			"msg":   "Failed to delete category",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error": false,
		"msg":   "Category deleted",
	})
}
