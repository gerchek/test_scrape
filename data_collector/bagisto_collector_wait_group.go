package data_collector

/*
import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"scraper_trendyol/couch_db"
	"scraper_trendyol/models/bagisto"
	"scraper_trendyol/pkg/const"
)

type BagistoCollectorWG struct {
	Url string
}

func NewBagistoCollectorWG() BagistoCollectorWG {
	return BagistoCollectorWG{
		Url: "http://176.53.65.160/api/",
	}
}

func (bc BagistoCollectorWG) GetDescendantCategoriesWG(url string) {
	link := bc.Url + url

	resp, err := http.Get(link)

	if err != nil {
		logrus.Error("GetDescendantCategories error: ", err)
	}

	defer resp.Body.Close()

	var respData bagisto.BagistoResponse

	byteData, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(byteData, &respData)

	if err != nil {
		logrus.Errorf("error convert to byte: %e", err)
	}

	// root category
	respCategory := respData.Data[0]

	wg := sync.WaitGroup{}

	// Iterate L1 categories
	for _, catL1 := range respCategory.Children {
		// if display_mode is products_and_description, then it is category
		if consts.DISPLAY_MODE == catL1.DisplayMode {
			// Iterate L2 categories
			for _, catL2 := range catL1.Children {
				// Iterate L3 category
				for _, catL3 := range catL2.Children {
					baseLink := "https://public.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/"
					// Category link to iterate
					linkCategory := baseLink + catL3.TrendyolURL

					categoryInfo := bagisto.Category{
						L1:   strconv.Itoa(catL1.ID),
						L2:   strconv.Itoa(catL2.ID),
						L3:   strconv.Itoa(catL3.ID),
						Slug: catL1.Slug,
					}

					wg.Add(1)

					go func(link string) {

						products, errGetProduct := GetProductsByPage(link, categoryInfo)

						if errGetProduct != nil {
							logrus.Error("error GetProductsByPage ", errGetProduct)
							return
						}

						couchDb, _ := couch_db.NewCouchDB("trendyol_data")
						couchDb.InsertData(products)

						wg.Done()
					}(linkCategory)
				}
			}
		}
	}
	wg.Wait()
}
*/
