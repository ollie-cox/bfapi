package bfapi

//
func DefaultMarketSub(id string) SubsMessage {
	return SubsMessage{
		Op:                  MarketSub,
		ID:                  1,
		SegmentationEnabled: true, // <- watch this
		HeartbeatMs:         2000,
		ConflateMs:          0,
		MarketFilter: &MarketStreamFilter{
			MarketIds: []string{id},
		},
		MarketDataFilter: &MarketDataFilter{
			Fields: []string{
				"EX_ALL_OFFERS",
				"EX_TRADED",
				"EX_TRADED_VOL",
				"EX_MARKET_DEF",
				"SP_TRADED",
				"SP_PROJECTED",
				"EX_LTP",
			},
		},
	}
}

//
func DefaultOrderSub() SubsMessage {
	return SubsMessage{
		Op:                  OrderSub,
		SegmentationEnabled: false,
		HeartbeatMs:         2000,
		OrderFilter:         &OrderFilter{IncludeOverallPosition: true},
	}
}

//
func DefaultListMarketCatalogueArg(filter MarketListFilter) ListMarketCatalogueArg {
	return ListMarketCatalogueArg{
		Filter: filter,
		Sort:   "FIRST_TO_START",
		MarketProjection: []string{
			"MARKET_START_TIME",
			"RUNNER_DESCRIPTION",
			"EVENT",
			"MARKET_DESCRIPTION",
			"EVENT_TYPE"}, //, "RUNNER_METADATA", "EVENT_TYPE",
		MaxResults: 200,
	}
}
