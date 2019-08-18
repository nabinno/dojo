package main

import (
	"errors"
	"fmt"
	"sync"
)

type fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

var fetched = struct {
	m map[string]error
	sync.Mutex
}{m: make(map[string]error)}

var errLoading = errors.New("url load in progress") // sentinel value

// Crawl ...
func Crawl(url string, depth int, f fetcher) ([]string, error) {
	if depth <= 0 {
		return []string{url}, fmt.Errorf("depth zero: %s", url)
	}

	fetched.Lock()
	if _, ok := fetched.m[url]; ok {
		fetched.Unlock()
		return []string{url}, fmt.Errorf("already fetched: %s", url)
	}
	fetched.m[url] = errLoading
	fetched.Unlock()

	_, urls, err := f.Fetch(url)
	fetched.Lock()
	fetched.m[url] = err
	fetched.Unlock()
	if err != nil {
		return []string{url}, err
	}

	rc := []string{url}
	for i := 0; i < len(urls); i++ {
		u := urls[i]
		if childURLs, childErr := Crawl(u, depth-1, f); childErr == nil {
			for _, childURL := range childURLs {
				rc = append(rc, childURL)
			}
		}
	}
	return rc, nil
}
