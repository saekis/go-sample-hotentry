package client_test

import (
	"net/http"
	"testing"

	"github.com/saekis/go-sample-hotentry/client"
	"github.com/saekis/go-sample-hotentry/config"
	"github.com/saekis/go-sample-hotentry/model/hatena"
)

type mockParser struct{}

func (mockParser) ParseToEntryList(*http.Response, int) (*hatena.EntryList, error) {
	return &hatena.EntryList{
		hatena.Entry{},
		hatena.Entry{},
		hatena.Entry{},
	}, nil
}

type mockHTTPClient struct{}

func (mockHTTPClient) Get(url string) (*http.Response, error) {
	return &http.Response{}, nil
}

func TestGetHotentryList_Success(t *testing.T) {
	mockhatenaclient, err := client.NewHatenaClient(mockHTTPClient{}, &config.Hatena{}, mockParser{})
	if err != nil {
		t.Errorf("got: %v\nwant: %v", err, nil)
	}

	_, err = mockhatenaclient.GetHotentryList()
	if err != nil {
		t.Errorf("got: %v\nwant: %v", err, nil)
	}
}
