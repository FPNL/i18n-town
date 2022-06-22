package anime

import (
	"errors"
	"fmt"
	"testing"
)

type fakeS struct {
	i    string
	love string
	you  string
}

func TestQuickStart(t *testing.T) {
	// dashboard -> series -> episode
	a := NewDashboard(Options{}) // global Options
	s := a.NewSeries("Spy and family")
	begin, e1p, e2p := 0, 5, 2

	e1 := NewEpisode(func(user User) any {
		// data transform
		d, ok := user.Data().(int)
		if !ok {
			return user.StopFor(errors.New("this is not int"))
		}

		d += e1p

		// how can user skip episode?
		if d > 10 {
			// 第二種寫法要考量的是
			return user.SkipTo(3, d-10) // what is five?
		}

		return d
	})
	e2 := NewEpisode(func(u User) any {
		d, ok := u.Data().(int)
		if !ok {
			return u.StopFor(errors.New("data is not int"))
		}
		d += e2p
		return d
	})
	e3 := NewEpisode(func(u User) any {
		answer, ok := u.Data().(int)
		if !ok {
			return u.StopFor(errors.New("not int"))
		}

		if answer != (begin + e1p + e2p) {
			t.Fatalf("%d + %d + %d = %d ????", begin, e1p, e2p, answer)
		}

		return nil
	})
	episodes := []*Episode{
		e1, e2, e3,
	}
	s.AddEpisode(episodes)
	a.MakeAllRelease() // will make all series release

	data := "begin"
	for i := 0; i < 10; i++ {
		err := a.Play("Spy and family", data)
		if err != nil {
			t.Fatal(err)
		}
	}
	x := a.List()
	fmt.Println(x)
}

/*
*** dashboard ***
a.List()
a.MakeAllRelease()
a.PauseAll()
a.Pause(s)
a.Play(s, data) // like Youtube AutoPlay, Play continually
// or 指定播放某系列的某集 a.Play("Spy and family/1", data)
a.PlayOnce(s, data) traceId
a.Remove(s) // Can't remove specified episode

*** series ***
s.Status() // show episodes status... and average executing time

*** episode ***
// episode Play process: click -> waiting -> playing -> Finished(leave without status)
//											   |_> pausing
//											   |_> Failed
e.ShowLog(int traceId)
e.GetFailing()
e.GetStatus()

// Target: I want to fix a series I can't stop the others, means it can fix on run-time.
*/
