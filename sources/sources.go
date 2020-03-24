package sources

import (
	"github.com/fnlbhq/fred/query"
)

const (
	sources  = "fred/sources"         // Get all sources of economic data.
	source   = "fred/source"          // Get a source of economic data.
	releases = "fred/source/releases" // Get the releases for a source.
)

func Sources() query.Query {
	return query.NewQuery(sources)
}

func Source() query.Query {
	return query.NewQuery(source)
}

func Releases() query.Query {
	return query.NewQuery(releases)
}
