package main

import (
	"log"
	"net/http"
	"strconv"

	"blog.com/blog/models"

	"github.com/gin-gonic/gin"
)

func main() {
	err := models.ConnectDatabase()
	checkErr(err)
	r := gin.Default()
	//API auhtor
	v1 := r.Group("/api/v1")
	{
		v1.GET("author", getAuthor)

		v1.GET("posts", getPost)
		v1.GET("posts/:id", getPostById)
		v1.POST("posts", addPost)
		v1.PUT("posts/:id", updatePost)
		v1.DELETE("posts/:id", deletePost)

		v1.GET("comment", getComment)
		v1.GET("comment/:id", getCommentById)
		v1.POST("comment", addComment)
		v1.PUT("comment/:id", updateComment)
		v1.DELETE("comment/:id", deleteComment)

		v1.GET("tag", getTag)
		v1.GET("tag/:id", getTagById)
		v1.POST("tag", addTag)
		v1.PUT("tag/:id", updateTag)
		v1.DELETE("tag/:id", deleteTag)
	}
	r.Run()
}

// Author handler functions----------------------------------------
func getAuthor(c *gin.Context) {
	author, err := models.GetAuthor(5)
	checkErr(err)
	if author == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Record Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": author})
	}
}

// Post handler functions
func getPost(c *gin.Context) {
	posts, err := models.GetPosts(5)
	checkErr(err)
	if posts == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Record Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": posts})
	}
}
func getPostById(c *gin.Context) {
	id := c.Param("id")
	posts, err := models.GetPersonById(id)
	checkErr(err)
	//if the post is blank we can assume nothing is found
	if posts.Title == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Record Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": posts})
	}

}
func addPost(c *gin.Context) {
	var json models.Posts
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success, err := models.AddPosts(json)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func updatePost(c *gin.Context) {
	var json models.Posts
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	postID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	success, err := models.UpdatePost(json, postID)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func deletePost(c *gin.Context) {
	postId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	success, err := models.DeletePost(postId)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

// Comment handler functions-------------------------------------
func getComment(c *gin.Context) {
	comment, err := models.GetComment(5)
	checkErr(err)
	if comment == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Record Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": comment})
	}
}
func getCommentById(c *gin.Context) {
	id := c.Param("id")
	comment, err := models.GetCommentById(id)
	checkErr(err)
	//if the comment is blank we can assume nothing is found
	if comment.Comment_data == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Record Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": comment})
	}

}
func addComment(c *gin.Context) {
	var json models.Comment
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success, err := models.AddComment(json)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func updateComment(c *gin.Context) {
	var json models.Comment
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	success, err := models.UpdateComment(json, commentID)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func deleteComment(c *gin.Context) {
	commentId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	success, err := models.DeleteComment(commentId)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

// Tag handler functions----------------------------------------
func getTag(c *gin.Context) {
	tags, err := models.GetTags(5)
	checkErr(err)
	if tags == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Record Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": tags})
	}
}

func getTagById(c *gin.Context) {
	id := c.Param("id")
	tags, err := models.GetTagById(id)
	checkErr(err)
	//if the tag is blank we can assume nothing is found
	if tags.Tag_data == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No Record Found"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"data": tags})
	}
}

func addTag(c *gin.Context) {
	var json models.Tag
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	success, err := models.AddTags(json)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func updateTag(c *gin.Context) {
	var json models.Tag
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tagID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	success, err := models.UpdateTag(json, tagID)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}
func deleteTag(c *gin.Context) {
	tagId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	success, err := models.DeleteTag(tagId)
	if success {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
