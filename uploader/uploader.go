package uploader

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"scraper_trendyol/models/bagisto"
	"scraper_trendyol/pkg/helper"
	"scraper_trendyol/pkg/logging"
	"strconv"

	"time"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Uploader struct {
	CreateUrl string
	UpdateUrl string
}

func NewUploader() (Uploader, error) {

	err := godotenv.Load()
	if err != nil {
		logging.Error(err)
		return Uploader{}, err
	}

	uploader := Uploader{
		CreateUrl: os.Getenv("remote_url") + os.Getenv("create_url"),
		UpdateUrl: os.Getenv("remote_url") + os.Getenv("update_url"),
	}

	return uploader, nil
}

func (uploader Uploader) UploadToRemoteServer(product bagisto.BagistoProductModel, operation string) error {

	logrus.Infoln("product: ", product.Name, " code: ", product.ProductNumber)

	var url string

	if operation == "create" {
		url = uploader.CreateUrl
	} else {
		url = uploader.UpdateUrl
	}

	start := time.Now()

	client := &http.Client{
		Timeout: time.Second * 50,
	}

	jsonData, err := json.Marshal(product)
	if err != nil {
		logging.Error(err)
		return err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logging.Error(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "0a358dd1-2b07-4cdf-9d9a-a68dac6bb5fc")

	response, err := client.Do(req)

	if err != nil {
		logging.Error(err)
		return err
	}

	if response.StatusCode != 200 {
		body, _ := ioutil.ReadAll(response.Body)
		bodyString := string(body)

		logging.Error(bodyString)

		errMsg := " bagisto insert error with code:  " + strconv.Itoa(response.StatusCode)
		errInsert := errors.New(errMsg)
		logging.Error(errInsert)
		return errInsert
	}

	defer response.Body.Close()

	helper.InsertedProductCount++

	duration := time.Since(start)

	logrus.Infoln("product insert duration: ", duration)

	return nil
}
