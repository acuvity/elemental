// Author: Antoine Mercadal
// See LICENSE file for full LICENSE
// Copyright 2016 Aporeto.

package elemental

import "reflect"

// extractFieldNames returns all the field Name of the given
// object using reflection.
func extractFieldNames(obj interface{}) []string {

	val := reflect.ValueOf(obj).Elem()
	c := val.NumField()
	fields := make([]string, c)

	for i := 0; i < c; i++ {
		fields[i] = val.Type().Field(i).Name
	}

	return fields
}

// areFieldValuesEqual checks if the value of the given field name are
// equal in both given objects using reflection.
func areFieldValuesEqual(field string, o1, o2 interface{}) bool {

	return reflect.ValueOf(o1).Elem().FieldByName(field).Interface() == reflect.ValueOf(o2).Elem().FieldByName(field).Interface()
}

// isFieldValueZero check if the value of the given field is set to its zero value.
//
// For a Slic
func isFieldValueZero(field string, o interface{}) bool {

	v := reflect.ValueOf(o).Elem().FieldByName(field)

	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.IsNil() || v.Len() == 0
	default:
		return v.Interface() == reflect.Zero(reflect.TypeOf(v.Interface())).Interface()
	}
}
