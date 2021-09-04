package async

import (
	"context"
)

// Run starts a new goroutine and returns a promised result.
func Run[E any](ctx context.Context, fn func(ctx context.Context) (E, error)) Promise[E] {
	var future = make(chan result[E], 1)
	go func() {
		defer close(future)
		var v, e = fn(ctx)
		future <- result[E]{
			value: v,
			err:   e,
		}
	}()
	return Promise[E](future)
}

// Promise is a deferred result of function execution.
type Promise[E any] <-chan result[E]

// ErrPromiseClosed means that the promise was rejected without error.
// Usually it means that that source goroutine paniced.
const ErrPromiseClosed err = "promise is closed"

// Await polls the promise until context is canceled or result is returned.
// Returns context errors, if context is canceled.
// Returns ErrPromiseClosed is promise was closed without result.
func (promise Promise[E]) Await(ctx context.Context) (E, error) {
	select {
	case <-ctx.Done():
		return empty[E](), ctx.Err()
	case res, ok := <-promise:
		if !ok {
			return empty[E](), ErrPromiseClosed
		}
		return res.value, res.err
	}
}

type result[E any] struct {
	value E
	err   error
}
