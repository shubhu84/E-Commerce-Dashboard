package Controllers

import (
	"Shubham/Models"
	"Shubham/Utils"
	"encoding/json"
	"fmt"
	"strconv"
	"github.com/gorilla/mux"
	"net/http"
	"net/url"
)

func AddUser(w http.ResponseWriter, r *http.Request) {
	user := &Models.User{}

    Utils.ParseData(r,user)
	user.CreateUser()
	userID := strconv.Itoa(int(user.ID))

	//creation of Cart And WishList
	data:=url.Values{
		"userID":{userID},
	}
	resp1,_:=http.PostForm("http://localhost:9004/cart/create", data)
	resp2,_:=http.PostForm("http://localhost:9004/wishlist/create", data)
	_,_=resp1,resp2


	res, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Hello There is No Data in Request", http.StatusBadRequest)
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}


func GetUsers(w http.ResponseWriter , r *http.Request){
    users := Models.GetAllUsers()
	jsonUSR, _ := json.Marshal(users)
	w.Header().Set("content-type", "application/json")
	w.Write(jsonUSR)
}

func GetUserById(w http.ResponseWriter , r *http.Request){
	params := mux.Vars(r)
	userID ,_:= strconv.ParseInt(params["Id"],10,64)
	user:=Models.FindUserById(userID)
	res, _:=json.Marshal(user)
    w.Header().Set("Content-Type" , "application/json")
    w.Write(res)
	 
}
func CountUsers(w http.ResponseWriter , r *http.Request){
    userCount:=Models.GetAllUsers()
    count := 0
    for i := range(userCount){
        count ++
        _=i
    }
    w.Header().Set("content-type" , "text")
    fmt.Fprint(w, count)
        
}