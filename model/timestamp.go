package model

import (
	"normal_pkg/tools"
	"strconv"
	"time"
)

type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(data []byte) (err error) {
	tm, err := tools.BytesToInt(data)
	*t = Timestamp(time.UnixMilli(tm))
	return err
}

func (t Timestamp) MarshalJSON() ([]byte, error) {
	tm := time.Time(t)
	r := tm.UnixMilli()
	bytes, err := tools.IntToBytes(int(r))
	return bytes, err
}

func (t Timestamp) String() string {
	return strconv.FormatInt(time.Time(t).UnixMilli(), 10)
}
