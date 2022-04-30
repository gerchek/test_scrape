package helper

import (
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	"scraper_trendyol/models"
	"scraper_trendyol/models/trendyol"
)

func IsItemAdded(items []trendyol.Item, keyItem trendyol.Item) bool {
	for _, item := range items {
		if item.ID == keyItem.ID {
			return true
		}
	}
	return false
}

func IsVariantAdded(items []*gabs.Container, keyItem models.SizeVariant) bool {
	for _, item := range items {

		var SizeVariants models.SizeVariant
		byteData, _ := json.Marshal(item.Data())
		json.Unmarshal(byteData, &SizeVariants)

		fmt.Println("SIZE_OPTION: ", SizeVariants, " : ", keyItem)
		if SizeVariants.Size == keyItem.Size && SizeVariants.Price == keyItem.Price {
			return true
		}
	}

	return false
}

func RemoveDuplicateValues(items []models.SizeVariant) []models.SizeVariant {
	keys := make(map[int]bool)
	list := make([]models.SizeVariant, 0)

	for _, item := range items {
		if _, value := keys[item.ItemNumber]; !value {
			keys[item.ItemNumber] = true
			list = append(list, item)
		}
	}

	return list
}

func IsLCWSizeVariantsAdded(items []models.SizeVariant, keyItem models.SizeVariant) bool {
	for _, option := range items {
		if option.Size == keyItem.Size {
			return true
		}
	}
	return false
}
