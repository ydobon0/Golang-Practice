package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func handler(ww http.ResponseWriter, _ *http.Request) { // r is what we read from client and w is what we write back
	fmt.Fprint(ww, "Hello. Thanks for connecting!") //sends formatted output to client which is w
}

func handler1(ww http.ResponseWriter, _ *http.Request) { // r is what we read from client and w is what we write back
	fmt.Fprint(ww, "Handler 1") //sends formatted output to client which is w
}

func now(ww http.ResponseWriter, _ *http.Request) { // r is what we read from client and w is what we write back
	//fmt.Fprint(ww, time.Now().Format("15:04:05\n"))
	now := time.Now().Local()
	p := make(map[string]string)
	p["UTC"] = now.Local().Format(time.ANSIC)
	ww.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ww).Encode(p)
}

func utc(ww http.ResponseWriter, _ *http.Request) { // r is what we read from client and w is what we write back
	//fmt.Fprint(ww, time.Now().Format("15:04:05\n"))
	now := time.Now().Local()
	p := make(map[string]string)
	p["UTC"] = now.Local().UTC().Format(time.ANSIC)
	ww.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ww).Encode(p)
}

func ny(ww http.ResponseWriter, _ *http.Request) { // r is what we read from client and w is what we write back
	//fmt.Fprint(ww, time.Now().Format("15:04:05\n"))

	loc, _ := time.LoadLocation("America/New_York")
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["NY"] = now.Format(time.ANSIC)
	ww.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ww).Encode(p)
}

func tokyo(ww http.ResponseWriter, _ *http.Request) { // r is what we read from client and w is what we write back
	//fmt.Fprint(ww, time.Now().Format("15:04:05\n"))

	loc, _ := time.LoadLocation("Asia/Tokyo")
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["NY"] = now.Format(time.ANSIC)
	ww.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ww).Encode(p)
}

func la(ww http.ResponseWriter, _ *http.Request) { // r is what we read from client and w is what we write back
	//fmt.Fprint(ww, time.Now().Format("15:04:05\n"))

	loc, _ := time.LoadLocation("America/Los_Angeles")
	now := time.Now().In(loc)
	p := make(map[string]string)
	p["NY"] = now.Format(time.ANSIC)
	ww.Header().Set("Content-Type", "application/json")
	json.NewEncoder(ww).Encode(p)
}

func main() {
	rr := mux.NewRouter()
	rr.HandleFunc("/", handler).Methods("GET") // Client wants to get from server so we use "GET". to post to server, we use "POST"
	rr.HandleFunc("/us", handler1).Methods("GET")
	rr.HandleFunc("/now", now).Methods("GET")
	rr.HandleFunc("/UTC", utc).Methods("GET")
	rr.HandleFunc("/NY", ny).Methods("GET")
	rr.HandleFunc("/tokyo", tokyo).Methods("GET")
	rr.HandleFunc("/la", la).Methods("GET")

	http.Handle("/", rr)
	http.ListenAndServe(":8000", nil)
}
