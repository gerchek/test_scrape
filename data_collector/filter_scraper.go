package data_collector

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"scraper_trendyol/models/trendyol"
	consts "scraper_trendyol/pkg/const"
	"scraper_trendyol/pkg/helper"

	"github.com/gocolly/colly/v2"
)

/*
	Collects the filter data of trendyol.com, displayed on the left side of the site.
*/
type TYFilterDataScraper struct {
	Collector *colly.Collector
}

func NewTYScraper() TYFilterDataScraper {
	return TYFilterDataScraper{Collector: colly.NewCollector()}
}

// ScrapeCategory Scrape the categories
func (scraper *TYFilterDataScraper) ScrapeCategory() trendyol.SortedItems {

	// SortedItems: Sorts the filtered data. prevents duplication of the data.
	sortedItem := trendyol.SortedItems{}

	scraper.Collector.OnHTML(".main-nav > .tab-link", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(index int, catHeaderElm *colly.HTMLElement) {
			link := catHeaderElm.Attr("href")

			categoryType := catHeaderElm.Attr("class")

			if categoryType == "category-header" {
				//fmt.Println("category-header")
			} else if categoryType == "sub-category-header" {
				fmt.Printf("Link found: %d: %q -> %s\n", index, catHeaderElm.Text, link)

				// fetch filtered data,
				fetchFilterData(link, &sortedItem)
			} else {
				fmt.Println("3rd Level category: ", link)
				//filterCollector.Visit(e.Request.AbsoluteURL(link))
			}
		})
	})

	scraper.Collector.OnScraped(func(response *colly.Response) {
		fmt.Printf("data_collector  completed")
		fmt.Printf("Brand count: %d ", len(sortedItem.Brands))
		fmt.Printf("Category count: %d ", len(sortedItem.Categories))
	})

	scraper.Collector.Visit("https://www.trendyol.com/")

	return sortedItem
}

func fetchFilterData(link string, sortedItems *trendyol.SortedItems) {
	baseLink := "https://public.trendyol.com/discovery-web-searchgw-service/v2/api/aggregations/"
	params := "?culture=tr-TR&storefrontId=1&categoryRelevancyEnabled=false&priceAggregationType=DYNAMIC_GAUSS&searchTestTypeAbValue=B"

	queryLink := baseLink + link + params
	resp, err := http.Get(queryLink)

	if err != nil {
		err := fmt.Errorf("error http get: %e", err)
		fmt.Println(err)
	}

	defer resp.Body.Close()

	var respData trendyol.TrendyolFilterResponse
	byteData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteData, &respData)
	if err != nil {
		err := fmt.Errorf("error convert to byte: %e", err)
		fmt.Println(err)
	}

	//fmt.Println("DATA: ", respData)

	aggregations := respData.Result.Aggregations

	for _, aggregation := range aggregations {

		// Get brands
		if aggregation.Type == consts.BRAND_FK {
			for _, value := range aggregation.Values {
				item := trendyol.Item{
					ID:             value.ID,
					Text:           value.Text,
					BeautifiedName: value.BeautifiedName,
					Type:           value.Type,
					URL:            value.URL,
				}

				if !helper.IsItemAdded(sortedItems.Brands, item) {
					sortedItems.Brands = append(sortedItems.Brands, item)
				} else {
					fmt.Println("Brand already added: ")
				}
			}
		}

		// Get categories
		if aggregation.Type == consts.CATEGORY_FK {
			for _, value := range aggregation.Values {
				item := trendyol.Item{
					ID:             value.ID,
					Text:           value.Text,
					BeautifiedName: value.BeautifiedName,
					Type:           value.Type,
					URL:            value.URL,
				}

				if !helper.IsItemAdded(sortedItems.Categories, item) {
					sortedItems.Categories = append(sortedItems.Categories, item)
				} else {
					fmt.Println("Brand already added: ")
				}
			}
		}
	}
}
