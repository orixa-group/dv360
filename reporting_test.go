package dv360

import (
	"testing"
	"time"
)

func TestGetAccountReport(t *testing.T) {
	initTestSession()

	from, _ := time.Parse(time.DateOnly, "2026-03-01")
	to, _ := time.Parse(time.DateOnly, "2026-03-31")

	metrics, err := GetAccountReport("802111824", from, to)
	assertNotEmptyResult(t, metrics, err)

	for _, metric := range metrics {
		if len(metric.CampaignName) == 0 {
			t.Errorf("invalid campaign name")
		}
		if metric.AdvertiserTotalCost == 0 {
			t.Errorf("metric advertiser total cost is zero")
		}
		if metric.PartnerTotalCost == 0 {
			t.Errorf("metric partner total cost is zero")
		}
	}
}
