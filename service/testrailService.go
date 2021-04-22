package util

import (
	"github.com/educlos/testrail"
	"github.com/nansuri/gp-server/config"
)

func TestrailBridge(email string, password string) *testrail.Client {

	return testrail.NewClient(config.TestrailUrl, email, password)
}
