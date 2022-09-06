package Controllers
import(
    //"fmt"
    "encoding/json"
    "net/http"
   
    //"strconv"
    "Shubham/Models"
    "Shubham/Utils"
    
    

)

func CreateWishlist(w http.ResponseWriter, r *http.Request){
    wishlist := &Models.Wishlist{}
    wishlist.UserID = r.FormValue("userID") 
    wishlist.CreateWish()
    
}
func AddItemsInWishlist(w http.ResponseWriter , r *http.Request){
    if Utils.IsvalidUser(r){
        item:= &Models.Wishlists_product{}
        Utils.ParseData(r, item)
        item.AddToWishlist()
        res, _:=json.Marshal(item)
        w.Header().Set("Content-Type" , "application/json")
        w.Write(res)
        
    }
}