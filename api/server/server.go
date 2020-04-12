package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"api/dataservice"

	"github.com/gorilla/mux"
	"github.com/gorilla/handlers"
	log "github.com/sirupsen/logrus"
)

// API endpoint handler for /orchestratorStats
func GetOrchestratorStats(w http.ResponseWriter, req *http.Request) {
	log.Infoln("GET /orchestratorStats")
	query := req.URL.Query()
	excludeUnavailable, err := strconv.ParseBool(query.Get("excludeUnavailable"))
	if err!=nil {
		excludeUnavailable = true
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataservice.FetchOrchestratorStatistics(excludeUnavailable))
}

// API endpoint handler for /priceHistory/{address}
func GetOrchestratorPriceHistory(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Infof("GET /priceHistory/%s", params["address"])
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dataservice.FetchPricingHistory(params["address"]))
}

// Starts the server on port number passed as serverPort
func StartServer(serverPort string) {
	router := mux.NewRouter()
	router.HandleFunc("/orchestratorStats", GetOrchestratorStats).Methods("GET")
	router.HandleFunc("/priceHistory/{address}", GetOrchestratorPriceHistory).Methods("GET")
	log.Infoln("Starting server at PORT", serverPort)
	log.Fatalln("Error in starting server", http.ListenAndServe(serverPort, handlers.CORS()(router)))
}