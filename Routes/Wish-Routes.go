package Routes

import(
    "github.com/gorilla/mux"
    "Shubham/Controllers"
)

var WishRoutes= func (router *mux.Router)  {
    router.HandleFunc("/wishlist/create", Controllers.CreateWishlist).Methods("POST")
   router.HandleFunc("/wishlist/add_items/{userID}", Controllers.AddItemsInWishlist).Methods("POST")
//   router.HandleFunc("/cart/{userID}", Controllers.GetItemsinCart).Methods("GET")
//   router.HandleFunc("/cart/CountItems", Controllers.NumberOfItems).Methods("GET")
  
}