package tags

import (
	"github.com/fnlbhq/fred/query"
)

const (
	tags        = "fred/tags"         // Get all tags, search for tags, or get tags by name.
	relatedTags = "fred/related_tags" // Get the related tags for one or more tags.
	series      = "fred/tags/series"  // Get the series matching tags.
)

func Tags() query.Query {
	return query.NewQuery(tags)
}

func RelatedTags() query.Query {
	return query.NewQuery(relatedTags)
}

func Series() query.Query {
	return query.NewQuery(series)
}
