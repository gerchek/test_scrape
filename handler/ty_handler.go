package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"scraper_trendyol/data_collector"
	"scraper_trendyol/excel_parser"
	"scraper_trendyol/models"
	"scraper_trendyol/models/couch_db_model"
	"scraper_trendyol/pkg/helper"
	"scraper_trendyol/pkg/logging"

	"strconv"
	"sync/atomic"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// type Ty_handler struct {
// 	// couch database client
// 	tt excel_parser.ExcelParser
// }

const (
	stateUnlocked uint32 = iota
	stateLocked
)

var (
	locker = stateUnlocked
	// tt     excel_parser.ExcelParser
)

func InitScraper(w http.ResponseWriter, r *http.Request) {

	// lock the request
	if !atomic.CompareAndSwapUint32(&locker, stateUnlocked, stateLocked) {
		w.WriteHeader(http.StatusTooManyRequests)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": "Scrape in progress!",
		})

		return
	}
	defer atomic.StoreUint32(&locker, stateUnlocked)

	keys := r.URL.Query()["product-limit"]

	productLimitStr := keys[0]

	logrus.Infoln("InitTrendyolScraper")
	logrus.Infoln("Total product limit: ", productLimitStr)

	productLimit, _ := strconv.Atoi(productLimitStr)

	helper.TotalProductLimit = productLimit
	helper.InsertedProductCount = 0

	bagisto, err := data_collector.NewBagisto()
	if err != nil {
		logging.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rootCategory, err := bagisto.GetRootCategory()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scraper, err := data_collector.NewScraper()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	scraper.BeginCollectingData(rootCategory)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": "Trendyol scraper completed successfully!",
	})
}

func ParseLink(w http.ResponseWriter, r *http.Request) {

	link := r.URL.Query().Get("url")

	logrus.Info("link: ", link)

	linkParser := data_collector.NewLinkParser(link)
	productGroupId, err := linkParser.ParseLink()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg":            "Link parsed successfully",
		"productGroupId": strconv.Itoa(productGroupId),
	})
}

func ParseExcel(w http.ResponseWriter, r *http.Request) {
	ep, err := excel_parser.NewExcelParser()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	err = ep.ParseExcelAndInsert()

	msg := "categories updated successfully"

	if err != nil {
		msg = err.Error()
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"msg": msg,
	})
}

func InitUpdater(w http.ResponseWriter, r *http.Request) {

	updater, err := data_collector.NewUpdater()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	errUpdater := updater.InitUpdater()

	if errUpdater != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"msg": err.Error(),
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"msg": "updated products",
	})
}

// ----------------------------------------------------------------------------------------------------------------

type people struct {
	Number int `json:"number"`
}

func GetExcel(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://admin:admin@localhost:5984/ty_categories/_all_docs?include_docs=true")
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	var response couch_db_model.CategoryModelResponse
	err = json.Unmarshal(data, &response)

	if err != nil {
		fmt.Println("error")
	}

	rows := response.Rows

	tmpl := template.Must(template.ParseFiles("./views/pages/data_table.html", "./views/partials/vertical_menu.html", "./views/layouts/default.html"))
	tmpl.ExecuteTemplate(w, "default", rows)

	// for i := 1; i < len(rows); i++ {
	// 	// sum += i
	// 	fmt.Println(rows[i])
	// }

	// fmt.Println(len(rows))

}

func CategoryDelete(w http.ResponseWriter, r *http.Request) {

	code := mux.Vars(r)["id"]

	fmt.Println(code)
	resp, err := http.Get("http://admin:admin@localhost:5984/ty_categories/" + code)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(data)

	var response models.Category
	err = json.Unmarshal(data, &response)

	if err != nil {
		fmt.Println("error")
	}

	rows := response

	req, err := http.NewRequest(http.MethodDelete, "http://admin:admin@localhost:5984/ty_categories/"+rows.ID+"?rev="+rows.Rev, nil)

	// fmt.Println(req)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("hello")
		http.Redirect(w, r, "/GetExcelData", http.StatusSeeOther)
	}
	// Create client
	client := &http.Client{}

	resp_1, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp_1.Body.Close()

}

func GetCategoryData(w http.ResponseWriter, r *http.Request) {

	code := mux.Vars(r)["id"]

	fmt.Println(code)
	resp, err := http.Get("http://admin:admin@localhost:5984/ty_categories/" + code)
	if err != nil {
		fmt.Println("error")
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)

	// fmt.Println(data)

	var response models.Category
	err = json.Unmarshal(data, &response)

	if err != nil {
		fmt.Println("error")
	}

	rows := response

	fmt.Println(rows)

}

func UpdateCategoryData(w http.ResponseWriter, r *http.Request) {

	// 1.
	payload, err := json.Marshal(map[string]interface{}{
		"_id":       "1000",
		"_rev":      "7-b2ec2b739a0c8b9d3fafcc9adef6e77a",
		"createdAt": "2022-02-26 13:21:36",
		"id":        "1001",
		"name":      "Bebek Kremi ve Yağı salamlar",
		"order":     "1",
		"parent_id": "2890",
		"sarga_id":  "NULL",
		"slug":      "bebek-islak-mendil",
		"updatedAt": "2022-02-26 13:21:36",
		"weight":    "0.15",
	})
	if err != nil {
		// log.Fatal(err)
		fmt.Println("1 error")
	}

	// 2.
	client := &http.Client{}
	url := "http://admin:admin@localhost:5984/ty_categories/1000"

	// 3.
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		log.Fatal(err)
	}

	// 4.
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	// 5.
	defer resp.Body.Close()

	// 6.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(string(body))

	// tmpl := template.Must(template.ParseFiles("./views/pages/get_data.html", "./views/partials/vertical_menu.html", "./views/layouts/default.html"))
	// tmpl.ExecuteTemplate(w, "default", "rows")

}
