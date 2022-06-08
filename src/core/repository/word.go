package repository

import (
	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/lib/idatabase"
	"github.com/FPNL/i18n-town/src/lib/ierror"
)

type Interface_word interface {
	Insert(*entity.Word) error
	DeleteByIds([]int) error
	Truncate() error
	SelectAll() ([]entity.Word, error)
	UpdateWords([]entity.Word) error
}

type Struct_word struct{}

var wordRepo = Struct_word{}

func WordRepo() Interface_word {
	// return &Struct_word{}
	// or
	return &wordRepo
}

func (r *Struct_word) Insert(w *entity.Word) error {
	db := idatabase.ConnectDB()
	if _, ok := db.Words.Exist(w); ok {
		return ierror.NewValidateErr("word duplicate")
	}
	return db.Words.Insert(*w)
}

func (r *Struct_word) DeleteByIds(ids []int) error {
	db := idatabase.ConnectDB()

	for _, id := range ids {
		i, ok := db.Words.Exist(&entity.Word{Id: id})
		if ok {
			err := db.Words.DeleteByIndex(i)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (r *Struct_word) Truncate() error {
	db := idatabase.ConnectDB()
	db.Words.Truncate()
	return nil
}

func (r *Struct_word) SelectAll() ([]entity.Word, error) {
	db := idatabase.ConnectDB()
	return db.Words.Columns, nil
}

func (r *Struct_word) UpdateWords(w []entity.Word) error {
	db := idatabase.ConnectDB()
	ws := make(map[int]entity.Word)
	for _, word := range w {
		i, ok := db.Words.Exist(&word)
		if !ok {
			return ierror.NewValidateErr("word not exits")
		}
		ws[i] = word
	}
	for index, word := range ws {
		err := db.Words.UpdateByIndex(index, word)
		if err != nil {
			return err
		}
	}
	return nil
}
