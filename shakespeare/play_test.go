package shakespeare

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPlayHTTPError(t *testing.T) {
	_, err := getPlay("")
	require.Error(t, err, "returned Status 404")
}

func TestGetPlaySuccess(t *testing.T) {
	p, err := getPlay(plays["hamlet"])
	require.NoError(t, err, "unexpected error")
	assert.NotEmpty(t, p.TITLE)
}
