package main

type Payment struct {
	ReferenceNo    string `json:"referenceNo"`
	AuthNo         string `json:"authNo"`
	CcTransType    string `json:"ccTransType"`
	MRefNo         string `json:"mRefNo"`
	IssuBankCd     string `json:"issuBankCd"`
	IssuBankNm     string `json:"issuBankNm"`
	TXid           string `json:"tXid"`
	TransTm        string `json:"transTm"`
	MitraCd        string `json:"mitraCd"`
	RecurringToken string `json:"recurringToken"`
	ResultCd       string `json:"resultCd"`
	TransDt        string `json:"transDt"`
	AcquBankCd     string `json:"acquBankCd"`
	AcquBankNm     string `json:"acquBankNm"`
	InstmntType    string `json:"instmntType"`
	InstmntMon     string `json:"instmntMon"`
	PayMethod      string `json:"payMethod"`
	ReceiptCode    string `json:"receiptCode"`
	CardExpYymm    string `json:"cardExpYymm"`
	CardNo         string `json:"cardNo"`
	Description    string `json:"description"`
	ResultMsg      string `json:"resultMsg"`
	GoodsNm        string `json:"goodsNm"`
	PreauthToken   string `json:"preauthToken"`
	Amt            string `json:"amt"`
	BillingNm      string `json:"billingNm"`
	Currency       string `json:"currency"`
}
