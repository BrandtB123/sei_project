package router

import (
	"net/http"
	"sei_project/middleware"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/totalTxs", middleware.GetTransactionsInPastNBlocks)
	mux.HandleFunc("/api/blocksByProposer", middleware.GetProposerBlocks)
	mux.HandleFunc("/api/topNPeers", middleware.GetNPeersOverNBlocks)
	return mux
}
