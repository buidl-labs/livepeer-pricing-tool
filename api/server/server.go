package server

import (
	"encoding/json"
	"math/big"
	"net/http"
	"strconv"

	"api/dataservice"
	"api/model"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func reformatPPPValue(ppp string) float64 {
	priceRat, ok := new(big.Rat).SetString(ppp)
	if !ok {
		// TODO: return error
		// The returned error can be used to exclude orchestrators from the eventual list
		return 0
	}
	priceFloat, _ := priceRat.Float64()
	return priceFloat
}

func reformatOrchestrator(x model.DBOrchestrator) model.APIOrchestrator {
	orch := model.APIOrchestrator{}
	n1 := new(big.Int)
	n2 := new(big.Int)
	orch.Address = x.Address
	orch.ServiceURI = x.ServiceURI
	orch.LastRewardRound = x.LastRewardRound
	orch.RewardCut = x.RewardCut
	orch.FeeShare = x.FeeShare
	n1, ok := n1.SetString(x.DelegatedStake, 10)
	if !ok {
		log.Errorln("SetString: error")
	}
	orch.DelegatedStake = n1
	orch.ActivationRound = x.ActivationRound
	n2, ok = n2.SetString(x.DeactivationRound, 10)
	if !ok {
		log.Errorln("SetString: error")
	}
	orch.DeactivationRound = n2
	orch.Active = x.Active
	orch.Status = x.Status
	orch.PricePerPixel = reformatPPPValue(x.PricePerPixel)
	orch.UpdatedAt = x.UpdatedAt
	return orch
}

func reformatPriceHistory(x model.DBPriceHistory) model.APIPriceHistory {
	ph := model.APIPriceHistory{}
	ph.Time = x.Time
	ph.PricePerPixel = reformatPPPValue(x.PricePerPixel)
	return ph
}

// API endpoint handler for /orchestratorStats
func GetOrchestratorStats(w http.ResponseWriter, req *http.Request) {
	log.Infoln("GET /orchestratorStats")
	query := req.URL.Query()
	excludeUnavailable, err := strconv.ParseBool(query.Get("excludeUnavailable"))
	if err != nil {
		excludeUnavailable = true
	}
	dborchs := dataservice.FetchOrchestratorStatistics(excludeUnavailable)
	data := []model.APIOrchestrator{}
	for _, x := range dborchs {
		data = append(data, reformatOrchestrator(x))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// API endpoint handler for /priceHistory/{address}
func GetOrchestratorPriceHistory(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	log.Infof("GET /priceHistory/%s", params["address"])
	dbphs := dataservice.FetchPricingHistory(params["address"])
	data := []model.APIPriceHistory{}
	for _, x := range dbphs {
		data = append(data, reformatPriceHistory(x))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

// Starts the server on port number passed as serverPort
func StartServer(serverPort string) {
	router := mux.NewRouter()
	router.HandleFunc("/orchestratorStats", GetOrchestratorStats).Methods("GET")
	router.HandleFunc("/priceHistory/{address}", GetOrchestratorPriceHistory).Methods("GET")
	log.Infoln("Starting server at PORT", serverPort)
	log.Fatalln("Error in starting server", http.ListenAndServe(serverPort, handlers.CORS()(router)))
}
