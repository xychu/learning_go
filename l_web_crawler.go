package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 {
		return
	}

	fetched := make(map[string]bool)
	type target struct {
		urls  []string
		depth int
	}

	doFetch := func(url string, ch chan target) {
		body, urls, err := fetcher.Fetch(url)
		fetched[url] = true
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("found: %s %q\n", url, body)
		}
		ch <- target{urls, depth - 1}
	}

	url_ch := make(chan target)

	go doFetch(url, url_ch)

	for count := 1; count > 0; {
		next := <-url_ch
		count--
		if next.depth <= 0 {
			return
		}
		for _, u := range next.urls {
			if _, fetched := fetched[u]; fetched {
				continue
			}
			fetched[u] = true
			count++
			go doFetch(u, url_ch)
		}
	}

	return
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
