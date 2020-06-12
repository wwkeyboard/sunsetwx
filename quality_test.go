package quality

import (
	"testing"
)

func TestFromJson(t *testing.T) {
	fc, err := FromJSON([]byte(sample()))
	if err != nil {
		t.Error(err)
	}

	if len(fc.Features) != 1 {
		t.Errorf("didn't parse the correct number of Features")
	}

}

func sample() string {
	samp := `{
		"type": "FeatureCollection",
		"features": [
		  {
			"type": "Feature",
			"geometry": {
			  "type": "Point",
			  "coordinates": [
				-88.2183,
				40.1225
			  ]
			},
			"properties": {
			  "type": "Sunrise",
			  "quality": "Poor",
			  "quality_percent": 23.14,
			  "quality_value": -271.859,
			  "temperature": 16.54,
			  "last_updated": "2020-06-11T18:00:00Z",
			  "imported_at": "2020-06-11T21:30:48Z",
			  "dawn": {
				"astronomical": "2020-06-12T08:20:00Z",
				"nautical": "2020-06-12T09:09:00Z",
				"civil": "2020-06-12T09:50:00Z"
			  },
			  "valid_at": "2020-06-12T10:23:00Z",
			  "source": "NAM",
			  "distance": 1.587
			}
		  }
		]
	  }`
	return samp
}
