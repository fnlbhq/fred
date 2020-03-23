package fred

import (
	"fmt"
	"testing"

	"github.com/fnlbhq/fred/series"
)

func init() {
	// os.Setenv("FRED_API_KEY", "")
}

func TestUpdate(t *testing.T) {

	//limit := 1000
	// offset := 0

	q := series.Updates().Limit("100").Offset("100000")

	fmt.Println(q.String())
	// r, _ := q.Get()

	// fmt.Println(r.PrettyJSON())
}
