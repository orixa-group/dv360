package dv360

import (
	"os"
	"testing"
	"time"
)

func Test_convertCsvToArray(t *testing.T) {
	r, err := os.Open("tests/campaign_report.csv")
	if nil != err {
		t.Fatal(err)
	} else {
		defer r.Close()
	}

	metrics, err := convertCsvToArray[CampaignMetric](r)
	assertNotEmptyResult(t, metrics, err)

	minDate, _ := time.Parse(time.DateOnly, "2026-03-01")
	maxDate, _ := time.Parse(time.DateOnly, "2026-03-31")

	// Must be 40 rows of metrics in sample file
	if expected, current := 40, len(metrics); expected != current {
		t.Fatalf("want %d metrics, got %d", expected, current)
	}

	for _, metric := range metrics {
		if len(metric.CampaignId) == 0 {
			t.Error("metric should have a campaign_id")
		}
		if metric.Impressions == 0 {
			t.Error("metric should have some impressions")
		}
		if nil == metric.Date {
			t.Error("metric should have a date")
		} else if metric.Date.Before(minDate) || metric.Date.After(maxDate) {
			t.Errorf(
				"metric date '%s' out of range [%s - %s]",
				metric.Date.Format(time.DateOnly),
				minDate.Format(time.DateOnly),
				maxDate.Format(time.DateOnly),
			)
		}
	}
}
