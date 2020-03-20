package fred

import (
	"fmt"
	"testing"

	"github.com/fnlbhq/fred/series"
)

func TestUpdate(t *testing.T) {
	q, _ := NewQuery(series.Updates)
	r, _ := q.Get()
	fmt.Println(r.PrettyJSON())
}
