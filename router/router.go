package router

import (
	"net/http"
	"sei_project/middleware"
)

func Router() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/calculateAllocations", middleware.Api())
	return mux
}
