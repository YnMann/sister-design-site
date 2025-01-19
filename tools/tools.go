package tools

import "time"

// ConvertTimeToJSTime converts time to JS time format
func ConvertTimeToJSTime(t time.Time) string {
	return t.Format("2006-01-02T15:04:05.000Z")
}

// IsLastInMap checks if the key is the last in the map
func IsLastInMap(key string, params map[string]string) bool {
	i := 0
	for k := range params {
		if k == key {
			break
		}
		i++
	}
	return i == len(params)-1
}

// AppendIfNotNil adds an element to a slice if it is not nil.
func AppendIfNotNil[T any](field *T, slice *[]T) {
	if field != nil {
		*slice = append(*slice, *field)
	}
}
