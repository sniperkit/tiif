package engine

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

type Google struct {
}

func (g *Google) Flag() (flag, usage string) {
	flag, usage = "g", "Google"
	return
}

func (g *Google) URL(q string, domains []string, limit int) string {
	var dms string
	for i, dm := range domains {
		if i > 0 {
			dms += " OR"
		}
		dms += " site:" + dm
	}

	return "https://www.google.com/search?q=" + url.QueryEscape(q+dms)
}

func (g *Google) Parse(doc *goquery.Document, limit int) ([]*Result, error) {
	results := []*Result{}
	rNodes := doc.Find("#ires ol .g") // h3.r a")

	rNodes.Each(func(i int, n *goquery.Selection) {
		a := n.Find("h3.r a")
		rawURL, _ := a.Attr("href")
		googleURL, err := url.Parse(rawURL)
		if err != nil {
			return
		}

		u, ok := googleURL.Query()["q"]
		if !ok {
			return
		}

		results = append(results, &Result{
			singleLine(a.Text()),
			singleLine(n.Find(".st").Text()),
			u[0],
		})
	})

	return results, nil
}
