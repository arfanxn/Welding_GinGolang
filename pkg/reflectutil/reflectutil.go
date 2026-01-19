package reflectutil

import (
	"encoding"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

/* ============================================
   Struct Helpers
=============================================== */

// DerefAndInit dereferences pointers and allocates nil ones so
// the result is always a concrete, settable underlying value.
func DerefAndInit(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		v = v.Elem()
	}
	return v
}

// Deref dereferences pointers until it reaches a non-pointer.
func Deref(v reflect.Value) reflect.Value {
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return v // cannot dereference further
		}
		v = v.Elem()
	}
	return v
}

func DerefType(t reflect.Type) reflect.Type {
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	return t
}

// GetStructName deeply dereferences pointers and returns the struct name.
func GetStructName(v any) string {
	t := DerefType(reflect.TypeOf(v))
	return t.Name()
}

// GetStructField deeply dereferences and returns a struct field by name.
func GetStructField(v any, name string) (any, bool) {
	val := Deref(reflect.ValueOf(v))

	if val.Kind() != reflect.Struct {
		return nil, false
	}

	field := val.FieldByName(name)
	if !field.IsValid() || !field.CanInterface() {
		return nil, false
	}

	return field.Interface(), true
}

/* ============================================
   Helpers
=============================================== */

func isTimeDurationType(t reflect.Type) bool {
	return t.PkgPath() == "time" && t.Name() == "Duration"
}

/* ============================================
   SetValueFromString
=============================================== */

func SetValueFromString(fieldVal reflect.Value, str string) error {
	// Handle pointer types
	fieldVal = DerefAndInit(fieldVal)

	t := fieldVal.Type()

	// Special: time.Duration
	if isTimeDurationType(t) {
		d, err := time.ParseDuration(str)
		if err != nil {
			return err
		}
		fieldVal.SetInt(int64(d))
		return nil
	}

	switch fieldVal.Kind() {
	case reflect.String:
		fieldVal.SetString(str)
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return setInt(fieldVal, str)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return setUint(fieldVal, str)
	case reflect.Float32, reflect.Float64:
		return setFloat(fieldVal, str)
	case reflect.Bool:
		return setBool(fieldVal, str)
	default:
		// custom TextUnmarshaler
		if fieldVal.CanAddr() {
			if unmarshaler, ok := fieldVal.Addr().Interface().(encoding.TextUnmarshaler); ok {
				return unmarshaler.UnmarshalText([]byte(str))
			}
		}
		return fmt.Errorf("unsupported type: %s", t.String())
	}

	return nil
}

func setInt(v reflect.Value, s string) error {
	x, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	v.SetInt(x)
	return nil
}

func setUint(v reflect.Value, s string) error {
	x, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return err
	}
	v.SetUint(x)
	return nil
}

func setFloat(v reflect.Value, s string) error {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	v.SetFloat(x)
	return nil
}

func setBool(v reflect.Value, s string) error {
	x, err := strconv.ParseBool(s)
	if err != nil {
		return err
	}
	v.SetBool(x)
	return nil
}
