package Models
import (
    "gorm.io/gorm"
)

type Wishlist struct{
    gorm.Model
    UserID string `gorm:"index"`
    Total float32 
    Products []Product `gorm:"many2many:wishlists_products"`
}


type Wishlists_product struct{
    Wishlist_id uint `json:"wishlist_id"`
    Product_id uint `json:"product_id"`
    Quantity int `json:"quantity"`
}


func (u *Wishlist)CreateWish(){
    db.Create(u)
    
}


func (wish * Wishlists_product)AddToWishlist(){
    var wishlistItem Wishlists_product
    db.Where("cart_id = ? and product_id =?",wish.Wishlist_id,wish.Product_id).Find(&wishlistItem)
    if wishlistItem.Wishlist_id == 0{
        db.Create(wish)
    }else{
        db.Model(Wishlists_product{}).Where("wishlist_id = ? and product_id =?",wish.Wishlist_id,wish.Product_id).Update("quantity = ?",wish.Quantity + wishlistItem.Quantity)
    } 
}

