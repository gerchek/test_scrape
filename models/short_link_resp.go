package models

type ShortLinkResponse struct {
	Product struct {
		Attributes          []interface{} `json:"attributes"`
		AlternativeVariants []interface{} `json:"alternativeVariants"`
		Variants            []struct {
			AttributeID    int    `json:"attributeId"`
			AttributeName  string `json:"attributeName"`
			AttributeType  string `json:"attributeType"`
			AttributeValue string `json:"attributeValue"`
			Stamps         []struct {
				Type int    `json:"type"`
				Text string `json:"text"`
			} `json:"stamps"`
			Price struct {
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
			} `json:"price"`
			FulfilmentType                   string      `json:"fulfilmentType"`
			AttributeBeautifiedValue         string      `json:"attributeBeautifiedValue"`
			IsWinner                         bool        `json:"isWinner"`
			ListingID                        string      `json:"listingId"`
			Stock                            interface{} `json:"stock"`
			Sellable                         bool        `json:"sellable"`
			AvailableForClaim                bool        `json:"availableForClaim"`
			Barcode                          string      `json:"barcode"`
			ItemNumber                       int         `json:"itemNumber"`
			DiscountedPriceInfo              string      `json:"discountedPriceInfo"`
			HasCollectable                   bool        `json:"hasCollectable"`
			RushDeliveryMerchantListingExist bool        `json:"rushDeliveryMerchantListingExist"`
			LowerPriceMerchantListingExist   bool        `json:"lowerPriceMerchantListingExist"`
		} `json:"variants"`
		OtherMerchants []interface{} `json:"otherMerchants"`
		Campaign       struct {
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
		Color     interface{} `json:"color"`
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
		CategoryTopRankings   interface{}   `json:"categoryTopRankings"`
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
		} `json:"gender"`
		URL              string   `json:"url"`
		Images           []string `json:"images"`
		IsSellable       bool     `json:"isSellable"`
		IsBasketDiscount bool     `json:"isBasketDiscount"`
		HasStock         bool     `json:"hasStock"`
		Price            struct {
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
		} `json:"price"`
		IsFreeCargo bool `json:"isFreeCargo"`
		Promotions  []struct {
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
		CargoRemainingDays int           `json:"cargoRemainingDays"`
		IsMarketplace      bool          `json:"isMarketplace"`
		ProductStamps      []interface{} `json:"productStamps"`
		HasHTMLContent     bool          `json:"hasHtmlContent"`
		FavoriteCount      int           `json:"favoriteCount"`
		UxLayout           string        `json:"uxLayout"`
		IsDigitalGood      bool          `json:"isDigitalGood"`
		IsRunningOut       bool          `json:"isRunningOut"`
		ScheduledDelivery  bool          `json:"scheduledDelivery"`
		RatingScore        struct {
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
	} `json:"product"`
	HTMLContent interface{} `json:"htmlContent"`
	User        struct {
		LoggedIn bool `json:"loggedIn"`
		IsBuyer  bool `json:"isBuyer"`
	} `json:"user"`
	Configuration struct {
		HTMLContentCSSURL              string `json:"htmlContentCssUrl"`
		HTMLContentJsURL               string `json:"htmlContentJsUrl"`
		StorefrontID                   int    `json:"storefrontId"`
		ShowPaymentMethodsAndShippings bool   `json:"showPaymentMethodsAndShippings"`
		ShowDiscountInformation        bool   `json:"showDiscountInformation"`
		IsDynamicRender                bool   `json:"isDynamicRender"`
		Widgets                        []struct {
			Name  string `json:"name"`
			Order int    `json:"order"`
		} `json:"widgets"`
		MaxWidgetCount int `json:"maxWidgetCount"`
		CallToActions  []struct {
			Name     string `json:"name"`
			Selector string `json:"selector"`
		} `json:"callToActions"`
		Discovers []struct {
			Name  string `json:"name"`
			Link  string `json:"link"`
			Image string `json:"image"`
		} `json:"discovers"`
		DiscoverWidgetTitle                         string     `json:"discoverWidgetTitle"`
		IsInternational                             bool       `json:"isInternational"`
		Culture                                     string     `json:"culture"`
		CdnURL                                      string     `json:"cdnUrl"`
		ClaimInfoText                               string     `json:"claimInfoText"`
		QuestionsAndAnswersEnabled                  bool       `json:"questionsAndAnswersEnabled"`
		RatingReviewEnabled                         bool       `json:"ratingReviewEnabled"`
		OtherMerchantsEnabled                       bool       `json:"otherMerchantsEnabled"`
		RatingReviewLikesEnabled                    bool       `json:"ratingReviewLikesEnabled"`
		NewRatingSummaryEnabled                     bool       `json:"newRatingSummaryEnabled"`
		RedirectReviewsPageEnabled                  bool       `json:"redirectReviewsPageEnabled"`
		LanguageCode                                string     `json:"languageCode"`
		SizeCharts                                  [][]string `json:"sizeCharts"`
		EnhancedEcommerceEnabled                    bool       `json:"enhancedEcommerceEnabled"`
		AddReviewEnabled                            bool       `json:"addReviewEnabled"`
		ReviewReportAbuseEnabled                    bool       `json:"reviewReportAbuseEnabled"`
		ShowImageOnProductCommentsEnabled           bool       `json:"showImageOnProductCommentsEnabled"`
		CrossProductsTitle                          string     `json:"crossProductsTitle"`
		CrossProductsEnabled                        bool       `json:"crossProductsEnabled"`
		RecommendationEnabled                       bool       `json:"recommendationEnabled"`
		ProductGroupEnabled                         bool       `json:"productGroupEnabled"`
		HTMLContentEnabled                          bool       `json:"htmlContentEnabled"`
		ProductAttributesEnabled                    bool       `json:"productAttributesEnabled"`
		PublishCriteriaEnabled                      bool       `json:"publishCriteriaEnabled"`
		SellerAreaEnabled                           bool       `json:"sellerAreaEnabled"`
		SellerPointLowLimit                         int        `json:"sellerPointLowLimit"`
		SizeChartURL                                string     `json:"sizeChartUrl"`
		ProductDetailMetaDescription                string     `json:"productDetailMetaDescription"`
		SchemaJSONEnabled                           bool       `json:"schemaJsonEnabled"`
		SiteAddress                                 string     `json:"siteAddress"`
		NotifyMeEnabled                             bool       `json:"notifyMeEnabled"`
		NotifyMeCount                               int        `json:"notifyMeCount"`
		RecommendationAbTestValue                   string     `json:"recommendationAbTestValue"`
		LastProductCountAbTestValue                 string     `json:"lastProductCountAbTestValue"`
		RecoCardAbTestValue                         string     `json:"recoCardAbTestValue"`
		AbTestingCookieName                         string     `json:"abTestingCookieName"`
		ClientSideReviewsEnabled                    bool       `json:"clientSideReviewsEnabled"`
		ClientSideHTMLContentEnabled                bool       `json:"clientSideHtmlContentEnabled"`
		AlternativeVariantsEnabled                  bool       `json:"alternativeVariantsEnabled"`
		RelatedCategoryEnabled                      bool       `json:"relatedCategoryEnabled"`
		RelatedCategoryAbTestVariant                string     `json:"relatedCategoryAbTestVariant"`
		RelatedCategoryAbTestValue                  string     `json:"relatedCategoryAbTestValue"`
		RelatedCategoryTitleTooltipThreshold        int        `json:"relatedCategoryTitleTooltipThreshold"`
		RelatedCategoryCountLimit                   int        `json:"relatedCategoryCountLimit"`
		RelatedCategoryImageVirtualBrandIds         []int      `json:"relatedCategoryImageVirtualBrandIds"`
		RelatedCategoryVirtualBrandImagePath        string     `json:"relatedCategoryVirtualBrandImagePath"`
		DefaultBrandCategoryCombinationImageURL     string     `json:"defaultBrandCategoryCombinationImageUrl"`
		LegalRequirementCacheDuration               int        `json:"legalRequirementCacheDuration"`
		DigitalGoodsDeliveryText                    string     `json:"digitalGoodsDeliveryText"`
		RecoCrossCustomStampsEnabled                bool       `json:"recoCrossCustomStampsEnabled"`
		CanShowSizeChartButton                      bool       `json:"canShowSizeChartButton"`
		ProductDetailImprovedBreadcrumbEnabled      bool       `json:"productDetailImprovedBreadcrumbEnabled"`
		SellerShippingEnabled                       bool       `json:"sellerShippingEnabled"`
		ProductDetailReportAbuseEnabled             bool       `json:"productDetailReportAbuseEnabled"`
		ProductDetailReportAbuseItems               string     `json:"productDetailReportAbuseItems"`
		ScheduledDeliveryWarningMessage             string     `json:"scheduledDeliveryWarningMessage"`
		InstallmentStampAmountText                  string     `json:"installmentStampAmountText"`
		StarredAttributesLimit                      int        `json:"starredAttributesLimit"`
		SellerStoreLinkEnabled                      bool       `json:"sellerStoreLinkEnabled"`
		FeaturedCardFavButtonEnabled                bool       `json:"featuredCardFavButtonEnabled"`
		OpenReviewModalEnabled                      bool       `json:"openReviewModalEnabled"`
		InstallmentCountToDisplay                   int        `json:"installmentCountToDisplay"`
		CurrencySymbol                              string     `json:"currencySymbol"`
		NewSearchEnabled                            bool       `json:"newSearchEnabled"`
		StampType                                   string     `json:"stampType"`
		MemberGwURL                                 string     `json:"memberGwUrl"`
		GetNotifyMePreferencesFromMemberGw          bool       `json:"getNotifyMePreferencesFromMemberGw"`
		PublicProductGwURL                          string     `json:"publicProductGwUrl"`
		PublicSdcProductGwURL                       string     `json:"publicSdcProductGwUrl"`
		PublicMdcProductGwURL                       string     `json:"publicMdcProductGwUrl"`
		PublicMdcRecoGwURL                          string     `json:"publicMdcRecoGwUrl"`
		ProductRecommendationVersion                int        `json:"productRecommendationVersion"`
		ProductDetailSimilarProductsButtonAbTest    string     `json:"productDetailSimilarProductsButtonAbTest"`
		PublicSdcCheckoutGwURL                      string     `json:"publicSdcCheckoutGwUrl"`
		PublicMdcCheckoutGwURL                      string     `json:"publicMdcCheckoutGwUrl"`
		AddToBasketOnCheckoutgwEnabled              bool       `json:"addToBasketOnCheckoutgwEnabled"`
		CollectionRecommendationEnabled             bool       `json:"collectionRecommendationEnabled"`
		ShowCargoRemainingDays                      bool       `json:"showCargoRemainingDays"`
		PdpLinearVariantsEnabled                    bool       `json:"pdpLinearVariantsEnabled"`
		ShowLinearAlternativeVariants               bool       `json:"showLinearAlternativeVariants"`
		SellerStoreURL                              string     `json:"sellerStoreUrl"`
		MinSellerFollowerCount                      int        `json:"minSellerFollowerCount"`
		ShowSellerFollowerCount                     bool       `json:"showSellerFollowerCount"`
		QuestionCountToShow                         int        `json:"questionCountToShow"`
		CollectableCouponsEnabled                   bool       `json:"collectableCouponsEnabled"`
		PublicSdcCouponGwURL                        string     `json:"publicSdcCouponGwUrl"`
		PublicMdcCouponGwURL                        string     `json:"publicMdcCouponGwUrl"`
		ShowInstallmentCampaign                     bool       `json:"showInstallmentCampaign"`
		VasServiceEnabled                           bool       `json:"vasServiceEnabled"`
		ShowCrossPromotions                         bool       `json:"showCrossPromotions"`
		ShowFastMerchant                            bool       `json:"showFastMerchant"`
		ShowWalletRebate                            bool       `json:"showWalletRebate"`
		ShowDigitalGoodsRebate                      bool       `json:"showDigitalGoodsRebate"`
		PudoBannerMoreInformationText               string     `json:"pudoBannerMoreInformationText"`
		PudoBannerEnabled                           bool       `json:"pudoBannerEnabled"`
		ShowTopCategoryRankingEnabled               bool       `json:"showTopCategoryRankingEnabled"`
		ShowFlashSales                              bool       `json:"showFlashSales"`
		FlashSalesTimeSlots                         []int      `json:"flashSalesTimeSlots"`
		PublicMdcSocialGwURL                        string     `json:"publicMdcSocialGwUrl"`
		PublicSdcSocialGwURL                        string     `json:"publicSdcSocialGwUrl"`
		MinInstallmentAmountEnabled                 bool       `json:"minInstallmentAmountEnabled"`
		MinInstallmentAmountMinPrice                int        `json:"minInstallmentAmountMinPrice"`
		MinInstallmentAmountMinInstallment          int        `json:"minInstallmentAmountMinInstallment"`
		ShowCheapSeller                             bool       `json:"showCheapSeller"`
		StickersEnabled                             bool       `json:"stickersEnabled"`
		ProductGroupSantralURL                      string     `json:"productGroupSantralUrl"`
		PublicWebSfxProductRecommendationServiceURL string     `json:"publicWebSfxProductRecommendationServiceUrl"`
		PublicMdcContractServiceURL                 string     `json:"publicMdcContractServiceUrl"`
		ThresholdDayForLongTermDelivery             int        `json:"thresholdDayForLongTermDelivery"`
		InfoForLongTermDelivery                     string     `json:"infoForLongTermDelivery"`
	} `json:"configuration"`
}
