package Controllers

import (
	"Shubham/Models"
	"Shubham/Utils"
	"encoding/json"
	
	"github.com/gorilla/mux"
	"fmt"
	"net/http"
	"strconv"
)

func AddProduct(w http.ResponseWriter, r *http.Request) {
    product := &Models.Product{}


    Utils.ParseData(r,product)
	product.CreateProduct()
	res, err := json.Marshal(product)
	if err != nil {
		http.Error(w, "Hello There is No Data in Request", http.StatusBadRequest)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")

	w.Write(res)
}


func GetProduct(w http.ResponseWriter , r *http.Request){
    product := Models.GetAllProduct()
    //fmt.Println(product)
	jsonPROD, _ := json.Marshal(product)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonPROD)
}
func GetProductById(w http.ResponseWriter , r *http.Request){
	params := mux.Vars(r)
	productID ,_:= strconv.Atoi(params["Id"])
	product:=Models.FindProductById(productID)
	res, _:=json.Marshal(product)
	w.Header().Set("Content-Type" , "application/json")
	w.Write(res)
	
}

func CountProduct(w http.ResponseWriter , r *http.Request){
        prodCount:=Models.GetAllProduct()
        count := 0
        for i:= range(prodCount){
            count ++
            _=i
        }
        w.Header().Set("content-type" , "text")
        fmt.Fprint(w, count)
        
}