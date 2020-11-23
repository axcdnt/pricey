package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type appConfig struct {
	url string
	pattern string
	targetPrice int
	interval time.Duration
	currency string
}

func main() {
	appConfig := newConfig()
	appConfig.validate()

	// the first visit before waiting
	visit(appConfig)

	for range time.NewTicker(appConfig.interval).C {
		// keep visiting from time to time
		visit(appConfig)
	}
}

func newConfig() appConfig {
	// the flags here serve as a configuration to the app
	url := flag.String("url","", "the url to be crawled <required>")
	pattern := flag.String("pattern", "", "the corresponding css rule <required>")
	targetPrice := flag.Int("target", 0, "the desired price <required>")
	interval := flag.Duration("interval", 5 * time.Second, "the time interval for crawling")
	currency := flag.String("currency", "", "the currency (US$ (US), R$ (BR)...")
	flag.Parse()

	return appConfig {
		*url,
		*pattern,
		*targetPrice,
		*interval,
		*currency,
	}
}

func (a *appConfig) validate() {
	if len(a.url) == 0 {
		log.Fatalln("url is required!")
	}

	if len(a.pattern) == 0 {
		log.Fatalln("search pattern is required!")
	}

	if a.targetPrice <= 0 {
		log.Fatalln("targetPrice is required!")
	}
}

func visit(appConfig appConfig) {
	c := colly.NewCollector()
	c.OnHTML(appConfig.pattern, func(e *colly.HTMLElement) {
		symbols := []string{appConfig.currency, ",", "."}
		price, err := stripSymbols(e.Text, symbols)
		if err != nil {
			log.Printf("could not convert price: %v", err)
		} else {
			fmt.Printf("debug => expected: %d, current: %d\n", appConfig.targetPrice, price)
			if appConfig.targetPrice > 0.0 && price <= appConfig.targetPrice {
				log.Printf("wow! the price (%d) looks good now!\n", price)
			} else {
				log.Printf("the current price is (%d), I'll keep looking...\n", price)
			}
		}
	})

	err := c.Visit(appConfig.url)
	if err != nil {
		log.Fatalf("an error occurred while visiting url: %v", err)
	}
}

func stripSymbols(str string, symbols []string) (int, error) {
	if len(symbols) == 0 {
		return strconv.Atoi(strings.Trim(str, " "))
	}

	// an alternative impl without recursion could be done using regex:
	// var re = regexp.MustCompile(`(US\$|\.|\,)`)
	// s := re.ReplaceAllString(str, "")
	return stripSymbols(strings.Replace(str, symbols[0], "", -1), symbols[1:])
}

