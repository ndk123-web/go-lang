package routers

import (
	"github.com/gorilla/mux"
	"github.com/ndk123-web/muxrouter/controllers"
)

func Router(r *mux.Router) {
	r.HandleFunc("/", controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/create", controllers.InsertUser).Methods("POST")
	r.HandleFunc("/update", controllers.UpdateUser).Methods("PUT")
	r.HandleFunc("/delete", controllers.DeleteUser).Methods("DELETE")
	r.HandleFunc("/read", controllers.ReadUser).Methods("GET")
}
