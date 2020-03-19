package series

const (
	Series = "fred/series"                         // Get an economic data series.
	Categories = "fred/series/categories"          // Get the categories for an economic data series.
	Observations = "fred/series/observations"      // Get the observations or data values for an economic data
	Release = "fred/series/release"                // Get the release for an economic data series.
	Search = "fred/series/search"                  // Get economic data series that match keywords.
	SearchTags = "fred/series/search/tags"         // Get the tags for a series search.
	SearchRelatedTags = "fred/series/search/related_tags" // Get the related tags for a series search.
	Tags = "fred/series/tags"                      // Get the tags for an economic data series.
	Updates = "fred/series/updates"                // Get economic data series sorted by when observations were updated on the FRED® server.
	VintageDates = "fred/series/vintagedates"           // Get the dates in history when a series' data values were revised or new data values were released.
)