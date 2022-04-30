package trendyol

type TrendyolProductResponse struct {
	IsSuccess  bool        `json:"isSuccess"`
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error"`
	Result     struct {
		SlpName         string                `json:"slpName"`
		Products        []TrendyolProductModel `json:"products"`
		TotalCount      int                   `json:"totalCount"`
		RoughTotalCount string                `json:"roughTotalCount"`
		SearchStrategy  string                `json:"searchStrategy"`
		Title           string                `json:"title"`
		UxLayout        string                `json:"uxLayout"`
		QueryTerm       string                `json:"queryTerm"`
		PageIndex       int                   `json:"pageIndex"`
		Widgets         []interface {
		} `json:"widgets"`
	} `json:"result"`
	Headers struct {
		Tysidecarcachable string `json:"tysidecarcachable"`
	} `json:"headers"`
}

type TrendyolProductModel struct {
	ID       int      `json:"id"`
	Name     string   `json:"name"`
	Images   []string `json:"images"`
	ImageAlt string   `json:"imageAlt"`
	Brand    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"brand"`
	Tax          int    `json:"tax"`
	BusinessUnit string `json:"businessUnit"`
	RatingScore  struct {
		AverageRating float64 `json:"averageRating"`
		TotalCount    int     `json:"totalCount"`
	} `json:"ratingScore"`
	ShowSexualContent bool `json:"showSexualContent"`
	ProductGroupID    int  `json:"productGroupId"`
	InStock           bool `json:"inStock"`
	Sections          []struct {
		ID string `json:"id"`
	} `json:"sections"`
	Variants []struct {
		AttributeValue string `json:"attributeValue"`
		AttributeName  string `json:"attributeName"`
		Price          struct {
			DiscountedPrice float64 `json:"discountedPrice"`
			BuyingPrice     int     `json:"buyingPrice"`
			OriginalPrice   float64 `json:"originalPrice"`
			SellingPrice    float64 `json:"sellingPrice"`
		} `json:"price"`
		ListingID           string `json:"listingId"`
		CampaignID          int    `json:"campaignId"`
		MerchantID          int    `json:"merchantId"`
		DiscountedPriceInfo string `json:"discountedPriceInfo"`
	} `json:"variants"`
	CategoryHierarchy string `json:"categoryHierarchy"`
	CategoryID        int    `json:"categoryId"`
	CategoryName      string `json:"categoryName"`
	URL               string `json:"url"`
	MerchantID        int    `json:"merchantId"`
	CampaignID        int    `json:"campaignId"`
	Price             struct {
		SellingPrice             float64 `json:"sellingPrice"`
		OriginalPrice            float64 `json:"originalPrice"`
		ManipulatedOriginalPrice float64 `json:"manipulatedOriginalPrice"`
		DiscountedPrice          float64 `json:"discountedPrice"`
		BuyingPrice              int     `json:"buyingPrice"`
	} `json:"price"`
	Promotions           []interface{} `json:"promotions"`
	RushDeliveryDuration int           `json:"rushDeliveryDuration"`
	FreeCargo            bool          `json:"freeCargo"`
	Margin               int           `json:"margin"`
	CampaignName         string        `json:"campaignName"`
	ListingID            string        `json:"listingId"`
	WinnerVariant        string        `json:"winnerVariant"`
	ItemNumber           int           `json:"itemNumber"`
	DiscountedPriceInfo  string        `json:"discountedPriceInfo"`
	Stamps               []struct {
		ImageURL    string  `json:"imageUrl"`
		Type        string  `json:"type"`
		Position    string  `json:"position"`
		AspectRatio float64 `json:"aspectRatio"`
		Priority    int     `json:"priority"`
	} `json:"stamps,omitempty"`
}
