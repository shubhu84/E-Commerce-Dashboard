package Controllers

import (
   // "log"
	"encoding/json"

	"net/http"

	//"strconv"
	"Shubham/Models"
	"Shubham/Utils"
)

func CreateCart(w http.ResponseWriter, r *http.Request){
    cart := &Models.Cart{}
    cart.UserID = r.FormValue("userID") 
    cart.CreateNewCart()
    

}

func AddItemsInCart(w http.ResponseWriter, r *http.Request){
    if Utils.IsvalidUser(r){
        item:= &Models.Carts_product{}
        Utils.ParseData(r, item)
        item.AddToCart()
        res, _:=json.Marshal(item)
        w.Header().Set("Content-Type" , "application/json")
        w.Write(res)
        
    }
 }

func GetItemsinCart(w http.ResponseWriter, r *http.Request){

}
 func NumberOfItems(w http.ResponseWriter, r *http.Request){

 }
