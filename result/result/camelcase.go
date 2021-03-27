package result

import "encoding/json"

type Result struct {
	Start            string        `json:"realtimeStart" gorm:"column:realtime_start"`
	End              string        `json:"realtimeEnd" gorm:"column:realtime_end"`
	ObservationStart string        `json:"observationStart" gorm:"column:observation_start"`
	ObservationEnd   string        `json:"observationEnd" gorm:"column:observation_end"`
	Units            string        `json:"units" gorm:"column:units"`
	OutputType       int           `json:"outputType" gorm:"column:output_type"`
	FileType         string        `json:"fileType" gorm:"column:file_type"`
	OrderBy          string        `json:"orderBy" gorm:"column:order_by"`
	SortOrder        string        `json:"sortOrder" gorm:"column:sort_order"`
	Count            int           `json:"count" gorm:"column:count"`
	Offset           int           `json:"offset" gorm:"column:offset"`
	Limit            int           `json:"limit" gorm:"column:limit"`
	Series           []Series      `json:"series,omitempty" gorm:"column:series"`
	Observations     []Observation `json:"observations,omitempty" gorm:"column:observations"`
	Releases         []Release     `json:"releases,omitempty" gorm:"column:releases"`
	Categories       []Category    `json:"categories,omitempty" gorm:"column:categories"`
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
	ID                      string `json:"id" csv:"id" gorm:"column:id"`
	RealtimeStart           string `json:"realtimeStart" csv:"realtimeStart" gorm:"column:realtime_start"`
	RealtimeEnd             string `json:"realtimeEnd" csv:"realtimeEnd" gorm:"column:realtime_end"`
	Title                   string `json:"title" csv:"title" gorm:"column:title"`
	ObservationStart        string `json:"observationStart" csv:"observationStart" gorm:"column:observation_start"`
	ObservationEnd          string `json:"observationEnd" csv:"observationEnd" gorm:"column:observation_end"`
	Frequency               string `json:"frequency" csv:"frequency" gorm:"column:frequency"`
	FrequencyShort          string `json:"frequencyShort" csv:"frequencyShort" gorm:"column:frequency_short"`
	Units                   string `json:"units" csv:"units" gorm:"column:units"`
	UnitsShort              string `json:"unitsShort" csv:"unitsShort" gorm:"column:units_short"`
	SeasonalAdjustment      string `json:"seasonalAdjustment" csv:"seasonalAdjustment" gorm:"column:seasonal_adjustment"`
	SeasonalAdjustmentShort string `json:"seasonalAdjustmentShort" csv:"seasonalAdjustmentShort" gorm:"column:seasonal_adjustment_short"`
	LastUpdated             string `json:"lastUpdated" csv:"lastUpdated" gorm:"column:last_updated"`
	Popularity              int    `json:"popularity" csv:"popularity" gorm:"column:popularity"`
	Notes                   string `json:"notes" gorm:"column:notes"`
}

type Observation struct {
	Date          string `json:"date" gorm:"column:date"`
	RealtimeStart string `json:"realtimeStart" gorm:"column:realtime_start"`
	RealtimeEnd   string `json:"realtimeEnd" gorm:"column:realtime_end"`
	Value         string `json:"value" gorm:"column:value"`
}

type Release struct {
	ID            int `json:"id" gorm:"column:id"`
	RealtimeStart string `json:"realtimeStart" gorm:"column:realtime_start"`
	RealtimeEnd   string `json:"realtimeEnd" gorm:"column:realtime_end"`
	Name          string `json:"name" gorm:"column:name"`
	PressRelease  string `json:"pressRelease" gorm:"column:press_release"`
	Link          string `json:"link" gorm:"column:link"`
}

type Category struct {
	ID       int    `json:"id" gorm:"column:id"`
	Name     string `json:"name" gorm:"column:name"`
	ParentID int    `json:"parentId" gorm:"column:parent_id"`
}

