package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/drizlye0/GopherSocial/internal/store"
	"github.com/drizlye0/GopherSocial/internal/store/cache"
	"go.uber.org/zap"
)

func newTestApplication(t *testing.T) *application {
	t.Helper()

	logger := zap.NewNop().Sugar()
	store := store.NewMockStore()
	cacheStorage := cache.NewMockStore()

	return &application{
		logger:       logger,
		store:        store,
		cacheStorage: cacheStorage,
	}
}

func executeRequest(req *http.Request, mux http.Handler) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	mux.ServeHTTP(rr, req)

	return rr
}
