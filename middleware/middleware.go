package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sei_project/models"
	"strconv"
)

func GetProposerBlocks(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	vars := r.URL.Query()
	proposer, ok := vars["proposer"]
	if !ok {
		fmt.Println("Must need validator")
		w.WriteHeader(400)
	}
	blockHeights, err := models.GetProposedBlocksByValidator(proposer[0])
	if err != nil {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(blockHeights)
}

func GetTransactionsInPastNBlocks(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	fmt.Println("SUCCESS")
	vars := r.URL.Query()
	nVars, ok := vars["n"]
	if !ok {
		fmt.Println("Must need n")
		w.WriteHeader(400)
	}
	n, err := strconv.Atoi(nVars[0])
	if err != nil {
		w.WriteHeader(400)
	}
	totalTxs, err := models.GetTransactionsInPastNBlocks(n)
	if err != nil {
		w.WriteHeader(400)
	}

	json.NewEncoder(w).Encode(totalTxs)
}

func GetNPeersOverNBlocks(w http.ResponseWriter,
	r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "content-type")
	vars := r.URL.Query()
	nVars, ok := vars["n"]
	if !ok {
		fmt.Println("Must need validator")
		w.WriteHeader(400)
	}

	n, err := strconv.Atoi(nVars[0])
	if err != nil {
		w.WriteHeader(400)
	}
	topNPeers, err := models.GetTopNPeersByScore(n)
	if err != nil {
		w.WriteHeader(400)
	}
	json.NewEncoder(w).Encode(topNPeers)
}
