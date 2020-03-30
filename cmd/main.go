package main

import (
	"github.com/SergeyDavidenko/subscription/api"
	"github.com/SergeyDavidenko/subscription/utils"
	log "github.com/sirupsen/logrus"
)

func init() {
	utils.InitServer("")
}
func main() {
	log.Info("Start app subscription")
	api.WEBServerRun()
}
