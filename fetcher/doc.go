package fetcher

import (
	. "campop/commonutil"
	"fmt"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

var netClient = &http.Client{}

func init() {
	t := &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConnsPerHost: 10,
		MaxIdleConns:        10,
		IdleConnTimeout:     30 * time.Second,
		TLSHandshakeTimeout: 10 * time.Second,
	}
	netClient = &http.Client{Transport: t}

}

func getAllCameraLinks() (links []string) {
	links = getAllProductLinks("http://production.xitek.com/post.php?a=list&parent=2&cid=0&bid=0&show=1&order=title&sort=ASC&s_p_a=&page=", cameraPagesCount())
	return
}

func getAllLensLinks() (links []string) {
	links = getAllProductLinks("http://production.xitek.com/post.php?a=list&parent=9&cid=0&bid=0&show=1&order=title&sort=ASC&s_p_a=&page=", lensPagesCount())
	return
}

func getAllProductLinks(base_url, pagesCount string) (links []string) {
	// Channels
	productLinks := make(chan string)
	pageFecthed := make(chan bool)

	maxPage, err := strconv.Atoi(pagesCount)
	CheckError("", err)
	fmt.Printf("Fetching.")

	for i := 1; i <= maxPage; i++ {
		url := base_url + fmt.Sprintf("%d", i)
		go getSinglePageProductLinks(url, productLinks, pageFecthed)
	}

	for parsedPage := 1; parsedPage < maxPage; {
		select {
		case url := <-productLinks:
			links = append(links, url)
		case <-pageFecthed:
			fmt.Printf(".")
			parsedPage++
		}
	}
	return
}

func getSinglePageProductLinks(url string, urlsCh chan string, pageFecthed chan bool) {
	doc := getDoc(url)

	defer func() {
		pageFecthed <- true
	}()

	doc.Find("#list-box .t-name a").Each(func(i int, a *goquery.Selection) {
		href, ok := a.Attr("href")
		if ok {
			href = WujiHost + strings.Replace(href, "post", "basic", -1)
			urlsCh <- href
		}
	})
	return
}

func getGbkDoc(url string) (doc *goquery.Document, err error) {
	resp, err := netClient.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Println("status code error: %d %s", resp.StatusCode, resp.Status)
		return
	}

	reader := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	doc, err = goquery.NewDocumentFromReader(reader)
	CheckError("", err)
	return
}

func getDoc(url string) (doc *goquery.Document) {
	resp, err := netClient.Get(url)
	CheckError("", err)

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatal("status code error: %d %s", resp.StatusCode, resp.Status)
	}

	doc, err = goquery.NewDocumentFromReader(resp.Body)
	CheckError("", err)
	return
}

func cameraPagesCount() string {
	return getTotalPagesCount("http://production.xitek.com/post-a-list-parent-2.html")
}

func lensPagesCount() string {
	return getTotalPagesCount("http://production.xitek.com/post-a-list-parent-9.html")
}

func getTotalPagesCount(url string) string {
	doc := getDoc(url)
	lastPageUrl, ok := doc.Find("#list-box .page-last").Last().Attr("href")
	if ok {
		arr := strings.Split(lastPageUrl, "=")
		maxPage := arr[len(arr)-1]
		return maxPage
	} else {
		return "1"
	}
}

func getCategoryLinks() (links []string) {
	url := "http://production.xitek.com/index.html"
	doc := getDoc(url)
	doc.Find("#comment .c-red").Each(func(i int, a *goquery.Selection) {
		// For each item found, get the band and title
		href, ok := a.Attr("href")
		if ok {
			links = append(links, WujiHost+href)
		}
	})
	return
}
