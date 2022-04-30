package models

import "scraper_trendyol/models/trendyol"

type SizeVariant struct {
	ItemNumber int            `json:"itemNumber"`
	Price      trendyol.Price `json:"price"`
	Size       string         `json:"size"`
	Stock      int            `json:"stock"`
}
