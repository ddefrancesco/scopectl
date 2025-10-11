package handlers

import (
	"encoding/json"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	etxClient "github.com/ddefrancesco/scopectl/restclient"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestGotoCommandHandler_Success(t *testing.T) {
	t.Skip()
	// Start a test HTTP server that returns the expected successful response
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		_ = json.NewEncoder(w).Encode(etxClient.ScopeResponse{Code: 202, Response: "ok"})
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	host, port, _ := net.SplitHostPort(u.Host)

	viper.Set("environment", "test")
	viper.Set("environments.test.url", u.Scheme+"://"+host)
	viper.Set("environments.test.port", port)

	pmap := map[string]string{"ngc": "1234"}
	resp, err := GotoCommandHandler(pmap)
	assert.NoError(t, err)
	if assert.NotNil(t, resp) {
		assert.Equal(t, 202, resp.Code)
		assert.Equal(t, "ok", resp.Response)
	}
}

func TestGotoCommandHandler_Fail(t *testing.T) {
	t.Skip()
	// Start a test HTTP server that returns a server error
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": "fail"})
	}))
	defer srv.Close()

	u, _ := url.Parse(srv.URL)
	host, port, _ := net.SplitHostPort(u.Host)

	viper.Set("environment", "test")
	viper.Set("environments.test.url", u.Scheme+"://"+host)
	viper.Set("environments.test.port", port)

	pmap := map[string]string{"ngc": "1234"}
	resp, err := GotoCommandHandler(pmap)
	assert.Error(t, err)
	assert.Nil(t, resp)
}
