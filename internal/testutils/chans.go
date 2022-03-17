package testutils

import "time"

func ClosedChan[E any]() chan E {
	ch := make(chan E)
	close(ch)
	return ch
}

func ChanWithValue[E any](v E) chan E {
	ch := make(chan E, 1)
	ch <- v
	return ch
}

func ChanSendAfter[E any](v E, dt time.Duration) chan E {
	ch := make(chan E, 1)
	time.AfterFunc(dt, func() {
		ch <- v
	})
	return ch
}
