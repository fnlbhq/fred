package query

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/fnlbhq/fred/result"

	"github.com/fnlbhq/fred/query/argument"
)

const fredUrl = "https://api.stlouisfed.org"
const apiKeyEnvVar = "FRED_API_KEY"

func NewQuery(action string) Query {

	u, _ := url.Parse(fmt.Sprintf("%s/%s", fredUrl, action))

	apiKey := os.Getenv(apiKeyEnvVar)

	q := u.Query()
	q.Add("file_type", "json")

	if apiKey != "" {
		q.Add(argument.APIKey, apiKey)
	}

	u.RawQuery = q.Encode()

	return Query{u: u.String()}
}

type Query struct {
	u string
}

func (q Query) AddParameter(argument, value string) Query {
	u, _ := url.Parse(q.u)
	query := u.Query()
	query.Set(argument, fmt.Sprintf("%v", value))
	u.RawQuery = query.Encode()
	q.u = u.String()
	return q
}

func (q Query) APIKey(value string) Query {
	return q.AddParameter(argument.APIKey, value)
}

func (q Query) CategoryID(value string) Query {
	return q.AddParameter(argument.CategoryId, value)
}

func (q Query) RealtimeStart(value string) Query {
	return q.AddParameter(argument.RealTimeStart, value)
}

func (q Query) RealtimeEnd(value string) Query {
	return q.AddParameter(argument.RealTimeEnd, value)
}

func (q Query) Limit(value string) Query {
	return q.AddParameter(argument.Limit, value)
}

func (q Query) Offset(value string) Query {
	return q.AddParameter(argument.Offset, value)
}

func (q Query) OrderBy(value string) Query {
	return q.AddParameter(argument.OrderBy, value)
}

func (q Query) SortOrder(value string) Query {
	return q.AddParameter(argument.SortOrder, value)
}

func (q Query) FilterVariable(value string) Query {
	return q.AddParameter(argument.FilterVariable, value)
}

func (q Query) FilterValue(value string) Query {
	return q.AddParameter(argument.FilterValue, value)
}

func (q Query) TagNames(value string) Query {
	return q.AddParameter(argument.TagNames, value)
}

func (q Query) ExcludeTagNames(value string) Query {
	return q.AddParameter(argument.ExcludeTagNames, value)
}

func (q Query) TagGroupID(value string) Query {
	return q.AddParameter(argument.TagGroupId, value)
}

func (q Query) SearchText(value string) Query {
	return q.AddParameter(argument.SearchText, value)
}

func (q Query) IncludeReleaseDatesWithNoData(value string) Query {
	return q.AddParameter(argument.IncludeReleaseDatesWithNoData, value)
}

func (q Query) ReleaseID(value string) Query {
	return q.AddParameter(argument.ReleaseId, value)
}

func (q Query) ElementID(value string) Query {
	return q.AddParameter(argument.ElementId, value)
}

func (q Query) IncludeObservationValues(value string) Query {
	return q.AddParameter(argument.IncludeObservationValues, value)
}

func (q Query) ObservationDate(value string) Query {
	return q.AddParameter(argument.ObservationDate, value)
}

func (q Query) SeriesID(value string) Query {
	return q.AddParameter(argument.SeriesId, value)
}

func (q Query) ObservationStart(value string) Query {
	return q.AddParameter(argument.ObservationStart, value)
}

func (q Query) ObservationEnd(value string) Query {
	return q.AddParameter(argument.ObservationEnd, value)
}

func (q Query) Units(value string) Query {
	return q.AddParameter(argument.Units, value)
}

func (q Query) Frequency(value string) Query {
	return q.AddParameter(argument.Frequency, value)
}

func (q Query) AggregationMethod(value string) Query {
	return q.AddParameter(argument.AggregationMethod, value)
}

func (q Query) OutputType(value string) Query {
	return q.AddParameter(argument.OutputType, value)
}

func (q Query) VintageDates(value string) Query {
	return q.AddParameter(argument.VintageDates, value)
}

func (q Query) TagSearchText(value string) Query {
	return q.AddParameter(argument.TagSearchText, value)
}

func (q Query) SeriesSearchText(value string) Query {
	return q.AddParameter(argument.SeriesSearchText, value)
}

func (q Query) SourceID(value string) Query {
	return q.AddParameter(argument.SourceId, value)
}

func (q Query) String() string {
	return q.u
}

func (q Query) Get() (*result.Result, error) {
	resp, err := http.Get(q.u)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result result.Result

	err = json.Unmarshal(body, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}
