package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

type Orchestrator struct {
	Address           string `json:"Address"`
	ServiceURI        string `json:"ServiceURI"`
	LastRewardRound   int    `json:"LastRewardRound"`
	RewardCut         int    `json:"RewardCut"`
	FeeShare          int    `json:"FeeShare"`
	DelegatedStake    int64  `json:"DelegatedStake"`
	ActivationRound   int    `json:"ActivationRound"`
	DeactivationRound int64  `json:"DeactivationRound"`
	Active            bool   `json:"Active"`
	Status            string `json:"Status"`
	PricePerPixel     string `json:"PricePerPixel"`
}

type OrchestratorStats struct {
	Address           string `json:"Address"`
	ServiceURI        string `json:"ServiceURI"`
	LastRewardRound   int    `json:"LastRewardRound"`
	RewardCut         int    `json:"RewardCut"`
	FeeShare          int    `json:"FeeShare"`
	DelegatedStake    int64  `json:"DelegatedStake"`
	ActivationRound   int    `json:"ActivationRound"`
	DeactivationRound int64  `json:"DeactivationRound"`
	Active            bool   `json:"Active"`
	Status            string `json:"Status"`
	PricePerPixel     string `json:"PricePerPixel"`
	UpdatedAt         int    `json:"UpdatedAt"`
}

type PriceHistory struct {
	Address       string `json:"Address"`
	PricePerPixel string `json:"PricePerPixel"`
	Time          int    `json:"Time"`
}

func initializeDB() {
	database, _ := sql.Open("sqlite3", "./pricing_tool.db")
	statement, _ := database.Prepare("CREATE TABLE IF NOT EXISTS Orchestrators (Address TEXT PRIMARY KEY, ServiceURI TEXT, LastRewardRound INTEGER, RewardCut INTEGER, FeeShare INTEGER, DelegatedState INTEGER, ActivationRound INTEGER, DeactivationRound INTEGER, Active INTEGER, Status TEXT, PricePerPixel INTEGER, UpdateAt INTEGER)")
	statement.Exec()
	statement, _ = database.Prepare("CREATE TABLE IF NOT EXISTS PriceHistory (Address TEXT, Time INTEGER, PricePerPixel INTEGER)")
	statement.Exec()
}

func fetchAndStore() {

	database, _ := sql.Open("sqlite3", "./pricing_tool.db")

	fmt.Println("Fetching the data...")
	response, err := http.Get("http://localhost:7935/registeredOrchestrators")
	if err != nil {
		fmt.Println("The HTTP request failed with error", err)
	} else {
		fmt.Println("The HTTP reqeust succeeded.")
		// data, _ := ioutil.ReadAll(response.Body)
		// fmt.Println(string(data))
	}

	orchestrators := []Orchestrator{}
	err = json.NewDecoder(response.Body).Decode(&orchestrators)
	if err != nil {
		fmt.Println("Error JSON parsing", err.Error())
	} else {
		fmt.Println("JSON parsing successful")
	}

	for i, x := range orchestrators {
		fmt.Println(i)

		statement, _ := database.Prepare("INSERT INTO Orchestrators (Address, ServiceURI, LastRewardRound, RewardCut, FeeShare, DelegatedState, ActivationRound, DeactivationRound, Active, Status, PricePerPixel, UpdateAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
		_, err = statement.Exec(x.Address, x.ServiceURI, x.LastRewardRound, x.RewardCut, x.FeeShare, x.DelegatedStake, x.ActivationRound, x.DeactivationRound, x.Active, x.Status, x.PricePerPixel, time.Now().Unix())
		if err != nil {
			fmt.Println("Orchestrator data already exists. Updating orchestrator data for", x.Address)
			statement, _ = database.Prepare("UPDATE Orchestrators SET ServiceURI=?, LastRewardRound=?, RewardCut=?, FeeShare=?, DelegatedState=?, ActivationRound=?, DeactivationRound=?, Active=?, Status=?, PricePerPixel=?, UpdateAt=? WHERE Address=?")
			_, err = statement.Exec(x.ServiceURI, x.LastRewardRound, x.RewardCut, x.FeeShare, x.DelegatedStake, x.ActivationRound, x.DeactivationRound, x.Active, x.Status, x.PricePerPixel, time.Now().Unix(), x.Address)
			if err != nil {
				fmt.Println("Error in updating orchestrator data.")
			}
		}
		statement, _ = database.Prepare("INSERT INTO PriceHistory (Address, Time, PricePerPixel) VALUES (?, ?, ?)")
		_, err = statement.Exec(x.Address, time.Now().Unix(), x.PricePerPixel)
		if err != nil {
			fmt.Println("Error in inserting data in PriceHistory Table", err.Error())
		}
	}
}

// Method to get data from Orchestrator Table
func GetOrchestratorStats(w http.ResponseWriter, req *http.Request) {
	database, _ := sql.Open("sqlite3", "./pricing_tool.db")
	rows, _ := database.Query("SELECT * FROM Orchestrators")
	orchestrators := []OrchestratorStats{}
	x := OrchestratorStats{}
	for rows.Next() {
		rows.Scan(&x.Address, &x.ServiceURI, &x.LastRewardRound, &x.RewardCut, &x.FeeShare, &x.DelegatedStake, &x.ActivationRound, &x.DeactivationRound, &x.Active, &x.Status, &x.PricePerPixel, &x.UpdatedAt)
		orchestrators = append(orchestrators, x)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(orchestrators)
}

// Method to get data from PriceHistory Table based on orchestrator ID
func GetOrchestratorPriceHistory(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	database, _ := sql.Open("sqlite3", "./pricing_tool.db")
	rows, _ := database.Query("SELECT * FROM PriceHistory WHERE Address=? ORDER BY Time DESC", params["address"])
	data := []PriceHistory{}
	x := PriceHistory{}
	for rows.Next() {
		rows.Scan(&x.Address, &x.Time, &x.PricePerPixel)
		data = append(data, x)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {

	initializeDB()

	fetchAndStore()

	router := mux.NewRouter()
	router.HandleFunc("/orchestratorStats", GetOrchestratorStats).Methods("GET")
	router.HandleFunc("/priceHistory/{address}", GetOrchestratorPriceHistory).Methods("GET")
	http.ListenAndServe(":12345", router)

}
