package collection

import (
	"bytes"
	"encoding/gob"
)

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
	} else {
		switch values[0].(string) {
		case "=":
			for _, value := range c {
				if value[key] == values[1] {
					d = append(d, value)
				}
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

func copyMap(m map[string]interface{}) map[string]interface{} {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	err := enc.Encode(m)
	if err != nil {
		panic(err)
	}
	var cm map[string]interface{}
	err = dec.Decode(&cm)
	if err != nil {
		panic(err)
	}
	return cm
}

func isTrue(a interface{}) bool {
	switch a.(type) {
	case uint:
		return a.(uint) != uint(0)
	case uint8:
		return a.(uint8) != uint8(0)
	case uint16:
		return a.(uint16) != uint16(0)
	case uint32:
		return a.(uint32) != uint32(0)
	case uint64:
		return a.(uint64) != uint64(0)
	case int:
		return a.(int) != int(0)
	case int8:
		return a.(int8) != int8(0)
	case int16:
		return a.(int16) != int16(0)
	case int32:
		return a.(int32) != int32(0)
	case int64:
		return a.(int64) != int64(0)
	case float32:
		return a.(float32) != float32(0)
	case float64:
		return a.(float64) != float64(0)
	case string:
		return a.(string) != ""
	case bool:
		return a.(bool)
	default:
		return false
	}
}
