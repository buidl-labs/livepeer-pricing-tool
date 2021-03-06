package dataservice

import (
	"database/sql"
	"time"

	"api/model"

	_ "github.com/mattn/go-sqlite3"
	log "github.com/sirupsen/logrus"
)

const dbFilePath = "./PricingToolDB.sqlite3"

var sqldb *sql.DB

// Initializes the pricing tool database
func DBInit() {

	database, err := sql.Open("sqlite3", dbFilePath)
	if err != nil {
		log.Fatalln("Error in creating DB", err.Error())
	}
	sqldb = database

	statement, err := database.Prepare(`
		CREATE TABLE IF NOT EXISTS Orchestrators (
			Address TEXT PRIMARY KEY, 
			ServiceURI TEXT, 
			LastRewardRound INTEGER, 
			RewardCut INTEGER, 
			FeeShare INTEGER, 
			DelegatedState TEXT, 
			ActivationRound INTEGER, 
			DeactivationRound TEXT, 
			Active INTEGER, 
			Status TEXT, 
			PricePerPixel STRING, 
			UpdatedAt INTEGER
		)
	`)
	if err != nil {
		log.Fatalln("Error in creating DB", err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln("Error in creating DB", err.Error())
	}

	statement, err = database.Prepare(`
		CREATE TABLE IF NOT EXISTS PriceHistory (
			Address TEXT, 
			Time INTEGER, 
			PricePerPixel STRING
		)
	`)
	if err != nil {
		log.Fatalln("Error in creating DB", err.Error())
	}
	_, err = statement.Exec()
	if err != nil {
		log.Fatalln("Error in creating DB", err.Error())
	}

	log.Info("DB created successfully.")
}

// Adds orchestrator statistics to the database
func InsertOrchestrator(x model.Orchestrator) {
	statement, err := sqldb.Prepare("INSERT INTO Orchestrators (Address, ServiceURI, LastRewardRound, RewardCut, FeeShare, DelegatedState, ActivationRound, DeactivationRound, Active, Status, PricePerPixel, UpdatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Errorln("Error in inserting orchestrator", x.Address)
		log.Errorln(err.Error())
	}
	_, err = statement.Exec(x.Address, x.ServiceURI, x.LastRewardRound, x.RewardCut, x.FeeShare, x.DelegatedStake.String(), x.ActivationRound, x.DeactivationRound.String(), x.Active, x.Status, x.PricePerPixel, time.Now().Unix())
	if err != nil {
		log.Errorln("Error in inserting orchestrator", x.Address)
		log.Errorln(err.Error())
	}
}

// Updates orchestrator statistics in the database
func UpdateOrchestrator(x model.Orchestrator) {
	statement, err := sqldb.Prepare("UPDATE Orchestrators SET ServiceURI=?, LastRewardRound=?, RewardCut=?, FeeShare=?, DelegatedState=?, ActivationRound=?, DeactivationRound=?, Active=?, Status=?, PricePerPixel=?, UpdatedAt=? WHERE Address=?")
	if err != nil {
		log.Errorln("Error in updating orchestrator", x.Address)
		log.Errorln(err.Error())
	}
	_, err = statement.Exec(x.ServiceURI, x.LastRewardRound, x.RewardCut, x.FeeShare, x.DelegatedStake.String(), x.ActivationRound, x.DeactivationRound.String(), x.Active, x.Status, x.PricePerPixel, time.Now().Unix(), x.Address)
	if err != nil {
		log.Errorln("Error in updating orchestrator", x.Address)
		log.Errorln(err.Error())
	}
}

// Add price history to the database
func InsertPriceHistory(x model.Orchestrator) {
	statement, err := sqldb.Prepare("INSERT INTO PriceHistory (Address, Time, PricePerPixel) VALUES (?, ?, ?)")
	if err != nil {
		log.Errorln("Error in inserting price history", x.Address)
		log.Errorln(err.Error())
	}
	_, err = statement.Exec(x.Address, time.Now().Unix(), x.PricePerPixel)
	if err != nil {
		log.Errorln("Error in inserting price history", x.Address)
		log.Errorln(err.Error())
	}
}

// Fetching orchestrator statistics
func FetchOrchestratorStatistics(excludeUnavailable bool) []model.DBOrchestrator {

	rows, err := sqldb.Query("SELECT * FROM Orchestrators")
	if err != nil {
		log.Errorln("Error in fetching orchestrator statistics")
		log.Errorln(err.Error())
	}
	orchestrators := []model.DBOrchestrator{}
	x := model.DBOrchestrator{}
	for rows.Next() {
		rows.Scan(&x.Address, &x.ServiceURI, &x.LastRewardRound, &x.RewardCut, &x.FeeShare, &x.DelegatedStake, &x.ActivationRound, &x.DeactivationRound, &x.Active, &x.Status, &x.PricePerPixel, &x.UpdatedAt)
		if excludeUnavailable == true && x.PricePerPixel == "0" {
			continue
		}
		orchestrators = append(orchestrators, x)
	}
	return orchestrators
}

// Fetcing pricing history
func FetchPricingHistory(address string) []model.DBPriceHistory {

	rows, err := sqldb.Query("SELECT * FROM PriceHistory WHERE Address=? ORDER BY Time DESC", address)
	if err != nil {
		log.Errorln("Error in inserting price history for", address)
		log.Errorln(err.Error())
	}
	data := []model.DBPriceHistory{}
	x := model.DBPriceHistory{}
	for rows.Next() {
		rows.Scan(&x.Address, &x.Time, &x.PricePerPixel)
		data = append(data, x)
	}
	return data
}

// checking for existence of an orchestrator in table
func IfOrchestratorExists(address string) bool {
	count := 0
	rows, err := sqldb.Query("SELECT * FROM Orchestrators WHERE Address=?", address)
	if err != nil {
		log.Errorln("Error in checking existence of orchestrator", address)
		log.Errorln(err.Error())
	}
	for rows.Next() {
		count += 1
	}
	if count == 0 {
		return false
	} else {
		return true
	}
}
