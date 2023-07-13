package main

type Account struct {
	IMid            string `json:"iMid"`
	MerchantToken   string `json:"merchantToken"`
	TimeStamp       string `json:"timeStamp"`
	PayMethod       string `json:"payMethod"`
	Currency        string `json:"currency"`
	Amt             string `json:"amt"`
	ReferenceNo     string `json:"referenceNo"`
	GoodsNm         string `json:"goodsNm"`
	BillingNm       string `json:"billingNm"`
	BillingPhone    string `json:"billingPhone"`
	BillingEmail    string `json:"billingEmail"`
	BillingAddr     string `json:"billingAddr"`
	BillingCity     string `json:"billingCity"`
	BillingState    string `json:"billingState"`
	BillingPostCd   string `json:"billingPostCd"`
	BillingCountry  string `json:"billingCountry"`
	DeliveryNm      string `json:"deliveryNm"`
	DeliveryPhone   string `json:"deliveryPhone"`
	DeliveryAddr    string `json:"deliveryAddr"`
	DeliveryCity    string `json:"deliveryCity"`
	DeliveryState   string `json:"deliveryState"`
	DeliveryPostCd  string `json:"deliveryPostCd"`
	DeliveryCountry string `json:"deliveryCountry"`
	DbProcessUrl    string `json:"dbProcessUrl"`
	Vat             string `json:"vat"`
	Fee             string `json:"fee"`
	NotaxAmt        string `json:"notaxAmt"`
	Description     string `json:"description"`
	ReqDt           string `json:"reqDt"`
	ReqTm           string `json:"reqTm"`
	ReqDomain       string `json:"reqDomain"`
	ReqServerIP     string `json:"reqServerIP"`
	ReqClientVer    string `json:"reqClientVer"`
	UserIP          string `json:"userIP"`
	UserSessionID   string `json:"userSessionID"`
	UserAgent       string `json:"userAgent"`
	UserLanguage    string `json:"userLanguage"`
	CartData        string `json:"cartData"`
	InstmntType     string `json:"instmntType"`
	InstmntMon      string `json:"instmntMon"`
	RecurrOpt       string `json:"recurrOpt"`
	BankCd          string `json:"bankCd"`
	VacctValidDt    string `json:"vacctValidDt"`
	VacctValidTm    string `json:"vacctValidTm"`
	MerFixAcctId    string `json:"merFixAcctId"`
	MitraCd         string `json:"mitraCd"`
}

func NewAccount(iMid string, merchantToken string, timeStamp string, refNo string, amount string) *Account {
	return &Account{
		IMid:            iMid,
		MerchantToken:   merchantToken,
		TimeStamp:       timeStamp,
		PayMethod:       "00",
		Currency:        "IDR",
		Amt:             amount,
		ReferenceNo:     refNo,
		GoodsNm:         "Test Transaction Nicepay",
		BillingNm:       "John Doe",
		BillingPhone:    "John Doe",
		BillingEmail:    "email@merchant.com",
		BillingAddr:     "Jalan Bukit Berbunga 22",
		BillingCity:     "Jakarta",
		BillingState:    "DKI Jakarta",
		BillingPostCd:   "12345",
		BillingCountry:  "Indonesia",
		DeliveryNm:      "email@merchant.com",
		DeliveryPhone:   "08123456789",
		DeliveryAddr:    "Jalan Bukit Berbunga 22",
		DeliveryCity:    "Jakarta",
		DeliveryState:   "DKI Jakarta",
		DeliveryPostCd:  "12345",
		DeliveryCountry: "Indonesia",
		DbProcessUrl:    "https://merchant.com/api/dbProcessUrl/Notif",
		Vat:             "",
		Fee:             "",
		NotaxAmt:        "",
		Description:     "",
		ReqDt:           "",
		ReqTm:           "",
		ReqDomain:       "merchant.com",
		ReqServerIP:     "127.0.0.1",
		ReqClientVer:    "",
		UserIP:          "127.0.0.1",
		UserSessionID:   "697D6922C961070967D3BA1BA5699C2C",
		UserAgent:       "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML,like Gecko) Chrome/60.0.3112.101 Safari/537.36",
		UserLanguage:    "ko-KR,en-US;q=0.8,ko;q=0.6,en;q=0.4",
		CartData:        "{\"count\":\"1\",\"item\":[{\"goods_id\":\"BB12345678\",\"goods_detail\":\"BB12345678\",\"goods_name\":\"Pasar Modern\",\"goods_amt\":\"25145\",\"goods_type\":\"Sembako\",\"goods_url\":\"http://merchant.com/cellphones/iphone5s_64g\",\"goods_quantity\":\"1\",\"goods_sellers_id\":\"SEL123\",\"goods_sellers_name\":\"Sellers 1\"}]}",
		InstmntType:     "2",
		InstmntMon:      "1",
		RecurrOpt:       "0",
		BankCd:          "",
		VacctValidDt:    "",
		VacctValidTm:    "",
		MerFixAcctId:    "",
		MitraCd:         "",
	}
}
