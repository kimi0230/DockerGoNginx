package validatefunction

import (
	"reflect"

	"gopkg.in/go-playground/validator.v8"
)

func AwsomeValidate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if data, ok := field.Interface().(string); ok {
		if data != "awsome" {
			return false
		}
	}
	return true
}
