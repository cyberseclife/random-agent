package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Agent struct to categorize our user agents
type Agent struct {
	UserAgent string
	Type      string // "mobile" or "desktop"
	OS        string // "windows", "linux", "mac", "android", "ios"
}

var agents = []Agent{
	// --- DESKTOP ---
	// Windows
	{UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36", Type: "desktop", OS: "windows"},
	{UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Edg/119.0.0.0", Type: "desktop", OS: "windows"},
	{UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:120.0) Gecko/20100101 Firefox/120.0", Type: "desktop", OS: "windows"},
	// Mac
	{UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.1 Safari/605.1.15", Type: "desktop", OS: "mac"},
	{UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36", Type: "desktop", OS: "mac"},
	{UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/119.0", Type: "desktop", OS: "mac"},
	// Linux
	{UserAgent: "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36", Type: "desktop", OS: "linux"},
	{UserAgent: "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/119.0", Type: "desktop", OS: "linux"},

	// --- MOBILE ---
	// iOS
	{UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 17_1_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1", Type: "mobile", OS: "ios"},
	{UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 16_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.6 Mobile/15E148 Safari/604.1", Type: "mobile", OS: "ios"},
	{UserAgent: "Mozilla/5.0 (iPad; CPU OS 17_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.0 Mobile/15E148 Safari/604.1", Type: "mobile", OS: "ios"},
	// Android
	{UserAgent: "Mozilla/5.0 (Linux; Android 14; SM-S918B) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.6099.43 Mobile Safari/537.36", Type: "mobile", OS: "android"},
	{UserAgent: "Mozilla/5.0 (Linux; Android 13; Pixel 7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Mobile Safari/537.36", Type: "mobile", OS: "android"},
	{UserAgent: "Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Mobile Safari/537.36", Type: "mobile", OS: "android"},
	{UserAgent: "Mozilla/5.0 (Android 14; Mobile; rv:120.0) Gecko/120.0 Firefox/120.0", Type: "mobile", OS: "android"},
}

func main() {
	// Flags
	mobileFlag := flag.Bool("mobile", false, "Generate only mobile User-Agents")
	desktopFlag := flag.Bool("desktop", false, "Generate only desktop User-Agents")
	osFlag := flag.String("os", "", "Filter by specific OS (windows, linux, mac, android, ios)")
	
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `
Random-Agent: A simple utility to generate random User-Agent strings.

Usage: random-agent [flags]

Description:
  If no flags are provided, the program generates a random User-Agent 
  selected from the entire available pool (Desktop & Mobile, all OSs).

Flags:
`)
		flag.PrintDefaults()
	}
	flag.Parse()

	// Logic
	var candidates []string

	for _, a := range agents {
		// Filter by Type (-mobile / -desktop)
		if *mobileFlag && a.Type != "mobile" {
			continue
		}
		if *desktopFlag && a.Type != "desktop" {
			continue
		}

		// Filter by OS (-os)
		if *osFlag != "" {
			if strings.ToLower(*osFlag) != a.OS {
				continue
			}
		}

		candidates = append(candidates, a.UserAgent)
	}

	if len(candidates) == 0 {
		fmt.Println("No User-Agents found matching your criteria.")
		os.Exit(1)
	}

	// Random Selection
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Println(candidates[r.Intn(len(candidates))])
}