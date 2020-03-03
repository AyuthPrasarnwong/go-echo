package helpers

import (
	"encoding/csv"
	"os"
	"reflect"
)

// FillStruct fill struct
func FillStruct(data map[string]interface{}, result interface{}) {
	t := reflect.ValueOf(result).Elem()
	for k, v := range data {
		val := t.FieldByName(k)
		val.Set(reflect.ValueOf(v))
	}
}

// ReadCsv fill struct
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

// InArray find value in array
func InArray(array interface{}, val interface{}) (index int, exists bool) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

// IsEmptyString check is empty string
func IsEmptyString(v string) bool {
	if v == "" {
		return false
	}
	return true
}

// WrapStringNil check empty sring then return nil
func WrapStringNil(v string) interface{} {
	if v == "" {
		return nil
	}
	return v
}

// // DateFormat parse and format time
// func DateFormat(layout string, v string) string {
// 	date, err := time.Parse(layout, v)
// 	if err != nil {

// 	}
// 	return date.Format(v)
// }
