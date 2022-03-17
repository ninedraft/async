package async

import (
	"context"
	"sync"
)

type Done = <-chan struct{}

func Cancel() (Done, context.CancelFunc) {
	var done = make(chan struct{})
	var once sync.Once
	return done, func() {
		once.Do(func() {
			close(done)
		})
	}
}
