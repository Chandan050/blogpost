package controller

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/Chandan050/blogwebsite/database"
	"github.com/Chandan050/blogwebsite/models"
	"github.com/Chandan050/blogwebsite/util"
	"github.com/gin-gonic/gin"
)

func CreatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		var blogpost models.Blog
		if err := c.ShouldBindJSON(&blogpost); err != nil {
			fmt.Println("unable to parse body")
		}
		if err := database.DB.Create(&blogpost).Error; err != nil {

			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid payload"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Congratotion!, your post is live"})

	}
}

func AllPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		page, _ := strconv.Atoi(c.Query("page"))
		if page == 0 {
			page = 1
		}
		limit := 5
		var total int64
		offset := limit * (page - 1)
		var posts []models.Blog
		database.DB.Preload("User").Offset(offset).Limit(limit).Find(&posts)
		database.DB.Model(&models.Blog{}).Count(&total)
		var Lpage float64 = float64(int(total) / limit)

		c.JSON(http.StatusOK, gin.H{"data": posts,
			"meta": gin.H{"total": total,
				"page":      page,
				"last page": math.Ceil(Lpage)}})
	}
}
func DetailsPost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		var post models.Blog
		database.DB.Where("id=?", id).Preload("User").First(&post)
		c.JSON(http.StatusOK, gin.H{"data": post})
	}
}

func UpdatePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		blog := models.Blog{
			Id: uint(id),
		}
		if err := c.ShouldBindJSON(&blog); err != nil {
			fmt.Println("unable to parse body")
		}
		database.DB.Model(&blog).Updates(blog)
		c.JSON(http.StatusOK, gin.H{"message": "post updated sussfully"})

	}
}

func UniuePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		id, _ := util.Parsejwt(cookie)
		var blog []models.Blog

		database.DB.Model(&blog).Where("user_id=?", id).Preload("User").Find(&blog)

		c.JSON(http.StatusOK, gin.H{"blog": blog})

	}
}

func DeletePost() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		blog := models.Blog{
			Id: uint(id),
		}
		res := database.DB.Delete(&blog)
		if res.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "oops!,record not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "post deleted succesfully"})
	}
}
