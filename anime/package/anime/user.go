package anime

import (
	"fmt"
	"log"

	"github.com/FPNL/anime/utils"
	"github.com/google/uuid"
)

type TraceId uuid.UUID //  [16]byte = 128bits

type worker struct {
	isStop         bool
	isSkip         bool
	watchingEp     int
	traceId        TraceId
	watchingSeries string
	data           any
	err            error
	series         *Series
}

type User interface {
	Data() any
	WatchingEp() int
	TraceId() TraceId
	StopFor(err error) error
	SkipTo(number int, data any) any
}

func newWorker(s *Series, data any) worker {
	id := TraceId(uuid.New())
	return worker{
		traceId: id,
		data:    data,
		series:  s,
	}
}

func (w *worker) SkipTo(n int, data any) any {
	if w.series.isOverBound(n) {
		return utils.ErrHappened("Series %s not include ep%d (engineers start from 0)", w.series.Name, n)
	}
	w.watchingEp = n
	w.isSkip = true
	return data
}

// Data here is copy data
func (w *worker) Data() any {
	return w.data
}

func (w *worker) WatchingEp() int {
	return w.watchingEp
}

func (w *worker) TraceId() TraceId {
	return w.traceId
}

func (w *worker) StopFor(err error) error {
	w.err = err
	w.isStop = true
	return err
}

func (w *worker) Play(series *Series, i int) {
	var err error
	episodes := series.episodes

	for w.watchingEp = i; w.watchingEp < len(episodes); w.watchingEp++ {
		ep := episodes[w.watchingEp]
		ep.enroll(w)

		w.data, err = ep.executeFn(w)
		if err != nil {
			_ = w.StopFor(err)
		}

		if w.isStop {
			log.Println(err)
			// worker wouldn't disappear for ep enrolled it
			break
		} else if w.isSkip {
			fmt.Println("worker skip to", w.watchingSeries, w.watchingEp)
			w.watchingEp = w.watchingEp - 1 // for loop will automatically plus one, so here should minus one
			w.isSkip = false
		}

		ep.remove(w)
	}
}
