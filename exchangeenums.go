package bfapi

import (
	"fmt"
)

//
type MarketProjection uint8

const (
	// MPEvent "EVENT", If not selected then the event will not be returned with marketCatalogue
	MPEvent MarketProjection = iota + 1
	// MPCompetition "COMPETITION", If not selected then the competition will not be returned with marketCatalogue
	MPCompetition
	// MPEventType "EVENT_TYPE", If not selected then the eventType will not be returned with marketCatalogue
	MPEventType
	// MPMarketStartTime "MARKET_START_TIME", If not selected then the start time will not be returned with marketCatalogue
	MPMarketStartTime
	// MPMarketDescription "MARKET_DESCRIPTION", If not selected then the description will not be returned with marketCatalogue
	MPMarketDescription
	// MPRunnerDescription "RUNNER_DESCRIPTION", If not selected then the runners will not be returned with marketCatalogue
	MPRunnerDescription
	// MPRunnerMetadata "RUNNER_METADATA", If not selected then the runner metadata will not be returned with marketCatalogue. If selected then RUNNER_DESCRIPTION will also be returned regardless of whether it is included as a market projection.
	MPRunnerMetadata
)

//
func (mp MarketProjection) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"EVENT"`:
		mp = MPEvent
	case `"COMPETITION"`:
		mp = MPCompetition
	case `"EVENT_TYPE"`:
		mp = MPEventType
	case `"MARKET_START_TIME"`:
		mp = MPMarketStartTime
	case `"MARKET_DESCRIPTION"`:
		mp = MPMarketDescription
	case `"RUNNER_DESCRIPTION"`:
		mp = MPRunnerDescription
	case `"RUNNER_METADATA"`:
		mp = MPRunnerMetadata
	default:
		return fmt.Errorf("Invalid MarketProjection value received")
	}

	return nil
}

//
func (mp MarketProjection) MarshalJSON() ([]byte, error) {
	switch mp {
	case MPEvent:
		return []byte(`"EVENT"`), nil
	case MPCompetition:
		return []byte(`"COMPETITION"`), nil
	case MPEventType:
		return []byte(`"EVENT_TYPE"`), nil
	case MPMarketStartTime:
		return []byte(`"MARKET_START_TIME"`), nil
	case MPMarketDescription:
		return []byte(`"MARKET_DESCRIPTION"`), nil
	case MPRunnerDescription:
		return []byte(`"RUNNER_DESCRIPTION"`), nil
	case MPRunnerMetadata:
		return []byte(`"RUNNER_METADATA"`), nil
	default:
		return nil, fmt.Errorf("Invalid MarketProjection value received")
	}
}

//
type PriceData uint8

const (
	// PDSPAvailable "SP_AVAILABLE", Amount available for the BSP auction.
	PDSPAvailable PriceData = iota + 1
	// PDSPTraded "SP_TRADED", Amount traded in the BSP auction.
	PDSPTraded
	// PDExBestOffers "EX_BEST_OFFERS", Only the best prices available for each runner, to requested price depth.
	PDExBestOffers
	// PDExAllOffers "EX_ALL_OFFERS", EX_ALL_OFFERS trumps EX_BEST_OFFERS if both settings are present
	PDExAllOffers
	// PDExTraded "EX_TRADED", Amount traded on the exchange.
	PDExTraded
)

//
func (pd PriceData) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"SP_AVAILABLE"`:
		pd = PDSPAvailable
	case `"SP_TRADED"`:
		pd = PDSPTraded
	case `"EX_BEST_OFFERS"`:
		pd = PDExBestOffers
	case `"EX_ALL_OFFERS"`:
		pd = PDExAllOffers
	case `"EX_TRADED"`:
		pd = PDExTraded
	default:
		return fmt.Errorf("Invalid PriceData value received")
	}

	return nil
}

//
func (pd PriceData) MarshalJSON() ([]byte, error) {
	switch pd {
	case PDSPAvailable:
		return []byte(`"SP_AVAILABLE"`), nil
	case PDSPTraded:
		return []byte(`"SP_TRADED"`), nil
	case PDExBestOffers:
		return []byte(`"EX_BEST_OFFERS"`), nil
	case PDExAllOffers:
		return []byte(`"EX_ALL_OFFERS"`), nil
	case PDExTraded:
		return []byte(`"EX_TRADED"`), nil
	default:
		return nil, fmt.Errorf("Invalid PriceData value received")
	}
}

//
type MatchProjection uint8

const (
	// MPNoRollUp "NO_ROLLUP", No rollup, return raw fragments
	MPNoRollUp MatchProjection = iota + 1
	// MPRolledUpByPrice "ROLLED_UP_BY_PRICE", Rollup matched amounts by distinct matched prices per side
	MPRolledUpByPrice
	// MPRolledUpByAvgPrice "ROLLED_UP_BY_AVG_PRICE", Rollup matched amounts by average matched price per side
	MPRolledUpByAvgPrice
)

//
func (mp MatchProjection) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"NO_ROLLUP"`:
		mp = MPNoRollUp
	case `"ROLLED_UP_BY_PRICE"`:
		mp = MPRolledUpByPrice
	case `"ROLLED_UP_BY_AVG_PRICE"`:
		mp = MPRolledUpByAvgPrice
	default:
		return fmt.Errorf("Invalid MatchProjection value received")
	}

	return nil
}

//
func (mp MatchProjection) MarshalJSON() ([]byte, error) {
	switch mp {
	case MPNoRollUp:
		return []byte(`"NO_ROLLUP"`), nil
	case MPRolledUpByPrice:
		return []byte(`"ROLLED_UP_BY_PRICE"`), nil
	case MPRolledUpByAvgPrice:
		return []byte(`"ROLLED_UP_BY_AVG_PRICE"`), nil
	default:
		return nil, fmt.Errorf("Invalid MatchProjection value received")
	}
}

//
type OrderProjection uint8

const (
	// OPAll "ALL", EXECUTABLE and EXECUTION_COMPLETE orders
	OPAll OrderProjection = iota + 1
	// OPExecutable "EXECUTABLE", An order that has a remaining unmatched portion
	OPExecutable
	// OPExecutionComplete "EXECUTION_COMPLETE", An order that does not have any remaining unmatched portion
	OPExecutionComplete
)

//
func (op OrderProjection) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"ALL"`:
		op = OPAll
	case `"EXECUTABLE"`:
		op = OPExecutable
	case `"EXECUTION_COMPLETE"`:
		op = OPExecutionComplete
	default:
		return fmt.Errorf("Invalid OrderProjection value received")
	}

	return nil
}

//
func (op OrderProjection) MarshalJSON() ([]byte, error) {
	switch op {
	case OPAll:
		return []byte(`"ALL"`), nil
	case OPExecutable:
		return []byte(`"EXECUTABLE"`), nil
	case OPExecutionComplete:
		return []byte(`"EXECUTION_COMPLETE"`), nil
	default:
		return nil, fmt.Errorf("Invalid OrderProjection value received")
	}
}

//
type MarketStatus uint8

const (
	// MSInactive "INACTIVE", The market has been created but isn't yet available.
	MSInactive MarketStatus = iota + 1
	// MSOpen "OPEN", The market is open for betting.
	MSOpen
	// MSSuspended "SUSPENDED", The market is suspended and not available for betting.
	MSSuspended
	// MSClosed "CLOSED", The market has been settled and is no longer available for betting.
	MSClosed
)

//
func (ms MarketStatus) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"INACTIVE"`:
		ms = MSInactive
	case `"OPEN"`:
		ms = MSOpen
	case `"SUSPENDED"`:
		ms = MSSuspended
	case `"CLOSED"`:
		ms = MSClosed
	default:
		return fmt.Errorf("Invalid MarketStatus value received")
	}

	return nil
}

//
func (ms MarketStatus) MarshalJSON() ([]byte, error) {
	switch ms {
	case MSInactive:
		return []byte(`"INACTIVE"`), nil
	case MSOpen:
		return []byte(`"OPEN"`), nil
	case MSSuspended:
		return []byte(`"SUSPENDED"`), nil
	case MSClosed:
		return []byte(`"CLOSED"`), nil
	default:
		return nil, fmt.Errorf("Invalid MarketStatus value received")
	}
}

//
type RunnerStatus uint8

const (
	// RSActive "ACTIVE"
	RSActive RunnerStatus = iota + 1
	// RSLoser "LOSER"
	RSLoser
	// RSWinner "WINNER"
	RSWinner
	// RSPlaced "PLACED",	The runner was placed, applies to EACH_WAY marketTypes only.
	RSPlaced
	// RSRemovedVacant "REMOVED_VACANT" applies to Greyhounds. Greyhound markets always return a fixed number of runners (traps). If a dog has been removed, the trap is shown as vacant.
	RSRemovedVacant
	// RSRemoved "REMOVED"
	RSRemoved
	// RSHidden "HIDDEN",	The selection is hidden from the market.  This occurs in Horse Racing markets were runners is hidden when it is doesn’t hold an official entry following an entry stage. This could be because the horse was never entered or because they have been scratched from a race at a declaration stage. All matched customer bet prices are set to 1.0 even if there are later supplementary stages. Should it appear likely that a specific runner may actually be supplemented into the race this runner will be reinstated with all matched customer bets set back to the original prices.
	RSHidden
)

//
func (rs RunnerStatus) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"ACTIVE"`:
		rs = RSActive
	case `"LOSER"`:
		rs = RSLoser
	case `"WINNER"`:
		rs = RSWinner
	case `"PLACED"`:
		rs = RSPlaced
	case `"REMOVED_VACANT"`:
		rs = RSRemovedVacant
	case `"REMOVED"`:
		rs = RSRemoved
	case `"HIDDEN"`:
		rs = RSHidden
	default:
		return fmt.Errorf("Invalid RunnerStatus value received")
	}

	return nil
}

//
func (rs RunnerStatus) MarshalJSON() ([]byte, error) {
	switch rs {
	case RSActive:
		return []byte(`"ACTIVE"`), nil
	case RSLoser:
		return []byte(`"LOSER"`), nil
	case RSWinner:
		return []byte(`"WINNER"`), nil
	case RSPlaced:
		return []byte(`"PLACED"`), nil
	case RSRemovedVacant:
		return []byte(`"REMOVED_VACANT"`), nil
	case RSRemoved:
		return []byte(`"REMOVED"`), nil
	case RSHidden:
		return []byte(`"HIDDEN"`), nil
	default:
		return nil, fmt.Errorf("Invalid RunnerStatus value received")
	}
}

//
type TimeGranularity uint8

const (
	// TGDays "DAYS"
	TGDays TimeGranularity = iota + 1
	// TGHours "HOURS"
	TGHours
	// TGMinutes "MINUTES"
	TGMinutes
)

//
func (tg TimeGranularity) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"DAYS"`:
		tg = TGDays
	case `"HOURS"`:
		tg = TGHours
	case `"MINUTES"`:
		tg = TGMinutes
	default:
		return fmt.Errorf("Invalid TimeGranularity value received")
	}

	return nil
}

//
func (tg TimeGranularity) MarshalJSON() ([]byte, error) {
	switch tg {
	case TGDays:
		return []byte(`"DAYS"`), nil
	case TGHours:
		return []byte(`"HOURS"`), nil
	case TGMinutes:
		return []byte(`"MINUTES"`), nil
	default:
		return nil, fmt.Errorf("Invalid TimeGranularity value received")
	}
}

//
type Side uint8

const (
	// SBack "BACK", To back a team, horse or outcome is to bet on the selection to win. For Line markets a Back bet refers to a SELL line. A SELL line will win if the outcome is LESS THAN the taken line (price).
	SBack Side = iota + 1
	// SLay "LAY", To lay a team, horse, or outcome is to bet on the selection to lose. For line markets a Lay bet refers to a BUY line. A BUY line will win if the outcome is MORE THAN the taken line (price).
	SLay
)

//
func (s Side) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"BACK"`:
		s = SBack
	case `"LAY"`:
		s = SLay
	default:
		return fmt.Errorf("Invalid Side value received")
	}

	return nil
}

//
func (s Side) MarshalJSON() ([]byte, error) {
	switch s {
	case SBack:
		return []byte(`"BACK"`), nil
	case SLay:
		return []byte(`"LAY"`), nil
	default:
		return nil, fmt.Errorf("Invalid Side value received")
	}
}

//
type OrderStatus uint8

const (
	// OSPending "PENDING", An asynchronous order is yet to be processed. Once the bet has been processed by the exchange  (including waiting for any in-play delay), the result will be reported and available on the Exchange Stream API and API NG. Not a valid search criteria on MarketFilter
	OSPending OrderStatus = iota + 1
	// OSExecutionComplete "EXECUTION_COMPLETE", An order that does not have any remaining unmatched portion.
	OSExecutionComplete
	// OSExecutable "EXECUTABLE", An order that has a remaining unmatched portion.
	OSExecutable
	// OSExpired "EXPIRED", The order is no longer available for execution due to its time in force constraint.  In the case of FILL_OR_KILL orders, this means the order has been killed because it could not be filled to your specifications.  Not a valid search criteria on MarketFilter
	OSExpired
)

//
func (os OrderStatus) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"PENDING"`:
		os = OSPending
	case `"EXECUTION_COMPLETE"`:
		os = OSExecutionComplete
	case `"EXECUTABLE"`:
		os = OSExecutable
	case `"EXPIRED"`:
		os = OSExpired
	default:
		return fmt.Errorf("Invalid OrderStatus value received")
	}

	return nil
}

//
func (os OrderStatus) MarshalJSON() ([]byte, error) {
	switch os {
	case OSPending:
		return []byte(`"PENDING"`), nil
	case OSExecutionComplete:
		return []byte(`"EXECUTION_COMPLETE"`), nil
	case OSExecutable:
		return []byte(`"EXECUTABLE"`), nil
	case OSExpired:
		return []byte(`"EXPIRED"`), nil
	default:
		return nil, fmt.Errorf("Invalid OrderStatus value received")
	}
}

//
type OrderBy uint8

const (
	// OBByBet "BY_BET", @Deprecated Use BY_PLACE_TIME instead. Order by placed time, then bet id.
	OBByBet OrderBy = iota + 1
	// OBByMarket "BY_MARKET", Order by market id, then placed time, then bet id.
	OBByMarket
	// OBByMatchTime "BY_MATCH_TIME", Order by time of last matched fragment (if any), then placed time, then bet id. Filters out orders which have no matched date. The dateRange filter (if specified) is applied to the matched date.
	OBByMatchTime
	// OBByPlaceTime "BY_PLACE_TIME", Order by placed time, then bet id. This is an alias of to be deprecated BY_BET. The dateRange filter (if specified) is applied to the placed date.
	OBByPlaceTime
	// OBBySettledTime "BY_SETTLED_TIME", Order by time of last settled fragment (if any due to partial market settlement), then by last match time, then placed time, then bet id. Filters out orders which have not been settled. The dateRange filter (if specified) is applied to the settled date.
	OBBySettledTime
	// OBByVoidTime "BY_VOID_TIME", Order by time of last voided fragment (if any), then by last match time, then placed time, then bet id. Filters out orders which have not been voided. The dateRange filter (if specified) is applied to the voided date.
	OBByVoidTime
)

//
func (ob OrderBy) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"BY_BET"`:
		ob = OBByBet
	case `"BY_MARKET"`:
		ob = OBByMarket
	case `"BY_MATCH_TIME"`:
		ob = OBByMatchTime
	case `"BY_PLACE_TIME"`:
		ob = OBByPlaceTime
	case `"BY_SETTLED_TIME"`:
		ob = OBBySettledTime
	case `"BY_VOID_TIME"`:
		ob = OBByVoidTime
	default:
		return fmt.Errorf("Invalid OrderBy value received")
	}

	return nil
}

//
func (ob OrderBy) MarshalJSON() ([]byte, error) {
	switch ob {
	case OBByBet:
		return []byte(`"BY_BET"`), nil
	case OBByMarket:
		return []byte(`"BY_MARKET"`), nil
	case OBByMatchTime:
		return []byte(`"BY_MATCH_TIME"`), nil
	case OBByPlaceTime:
		return []byte(`"BY_PLACE_TIME"`), nil
	case OBBySettledTime:
		return []byte(`"BY_SETTLED_TIME"`), nil
	case OBByVoidTime:
		return []byte(`"BY_VOID_TIME"`), nil
	default:
		return nil, fmt.Errorf("Invalid OrderBy value received")
	}
}

//
type SortDir uint8

const (
	// SDEarliestToLatest "EARLIEST_TO_LATEST", Order from earliest value to latest e.g. lowest betId is first in the results.
	SDEarliestToLatest SortDir = iota + 1
	// SDLatestToEarliest "LATEST_TO_EARLIEST", Order from the latest value to the earliest e.g. highest betId is first in the results.
	SDLatestToEarliest
)

//
func (sd SortDir) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"EARLIEST_TO_LATEST"`:
		sd = SDEarliestToLatest
	case `"LATEST_TO_EARLIEST"`:
		sd = SDLatestToEarliest
	default:
		return fmt.Errorf("Invalid SortDir value received")
	}

	return nil
}

//
func (sd SortDir) MarshalJSON() ([]byte, error) {
	switch sd {
	case SDEarliestToLatest:
		return []byte(`"EARLIEST_TO_LATEST"`), nil
	case SDLatestToEarliest:
		return []byte(`"LATEST_TO_EARLIEST"`), nil
	default:
		return nil, fmt.Errorf("Invalid SortDir value received")
	}
}

//
type OrderType uint8

const (
	// OTLimit "LIMIT", A normal exchange limit order for immediate execution
	OTLimit OrderType = iota + 1
	// OTLimitOnClose "LIMIT_ON_CLOSE", Limit order for the auction (SP)
	OTLimitOnClose
	// OTMarketOnClose "MARKET_ON_CLOSE", Market order for the auction (SP)
	OTMarketOnClose
)

//
func (ot OrderType) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"LIMIT"`:
		ot = OTLimit
	case `"LIMIT_ON_CLOSE"`:
		ot = OTLimitOnClose
	case `"MARKET_ON_CLOSE"`:
		ot = OTMarketOnClose
	default:
		return fmt.Errorf("Invalid OrderType value received")
	}

	return nil
}

//
func (ot OrderType) MarshalJSON() ([]byte, error) {
	switch ot {
	case OTLimit:
		return []byte(`"LIMIT"`), nil
	case OTLimitOnClose:
		return []byte(`"LIMIT_ON_CLOSE"`), nil
	case OTMarketOnClose:
		return []byte(`"MARKET_ON_CLOSE"`), nil
	default:
		return nil, fmt.Errorf("Invalid OrderType value received")
	}
}

//
type MarketSort uint8

const (
	// MSMinimumTraded "MINIMUM_TRADED", Minimum traded volume
	MSMinimumTraded MarketSort = iota + 1
	// MSMaximumTraded "MAXIMUM_TRADED",    Maximum traded volume
	MSMaximumTraded
	// MSMinimumAvailable "MINIMUM_AVAILABLE", Minimum available to match
	MSMinimumAvailable
	// MSMaximumAvailable "MAXIMUM_AVAILABLE", Maximum available to match
	MSMaximumAvailable
	// MSFirstToStart "FIRST_TO_START",    The closest markets based on their expected start time
	MSFirstToStart
	// MSLastToStart "LAST_TO_START",     The most distant markets based on their expected start time
	MSLastToStart
)

//
func (ms MarketSort) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"MINMIUM_TRADED"`:
		ms = MSMinimumTraded
	case `"MAXIMUM_TRADED"`:
		ms = MSMaximumTraded
	case `"MINIMUM_AVAILABLE"`:
		ms = MSMinimumTraded
	case `"MAXIMUM_AVAILABLE"`:
		ms = MSMaximumAvailable
	case `"FIRST_TO_START"`:
		ms = MSFirstToStart
	case `"LAST_TO_START"`:
		ms = MSLastToStart
	default:
		return fmt.Errorf("Invalid MarketSort value received")
	}

	return nil
}

//
func (ms MarketSort) MarshalJSON() ([]byte, error) {
	switch ms {
	case MSMinimumTraded:
		return []byte(`"MINMIUM_TRADED"`), nil
	case MSMaximumTraded:
		return []byte(`"MAXIMUM_TRADED"`), nil
	case MSMinimumAvailable:
		return []byte(`"MINIMUM_AVAILABLE"`), nil
	case MSMaximumAvailable:
		return []byte(`"MAXIMUM_AVAILABLE"`), nil
	case MSFirstToStart:
		return []byte(`"FIRST_TO_START"`), nil
	case MSLastToStart:
		return []byte(`"LAST_TO_START"`), nil
	default:
		return nil, fmt.Errorf("Invalid MarketSort value received")
	}
}

//
type MarketBettingType uint8

const (
	// MBTOdds "ODDS", Odds Market - Any market that doesn't fit any any of the below categories.
	MBTOdds MarketBettingType = iota + 1
	// MBTLine "LINE", Line Market - LINE markets operate at even-money odds of 2.0. However, price for these markets refers to the line positions available as defined by the markets min-max range and interval steps. Customers either Buy a line (LAY bet, winning if outcome is greater than the taken line (price)) or Sell a line (BACK bet, winning if outcome is less than the taken line (price)). If settled outcome equals the taken line, stake is returned.
	MBTLine
	// MBTRange "RANGE", Range Market - Now Deprecated
	MBTRange
	// MBTAsianHandicapDoubleLine "ASIAN_HANDICAP_DOUBLE_LINE", Asian Handicap Market - A traditional Asian handicap market. Can be identified by marketType ASIAN_HANDICAP
	MBTAsianHandicapDoubleLine
	// MBTAsianHandicapSingleLine "ASIAN_HANDICAP_SINGLE_LINE", Asian Single Line Market - A market in which there can be 0 or multiple winners. e,.g marketType TOTAL_GOALS
	MBTAsianHandicapSingleLine
	// MBTFixedOdds "FIXED_ODDS", Sportsbook Odds Market. This type is deprecated and will be removed in future releases, when Sportsbook markets will be represented as ODDS market but with a different product type.
	MBTFixedOdds
)

//
func (mbt MarketBettingType) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"ODDS"`:
		mbt = MBTOdds
	case `"LINE"`:
		mbt = MBTLine
	case `"RANGE"`:
		mbt = MBTRange
	case `"ASIAN_HANDICAP_DOUBLE_LINE"`:
		mbt = MBTAsianHandicapDoubleLine
	case `"ASIAN_HANDICAP_SINGLE_LINE"`:
		mbt = MBTAsianHandicapSingleLine
	case `"FIXED_ODDS"`:
		mbt = MBTFixedOdds
	default:
		return fmt.Errorf("Invalid MarketBettingType value received")
	}

	return nil
}

//
func (mbt MarketBettingType) MarshalJSON() ([]byte, error) {
	switch mbt {
	case MBTOdds:
		return []byte(`"ODDS"`), nil
	case MBTLine:
		return []byte(`"LINE"`), nil
	case MBTRange:
		return []byte(`"RANGE"`), nil
	case MBTAsianHandicapDoubleLine:
		return []byte(`"ASIAN_HANDICAP_DOUBLE_LINE"`), nil
	case MBTAsianHandicapSingleLine:
		return []byte(`"ASIAN_HANDICAP_SINGLE_LINE"`), nil
	case MBTFixedOdds:
		return []byte(`"FIXED_ODDS"`), nil
	default:
		return nil, fmt.Errorf("Invalid MarketBettingType value received")
	}
}

//
type ExecutionReportStatus uint8

const (
	// ERSSuccess "SUCCESS", Order processed successfully
	ERSSuccess ExecutionReportStatus = iota + 1
	// ERSFailure "FAILURE", Order failed.
	ERSFailure
	// ERSProcessedWithErrors "PROCESSED_WITH_ERRORS", The order itself has been accepted, but at least one (possibly all) actions have generated errors. This error only occurs for replaceOrders, cancelOrders and updateOrders operations. The placeOrders operation will not return PROCESSED_WITH_ERRORS status as it is an atomic operation.
	ERSProcessedWithErrors
	// ERSTimeout "TIMEOUT", The order timed out & the status of the bet is unknown.  If a TIMEOUT error occurs on a placeOrders/replaceOrders request, you should check listCurrentOrders to verify the status of your bets before placing further orders. Please Note: Timeouts will occur after 5 seconds of attempting to process the bet but please allow up to 2 minutes for a timed out order to appear.
	ERSTimeout
)

//
func (ers ExecutionReportStatus) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"SUCCESS"`:
		ers = ERSSuccess
	case `"FAILURE"`:
		ers = ERSFailure
	case `"PROCESSED_WITH_ERRORS"`:
		ers = ERSProcessedWithErrors
	case `"TIMEOUT"`:
		ers = ERSTimeout
	default:
		return fmt.Errorf("Invalid ExecutionReportStatus value received")
	}

	return nil
}

//
func (ers ExecutionReportStatus) MarshalJSON() ([]byte, error) {
	switch ers {
	case ERSSuccess:
		return []byte(`"SUCCESS"`), nil
	case ERSFailure:
		return []byte(`"FAILURE"`), nil
	case ERSProcessedWithErrors:
		return []byte(`"PROCESSED_WITH_ERRORS"`), nil
	case ERSTimeout:
		return []byte(`"TIMEOUT"`), nil
	default:
		return nil, fmt.Errorf("Invalid ExecutionReportStatus value received")
	}
}

//
type ExecutionReportErrorCode uint8

const (
	// ERECErrorInMatcher "ERROR_IN_MATCHER", The matcher is not healthy
	ERECErrorInMatcher ExecutionReportErrorCode = iota + 1
	// ERECProcessedWithErrors "PROCESSED_WITH_ERRORS", The order itself has been accepted, but at least one (possibly all) actions have generated errors
	ERECProcessedWithErrors
	// ERECBetActionError "BET_ACTION_ERROR", There is an error with an action that has caused the entire order to be rejected. Check the instructionReports errorCode for the reason for the rejection of the order.
	ERECBetActionError
	// ERECInvalidAccountState "INVALID_ACCOUNT_STATE", Order rejected due to the account's status (suspended, inactive, dup cards)
	ERECInvalidAccountState
	// ERECInvalidWalletStatus "INVALID_WALLET_STATUS", Order rejected due to the account's wallet's status
	ERECInvalidWalletStatus
	// ERECInsufficientFunds "INSUFFICIENT_FUNDS", Account has exceeded its exposure limit or available to bet limit
	ERECInsufficientFunds
	// ERECLossLimitExceeded "LOSS_LIMIT_EXCEEDED", The account has exceed the self imposed loss limit
	ERECLossLimitExceeded
	// ERECMarketSuspended "MARKET_SUSPENDED", Market is suspended
	ERECMarketSuspended
	// ERECMarketNotOpenForBetting "MARKET_NOTOPEN_FOR_BETTING", Market is not open for betting. It is either not yet active, suspended or closed awaiting settlement.
	ERECMarketNotOpenForBetting
	// ERECDuplicateTransaction "DUPLICATE_TRANSACTION", Duplicate customer reference data submitted - Please note: There is a time window associated with the de-duplication of duplicate submissions which is 60 second
	ERECDuplicateTransaction
	// ERECInvalidOrder "INVALID_ORDER", Order cannot be accepted by the matcher due to the combination of actions. For example, bets being edited are not on the same market, or order includes both edits and placement
	ERECInvalidOrder
	// ERECInvalidMarketID "INVALID_MARKET_ID", Market doesn't exist
	ERECInvalidMarketID
	// ERECPermissionDenied "PERMISSION_DENIED", Business rules do not allow order to be placed. You are either attempting to place the order using a Delayed Application Key or from a restricted jurisdiction (i.e. USA)
	ERECPermissionDenied
	// ERECDuplicateBetIDs "DUPLICATE_BETIDS", duplicate bet ids found
	ERECDuplicateBetIDs
	// ERECNoActionRequired "NO_ACTION_REQUIRED", Order hasn't been passed to matcher as system detected there will be no state change
	ERECNoActionRequired
	// ERECServiceUnavailable "SERVICE_UNAVAILABLE", The requested service is unavailable
	ERECServiceUnavailable
	// ERECRejectedByRegulator "REJECTED_BY_REGULATOR", The regulator rejected the order. On the Italian Exchange this error will occur if more than 50 bets are sent in a single placeOrders request.
	ERECRejectedByRegulator
	// ERECNoChasing "NO_CHASING", A specific error code that relates to Spanish Exchange markets only which indicates that the bet placed contravenes the Spanish regulatory rules relating to loss chasing.
	ERECNoChasing
	// ERECRegulatorIsNotAvailable "REGULATOR_IS_NOTAVAILABLE", The underlying regulator service is not available.
	ERECRegulatorIsNotAvailable
	// ERECTooManyInstructions "TOO_MANY_INSTRUCTIONS", The amount of orders exceeded the maximum amount allowed to be executed
	ERECTooManyInstructions
	// ERECInvalidMarketVersion "INVALID_MARKET_VERSION", The supplied market version is invalid. Max length allowed for market version is 12.
	ERECInvalidMarketVersion
)

//
func (erec ExecutionReportErrorCode) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"ERROR_IN_MATCHER"`:
		erec = ERECErrorInMatcher
	case `"PROCESSED_WITH_ERRORS"`:
		erec = ERECProcessedWithErrors
	case `"BET_ACTION_ERROR"`:
		erec = ERECBetActionError
	case `"INVALID_ACCOUNT_STATE"`:
		erec = ERECInvalidAccountState
	case `"INVALID_WALLET_STATUS"`:
		erec = ERECInvalidWalletStatus
	case `"INSUFFICIENT_FUNDS"`:
		erec = ERECInsufficientFunds
	case `"LOSS_LIMIT_EXCEEDED"`:
		erec = ERECLossLimitExceeded
	case `"MARKET_SUSPENDED"`:
		erec = ERECMarketSuspended
	case `"MARKET_NOTOPEN_FOR_BETTING"`:
		erec = ERECMarketNotOpenForBetting
	case `"DUPLICATE_TRANSACTION"`:
		erec = ERECDuplicateTransaction
	case `"INVALID_ORDER"`:
		erec = ERECInvalidOrder
	case `"INVALID_MARKET_ID"`:
		erec = ERECInvalidMarketID
	case `"PERMISSION_DENIED"`:
		erec = ERECPermissionDenied
	case `"DUPLICATE_BETIDS"`:
		erec = ERECDuplicateBetIDs
	case `"NO_ACTION_REQUIRED"`:
		erec = ERECNoActionRequired
	case `"SERVICE_UNAVAILABLE"`:
		erec = ERECServiceUnavailable
	case `"REJECTED_BY_REGULATOR"`:
		erec = ERECRejectedByRegulator
	case `"NO_CHASING"`:
		erec = ERECNoChasing
	case `"REGULATOR_IS_NOTAVAILABLE"`:
		erec = ERECRegulatorIsNotAvailable
	case `"TOO_MANY_INSTRUCTIONS"`:
		erec = ERECTooManyInstructions
	case `"INVALID_MARKET_VERSION"`:
		erec = ERECInvalidMarketVersion
	default:
		return fmt.Errorf("Invalid ExecutionReportErrorCode value received")
	}

	return nil
}

//
func (erec ExecutionReportErrorCode) MarshalJSON() ([]byte, error) {
	switch erec {
	case ERECErrorInMatcher:
		return []byte(`"ERROR_IN_MATCHER"`), nil
	case ERECProcessedWithErrors:
		return []byte(`"PROCESSED_WITH_ERRORS"`), nil
	case ERECBetActionError:
		return []byte(`"BET_ACTION_ERROR"`), nil
	case ERECInvalidAccountState:
		return []byte(`"INVALID_ACCOUNT_STATE"`), nil
	case ERECInvalidWalletStatus:
		return []byte(`"INVALID_WALLET_STATUS"`), nil
	case ERECInsufficientFunds:
		return []byte(`"INSUFFICIENT_FUNDS"`), nil
	case ERECLossLimitExceeded:
		return []byte(`"LOSS_LIMIT_EXCEEDED"`), nil
	case ERECMarketSuspended:
		return []byte(`"MARKET_SUSPENDED"`), nil
	case ERECMarketNotOpenForBetting:
		return []byte(`"MARKET_NOTOPEN_FOR_BETTING"`), nil
	case ERECDuplicateTransaction:
		return []byte(`"DUPLICATE_TRANSACTION"`), nil
	case ERECInvalidOrder:
		return []byte(`"INVALID_ORDER"`), nil
	case ERECInvalidMarketID:
		return []byte(`"INVALID_MARKET_ID"`), nil
	case ERECPermissionDenied:
		return []byte(`"PERMISSION_DENIED"`), nil
	case ERECDuplicateBetIDs:
		return []byte(`"DUPLICATE_BETIDS"`), nil
	case ERECNoActionRequired:
		return []byte(`"NO_ACTION_REQUIRED"`), nil
	case ERECServiceUnavailable:
		return []byte(`"SERVICE_UNAVAILABLE"`), nil
	case ERECRejectedByRegulator:
		return []byte(`"REJECTED_BY_REGULATOR"`), nil
	case ERECNoChasing:
		return []byte(`"NO_CHASING"`), nil
	case ERECRegulatorIsNotAvailable:
		return []byte(`"REGULATOR_IS_NOTAVAILABLE"`), nil
	case ERECTooManyInstructions:
		return []byte(`"TOO_MANY_INSTRUCTIONS"`), nil
	case ERECInvalidMarketVersion:
		return []byte(`"INVALID_MARKET_VERSION"`), nil
	default:
		return nil, fmt.Errorf("Invalid ExecutionReportErrorCode value received")
	}
}

//
type PersistenceType uint8

const (
	// PTLapse "LAPSE", Lapse the order when the market is turned in-play
	PTLapse PersistenceType = iota + 1
	// PTPersist "PERSIST", Persist the order to in-play. The bet will be place automatically into the in-play market at the start of the event.
	PTPersist
	// PTMarketOnClose "MARKET_ON_CLOSE", Put the order into the auction (SP) at turn-in-play
	PTMarketOnClose
)

//
func (pt PersistenceType) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"LAPSE"`:
		pt = PTLapse
	case `"PERSIST"`:
		pt = PTPersist
	case `"MARKET_ON_CLOSE"`:
		pt = PTMarketOnClose
	default:
		return fmt.Errorf("Invalid PersistenceType value received")
	}

	return nil
}

//
func (pt PersistenceType) MarshalJSON() ([]byte, error) {
	switch pt {
	case PTLapse:
		return []byte(`"LAPSE"`), nil
	case PTPersist:
		return []byte(`"PERSIST"`), nil
	case PTMarketOnClose:
		return []byte(`"MARKET_ON_CLOSE"`), nil
	default:
		return nil, fmt.Errorf("Invalid PersistenceType value received")
	}
}

//
type InstructionReportStatus uint8

const (
	// IRSSuccess "SUCCESS", The instruction was successful.
	IRSSuccess InstructionReportStatus = iota + 1
	// IRSFailure "FAILURE", The instruction failed.
	IRSFailure
	// IRSTimeout "TIMEOUT", The instruction timed out & the status of the bet is unknown.  If a TIMEOUT error occurs on a placeOrders/replaceOrders request, you should check listCurrentOrders to verify the status of your bets before placing further orders. Please Note: Timeouts will occur after 5 seconds of attempting to process the bet but please allow up to 2 minutes for timed out order to appear.
	IRSTimeout
)

//
func (irs InstructionReportStatus) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"SUCCESS"`:
		irs = IRSSuccess
	case `"FAILURE"`:
		irs = IRSFailure
	case `"TIMEOUT"`:
		irs = IRSTimeout
	default:
		return fmt.Errorf("Invalid InstructionReportStatus value received")
	}

	return nil
}

//
func (irs InstructionReportStatus) MarshalJSON() ([]byte, error) {
	switch irs {
	case IRSSuccess:
		return []byte(`"SUCCESS"`), nil
	case IRSFailure:
		return []byte(`"FAILURE"`), nil
	case IRSTimeout:
		return []byte(`"TIMEOUT"`), nil
	default:
		return nil, fmt.Errorf("Invalid InstructionReportStatus value received")
	}
}

//
type InstructionReportErrorCode uint8

const (
	// IRECInvalidBetSize "INVALID_BET_SIZE", bet size is invalid for your currency or your regulator
	IRECInvalidBetSize InstructionReportErrorCode = iota + 1
	// IRECInvalidRunner "INVALID_RUNNER", Runner does not exist, includes vacant traps in greyhound racing
	IRECInvalidRunner
	// IRECBetTakenOrLapsed "BET_TAKEN_OR_LAPSED", Bet cannot be cancelled or modified as it has already been taken or has been cancelled/lapsed Includes attempts to cancel/modify market on close BSP bets and cancelling limit on close BSP bets. The error may be returned on placeOrders request if for example a bet is placed at the point when a market admin event takes place (i.e. market is turned in-play)
	IRECBetTakenOrLapsed
	// IRECBetInProgress "BET_IN_PROGRESS", No result was received from the matcher in a timeout configured for the system
	IRECBetInProgress
	// IRECRunnerRemoved "RUNNER_REMOVED", Runner has been removed from the event
	IRECRunnerRemoved
	// IRECMarketNotOpenForBetting "MARKET_NOTOPEN_FOR_BETTING", Attempt to edit a bet on a market that has closed.
	IRECMarketNotOpenForBetting
	// IRECLossLimitExceeded LOSS_LIMIT_EXCEEDED", The action has caused the account to exceed the self imposed loss limit
	IRECLossLimitExceeded
	// IRECMarketNotOpenForBspBetting MARKET_NOTOPEN_FOR_BSP_BETTING", Market now closed to bsp betting. Turned in-play or has been reconciled
	IRECMarketNotOpenForBspBetting
	// IRECInvalidPriceEdit INVALID_PRICE_EDIT", Attempt to edit down the price of a bsp limit on close lay bet, or edit up the price of a limit on close back bet
	IRECInvalidPriceEdit
	// IRECInvalidOdds INVALID_ODDS", Odds not on price ladder - either edit or placement
	IRECInvalidOdds
	// IRECInsufficientFunds "INSUFFICIENT_FUNDS", Insufficient funds available to cover the bet action. Either the exposure limit or available to bet limit would be exceeded
	IRECInsufficientFunds
	// IRECInvalidPersistenceType "INVALID_PERSISTENCE_TYPE", Invalid persistence type for this market, e.g. KEEP for a non in-play market.
	IRECInvalidPersistenceType
	// IRECErrorInMatcher "ERROR_IN_MATCHER", A problem with the matcher prevented this action completing successfully
	IRECErrorInMatcher
	// IRECInvalidBackLayCombination "INVALID_BACK_LAY_COMBINATION", The order contains a back and a lay for the same runner at overlapping prices. This would guarantee a self match. This also applies to BSP limit on close bets
	IRECInvalidBackLayCombination
	// IRECErrorInOrder "ERROR_IN_ORDER", The action failed because the parent order failed
	IRECErrorInOrder
	// IRECInvalidBidType "INVALID_BID_TYPE", Bid type is mandatory
	IRECInvalidBidType
	// IRECInvalidBetID "INVALID_BET_ID", Bet for id supplied has not been found
	IRECInvalidBetID
	// IRECCancelledNotPlaced "CANCELLED_NOTPLACED", Bet cancelled but replacement bet was not placed
	IRECCancelledNotPlaced
	// IRECRelatedActionFailed "RELATED_ACTION_FAILED", Action failed due to the failure of a action on which this action is dependent
	IRECRelatedActionFailed
	// IRECNoActionRequired "NO_ACTION_REQUIRED", the action does not result in any state change. eg changing a persistence to it's current value
	IRECNoActionRequired
	// IRECTimeInForceConflict "TIME_IN_FORCE_CONFLICT", You may only specify a time in force on either the place request OR on individual limit order instructions (not both),  since the implied behaviors are incompatible.
	IRECTimeInForceConflict
	// IRECUnexpectedPersistenceType "UNEXPECTED_PERSISTENCE_TYPE",  You have specified a persistence type for a FILL_OR_KILL order, which is nonsensical because no umatched portion can remain after the order has been placed.
	IRECUnexpectedPersistenceType
	// IRECInvalidOrderType "INVALID_ORDER_TYPE",  You have specified a time in force of FILL_OR_KILL, but have included a non-LIMIT order type.
	IRECInvalidOrderType
	// IRECUnexpectedMinFillSize "UNEXPECTED_MIN_FILL_SIZE",  You have specified a minFillSize on a limit order, where the limit order's time in force is not FILL_OR_KILL. Using minFillSize is not supported where the time in force of the request (as opposed to an order) is FILL_OR_KILL.
	IRECUnexpectedMinFillSize
	// IRECInvalidCustomerOrderRef "INVALID_CUSTOMER_ORDER_REF",  The supplied customer order reference is too long.
	IRECInvalidCustomerOrderRef
	// IRECInvalidMinFillSize "INVALID_MIN_FILL_SIZE", The minFillSize must be greater than zero and less than or equal to the order's size. The minFillSize cannot be less than the minimum bet size for your currency
	IRECInvalidMinFillSize
)

//
func (irec InstructionReportErrorCode) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"INVALID_BET_SIZE"`:
		irec = IRECInvalidBetSize
	case `"INVALID_RUNNER"`:
		irec = IRECInvalidRunner
	case `"BET_TAKEN_OR_LAPSED"`:
		irec = IRECBetTakenOrLapsed
	case `"BET_IN_PROGRESS"`:
		irec = IRECBetInProgress
	case `"RUNNER_REMOVED"`:
		irec = IRECRunnerRemoved
	case `"MARKET_NOTOPEN_FOR_BETTING"`:
		irec = IRECMarketNotOpenForBetting
	case `"LOSS_LIMIT_EXCEEDED"`:
		irec = IRECLossLimitExceeded
	case `"MARKET_NOTOPEN_FOR_BSP_BETTING"`:
		irec = IRECMarketNotOpenForBspBetting
	case `"INVALID_PRICE_EDIT"`:
		irec = IRECInvalidPriceEdit
	case `"INVALID_ODDS"`:
		irec = IRECInvalidOdds
	case `"INSUFFICIENT_FUNDS"`:
		irec = IRECInsufficientFunds
	case `"INVALID_PERSISTENCE_TYPE"`:
		irec = IRECInvalidPersistenceType
	case `"ERROR_IN_MATCHER"`:
		irec = IRECErrorInMatcher
	case `"INVALID_BACK_LAY_COMBINATION"`:
		irec = IRECInvalidBackLayCombination
	case `"ERROR_IN_ORDER"`:
		irec = IRECErrorInOrder
	case `"INVALID_BID_TYPE"`:
		irec = IRECInvalidBidType
	case `"INVALID_BET_ID"`:
		irec = IRECInvalidBetID
	case `"CANCELLED_NOTPLACED"`:
		irec = IRECCancelledNotPlaced
	case `"RELATED_ACTION_FAILED"`:
		irec = IRECRelatedActionFailed
	case `"NO_ACTION_REQUIRED"`:
		irec = IRECNoActionRequired
	case `"TIME_IN_FORCE_CONFLICT"`:
		irec = IRECTimeInForceConflict
	case `"UNEXPECTED_PERSISTENCE_TYPE"`:
		irec = IRECUnexpectedPersistenceType
	case `"INVALID_ORDER_TYPE"`:
		irec = IRECInvalidOrderType
	case `"UNEXPECTED_MIN_FILL_SIZE"`:
		irec = IRECUnexpectedMinFillSize
	case `"INVALID_CUSTOMER_ORDER_REF"`:
		irec = IRECInvalidCustomerOrderRef
	case `"INVALID_MIN_FILL_SIZE"`:
		irec = IRECInvalidBackLayCombination
	default:
		return fmt.Errorf("Invalid InstructionReportErrorCode value received")
	}

	return nil
}

//
func (irec InstructionReportErrorCode) MarshalJSON() ([]byte, error) {
	switch irec {
	case IRECInvalidBetSize:
		return []byte(`"INVALID_BET_SIZE"`), nil
	case IRECInvalidRunner:
		return []byte(`"INVALID_RUNNER"`), nil
	case IRECBetTakenOrLapsed:
		return []byte(`"BET_TAKEN_OR_LAPSED"`), nil
	case IRECBetInProgress:
		return []byte(`"BET_IN_PROGRESS"`), nil
	case IRECRunnerRemoved:
		return []byte(`"RUNNER_REMOVED"`), nil
	case IRECMarketNotOpenForBetting:
		return []byte(`"MARKET_NOTOPEN_FOR_BETTING"`), nil
	case IRECLossLimitExceeded:
		return []byte(`"LOSS_LIMIT_EXCEEDED"`), nil
	case IRECMarketNotOpenForBspBetting:
		return []byte(`"MARKET_NOTOPEN_FOR_BSP_BETTING"`), nil
	case IRECInvalidPriceEdit:
		return []byte(`"INVALID_PRICE_EDIT"`), nil
	case IRECInvalidOdds:
		return []byte(`"INVALID_ODDS"`), nil
	case IRECInsufficientFunds:
		return []byte(`"INSUFFICIENT_FUNDS"`), nil
	case IRECInvalidPersistenceType:
		return []byte(`"INVALID_PERSISTENCE_TYPE"`), nil
	case IRECErrorInMatcher:
		return []byte(`"ERROR_IN_MATCHER"`), nil
	case IRECInvalidBackLayCombination:
		return []byte(`"INVALID_BACK_LAY_COMBINATION"`), nil
	case IRECErrorInOrder:
		return []byte(`"ERROR_IN_ORDER"`), nil
	case IRECInvalidBidType:
		return []byte(`"INVALID_BID_TYPE"`), nil
	case IRECInvalidBetID:
		return []byte(`"INVALID_BET_ID"`), nil
	case IRECCancelledNotPlaced:
		return []byte(`"CANCELLED_NOTPLACED"`), nil
	case IRECRelatedActionFailed:
		return []byte(`"RELATED_ACTION_FAILED"`), nil
	case IRECNoActionRequired:
		return []byte(`"NO_ACTION_REQUIRED"`), nil
	case IRECTimeInForceConflict:
		return []byte(`"TIME_IN_FORCE_CONFLICT"`), nil
	case IRECUnexpectedPersistenceType:
		return []byte(`"UNEXPECTED_PERSISTENCE_TYPE"`), nil
	case IRECInvalidOrderType:
		return []byte(`"INVALID_ORDER_TYPE"`), nil
	case IRECUnexpectedMinFillSize:
		return []byte(`"UNEXPECTED_MIN_FILL_SIZE"`), nil
	case IRECInvalidCustomerOrderRef:
		return []byte(`"INVALID_CUSTOMER_ORDER_REF"`), nil
	case IRECInvalidMinFillSize:
		return []byte(`"INVALID_MIN_FILL_SIZE"`), nil
	default:
		return nil, fmt.Errorf("Invalid InstructionReportErrorCode value received")
	}
}

//
type RollupModel uint8

const (
	// RMStake "STAKE", The volumes will be rolled up to the minimum value which is >= rollupLimit.
	RMStake RollupModel = iota + 1
	// RMPayout "PAYOUT", The volumes will be rolled up to the minimum value where the payout( price * volume ) is >= rollupLimit. On a LINE market, volumes will be rolled up where payout( 2.0 * volume ) is >= rollupLimit
	RMPayout
	// RMManagedLiability "MANAGED_LIABILITY", The volumes will be rolled up to the minimum value which is >= rollupLimit, until a lay price threshold. There after, the volumes will be rolled up to the minimum value such that the liability >= a minimum liability. Not supported as yet.
	RMManagedLiability
	// RMNone "NONE", No rollup will be applied. However the volumes will be filtered by currency specific minimum stake unless overridden specifically for the channel.
	RMNone
)

//
func (rm RollupModel) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"STAKE"`:
		rm = RMStake
	case `"PAYOUT"`:
		rm = RMPayout
	case `"MANAGED_LIABILITY"`:
		rm = RMManagedLiability
	case `"NONE"`:
		rm = RMNone
	default:
		return fmt.Errorf("Invalid RollupModel value received")
	}

	return nil
}

//
func (rm RollupModel) MarshalJSON() ([]byte, error) {
	switch rm {
	case RMStake:
		return []byte(`"STAKE"`), nil
	case RMPayout:
		return []byte(`"PAYOUT"`), nil
	case RMManagedLiability:
		return []byte(`"MANAGED_LIABILITY"`), nil
	case RMNone:
		return []byte(`"NONE"`), nil
	default:
		return nil, fmt.Errorf("Invalid RollupModel value received")
	}
}

//
type GroupBy uint8

const (
	// GBEventType "EVENT_TYPE", A roll up of settled P&L, commission paid and number of bet orders, on a specified event type
	GBEventType GroupBy = iota + 1
	// GBEvent "EVENT", A roll up of settled P&L, commission paid and number of bet orders, on a specified event
	GBEvent
	// GBMarket "MARKET", A roll up of settled P&L, commission paid and number of bet orders, on a specified market
	GBMarket
	// GBSide "SIDE", An averaged roll up of settled P&L, and number of bets, on the specified side of a specified selection within a specified market, that are either settled or voided
	GBSide
	// GBBet "BET", The P&L, commission paid, side and regulatory information etc, about each individual bet order
	GBBet
)

//
func (gb GroupBy) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"EVENT_TYPE"`:
		gb = GBEventType
	case `"EVENT"`:
		gb = GBEvent
	case `"MARKET"`:
		gb = GBMarket
	case `"SIDE"`:
		gb = GBSide
	case `"BET"`:
		gb = GBBet
	default:
		return fmt.Errorf("Invalid GroupBy value received")
	}

	return nil
}

//
func (gb GroupBy) MarshalJSON() ([]byte, error) {
	switch gb {
	case GBEventType:
		return []byte(`"EVENT_TYPE"`), nil
	case GBEvent:
		return []byte(`"EVENT"`), nil
	case GBMarket:
		return []byte(`"MARKET"`), nil
	case GBSide:
		return []byte(`"SIDE"`), nil
	case GBBet:
		return []byte(`"BET"`), nil
	default:
		return nil, fmt.Errorf("Invalid GroupBy value received")
	}
}

//
type BetStatus uint8

const (
	// BSSettled "SETTLED", A matched bet that was settled normally
	BSSettled BetStatus = iota + 1
	// BSVoided "VOIDED", A matched bet that was subsequently voided by Betfair, before, during or after settlement
	BSVoided
	// BSLapsed "LAPSED", Unmatched bet that was cancelled by Betfair (for example at turn in play).
	BSLapsed
	// BSCancelled "CANCELLED", Unmatched bet that was cancelled by an explicit customer action.
	BSCancelled
)

//
func (bs BetStatus) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"SETTLED"`:
		bs = BSSettled
	case `"VOIDED"`:
		bs = BSVoided
	case `"LAPSED"`:
		bs = BSLapsed
	case `"CANCELLED"`:
		bs = BSCancelled
	default:
		return fmt.Errorf("Invalid BetStatus value received")
	}

	return nil
}

//
func (bs BetStatus) MarshalJSON() ([]byte, error) {
	switch bs {
	case BSSettled:
		return []byte(`"SETTLED"`), nil
	case BSVoided:
		return []byte(`"VOIDED"`), nil
	case BSLapsed:
		return []byte(`"LAPSED"`), nil
	case BSCancelled:
		return []byte(`"CANCELLED"`), nil
	default:
		return nil, fmt.Errorf("Invalid BetStatus value received")
	}
}

//
type TimeInForce uint8

const (
	// TIFFillOrKill "FILL_OR_KILL", Execute the transaction immediately and completely (filled to size or between minFillSize and size) or not at all (cancelled). For LINE markets Volume Weighted Average Price (VWAP) functionality is disabled
	TIFFillOrKill TimeInForce = iota + 1
)

//
func (tif TimeInForce) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"FILL_OR_KILL"`:
		tif = TIFFillOrKill
	default:
		return fmt.Errorf("Invalid TimeInForce value received")
	}

	return nil
}

//
func (tif TimeInForce) MarshalJSON() ([]byte, error) {
	switch tif {
	case TIFFillOrKill:
		return []byte(`"FILL_OR_KILL"`), nil
	default:
		return nil, fmt.Errorf("Invalid TimeInForce value received")
	}
}

//
type BetTargetType uint8

const (
	// BTTBackersProfit "BACKERSPROFIT", The payout requested minus the calculated size at which this LimitOrder is to be placed. BetTargetType bets are invalid for LINE markets
	BTTBackersProfit BetTargetType = iota + 1
	// BTTPayout "PAYOUT", The total payout requested on a LimitOrder
	BTTPayout
)

//
func (btt BetTargetType) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"BACKERSPROFIT`:
		btt = BTTBackersProfit
	case `"PAYOUT"`:
		btt = BTTPayout
	default:
		return fmt.Errorf("Invalid BetTargetType value received")
	}

	return nil
}

//
func (btt BetTargetType) MarshalJSON() ([]byte, error) {
	switch btt {
	case BTTBackersProfit:
		return []byte(`"BACKERSPROFIT`), nil
	case BTTPayout:
		return []byte(`"PAYOUT"`), nil
	default:
		return nil, fmt.Errorf("Invalid BetTargetType value received")
	}
}

//
type PriceLadderType uint8

const (
	// PLTClassic "CLASSIC", Price ladder increments traditionally used for Odds Markets.
	PLTClassic PriceLadderType = iota + 1
	// PLTFinest "FINEST", Price ladder with the finest available increment, traditionally used for Asian Handicap markets.
	PLTFinest
	// PLTLineRange "LINE_RANGE", Price ladder used for LINE markets. Refer to MarketLineRangeInfo for more details.
	PLTLineRange
)

//
func (plt PriceLadderType) UnmarshallJSON(data []byte) error {
	switch string(data) {
	case `"CLASSIC"`:
		plt = PLTClassic
	case `"FINEST"`:
		plt = PLTFinest
	case `"LINE_RANGE"`:
		plt = PLTLineRange
	default:
		return fmt.Errorf("Invalid PriceLadderType value received")
	}

	return nil
}

//
func (plt PriceLadderType) MarshalJSON() ([]byte, error) {
	switch plt {
	case PLTClassic:
		return []byte(`"CLASSIC"`), nil
	case PLTFinest:
		return []byte(`"FINEST"`), nil
	case PLTLineRange:
		return []byte(`"LINE_RANGE"`), nil
	default:
		return nil, fmt.Errorf("Invalid PriceLadderType value received")
	}
}
