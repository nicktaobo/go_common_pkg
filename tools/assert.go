package tools

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

func Choose[T any](b bool, src T, dst T) T {
	if b {
		return src
	}
	return dst
}

func Choose1(b bool, src any, dst any) any {
	if b {
		return src
	}
	return dst
}
