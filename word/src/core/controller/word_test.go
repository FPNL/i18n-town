package controller

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/FPNL/i18n-town/src/core/entity"
	mock_service "github.com/FPNL/i18n-town/src/mocks/service"
	"github.com/FPNL/i18n-town/src/tools"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func wordService(t *testing.T) *mock_service.MockIWordService {
	mockCtl := gomock.NewController(t)
	m := mock_service.NewMockIWordService(mockCtl)
	return m
}

func TestWordHandler_FetchAllWords(t *testing.T) {
	m := wordService(t)
	m.EXPECT().FetchAllWords().Return([]entity.Word{}, nil)

	hd := Word(m)
	path := "/"
	r := tools.SetupRouter(path, hd.FetchAllWords)
	w := tools.ServeHTTP(r, http.MethodGet, path, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "[]", w.Body.String())
}

func TestWordHandler_AddOneWord(t *testing.T) {
	m := wordService(t)
	m.EXPECT().AddOneWord(gomock.Any()).Return(nil)

	hd := Word(m)
	path := "/"
	mockWord := entity.Word{Id: 0, Lang: "chi", Word: "go", Tag: "go"}
	query := fmt.Sprintf("?tag=%q&lang=%q&word=%q", mockWord.Tag, mockWord.Lang, mockWord.Word)
	router := tools.SetupRouter(path, hd.AddOneWord)
	res := tools.ServeHTTP(router, http.MethodGet, path+query, nil)

	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}

func TestWordHandler_AddManyWords(t *testing.T) {
	m := wordService(t)
	m.EXPECT().AddManyWords(gomock.Any()).Return(nil)
	hd := Word(m)

	path := "/"
	reqBody := tools.MakeJsonBody([]map[string]string{
		{
			"tag":  "hello",
			"lang": "eng",
			"word": "hello",
		},
		{
			"tag":  "hello2",
			"lang": "eng",
			"word": "hello2",
		},
	})
	router := tools.SetupRouter(path, hd.AddManyWords)
	res := tools.ServeHTTP(router, http.MethodGet, path, reqBody)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}

func TestWordHandler_UpdateOneWord(t *testing.T) {
	m := wordService(t)
	id, word := 1, "hello"
	m.EXPECT().UpdateOneWord(id, word).Return(nil)
	hd := Word(m)

	path := "/"
	query := fmt.Sprintf(`?id=%d&word=%s`, id, word)
	router := tools.SetupRouter(path, hd.UpdateOneWord)
	res := tools.ServeHTTP(router, http.MethodGet, path+query, nil)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}

func TestWordHandler_UpdateManyWords(t *testing.T) {
	m := wordService(t)
	m.EXPECT().UpdateManyWords(gomock.Any()).Return(nil)
	hd := Word(m)

	path := "/"
	body := tools.MakeJsonBody(map[int]string{
		1: "no",
		2: "no2",
	})
	router := tools.SetupRouter(path, hd.UpdateManyWords)
	res := tools.ServeHTTP(router, http.MethodGet, path, body)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}

func TestWordHandler_DeleteOneWord(t *testing.T) {
	m := wordService(t)
	id := 1
	m.EXPECT().DeleteOneWord(id).Return(nil)
	hd := Word(m)

	path := "/"
	query := fmt.Sprintf(`?id=%d`, id)
	router := tools.SetupRouter(path, hd.DeleteOneWord)
	res := tools.ServeHTTP(router, http.MethodGet, path+query, nil)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}

func TestWordHandler_DeleteManyWord(t *testing.T) {
	m := wordService(t)
	m.EXPECT().DeleteManyWord(gomock.Any()).Return(nil)
	hd := Word(m)

	path := "/"
	body := tools.MakeJsonBody([]int{1, 2, 3})
	router := tools.SetupRouter(path, hd.DeleteManyWord)
	res := tools.ServeHTTP(router, http.MethodGet, path, body)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}

func TestWordHandler_DeleteAll(t *testing.T) {
	m := wordService(t)
	m.EXPECT().DeleteAll().Return(nil)
	hd := Word(m)

	path := "/"
	router := tools.SetupRouter(path, hd.DeleteAll)
	res := tools.ServeHTTP(router, http.MethodGet, path, nil)
	assert.Equal(t, http.StatusOK, res.Code)
	assert.Equal(t, "ok", res.Body.String())
}
