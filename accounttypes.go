package bfapi

// AccountFundsResponse - response from getAccountFunds
type AccountFundsResponse struct {
	AvailableBalance   float64 `json:"availableToBetBalance,omitempty"` //Amount available to bet.
	Exposure           float64 `json:"exposure,omitempty"`              //Current exposure.
	RetainedCommission float64 `json:"retainedCommission,omitempty"`    //Sum of retained commission.
	ExposureLimit      float64 `json:"exposureLimit,omitempty"`         //Exposure limit.
	DiscountRate       float64 `json:"discountRate,omitempty"`          //User Discount Rate.
	PointsBalance      int     `json:"pointsBalance,omitempty"`         //The Betfair points balance
}

// AccountDetailsResponse - response from getAccountDetails
type AccountDetailsResponse struct {
	CurrencyCode  string  `json:"currencyCode,omitempty"`
	FirstName     string  `json:"firstName,omitempty"`
	LastName      string  `json:"lastName,omitempty"`
	LocaleCode    string  `json:"localeCode,omitempty"`
	Region        string  `json:"region,omitempty"`
	Timezone      string  `json:"timezone,omitempty"`
	DiscountRate  float64 `json:"discountRate,omitempty"`
	PointsBalance int     `json:"pointsBalance,omitempty"`
	CountryCode   string  `json:"countryCode,omitempty"`
}
