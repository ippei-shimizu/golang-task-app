package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Task struct {
    ID uint `json:"id" gorm:"primary_key"`
    Title string `json:"title"`
    Description string `json:"description"`
    Status string `json:"status"`
    CreatedAt time.Time `json:"created_at"`
}

func main() {
    var err error
    dsn := "host=db user=postgres password=example dbname=golang_task_app port=5432 sslmode=disable TimeZone=Asia/Tokyo"
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("failed to connect database", err)
    }

    db.AutoMigrate(&Task{})

    r := gin.Default()

    r.GET("/tasks", getTasks)
    r.POST("/tasks", createTask)
    r.DELETE("/tasks/:id", deleteTask)
    r.PUT("/tasks/:id", updateTask)

    r.Run(":8080")
}

func getTasks(c *gin.Context) {
    var tasks []Task
    db.Find(&tasks)
    c.JSON(http.StatusOK, gin.H{"data": tasks})
}

func createTask(c *gin.Context) {
    var task Task
    if err := c.ShouldBindJSON(&task); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&task)
    c.JSON(http.StatusCreated, task)
}

func deleteTask(c *gin.Context) {
    id := c.Param("id")
    var task Task
    if err := db.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
    }

    db.Delete(&task)
    c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

func updateTask(c *gin.Context) {
    id := c.Param("id")
    var task Task
    if err := db.First(&task, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "record not found"})
        return
    }

    var input Task
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    }

    db.Model(&task).Updates(input)
    c.JSON(http.StatusOK, task)
}
