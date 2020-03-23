package fred

import (
	"os"
)

func init() {
	os.Setenv("FRED_API_KEY", "")
}

/*
func TestUpdate(t *testing.T) {

	//limit := 1000
	// offset := 0

	q, _ := NewQuery(series.Updates)

	q.Argument(argument.Limit, "100")
	q.Argument(argument.Offset, "100000")
	fmt.Println(q.String())
	r, _ := q.Get()

	fmt.Println(r.PrettyJSON())
}
*/
