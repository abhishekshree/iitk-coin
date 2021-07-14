package config

var (
	INITIAL_COINS          = 100
	TAX_PERCENT_INTRABATCH = 2.0
	TAX_PERCENT_INTERBATCH = 33.0
	MAX_BALANCE            = 20000.0
)

var REDEEM_LIST = map[string]float64{
	"A": 1000,
	"B": 10000,
	"C": 100.1,
}

var REDEEMABLE_ITEMS = len(REDEEM_LIST)
