package routes

import (
	"Golang_latihan/test_restapi_crud/controllers"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	r.GET("/tasks", controllers.FindTasks)
	r.POST("/tasks", controllers.CreateTask)
	r.GET("/tasks/:id", controllers.FindTask)
	r.PATCH("/tasks/:id", controllers.UpdateTask)
	r.DELETE("tasks/:id", controllers.DeleteTask)
	return r
}
