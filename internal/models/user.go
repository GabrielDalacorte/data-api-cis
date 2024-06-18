package models

type User struct {
    ID        int    `json:"id"`
    Email     string `json:"email"`
    Username  string `json:"username"`
    Password  string `json:"password"`
    FirstName string `json:"first_name"`
    LastName  string `json:"last_name"`
}
