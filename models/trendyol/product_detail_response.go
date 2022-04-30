package trendyol

type TrendyolProductDetailResponse struct {
	IsSuccess  bool                       `json:"isSuccess"`
	StatusCode int                        `json:"statusCode"`
	Error      interface{}                `json:"error"`
	Result     TrendyolProductDetailModel `json:"result"`
	Headers    struct {
		Tysidecarcachable string `json:"tysidecarcachable"`
	} `json:"headers"`
}

type TrendyolProductDetailModel struct {
	Attributes []struct {
		Key struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"key"`
		Value struct {
			Name string `json:"name"`
			ID   int    `json:"id"`
		} `json:"value"`
		Starred bool `json:"starred"`
	} `json:"attributes"`
	AlternativeVariants []AlternativeVariant `json:"alternativeVariants"`
	Variants            []Variant            `json:"variants"`
	OtherMerchants      []interface{}        `json:"otherMerchants"`
	Campaign            struct {
		ID                 int    `json:"id"`
		Name               string `json:"name"`
		StartDate          string `json:"startDate"`
		EndDate            string `json:"endDate"`
		IsMultipleSupplied bool   `json:"isMultipleSupplied"`
		StockTypeID        int    `json:"stockTypeId"`
		URL                string `json:"url"`
		ShowTimer          bool   `json:"showTimer"`
	} `json:"campaign"`
	Category struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Hierarchy      string `json:"hierarchy"`
		Refundable     bool   `json:"refundable"`
		BeautifiedName string `json:"beautifiedName"`
		IsVASEnabled   bool   `json:"isVASEnabled"`
	} `json:"category"`
	Brand struct {
		IsVirtual      bool   `json:"isVirtual"`
		BeautifiedName string `json:"beautifiedName"`
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Path           string `json:"path"`
	} `json:"brand"`
	Color     string `json:"color"`
	MetaBrand struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		BeautifiedName string `json:"beautifiedName"`
		IsVirtual      bool   `json:"isVirtual"`
		Path           string `json:"path"`
	} `json:"metaBrand"`
	ShowVariants         bool          `json:"showVariants"`
	ShowSexualContent    bool          `json:"showSexualContent"`
	BrandCategoryBanners []interface{} `json:"brandCategoryBanners"`
	AllVariants          []struct {
		ItemNumber int     `json:"itemNumber"`
		Value      string  `json:"value"`
		InStock    bool    `json:"inStock"`
		Currency   string  `json:"currency"`
		Barcode    string  `json:"barcode"`
		Price      float64 `json:"price"`
	} `json:"allVariants"`
	OtherMerchantVariants []interface{} `json:"otherMerchantVariants"`
	InstallmentBanner     interface{}   `json:"installmentBanner"`
	IsVasEnabled          bool          `json:"isVasEnabled"`
	OriginalCategory      struct {
		ID             int    `json:"id"`
		Name           string `json:"name"`
		Hierarchy      string `json:"hierarchy"`
		Refundable     bool   `json:"refundable"`
		BeautifiedName string `json:"beautifiedName"`
		IsVASEnabled   bool   `json:"isVASEnabled"`
	} `json:"originalCategory"`
	Landings            []interface{} `json:"landings"`
	ID                  int           `json:"id"`
	ProductCode         string        `json:"productCode"`
	Name                string        `json:"name"`
	NameWithProductCode string        `json:"nameWithProductCode"`
	Description         string        `json:"description"`
	ContentDescriptions []struct {
		Description string `json:"description"`
		Bold        bool   `json:"bold"`
	} `json:"contentDescriptions"`
	ProductGroupID int    `json:"productGroupId"`
	Tax            int    `json:"tax"`
	BusinessUnit   string `json:"businessUnit"`
	MaxInstallment int    `json:"maxInstallment"`
	Gender         struct {
		Name string `json:"name"`
		ID   int    `json:"id"`
	} `json:"gender"`
	URL              string   `json:"url"`
	Images           []string `json:"images"`
	IsSellable       bool     `json:"isSellable"`
	IsBasketDiscount bool     `json:"isBasketDiscount"`
	HasStock         bool     `json:"hasStock"`
	Price            Price    `json:"price"`
	IsFreeCargo      bool     `json:"isFreeCargo"`
	Promotions       []struct {
		PromotionRemainingTime string `json:"promotionRemainingTime"`
		Type                   int    `json:"type"`
		Text                   string `json:"text"`
		ID                     int    `json:"id"`
		Link                   string `json:"link"`
	} `json:"promotions"`
	Merchant struct {
		IsSearchableMerchant bool          `json:"isSearchableMerchant"`
		Stickers             []interface{} `json:"stickers"`
		ID                   int           `json:"id"`
		Name                 string        `json:"name"`
		OfficialName         string        `json:"officialName"`
		CityName             string        `json:"cityName"`
		TaxNumber            string        `json:"taxNumber"`
		SellerScore          float64       `json:"sellerScore"`
		SellerScoreColor     string        `json:"sellerScoreColor"`
		DeliveryProviderName string        `json:"deliveryProviderName"`
		SellerLink           string        `json:"sellerLink"`
	} `json:"merchant"`
	DeliveryInformation struct {
		IsRushDelivery bool   `json:"isRushDelivery"`
		DeliveryDate   string `json:"deliveryDate"`
	} `json:"deliveryInformation"`
	CargoRemainingDays int  `json:"cargoRemainingDays"`
	IsMarketplace      bool `json:"isMarketplace"`
	ProductStamps      []struct {
		Type          string  `json:"type"`
		ImageURL      string  `json:"imageUrl"`
		Position      string  `json:"position"`
		AspectRatio   float64 `json:"aspectRatio"`
		Priority      int     `json:"priority"`
		PriceTagStamp bool    `json:"priceTagStamp,omitempty"`
	} `json:"productStamps"`
	HasHTMLContent    bool   `json:"hasHtmlContent"`
	FavoriteCount     int    `json:"favoriteCount"`
	UxLayout          string `json:"uxLayout"`
	IsDigitalGood     bool   `json:"isDigitalGood"`
	IsRunningOut      bool   `json:"isRunningOut"`
	ScheduledDelivery bool   `json:"scheduledDelivery"`
	RatingScore       struct {
		AverageRating     float64 `json:"averageRating"`
		TotalRatingCount  int     `json:"totalRatingCount"`
		TotalCommentCount int     `json:"totalCommentCount"`
	} `json:"ratingScore"`
	ShowStarredAttributes    bool   `json:"showStarredAttributes"`
	ReviewsURL               string `json:"reviewsUrl"`
	QuestionsURL             string `json:"questionsUrl"`
	SellerQuestionEnabled    bool   `json:"sellerQuestionEnabled"`
	SizeExpectationAvailable bool   `json:"sizeExpectationAvailable"`
	CrossPromotionAward      struct {
		AwardType  interface{} `json:"awardType"`
		AwardValue interface{} `json:"awardValue"`
		ContentID  int         `json:"contentId"`
		MerchantID int         `json:"merchantId"`
	} `json:"crossPromotionAward"`
	RushDeliveryMerchantListingExist bool `json:"rushDeliveryMerchantListingExist"`
	LowerPriceMerchantListingExist   bool `json:"lowerPriceMerchantListingExist"`
	ShowValidFlashSales              bool `json:"showValidFlashSales"`
	ShowExpiredFlashSales            bool `json:"showExpiredFlashSales"`
	WalletRebate                     struct {
		MinPrice    int     `json:"minPrice"`
		MaxPrice    int     `json:"maxPrice"`
		RebateRatio float64 `json:"rebateRatio"`
	} `json:"walletRebate"`
	IsArtWork bool `json:"isArtWork"`
}

type Price struct {
	ProfitMargin    int `json:"profitMargin"`
	DiscountedPrice struct {
		Text  string  `json:"text"`
		Value float64 `json:"value"`
	} `json:"discountedPrice"`
	SellingPrice struct {
		Text  string  `json:"text"`
		Value float64 `json:"value"`
	} `json:"sellingPrice"`
	OriginalPrice struct {
		Text  string  `json:"text"`
		Value float64 `json:"value"`
	} `json:"originalPrice"`
	Currency string `json:"currency"`
}

type Variant struct {
	AttributeID              int           `json:"attributeId"`
	AttributeName            string        `json:"attributeName"`
	AttributeType            string        `json:"attributeType"`
	AttributeValue           string        `json:"attributeValue"`
	Stamps                   []interface{} `json:"stamps"`
	Price                    Price         `json:"price"`
	FulfilmentType           string        `json:"fulfilmentType"`
	AttributeBeautifiedValue string        `json:"attributeBeautifiedValue"`
	IsWinner                 bool          `json:"isWinner"`
	ListingID                string        `json:"listingId"`
	Stock                    interface{}   `json:"stock"`
	Sellable                 bool          `json:"sellable"`
	AvailableForClaim        bool          `json:"availableForClaim"`
	Barcode                  string        `json:"barcode"`
	ItemNumber               int           `json:"itemNumber"`
	DiscountedPriceInfo      string        `json:"discountedPriceInfo"`
	HasCollectable           bool          `json:"hasCollectable"`
	RushDeliveryMerchantListingExist bool `json:"rushDeliveryMerchantListingExist"`
	LowerPriceMerchantListingExist   bool `json:"lowerPriceMerchantListingExist"`
}

type AlternativeVariant struct {
	AttributeValue           string      `json:"attributeValue"`
	AttributeBeautifiedValue string      `json:"attributeBeautifiedValue"`
	CampaignID               int         `json:"campaignId"`
	MerchantID               int         `json:"merchantId"`
	URLQuery                 string      `json:"urlQuery"`
	ListingID                string      `json:"listingId"`
	ItemNumber               int         `json:"itemNumber"`
	Barcode                  string      `json:"barcode"`
	Stock                    interface{} `json:"stock"`
	Quantity                 int         `json:"quantity"`
	Price                    Price       `json:"price"`
}
