package main

import (
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/jebrial/learnlink/models"
	//"io/ioutil"
	"log"
)

// type Config struct {
// 	dbUrl string
// }

func dbWare() gin.HandlerFunc {
	// // //load the config
	// file, err := ioutil.ReadFile("config.json")
	// if err != nil {
	// 	log.Panic(err)
	// }
	// var config Config
	// err = json.Unmarshal(file, &config)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// connect to the database

	db, err2 := models.NewDB("")
	if err2 != nil {
		log.Panic(err2)
	}
	return func(c *gin.Context) {

		c.Set("db", db)
		c.Next()
	}

}

func main() {
	ginServer := gin.Default()
	ginServer.Use(dbWare())
	//user routes
	ginServer.GET("/user/all", usersIndex)
	ginServer.GET("/user/find/:email", userSearch)
	ginServer.POST("/user/new", userAdd)
	ginServer.DELETE("/user/delete/:email", userRemove)
	//course routes
	ginServer.GET("/course/all/:email", courseIndex)
	ginServer.POST("/course/new", courseAdd)
	ginServer.POST("/course/update/:id", courseUpdate)
	ginServer.DELETE("/course/delete/:id", courseRemove)

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

// func courseIndex(ctx context.Context, w http.ResponseWriter, r *http.Request) {
// 	if r.Method != "GET" {
// 		http.Error(w, http.StatusText(405), 405)
// 		return
// 	}

// 	courses, err := models.ListCourse(ctx)
// 	if err != nil {
// 		http.Error(w, http.StatusText(500), 500)
// 		return
// 	}

// 	for _, user := range users {
// 		fmt.Fprintf(w, "%d, %s, %s\n", user.Id, user.Name, user.Email)
// 	}
// }
