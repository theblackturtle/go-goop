package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	var concurrency int
	flag.IntVar(&concurrency, "c", 1, "Concurrency level")

	var delay int
	flag.IntVar(&delay, "d", 1, "Delay between requests")

	var numPage int
	flag.IntVar(&numPage, "p", 2, "Number of page for searching")

	var getFull bool
	flag.BoolVar(&getFull, "f", false, "Get unparsed results")
	flag.Parse()

	FBCookie := os.Getenv("FBCookie")
	if FBCookie == "" {
		fmt.Fprintln(os.Stderr, "Please set your Facebook Cookie")
		os.Exit(0)
	}

	jobs := make(chan string)
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		c := getClient()
		go func() {
			defer wg.Done()
			for dork := range jobs {
				for page := 0; page <= numPage; page++ {
					urls := search(c, dork, FBCookie, page, getFull)

					if len(urls) == 0 {
						break
					}
					for _, url := range urls {
						fmt.Fprintln(os.Stderr, url)
					}
					time.Sleep(time.Duration(delay) * time.Second)
				}
			}
		}()
	}
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		jobs <- sc.Text()
	}
	close(jobs)
	wg.Wait()
}
