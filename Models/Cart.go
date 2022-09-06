package Models
import (
    "gorm.io/gorm"   
)

type Cart struct{
    gorm.Model
    UserID string `gorm:"index"`
    Total float32 
    Products []Product `gorm:"many2many:carts_products"`
}

type Carts_product struct{
    Cart_id uint `json:"cart_id"`
    Product_id uint `json:"product_id"`
    Quantity int `json:"quantity"`
}
func (u *Cart)CreateNewCart(){
    db.Create(u)
    
}


func  (u *Carts_product)AddToCart(){
    var cartItem Carts_product
    db.Where("cart_id = ? and product_id =?",u.Cart_id,u.Product_id).Find(&cartItem)
    if cartItem.Cart_id == 0{
        db.Create(u)
    }else{
        db.Model(Carts_product{}).Where("cart_id = ? and product_id =?",u.Cart_id,u.Product_id).Update("quantity = ?",u.Quantity+cartItem.Quantity)
    } 
 }