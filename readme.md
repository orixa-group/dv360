# Google DV360 & DBM Api

## Google Documentations

- [Reporting](https://developers.google.com/bid-manager/reference/rest/v2/filters-metrics#metrics)
- [DV360](https://developers.google.com/display-video/api/guides/getting-started/overview?hl=fr)

## Generate Refresh token
- Set environment variables `GOOGLE_API_KEY` `GOOGLE_API_SECRET` `GOOGLE_OAUTH_REDIRECT_URL`
- Launch unit test `go test . -v TestGetOauthTokens` to display OAuth install url
- Open link & validate process to get OAuth authorization code
- Launch again unit test `go test . -v TestGetOauthTokens` with auth code in `GOOGLE_OAUTH_CODE` environment variable to display refresh token

## Usage

```go
package main

import (
	"github.com/orixa-group/dv360"
	"log"
	"time"
)

func main() {
	dv360.WithApiKeys("API_KEY", "API_SECRET")
	dv360.WithRefreshToken("REFRESH_TOKEN")

	metrics, err := dv360.GetAccountReport("1234567", time.Now().AddDate(0, -1, 0), time.Now())
	if nil != err {
		log.Fatal(err)
	}
	for _, m := range metrics {
		// process 
	}
}


```

See unit tests for more usages :)
