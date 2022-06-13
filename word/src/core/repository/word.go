package repository

import (
	"context"
	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/core/model"
	"github.com/FPNL/i18n-town/src/lib/ierror"
	"github.com/go-redis/redis/v9"
)

type IWordRepo interface {
	Insert(context.Context, *entity.Word) error
	DeleteByIds(context.Context, []int) error
	Truncate(context.Context) error
	SelectAll(context.Context) ([]entity.Word, error)
	UpdateWords(context.Context, []entity.Word) error
}

type wordRepo struct {
	rcache    redis.Cmdable
	wordModel model.IWord
}

var singleWord = wordRepo{}

func Word(wordModel model.IWord, rcache redis.Cmdable) IWordRepo {
	singleWord.rcache = rcache
	singleWord.wordModel = wordModel
	return &singleWord
}

func (r *wordRepo) Insert(ctx context.Context, w *entity.Word) error {
	if ok, err := r.wordModel.Exist(ctx, w); err != nil {
		return err
	} else if ok {
		return ierror.NewValidateErr("word duplicate")
	}

	return r.wordModel.Insert(ctx, *w)
}

func (r *wordRepo) DeleteByIds(ctx context.Context, ids []int) error {
	for _, id := range ids {
		err := r.wordModel.DeleteById(ctx, id)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *wordRepo) Truncate(ctx context.Context) error {
	return r.wordModel.Truncate(ctx)
}

func (r *wordRepo) SelectAll(ctx context.Context) ([]entity.Word, error) {
	return r.wordModel.SelectAll(ctx)
}

func (r *wordRepo) UpdateWords(ctx context.Context, w []entity.Word) error {
	for _, word := range w {
		ok, err := r.wordModel.Exist(ctx, &word)
		if err != nil {
			return err
		} else if !ok {
			return ierror.NewValidateErr("word not exits")
		}
	}
	for _, word := range w {
		err := r.wordModel.UpdateById(ctx, &word)
		if err != nil {
			return err
		}
	}
	return nil
}
