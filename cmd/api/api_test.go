package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/drizlye0/GopherSocial/internal/ratelimiter"
)

func TestRateLimiterMiddleware(t *testing.T) {
	cfg := config{
		rateLimiter: ratelimiter.Config{
			RequestPerTimeFrame: 20,
			TimeFrame:           time.Second * 5,
			Enabled:             true,
		},
		addr: ":8080",
	}

	app := newTestApplication(t, cfg)
	ts := httptest.NewServer(app.mount())
	defer ts.Close()

	client := &http.Client{}
	mockIP := "192.168.1.1"
	marginOfError := 2

	for i := 0; i < cfg.rateLimiter.RequestPerTimeFrame+marginOfError; i++ {
		req, err := http.NewRequest(http.MethodGet, ts.URL+"/v1/health", nil)
		if err != nil {
			t.Fatalf("could not create the request: %v", err)
		}

		req.Header.Set("X-Forwarded-For", mockIP)

		resp, err := client.Do(req)
		if err != nil {
			t.Fatalf("coudl not send the request: %v", err)
		}
		defer resp.Body.Close()

		if i < cfg.rateLimiter.RequestPerTimeFrame {
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v: ", resp.StatusCode)
			}
		} else {
			if resp.StatusCode != http.StatusTooManyRequests {
				t.Errorf("expected status Too Many Request; got: %v", resp.StatusCode)
			}
		}

	}
}
