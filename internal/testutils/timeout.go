package testutils

import (
	"context"
	"sync"
	"testing"
	"time"
)

func Timeout(t testing.TB, timeout time.Duration) context.CancelFunc {
	t.Helper()
	timer := time.AfterFunc(timeout, func() {
		t.Fatalf("timeout (%s)", timeout)
	})
	stop := func() {
		t.Helper()
		timer.Stop()
	}
	once := sync.Once{}
	return func() {
		t.Helper()
		once.Do(stop)
	}
}
