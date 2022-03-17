package async_test

import (
	"testing"
	"time"

	"github.com/ninedraft/async"
	"github.com/ninedraft/async/internal/testutils"
)

func TestAwait(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch := testutils.ChanWithValue("doot")

	v, ok := async.Await(done, ch)
	testutils.RequireEq(test, true, ok, "await.ok")
	testutils.RequireEq(test, "doot", v, "await result")

	cancel()
	v, ok = async.Await(done, ch)

	testutils.RequireEq(test, false, ok, "await.ok")
	testutils.RequireEq(test, "", v, "await result")
}

func TestAwait2_CancelCh1(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch1 := testutils.ClosedChan[int]()
	ch2 := testutils.ChanWithValue(2)

	v, ok := async.Await2(done, ch1, ch2)

	testutils.RequireEq(test, true, ok, "await.ok")
	testutils.RequireEq(test, 2, v, "await result")
}

func TestAwait2_CancelCh2(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch1 := testutils.ChanSendAfter(1, 100*time.Millisecond)
	ch2 := testutils.ClosedChan[int]()

	v, ok := async.Await2(done, ch1, ch2)

	testutils.RequireEq(test, true, ok, "await.ok")
	testutils.RequireEq(test, 1, v, "await result")
}

func TestAwait2_CancelAll(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)

	v, ok := async.Await2(done, ch1, ch2)

	testutils.RequireEq(test, false, ok, "await.ok")
	testutils.RequireEq(test, 0, v, "await result")
}

func TestAwait2_CloseAll(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch1 := testutils.ClosedChan[int]()
	ch2 := testutils.ClosedChan[int]()

	v, ok := async.Await2(done, ch1, ch2)

	testutils.RequireEq(test, false, ok, "await.ok")
	testutils.RequireEq(test, 0, v, "await result")
}

func TestAwait3_CloseCh1(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch1 := testutils.ClosedChan[int]()
	ch2 := testutils.ChanSendAfter(2, 100*time.Millisecond)
	ch3 := testutils.ChanSendAfter(3, 100*time.Millisecond)

	v, ok := async.Await3(done, ch1, ch2, ch3)

	testutils.RequireEq(test, true, ok, "await.ok")
	testutils.RequireOneOf(test, []int{2, 3}, v, "await result")
}

func TestAwait3_CloseCh2(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch1 := testutils.ChanSendAfter(1, 100*time.Millisecond)
	ch2 := testutils.ClosedChan[int]()
	ch3 := testutils.ChanSendAfter(3, 100*time.Millisecond)

	v, ok := async.Await3(done, ch1, ch2, ch3)

	testutils.RequireEq(test, true, ok, "await.ok")
	testutils.RequireOneOf(test, []int{1, 3}, v, "await result")
}

func TestAwait3_CloseCh3(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch1 := testutils.ChanSendAfter(1, 100*time.Millisecond)
	ch2 := testutils.ChanSendAfter(2, 100*time.Millisecond)
	ch3 := testutils.ClosedChan[int]()

	v, ok := async.Await3(done, ch1, ch2, ch3)

	testutils.RequireEq(test, true, ok, "await.ok")
	testutils.RequireOneOf(test, []int{1, 2}, v, "await result")
}

func TestAwait3_CancelAll(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	cancel()

	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	v, ok := async.Await3(done, ch1, ch2, ch3)

	testutils.RequireEq(test, false, ok, "await.ok")
	testutils.RequireEq(test, 0, v, "await result")
}

func TestAwait3_CloseAll(test *testing.T) {
	cancelTimeout := testutils.Timeout(test, time.Second)
	defer cancelTimeout()

	done, cancel := async.Cancel()
	defer cancel()

	ch1 := testutils.ClosedChan[int]()
	ch2 := testutils.ClosedChan[int]()
	ch3 := testutils.ClosedChan[int]()

	v, ok := async.Await3(done, ch1, ch2, ch3)

	testutils.RequireEq(test, false, ok, "await.ok")
	testutils.RequireEq(test, 0, v, "await result")
}
