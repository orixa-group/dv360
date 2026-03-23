package dv360

import (
	"context"
	"fmt"
	"google.golang.org/api/doubleclickbidmanager/v2"
	"math"
	"net/http"
	"time"
)

func GetAccountReport(accountId string, from, to time.Time) ([]*CampaignMetric, error) {
	reportingService := getDBMService(context.Background()).Queries

	q := newReportingQuery(accountId, from, to)

	resp, err := reportingService.
		Create(q).
		Do()

	if nil != err {
		return nil, err
	}

	url, err := getReportCsvUrl(reportingService, resp.QueryId)
	if nil != err {
		return nil, err
	}

	r, err := http.Get(url)
	if nil != err {
		return nil, err
	} else if r.StatusCode >= http.StatusBadRequest {
		return nil, fmt.Errorf("http error: %d", r.StatusCode)
	}
	defer r.Body.Close()

	return convertCsvToArray[CampaignMetric](r.Body)
}

func getReportCsvUrl(service *doubleclickbidmanager.QueriesService, queryId int64) (string, error) {
	runResp, err := service.
		Run(queryId, &doubleclickbidmanager.RunQueryRequest{}).
		Do()

	if nil != err {
		return "", err
	}

	var status *doubleclickbidmanager.Report

	for i := float64(1); i <= 12; i++ {
		// incremental wait
		time.Sleep(time.Duration(1000*math.Pow(1.44, i)) * time.Millisecond)

		status, err = service.Reports.
			Get(runResp.Key.QueryId, runResp.Key.ReportId).
			Do()

		if nil != err {
			return "", err
		} else if "DONE" == status.Metadata.Status.State {
			return status.Metadata.GoogleCloudStoragePath, nil
			// success
		} else if "FAILED" == status.Metadata.Status.State {
			return "", err
		}
	}

	return "", fmt.Errorf("report generation timeout (last status: %s", status.Metadata.Status.State)
}
