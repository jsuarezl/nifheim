package main

import (
	"github.com/Beelzebu/nifheim/handler"
	"github.com/Beelzebu/nifheim/storage"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	router := mux.NewRouter().StrictSlash(true)
	router.Use(CORSMiddleware)
	punishments := router.PathPrefix("/punishments").Methods("GET").Subrouter() // punishments are read only
	punishments.HandleFunc("/punishments", handler.GetPunishments)
	punishments.HandleFunc("/punishments/:id", handler.GetPunishmentById)
	err = http.ListenAndServe(":5000", router)
	if err != nil {
		panic(err)
	}
	defer storage.Close()
}

func CORSMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
