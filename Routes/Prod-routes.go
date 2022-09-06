package Routes

import(
    "github.com/gorilla/mux"
    "Shubham/Controllers"
)

var ProductRoutes= func (router *mux.Router)  {
    router.HandleFunc("/product/add_product", Controllers.AddProduct).Methods("POST")
  router.HandleFunc("/product/getproducts", Controllers.GetProduct).Methods("GET")
  router.HandleFunc("/product/count", Controllers.CountProduct).Methods("GET")
  router.HandleFunc("/product/{Id}", Controllers.GetProductById).Methods("GET")
}