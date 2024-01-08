package main

// Fetcher interface that retrieves data
type Fetcher interface {
	FetchData() ([]byte, error)
}
