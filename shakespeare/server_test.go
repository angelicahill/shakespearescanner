package shakespeare

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppendixHandler(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "https://example.com/", nil)
	AppendixHandler(writer, request)
	if writer.Code != 200 {
		t.Errorf("Invalid or Unexpected Code - value recieve %v", writer.Code)
	}
	if writer.Body.Len() == 0 {
		t.Error("Body is Empty.")
	}
}

func TestRun2(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "https://example.com?play=kinglear&word=fair", nil)
	Run2(writer, request)
	if writer.Code != 200 {
		t.Errorf("Invalid or Unexpected Code - value recieve %v", writer.Code)
	}
	if writer.Body.Len() == 0 {
		t.Error("Body is Empty.")
	}
	if !strings.Contains(writer.Body.String(), "fair showed up in your play 3 times in ACT III") {
		t.Fatal("expected the word fair to appear 3 times in ACT III")
	}
}
func TestRun2404(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "https://example.com?play=melon&word=fair", nil)
	Run2(writer, request)
	if writer.Code != 404 {
		t.Errorf("Did not get 404 Code as expected - value recieve %v", writer.Code)
	}
}
