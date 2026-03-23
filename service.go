package dv360

import (
	"context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/displayvideo/v4"
	"google.golang.org/api/doubleclickbidmanager/v2"
	"google.golang.org/api/option"
)

func getDV360Service(ctx context.Context) *displayvideo.Service {
	conf := &oauth2.Config{
		ClientID:     apiKey,
		ClientSecret: apiSecret,
		Endpoint:     google.Endpoint,
	}

	srv, _ := displayvideo.NewService(
		ctx,
		option.WithHTTPClient(conf.Client(ctx, token)),
	)

	return srv
}

// getDBMService Google DoubleClick Bid Manager service (reporting)
func getDBMService(ctx context.Context) *doubleclickbidmanager.Service {
	conf := &oauth2.Config{
		ClientID:     apiKey,
		ClientSecret: apiSecret,
		Endpoint:     google.Endpoint,
	}

	srv, _ := doubleclickbidmanager.NewService(
		ctx,
		option.WithHTTPClient(conf.Client(ctx, token)),
	)

	return srv
}
