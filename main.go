package fred

import (
	"fmt"
	"net/url"
	"os"

	"github.com/fnlbhq/fred/query"

	"github.com/fnlbhq/fred/argument"
)

const fredUrl = "https://api.stlouisfed.org"
const apiKeyEnvVar = "FRED_API_KEY"

func NewQuery(action string) (*query.Query, error) {

	u, err := url.Parse(fmt.Sprintf("%s/%s", fredUrl, action))

	if err != nil {
		return nil, err
	}

	apiKey := os.Getenv(apiKeyEnvVar)

	q := u.Query()
	q.Add("file_type", "json")

	if apiKey != "" {
		q.Add(argument.APIKey, apiKey)
	}

	u.RawQuery = q.Encode()

	return &query.Query{URL: *u}, nil
}
