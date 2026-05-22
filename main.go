// go-http-ping: sends HTTP requests and reports latency
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"time"
)

func ping(url string, n int) {
	fmt.Printf("\n🏓 Pinging %s (%d times)\n", url, n)
	var total time.Duration
	ok := 0
	client := &http.Client{Timeout: 10 * time.Second}
	for i := 0; i < n; i++ {
		start := time.Now()
		resp, err := client.Get(url)
		dur := time.Since(start)
		if err != nil {
			fmt.Printf("  #%d ❌ %v\n", i+1, err)
		} else {
			resp.Body.Close()
			total += dur
			ok++
			emoji := "✅"
			if resp.StatusCode >= 400 { emoji = "⚠️ " }
			fmt.Printf("  #%d %s %d  %v\n", i+1, emoji, resp.StatusCode, dur.Round(time.Millisecond))
		}
		time.Sleep(200 * time.Millisecond)
	}
	if ok > 0 {
		fmt.Printf("  avg: %v | success: %d/%d\n", (total/time.Duration(ok)).Round(time.Millisecond), ok, n)
	}
}

func main() {
	n := flag.Int("n", 3, "number of pings")
	flag.Parse()
	if flag.NArg() == 0 { fmt.Fprintln(os.Stderr, "Usage: http-ping [-n N] <url>..."); os.Exit(1) }
	for _, url := range flag.Args() { ping(url, *n) }
}
