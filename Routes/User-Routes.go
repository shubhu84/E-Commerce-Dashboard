package Routes

import (
	"Shubham/Controllers"

	"github.com/gorilla/mux"
)

var UserRoutes = func(router *mux.Router) {
	router.HandleFunc("/user/add_user", Controllers.AddUser).Methods("POST")
	router.HandleFunc("/user/getUser", Controllers.GetUsers).Methods("GET")
  router.HandleFunc("/user/count", Controllers.CountUsers).Methods("GET")
	router.HandleFunc("/user/{Id}", Controllers.GetUserById).Methods("GET")
 

}
