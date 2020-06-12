package quality

import (
	"encoding/json"
	"time"
)

// FeatureCollection returned from the SunsetWX API
type FeatureCollection struct {
	Type     string     `json:"type"`
	Features []Features `json:"features"`
}

// Geometry the feature is valid for
type Geometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Dawn times
type Dawn struct {
	Astronomical time.Time `json:"astronomical"`
	Nautical     time.Time `json:"nautical"`
	Civil        time.Time `json:"civil"`
}

// Properties of the Feature
type Properties struct {
	Type           string    `json:"type"`
	Quality        string    `json:"quality"`
	QualityPercent float64   `json:"quality_percent"`
	QualityValue   float64   `json:"quality_value"`
	Temperature    float64   `json:"temperature"`
	LastUpdated    time.Time `json:"last_updated"`
	ImportedAt     time.Time `json:"imported_at"`
	Dawn           Dawn      `json:"dawn"`
	ValidAt        time.Time `json:"valid_at"`
	Source         string    `json:"source"`
	Distance       float64   `json:"distance"`
}

// Features that have a quality estimate about them
type Features struct {
	Type       string     `json:"type"`
	Geometry   Geometry   `json:"geometry"`
	Properties Properties `json:"properties"`
}

//FromJSON create a FeatureCollection
func FromJSON(data []byte) (*FeatureCollection, error) {
	var fc FeatureCollection
	err := json.Unmarshal(data, &fc)
	if err != nil {
		return nil, err
	}
	return &fc, nil
}
