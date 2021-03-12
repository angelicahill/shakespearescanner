package shakespeare

import (
	"github.com/stretchr/testify/require"
	"testing"
)

//Think of possible ways to use this here: https://pkg.go.dev/github.com/stretchr/testify/assert
//assert.Error and assert.NotError
func TestGetPlayHTTPError(t *testing.T) {
	_, err := getPlay("")
	require.Error(t, err, "returned Status 404")
}

func TestGetPlaySuccess(t *testing.T) {
	p, err := getPlay(plays["hamlet"])
	if err != nil {
		t.Error("Unexpected Error.", err)
	}
	if p.TITLE == "" {
		t.Error("expected Title. Got Empty Value.")
	}
}
