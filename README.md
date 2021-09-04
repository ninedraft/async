# async

This experimental package provides some utilities for async tasks.

```
package async // import "github.com/ninedraft/async"


CONSTANTS

const ErrPromiseClosed err = "promise is closed"
    ErrPromiseClosed means that the promise was rejected without error. Usually
    it means that that source goroutine paniced.


FUNCTIONS

func Push[E any](ctx context.Context, ch chan<- E, v E) bool
func Run[E any](ctx context.Context, fn func(ctx context.Context) (E, error)) Promise[E]
    Run starts a new goroutine and returns a promised result.


TYPES

type Promise[E any] <-chan result[E]
    Promise is a deferred result of function execution.
```