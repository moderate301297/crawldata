package main

import (
	"log"
	"strings"

	"./dbconnection"
	"./savedata"
	"github.com/buger/jsonparser"
)

type DataWeb struct {
	Body string
}

func QueryAll() (value []string) {
	c := dbconnection.Session.DB("data_walmart").C("data")
	var results []DataWeb
	err1 := c.Find(nil).All(&results)
	if err1 != nil {
		log.Println("ERROR : ", err1)
	}
	for _, body := range results {
		value = append(value, body.Body)
	}
	return value
}

func OptimizeUrl(value string) (url string) {
	if strings.Index(value, "http") == 0 {
		url = value
		return url
	} else {
		url = "https://www.walmart.com" + value
		return url
	}
}

func main() {
	value := QueryAll()
	for _, body := range value {
		data := []byte(body)

		dataProduct, _, _, _ := jsonparser.Get(data, "preso")

		dataPath, _, _, _ := jsonparser.Get(dataProduct, "adContext", "categoryPathName")
		jsonparser.ArrayEach(dataProduct, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			title, _, _, _ := jsonparser.Get(value, "title")
			description, _, _, _ := jsonparser.Get(value, "description")
			imageUrl, _, _, _ := jsonparser.Get(value, "imageUrl")
			url, _, _, _ := jsonparser.Get(value, "productPageUrl")
			link := OptimizeUrl(string(url))
			savedata.SaveData(string(title), string(description), string(dataPath), string(imageUrl), link)
		}, "items")
	}
}
