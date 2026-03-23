package dv360

import "google.golang.org/api/googleapi"

func withPartnerId(partnerId string) googleapi.CallOption {
	return googleapi.QueryParameter("partnerId", partnerId)
}
