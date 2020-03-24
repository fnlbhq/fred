package releases

import (
	"github.com/fnlbhq/fred/query"
)

const (
	releases         = "fred/releases"             // Get all releases of economic data.
	allReleasesDates = "fred/releases/dates"       // Get release dates for all releases of economic data.
	release          = "fred/release"              // Get a release of economic data.
	releaseDates     = "fred/release/dates"        // Get release dates for a release of economic data.
	releaseSeries    = "fred/release/series"       // Get the series on a release of economic data.
	releaseSources   = "fred/release/sources"      // Get the sources for a release of economic data.
	releaseTags      = "fred/release/tags"         // Get the tags for a release.
	relatedTags      = "fred/release/related_tags" // Get the related tags for a release.
	releaseTables    = "fred/release/tables"       // Get the release tables for a given release.
)

func Releases() query.Query {
	return query.NewQuery(releases)
}

func AllReleasesDates() query.Query {
	return query.NewQuery(allReleasesDates)
}

func Release() query.Query {
	return query.NewQuery(release)
}

func ReleaseDates() query.Query {
	return query.NewQuery(releaseDates)
}

func ReleaseSeries() query.Query {
	return query.NewQuery(releaseSeries)
}

func ReleaseSources() query.Query {
	return query.NewQuery(releaseSources)
}

func ReleaseTags() query.Query {
	return query.NewQuery(releaseTags)
}

func RelatedTags() query.Query {
	return query.NewQuery(relatedTags)
}

func ReleaseTables() query.Query {
	return query.NewQuery(releaseTables)
}
