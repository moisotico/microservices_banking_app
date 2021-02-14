package main

import (
	"github.com/moisotico/banking/app"
	"github.com/moisotico/banking/logger"
)

func main() {
	//log.Println("Starting our application...")
	//Log using Zap
	logger.Info("Starting the application...")
	app.Start()
}
