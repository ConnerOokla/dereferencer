package dereference

import (
	"reflect"
)

// Dereference function to dereference pointers in a nested struct and skip unsupported types
func Dereference(v interface{}) interface{} {
	val := reflect.ValueOf(v)

	// If the value is a pointer, dereference it
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	derefStruc := reflect.New(val.Type()).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		newField := derefStruc.Field(i)

		switch field.Kind() {
		case reflect.Func:
			continue

		case reflect.Ptr:
			if !field.IsNil() {
				derefValue := field.Elem()

				switch newField.Kind() {
				case reflect.Ptr:
					newField.Set(reflect.New(derefValue.Type()))
					newField.Elem().Set(derefValue)
				default:
					newField.Set(derefValue)
				}
			}

		case reflect.Struct:
			newField.Set(reflect.ValueOf(Dereference(field.Interface())))

		default:
			if newField.CanSet() {
				newField.Set(field)
			}
		}
	}

	return derefStruc.Interface()
}
