package anime

type episodeCallback func(u User) any

type Episode struct {
	fn         episodeCallback
	workersMap map[TraceId]*worker
}

func NewEpisode(fn episodeCallback) *Episode {
	ep := Episode{fn: fn}
	ep.workersMap = make(map[TraceId]*worker)
	return &ep
}

func (e *Episode) enroll(w *worker) {
	if e.workersMap[w.traceId] != nil {
		// 系統設計錯誤，直接 crash program
		panic("episode's worker duplicated!")
	}
	e.workersMap[w.traceId] = w
}

func (e *Episode) remove(w *worker) {
	delete(e.workersMap, w.traceId)
}

func (e *Episode) executeFn(w *worker) (any, error) {
	data := e.fn(w)
	err, isErr := data.(error)
	if isErr {
		return nil, err
	}
	return data, nil
}

func (e *Episode) ShowLog(traceId int) (log []byte) {
	panic("Implement Me!")
	return
}

func (e *Episode) Failing() (user []User) {
	panic("Implement Me!")
	return
}

func (e *Episode) Status() (user []User) {
	panic("Implement Me!")
	return
}
