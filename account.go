package dv360

import (
	"context"
	"google.golang.org/api/displayvideo/v4"
)

func GetPartners() ([]*displayvideo.Partner, error) {
	resp, err := getDV360Service(context.Background()).Partners.List().Do()
	if err != nil {
		return nil, err
	}

	return resp.Partners, nil
}

func GetPartnerAccounts(partnerId string) ([]*displayvideo.Advertiser, error) {
	resp, err := getDV360Service(context.Background()).Advertisers.List().Do(
		withPartnerId(partnerId),
	)

	if nil != err {
		return nil, err
	}

	return resp.Advertisers, nil
}
