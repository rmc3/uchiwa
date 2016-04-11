package uchiwa

import "github.com/rmc3/uchiwa/uchiwa/logger"

// ResolveEvent sends a DELETE request in order to
// resolve an event for a given check on a given client
func (u *Uchiwa) ResolveEvent(check, client, dc string) error {
	api, err := getAPI(u.Datacenters, dc)
	if err != nil {
		logger.Warning(err)
		return err
	}

	err = api.DeleteEvent(check, client)
	if err != nil {
		logger.Warning(err)
		return err
	}

	return nil
}
