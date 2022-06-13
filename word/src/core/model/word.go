package model

import (
	"context"
	"fmt"

	"github.com/FPNL/i18n-town/src/core/entity"
	"github.com/FPNL/i18n-town/src/lib/ierror"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	TableName = "word"
)

type IWord interface {
	Insert(context.Context, entity.Word) error
	DeleteById(context.Context, int) error
	Exist(context.Context, *entity.Word) (bool, error)
	SelectAll(context.Context) ([]entity.Word, error)
	Truncate(context.Context) error
	UpdateById(context.Context, *entity.Word) error
}
type wordModel struct {
	db *pgxpool.Pool
}

var singleWord = wordModel{}

func Word(db *pgxpool.Pool) IWord {
	singleWord.db = db
	return &singleWord
}

func (md *wordModel) SelectAll(ctx context.Context) ([]entity.Word, error) {
	sentence := fmt.Sprintf(`SELECT * FROM %q`, TableName)
	rows, err := md.db.Query(ctx, sentence)
	if err != nil {
		return nil, err
	}
	answer := make([]entity.Word, 0, 10)
	for rows.Next() {
		var w entity.Word
		err := rows.Scan(&w.Id, &w.Tag, &w.Lang, &w.Word)
		if err != nil {
			return nil, err
		}
		answer = append(answer, w)
	}

	return answer, nil
}

func (md *wordModel) DeleteById(ctx context.Context, id int) error {
	sentence := fmt.Sprintf(`DELETE FROM %q WHERE id = %v;`, TableName, id)
	_, err := md.db.Exec(ctx, sentence)
	return err
}

func (md *wordModel) Exist(ctx context.Context, word *entity.Word) (bool, error) {
	sentence := fmt.Sprintf(`SELECT COUNT(1) FROM %q`, TableName)
	if word.Id > 0 {
		sentence = fmt.Sprintf(`%s WHERE id = %d;`,
			sentence, word.Id,
		)
	} else {
		sentence = fmt.Sprintf(`%s WHERE tag = '%s' AND lang = '%s';`,
			sentence, word.Tag, word.Lang,
		)
	}

	var count int
	err := md.db.QueryRow(ctx, sentence).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (md *wordModel) UpdateById(ctx context.Context, w *entity.Word) error {
	sentence := fmt.Sprintf(
		`UPDATE %q SET word = '%s' WHERE id = %d;`,
		TableName, w.Word, w.Id,
	)

	_, err := md.db.Exec(ctx, sentence)
	if err != nil {
		return err
	}

	return nil
}

func (md *wordModel) Truncate(ctx context.Context) error {
	sentence := `TRUNCATE TABLE word RESTART IDENTITY;`
	_, err := md.db.Exec(ctx, sentence)
	return err
}

func (md *wordModel) Insert(ctx context.Context, w entity.Word) error {
	sentence := fmt.Sprintf(
		`INSERT INTO %s(tag, lang, word) VALUES ('%s', '%s', '%s');`,
		TableName, w.Tag, w.Lang, w.Word,
	)

	rows, err := md.db.Exec(ctx, sentence)
	if err != nil {
		return err
	} else if rows.RowsAffected() < 1 {
		return ierror.NewValidateErr("沒寫入？？")
	}

	return nil
}

//
//func (table Word) DeleteById(i int) error {
//	db := idatabase.Connect()
//	wt.Mux.Lock()
//	defer wt.Mux.Unlock()
//	wt.Columns = sliceRemove(wt.Columns, i)
//	return nil
//}
//
//func (table Word) Exist(w *entity.Word) (int, bool) {
//	wt.Mux.RLock()
//	defer wt.Mux.RUnlock()
//	for i, column := range wt.Columns {
//		if column.Id == w.Id ||
//			column.Tag == w.Tag && column.Lang == w.Lang {
//			return i, true
//		}
//	}
//
//	return -1, false
//}
//
//func (table Word) UpdateByIndex(i int, w entity.Word) error {
//	wt.Mux.Lock()
//	defer wt.Mux.Unlock()
//	// FIXME: It's impossible in real
//	if w.Lang != "" {
//		wt.Columns[i].Lang = w.Lang
//	}
//	if w.Tag != "" {
//		wt.Columns[i].Tag = w.Tag
//	}
//	if w.Word != "" {
//		wt.Columns[i].Word = w.Word
//	}
//
//	return nil
//}
//
//func (table Word) Truncate() {
//	wt.Mux.Lock()
//	defer wt.Mux.Unlock()
//	wt.Columns = make(Columns, 0, 10)
//}
//
//func sliceRemove(column Columns, s int) Columns {
//	return append(column[:s], column[s+1:]...)
//}
