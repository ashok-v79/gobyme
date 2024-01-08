package main

import (
	"fmt"
)

// Fetcher interface that retrieves data
// type Fetcher interface {
// 	FetchData() ([]byte, error)
// }

// Real implementation of Fetcher
type RealFetcher struct{}

// FetchData fetches data. In real implementation, it might fetch data from an external API
func (f *RealFetcher) FetchData() ([]byte, error) {
	// Simulate fetching data. In a real application, this would be an API call, for example.
	return []byte("real data"), nil
}

// myFunc uses a Fetcher to fetch data
func myFunc(f Fetcher) ([]byte, error) {
	return f.FetchData()
}

func main() {
	fetcher := &RealFetcher{}
	data, err := myFunc(fetcher)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	fmt.Println("Fetched data:", string(data))
}
