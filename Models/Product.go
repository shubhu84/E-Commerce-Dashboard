package Models
import (
    "gorm.io/gorm"
)

type Product struct{
    gorm.Model
    Name string `json:"name"`
    Type string `json:"type"`
    Price float32 `json:"price"`
}

func (u *Product)CreateProduct(){
    db.Create(&u)
    
}

func  GetAllProduct()[]Product{
    var products []Product
    db.Find(&products)
    return products

}
func FindProductById(id int)Product{
    var product Product
    db.Where("id = ?" ,id).Find(&product)
    return product
}