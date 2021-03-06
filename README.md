## FRED  
#### A client for the St. Louis Fed's FRED [Economic Data API](https://research.stlouisfed.org/docs/api/fred/), written in [Go](https://golang.org).  
FRED aims to provide a simple, developer-friendly interface that largely mirrors the patterns established by the  [API](https://research.stlouisfed.org/docs/api/fred/).
### Example
```go

// Fetch the latest 1000 series updates, offset by 2000
q := series.Updates().Limit("1000").Offset("2000")

// View the query url
// url := q.String() 

// Execute the request
r, _ := q.Get()

// Convert the fred to JSON and print the fred
fmt.Println(r.PrettyJSON())


// Fetch the US unemployment rate via the series observation method
result, err := GetSeriesObservations("UNRATE")

// Build a query to fetch the US unemployment rate
result, err := series.Observations().AddParameter(argument.SeriesId, "UNRATE").Get()
``` 
