package async

func empty[E any]() E {
	// we can write something like
	//  func[E any]() (_ E) { return }
	// but we must not
	var v E
	return v
}

type err string

func (e err) Error() string { return string(e) }
