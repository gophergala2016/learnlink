package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jebrial/learnlink/models"
	"log"
	"os"
)

type Conf struct {
	Url string
}

func dbWare(url string) gin.HandlerFunc {
	// connect to the database
	db, err := models.NewDB(url)
	if err != nil {
		log.Panic(err)
	}
	return func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	}
}

func main() {
	//load the config
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Panic(err)
	}
	decoder := json.NewDecoder(file)
	conf := Conf{}

	err = decoder.Decode(&conf)
	if err != nil {
		log.Panic(err)
	}

	// Set up server
	ginServer := gin.Default()
	ginServer.Use(dbWare(conf.Url))
	//user routes
	//ginServer.POST("/login", )
	ginServer.POST("/login/new", userAdd)

	ginServer.GET("/api/user/all", usersIndex)
	ginServer.GET("/api/user/find/:email", userSearch)
	ginServer.DELETE("/api/user/delete/:email", userRemove)
	//course routes
	ginServer.GET("/api/course/all/:email", courseIndex)
	ginServer.POST("/api/course/new", courseAdd)
	ginServer.PUT("/api/course/update/:id", courseUpdate)
	ginServer.DELETE("/api/course/delete/:id", courseRemove)

	ginServer.Run(":3001")
}

func usersIndex(ctx *gin.Context) {
	users, err := models.AllUsers(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "No users found"})
		return
	}
	ctx.JSON(200, users)
}

func userSearch(ctx *gin.Context) {
	user, err := models.FindUser(ctx)
	if err != nil {
		ctx.JSON(404, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(200, user)
}

func userAdd(ctx *gin.Context) {
	_, err := models.AddUser(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "New user added"})
}

func userRemove(ctx *gin.Context) {
	_, err := models.RemoveUser(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "User removed"})
}

func courseIndex(ctx *gin.Context) {
	courses, err := models.AllCourses(ctx)
	if err != nil {
		log.Fatal(err)
		ctx.JSON(404, gin.H{"error": "No courses found"})
		return
	}
	ctx.JSON(200, courses)
}

func courseAdd(ctx *gin.Context) {
	_, err := models.AddCourse(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "New course added"})
}

func courseUpdate(ctx *gin.Context) {
	_, err := models.UpdateCourse(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Error processing request"})
		return
	}
	ctx.JSON(200, gin.H{"success": "Course updated"})
}

func courseRemove(ctx *gin.Context) {
	_, err := models.RemoveCourse(ctx)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Internal error"})
		return
	}
	ctx.JSON(200, gin.H{"success": "Course removed"})
}
