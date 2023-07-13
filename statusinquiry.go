package main

type Inquiry struct {
	TimeStamp     string `json:"timeStamp"`
	TXid          string `json:"tXid"`
	IMid          string `json:"iMid"`
	ReferenceNo   string `json:"referenceNo"`
	Amt           string `json:"amt"`
	MerchantToken string `json:"merchantToken"`
}

func NewStatusInquiry(iMid string, merchantToken string, timeStamp string, tXid string, referenceNo string, amt string) *Inquiry {
	return &Inquiry{
		TimeStamp:     timeStamp,
		TXid:          tXid,
		IMid:          iMid,
		ReferenceNo:   referenceNo,
		Amt:           amt,
		MerchantToken: merchantToken,
	}
}
