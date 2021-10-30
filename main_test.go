package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	big "github.com/allegro/bigcache/v3"
)

func TestGetEmpty(t *testing.T) {
	cache, _ = big.NewBigCache(big.DefaultConfig(1 * time.Minute))
	req := httptest.NewRequest(http.MethodGet, ROOM+"A", nil)
	w := httptest.NewRecorder()
	ffHandler(ROOM)(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "body not found\n" {
		t.Errorf("expected ABC got %v", string(data))
	}
}

func TestPostGet(t *testing.T) {
	const BODY = "AAAAAA"
	cache, _ = big.NewBigCache(big.DefaultConfig(1 * time.Minute))
	ss := strings.NewReader(BODY)
	req := httptest.NewRequest(http.MethodPost, ROOM+"A", ss)
	w := httptest.NewRecorder()
	ffHandler(ROOM)(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data) != "ok" {
		t.Errorf("expected ABC got %v", string(data))
	}

	req = httptest.NewRequest(http.MethodGet, ROOM+"A", nil)
	w = httptest.NewRecorder()
	ffHandler(ROOM)(w, req)
	res2 := w.Result()
	defer res2.Body.Close()
	data2, err := io.ReadAll(res2.Body)
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
	if string(data2) != BODY {
		t.Errorf("expected %s got %v", BODY, string(data))
	}
}
