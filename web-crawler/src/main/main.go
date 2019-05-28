package main

/*

Exercise: Web Crawler

In this exercise you'll use Go's concurrency features to parallelize a web crawler.
Modify the Crawl function to fetch URLs in parallel without fetching the same URL twice.
Hint: you can keep a cache of the URLs that have been fetched on a map, but maps alone are not safe for concurrent use!

*/

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // 1

var urls = UniqueUrls{list: make(map[string]bool)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	if urls.list[url] {
		return
	}

	urls.set(url)

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	//ch <- FetchResult{url, body}
	fmt.Printf("found: %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}


	return
}

func main() {

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
