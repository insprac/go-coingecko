package v3

type Coin struct {
	ID                           string                  `json:"id"`
	Symbol                       string                  `json:"symbol"`
	Name                         string                  `json:"name"`
	AssetPlatformID              string                  `json:"asset_platform_id"`
	BlockTimeInMinutes           uint16                  `json:"block_time_in_minutes"`
	HashingAlgorithm             string                  `json:"hashing_algorithm"`
	Categories                   []string                `json:"categories"`
	PublicNotice                 string                  `json:"public_notice"`
	AdditionalNotice             string                  `json:"additional_notice"`
	Localization                 map[string]string       `json:"localization"`
	Description                  map[string]string       `json:"description"`
	Links                        CoinLinks               `json:"links"`
	Image                        ImageLinks              `json:"image"`
	CountryOrigin                string                  `json:"country_origin"`
	GenesisDate                  string                  `json:"genesis_date"`
	SentimentVotesUpPercentage   float64                 `json:"sentiment_votes_up_percentage"`
	SentimentVotesDownPercentage float64                 `json:"sentiment_votes_down_percentage"`
	MarketCapRank                uint16                  `json:"market_cap_rank"`
	CoinGeckoRank                uint16                  `json:"coingecko_rank"`
	CoinGeckoScore               float64                 `json:"coingecko_score"`
	DeveloperScore               float64                 `json:"developer_score"`
	CommunityScore               float64                 `json:"community_score"`
	LiquidityScore               float64                 `json:"liquidity_score"`
	PublicInterestScore          float64                 `json:"public_interest_score"`
	MarketData                   CoinMarketData          `json:"market_data"`
	PublicInterestStats          CoinPublicInterestStats `json:"public_interest_stats"`
	StatusUpdates                []CoinStatusUpdate      `json:"status_updates"`
	LastUpdated                  string                  `json:"last_updated"`
	Tickers                      []CoinTicker            `json:"tickers"`
}

type CoinLinks struct {
	Homepage                    []string `json:"homepage"`
	BlockchainSite              []string `json:"blockchain_site"`
	OfficialForumURL            []string `json:"official_forum_url"`
	ChatURL                     []string `json:"chat_url"`
	AnnouncementURL             []string `json:"announcement_url"`
	TwitterScreenName           string   `json:"twitter_screen_name"`
	FacebookUsername            string   `json:"facebook_username"`
	BitcointalkThreadIdentifier string   `json:"facebook_username"`
	TelegramChannelIdentifier   string   `json:"telegram_channel_identifier"`
	SubredditURL                string   `json:"subreddit_url"`
	ReposURL                    RepoURLs `json:"repos_url"`
}

type RepoURLs struct {
	GitHub    []string `json:"github"`
	Bitbucket []string `json:"bitbucket"`
}

type ImageLinks struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type CoinMarketData struct {
	CurrentPrice                           map[string]float64 `json:"current_price"`
	ROI                                    ROI                `json:"roi"`
	ATH                                    map[string]float64 `json:"ath"`
	ATHChangePercentage                    map[string]float64 `json:"ath_change_percentage"`
	ATHDate                                map[string]string  `json:"ath_date"`
	ATL                                    map[string]float64 `json:"atl"`
	ATLChangePercentage                    map[string]float64 `json:"atl_change_percentage"`
	ATLDate                                map[string]string  `json:"atl_date"`
	MarketCapRank                          uint16             `json:"market_cap_rank"`
	FullyDilutedValuation                  map[string]uint64  `json:"fully_diluted_valuation"`
	High24h                                map[string]float64 `json:"high_24h"`
	Low24h                                 map[string]float64 `json:"low_24h"`
	PriceChange24h                         float64            `json:"price_change_24h"`
	PriceChangePercentage24h               float64            `json:"price_change_percentage_24h"`
	PriceChangePercentage7d                float64            `json:"price_change_percentage_7d"`
	PriceChangePercentage14d               float64            `json:"price_change_percentage_14d"`
	PriceChangePercentage30d               float64            `json:"price_change_percentage_30d"`
	PriceChangePercentage60d               float64            `json:"price_change_percentage_60d"`
	PriceChangePercentage200d              float64            `json:"price_change_percentage_200d"`
	PriceChangePercentage1y                float64            `json:"price_change_percentage_1y"`
	MarketCapChange24h                     float64            `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h           float64            `json:"market_cap_change_percentage_24h"`
	PriceChange24hInCurrency               map[string]float64 `json:"price_change_24h_in_currency"`
	PriceChangePercentage1hInCurrency      map[string]float64 `json:"price_change_percentage_1h_in_currency"`
	PriceChangePercentage24hInCurrency     map[string]float64 `json:"price_change_percentage_24h_in_currency"`
	PriceChangePercentage7dInCurrency      map[string]float64 `json:"price_change_percentage_7d_in_currency"`
	PriceChangePercentage14dInCurrency     map[string]float64 `json:"price_change_percentage_14d_in_currency"`
	PriceChangePercentage30dInCurrency     map[string]float64 `json:"price_change_percentage_30d_in_currency"`
	PriceChangePercentage60dInCurrency     map[string]float64 `json:"price_change_percentage_60d_in_currency"`
	PriceChangePercentage200dInCurrency    map[string]float64 `json:"price_change_percentage_200d_in_currency"`
	PriceChangePercentage1yInCurrency      map[string]float64 `json:"price_change_percentage_1y_in_currency"`
	MarketCapChange24hInCurrency           map[string]float64 `json:"market_cap_change_24h_in_currency"`
	MarketCapChangePercentage24hInCurrency map[string]float64 `json:"market_cap_change_percentage_24h_in_currency"`
	TotalSupply                            uint64             `json:"total_supply"`
	MaxSupply                              uint64             `json:"max_supply"`
	CirculatingSupply                      uint64             `json:"circulating_supply"`
	LastUpdated                            string             `json:"last_updated"`
}

type CoinCommunityData struct {
	FacebookLikes            uint64  `json:"facebook_likes"`
	TwitterFollowers         uint64  `json:"twitter_followers"`
	RedditAveragePosts48h    float64 `json:"reddit_average_posts_48h"`
	RedditAverageComments48h float64 `json:"reddit_average_comments_48h"`
	RedditSubscribers        uint64  `json:"reddit_subscribers"`
	RedditAccountsActive48h  uint64  `json:"reddit_accounts_active_48h"`
	TelegramChannelUserCount uint64  `json:"telegram_channel_user_count"`
}

type CoinDeveloperData struct {
	Forks                          uint32                              `json:"forks"`
	Stars                          uint32                              `json:"stars"`
	Subscribers                    uint32                              `json:"subscribers"`
	TotalIssues                    uint32                              `json:"total_issues"`
	ClosedIssues                   uint32                              `json:"closed_issues"`
	PullRequestsMerged             uint32                              `json:"pull_requests_merged"`
	PullRequestContributors        uint32                              `json:"pull_request_contributors"`
	CodeAdditionsDeletions4Weeks   CoinDeveloperDataDeletionsAdditions `json:"code_additions_deletions_4_weeks"`
	CommitCount4Weeks              uint32                              `json:"commit_count_4_weeks"`
	Last4WeeksCommitActivitySeries []uint16                            `json:"last_4_weeks_commit_activity_series"`
}

type CoinDeveloperDataDeletionsAdditions struct {
	Additions uint32 `json:"additions"`
	Deletions uint32 `json:"deletions"`
}

type CoinPublicInterestStats struct {
	AlexaRank   uint32 `json:"alexa_rank"`
	BingMatches uint32 `json:"bing_matches"`
}

type CoinStatusUpdate struct {
	Description string                  `json:"description"`
	Category    string                  `json:"category"`
	CreatedAt   string                  `json:"created_at"`
	User        string                  `json:"user"`
	UserTitle   string                  `json:"user_title"`
	Pin         bool                    `json:"pin"`
	Project     CoinStatusUpdateProject `json:"project"`
}

type CoinStatusUpdateProject struct {
	ID     string     `json:"id"`
	Symbol string     `json:"symbol"`
	Name   string     `json:"name"`
	Type   string     `json:"type"`
	Image  ImageLinks `json:"image"`
}

type CoinTicker struct {
	Base                   string             `json:"base"`
	Target                 string             `json:"target"`
	Market                 CoinTickerMarket   `json:"market"`
	Last                   float64            `json:"last"`
	Volume                 string             `json:"volume"`
	ConvertedLast          map[string]float64 `json:"converted_last"`
	TrustScore             string             `json:"trust_score"`
	BidAskSpreadPercentage float64            `json:"bid_ask_spread_percentage"`
	Timestamp              string             `json:"timestamp"`
	LastTradedAt           string             `json:"last_traded_at"`
	LastFetchAt            string             `json:"last_fetch_at"`
	IsAnomaly              bool               `json:"is_anomaly"`
	IsStale                bool               `json:"is_stale"`
	TradeURL               string             `json:"trade_url"`
	TokenInfoURL           string             `json:"token_info_url"`
	CoinID                 string             `json:"coin_id"`
	TargetCoinID           string             `json:"target_coin_id"`
}

type CoinTickerMarket struct {
	Name                string `json:"name"`
	Identifier          string `json:"identifier"`
	HasTradingIncentive bool   `json:"has_trading_incentive"`
}

type Market struct {
	ID                           string  `json:"id"`
	Symbol                       string  `json:"symbol"`
	Name                         string  `json:"name"`
	Image                        string  `json:"image"`
	CurrentPrice                 float64 `json:"current_price"`
	MarketCap                    uint64  `json:"market_cap"`
	MarketCapRank                uint16  `json:"market_cap_rank"`
	FullyDilutedValuation        uint64  `json:"fully_diluted_valuation"`
	TotalVolume                  uint64  `json:"total_volume"`
	High24h                      float64 `json:"high_24h"`
	Low24h                       float64 `json:"low_24h"`
	PriceChange24h               float64 `json:"price_change_24h"`
	PriceChangePercentage24h     float64 `json:"price_change_percentage_24h"`
	MarketCapChange24h           int16   `json:"market_cap_change_24h"`
	MarketCapChangePercentage24h float64 `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            uint64  `json:"circulating_supply"`
	TotalSupply                  uint64  `json:"total_supply"`
	MaxSupply                    uint64  `json:"max_supply"`
	ATH                          float64 `json:"ath"`
	ATHChangePercentage          float64 `json:"athChange_percentage"`
	ATHDate                      string  `json:"ath_date"`
	ATL                          float64 `json:"atl"`
	ATLChangePercentage          float64 `json:"atl_change_percentage"`
	ATLDate                      string  `json:"atl_date"`
	ROI                          ROI     `json:"roi"`
	LastUpdated                  string  `json:"last_updated"`
}

type ROI struct {
	Times      float64 `json:"times"`
	Currency   string  `json:"currency"`
	Percentage float64 `json:"percentage"`
}
