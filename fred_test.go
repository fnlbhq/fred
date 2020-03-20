package fred

import (
	"fmt"
	"testing"
)

/*
func TestMain(m *testing.M) {
	q, _ := NewQuery(series.Categories)

	q.With(argument.AggregationMethod, "ABCDEFG").String()
	fmt.Println(q.String())

}

func TestFoo(t *testing.T) {
	q, err := NewQuery(categories.Category)
	fmt.Println(err)
	result, err := q.With(argument.CategoryId, "125").Get()
	fmt.Println(err)
	fmt.Println(result)

	fmt.Println(q.With(argument.CategoryId, "125").String())
}
*/

func TestUpdate(t *testing.T) {
	result, err := Updates()
	fmt.Println(err)
	fmt.Println(result.PrettyJSON())
}
