package shakespeare

import "testing"

//Think of possible ways to use this here: https://pkg.go.dev/github.com/stretchr/testify/assert
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
