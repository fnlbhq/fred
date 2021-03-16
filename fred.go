package fred

import (
	"fmt"
	"github.com/fnlbhq/fred/query/argument"
	"github.com/fnlbhq/fred/result"

	"github.com/fnlbhq/fred/series"
)

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

func GetUpdates(offset, limit string) string {
	r := series.Updates().Offset(offset).Limit(limit)
	fmt.Println(r.String())
	rr, _ := r.Get()
	j, _ := rr.JSON()
	return j
}

// Assumes that the Fred API is stored in an environment variable: FRED_API_KEY
func GetSeriesObservations(seriesId string) (*result.Result, error) {
	result, err := series.
		Observations().
		AddParameter(argument.SeriesId, seriesId).
		Get()

	if err != nil {
		return nil, err
	}
	return result, nil
}
