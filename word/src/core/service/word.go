package service

import (
	"context"
	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/core/repository"
)

type IWordService interface {
	AddOneWord(*entity.Word) error
	AddManyWords([]entity.Word) error
	DeleteOneWord(int) error
	DeleteManyWord([]int) error
	DeleteAll() error
	UpdateOneWord(int, string) error
	UpdateManyWords(map[int]string) error
	FetchAllWords() ([]entity.Word, error)
}

type wordService struct {
	wordRepo repository.IWordRepo
}

var singletonWord = wordService{}

func Word(wordRepo repository.IWordRepo) IWordService {
	singletonWord.wordRepo = wordRepo
	return &singletonWord
}

func (service *wordService) FetchAllWords() ([]entity.Word, error) {
	return service.wordRepo.SelectAll(context.Background())
}

func (service *wordService) AddOneWord(w *entity.Word) error {
	return service.wordRepo.Insert(context.Background(), w)
}

func (service *wordService) AddManyWords(ww []entity.Word) error {
	for _, w := range ww {
		err := service.wordRepo.Insert(context.Background(), &w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *wordService) DeleteOneWord(i int) error {
	return service.wordRepo.DeleteByIds(context.Background(), []int{i})
}

func (service *wordService) DeleteManyWord(i []int) error {
	return service.wordRepo.DeleteByIds(context.Background(), i)
}

func (service *wordService) DeleteAll() error {
	return service.wordRepo.Truncate(context.Background())
}

func (service *wordService) UpdateOneWord(id int, s string) error {
	w := []entity.Word{
		{Id: id, Word: s},
	}
	return service.wordRepo.UpdateWords(context.Background(), w)
}

func (service *wordService) UpdateManyWords(w map[int]string) error {
	ww := make([]entity.Word, 0)
	for i, s := range w {
		ww = append(ww, entity.Word{Id: i, Word: s})
	}
	return service.wordRepo.UpdateWords(context.Background(), ww)
}
