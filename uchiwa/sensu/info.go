package sensu

import (
	"encoding/json"
	"fmt"

	"github.com/rmc3/uchiwa/uchiwa/structs"
)

// GetInfo returns a pointer to a structs.Info struct containing the
// Sensu version and the transport and Redis connection information
func (s *Sensu) GetInfo() (*structs.Info, error) {
	body, _, err := s.getBytes("info")
	if err != nil {
		return nil, err
	}

	var info structs.Info
	if err := json.Unmarshal(body, &info); err != nil {
		return nil, fmt.Errorf("Parsing JSON-encoded response body: %v", err)
	}

	return &info, nil
}
