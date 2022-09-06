package Models
import(
    "Shubham/Config"
    "gorm.io/gorm"
)
var db *gorm.DB
func init(){
    db =Config.GetDB()
    Config.MakeMigration(User{})
    Config.MakeMigration(Product{})
    Config.MakeMigration(Cart{})
    Config.MakeMigration(Carts_product{})
    Config.MakeMigration(Wishlist{})
    Config.MakeMigration(Wishlists_product{})
}