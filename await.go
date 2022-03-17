package async

import "sync"

func Await[E any](done Done, ch <-chan E) (E, bool) {
	select {
	case <-done:
		return empty[E](), false
	case v, ok := <-ch:
		return v, ok
	}
}

func Await2[E any](done Done, ch1, ch2 <-chan E) (E, bool) {
	for i := 0; i < 2; i++ {
		select {
		case <-done:
			return empty[E](), false
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			return v, true
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			return v, true
		}
	}
	return empty[E](), false
}

func Await3[E any](done Done, ch1, ch2, ch3 <-chan E) (E, bool) {
	for i := 0; i < 3; i++ {
		select {
		case <-done:
			return empty[E](), false
		case v, ok := <-ch1:
			if !ok {
				ch1 = nil
				continue
			}
			return v, true
		case v, ok := <-ch2:
			if !ok {
				ch2 = nil
				continue
			}
			return v, true
		case v, ok := <-ch3:
			if !ok {
				ch3 = nil
				continue
			}
			return v, true
		}
	}
	return empty[E](), false
}

func AwaitN[E any](done Done, chch ...<-chan E) (E, bool) {
	wg := &sync.WaitGroup{}
	wg.Add(len(chch))
	result := make(chan E, 1)

	groupDone, cancel := Cancel()
	wg.Add(1)
	go func() {
		defer wg.Done()
		Await(done, groupDone)
		cancel()
	}()
	for _, ch := range chch {
		go func(ch <-chan E) {
			defer wg.Done()
			v, ok := Await(groupDone, ch)
			if !ok {
				return
			}
			select {
			case result <- v:
				cancel()
			case <-groupDone:
			case <-done:
			}
		}(ch)
	}
	go func() {
		wg.Wait()
		cancel()
		close(result)
	}()

	return Await(done, result)
}
