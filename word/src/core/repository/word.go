package repository

import (
	"errors"
	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/lib/ierror"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

type IWord interface {
	AddCommittedWords([]entity.CommittedWord) error
	FetchCommittedWords(*entity.CommittedWord, *entity.Pagination) ([]entity.CommittedWord, error)
	UpdateCommittedWords([]entity.CommittedWord) error
	Count(*entity.CommittedWord) (uint, error)
	DeleteCommittedWords([]uint) error
}

type wordRepo struct {
	cache redis.Cmdable
	db    *gorm.DB
}

var singleWord = wordRepo{}

func Word(model *gorm.DB, cache redis.Cmdable) IWord {
	singleWord.cache = cache
	singleWord.db = model
	return &singleWord
}

func (r *wordRepo) Count(condition *entity.CommittedWord) (uint, error) {
	tx := r.db.Model(&entity.CommittedWord{})

	// TODO 做 count
	tx.Where(condition)
	//tx.Where("", condition.Lang)
	//tx.Where("", condition.Word)
	//tx.Where("", condition.Tag)
	//tx.Where("", condition.Organize_ID)

	var count int64
	tx.Count(&count)

	return uint(count), tx.Error
}

func (r *wordRepo) AddCommittedWords(words []entity.CommittedWord) error {
	tx := r.db.Create(&words)
	if tx.RowsAffected != int64(len(words)) {
		return ierror.NewServerErr("寫入數量 %d 跟預計寫入數量 %d 不符合", tx.RowsAffected, len(words))
	}
	return tx.Error
}

func (r *wordRepo) FetchCommittedWords(condition *entity.CommittedWord, page *entity.Pagination) ([]entity.CommittedWord, error) {
	var words []entity.CommittedWord
	tx := r.db.Limit(page.Amount).Offset(page.Amount * (page.Page - 1))
	tx.Order(page.Order)
	tx.Find(&words, condition)
	if errors.As(tx.Error, &gorm.ErrRecordNotFound) {
		return words, nil
	}

	return words, tx.Error
}

func (r *wordRepo) UpdateCommittedWords(words []entity.CommittedWord) error {
	for _, word := range words {
		ctx := r.db.Model(&word).Updates(map[string]interface{}{
			"word":        word.Word,
			"commit_user": word.CommitUser_ID,
		})
		if ctx.Error != nil {
			return ctx.Error
		}
	}

	return nil
}

func (r *wordRepo) DeleteCommittedWords(ids []uint) error {
	return r.db.Where("id IN ?", ids).Delete(&entity.CommittedWord{}).Error
}

//func (r *wordRepo) Insert(ctx context.Context, w *entity.CommittedWord) error {
//	if ok, err := r.wordModel.Exist(ctx, w); err != nil {
//		return err
//	} else if ok {
//		return ierror.NewValidateErr("word duplicate")
//	}
//
//	return r.wordModel.Insert(ctx, *w)
//}
//
//func (r *wordRepo) DeleteByIds(ctx context.Context, ids []int) error {
//	for _, id := range ids {
//		err := r.wordModel.DeleteById(ctx, id)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func (r *wordRepo) Truncate(ctx context.Context) error {
//	return r.wordModel.Truncate(ctx)
//}
//
//func (r *wordRepo) SelectAll(ctx context.Context) ([]entity.CommittedWord, error) {
//	return r.wordModel.SelectAll(ctx)
//}
//
//func (r *wordRepo) UpdateWords(ctx context.Context, w []entity.CommittedWord) error {
//	for _, word := range w {
//		ok, err := r.wordModel.Exist(ctx, &word)
//		if err != nil {
//			return err
//		} else if !ok {
//			return ierror.NewValidateErr("word not exits")
//		}
//	}
//	for _, word := range w {
//		err := r.wordModel.UpdateById(ctx, &word)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}
