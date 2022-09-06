package Config

import (
    "gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var(
    db * gorm.DB
)
const Dburl= "root:@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"
func Connect(){
    d,err:=gorm.Open(mysql.Open(Dburl) ,&gorm.Config{})
    if err !=nil{
        panic(err)
    }
    db = d
}
func GetDB() *gorm.DB{
    Connect() 
    return db 
}
func MakeMigration(x interface{}){
    db.AutoMigrate(&x)
}