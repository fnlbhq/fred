package fred

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/fnlbhq/fred/series"

	"github.com/fnlbhq/fred/argument"
)

const fredUrl = "https://api.stlouisfed.org"
const apiKeyEnvVar = "FRED_API_KEY"

func NewQuery(action string) (*Query, error) {

	u, err := url.Parse(fmt.Sprintf("%s/%s", fredUrl, action))

	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv(apiKeyEnvVar)

	query := u.Query()
	query.Add("file_type", "json")

	if apiKey != "" {
		query.Add(argument.APIKey, apiKey)
	}

	u.RawQuery = query.Encode()

	return &Query{URL: *u}, nil
}

type Query struct {
	url.URL
}

func (q *Query) With(argument, value string) *Query {
	query := q.URL.Query()
	query.Add(argument, fmt.Sprintf("%v", value))
	q.URL.RawQuery = query.Encode()
	copy, _ := url.Parse(q.URL.String())
	return &Query{URL: *copy}
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

func SeriesInRealtimeRange(seriesId, start, end string) (*Result, error) {
	q, err := NewQuery(series.Observations)

	if err != nil {
		return nil, err
	}

	return q.With(argument.SeriesId, seriesId).
		With(argument.RealTimeStart, start).
		With(argument.RealTimeEnd, end).
		Get()
}

func Series(seriesId string) (*Result, error) {
	q, err := NewQuery(series.Observations)

	if err != nil {
		return nil, err
	}

	return q.With(argument.SeriesId, seriesId).Get()
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

func Updates() (*Result, error) {
	return Series(series.Updates)
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
