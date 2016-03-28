package extractors

import (
	"golang.org/x/net/html"
	"github.com/yhat/scrape"
	"github.com/napisani/wootscrape/matchers"
	"golang.org/x/net/html/atom"
)
type WootDeal struct{
	Title string
	MinPrice string
	MaxPrice string
	Category string
}

func getInternalWootDealAttrs(attr string, offerNode *html.Node) string{
	attrNode, _ := scrape.Find(offerNode, scrape.ByClass(attr))
	return scrape.Text(attrNode);
}

func GetWootDeal(offerNode *html.Node) WootDeal {
	deal := WootDeal{};
	//label := offerNode.Data
	titleNode, _ := scrape.Find(offerNode, scrape.ByClass("title"))
	title := scrape.Text(titleNode);


	priceNode, _ := scrape.Find(offerNode, scrape.Matcher(matchers.GetWootPriceSpan()))

	maxNode, _ := scrape.Find(priceNode, scrape.ByClass("max"))
	minNode, _ := scrape.Find(priceNode, scrape.ByClass("min"))

	if maxNode == nil || minNode == nil{
		deal.MaxPrice = scrape.Text(priceNode)
		deal.MinPrice = scrape.Text(priceNode)
	}else{
		deal.MaxPrice = scrape.Text(maxNode)
		deal.MinPrice = scrape.Text(minNode)
	}
	parentSectionNode, _ := scrape.FindParent(offerNode, scrape.ByTag(atom.Section))
	parentLiNode, _ := scrape.FindParent(parentSectionNode, scrape.ByTag(atom.Li))
	tabLinkNode, _ := scrape.Find(parentLiNode, scrape.ByClass("tab"))
	categoryNode, _ := scrape.Find(tabLinkNode, scrape.ByTag(atom.Div))
	deal.Category = scrape.Text(categoryNode);
	//title := scrape.Attr(label, "class");
	//title := label.NextSibling;
	//priceDiv := title.NextSibling;

	deal.Title = title;

	return deal;
}