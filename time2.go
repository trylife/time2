package time2

import (
	"time"
)

const TimeYmdLayout = "2006-01-02"

type t2 struct {
	origin time.Time
	loc    *time.Location
}

func Now() *t2 {
	return &t2{
		origin: time.Now(),
	}
}

func Load(t time.Time) *t2 {
	return &t2{
		origin: t,
	}
}

func (t *t2) Loc(name string) *t2 {
	t.loc, _ = time.LoadLocation(name)
	return t
}

func (t *t2) CurrentDayStart() *time.Time {
	date := string(t.origin.Format(TimeYmdLayout))
	tm, _ := time.ParseInLocation(TimeYmdLayout, date, t.loc)
	return &tm
}
