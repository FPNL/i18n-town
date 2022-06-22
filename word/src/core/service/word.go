package service

import (
	"context"
	"gorm.io/gorm"

	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/core/repository"
	"github.com/FPNL/i18n-town/src/lib/ierror"

	pb "github.com/FPNL/i18n-town/src/lib/igrpc"
)

type IWord interface {
	AddCommittedWords(context.Context, *pb.User, []entity.CommittedWord) error
	FetchCommittedWords(context.Context, *pb.User, *entity.SearchCondition_CommittedWord) ([]entity.CommittedWord, uint, error)
	UpdateCommittedWords(context.Context, *pb.User, map[uint]string) error
	DeleteCommittedWords(context.Context, *pb.User, []uint) error
	//
	//FetchStageWords(context.Context, *entity.User, *entity.SearchCondition_StageWord) ([]entity.StageWord, error)
	//AdviseWords(context.Context, *entity.User, []entity.StageWord) error
	//CommitWords(context.Context, *entity.User, []entity.StageWord) error
	//DiscardStageWords(context.Context, *entity.User, []entity.StageWord) error
}

type wordService struct {
	wordRepo repository.IWord
}

var singletonWord = wordService{}

func Word(wordRepo repository.IWord) IWord {
	singletonWord.wordRepo = wordRepo
	return &singletonWord
}

func (sv *wordService) AddCommittedWords(ctx context.Context, user *pb.User, words []entity.CommittedWord) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		if user.Role > pb.Role_Maintainer {
			return ierror.NewValidateErr("身份不夠高")
		}

		for i := range words {
			words[i].CommitUser_ID = uint(user.GetID())
		}

		return sv.wordRepo.AddCommittedWords(words)
	}
}

// FetchCommittedWords 任何角色都可以呼叫
func (sv *wordService) FetchCommittedWords(ctx context.Context, user *pb.User, conditions *entity.SearchCondition_CommittedWord) ([]entity.CommittedWord, uint, error) {
	select {
	case <-ctx.Done():
		return nil, 0, nil
	default:
		conditions.CommittedWord.Organize_ID = uint(user.Organize)
		count, err := sv.wordRepo.Count(&conditions.CommittedWord)
		if err != nil {
			return nil, 0, err
		}
		words, err := sv.wordRepo.FetchCommittedWords(&conditions.CommittedWord, &conditions.Pagination)
		if err != nil {
			return nil, 0, err
		}

		return words, count, nil
	}
}

func (sv *wordService) UpdateCommittedWords(ctx context.Context, user *pb.User, words map[uint]string) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		if user.Role > pb.Role_Maintainer {
			return ierror.NewValidateErr("身份不夠高")
		}

		var committedWords = make([]entity.CommittedWord, 0, len(words))
		for id, str := range words {
			word := entity.CommittedWord{
				StageWord: entity.StageWord{
					Model: gorm.Model{
						ID: id,
					},
					Word: str,
				},
				CommitUser_ID: uint(user.GetID()),
			}

			committedWords = append(committedWords, word)
		}

		return sv.wordRepo.UpdateCommittedWords(committedWords)
	}
}

func (sv *wordService) DeleteCommittedWords(ctx context.Context, user *pb.User, ids []uint) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		if user.Role > pb.Role_Maintainer {
			return ierror.NewValidateErr("身份不夠高")
		}

		return sv.wordRepo.DeleteCommittedWords(ids)
	}
}

//func (service *wordService) AddOneWord(ctx context.Context, w *entity.Word) error {
//	return service.wordRepo.Insert(context.Background(), w)
//}
//
//func (service *wordService) AddManyWords(ctx context.Context, ww []entity.Word) error {
//	for _, w := range ww {
//		err := service.wordRepo.Insert(context.Background(), &w)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}
//
//func (service *wordService) DeleteOneWord(ctx context.Context, i int) error {
//	return service.wordRepo.DeleteByIds(context.Background(), []int{i})
//}
//
//func (service *wordService) DeleteManyWord(ctx context.Context, i []int) error {
//	return service.wordRepo.DeleteByIds(context.Background(), i)
//}
//
//func (service *wordService) DeleteAll(ctx context.Context) error {
//	return service.wordRepo.Truncate(context.Background())
//}
//
//func (service *wordService) UpdateOneWord(ctx context.Context, id int, s string) error {
//	w := []entity.Word{
//		{Id: id, Word: s},
//	}
//	return service.wordRepo.UpdateWords(context.Background(), w)
//}
//
//func (service *wordService) UpdateManyWords(ctx context.Context, w map[int]string) error {
//	ww := make([]entity.Word, 0)
//	for i, s := range w {
//		ww = append(ww, entity.Word{Id: i, Word: s})
//	}
//	return service.wordRepo.UpdateWords(context.Background(), ww)
//}
