package anime

import (
	"golang.org/x/exp/maps"
	"strconv"
	"strings"

	"github.com/FPNL/anime/utils"
)

// TODO: If under 64 bytes return value
// TODO: data align

func NewDashboard(o Options) *Dashboard {
	d := &Dashboard{}
	d.series = make(map[string]*Series)
	d.option = o
	return d
}

type Options struct{}
type Dashboard struct {
	series map[string]*Series
	option Options
}

type Status struct{}

func (d *Dashboard) NewSeries(s string) *Series {
	var sr = &Series{
		Name: s,
	}
	d.series[s] = sr
	return sr
}

func (d *Dashboard) MakeAllRelease() {
	for _, series := range d.series {
		series.release = true
	}
}

// Play series and send data, s can be "Spy and family" or
// "Spy and family/3" to specify which number of ep
func (d *Dashboard) Play(name string, data any) error {
	name, number, err := parseName(name)
	if err != nil {
		return err
	}

	series, ok := d.series[name]
	if !ok { // 代表不存在
		return utils.ErrHappened("Series %s Not Exist", name)
	}

	if series.isOverBound(number) {
		return utils.ErrHappened("Series %s not include ep%d (engineers start from 0)", name, number)
	}

	if !series.release {
		return utils.ErrHappened("Series %s is not release", name)
	}

	// Story Inverse... Here!!
	// Now the vision switch to user who watch anime

	//TODO: Figure out is this in heap or stack
	// if is in heap, how to covert it into stack
	w := newWorker(series, data)

	w.Play(series, number)

	return nil
}

func parseName(name string) (string, int, error) {
	switch s := strings.Split(name, "/"); len(s) {
	case 1:
		return name, 0, nil
	case 2:
		number, err := strconv.Atoi(s[1])
		if err != nil {
			return "", 0, err
		}
		return s[0], number, nil
	default:
		return "", 0, utils.ErrHappened("Everything is wrong")
	}
}

func (d *Dashboard) List() []*Series {
	return maps.Values(d.series)
}

func (d *Dashboard) PauseAll() error {
	panic("Implement Me!")
	return nil
}

func (d *Dashboard) Pause(series Series) error {
	panic("Implement Me!")
	return nil
}

func (d *Dashboard) PlayOnce(series Series, data any) error {
	panic("Implement Me!")
	return nil
}

func (d *Dashboard) Remove(series Series) error {
	panic("Implement Me!")
	return nil
}
