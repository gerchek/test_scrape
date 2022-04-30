package data_collector

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"regexp"
	"scraper_trendyol/couch_db"
	"scraper_trendyol/excel_parser"
	"scraper_trendyol/models/bagisto"
	"scraper_trendyol/models/trendyol"
	consts "scraper_trendyol/pkg/const"
	"scraper_trendyol/pkg/helper"
	"scraper_trendyol/pkg/logging"
	"strconv"
	"strings"
	"sync"

	"github.com/Jeffail/gabs/v2"
	"github.com/sirupsen/logrus"
)

// TODO:
// 1. By Brand
// 2. By Count in stock

type Scraper struct {
	// couch database
	couchDBClient couch_db.CouchDBClient

	excelParser excel_parser.ExcelParser

	categories []int
}

func NewScraper() (Scraper, error) {

	scraper := Scraper{}

	cdb, err := couch_db.NewCouchDB(os.Getenv("db_trendyol"))

	if err != nil {
		return scraper, err
	}

	ep, err := excel_parser.NewExcelParser()

	if err != nil {
		return scraper, err
	}
	scraper.couchDBClient = cdb

	scraper.excelParser = ep

	return scraper, nil
}

func (scraper Scraper) BeginCollectingData(rootCategory bagisto.BagistoCategory) {

	// Iterate L1 categories
	for _, catL1 := range rootCategory.Children {
		// if display_mode is products_and_description, then it is category
		if consts.DISPLAY_MODE == catL1.DisplayMode {
			scraper.categories = append(scraper.categories, catL1.ID)

			// Iterate L2 categories
			for _, catL2 := range catL1.Children {
				scraper.categories = append(scraper.categories, catL2.ID)

				// Iterate L3 category
				for _, catL3 := range catL2.Children {

					scraper.categories = append(scraper.categories, catL3.ID)

					logrus.Infoln("catL3: ", catL3.Name, ", limit: ", catL3.ProductLimit)

					if len(strings.Trim(catL3.TrendyolURL, " ")) != 0 {
						isProductLimitReached := scraper.iterateCategoryPage(catL3.TrendyolURL, catL3.ProductLimit)

						if isProductLimitReached {
							return
						}
					}
				}
			}
		}
	}
}

//iterateCategoryPage iterates category link with pagination and collects data.
// 24 products retrieved in each iteration.
func (scraper Scraper) iterateCategoryPage(categoryUrl string, productLimit int) bool {
	baseLink := "https://public.trendyol.com/discovery-web-searchgw-service/v2/api/infinite-scroll/"

	// category link to iterate
	linkCategory := baseLink + categoryUrl

	i := strings.Index(linkCategory, "?")

	if i > -1 {
		linkCategory += "&"
	} else {
		linkCategory += "?"
	}

	// FIXME: change product limit
	// iterate pages. init page from 1
	for i := 1; i < productLimit/24; i++ {

		logrus.Infoln("page: ", i)

		url := linkCategory + "pi=" + strconv.Itoa(i)

		if helper.IsProductLimitReached() {
			return true
		}

		scraper.getProductsByPage(url)
	}
	// all category products scraped, but product limit not reached
	// so, continue to next category
	return false
}

//getProductsByPage Gets products using pagination. url param is category link to iterate.
func (scraper Scraper) getProductsByPage(url string) {

	logrus.Infoln("url: ", url)

	resp, err := http.Get(url)
	if err != nil {
		logging.Error(err)
		return
	}

	defer resp.Body.Close()

	var respData trendyol.TrendyolProductResponse
	byteData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteData, &respData)
	if err != nil {
		logging.Error(err)
		return
	}

	// iterate products and convert to json
	for _, product := range respData.Result.Products {

		isProductInserted := scraper.couchDBClient.IsProductExists(strconv.Itoa(product.ProductGroupID))

		logrus.Infoln("isProductInserted: ", isProductInserted)

		if !isProductInserted {
			json, errGPD := scraper.GetProductDetailWithOptions(product.ID, product.ProductGroupID)
			if errGPD == nil {
				// return breaks the loop and exits from function
				// don't use return here
				// insertErr := scraper.InsertData(json)
				// if insertErr != nil {
				// 	return
				// }
				scraper.InsertData(json)
			}
		} else {
			logrus.Infoln("product already inserted")
		}

	}
}

//GetProductDetailWithOptions returns JSON with variants
func (scraper Scraper) GetProductDetailWithOptions(productId, productGroupId int) (*gabs.Container, error) {

	primaryProductDetail, err := scraper.GetProductDetails(strconv.Itoa(productId))

	//FIXME: filter brands with category
	// primaryProductDetail.Brand.BeautifiedName ==

	if err != nil {
		return nil, err
	}

	productDetailJSON := scraper.createJSONFromModel(primaryProductDetail, scraper.categories)

	colorVariants, err := getProductColorVariants(productGroupId)

	if err != nil {
		return nil, err
	}

	// Get the color variant products
	for _, slicingAttribute := range colorVariants.Result.SlicingAttributes {
		for _, attribute := range slicingAttribute.Attributes {
			for _, content := range attribute.Contents {
				// Don't fetch primary product again.
				if content.ID != primaryProductDetail.ID {
					productVariantDetail, errGPD := scraper.GetProductDetails(strconv.Itoa(content.ID))
					if errGPD == nil {
						productVariantJSON := scraper.createJSONFromModel(productVariantDetail, scraper.categories)
						if productVariantJSON != nil {
							productDetailJSON.ArrayAppend(productVariantJSON, "color_variants")
						}
					}
				}
			}
		}
		// Color variant count
		if len(slicingAttribute.Attributes) > 0 {
			productDetailJSON.Set(len(slicingAttribute.Attributes), "color_variant_count")
		}
	}

	return productDetailJSON, nil
}

func (scraper Scraper) InsertData(json *gabs.Container) error {
	return scraper.couchDBClient.InsertProductData(json)
}

// GetProductDetails return JSON object. Merges color option
func (scraper Scraper) GetProductDetails(productId string) (trendyol.TrendyolProductDetailModel, error) {

	productDetailModel := trendyol.TrendyolProductDetailModel{}

	// set linearVariants false to get variants
	link := "https://public.trendyol.com/discovery-web-productgw-service/api/productDetail/" + productId + "?storefrontId=1&culture=tr-TR&linearVariants=false"

	resp, err := http.Get(link)

	if err != nil {
		logging.Error(err)
		return productDetailModel, err
	}

	defer resp.Body.Close()

	var respData trendyol.TrendyolProductDetailResponse
	byteData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteData, &respData)
	if err != nil {
		logging.Error(err)
		return productDetailModel, err
	}

	productDetailModel = respData.Result

	return productDetailModel, nil
}

//getProductColorVariants returns color options of product
func getProductColorVariants(productGroupId int) (trendyol.TrendyolProductVariantsResponse, error) {

	link := "https://public.trendyol.com/discovery-web-productgw-service/api/productGroup/" + strconv.Itoa(productGroupId) + "?storefrontId=1&culture=tr-TR"

	resp, err := http.Get(link)
	if err != nil {
		logging.Error(err)
		return trendyol.TrendyolProductVariantsResponse{}, err
	}

	defer resp.Body.Close()

	var respData trendyol.TrendyolProductVariantsResponse
	byteData, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(byteData, &respData)
	if err != nil {
		logging.Error(err)
		return trendyol.TrendyolProductVariantsResponse{}, err
	}

	return respData, nil
}

func (scraper Scraper) createJSONFromModel(model trendyol.TrendyolProductDetailModel, categories []int) *gabs.Container {

	// get weight from categories database
	weight := scraper.excelParser.GetCategoryWeight(model.Category.BeautifiedName)

	json := gabs.New()
	json.Set("trendyol", "vendor")
	json.Set("p-"+strconv.Itoa(model.ID), "sku")
	json.Set(strconv.Itoa(model.ID), "product_number")
	json.Set(model.ProductCode, "product_code")
	json.Set(strconv.Itoa(model.ProductGroupID), "product_group_id")
	json.Set(model.Name, "name")
	json.Set(model.IsSellable, "isSellable")
	json.Set(model.FavoriteCount, "favoriteCount")
	json.Set(weight, "weight")
	json.Set(model.NameWithProductCode, "name_with_product_code")
	json.Set("https://www.trendyol.com"+model.URL, "url_key")
	json.Set(model.Images, "images")
	json.Set(model.Brand.Name, "brand")
	json.Set(model.Gender.Name, "cinsiyet")
	json.Set(model.Description, "description")
	json.Set(model.ContentDescriptions, "descriptions")
	json.Set(model.Description, "short_description")
	json.Set(model.Color, "color")
	json.SetP(model.Price.OriginalPrice, "price.originalPrice")
	json.SetP(model.Price.SellingPrice, "price.sellingPrice")
	json.SetP(model.Price.DiscountedPrice, "price.discountedPrice")

	// if product is parsed by link, then assign it root category
	if len(categories) == 0 {
		categories = append(categories, 1)
	}

	json.Set(categories, "categories")

	//json.Set(model.Attributes, "attributes")
	attributes := make([]map[string]string, 0)
	for _, attr := range model.Attributes {
		var re = regexp.MustCompile(`/[^A-Z0-9]/ig`)
		keyStr := re.ReplaceAllString(attr.Key.Name, `_`)
		key := strings.ToLower(keyStr)
		attribute := map[string]string{
			key: attr.Value.Name,
		}

		attributes = append(attributes, attribute)
	}

	json.Set(attributes, "attributes")

	// FIXME: all product variant can be less than 3 count.
	// look fot this

	// if show variants, then it is configurable product.
	if model.ShowVariants {
		var variants []trendyol.Variant

		for i := 0; i < len(model.Variants); i++ {
			variant := model.Variants[i]

			// if stockType is nil, then it does not have count like, 1,2,3
			stockType := reflect.TypeOf(variant.Stock)

			if stockType == nil && variant.Sellable {
				variants = append(variants, variant)
			}
		}

		for i := 0; i < len(model.AlternativeVariants); i++ {
			alternativeVariant := model.AlternativeVariants[i]

			stockType := reflect.TypeOf(alternativeVariant.Stock)

			if stockType == nil && len(variants) > 0 {

				// get the first variant for attribute info
				fv := variants[0]

				variant := trendyol.Variant{
					AttributeID:    fv.AttributeID,
					AttributeName:  fv.AttributeName,
					AttributeType:  fv.AttributeType,
					AttributeValue: alternativeVariant.AttributeValue,
					Price:          alternativeVariant.Price,
				}
				variants = append(variants, variant)
			}
		}
		json.Set(variants, "size_variants")
	}

	return json
}

func GetProductsByPageChn(url string) ([]trendyol.TrendyolProductModel, error) {
	// TODO:
	// 1. By Brand
	// 2. By Count in stock

	totalProducts := make([]trendyol.TrendyolProductModel, 0)

	wg := sync.WaitGroup{}
	for i := 0; i <= 100; i++ {
		wg.Add(1)
		link := url + "?pi=" + strconv.Itoa(i)

		// get products using concurrency
		go func(link string) {
			resp, err := http.Get(link)
			if err != nil {
				logrus.Error("GetProductsByPageChn error: ", err)
			}

			defer resp.Body.Close()

			var respData trendyol.TrendyolProductResponse
			byteData, _ := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(byteData, &respData)
			if err != nil {
				logrus.Error("error convert to byte: ", err)
			}
			totalProducts = append(totalProducts, respData.Result.Products...)
			wg.Done()

		}(link)

	}
	wg.Wait()

	return totalProducts, nil
}
