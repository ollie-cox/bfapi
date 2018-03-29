package bfapi

import (
	"time"
)

//
type ListMarketCatalogueArg struct {
	Sort             string           `json:"sort,omitempty"`
	MarketProjection []string         `json:"marketProjection,omitempty"`
	MaxResults       int              `json:"maxResults,omitempty"`
	Filter           MarketListFilter `json:"filter,omitempty"`
}

//
type MarketListFilter struct {
	MarketIds      []string `json:"marketIds,omitempty"`
	EventTypes     []string `json:"eventTypeIds,omitempty"`
	Countries      []string `json:"marketCountries,omitempty"`
	TypeCodes      []string `json:"marketTypeCodes,omitempty"`
	CompetitionIds []string `json:"competitionIds,omitempty"`
}

//
type MarketCatalogue struct {
	MarketID        string             `json:"marketId,omitempty"`
	MarketName      string             `json:"marketName,omitempty"`
	MarketStartTime time.Time          `json:"marketStartTime,omitempty"`
	Description     *MarketDescription `json:"marketDescription,omitempty"`
	TotalMatched    float64            `json:"totalMatched,omitempty"`
	Runners         []RunnerCatalogue  `json:"runners,omitempty"`
	EventType       *EventType         `json:"eventType,omitempty"`
	Competition     *Competition       `json:"competition,omitempty"`
	Event           *Event             `json:"event,omitempty"`
}

//
type RunnerCatalogue struct {
	SelectionID  int64           `json:"selectionId,omitempty"`
	RunnerName   string          `json:"runnerName,omitempty"`
	Handicap     float64         `json:"handicap,omitempty"`
	SortPriority int             `json:"sortPriority,omitempty"`
	Metadata     *RunnerMetadata `json:"metadata,omitempty"`
}

//
type RunnerMetadata struct {
	SireName                 *string `json:"SIRE_NAME,omitempty"`
	ClothNumberAlpha         *string `json:"CLOTH_NUMBER_ALPHA,omitempty"`
	OfficialRating           *string `json:"OFFICIAL_RATING,omitempty"`
	ColoursDescription       *string `json:"COLOURS_DESCRIPTION,omitempty"`
	ColoursFileName          *string `json:"COLOURS_FILENAME,omitempty"`
	ForecastPriceDenominator *string `json:"FORECASTPRICE_DENOMINATOR,omitempty"`
	DamSireName              *string `json:"DAMSIRE_NAME,omitempty"`
	WeightValue              *string `json:"WEIGHT_VALUE,omitempty"`
	SexType                  *string `json:"SEX_TYPE,omitempty"`
	DaysSinceLastRun         *string `json:"DAYS_SINCE_LAST_RUN,omitempty"`
	Wearing                  *string `json:"WEARING,omitempty"`
	OwnerName                *string `json:"OWNER_NAME,omitempty"`
	DamYearBorn              *string `json:"DAM_YEAR_BORN,omitempty"`
	SireBred                 *string `json:"SIRE_BRED,omitempty"`
	JockeyName               *string `json:"JOCKEY_NAME,omitempty"`
	DamBred                  *string `json:"DAM_BRED,omitempty"`
	AdjustedRating           *string `json:"ADJUSTED_RATING,omitempty"`
	RunnerID                 *string `json:"runnerId,omitempty"`
	ClothNumber              *string `json:"CLOTH_NUMBER,omitempty"`
	SireYearBorn             *string `json:"SIRE_YEAR_BORN,omitempty"`
	TrainerName              *string `json:"TRAINER_NAME,omitempty"`
	ColourType               *string `json:"COLOUR_TYPE,omitempty"`
	Age                      *string `json:"AGE,omitempty"`
	DamsireBred              *string `json:"DAMSIRE_BRED,omitempty"`
	JockeyClaim              *string `json:"JOCKEY_CLAIM,omitempty"`
	Form                     *string `json:"FORM,omitempty"`
	ForecastPriceNumerator   *string `json:"FORECASTPRICE_NUMERATOR,omitempty"`
	Bred                     *string `json:"BRED,omitempty"`
	DamName                  *string `json:"DAM_NAME,omitempty"`
	DamSireYearBorn          *string `json:"DAMSIRE_YEAR_BORN,omitempty"`
	StallDraw                *string `json:"STALL_DRAW,omitempty"`
	WeightUnits              *string `json:"WEIGHT_UNITS,omitempty"`
}

//
type Event struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	CountryCode string    `json:"countryCode,omitempty"`
	Timezone    string    `json:"timezone,omitempty"`
	Venue       string    `json:"venue,omitempty"`
	OpenDate    time.Time `json:"openDate,omitempty"`
}

//
type EventType struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//
type Competition struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

//
type MarketDescription struct {
	PersistenceEnabled     bool                   `json:"persistenceEnabled,omitempty"`
	BspMarket              bool                   `json:"bspMarket,omitempty"`
	MarketTime             time.Time              `json:"marketTime,omitempty"`
	SuspendTime            time.Time              `json:"suspendTime,omitempty"`
	SettleTime             time.Time              `json:"settleTime,omitempty"`
	BettingType            string                 `json:"bettingType,omitempty"` //ODDS/LINE/RANGE/ASIAN_HANDICAP_DOUBLE_LINE/ASIAN_HANDICAP_SINGLE_LINE/FIXED_ODDS
	TurnInPlayEnabled      bool                   `json:"turnInPlayEnabled,omitempty"`
	MarketType             string                 `json:"marketType,omitempty"`
	Regulator              string                 `json:"regulator,omitempty"`
	MarketBaseRate         float64                `json:"marketBaseRate,omitempty"`
	DiscountAllowed        bool                   `json:"discountAllowed,omitempty"`
	Wallet                 string                 `json:"wallet,omitempty"`
	Rules                  string                 `json:"rules,omitempty"`
	RulesHasDate           bool                   `json:"rulesHasDate,omitempty"`
	EachWayDivisor         float64                `json:"eachWayDivisor,omitempty"`
	Clarifications         string                 `json:"clarifications,omitempty"`
	LineRangeInfo          MarketLineRangeInfo    `json:"lineRangeInfo,omitempty"`
	PriceLadderDescription PriceLadderDescription `json:"priceLadderDescription,omitempty"`
}

//
type MarketLineRangeInfo struct {
	MaxUnitValue float64 `json:"maxUnitValue,omitempty"`
	MinUnitValue float64 `json:"minUnitValue,omitempty"`
	Interval     float64 `json:"interval,omitempty"`
	MarketUnit   string  `json:"marketUnit,omitempty"`
}

//
type PriceLadderDescription struct {
	Type string `json:"type,omitempty"` //CLASSIC/FINEST/LINE_RANGE
}

//
type StartingPrice struct {
	BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
	ActualSP          float64     `json:"actualSP,omitempty"`
}

//
type ExchangePrice struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
}

//
type PriceProjection struct {
	PriceData             []string              `json:"priceData,omitempty"` //SP_AVAILABLE,SP_TRADED,EX_BEST_OFFERS,EX_ALL_OFFERS,EX_ALL_OFFERS,EX_TRADED
	ExBestOffersOverrides ExBestOffersOverrides `json:"exBestOffersOverrides,omitempty"`
	Virtualise            bool                  `json:"virtualise,omitempty"`
	RolloverStakes        bool                  `json:"rolloverStakes,omitempty"`
}

//
type ExBestOffersOverrides struct {
	BestPricesDepth          int     `json:"bestPricesDepth,omitempty"`
	RollupModel              string  `json:"rollupModel,omitempty"` //STAKE,PAYOUT,MANAGED_LIABILITY,NONE
	RollupLimit              int     `json:"rollupLimit,omitempty"`
	RollupLiabilityThreshold float64 `json:"rollupLiabilityThreshold,omitempty"`
	RollupLiabilityFactor    int     `json:"rollupLiabilityFactor,omitempty"`
}

//
type ListMarketBookArg struct {
	MarketIds                     []string        `json:"marketIds,omitempty"` //REQUIRED
	PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
	OrderProjection               string          `json:"orderProjection,omitempty"` //ALL,EXECUTABLE,EXECUTION_COMPLETE
	MatchProjection               string          `json:"matchProjection,omitempty"` //NO_ROLLUP,ROLLED_UP_BY_PRICE,ROLLED_UP_BY_AVG_PRICE
	IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
	PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
	CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
	CurrencyCode                  string          `json:"currencyCode,omitempty"`
	Locale                        string          `json:"locale,omitempty"`
	MatchedSince                  time.Time       `json:"matchedSince,omitempty"`
	BetIds                        []string        `json:"betIds,omitempty"`
}

//
type MarketBook struct {
	MarketID              string    `json:"marketId,omitempty"`
	IsMarketDataDelayed   bool      `json:"isMarketDataDelayed,omitempty"`
	Status                string    `json:"status,omitempty"` //INACTIVE,OPEN,SUSPENDED,CLOSED
	BetDelay              int       `json:"betDelay,omitempty"`
	BspReconciled         bool      `json:"bspReconciled,omitempty"`
	Complete              bool      `json:"complete,omitempty"`
	Inplay                bool      `json:"inplay,omitempty"`
	NumberOfWinners       int       `json:"numberOfWinners,omitempty"`
	NumberOfRunners       int       `json:"numberOfRunners,omitempty"`
	NumberOfActiveRunners int       `json:"numberOfActiveRunners,omitempty"`
	LastMatchTime         time.Time `json:"lastMatchTime,omitempty"`
	TotalMatched          float64   `json:"totalMatched,omitempty"`
	TotalAvailable        float64   `json:"totalAvailable,omitempty"`
	CrossMatching         bool      `json:"crossMatching,omitempty"`
	RunnersVoidable       bool      `json:"runnersVoidable,omitempty"`
	Version               int64     `json:"version,omitempty"`
	Runners               []Runner  `json:"runners,omitempty"`
	KeyLineDescription    string    `json:"keyLineDescription,omitempty"` //?
}

//
type Runner struct {
	SelectionID      int64          `json:"selectionId,omitempty"`
	Handicap         float64        `json:"handicap,omitempty"`
	Status           string         `json:"status,omitempty"` //ACTIVE, REMOVED, WINNER, PLACED, LOSER, HIDDEN
	AdjustmentFactor float64        `json:"adjustmentFactor,omitempty"`
	LastPriceTraded  float64        `json:"lastPriceTraded,omitempty"`
	TotalMatched     float64        `json:"totalMatched,omitempty"`
	RemovalDate      time.Time      `json:"removalDate,omitempty"`
	Sp               StartingPrices `json:"sp,omitempty"`
	Ex               ExchangePrices `json:"ex,omitempty"`
	// orders            Order            `json:"orders,omitempty"`
	// matches           Match      `json:"matches,omitempty"`
	// matchesByStrategy map[string]Matches        `json:"matchesByStrategy,omitempty"`
}

//
type StartingPrices struct {
	NearPrice         float64     `json:"nearPrice,omitempty"`
	FarPrice          float64     `json:"farPrice,omitempty"`
	BackStakeTaken    []PriceSize `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []PriceSize `json:"layLiabilityTaken,omitempty"`
	ActualSP          float64     `json:"actualSP,omitempty"`
}

//
type ExchangePrices struct {
	AvailableToBack []PriceSize `json:"availableToBack,omitempty"`
	AvailableToLay  []PriceSize `json:"availableToLay,omitempty"`
	TradedVolume    []PriceSize `json:"tradedVolume,omitempty"`
}

//
type PriceSize struct {
	Price float64 `json:"price,omitempty"`
	Size  float64 `json:"size,omitempty"`
}

//
type PlaceOrderArg struct {
	MarketID            string             `json:"marketId,omitempty"`
	Instructions        []PlaceInstruction `json:"instructions,omitempty"`
	CustomerRef         string             `json:"customerRef,omitempty"`
	MarketVersion       int                `json:"marketVersion,omitempty"`
	CustomerStrategyRef string             `json:"customerStrategyRef,omitempty"`
	Async               bool               `json:"async,omitempty"`
}

//
type PlaceInstruction struct {
	OrderType          string              `json:"orderType,omitempty"`
	SelectionID        int64               `json:"selectionId,omitempty"`
	Handicap           string              `json:"handicap,omitempty"`
	Side               string              `json:"side,omitempty"`
	LimitOrder         *LimitOrder         `json:"limitOrder,omitempty"`
	LimitOnCloseOrder  *LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`
	MarketOnCloseOrder *MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"`
	CustomerOrderRef   string              `json:"customerOrderRef,omitempty"`
}

//
type PlaceOrderRequest struct {
	MarketID            string             `json:"marketId,omitempty"`
	Instructions        []PlaceInstruction `json:"instructions,omitempty"`
	CustomerRef         string             `json:"customerRef,omitempty"`
	MarketVersion       MarketVersion      `json:"marketVersion,omitempty"`
	CustomerStrategyRef string             `json:"customerStrategyRef,omitempty"`
	Async               bool               `json:"async,omitempty"`
}

//
type PlaceExecutionReport struct {
	CustomerRef        string                   `json:"customerRef,omitempty"`
	MarketID           string                   `json:"marketId,omitempty"`
	Status             string                   `json:"status,omitempty"`
	ErrorCode          string                   `json:"errorCode,omitempty"`
	InstructionReports []PlaceInstructionReport `json:"instructionReports,omitempty"`
}

//
type PlaceInstructionReport struct {
	Status              string           `json:"status,omitempty"`
	ErrorCode           string           `json:"errorCode,omitempty"`
	OrderStatus         string           `json:"orderStatus,omitempty"`
	Instruction         PlaceInstruction `json:"instruction,omitempty"`
	BetID               string           `json:"betId,omitempty"`
	PlacedDate          string           `json:"placedDate,omitempty"`
	AveragePriceMatched int              `json:"averagePriceMatched,omitempty"`
	SizeMatched         int              `json:"sizeMatched,omitempty"`
}

//
type LimitOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
	Price     float64 `json:"price,omitempty"`
}

//
type MarketOnCloseOrder struct {
	Liability float64 `json:"liability,omitempty"`
}

//
type LimitOrder struct {
	Size            float64 `json:"size,omitempty"`
	Price           float64 `json:"price,omitempty"`
	PersistenceType string  `json:"persistenceType,omitempty"`
}

//
type MarketVersion struct {
	Version float64 `json:"version,omitempty"`
}

//
type CancelInstruction struct {
	BetID         string  `json:"betId,omitempty"`
	SizeReduction float64 `json:"sizeReduction,omitempty"`
}

//
type CancelOrderRequest struct {
	MarketID     string              `json:"marketId,omitempty"`
	Instructions []CancelInstruction `json:"instructions,omitempty"`
	CustomerRef  string              `json:"customerRef,omitempty"`
}

//
type CancelInstructionReport struct {
	Status        string            `json:"status,omitempty"`
	ErrorCode     string            `json:"errorCode,omitempty"`
	Instruction   CancelInstruction `json:"instruction,omitempty"`
	SizeCancelled float64           `json:"sizeCancelled,omitempty"`
	CancelDate    string            `json:"cancelledDate,omitempty"`
}

//
type CancelExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             string                    `json:"status,omitempty"`
	ErrorCode          string                    `json:"errorCode,omitempty"`
	MarketID           string                    `json:"marketId,omitempty"`
	InstructionReports []CancelInstructionReport `json:"instructionReports,omitempty"`
}

//
type ReplaceInstruction struct {
	BetID    string  `json:"betId,omitempty"`
	NewPrice float64 `json:"newPrice,omitempty"`
}

//
type ReplaceOrderRequest struct {
	MarketID      string               `json:"marketId,omitempty"`
	Instructions  []ReplaceInstruction `json:"instructions,omitempty"`
	CustomerRef   string               `json:"customerRef,omitempty"`
	MarketVersion MarketVersion        `json:"marketVersion,omitempty"`
	Async         bool                 `json:"async,omitempty"`
}

//
type ReplaceInstructionReport struct {
	Status                  string                  `json:"status,omitempty"`
	ErrorCode               string                  `json:"errorCode,omitempty"`
	CancelInstructionReport CancelInstructionReport `json:"cancelInstructionReport,omitempty"`
	PlaceInstructionReport  PlaceInstructionReport  `json:"placeInstructionReport,omitempty"`
}

//
type ReplaceExecutionReport struct {
	CustomerRef        string                     `json:"customerRef,omitempty"`
	Status             string                     `json:"status,omitempty"`
	ErrorCode          string                     `json:"errorCode,omitempty"`
	MarketID           string                     `json:"marketId,omitempty"`
	InstructionReports []ReplaceInstructionReport `json:"instructionReports,omitempty"`
}

//
type UpdateInstruction struct {
	BetID              string `json:"betId,omitempty"`
	NewPersistenceType string `json:"newPersistenceType,omitempty"`
}

//
type UpdateOrderRequest struct {
	MarketID     string              `json:"marketId,omitempty"`
	Instructions []UpdateInstruction `json:"instructions,omitempty"`
	CustomerRef  string              `json:"customerRef,omitempty"`
}

//
type UpdateInstructionReport struct {
	Status      string            `json:"status,omitempty"`
	ErrorCode   string            `json:"errorCode,omitempty"`
	Instruction UpdateInstruction `json:"instruction,omitempty"`
}

//
type UpdateExecutionReport struct {
	CustomerRef        string                    `json:"customerRef,omitempty"`
	Status             string                    `json:"status,omitempty"`
	ErrorCode          string                    `json:"errorCode,omitempty"`
	MarketID           string                    `json:"marketId,omitempty"`
	InstructionReports []UpdateInstructionReport `json:"instructionReports,omitempty"`
}

//
type ListClearedOrdersArgs struct {
	BetStatus              string     `json:"betStatus,omitempty"` //SETTLED | VOIDED | LAPSED | CANCELLED
	EventTypeIDs           []string   `json:"eventTypeIds,omitempty"`
	EventIDs               []string   `json:"eventIds,omitempty"`
	MarketIDs              []string   `json:"marketIds,omitempty"`
	RunnerIDs              []int64    `json:"runnerIds,omitempty"`
	BetIDs                 []string   `json:"betIds,omitempty"`
	CustomerOrderRefs      []string   `json:"customerOrderRefs,omitempty"`
	CustomerStrategyRefs   []string   `json:"customerStrategyRefs,omitempty"`
	Side                   string     `json:"side,omitempty"` // BACK | LAY
	SettledDateRange       *TimeRange `json:"settledDateRange,omitempty"`
	GroupBy                string     `json:"groupBy,omitempty"` // EVENT_TYPE | EVENT | MARKET | SIDE | BET
	IncludeItemDescription bool       `json:"includeItemDescription,omitempty"`
	Locale                 string     `json:"locale,omitempty"`
	FromRecord             int        `json:"fromRecord,omitempty"`
	RecordCount            int        `json:"recordCount,omitempty"`
}

//
type ClearedOrderSummaryReport struct {
	ClearedOrders []ClearedOrderSummary `json:"clearedOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

//
type ClearedOrderSummary struct {
	EventTypeID         string           `json:"eventTypeId"`
	EventID             string           `json:"eventId"`
	MarketID            string           `json:"marketId"`
	SelectionID         int64            `json:"selectionId"`
	Handicap            float64          `json:"handicap"`
	BetID               string           `json:"betId"`
	PlacedDate          time.Time        `json:"placedDate"`
	PersistenceType     string           `json:"persistenceType"` // LAPSE | PERSIST | MARKET_ON_CLOSE
	OrderType           string           `json:"orderType"`       // LIMIT | LIMIT_ON_CLOSE | MARKET_ON_CLOSE
	Side                string           `json:"side"`            // BACK | LAY
	ItemDescription     *ItemDescription `json:"itemDescription"`
	BetOutcome          string           `json:"betOutcome"`
	PriceRequested      float64          `json:"priceRequested"`
	SettledDate         time.Time        `json:"settledDate"`
	LastMatchedDate     time.Time        `json:"lastMatchedDate"`
	BetCount            int              `json:"betCount"`
	Commission          float64          `json:"commission"`
	PriceMatched        float64          `json:"priceMatched"`
	PriceReduced        bool             `json:"priceReduced"`
	SizeSettled         float64          `json:"sizeSettled"`
	Profit              float64          `json:"profit"`
	SizeCancelled       float64          `json:"sizeCancelled"`
	CustomerOrderRef    string           `json:"customerOrderRef"`
	CustomerStrategyRef string           `json:"customerStrategyRef"`
}

//
type ListCurrentOrdersArgs struct {
	BetIDs               []string   `json:"betIds,omitempty"`
	MarketIDs            []string   `json:"marketIds,omitempty"`
	OrderProjection      string     `json:"orderProjection,omitempty"` // ALL, EXECUTABLE, EXECUTION_COMPLETE
	CustomerOrderRefs    []string   `json:"customerOrderRefs,omitempty"`
	CustomerStrategyRefs []string   `json:"customerStrategyRefs,omitempty"`
	DateRange            *TimeRange `json:"dateRange,omitempty"`
	OrderBy              string     `json:"orderBy,omitempty"` // BY_MARKET,BY_MATCH_TIME,BY_PLACE_TIME,BY_SETTLED_TIME,BY_VOID_TIME
	SortDir              string     `json:"sortDir,omitempty"` //EARLIEST_TO_LATEST, LATEST_TO_EARLIEST
	FromRecord           int        `json:"fromRecord,omitempty"`
	RecordCount          int        `json:"recordCount,omitempty"`
}

//
type CurrentOrderSummaryReport struct {
	CurrentOrders []CurrentOrderSummary `json:"currentOrders"`
	MoreAvailable bool                  `json:"moreAvailable"`
}

//
type CurrentOrderSummary struct {
	BetID               string    `json:"betId"`
	MarketID            string    `json:"marketId"`
	SelectionID         int64     `json:"selectionId"`
	Handicap            float64   `json:"handicap"`
	PriceSize           PriceSize `json:"priceSize"`
	BspLiability        float64   `json:"bspLiability"`
	Side                string    `json:"side"`            // BACK | LAY
	Status              string    `json:"status"`          // EXECUTABLE,EXECUTION_COMPLETE
	PersistenceType     string    `json:"persistenceType"` // LAPSE | PERSIST | MARKET_ON_CLOSE
	OrderType           string    `json:"orderType"`       // LIMIT | LIMIT_ON_CLOSE | MARKET_ON_CLOSE
	PlacedDate          time.Time `json:"placedDate"`
	MatchedDate         time.Time `json:"matchedDate"`
	AveragePriceMatched float64   `json:"averagePriceMatched"`
	SizeMatched         float64   `json:"sizeMatched"`
	SizeRemaining       float64   `json:"sizeRemaining"`
	SizeLapsed          float64   `json:"sizeLapsed"`
	SizeCancelled       float64   `json:"sizeCancelled"`
	SizeVoided          float64   `json:"sizeVoided"`
	RegulatorAuthCode   string    `json:"regulatorAuthCode"`
	RegulatorCode       string    `json:"regulatorCode"`
	CustomerOrderRef    string    `json:"customerOrderRef"`
	CustomerStrategyRef string    `json:"customerStrategyRef"`
}

//
type ItemDescription struct {
	EventTypeDesc   string    `json:"eventTypeDesc"`
	EventDesc       string    `json:"eventDesc"`
	MarketDesc      string    `json:"marketDesc"`
	MarketType      string    `json:"marketType"`
	MarketStartTime time.Time `json:"marketStartTime"`
	RunnerDesc      string    `json:"runnerDesc"`
	NumberOfWinners int       `json:"numberOfWinners"`
	EachWayDivisor  float64   `json:"eachWayDivisor"`
}

//
type TimeRange struct {
	To   time.Time `json:"to"`
	From time.Time `json:"from"`
}
