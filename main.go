package main

import (
	// "log"

	"github.com/rakesh/banking/app/app"
	"github.com/rakesh/banking/app/logger"
)

func main() {
	// log.Println("starting our application....")
	logger.Info("starting our application....")
	app.Start()
}
