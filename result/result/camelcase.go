package result

import "encoding/json"

type Result struct {
	Start            string        `json:"realtimeStart"`
	End              string        `json:"realtimeEnd"`
	ObservationStart string        `json:"observationStart"`
	ObservationEnd   string        `json:"observationEnd"`
	Units            string        `json:"units"`
	OutputType       int           `json:"outputType"`
	FileType         string        `json:"fileType"`
	OrderBy          string        `json:"orderBy"`
	SortOrder        string        `json:"sortOrder"`
	Count            int           `json:"count"`
	Offset           int           `json:"offset"`
	Limit            int           `json:"limit"`
	Series           []Series      `json:"series,omitempty"`
	Observations     []Observation `json:"observations,omitempty"`
	Releases         []Release     `json:"releases,omitempty"`
	Categories       []Category    `json:"categories,omitempty"`
}

func (r *Result) JSON() (string, error) {
	jsonData, err := json.Marshal(r)

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

func (r *Result) PrettyJSON() (string, error) {
	jsonData, err := json.MarshalIndent(r, "", " ")

	if err != nil {
		return "", err
	}

	return string(jsonData), nil
}

type Series struct {
	ID                      string `json:"id" csv:"id"`
	RealtimeStart           string `json:"realtimeStart" csv:"realtimeStart"`
	RealtimeEnd             string `json:"realtimeEnd" csv:"realtimeEnd"`
	Title                   string `json:"title" csv:"title"`
	ObservationStart        string `json:"observationStart" csv:"observationStart"`
	ObservationEnd          string `json:"observationEnd" csv:"observationEnd"`
	Frequency               string `json:"frequency" csv:"frequency"`
	FrequencyShort          string `json:"frequencyShort" csv:"frequencyShort"`
	Units                   string `json:"units" csv:"units"`
	UnitsShort              string `json:"unitsShort" csv:"unitsShort"`
	SeasonalAdjustment      string `json:"seasonalAdjustment" csv:"seasonalAdjustment"`
	SeasonalAdjustmentShort string `json:"seasonalAdjustmentShort" csv:"seasonalAdjustmentShort"`
	LastUpdated             string `json:"lastUpdated" csv:"lastUpdated"`
	Popularity              int    `json:"popularity" csv:"popularity"`
	Notes                   string `json:"notes"`
}

type Observation struct {
	Date          string `json:"date"`
	RealtimeStart string `json:"realtimeStart"`
	RealtimeEnd   string `json:"realtimeEnd"`
	Value         string `json:"value"`
}

type Release struct {
	ID            int
	RealtimeStart string `json:"realtimeStart"`
	RealtimeEnd   string `json:"realtimeEnd"`
	Name          string `json:"name"`
	PressRelease  string `json:"pressRelease"`
	Link          string `json:"link"`
}

type Category struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ParentID int    `json:"parentId"`
}

