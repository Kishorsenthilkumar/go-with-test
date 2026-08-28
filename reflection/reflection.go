package reflection

import "reflect"

func Walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	switch val.Kind() {

	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)

			switch field.Kind() {
			case reflect.String:
				fn(field.String())

			case reflect.Struct:
				Walk(field.Interface(), fn)

			case reflect.Slice, reflect.Array:
				Walk(field.Interface(), fn)

			case reflect.Map:
				for _, key := range val.MapKeys() {
					Walk(val.MapIndex(key).Interface(), fn)
				}

			}

		}

	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			Walk(val.Index(i).Interface(), fn)
		}
	}

}
