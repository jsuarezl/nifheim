package handler

import (
	"encoding/json"
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
	err := json.NewEncoder(w).Encode(punishment.GetPunishments())
	if err != nil {
		fmt.Println(err)
	}
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		fmt.Println(err)
		rError(w, http.StatusBadRequest, "invalid id")
		return
	}
	storedPunishment := punishment.GetPunishment(id)
	if (storedPunishment == punishment.Punishment{}) {
		rJSON(w, http.StatusNotFound, "punishment not found")
		return
	}
	rJSON(w, http.StatusOK, storedPunishment)
}

func AddPunishment(w http.ResponseWriter, r *http.Request) {
	var newPunishment punishment.Punishment

	// Call BindJSON to bind the received rJSON to
	// newPunishment.
	if err := decodeJSON(r, &newPunishment); err != nil {
		return
	}

	// Add the new album to the slice.
	//albums = append(albums, newPunishment)
	rJSON(w, http.StatusCreated, newPunishment)
}
