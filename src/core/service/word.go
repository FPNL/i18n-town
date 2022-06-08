package service

import (
	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/core/repository"
)

type Interface_wordService interface {
	AddOneWord(*entity.Word) error
	AddManyWords([]entity.Word) error
	DeleteOneWord(int) error
	DeleteManyWord([]int) error
	DeleteAll() error
	UpdateOneWord(int, string) error
	UpdateManyWords(map[int]string) error
	FetchAllWords() ([]entity.Word, error)
}

type struct_wordService struct{}

var wordService = struct_wordService{}

func WordService() Interface_wordService {
	return &wordService
}

func (service *struct_wordService) FetchAllWords() ([]entity.Word, error) {
	return repository.WordRepo().SelectAll()
}

func (service *struct_wordService) AddOneWord(w *entity.Word) error {
	return repository.WordRepo().Insert(w)
}

func (service *struct_wordService) AddManyWords(ww []entity.Word) error {
	for _, w := range ww {
		err := repository.WordRepo().Insert(&w)
		if err != nil {
			return err
		}
	}

	return nil
}

func (service *struct_wordService) DeleteOneWord(i int) error {
	return repository.WordRepo().DeleteByIds([]int{i})
}

func (service *struct_wordService) DeleteManyWord(i []int) error {
	return repository.WordRepo().DeleteByIds(i)
}

func (service *struct_wordService) DeleteAll() error {
	return repository.WordRepo().Truncate()
}

func (service *struct_wordService) UpdateOneWord(id int, s string) error {
	w := []entity.Word{
		{Id: id, Word: s},
	}

	return repository.WordRepo().UpdateWords(w)
}

func (service *struct_wordService) UpdateManyWords(w map[int]string) error {
	ww := make([]entity.Word, 0)
	for i, s := range w {
		ww = append(ww, entity.Word{Id: i, Word: s})
	}
	return repository.WordRepo().UpdateWords(ww)
}
