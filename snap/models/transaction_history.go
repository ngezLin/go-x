package models

import "time"

type (
	TransactionHistoryListRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		FromDateTime       string `json:"fromDateTime"`
		ToDateTime         string `json:"toDateTime"`
		PageSize           string `json:"pageSize"`
		PageNumber         string `json:"pageNumber"`
		AdditionalData     T      `json:"additionalData"`
	}
	DetailDataTrx[T any] struct {
		DateTime       time.Time      `json:"dateTime"`
		Amount         Amount         `json:"amount"`
		Remark         string         `json:"remark"`
		SourceOfFunds  []SourceOfFund `json:"sourceOfFunds"`
		Status         string         `json:"status"`
		Type           string         `json:"type"`
		AdditionalInfo T              `json:"additionalInfo"`
	}
	TransactionHistoryListResponse[T, J any] struct {
		ResponseCode       string             `json:"responseCode"`
		ResponseMessage    string             `json:"responseMessage"`
		ReferenceNo        string             `json:"referenceNo"`
		PartnerReferenceNo string             `json:"partnerReferenceNo"`
		DetailData         []DetailDataTrx[J] `json:"detailData"`
		AdditionalInfo     T                  `json:"additionalInfo"`
	}
)

type (
	TransactionHistoryDetailRequest[T any] struct {
		OriginalPartnerReferenceNo string `json:"originalPartnerReferenceNo"`
		AdditionalData             T      `json:"additionalData"`
	}
	TransactionHistoryDetailResponse[T, J any] struct {
		ResponseCode       string         `json:"responseCode"`
		ResponseMessage    string         `json:"responseMessage"`
		ReferenceNo        string         `json:"referenceNo"`
		PartnerReferenceNo string         `json:"partnerReferenceNo"`
		Amount             Amount         `json:"amount"`
		CancelledTime      string         `json:"cancelledTime"`
		DateTime           string         `json:"dateTime"`
		RefundAmount       Amount         `json:"refundAmount"`
		Remark             string         `json:"remark"`
		SourceOfFunds      []SourceOfFund `json:"sourceOfFunds"`
		Status             string         `json:"status"`
		Type               string         `json:"type"`
		AdditionalInfo     T              `json:"additionalInfo"`
	}
)

type (
	BankStatementRequest[T any] struct {
		PartnerReferenceNo string `json:"partnerReferenceNo"`
		BankCardToken      string `json:"bankCardToken"`
		AccountNo          string `json:"accountNo"`
		FromDateTime       string `json:"fromDateTime"`
		ToDateTime         string `json:"toDateTime"`
		AdditionalData     T      `json:"additionalData"`
	}

	DetailDataBankStatement struct {
		DetailBalance           DetailBalance `json:"detailBalance"`
		Amount                  Amount        `json:"amount"`
		OriginAmount            Amount        `json:"originAmount"`
		TransactionDate         time.Time     `json:"transactionDate"`
		Remark                  string        `json:"remark"`
		TransactionId           string        `json:"transactionId"`
		Type                    string        `json:"type"`
		TransactionDetailStatus string        `json:"transactionDetailStatus"`
		DetailInfo              DetailInfo    `json:"detailInfo"`
	}

	BankStatementResponse[T any] struct {
		ResponseCode       string                    `json:"responseCode"`
		ResponseMessage    string                    `json:"responseMessage"`
		ReferenceNo        string                    `json:"referenceNo"`
		PartnerReferenceNo string                    `json:"partnerReferenceNo"`
		Balance            []BalanceEntry            `json:"balance"`
		TotalCreditEntries TotalEntries              `json:"totalCreditEntries"`
		TotalDebitEntries  TotalEntries              `json:"totalDebitEntries"`
		HasMore            string                    `json:"hasMore"`
		LastRecordDateTime time.Time                 `json:"lastRecordDateTime"`
		DetailData         []DetailDataBankStatement `json:"detailData"`
		AdditionalInfo     T                         `json:"additionalInfo"`
	}
)
