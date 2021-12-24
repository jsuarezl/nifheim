package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

func rError(w http.ResponseWriter, code int, message string) {
	rJSON(w, code, map[string]string{"error": message})
}

func rJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func decodeJSON(req *http.Request, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid request")
	}
	decoder := json.NewDecoder(req.Body)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
