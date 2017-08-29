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
	MarketIds  []string `json:"marketIds,omitempty"`
	EventTypes []string `json:"eventTypeIds,omitempty"`
	Countries  []string `json:"marketCountries,omitempty"`
	TypeCodes  []string `json:"marketTypeCodes,omitempty"`
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
	SelectionID  int               `json:"selectionId,omitempty"`
	RunnerName   string            `json:"runnerName,omitempty"`
	Handicap     float64           `json:"handicap,omitempty"`
	SortPriority int               `json:"sortPriority,omitempty"`
	Metadata     map[string]string `json:"metadata,omitempty"`
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
	NarketTime             time.Time              `json:"marketTime,omitempty"`
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
	BackStakeTaken    []ExchangeBet `json:"backStakeTaken,omitempty"`
	LayLiabilityTaken []ExchangeBet `json:"layLiabilityTaken,omitempty"`
	ActualSP          float64       `json:"actualSP,omitempty"`
}

//
type ExchangePrice struct {
	AvailableToBack []ExchangeBet `json:"availableToBack,omitempty"`
	AvailableToLay  []ExchangeBet `json:"availableToLay,omitempty"`
	TradedVolume    []ExchangeBet `json:"tradedVolume,omitempty"`
}

//
type ExchangeBet [2]float64

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
	SelectionID        int                 `json:"selectionId,omitempty"`
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