package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sei_project/models"
)

func GetValidatorBlocks(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	vars := r.URL.Query()
	validator, ok := vars["validator"]
	if !ok {
		fmt.Println("Must need validator")
		w.WriteHeader(400)
	}

	var blocks []models.Block
	json.NewEncoder(w).Encode(proratedResponse)
}
