# BFAPI - Betfair Api

Implementation of Betfairs REST Http Api

### Use

Simple library implementing the Betfair Api types and methods wrapping the requests


To initialize the api, pass in cert/key bytes, appkey and testing/prod endpoints (bool)
bfapi methods shouldnt be called until after Inits execution.
```go
var certBytes, keyBytes []byte
certBytes, _ = ioutil.ReadFile("client-2048.crt")
keyBytes, _ = ioutil.ReadFile("client-2048.key")

bfapi.Init(certBytes, keyBytes, config.Appkey, true)
```

After Init but before other methods can be called, you must login. CertLogin also adds the returned sessionToken value as a X-Authentication default headers. The sessionToken is returned for storage/logging and is not needed to be passed back in to complete following methods
```go
var sessionToken string
sessionToken, _ = bfapi.CertLogin()
```

The sessionToken provided by CertLogin, expires every 3 hours and KeepAlive must be called to extend this time (by another 3 hours)
```go
var err error = bfapi.KeepAlive()
```


Example of ListMarketCatalogue from outside package
```go
var cat = bfapi.ListMarketCatalogueArg{
    Filter: MarketFilter{
        EventTypes: []string{"4339"}, // greyhounds id code
        Countries:  []string{"GB"},
        TypeCodes:  []string{"WIN"},
    },
    Sort:             "FIRST_TO_START",
    MarketProjection: []string{"MARKET_START_TIME", "RUNNER_DESCRIPTION", "EVENT"},
    MaxResults:       1000,
}

var result []MarketCatalogue
result, _ = bfapi.ListMarketCatalogue(cat)
```

### TODO

* Documentation - Function/TypeDef comments
* Add more Betfair methods and types
