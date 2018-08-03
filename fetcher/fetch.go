package fetcher

import (
	. "campop/commonutil"
	"encoding/csv"
	"fmt"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const WujiHost = "http://production.xitek.com/"

func Run() {
	start := time.Now()
	fetchAndWriteCameras()
	elapsed := time.Since(start)
	fmt.Printf("\nCameras done took %s\n", elapsed)

	start = time.Now()
	fetchAndWriteLenses()
	elapsed = time.Since(start)
	fmt.Printf("\nLenses done took %s\n", elapsed)
}

func fetchAndWriteCameras() {
	fetchAndWrite(CameraColumnsMapKeys(), getAllCameraLinks(), "cameras.csv")
	return
}

func fetchAndWriteLenses() {
	fetchAndWrite(LensColumnsMapKeys(), getAllLensLinks(), "lenses.csv")
	return
}

func fetchAndWrite(keys, productLinks []string, csvName string) {

	productDocs := make(chan *goquery.Document)
	linksParse := make(chan bool)

	file, err := os.Create(csvName)
	CheckError("Cannot create file", err)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, url := range productLinks {
		go func(url string) {
			defer func() {
				linksParse <- true
			}()
			productDoc, err := getGbkDoc(url)
			if err != nil {
				return
			}
			productDocs <- productDoc
		}(url)
	}

	// fmt.Println("cpus:", runtime.NumCPU())
	// fmt.Println("archive:", runtime.NumGoroutine())
	fmt.Printf("\nWriting CSV.")

	for c := 0; c < len(productLinks); {
		select {
		case doc := <-productDocs:
			values := writeCsv(doc, keys)
			err = writer.Write(values)
			CheckError("Cannot write to file", err)
		case <-linksParse:
			if c%100 == 0 {
				fmt.Printf(".")
			}
			c++
		}
	}
}
