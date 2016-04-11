package daemon

import (
	"testing"

	"github.com/rmc3/uchiwa/uchiwa/structs"
	"github.com/stretchr/testify/assert"
)

func TestRawMetricsToAggregatedCoordinates(t *testing.T) {
	// Normal metrics
	raw1 := structs.SERawMetric{
		Points: [][]interface{}{{1000000000.0, 0.5}, {1000000001.0, 1.0}, {1000000002.0, 1.0}},
	}
	raw2 := structs.SERawMetric{
		Points: [][]interface{}{{1000000000.0, 2.0}, {1000000001.0, 3.0}, {1000000002.0, 4.0}},
	}
	raw3 := structs.SERawMetric{
		Points: [][]interface{}{{1000000000.0, 1.0}, {1000000001.0, 0.0}, {1000000002.0, 0.0}, {1000000003.0, 2.0}},
	}
	raw4 := structs.SERawMetric{
		Points: [][]interface{}{{1000000002.0, 2.5}},
	}

	metrics := []*structs.SERawMetric{&raw1, &raw2, &raw3, &raw4}

	expectedCoordinates := structs.SEMetric{
		Data: []structs.XY{
			{X: 1000000000000, Y: 3.5},
			{X: 1000000001000, Y: 4},
			{X: 1000000002000, Y: 7.5},
		},
	}

	coordinates := rawMetricsToAggregatedCoordinates(metrics)
	assert.Equal(t, &expectedCoordinates, coordinates)

	// Empty metrics
	metrics = []*structs.SERawMetric{&structs.SERawMetric{}}
	coordinates = rawMetricsToAggregatedCoordinates(metrics)
	assert.Equal(t, &structs.SEMetric{Data: []structs.XY{}}, coordinates)

	// Invalid point
	metrics = []*structs.SERawMetric{&structs.SERawMetric{Points: [][]interface{}{{100000000.0}}}}
	coordinates = rawMetricsToAggregatedCoordinates(metrics)
	assert.Equal(t, &structs.SEMetric{Data: []structs.XY{structs.XY{X: 0, Y: 0}}}, coordinates)

	// Single point metrics
	metrics = []*structs.SERawMetric{&raw4}
	coordinates = rawMetricsToAggregatedCoordinates(metrics)
	assert.Equal(t, &structs.SEMetric{Data: []structs.XY{structs.XY{X: 1.000000002e+12, Y: 2.5}}}, coordinates)
}
