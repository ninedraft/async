package async

import (
	"time"
	"context"
	"errors"
	"fmt"
	"testing"
)

func ExampleRun() {
	var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var doSomething = func(ctx context.Context, request string) (int, error) {return 0, nil}

	// two calls of doSomething will be executed simultaneously
	var result01 = Run(ctx, func(ctx context.Context) (int, error) {
		return doSomething(ctx, "first")
	})
	var result02 = Run(ctx, func(ctx context.Context) (int, error) {
		return doSomething(ctx, "second")
	})
	
	fmt.Println(result01.Await(ctx))
	fmt.Println(result02.Await(ctx))
}

func TestRun(test *testing.T) {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	var p = Run(ctx, func(ctx context.Context) (string, error) {
		return "test value", nil
	})
	var x, errAwait = p.Await(ctx)
	if errAwait != nil {
		test.Fatal(errAwait)
	}
	test.Log("x", x)
}

func TestRunError(test *testing.T) {
	var ctx, cancel = context.WithCancel(context.Background())
	defer cancel()
	var p = Run(ctx, func(ctx context.Context) (string, error) {
		return "", err("test error")
	})
	var _, errAwait = p.Await(ctx)
	if errAwait == nil {
		test.Fatalf("expected error, got nil")
	}
}

func TestRunCanceled(test *testing.T) {
	var ctx, cancel = context.WithCancel(context.Background())
	cancel()
	var p = Run(ctx, func(ctx context.Context) (string, error) {
		return "test value", nil
	})
	var _, errAwait = p.Await(ctx)
	var expected = context.Canceled
	if !errors.Is(errAwait, expected) {
		test.Fatalf("expected %q, got %v", expected, errAwait)
	}
}
