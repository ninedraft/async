package async

import (
	"context"
)

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

type Promise[E any] <-chan result[E]

const ErrPromiseClosed err = "promise is closed"

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
