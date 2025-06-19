package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drizlye0/GopherSocial/internal/auth"
	"github.com/drizlye0/GopherSocial/internal/store"
	"github.com/drizlye0/GopherSocial/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	store := store.NewMockStore()
	cacheStorage := cache.NewMockStore()
	testAuth := &auth.TestAuthenticator{}

	return &application{
		logger:        logger,
		store:         store,
		cacheStorage:  cacheStorage,
		authenticator: testAuth,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}

func checkTestResponse(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected status code %d, got %d", expected, actual)
	}
}
