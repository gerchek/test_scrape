package bagisto

type BagistoCategory struct {
	ID               int         `json:"id"`
	Name             string      `json:"name"`
	Slug             string      `json:"slug"`
	TrendyolURL      string      `json:"trendyol_url"`
	LcwURL           string      `json:"lcw_url"`
	DisplayMode      string      `json:"display_mode"`
	ImageURL         interface{} `json:"image_url"`
	CategoryIconPath interface{} `json:"category_icon_path"`
	ProductLimit     int         `json:"product_limit"`
	Children         []BagistoCategory
}

type BagistoResponse struct {
	Data []BagistoCategory
}
