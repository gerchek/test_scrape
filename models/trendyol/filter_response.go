package trendyol

type TrendyolFilterResponse struct {
	IsSuccess  bool        `json:"isSuccess"`
	StatusCode int         `json:"statusCode"`
	Error      interface{} `json:"error"`
	Result     struct {
		SelectedFilters []struct {
			ID             string `json:"id"`
			FilterField    string `json:"filterField"`
			Type           string `json:"type"`
			IsVisible      bool   `json:"isVisible"`
			Text           string `json:"text"`
			BeautifiedName string `json:"beautifiedName"`
			URL            string `json:"url"`
			Source         int    `json:"source"`
		} `json:"selectedFilters"`
		ResolvedQuery struct {
			Filters []struct {
				FilterKey   string   `json:"filterKey"`
				FilterType  string   `json:"filterType"`
				Type        string   `json:"type"`
				FilterField string   `json:"filterField"`
				Values      []string `json:"values"`
				Filters     []struct {
					Key            string `json:"key"`
					BeautifiedName string `json:"beautifiedName"`
					Name           string `json:"name"`
				} `json:"filters"`
			} `json:"filters"`
			SearchKey        string        `json:"searchKey"`
			AlternateFilters []interface{} `json:"alternateFilters"`
		} `json:"resolvedQuery"`
		ClearAllFiltersURL string `json:"clearAllFiltersUrl"`
		Aggregations       []struct {
			Group      string `json:"group"`
			Type       string `json:"type"`
			Title      string `json:"title"`
			FilterKey  string `json:"filterKey"`
			Order      int    `json:"order"`
			TotalCount int    `json:"totalCount"`
			Values     []struct {
				ID             string `json:"id"`
				Text           string `json:"text"`
				BeautifiedName string `json:"beautifiedName"`
				Count          int    `json:"count"`
				Filtered       bool   `json:"filtered"`
				FilterField    string `json:"filterField"`
				Type           string `json:"type"`
				URL            string `json:"url"`
			} `json:"values"`
			ID          string `json:"id"`
			FilterType  string `json:"filterType"`
			FilterField string `json:"filterField"`
		} `json:"aggregations"`
	} `json:"result"`
	Headers struct {
		Tysidecarcachable string `json:"tysidecarcachable"`
	} `json:"headers"`
}

type Item struct {
	ID             string `json:"id"`
	Text           string `json:"text"`
	BeautifiedName string `json:"beautifiedName"`
	Type           string `json:"type"`
	URL            string `json:"url"`
}

type SortedItems struct {
	Brands     []Item `json:"brands"`
	Categories []Item `json:"categories"`
}
