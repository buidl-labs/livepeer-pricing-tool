package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

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

func main() {

	initializeDB()

	fetchAndStore()

}
