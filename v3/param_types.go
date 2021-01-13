package v3

type GetCoinMarketsParams struct {
	VSCurrency            string   `q:"vs_currency" required:"true"`
	IDs                   []string `q:"ids"`
	Category              string   `q:"category"`
	Order                 string   `q:"order"`
	PerPage               uint16   `q:"per_page"`
	Page                  uint16   `q:"page"`
	Sparkline             bool     `q:"sparkline"`
	PriceChangePercentage string   `q:"price_change_percentage"`
}

type GetCoinParams struct {
	Localization  string `q:"localization"`
	Tickers       bool   `q:"tickers"`
	MarketData    bool   `q:"market_data"`
	CommunityData bool   `q:"community_data"`
	DeveloperData bool   `q:"developer_data"`
	Sparkline     bool   `q:"sparkline"`
}

type GetCoinTickersParams struct {
	ExchangeIDs         string `q:"exchange_ids"`
	IncludeExchangeLogo bool   `q:"include_exchange_logo"`
	Page                uint16 `q:"page"`
	Order               string `q:"order"`
	Depth               string `q:"depth"`
}

type GetCoinHistoryParams struct {
	Date         string `q:"date" required:"true"`
	Localization bool   `q:"localization"`
}

type GetCoinMarketChartParams struct {
	VSCurrency string `q:"vs_currency" required:"true"`
	Days       uint16 `q:"days" required:"true"`
	Interval   string `q:"interval"`
}

type GetCoinMarketChartRangeParams struct {
	VSCurrency string `q:"vs_currency" required:"true"`
	From       int64  `q:"from" required:"true"`
	To         int64  `q:"to" required:"true"`
}

type GetCoinStatusUpdatesParams struct {
	PerPage uint16 `q:"per_page"`
	Page    uint16 `q:"page"`
}
