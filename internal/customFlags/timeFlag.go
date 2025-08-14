package customflags

import "time"

type TimeValue struct {
	Time       *time.Time
	TimeLayout string
}

func (T TimeValue) String() string {
	if T.Time != nil {
		return T.Time.Format(T.TimeLayout)
	} else {
		return ""
	}
}

func (T TimeValue) Set(s string) error {
	if t, err := time.Parse(T.TimeLayout, s); err != nil {
		return err
	} else {
		(*T.Time) = t
	}
	return nil
}
