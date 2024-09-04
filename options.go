package options

func CreateWithDefault[T any](v T, options ...Applier[T]) *T {
	for _, apply := range options {
		apply(&v)
	}
	return &v
}

func Create[T any](options ...Applier[T]) *T {
	option := new(T)
	for _, apply := range options {
		apply(option)
	}
	return option
}

type Applier[T any] func(*T)
