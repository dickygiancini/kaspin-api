package main

type Credentials struct {
	ResultCd     string `json:"resultCd"`
	ResultMsg    string `json:"resultMsg"`
	TXid         string `json:"tXid"`
	ReferenceNo  string `json:"referenceNo"`
	PayMethod    string `json:"payMethod"`
	Amt          string `json:"amt"`
	Currency     string `json:"currency"`
	GoodsNm      string `json:"goodsNm"`
	BillingNm    string `json:"billingNm"`
	TransDt      string `json:"transDt"`
	TransTm      string `json:"transTm"`
	Description  string `json:"description"`
	BankCd       string `json:"bankCd"`
	VacctNo      string `json:"vacctNo"`
	VacctValidDt string `json:"vacctValidDt"`
	VacctValidTm string `json:"vacctValidTm"`
	MitraCd      string `json:"mitraCd"`
	PayNo        string `json:"payNo"`
	PayValidDt   string `json:"payValidDt"`
	PayValidTm   string `json:"payValidTm"`
	RequestURL   string `json:"requestURL"`
	PaymentExpDt string `json:"paymentExpDt"`
	PaymentExpTm string `json:"paymentExpTm"`
	QrContent    string `json:"qrContent"`
	QrUrl        string `json:"qrUrl"`
}
