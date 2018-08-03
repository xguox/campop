package es

import (
	. "campop/commonutil"
	"campop/model"
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/olivere/elastic"
)

var Client *elastic.Client

const host = "http://127.0.0.1:9200/"

func init() {
	errorlog := log.New(os.Stdout, "APP", log.LstdFlags)
	var err error
	Client, err = elastic.NewClient(elastic.SetErrorLog(errorlog), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	// info, code, err := Client.Ping(host).Do(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)
	// esversion, err := Client.ElasticsearchVersion(host)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Elasticsearch version %s\n", esversion)
}

func CreateMappings() {
	createIndexAndMapping(CameraIndexName, CameraMapping)
	createIndexAndMapping(LensIndexName, LensMapping)
}

func IndicesRemapping() {
	indexRemapping(CameraIndexName, CameraMapping)
	indexRemapping(LensIndexName, LensMapping)
}

func IndicesDelete() {
	deleteIndex(CameraIndexName)
	deleteIndex(LensIndexName)
}

func chechIndexExist(indexName string) (bool, error) {
	exists, err := Client.IndexExists(indexName).Do(context.Background())
	if err != nil {
		log.Println("chechIndexExist " + err.Error())
		return false, err
	}
	return exists, err
}

// create Mapping
func createIndexAndMapping(indexName, mapping string) error {
	// Use the IndexExists service to check if a specified index exists.
	ctx := context.Background()
	exists, err := chechIndexExist(indexName)

	if !exists {
		log.Println("es index not exists: " + indexName)
		log.Println("creating...")
		// Create a new index.
		createIndex, err := Client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
		if err != nil {
			log.Println("CreateIndex " + err.Error())
			return err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
			return errors.New("create index:" + indexName + ", not Ack nowledged")
		}

	}

	return err
}

func deleteIndex(indexName string) {
	exists, err := chechIndexExist(indexName)
	if err != nil {
		panic(err)
	}

	if exists {

		deleteIndex, err := Client.DeleteIndex(indexName).Do(context.Background())
		if err != nil {
			panic(err)
		}
		if !deleteIndex.Acknowledged {
			// Not acknowledged
		}
	}
}

func indexRemapping(indexName, mapping string) {
	deleteIndex(indexName)
	createIndexAndMapping(indexName, mapping)
}

func putMapping(indexName, typeName, mapping string) {
	exists, err := Client.TypeExists().Index(indexName).Type(typeName).Do(context.Background())
	if err != nil {
		panic(err)
	}
	if !exists {
		putResp, err := Client.PutMapping().Index(indexName).Type(typeName).BodyString(mapping).Do(context.TODO())

		if err != nil {
			panic(err)
		}
		fmt.Println(putResp)
	}
}

func insertDoc(indexName, typeName, idString string, data interface{}) {
	// 插入数据 (using JSON serialization)
	_, err := Client.Index().
		Index(indexName).
		Type(typeName).
		Id(idString).
		BodyJson(data).
		Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
}

func ProcessSearch(indexName, q string) {
	if indexName == "camera" {
		results, err := search(CameraIndexName, CameraTypeName, q)
		CheckError("", err)
		if results.Hits.TotalHits == 0 {
			fmt.Println("No result found.")
			return
		}

		var typ model.Camera

		for _, item := range results.Each(reflect.TypeOf(typ)) {
			fmt.Println("***************************************************")

			camera := item.(model.Camera)
			s := reflect.ValueOf(&camera).Elem()
			for i := 1; i < s.NumField(); i++ {
				f := s.Field(i)
				if f.Interface() != "" {
					name := CameraColumnsMapConstEnFirst[s.Type().Field(i).Name]
					nameLen := len([]rune(name))
					placeholder := strings.Repeat(" ", 20-nameLen*2)
					fmt.Printf(" %s%s|     %v\n", name, placeholder, f.Interface())
				}
			}
		}
	} else if indexName == "lens" {
		results, err := search(LensIndexName, LensTypeName, q)
		CheckError("", err)
		if results.Hits.TotalHits == 0 {
			fmt.Println("No result found.")
			return
		}
		var typ model.Lens
		for _, item := range results.Each(reflect.TypeOf(typ)) {
			fmt.Println("***************************************************")
			lens := item.(model.Lens)
			s := reflect.ValueOf(&lens).Elem()
			for i := 1; i < s.NumField(); i++ {
				f := s.Field(i)
				if f.Interface() != "" {
					name := LensColumnsMapConstEnFirst[s.Type().Field(i).Name]
					nameLen := len([]rune(name))
					placeholder := strings.Repeat(" ", 20-nameLen*2)
					fmt.Printf(" %s%s|     %v\n", name, placeholder, f.Interface())
				}
			}
		}
	}
}

func search(indexName, typeName, q string) (*elastic.SearchResult, error) {
	query := elastic.NewMultiMatchQuery(q, "name", "lens_mount", "sensor_type", "maximum_aperture", "focal_length", "battery", "annex").Type("best_fields")
	searchResult, err := Client.Search().
		Index(indexName).        // search in index "twitter"
		Query(query).            // specify the query
		From(0).Size(10).        // take documents 0-9
		Pretty(true).            // pretty print request and response JSON
		Do(context.Background()) // execute
	if err != nil {
		// Handle error
		panic(err)
		return nil, err
	}
	return searchResult, err
}
