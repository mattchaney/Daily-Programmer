package main
 
import (
"fmt"
"time"
)
 
type Fetcher interface {
// Fetch returns the body of URL and
// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
 
type Todo struct {
	url string
	depth int
}
 
func getPage(url string, depth int, todolist chan Todo, live chan int) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		live <- -1
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		todolist <- Todo{u, depth + 1}
	}
	live <- -1
}
 
func Crawl(url string, depth int, fetcher Fetcher) {
	todolist := make(chan Todo, 1)
	live := make(chan int)
	livecount := 0
	visited := make(map[string]bool)
	 
	todolist <- Todo{url, 0}
	 
	for {
		select {
			case todo := <-todolist:
				if todo.depth <= depth && !visited[todo.url] {
					livecount++
					visited[todo.url] = true
					go getPage(todo.url, todo.depth, todolist, live)
				}
			case c := <-live:
				livecount += c
			default:
				time.Sleep(1)
		}
		if livecount == 0 {
			break
		}
	}
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