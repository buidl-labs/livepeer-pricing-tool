package main

import (
	"os"
	
	"server/src/dbhandler"
	"server/src/datahandler"
	"server/src/server"
	
	log "github.com/sirupsen/logrus"
)

func main() {

	// Log file configurations
	var filename string = "PricingToolServer.log"
	
	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	
	log.SetFormatter(Formatter)
    if err != nil {
        log.Fatalln("Error in configuring log file.", err.Error())
    } else{
        log.SetOutput(file)
    }

	log.Info("Log file configured succesfully.")

	// Initialize the pricing tool DB
	dbhandler.DBInit()
	
	// Start a concurrent service for period polling of data from broadcaster endpoint
	go datahandler.PollForData()

	// Start the API server
	server.StartServer(":9000")
}