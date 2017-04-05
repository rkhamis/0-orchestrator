package tools

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPError(t *testing.T) {
	resp := &http.Response{
		StatusCode: http.StatusInternalServerError,
	}
	err := NewHTTPError(resp, "this is an %s", "error")
	assert.Equal(t, err.Error(), "this is an error")
	assert.NotNil(t, err.resp)
	assert.Equal(t, err.resp.StatusCode, http.StatusInternalServerError)
}
