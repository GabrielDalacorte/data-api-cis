package controllers

import (
    "data-api-cis/internal/models"
    "gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
    return db.Create(user).Error
}

func GetAllUsers(db *gorm.DB) ([]models.User, error) {
    var users []models.User
    if err := db.Find(&users).Error; err != nil {
        return nil, err
    }
    return users, nil
}

func DeleteUser(db *gorm.DB, id uint) error {
    return db.Delete(&models.User{}, id).Error
}
