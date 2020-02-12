package esi

import (
	"fmt"
	"time"
)

//ESITransaction models
type ESITransaction struct {
	ClientId      int64     `json:"client_id"`
	Date          time.Time `json:"date"`
	IsBuy         bool      `json:"is_buy"`
	IsPersonal    bool      `json:"is_personal"`
	JournalRefId  int64     `json:"journal_ref_id"`
	LocationId    int64     `json:"location_id"`
	Quantity      int64     `json:"quantity"`
	TransactionId int64     `json:"transaction_id"`
	TypeId        int32     `json:"type_id"`
	UnitPrice     float64   `json:"unit_price"`
}

type CharactersWalletTransactionsResponse struct {
	R       []ESITransaction
	Expires time.Time
}

func (esi *ESI) CharactersWalletTransactions(characterId int64, fromId int64) (*CharactersWalletTransactionsResponse, error) {

	url := fmt.Sprintf("%s/characters/%d/wallet/transactions/", ESIRootURL, characterId)
	if fromId != 0 {
		url = url + fmt.Sprintf("?from_id=%d", fromId)
	}

	result := make([]ESITransaction, 0)

	expires, _, err := auth("GET", esi.AccessToken, url, &result)
	if err != nil {
		return nil, err
	}

	return &CharactersWalletTransactionsResponse{R: result, Expires: expires}, nil
}
