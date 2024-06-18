package main

import (
    "fmt"
    "log"
    "os"
    "strconv"

    "data-api-cis/internal/controllers"
    "data-api-cis/internal/models"
    "github.com/gin-gonic/gin"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable",
        dbHost, dbUser, dbPassword, dbName)

    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
    }
}

func main() {
    initDB()

    r := gin.Default()

    r.POST("/users", func(c *gin.Context) {
        var user models.User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        if err := controllers.CreateUser(DB, &user); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, user)
    })

    r.GET("/users", func(c *gin.Context) {
        users, err := controllers.GetAllUsers(DB)
        if err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, users)
    })

    r.DELETE("/users/:id", func(c *gin.Context) {
        idStr := c.Param("id")
        id, err := strconv.ParseUint(idStr, 10, 32)
        if err != nil {
            c.JSON(400, gin.H{"error": "ID inválido"})
            return
        }
        if err := controllers.DeleteUser(DB, uint(id)); err != nil {
            c.JSON(500, gin.H{"error": err.Error()})
            return
        }
        c.JSON(200, gin.H{"message": "Usuário deletado com sucesso"})
    })

    r.Run(":8080")
}
