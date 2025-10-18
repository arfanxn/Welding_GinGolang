package reflectutil

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// SetValueFromString sets a reflect.Value from a string representation.
// It supports basic types, durations, and anything implementing encoding.TextUnmarshaler.
func SetValueFromString(fieldVal reflect.Value, str string) error {
	// Handle pointer types
	if fieldVal.Kind() == reflect.Pointer {
		if fieldVal.IsNil() {
			fieldVal.Set(reflect.New(fieldVal.Type().Elem()))
		}
		return SetValueFromString(fieldVal.Elem(), str)
	}

	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(str)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Special case: time.Duration
		if fieldVal.Type().PkgPath() == "time" && fieldVal.Type().Name() == "Duration" {
			if d, err := time.ParseDuration(str); err == nil {
				fieldVal.SetInt(int64(d))
				return nil
			} else {
				return err
			}
		}
		if v, err := strconv.ParseInt(str, 10, 64); err == nil {
			fieldVal.SetInt(v)
		} else {
			return err
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v, err := strconv.ParseUint(str, 10, 64); err == nil {
			fieldVal.SetUint(v)
		} else {
			return err
		}
	case reflect.Float32, reflect.Float64:
		if v, err := strconv.ParseFloat(str, 64); err == nil {
			fieldVal.SetFloat(v)
		} else {
			return err
		}
	case reflect.Bool:
		if v, err := strconv.ParseBool(str); err == nil {
			fieldVal.SetBool(v)
		} else {
			return err
		}
	default:
		// Handle custom types that implement encoding.TextUnmarshaler
		if fieldVal.CanAddr() {
			if unmarshaler, ok := fieldVal.Addr().Interface().(encoding.TextUnmarshaler); ok {
				return unmarshaler.UnmarshalText([]byte(str))
			}
		}

		return fmt.Errorf("unsupported type: %s", fieldVal.Type().String())
	}
	return nil
}

// GetValueAsString returns a string representation of a reflect.Value.
func GetValueAsString(fieldVal reflect.Value) (string, error) {
	if fieldVal.Kind() == reflect.Pointer {
		if fieldVal.IsNil() {
			return "", nil
		}
		return GetValueAsString(fieldVal.Elem())
	}

	switch fieldVal.Kind() {
	case reflect.String:
		return fieldVal.String(), nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// Special case: time.Duration
		if fieldVal.Type().PkgPath() == "time" && fieldVal.Type().Name() == "Duration" {
			return time.Duration(fieldVal.Int()).String(), nil
		}
		return strconv.FormatInt(fieldVal.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return strconv.FormatUint(fieldVal.Uint(), 10), nil
	case reflect.Float32, reflect.Float64:
		return strconv.FormatFloat(fieldVal.Float(), 'f', -1, 64), nil
	case reflect.Bool:
		return strconv.FormatBool(fieldVal.Bool()), nil
	default:
		// Handle custom types (encoding.TextMarshaler)
		if marshaler, ok := fieldVal.Interface().(encoding.TextMarshaler); ok {
			b, err := marshaler.MarshalText()
			return string(b), err
		}
		return "", fmt.Errorf("unsupported type: %s", fieldVal.Type().String())
	}
}

// GetStructName returns the name of a struct.
func GetStructName(v any) string {
	t := reflect.TypeOf(v)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t.Name()
}

func IsSlice(v any) bool {
	return reflect.TypeOf(v).Kind() == reflect.Slice
}
