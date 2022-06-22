package spy_and_family

import "github.com/FPNL/anime/package/anime"

const Name = "Spy and family"

func SetupSpyAndFamily(a *anime.Dashboard) {
	s := a.NewSeries(Name)
	s.AddEpisode([]*anime.Episode{
		anime.NewEpisode(ep1),
		anime.NewEpisode(ep2),
		anime.NewEpisode(ep3),
	})
}
