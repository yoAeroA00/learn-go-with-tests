package walk

import "reflect"

func walk(x any, fn func(string)) {
	visited := make(map[uintptr]bool)
	walkValue(x, fn, visited)
}

func walkValue(x any, fn func(string), visited map[uintptr]bool) {
	val := getValue(x, visited)

	numOfVal := 0
	var getField func(int) reflect.Value
	walkFunc := func(value reflect.Value) {
		walkValue(value.Interface(), fn, visited)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numOfVal = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numOfVal = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkFunc(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkFunc(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, value := range valFnResult {
			walkFunc(value)
		}
	}

	for i := 0; i < numOfVal; i++ {
		walkFunc(getField(i))
	}
}

func getValue(x any, visited map[uintptr]bool) (val reflect.Value) {
	val = reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		if val.IsNil() {
			return
		}

		ptr := val.Pointer()

		if visited[ptr] {
			return
		}

		visited[ptr] = true
		val = val.Elem()
	}

	return
}
