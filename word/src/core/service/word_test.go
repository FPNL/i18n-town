package service

import (
	"github.com/FPNL/i18n-town/src/core/entity"
	mock_repository "github.com/FPNL/i18n-town/src/mocks/repository"
	"github.com/golang/mock/gomock"
	"testing"
)

func wordRepository(t *testing.T) *mock_repository.MockIWordRepo {
	mockCtl := gomock.NewController(t)
	m := mock_repository.NewMockIWordRepo(mockCtl)
	return m
}

func TestWordService_AddOneWord(t *testing.T) {
	m := wordRepository(t)
	m.EXPECT().Insert(gomock.Any(), gomock.Any()).Return(nil)
	serv := Word(m)
	err := serv.AddManyWords([]entity.Word{
		{},
	})

	if err != nil {
		t.Fatalf("錯: %s", err.Error())
	}
}

// 目前沒什麼邏輯，時間緊迫，先不測後面
