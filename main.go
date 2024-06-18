package main

import (
  "fmt"
  "log"
  "os"
  "strconv"

  "data-api/controllers"
  "data-api/models"
  "github.com/gin-gonic/gin"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

var DB *gorm.DB

func initDB() {
  dbHost := "cis-db"
  dbUser := "admin"
  dbPassword := "admin"
  dbName := "cis"

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Asia/Shanghai",
    dbHost, dbUser, dbPassword, dbName)

  var err error
  DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatalf("Falha ao conectar ao banco de dados: %v", err)
  }
}

func main() {
  initDB()

  DB.AutoMigrate(&models.User{})

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

  r.Run()
}
