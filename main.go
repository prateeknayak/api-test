package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var ver = "1.0.0"
var sha = "a1b2c3def"

type appInfo struct {
	Version       string `json:"version"`
	LastCommitSHA string `json:"lastcommitsha"`
	Description   string `json:"description"`
}

type helloWorld struct {
	Msg string `json:"message"`
}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", root)
	r.HandleFunc("/version", version)

	s := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func version(w http.ResponseWriter, r *http.Request) {
	info := appInfo{
		Version:       ver,
		LastCommitSHA: sha,
		Description:   "pre-interview technical test",
	}

	infoJSON, err := json.Marshal(info)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(infoJSON)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	hw, err := json.Marshal(helloWorld{"Hello World!"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(hw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}