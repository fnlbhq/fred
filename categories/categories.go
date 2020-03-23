package categories

import "github.com/fnlbhq/fred/query"

const (
	category    = "fred/category"              // Get a category.
	children    = "fred/category/children"     // Get the child categories for a specified parent category.
	related     = "fred/category/related"      // Get the related categories for a category.
	series      = "fred/category/series"       // Get the series in a category.
	tags        = "fred/category/tags"         // Get the tags for a category.
	relatedTags = "fred/category/related_tags" // Get the related tags for a category.
)

func Category() *query.Query {
	return query.NewQuery(category)
}

func Children() *query.Query {
	return query.NewQuery(children)
}

func Related() *query.Query {
	return query.NewQuery(related)
}

func Series() *query.Query {
	return query.NewQuery(series)
}

func Tags() *query.Query {
	return query.NewQuery(tags)
}

func RelatedTags() *query.Query {
	return query.NewQuery(relatedTags)
}
