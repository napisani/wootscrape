package main

import (
	"fmt"
	"net/http"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"github.com/napisani/wootscrape/matchers"
	"github.com/napisani/wootscrape/extractors"
)

func main() {
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
	matcher := scrape.Matcher(matchers.GetWootOfferNode());
	// grab all articles and print them
	offerNodes := scrape.FindAll(root, matcher)
	for _, offerNode := range offerNodes {
		offer := extractors.GetWootDeal(offerNode);
		if offer.MaxPrice != offer.MinPrice {
			fmt.Printf("%20s) %60s :: %s - %s\n", offer.Category, offer.Title, offer.MinPrice, offer.MaxPrice);
		}else{
			fmt.Printf("%20s) %60s :: %s\n", offer.Category, offer.Title, offer.MinPrice);

		}

	}
}
