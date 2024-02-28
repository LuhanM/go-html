package schemas

import (
	"sync"
	"time"
)

type Viewers struct {
	counter int
	List    map[string]time.Time
	mu      sync.Mutex
}

func (v *Viewers) Increase(viewerId string) {
	v.mu.Lock()
	defer v.mu.Unlock()

	if time.Now().After(v.List[viewerId].Add(time.Hour * 24)) {
		v.counter++
		v.List[viewerId] = time.Now()
	}
}

func (v *Viewers) GetValue() int {
	return v.counter
}
