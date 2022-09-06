package Models

import (
    "gorm.io/gorm"
    "log"
)




type User struct{
    gorm.Model
    Name string `json:"name"`
    Phone int64 `json:"phone"`
    Email string `json:"email"`
    Address string `json:"address"`
    Gender string `json:"gender"`

}

func (u *User)CreateUser(){
    db.Create(u)
    
}

func  GetAllUsers()[]User{
    log.Println("users fetched from model")
    var users []User
    db.Find(&users)
    return users
}
func FindUserById(ID int64)User{
    var user User
    db.Where("id = ?", ID ).Find(&user)
    return user
}
// func NumberOfUsers()int64{

// }