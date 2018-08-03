package main

import (
	"campop/es"
	"campop/fetcher"
	"flag"
	"fmt"
	"os"
	"runtime"
)

func main() {
	flag.NewFlagSet("setup", flag.ExitOnError)
	flag.NewFlagSet("crawl", flag.ExitOnError)
	flag.NewFlagSet("es", flag.ExitOnError)
	flag.NewFlagSet("help", flag.ExitOnError)
	flag.NewFlagSet("version", flag.ExitOnError)
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)
	searchTypePtr := searchCmd.String("t", "camera", "Type <camera|lens>. (Required)")
	keywordPtr := searchCmd.String("q", "", "Query keywords. (Required)")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, helpText)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "setup":
		// 降低并发防止 read: connection reset by peer
		runtime.GOMAXPROCS(1)
		fetcher.Run()
		es.IndicesRemapping()
		es.ReadCamerasCsv()
		es.ReadLensesCsv()
	case "crawl":
		runtime.GOMAXPROCS(1)
		fetcher.Run()
	case "es":
		es.IndicesRemapping()
		es.ReadCamerasCsv()
		es.ReadLensesCsv()
	case "search":
		searchCmd.Parse(os.Args[2:])
	case "help":
		fmt.Println("document upcoming")
	case "version":
		fmt.Println("version: 0.0.1")
	default:
		fmt.Fprintf(os.Stderr, helpText)
		os.Exit(1)
	}

	if searchCmd.Parsed() {
		if *keywordPtr == "" {
			fmt.Printf("need some keyword ヾ(´･ ･｀｡)ノ\n\n")
			searchCmd.PrintDefaults()
			os.Exit(1)
		}

		es.ProcessSearch(*searchTypePtr, *keywordPtr)
		os.Exit(1)

	}

}

var helpText = `Campop is a simple tool for displaying photographic equipment.

Usage:
  campop [command]

Available Commands:
  setup       prepare everything you need (ok, just some cameras & lenses data)
  crawl       crawl the data to a CSV file
  es          manipulate the elasticsearch indices
  search      do some fuzzy search
  help        help about any command
  version     prints version.
` + "\n"
