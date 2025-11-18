package typeutil

import "errors"

func GetByTypeFromArray[I any, T any](arr []I) (T, error) {
	var zero T
	for _, v := range arr {
		if val, ok := any(v).(T); ok {
			return val, nil
		}
	}
	return zero, errors.New("type not found")
}

func MustGetByTypeFromArray[I any, T any](arr []I) T {
	val, err := GetByTypeFromArray[I, T](arr)
	if err != nil {
		panic(err)
	}
	return val
}
