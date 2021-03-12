package shakespeare

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestAppendixHandler(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "https://example.com/", nil)
	AppendixHandler(writer, request)
	assert.Equal(t, 200, writer.Code, "Invalid or Unexpected Code - value recieve %v", writer.Code)
	assert.Equal(t, 64, writer.Body.Len(), "Body is Empty.")
}

func TestRun2(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "https://example.com?play=kinglear&word=fair", nil)
	Run2(writer, request)
	assert.Equal(t, 200, writer.Code, "Invalid or Unexpected Code - value recieve %v", writer.Code)
	assert.True(t, strings.Contains(writer.Body.String(), "fair showed up in your play 3 times in ACT III"), "expected the word fair to appear 3 times in ACT III")
}
func TestRun2404(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "https://example.com?play=melon&word=fair", nil)
	Run2(writer, request)
	assert.Equal(t, 404, writer.Code, "Did not get 404 Code as expected - value recieve %v", writer.Code)
}
