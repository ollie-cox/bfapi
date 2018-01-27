package bfapi

import "time"

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

//GetSubTokenArg - argument for GetSubscriptionToken
type GetSubTokenArg struct {
	SubscriptionLength int    `json:"subscriptionLength,omitempty"` //number of days
	ClientReference    string `json:"clientReference,omitempty"`
}

//SubscriptionHistory - response from getApplicationSubscriptionHistory
type SubscriptionHistory struct {
	SubscriptionToken    string    `json:"subscriptionToken,omitempty"`    //Application key identifier
	ExpiryDateTime       time.Time `json:"expiryDateTime,omitempty"`       //Subscription Expiry date
	ExpiredDateTime      time.Time `json:"expiredDateTime,omitempty"`      //Subscription Expired date
	CreatedDateTime      time.Time `json:"createdDateTime,omitempty"`      //Subscription Create date
	ActivationDateTime   time.Time `json:"activationDateTime,omitempty"`   //Subscription Activation date
	CancellationDateTime time.Time `json:"cancellationDateTime,omitempty"` //Subscription Cancelled date
	SubscriptionStatus   string    `json:"subscriptionStatus,omitempty"`   //Subscription status {ALL, ACTIVATED, UNACTIVATED, CANCELLED, EXPIRED}
	ClientReference      string    `json:"clientReference,omitempty"`      //Client reference
}
