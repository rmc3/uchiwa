package uchiwa

import (
	"github.com/rmc3/uchiwa/uchiwa/logger"
	"github.com/rmc3/uchiwa/uchiwa/structs"
)

// IssueCheckExecution sends a POST request to the /stashes endpoint in order to create a stash
func (u *Uchiwa) IssueCheckExecution(data structs.CheckExecution) error {
	api, err := getAPI(u.Datacenters, data.Dc)
	if err != nil {
		logger.Warning(err)
		return err
	}

	_, err = api.IssueCheckExecution(data)
	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}
