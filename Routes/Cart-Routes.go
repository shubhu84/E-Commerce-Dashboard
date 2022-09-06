package Routes

import(
    "github.com/gorilla/mux"
    "Shubham/Controllers"
)

var CartRoutes= func (router *mux.Router)  {
    router.HandleFunc("/cart/create", Controllers.CreateCart).Methods("POST")
    router.HandleFunc("/cart/add_items/{userId}", Controllers.AddItemsInCart).Methods("POST")
//   router.HandleFunc("/cart/CountItems", Controllers.NumberOfItems).Methods("GET")
   
}