package bagisto

import "scraper_trendyol/models/trendyol"

type BagistoProductModel struct {
	Attributes        []map[string]string   `json:"attributes"`
	Brand             string                `json:"brand"`
	Categories        []int                 `json:"categories"`
	Cinsiyet          string                `json:"cinsiyet"`
	Color             string                `json:"color"`
	ColorVariantCount int                   `json:"color_variant_count"`
	ColorVariants     []BagistoProductModel `json:"color_variants"`
	Description       string                `json:"description"`
	IsSellable        bool                  `json:"isSellable"`
	FavoriteCount     bool                  `json:"favorite_count"`
	Descriptions      []struct {
		Description string `json:"description"`
		Bold        bool   `json:"bold"`
	} `json:"descriptions"`
	Images              []string           `json:"images"`
	Name                string             `json:"name"`
	NameWithProductCode string             `json:"name_with_product_code"`
	Price               trendyol.Price     `json:"price"`
	ProductCode         string             `json:"product_code"`
	ProductGroupID      string             `json:"product_group_id"`
	ProductNumber       string             `json:"product_number"`
	ShortDescription    string             `json:"short_description"`
	SizeVariants        []trendyol.Variant `json:"size_variants"`
	Sku                 string             `json:"sku"`
	Stock               interface{}        `json:"stock"`
	URLKey              string             `json:"url_key"`
	Vendor              string             `json:"vendor"`
	Weight              string             `json:"weight"`
}
