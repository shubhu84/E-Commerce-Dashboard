package Utils

import (
	"encoding/json"
	//  "io/ioutil"
	"Shubham/Models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// func ParseBody(r *http.Request, x interface{}){
//     defer r.Body.Close()
//     dataByte , err := ioutil.ReadAll(r.Body)
//     if err != nil{
//         panic(err)
//     }
//     if err:=json.Unmarshal([]byte(dataByte),x); err !=nil{
//         return
//     }

// }
func ParseData(r *http.Request, x interface{}) {
	json.NewDecoder(r.Body).Decode(&x)
}
func IsvalidUser(r *http.Request) bool {
	user := &Models.User{}
	params := mux.Vars(r)
	URL := "http://localhost:9006/user/" + params["userId"]
	resp, _ := http.Get(URL)
	json.NewDecoder(resp.Body).Decode(&user)
	log.Println(user)
	return true
}
