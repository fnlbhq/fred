package fred

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/gocarina/gocsv"

	"github.com/fnlbhq/fred/fred"

	"github.com/fnlbhq/fred/series"
)

func init() {
	// os.Setenv("FRED_API_KEY", "")
}

func TestUpdatesToCSV(t *testing.T) {

	limit := 1000
	offset := 0
	var count int
	var accum []fred.Series
	q, _ := series.Updates().Limit(strconv.Itoa(limit)).Offset(strconv.Itoa(offset)).Get()
	count = q.Count

	fmt.Printf("There will be %d entries\n", count)

	accum = append(accum, q.Series...)

	for offset+limit < count {
		offset += limit

		fmt.Println(offset)
		q, err := series.Updates().Limit(strconv.Itoa(limit)).Offset(strconv.Itoa(offset)).Get()

		if err != nil {
			panic(err)
		}

		accum = append(accum, q.Series...)
	}

	updatesFile, err := os.Create("fred_updates.csv")

	if err != nil {
		panic(err)
	}

	defer updatesFile.Close()

	if err := gocsv.MarshalFile(&accum, updatesFile); err != nil {
		panic(err)
	}

	fmt.Println(len(accum))
}
