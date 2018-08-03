package fetcher

import (
	. "campop/commonutil"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func writeCsv(doc *goquery.Document, keys []string) (values []string) {
	values = make([]string, len(keys))

	doc.Find(".spces-box .table-01 tbody tr").Each(func(i int, tr *goquery.Selection) {
		thText := tr.Find("th").Text()
		title := strings.TrimSpace(thText)
		idx := IndexOf(keys, title)
		if idx != -1 {
			val := tr.Find("td span.js-detail-col2").Text()
			values[idx] = val
		}
	})
	productNameEl := doc.Find("div.mnc h2.t-18-bb.c-red")
	values[1] = strings.Replace(productNameEl.Text(), "综述", "", -1)
	productHref, ok := doc.Find("div.mnc .summary-box .tpc-img a").Attr("href")
	if ok {
		values[0] = strings.Split(productHref, "-")[2]
	}
	return
}
