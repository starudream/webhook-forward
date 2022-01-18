package webhookForward

import (
	"net/http"
	"testing"

	"github.com/go-sdk/lib/testx"
)

func TestSend(t *testing.T) {
	w := handle(http.MethodGet, "/send?message=hello+world", nil, nil)
	testx.RequireEqual(t, http.StatusOK, w.Code)
}
