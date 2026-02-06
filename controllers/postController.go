package controllers

import (
	"goorm/initializers"
	"goorm/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	c.Bind(&body)
	// get the data from req body
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}
	// Return JSON response
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func Postindex(c *gin.Context) {
	// get the posts
	var posts []models.Post

	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	initializers.DB.First(&post, id)
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body  string
		Title string
	}
	c.Bind(&body)

	var post models.Post
	initializers.DB.First(&post, id)

	initializers.DB.Model(&post).Updates(models.Post{Body: body.Body, Title: body.Title})
	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {
	id := c.Param("id")
	var body struct {
		Body  string
		Title string
	}
	var post models.Post
	c.Bind(&body)
	initializers.DB.First(&post, id)
	initializers.DB.Delete(models.Post{}, id)

	c.JSON(http.StatusOK, gin.H{
		"post": post,
		"msg":  "Data Deleted Successgully",
	})
}
