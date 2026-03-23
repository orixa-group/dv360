package dv360

import (
	"fmt"
	"google.golang.org/api/doubleclickbidmanager/v2"
	"time"
)

func newReportingDateRange(from, to time.Time) *doubleclickbidmanager.DataRange {
	return &doubleclickbidmanager.DataRange{
		Range: "CUSTOM_DATES",
		CustomStartDate: &doubleclickbidmanager.Date{
			Month: int64(from.Month()),
			Year:  int64(from.Year()),
			Day:   int64(from.Day()),
		},
		CustomEndDate: &doubleclickbidmanager.Date{
			Month: int64(to.Month()),
			Year:  int64(to.Year()),
			Day:   int64(to.Day()),
		},
	}
}

// newReportingQuery
// Metrics & filters = https://developers.google.com/bid-manager/reference/rest/v2/filters-metrics
func newReportingQuery(accountId string, from, to time.Time) *doubleclickbidmanager.Query {
	return &doubleclickbidmanager.Query{
		Schedule: &doubleclickbidmanager.QuerySchedule{
			Frequency: "ONE_TIME",
		},
		Params: &doubleclickbidmanager.Parameters{
			GroupBys: []string{
				"FILTER_MEDIA_PLAN",      // campaign_id
				"FILTER_MEDIA_PLAN_NAME", // campaign_name
				"FILTER_DATE",
				"FILTER_ADVERTISER_CURRENCY",
				"FILTER_PARTNER_CURRENCY",
			},
			Filters: []*doubleclickbidmanager.FilterPair{
				{
					Type:  "FILTER_ADVERTISER",
					Value: accountId,
				},
			},
			Metrics: []string{
				"METRIC_IMPRESSIONS",
				"METRIC_CLICKS",
				"METRIC_TOTAL_CONVERSIONS",
				"METRIC_TOTAL_MEDIA_COST_ADVERTISER",
				"METRIC_TOTAL_MEDIA_COST_PARTNER",
			},
		},
		Metadata: &doubleclickbidmanager.QueryMetadata{
			DataRange: newReportingDateRange(from, to),
			Format:    "CSV",
			Title:     fmt.Sprintf("Customer report %s-%s", from.Format(time.DateOnly), to.Format(time.DateOnly)),
		},
	}
}
