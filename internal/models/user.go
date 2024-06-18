package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    Email     string `json:"email" gorm:"not null"`
    Username  string `json:"username" gorm:"size:256;not null"`
    Password  string `json:"password" gorm:"size:256;not null"`
    FirstName string `json:"first_name" gorm:"size:256;not null"`
    LastName  string `json:"last_name" gorm:"size:256;not null"`
}
