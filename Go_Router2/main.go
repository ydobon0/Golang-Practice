package main

import (
	"golang-rest-api/router"
	//"encoding/json"
	//"fmt"
	"net/http"
	//"time"
	//"github.com/gorilla/mux"
	"os"

	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
)

/*
func handler(w http.ResponseWriter, _ *http.Request){
	fmt.Fprint(w,"Hello World Thanks  for connect me  ")
}
func handler1(w http.ResponseWriter, _ *http.Request){
	fmt.Fprint(w,"Handler 1  ")
}
func Now(w http.ResponseWriter, _*http.Request){
	now:=time.Now()
	p:=make(map[string]string)
	p["now"]=now.Format(time.ANSIC)
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(p)
}*/

func main() {
	//fmt.Println("")
	/*
	  r:=mux.NewRouter()
	  r.HandleFunc("/",handler).Methods("GET")// client wants to get
	  r.HandleFunc("/us",handler1).Methods("GET")
	  r.HandleFunc("/now",Now).Methods("GET")
	    http.Handle("/", r)
	  http.ListenAndServe(":8000",nil)
	*/

	router := router.ConfigureRouter()

	// srv := &http.Server{
	//     Handler:      router,
	//     Addr:         ":3000",
	//     // Good practice: enforce timeouts for servers you create!
	//     WriteTimeout: 15 * time.Second,
	//     ReadTimeout:  15 * time.Second,
	// }
	addr := ":3000"

	logrus.WithField("addr", addr).Info("Starting server...")

	if err := http.ListenAndServe(addr, handlers.LoggingHandler(os.Stdout, router)); err != nil {
		logrus.WithField("event", "start server").Fatal(err)
	}
	//log.Fatal(http.ListenAndServe(":3000", router))

}
