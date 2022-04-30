package couch_db_model

import (
	"scraper_trendyol/models"
	"scraper_trendyol/models/bagisto"
)

type CountResponse struct {
	Rows []struct {
		Key   interface{} `json:"key"`
		Value int         `json:"value"`
	} `json:"rows"`
}

type ModelResponse struct {
	Rows []struct {
		Key   string                      `json:"key"`
		Value map[string]interface{}      `json:"value"`
		Doc   bagisto.BagistoProductModel `json:"doc"`
	} `json:"rows"`
}

type CategoryModelResponse struct {
	Rows []struct {
		Doc models.Category `json:"doc"`
	} `json:"rows"`
}
