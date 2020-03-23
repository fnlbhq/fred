package series

import (
	"github.com/fnlbhq/fred/query"
)

const (
	series            = "fred/series"                     // Get an economic data series.
	categories        = "fred/series/categories"          // Get the categories for an economic data series.
	observations      = "fred/series/observations"        // Get the observations or data values for an economic data
	release           = "fred/series/release"             // Get the release for an economic data series.
	search            = "fred/series/search"              // Get economic data series that match keywords.
	searchTags        = "fred/series/search/tags"         // Get the tags for a series search.
	searchRelatedTags = "fred/series/search/related_tags" // Get the related tags for a series search.
	tags              = "fred/series/tags"                // Get the tags for an economic data series.
	updates           = "fred/series/updates"             // Get economic data series sorted by when observations were updated on the FRED® server.
	vintageDates      = "fred/series/vintagedates"        // Get the dates in history when a series' data values were revised or new data values were released.
)

func Series() *query.Query {
	return query.NewQuery(series)
}

func Categories() *query.Query {
	return query.NewQuery(categories)
}

func Observations() *query.Query {
	return query.NewQuery(observations)
}

func Release() *query.Query {
	return query.NewQuery(release)
}

func Search() *query.Query {
	return query.NewQuery(search)
}

func SearchTags() *query.Query {
	return query.NewQuery(searchTags)
}

func SearchRelatedTags() *query.Query {
	return query.NewQuery(searchRelatedTags)
}

func Tags() *query.Query {
	return query.NewQuery(tags)
}

func Updates() *query.Query {
	return query.NewQuery(updates)
}

func VintageDates() *query.Query {
	return query.NewQuery(vintageDates)
}
