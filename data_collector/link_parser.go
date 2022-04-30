package data_collector

import (
	"errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"scraper_trendyol/pkg/logging"
	"strings"
)

type LinkParser struct {
	link string
}

func NewLinkParser(link string) LinkParser {
	return LinkParser{link: link}
}

func (l LinkParser) ParseLink() (int, error) {

	logrus.Infoln("link: ", l.link)

	productId := ""

	if isShortLink(l.link) {
		productId = getProductIdFromShortLink(l.link)
	} else {
		productId = getProductIdFromLink(l.link)
	}

	if len(productId) == 0 {
		parseErr := errors.New("can not parse product id")
		logging.Error(parseErr)
		return 0, parseErr
	}

	logrus.Infoln("productId: ", productId)

	scraper, err := NewScraper()
	if err != nil {
		logging.Error(err)
		return 0, err
	}

	// TODO: think about it, works twice
	productDetail, err := scraper.GetProductDetails(productId)
	if err != nil {
		logging.Error(err)
		return 0, err
	}

	json, err := scraper.GetProductDetailWithOptions(productDetail.ID, productDetail.ProductGroupID)
	if err != nil {
		logging.Error(err)
		return 0, err
	}

	err = scraper.InsertData(json)
	if err != nil {
		logging.Error(err)
		return 0, err
	}

	return productDetail.ProductGroupID, nil
}

func getProductIdFromShortLink(shortLink string) string {

	var productId string

	resp, err := http.Get(shortLink)

	if err != nil {
		logging.Error(err)
		return productId
	}

	defer resp.Body.Close()

	url := resp.Request.URL.Path

	logrus.Infoln("url: ", url)

	productId = getProductIdFromLink(url)

	logrus.Infoln("productId: ", productId)

	/*byteData, _ := ioutil.ReadAll(resp.Body)

	html := string(byteData)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		logging.Error(err)
		return productId
	}

	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		textValue := s.Text()

		if strings.Contains(textValue, "window.__PRODUCT_DETAIL_APP_INITIAL_STATE__=") {
			productId = parseProductId(textValue)
		}
	})*/

	return productId
}

func getProductIdFromLink(link string) string {
	var productId string

	if strings.Contains(link, "?") {
		link = strings.Split(link, "?")[0]
	}

	strArr := strings.Split(link, "-")
	productId = strArr[len(strArr)-1]

	return productId
}

func isShortLink(link string) bool {
	return !strings.Contains(link, "trendyol.com")
}

// func parseProductId(str string) string {

// 	var sb strings.Builder
// 	r := regexp.MustCompile(`=\s*(.*?)\s*};`)
// 	matches := r.FindAllStringSubmatch(str, -1)
// 	for _, v := range matches {
// 		sb.WriteString(v[1])
// 	}

// 	sb.WriteString("}")

// 	shortLinkResp := models.ShortLinkResponse{}

// 	json.Unmarshal([]byte(sb.String()), &shortLinkResp)

// 	return strconv.Itoa(shortLinkResp.Product.ID)
// }
