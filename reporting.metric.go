package dv360

import (
	"encoding/json"
	"time"
)

type DateMetric struct {
	time.Time
}

// UnmarshalJSON Handle YYYY/MM/DD format from Google DBM reports
func (d *DateMetric) UnmarshalJSON(bytes []byte) error {
	var dateAsString string

	if err := json.Unmarshal(bytes, &dateAsString); err != nil {
		return err
	}

	if day, err := time.Parse("2006/01/02", dateAsString); err != nil {
		return err
	} else {
		d.Time = day
	}

	return nil
}

type CampaignMetric struct {
	CampaignId          string      `json:"campaign_id"`
	CampaignName        string      `json:"campaign"`
	Date                *DateMetric `json:"date"`
	Impressions         int64       `json:"impressions,string"`
	Clicks              int64       `json:"clicks,string"`
	Conversions         float64     `json:"total_conversions,string"`
	AdvertiserTotalCost float64     `json:"total_media_cost_advertiser_currency,string"`
	PartnerTotalCost    float64     `json:"total_media_cost_partner_currency,string"`
}
