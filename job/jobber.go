package job

import "time"

type Jobber interface {
	Do() error
	DoAt(t time.Time) error
	ReDoAt(t time.Time,dt time.Duration) error
	StartDo() error
	Compare(j Jobber) bool
}
