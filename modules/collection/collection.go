package collection

type Collection []map[string]interface{}

// Where filters the collection by a given key / value pair.
func (c Collection) Where(key string, values ...interface{}) Collection {
	var d = make([]map[string]interface{}, 0)
	if len(values) < 1 {
		for _, value := range c {
			if isTrue(value[key]) {
				d = append(d, value)
			}
		}
	} else if len(values) < 2 {
		for _, value := range c {
			if value[key] == values[0] {
				d = append(d, value)
			}
		}
	} else if values[0].(string) == "=" {
		for _, value := range c {
			if value[key] == values[1] {
				d = append(d, value)
			}
		}
	}
	return d
}

func (c Collection) Length() int {
	return len(c)
}

func (c Collection) FirstGet(key string) interface{} {
	return c[0][key]
}

func isTrue(a interface{}) bool {
	switch a := a.(type) {
	case uint:
		return a != uint(0)
	case uint8:
		return a != uint8(0)
	case uint16:
		return a != uint16(0)
	case uint32:
		return a != uint32(0)
	case uint64:
		return a != uint64(0)
	case int:
		return a != int(0)
	case int8:
		return a != int8(0)
	case int16:
		return a != int16(0)
	case int32:
		return a != int32(0)
	case int64:
		return a != int64(0)
	case float32:
		return a != float32(0)
	case float64:
		return a != float64(0)
	case string:
		return a != ""
	case bool:
		return a
	default:
		return false
	}
}
