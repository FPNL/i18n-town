package idatabase

import (
	"github.com/FPNL/i18n-town/src/core/entity"
	"sync"
)

type ID int
type Columns []entity.Word
type WordsTable struct {
	Mux     sync.RWMutex
	Columns Columns
}

var autoIncrementId ID = 0

func (wt *WordsTable) Insert(w entity.Word) error {
	wt.Mux.Lock()
	defer wt.Mux.Unlock()
	w.Id = int(generateId())
	wt.Columns = append(wt.Columns, w)
	return nil
}

func generateId() ID {
	autoIncrementId++
	return autoIncrementId
}

func (wt *WordsTable) DeleteByIndex(i int) error {
	wt.Mux.Lock()
	defer wt.Mux.Unlock()
	wt.Columns = sliceRemove(wt.Columns, i)
	return nil
}

func (wt *WordsTable) Exist(w *entity.Word) (int, bool) {
	wt.Mux.RLock()
	defer wt.Mux.RUnlock()
	for i, column := range wt.Columns {
		if column.Id == w.Id ||
			column.Tag == w.Tag && column.Lang == w.Lang {
			return i, true
		}
	}

	return -1, false
}

func (wt *WordsTable) UpdateByIndex(i int, w entity.Word) error {
	wt.Mux.Lock()
	defer wt.Mux.Unlock()
	// FIXME: It's impossible in real
	if w.Lang != "" {
		wt.Columns[i].Lang = w.Lang
	}
	if w.Tag != "" {
		wt.Columns[i].Tag = w.Tag
	}
	if w.Word != "" {
		wt.Columns[i].Word = w.Word
	}

	return nil
}

func (wt *WordsTable) Truncate() {
	wt.Mux.Lock()
	defer wt.Mux.Unlock()
	wt.Columns = make(Columns, 0, 10)
}

func sliceRemove(column Columns, s int) Columns {
	return append(column[:s], column[s+1:]...)
}
