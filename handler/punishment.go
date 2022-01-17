package handler

import (
	"fmt"
	"github.com/Beelzebu/nifheim/data/punishment"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetPunishments(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(punishment.GetPunishments())
	if err != nil {
		fmt.Println(err)
	}
}

func GetPunishmentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		rError(w, http.StatusBadRequest, "invalid id")
		return
	}
	storedPunishment := punishment.GetPunishment(id)
	if (storedPunishment == punishment.Punishment{}) {
		rError(w, http.StatusNotFound, "punishment not found")
		return
	}
	rJSON(w, http.StatusOK, storedPunishment)
}

func AddPunishment(w http.ResponseWriter, r *http.Request) {
	var newPunishment punishment.Punishment
	if err := decodeJSON(r, &newPunishment); err != nil {
		return
	}
	// store punishment in database
	rJSON(w, http.StatusCreated, newPunishment)
}
