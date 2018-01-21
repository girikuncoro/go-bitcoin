package btc_client

const (
	BitcoinURL = "https://vip.bitcoin.co.id/tapi"
	TickerURL  = "https://vip.bitcoin.co.id/api/%s_idr/ticker"
)

const (
	ParamNonce   = "nonce"
	ParamMethod  = "method"
	ParamGetInfo = "getInfo"
)

const (
	HeaderContentType    = "Content-Type"
	HeaderFormUrlEncoded = "application/x-www-form-urlencoded"
	HeaderSignedMessage  = "Sign"
	HeaderApiKey         = "Key"
)
