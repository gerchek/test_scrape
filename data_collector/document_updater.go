package data_collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"scraper_trendyol/models/bagisto"
	models "scraper_trendyol/models/couch_db_model"
	"scraper_trendyol/pkg/logging"
	"strconv"

	"github.com/Jeffail/gabs/v2"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/sirupsen/logrus"
)

type Updater struct {
	scraper Scraper
}

func NewUpdater() (Updater, error) {
	updater := Updater{}

	scraper, err := NewScraper()

	if err != nil {
		return updater, err
	}

	updater.scraper = scraper

	return updater, nil
}

func getTotalDocumentCount() (int, error) {
	// _design/count/_view/countView is couch db function
	resp, err := http.Get(os.Getenv("couch_db_source") + os.Getenv("db_trendyol") + "/_design/count/_view/countView")

	if err != nil {
		logging.Error(err)
		return 0, err
	}

	defer resp.Body.Close()

	var response models.CountResponse

	body, _ := ioutil.ReadAll(resp.Body)

	// bodyString := string(body)
	// logrus.Infoln(bodyString)

	err = json.Unmarshal(body, &response)

	if err != nil {
		logging.Error(err)
		return 0, err
	}

	return response.Rows[0].Value, nil
}

func (updater Updater) InitUpdater() error {

	totalDocCount, err := getTotalDocumentCount()

	if err != nil {
		return err
	}

	i := 0
	for i < totalDocCount {
		currentProduct, _ := GetDataFromCouchDb(i)

		newProduct, _ := GetDataFromTrendyol(currentProduct, updater.scraper)

		isEqual := isEqual(currentProduct, newProduct)

		if !isEqual {
			logrus.Infoln("is not equal:  %s", i, currentProduct.Name)

			newProduct.Categories = nil

			newProduct.Categories = append(newProduct.Categories, currentProduct.Categories...)

			reqBodyBytes := new(bytes.Buffer)
			json.NewEncoder(reqBodyBytes).Encode(newProduct)

			jsonParsed, e := gabs.ParseJSON([]byte(reqBodyBytes.Bytes()))

			if e != nil {
				logrus.Infoln("ee ", e.Error())
			}

			// logrus.Infoln("aa: ", jsonParsed.String())

			updater.scraper.couchDBClient.UpdateProductData(jsonParsed)

		} else {
			logrus.Println("isEqual: %s", i, currentProduct.Name)
		}

		i++
	}

	return nil
}

func isEqual(currentProduct bagisto.BagistoProductModel, newProduct bagisto.BagistoProductModel) bool {

	// if !cmp.Equal(currentProduct, newProduct, cmpopts.IgnoreFields(currentProduct, "Categories")) {
	// 	logrus.Infoln("values are not the same %s", cmp.Diff(currentProduct, newProduct, cmpopts.IgnoreFields(currentProduct, "Categories")))
	// }

	return cmp.Equal(currentProduct, newProduct, cmpopts.IgnoreFields(currentProduct, "Attributes", "Brand", "Categories", "Cinsiyet", "Color" /* "ColorVariants", */, "ColorVariantCount", "Description", "Descriptions", "Name", "NameWithProductCode", "ShortDescription"))
}

func GetDataFromTrendyol(product bagisto.BagistoProductModel, scraper Scraper) (bagisto.BagistoProductModel, error) {

	var result bagisto.BagistoProductModel

	productId, _ := strconv.Atoi(product.ProductNumber)
	productGroupId, _ := strconv.Atoi(product.ProductGroupID)

	jsonProduct, err := scraper.GetProductDetailWithOptions(productId, productGroupId)

	if err != nil {
		logging.Error(err)
		return result, err
	}

	err = json.Unmarshal(jsonProduct.Bytes(), &result)
	if err != nil {
		logging.Error(err)
		return result, err
	}

	return result, err
}

func GetDataFromCouchDb(i int) (bagisto.BagistoProductModel, error) {
	var model bagisto.BagistoProductModel

	url := os.Getenv("couch_db_source") + os.Getenv("db_trendyol") + fmt.Sprintf("/_all_docs?include_docs=true&limit=1&skip=%v", i)

	resp, err := http.Get(url)

	if err != nil {
		logging.Error(err)
		return model, err
	}

	defer resp.Body.Close()

	var response models.ModelResponse

	data, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(data, &response)

	if err != nil {
		logging.Error(err)
		return model, err
	}

	return response.Rows[0].Doc, nil
}
