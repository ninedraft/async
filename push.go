package async

func Push[E any](done Done, dst chan<- E, v E) bool {
	select {
	case <-done:
		return false
	case dst <- v:
		return true
	}
}
