package controllers

import (
	"Golang_latihan/test_restapi_crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type CreateTaskInput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender`
	Phone  string `json:"phone`
	Email  string `json:"email`
}

type UpdateTaskInput struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Phone  string `json:"phone`
	Email  string `json:"email`
}

// GET /tasks
// Get all tasks
func FindTasks(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var tasks []models.Kontak
	db.Find(&tasks)

	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// POST /tasks
// Create new task
func CreateTask(c *gin.Context) {
	// Validate input
	var input CreateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	// Create task
	task := models.Kontak{ID: input.ID, Name: input.Name, Gender: input.Gender, Phone: input.Phone, Email: input.Email}

	db := c.MustGet("db").(*gorm.DB)
	db.Create(&task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GET /tasks/:id
// Find a task
func FindTask(c *gin.Context) { // Get model if exist
	var task models.Kontak

	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// PATCH /tasks/:id
// Update a task
func UpdateTask(c *gin.Context) {

	db := c.MustGet("db").(*gorm.DB)
	// Get model if exist
	var task models.Kontak
	if err := db.Where("id = ?", c.Param("id")).First(&task).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	// Validate input
	var input UpdateTaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//date := "2006-01-02"
	//deadline, _ := time.Parse(date, input.Deadline)

	var updatedInput models.Kontak
	updatedInput.Name = input.Name
	updatedInput.Gender = input.Gender
	updatedInput.Phone = input.Phone
	updatedInput.Email = input.Email

	db.Model(&task).Updates(updatedInput)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// DELETE /tasks/:id
// Delete a task
func DeleteTask(c *gin.Context) {
	// Get model if exist
	db := c.MustGet("db").(*gorm.DB)
	var book models.Kontak
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Data tidak ditemukan!!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
