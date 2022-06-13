package repository

import (
	"context"
	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/lib/icache"
	mock_model "github.com/FPNL/i18n-town/src/mocks/model"
	"github.com/go-redis/redis/v9"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func wordModel(t *testing.T) *mock_model.MockIWord {
	mockCtl := gomock.NewController(t)
	m := mock_model.NewMockIWord(mockCtl)
	return m
}

func mockCache(t *testing.T) redis.Cmdable {
	// TODO 需要用 gomock 創建假的，目前先用真的 redis
	return icache.Connect()
}

func TestWordRepo_Insert(t *testing.T) {
	ctx := context.Background()
	word := entity.Word{}
	m := wordModel(t)
	m.EXPECT().Insert(ctx, word).Return(nil)
	m.EXPECT().Exist(ctx, &word).Return(false, nil)

	cache := mockCache(t)

	repo := Word(m, cache)
	err := repo.Insert(ctx, &word)
	if err != nil {
		t.Fatalf("錯 : %s", err.Error())
	}
}

func TestWordRepo_SelectAll(t *testing.T) {
	ctx := context.Background()
	word := make([]entity.Word, 1)

	m := wordModel(t)
	m.EXPECT().SelectAll(ctx).Return(word, nil)

	cache := mockCache(t)

	repo := Word(m, cache)
	res, err := repo.SelectAll(ctx)
	if err != nil {
		t.Fatalf("錯 : %s", err.Error())
	}
	assert.Equal(t, res, word)
}

func TestWordRepo_UpdateWords(t *testing.T) {
	ctx := context.Background()
	w := entity.Word{Id: 1}
	word := []entity.Word{w}

	m := wordModel(t)
	m.EXPECT().UpdateById(ctx, &w).Return(nil)
	m.EXPECT().Exist(ctx, &w).Return(true, nil)

	cache := mockCache(t)

	repo := Word(m, cache)
	err := repo.UpdateWords(ctx, word)
	if err != nil {
		t.Fatalf("錯 : %s", err.Error())
	}
}
