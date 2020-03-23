package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/fnlbhq/fred/query/argument"
)

const fredUrl = "https://api.stlouisfed.org"
const apiKeyEnvVar = "FRED_API_KEY"

func NewQuery(action string) *Query {

	u, _ := url.Parse(fmt.Sprintf("%s/%s", fredUrl, action))

	apiKey := os.Getenv(apiKeyEnvVar)

	q := u.Query()
	q.Add("file_type", "json")

	if apiKey != "" {
		q.Add(argument.APIKey, apiKey)
	}

	u.RawQuery = q.Encode()

	return &Query{URL: *u}
}

type Query struct {
	url.URL
}

func (q *Query) AddParameter(argument, value string) *Query {
	query := q.URL.Query()
	query.Set(argument, fmt.Sprintf("%v", value))
	q.URL.RawQuery = query.Encode()
	return q
}

func (q *Query) APIKey(value string) *Query {
	return q.AddParameter(argument.APIKey, value)
}

func (q *Query) CategoryID(value string) *Query {
	return q.AddParameter(argument.CategoryId, value)
}

func (q *Query) RealtimeStart(value string) *Query {
	return q.AddParameter(argument.RealTimeStart, value)
}

func (q *Query) RealtimeEnd(value string) *Query {
	return q.AddParameter(argument.RealTimeEnd, value)
}

func (q *Query) Limit(value string) *Query {
	return q.AddParameter(argument.Limit, value)
}

func (q *Query) Offset(value string) *Query {
	return q.AddParameter(argument.Offset, value)
}

func (q *Query) OrderBy(value string) *Query {
	return q.AddParameter(argument.OrderBy, value)
}

func (q *Query) SortOrder(value string) *Query {
	return q.AddParameter(argument.SortOrder, value)
}

func (q *Query) FilterVariable(value string) *Query {
	return q.AddParameter(argument.FilterVariable, value)
}

func (q *Query) FilterValue(value string) *Query {
	return q.AddParameter(argument.FilterValue, value)
}

func (q *Query) TagNames(value string) *Query {
	return q.AddParameter(argument.TagNames, value)
}

func (q *Query) ExcludeTagNames(value string) *Query {
	return q.AddParameter(argument.ExcludeTagNames, value)
}

func (q *Query) TagGroupID(value string) *Query {
	return q.AddParameter(argument.TagGroupId, value)
}

func (q *Query) SearchText(value string) *Query {
	return q.AddParameter(argument.SearchText, value)
}

func (q *Query) IncludeReleaseDatesWithNoData(value string) *Query {
	return q.AddParameter(argument.IncludeReleaseDatesWithNoData, value)
}

func (q *Query) ReleaseID(value string) *Query {
	return q.AddParameter(argument.ReleaseId, value)
}

func (q *Query) ElementID(value string) *Query {
	return q.AddParameter(argument.ElementId, value)
}

func (q *Query) IncludeObservationValues(value string) *Query {
	return q.AddParameter(argument.IncludeObservationValues, value)
}

func (q *Query) ObservationDate(value string) *Query {
	return q.AddParameter(argument.ObservationDate, value)
}

func (q *Query) SeriesID(value string) *Query {
	return q.AddParameter(argument.SeriesId, value)
}

func (q *Query) ObservationStart(value string) *Query {
	return q.AddParameter(argument.ObservationStart, value)
}

func (q *Query) ObservationEnd(value string) *Query {
	return q.AddParameter(argument.ObservationEnd, value)
}

func (q *Query) Units(value string) *Query {
	return q.AddParameter(argument.Units, value)
}

func (q *Query) Frequency(value string) *Query {
	return q.AddParameter(argument.Frequency, value)
}

func (q *Query) AggregationMethod(value string) *Query {
	return q.AddParameter(argument.AggregationMethod, value)
}

func (q *Query) OutputType(value string) *Query {
	return q.AddParameter(argument.OutputType, value)
}

func (q *Query) VintageDates(value string) *Query {
	return q.AddParameter(argument.VintageDates, value)
}

func (q *Query) TagSearchText(value string) *Query {
	return q.AddParameter(argument.TagSearchText, value)
}

func (q *Query) SeriesSearchText(value string) *Query {
	return q.AddParameter(argument.SeriesSearchText, value)
}

func (q *Query) SourceID(value string) *Query {
	return q.AddParameter(argument.SourceId, value)
}

func (q *Query) String() string {
	return q.URL.String()
}

func (q *Query) Get() (*Result, error) {
	resp, err := http.Get(q.URL.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result Result

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}

type Result struct {
	Start            string `json:"realtime_start"`
	End              string `json:"realtime_end"`
	ObservationStart string `json:"observation_start"`
	ObservationEnd   string `json:"observation_end"`
	Units            string
	OutputType       int    `json:"output_type"`
	FileType         string `json:"file_type"`
	OrderBy          string `json:"order_by"`
	SortOrder        string `json:"sort_order"`
	Count            int
	Offset           int
	Limit            int
	Series           []Series      `json:"seriess,omitempty"`
	Observations     []Observation `json:",omitempty"`
	Releases         []Release     `json:",omitempty"`
	Categories       []Category    `json:",omitempty"`
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
	ID                      string `json:"id"`
	RealtimeStart           string `json:"realtime_start"`
	RealtimeEnd             string `json:"realtime_end"`
	Title                   string
	ObservationStart        string `json:"observation_start"`
	ObservationEnd          string `json:"observation_end"`
	Frequency               string
	FrequencyShort          string `json:"frequency_short"`
	Units                   string
	UnitsShort              string `json:"units_short"`
	SeasonalAdjustment      string `json:"seasonal_adjustment"`
	SeasonalAdjustmentShort string `json:"seasonal_adjustment_short"`
	LastUpdated             string `json:"last_updated"`
	Popularity              int
}

type Observation struct {
	Date          string
	RealtimeStart string `json:"realtime_start"`
	RealtimeEnd   string `json:"realtime_end"`
	Value         string
}

type Release struct {
	ID            int
	RealtimeStart string `json:"realtime_start"`
	RealtimeEnd   string `json:"realtime_end"`
	Name          string
	PressRelease  string `json:"press_release"`
	Link          string
}

type Category struct {
	ID       int
	Name     string
	ParentID int
}

// Common series
const (
	CDRatesNonJumbo                = "MMNRNJ"          // FDIC via FRED
	CDRatesJumbo                   = "MMNRJD"          // FDIC via FRED
	RealGDP                        = "A191RL1Q225SBEA" // BEA via FRED
	ConsumerPriceIndex             = "CPIAUCSL"        // Board of Governors of the Federal Reserve System
	CreditCardInterestRate         = "TERMCBCCALLNS"   // Board of Governors of the Federal Reserve System
	FederalFundsRate               = "FEDFUNDS"        // Board of Governors of the Federal Reserve System
	InitialClaimsFourWeekMovingAvg = "IC4WSA"          // US Employment & Training Admin via FRED
	IndustrialProductionIndex      = "INDPRO"          // Board of Governors of the Federal Reserve System
	InstitutionalMoneyFunds        = "WIMFSL"
	MortgageRates30USFixedAverage  = "MORTGAGE30US"  // Freddie Mac via Board of Governors of the Federal Reserve System
	MortgageRates15USFixedAverage  = "MORTGAGE15US"  // Freddie Mac via Board of Governors of the Federal Reserve System
	MortgageRates5USFixedAverage   = "MORTGAGE5US"   // Freddie Mac via Board of Governors of the Federal Reserve System
	TotalHousingStarts             = "HOUST"         // U.S. Census Bureau and U.S. Department of Housing and Urban Development
	TotalPayrolls                  = "PAYEMS"        // U.S. Bureau of Labor Statistics
	TotalVehicleSales              = "TOTALSA"       // U.S. Bureau of Economic Analysis
	RetailMoneyFunds               = "WRMFSL"        // Board of Governors of the Federal Reserve System
	UnemploymentRate               = "UNRATE"        // U.S. Bureau of Labor Statistics
	USRecessionProbabilities       = "RECPROUSM156N" // U.S. Bureau of Economic Analysis
)
