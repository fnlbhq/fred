package fred

import (
	"github.com/fnlbhq/fred/categories"
	"github.com/fnlbhq/fred/query"
	"github.com/fnlbhq/fred/releases"
	"github.com/fnlbhq/fred/result"
	"github.com/fnlbhq/fred/series"
	"github.com/fnlbhq/fred/sources"
	"github.com/fnlbhq/fred/tags"
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

type Query = query.Query
type Result = result.Result

// Categories
var Category = categories.Category
var CategoryChildren = categories.Children
var CategoryRelated = categories.Related
var CategorySeries = categories.Series
var CategoryTags = categories.Tags
var CategoryRelatedTags = categories.RelatedTags

// Releases
var Releases = releases.Releases
var ReleasesAllReleaseDates = releases.AllReleasesDates
var ReleasesRelease = releases.Release
var ReleasesReleaseDates = releases.ReleaseDates
var ReleasesReleaseSeries = releases.ReleaseSeries
var ReleasesReleaseSources = releases.ReleaseSources
var ReleasesReleaseTags = releases.ReleaseTags
var ReleasesRelatedTags = releases.RelatedTags
var ReleasesReleaseTables = releases.ReleaseTables

// Series
var Series = series.Series
var SeriesCategories = series.Categories
var SeriesObservations = series.Observations
var SeriesRelease = series.Release
var SeriesSearch = series.Search
var SeriesSearchTags = series.SearchTags
var SeriesSearchRelatedTags = series.SearchRelatedTags
var SeriesTags = series.Tags
var SeriesUpdates = series.Updates
var SeriesVintageDates = series.VintageDates

// Sources
var Sources = sources.Sources
var SourcesSource = sources.Source
var SourcesReleases = sources.Releases

// Tags
var Tags = tags.Tags
var TagsRelatedTags = tags.RelatedTags
var TagsSeries = tags.Series
