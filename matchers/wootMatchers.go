package matchers

import (
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	//"fmt"
	"strings"
)

type matcher func(n *html.Node)  bool

func GetWootOfferNode () matcher {
	return func (n *html.Node)  bool {
		if n.DataAtom == atom.Div {
			//fmt.Println(scrape.Attr(n, "class"));
			return strings.TrimSpace(scrape.Attr(n, "class")) == "offer"
		}
		return false
	}
}


func GetWootPriceSpan() matcher {
	return func (node *html.Node) bool {
		if(node.DataAtom == atom.Span ){
			return strings.TrimSpace(scrape.Attr(node, "class")) == "price"
		}
		return false;
	}
}

