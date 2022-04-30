package data_collector

import (
	"encoding/json"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"scraper_trendyol/models/bagisto"
	"scraper_trendyol/pkg/logging"
	"time"
)

type Bagisto struct {
	Url string
}

//NewBagisto return new instance
func NewBagisto() (Bagisto, error) {

	dc := Bagisto{
		Url: os.Getenv("remote_url"), //"http://176.53.65.160/api/",
	}

	return dc, nil
}

//GetRootCategory return root category of Bagisto
func (b Bagisto) GetRootCategory() (bagisto.BagistoCategory, error) {
	rootCategory := bagisto.BagistoCategory{}

	link := b.Url +os.Getenv("descendant_category_url") 

	netClient := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   10 * time.Second,
				KeepAlive: 1 * time.Minute,
			}).DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
			MaxIdleConns:        3000,
			MaxIdleConnsPerHost: 3000,
			IdleConnTimeout:     60 * time.Second,
		},
		Timeout: 120 * time.Second,
	}
	resp, err := netClient.Get(link)

	//resp, err := http.Get(link)
	if err != nil {
		logging.Error(err)
		return rootCategory, err
	}

	defer resp.Body.Close()

	var respData bagisto.BagistoResponse

	byteData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteData, &respData)
	if err != nil {
		logging.Error(err)
		return rootCategory, err
	}

	rootCategory = respData.Data[0]

	return rootCategory, nil
}

/*type DataCollector struct {
	Url      string
	Database string

	// couch database
	couchDBClient couch_db.CouchDBClient
}

//NewDataCollector gets database as string parameter.
//writes data to that database
func NewDataCollector(database string) (DataCollector, error) {

	dc := DataCollector{
		Url:      os.Getenv("remote_url"),
		Database: database,
	}

	var err error
	dc.couchDBClient, err = couch_db.NewCouchDB(os.Getenv("db_trendyol"))
	if err != nil {
		logging.Error(err)
		return dc, err
	}

	return dc, nil
}

//GetRootCategory return root category of Bagisto
func (dc DataCollector) GetRootCategory() (bagisto.BagistoCategory, error) {
	rootCategory := bagisto.BagistoCategory{}

	link := dc.Url + os.Getenv("descendant_category_url")

	resp, err := http.Get(link)
	if err != nil {
		logging.Error(err)
		return rootCategory, err
	}

	defer resp.Body.Close()

	var respData bagisto.BagistoResponse

	byteData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteData, &respData)
	if err != nil {
		logging.Error(err)
		return rootCategory, err
	}

	rootCategory = respData.Data[0]

	return rootCategory, nil
}

func (dc DataCollector) BeginCollectingData(rootCategory bagisto.BagistoCategory) {
	// Iterate L1 categories
	for _, catL1 := range rootCategory.Children {
		// if display_mode is products_and_description, then it is category
		if consts.DISPLAY_MODE == catL1.DisplayMode {
			// Iterate L2 categories
			for _, catL2 := range catL1.Children {
				// Iterate L3 category
				for _, catL3 := range catL2.Children {

					categoryInfo := bagisto.Category{
						L1:   strconv.Itoa(catL1.ID),
						L2:   strconv.Itoa(catL2.ID),
						L3:   strconv.Itoa(catL3.ID),
						Slug: catL1.Slug,
					}

					if dc.Database == consts.DB_LCW {
						logrus.Infoln("collect data from LCW")
					}

					if dc.Database == consts.DB_TRENDYOL {
						dc.collectTrendyolData(categoryInfo, catL3.TrendyolURL)
					}

				}
			}
		}
	}
}

//collectTrendyolData iterates category link with pagination and collects data.
// 24 products retrieved in each iteration.
// Then saves products data to CouchDB
// Then sends request for next page.
func (dc DataCollector) collectTrendyolData(categoryInfo bagisto.Category, categoryUrl string) {
	logrus.Infoln("collectTrendyolData")

	baseLink := "https://public.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/"

	// category link to iterate
	linkCategory := baseLink + categoryUrl

	// iterate pages. init page from 1
	for i := 1; i < 3; i++ {
		url := linkCategory + "?pi=" + strconv.Itoa(i)
		products, err := GetProductsByPage(url, categoryInfo)
		if err == nil {
			// insert data to couch DB
			dc.couchDBClient.InsertData(products)
		}
	}
}
*/
