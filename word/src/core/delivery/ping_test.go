package delivery

import (
	"github.com/FPNL/i18n-town/src/tools"
	"github.com/go-playground/assert/v2"
	"net/http"
	"testing"

	mock_service "github.com/FPNL/i18n-town/src/mocks/service"

	"github.com/golang/mock/gomock"
)

func pingService(t *testing.T) *mock_service.MockIPingService {
	mockCtl := gomock.NewController(t)
	m := mock_service.NewMockIPingService(mockCtl)
	return m
}
func TestPing(t *testing.T) {
	m := pingService(t)
	m.EXPECT().Pong().Return("pong", nil)

	hd := Ping(m)
	path := "/"
	r := tools.SetupRouter(path, hd.PinPong)
	w := tools.ServeHTTP(r, http.MethodGet, path, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
