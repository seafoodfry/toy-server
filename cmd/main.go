package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/fundoplicatedFundus/toy-server/cmd/app"
)

func main() {
	log.SetLevel(log.DebugLevel)

	command := app.NewServerCommand()
	if err := command.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
