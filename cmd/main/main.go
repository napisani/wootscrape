
package main

import (
	"fmt"
	"net/http"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"github.com/napisani/wootscrape/matchers"
	"github.com/napisani/wootscrape/extractors"
	"flag"
	"time"
)

func main() {
	watchFrequency := flag.Int("w", 60 * 5, "the watch frequency")
	flag.Parse();

	fmt.Println(*watchFrequency);

	for ;; {
		c := make(chan string)
		go getWootDeals(c);
		fmt.Print("\033[H\033[2J");
		for strDeal := range c {
			fmt.Print(strDeal);
		}
		time.Sleep(time.Duration(*watchFrequency)*time.Second);
	}
}

func getWootDeals(c chan string){

	// request and parse the front page
	resp, err := http.Get("http://woot.com")
	if err != nil {
		panic(err)
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		panic(err)
	}

	// define a matcher
	matcher := scrape.Matcher(matchers.GetWootOfferNode())
	// grab all articles and print them
	offerNodes := scrape.FindAll(root, matcher)
	for _, offerNode := range offerNodes {
		offer := extractors.GetWootDeal(offerNode)
		if offer.MaxPrice != offer.MinPrice {
			c <- fmt.Sprintf("%20s) %60s :: %s - %s\n", offer.Category, offer.Title, offer.MinPrice, offer.MaxPrice)
		}else{
			c <- fmt.Sprintf("%20s) %60s :: %s\n", offer.Category, offer.Title, offer.MinPrice)

		}

	}
	close(c)

}