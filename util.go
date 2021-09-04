package async

type err string

func (e err) Error() string { return string(e) }

func empty[E any]() (_ E) { return }
