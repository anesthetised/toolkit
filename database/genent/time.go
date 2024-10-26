package genent

import (
	"sync"
	"time"
)

// TimeDefaultNow returns a function
// which executes time.Now exactly once.
func TimeDefaultNow() func() time.Time {
	var (
		once sync.Once
		t    time.Time
	)

	return func() time.Time {
		once.Do(func() {
			t = time.Now()
		})

		return t
	}
}
