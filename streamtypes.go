package bfapi

import (
	"time"
)

//RequestMessages OP codes
const (
	Auth      string = "authentication"
	MarketSub string = "marketSubscription"
	OrderSub  string = "orderSubscription"
	Heartbeat string = "heartbeat"
)

//ResponseMessages OP codes
const (
	Conn   string = "connection"
	Status string = "status"
	Mcm    string = "mcm"
	Ocm    string = "ocm"
)

//
type AuthMessage struct {
	Op      string `json:"op,omitempty"`
	ID      int32  `json:"id,omitempty"`
	Session string `json:"session,omitempty"`
	AppKey  string `json:"appKey,omitempty"`
}

//
type HeartbeatMessage struct {
	Op string `json:"op,omitempty"`
	ID int32  `json:"id,omitempty"`
}

//
type SubsMessage struct {
	Op                  string              `json:"op,omitempty"`
	ID                  int32               `json:"id,omitempty"`
	SegmentationEnabled bool                `json:"segmentationEnabled,omitempty"`
	Clk                 string              `json:"clk,omitempty"`
	HeartbeatMs         int64               `json:"heartbeatMs,omitempty"`
	InitialClk          string              `json:"initialClk,omitempty"`
	ConflateMs          int64               `json:"conflateMs"`
	MarketFilter        *MarketStreamFilter `json:"marketFilter,omitempty"`
	MarketDataFilter    *MarketDataFilter   `json:"marketDataFilter,omitempty"`
	OrderFilter         *OrderFilter        `json:"orderFilter,omitempty"`
}

//
type MarketDataFilter struct {
	LadderLevels int32    `json:"ladderLevels,omitempty"`
	Fields       []string `json:"fields,omitempty"`
}

//
type MarketStreamFilter struct {
	MarketIds         []string `json:"marketIds,omitempty"`
	CountryCodes      []string `json:"countryCodes,omitempty"`
	BettingTypes      []string `json:"bettingTypes,omitempty"`
	TurnInPlayEnabled bool     `json:"turnInPlayEnabled,omitempty"`
	MarketTypes       []string `json:"marketTypes,omitempty"`
	Venues            []string `json:"venues,omitempty"`
	EventTypeIds      []string `json:"eventTypeIds,omitempty"`
	EventIds          []string `json:"eventIds,omitempty"`
	BspMarket         bool     `json:"bspMarket,omitempty"`
}

//
type OrderFilter struct {
	IncludeOverallPosition        bool     `json:"includeOverallPosition,omitempty"`
	AccountIds                    []int64  `json:"accountIds,omitempty"`
	CustomerStrategyRefs          []string `json:"customerStrategyRefs,omitempty"`
	PartitionMatchedByStrategyRef bool     `json:"partitionMatchedByStrategyRef,omitempty"`
}

//
type ConnectionMessage struct {
	Op           string `json:"op,omitempty"`
	ID           int    `json:"id,omitempty"`
	ConnectionID string `json:"connectionId,omitempty"`
}

//
type StatusMessage struct {
	Op               string `json:"op,omitempty"`
	ID               int    `json:"id,omitempty"`
	ErrorMessage     string `json:"errorMessage,omitempty"`
	ErrorCode        string `json:"errorCode,omitempty"`
	ConnectionID     string `json:"connectionId,omitempty"`
	ConnectionClosed bool   `json:"connectionClosed,omitempty"`
	StatusCode       string `json:"statusCode,omitempty"`
}

//
type ChangeMessage struct {
	Op          string         `json:"op,omitempty"`
	ID          int            `json:"id,omitempty"`
	Ct          string         `json:"ct,omitempty"`
	Clk         string         `json:"clk,omitempty"`
	HeartbeatMs int64          `json:"heartbeatMs,omitempty"`
	Pt          int64          `json:"pt,omitempty"`
	InitialClk  string         `json:"initialClk,omitempty"`
	ConflateMs  int64          `json:"conflateMs,omitempty"`
	SegmentType string         `json:"segmentType,omitempty"`
	Mc          []MarketChange `json:"mc,omitempty"`
	Oc          []OrderChange  `json:"oc,omitempty"`
}

//
type MarketChange struct {
	ID               string            `json:"id"`
	Rc               []RunnerChange    `json:"rc"`
	MarketDefinition *MarketDefinition `json:"marketDefinition"`
	Img              bool              `json:"img"`
	Con              bool              `json:"con"`
	Tv               float64           `json:"tv"`  //total volume
	Spn              float64           `json:"spn"` //starting price projection prices
	Spf              float64           `json:"sbf"` //starting price projection prices
	Ltp              float64           `json:"ltp"` //last traded price
}

//
type RunnerChange struct {
	ID    int64        `json:"id"`    //RunnerId
	Spn   float64      `json:"spn"`   //predicted sp
	Ltp   float64      `json:"ltp"`   //last traded price
	Tv    float64      `json:"tv"`    //traded volume
	Bdatb [][3]float64 `json:"bdatb"` //available to back
	Bdatl [][3]float64 `json:"bdatl"` //available to lay
	Batb  [][3]float64 `json:"batb"`  //available to back
	Batl  [][3]float64 `json:"batl"`  //available to lay
	Atb   [][2]float64 `json:"atb"`   //available to back
	Atl   [][2]float64 `json:"atl"`   //available to lay
	Trd   [][2]float64 `json:"trd"`   //traded bets
	Spb   [][2]float64 `json:"spb"`   //starting price bets
	Spl   [][2]float64 `json:"spl"`   //starting price lays
}

//
type MarketDefinition struct {
	MarketName            string             `json:"name,omitempty"`
	EventName             string             `json:"eventName,omitempty"`
	Status                string             `json:"status,omitempty"`
	Venue                 string             `json:"venue,omitempty"`
	SettledTime           time.Time          `json:"settledTime,omitempty"`
	Timezone              string             `json:"timezone,omitempty"`
	EachWayDivisor        float64            `json:"eachWayDivisor,omitempty"`
	Regulators            []string           `json:"regulators,omitempty"`
	MarketType            string             `json:"marketType,omitempty"`
	MarketBaseRate        float64            `json:"marketBaseRate,omitempty"`
	NumberOfWinners       int32              `json:"numberOfWinners,omitempty"`
	CountryCode           string             `json:"countryCode,omitempty"`
	InPlay                bool               `json:"inPlay,omitempty"`
	BetDelay              int32              `json:"betDelay,omitempty"`
	BspMarket             bool               `json:"bspMarket,omitempty"`
	BettingType           string             `json:"bettingType,omitempty"`
	NumberOfActiveRunners int32              `json:"numberOfActiveRunners,omitempty"`
	EventID               string             `json:"eventId,omitempty"`
	CrossMatching         bool               `json:"crossMatching,omitempty"`
	RunnersVoidable       bool               `json:"runnersVoidable,omitempty"`
	TurnInPlayEnabled     bool               `json:"turnInPlayEnabled,omitempty"`
	SuspendTime           time.Time          `json:"suspendTime,omitempty"`
	DiscountAllowed       bool               `json:"discountAllowed,omitempty"`
	PersistenceEnabled    bool               `json:"persistenceEnabled,omitempty"`
	Runners               []RunnerDefinition `json:"runners,omitempty"`
	Version               int64              `json:"version,omitempty"`
	EventTypeID           string             `json:"eventTypeId,omitempty"`
	Complete              bool               `json:"complete,omitempty"`
	OpenDate              time.Time          `json:"openDate,omitempty"`
	MarketTime            time.Time          `json:"marketTime,omitempty"`
	BspReconciled         bool               `json:"bspReconciled,omitempty"`
}

//
type RunnerDefinition struct {
	Status           string    `json:"status,omitempty"`
	SortPriority     int       `json:"sortPriority,omitempty"`
	Sp               float64   `json:"bsp,omitempty"`
	SelectionID      int64     `json:"id,omitempty"`
	SelectionName    string    `json:"name,omitempty"`
	AdjustmentFactor float64   `json:"adjustmentFactor,omitempty"`
	RemovalDate      time.Time `json:"removalDate,omitempty"`
	Handicap         float64   `json:"hc,omitempty"`
}

//
type OrderChange struct {
	MarketID string        `json:"id"`
	Runners  []RunnerOrder `json:"orc"`
}

//
type RunnerOrder struct {
	Image        bool             `json:"fullImage"`
	SelectionID  int64            `json:"id"`
	Unmatched    []UnmatchedOrder `json:"uo"`
	MatchedBacks [][2]float64     `json:"mb"`
	MatchedLays  [][2]float64     `json:"ml"`
}

//
type UnmatchedOrder struct {
	ID           string  `json:"id"`
	Ref          string  `json:"rfo"`
	Price        float64 `json:"p"`
	Size         float64 `json:"s"`
	Side         string  `json:"side"`
	Status       string  `json:"status"` //(E = EXECUTABLE, EC = EXECUTION_COMPLETE)
	PerType      string  `json:"pt"`     //(L = LAPSE, P = PERSIST, MOC = Market On Close)
	OrderType    string  `json:"ot"`     //(L = LIMIT, MOC = MARKET_ON_CLOSE, LOC = LIMIT_ON_CLOSE)
	PlacedAt     int64   `json:"pd"`
	MatchedAt    int64   `json:"md"`
	LapsedAt     int64   `json:"ld"`
	AvgPrice     float64 `json:"avp"`
	SizeMatched  float64 `json:"sm"`
	SizeRemaing  float64 `json:"sr"`
	SizeLapse    float64 `json:"sl"`
	SizeCanceled float64 `json:"sc"`
	SizeVoided   float64 `json:"sv"`
}
