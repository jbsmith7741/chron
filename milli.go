package chron

import (
	"time"

	"github.com/dustinevan/chron/dura"
)

type Milli struct {
	time.Time
}

func NewMilli(year int, month time.Month, day, hour, min, sec, milli int) Milli {
	return Milli{time.Date(year, month, day, hour, min, sec, milli*1000000, time.UTC)}
}

func ThisMilli() Milli {
	return Now().AsMilli()
}

func MilliOf(t time.Time) Milli {
	return Milli{t.Truncate(time.Millisecond)}
}

func (m Milli) AsYear() Year           { return YearOf(m.Time) }
func (m Milli) AsMonth() Month         { return MonthOf(m.Time) }
func (m Milli) AsDay() Day             { return DayOf(m.Time) }
func (m Milli) AsHour() Hour           { return HourOf(m.Time) }
func (m Milli) AsMinute() Minute       { return MinuteOf(m.Time) }
func (m Milli) AsSecond() Second       { return SecondOf(m.Time) }
func (m Milli) AsMilli() Milli         { return MilliOf(m.Time) }
func (m Milli) AsMicro() Micro         { return MicroOf(m.Time) }
func (m Milli) AsTimeExact() TimeExact { return TimeOf(m.Time) }
func (m Milli) AsTime() time.Time      { return m.Time }

func (m Milli) Increment(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(l.Years(), l.Months(), l.Days()).Add(l.Duration())}
}

func (m Milli) Decrement(l dura.Time) TimeExact {
	return TimeExact{m.AddDate(-1*l.Years(), -1*l.Months(), -1*l.Days()).Add(-1 * l.Duration())}
}

func (m Milli) AddN(n int) Milli {
	return Milli{m.Add(time.Duration(int(time.Millisecond) * n))}
}

// span.Time implementation
func (m Milli) Start() Time {
	return m.AsTimeExact()
}

func (m Milli) End() Time {
	return m.AddN(1).Decrement(dura.Nano)
}

func (m Milli) Contains(t Span) bool {
	return !m.Before(t) && !m.After(t)
}

func (m Milli) Before(t Span) bool {
	return m.End().AsTime().Before(t.Start().AsTime())
}

func (m Milli) After(t Span) bool {
	return m.Start().AsTime().After(t.End().AsTime())
}

func (m Milli) Duration() dura.Time {
	return dura.Milli
}
