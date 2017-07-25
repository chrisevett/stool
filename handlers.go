package main

import (
	"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to the gAPI")
}

func GameIndex(w http.ResponseWriter, r *http.Request) {
	games := Games{
		Game{Name: "Splendor", Players: 4, Id: 1},
		Game{Name: "Love Letter", Players: 4, Id: 2},
	}

	// json.NewEncoder(w).Encode(games)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(games); err != nil {
		panic(err)
	}
}

func ShowGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	fmt.Fprintln(w, "Game:", gameId)
}