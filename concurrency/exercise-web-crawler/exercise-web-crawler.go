package main

import (
	"errors"
	"fmt"
	"maps"
	"slices"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type URLCache struct {
	mu   sync.RWMutex
	urls map[string]error
}

func NewURLCache() *URLCache {
	return &URLCache{urls: make(map[string]error)}
}

var errLoading = errors.New("url load in progress") // sentinel value

// TryVisit returns true if the URL hasn't been visited yet, otherwise marks it as loading.
func (c *URLCache) TryVisit(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.urls[url]; ok {
		return false
	}
	c.urls[url] = errLoading
	return true
}

func (c *URLCache) Update(url string, err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.urls[url] = err
}

func (c *URLCache) PrintStats() {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Println("Fetching stats\n--------------")
	for _, url := range slices.Sorted(maps.Keys(c.urls)) {
		err := c.urls[url]
		if err != nil {
			fmt.Printf("%v failed: %v\n", url, err)
		} else {
			fmt.Printf("%v was fetched\n", url)
		}
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *URLCache, wg *sync.WaitGroup) {
	if depth <= 0 {
		fmt.Printf("<- Done with %v, depth 0.\n", url)
		return
	}

	if !cache.TryVisit(url) {
		return
	}

	fmt.Printf("-> Fetching: %v\n", url)
	body, urls, err := fetcher.Fetch(url)
	cache.Update(url, err)

	if err != nil {
		fmt.Printf("<- Error fetching %v: %v\n", url, err)
		return
	}
	fmt.Printf("<- Found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Go(func() {
			Crawl(u, depth-1, fetcher, cache, wg)
		})
	}
}

func main() {
	fetched := NewURLCache()
	wg := &sync.WaitGroup{}
	wg.Go(func() {
		Crawl("https://golang.org/", 4, fetcher, fetched, wg)
	})
	wg.Wait()
	fetched.PrintStats()

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
