// Package types provee tipos de utilidad para el resto de los paquetes.
package types

type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

type Iterator[T any] interface {
	HasNext() bool
	Next() (T, error)
}
