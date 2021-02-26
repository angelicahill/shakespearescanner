package shakespeare

import "testing"

//To run in terminal - go test -v ./... - will run all test files I have in all sub-directories.
func TestGetPlayHTTPError(t *testing.T) {
	_, err := getPlay("")
	if err == nil {
		t.Error("Expecting Error - error not returned.")
	}
}

func TestGetPlaySuccess(t *testing.T) {
	p, err := getPlay(plays["hamlet"])
	if err != nil {
		t.Error("Unexpected Error.", err)
	}
	if p.TITLE == "" {
		t.Error("Expected Title. Got Empty Value.")
	}
}
