package time2_test

import (
	"testing"
	"github.com/trylife/time2"
)

func TestTime2_CurrentDayStart(t *testing.T) {
	t2 := time2.Now()
	st := t2.CurrentDayStart()
	t.Log(st)
	t.Log(st, st.Hour(), st.Minute(), st.Second(), st.Nanosecond())

	if st.Hour() != 0 {
		t.Error("Not 0 Hour")
	}
	if st.Minute() != 0 {
		t.Error("Not 0 Minute")
	}
	if st.Second() != 0 {
		t.Error("Not 0 Second")
	}
}
