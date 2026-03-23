package dv360

import (
	"strconv"
	"testing"
)

func assertNotEmptyResult[T any](t *testing.T, results []*T, err error) {
	if err != nil {
		t.Fatal(err)
	} else if len(results) == 0 {
		t.Fatal("no results found")
	}
}

func TestGetAccounts(t *testing.T) {
	initTestSession()

	partners, err := GetPartners()
	assertNotEmptyResult(t, partners, err)

	accounts, err := GetPartnerAccounts(strconv.Itoa(int(partners[0].PartnerId)))
	assertNotEmptyResult(t, accounts, err)

	for _, account := range accounts {
		t.Logf("%d %s", account.AdvertiserId, account.DisplayName)
	}
}
