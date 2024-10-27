package pogo

import (
	"fmt"
	"reflect"
	"strings"
)

// Function to map query results to struct fields based on `pogo` tag.
func SuperQuery[T any](db *Database, query string, recipient *[]T, args ...any) error {
	// Perform the query
	if strings.Contains(query, ":fields") {
		fields := get_tags(recipient)
		query = strings.ReplaceAll(query, ":fields", strings.Join(fields, ","))
	}
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	// Use reflection to ensure recipient is a pointer to a slice
	val := reflect.ValueOf(recipient)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Slice {
		return fmt.Errorf("recipient must be a pointer to a slice")
	}

	// Get the slice element type (e.g., the User struct type)
	elemType := val.Elem().Type().Elem()

	// Prepare a slice for scanning field values
	columns, err := rows.Columns()
	if err != nil {
		return err
	}

	for rows.Next() {
		// Create a new instance of the struct (e.g., User)
		newElem := reflect.New(elemType).Elem()

		// Create a map for storing pointers to the fields based on `pogo` tag
		fieldMap := make(map[string]interface{})

		// Scan the struct fields and prepare for scanning the values from the row
		fieldPointers := make([]interface{}, len(columns))
		for i, column := range columns { // Iterate through the struct fields to match the `pogo` tags
			for j := 0; j < newElem.NumField(); j++ {
				field := newElem.Type().Field(j)
				tag := field.Tag.Get("pogo")

				// If the column matches the `pogo` tag, store a pointer to the field
				if tag == column {
					fieldMap[column] = newElem.Field(j).Addr().Interface()
					break
				}
			}

			// Assign the pointer for the current column (if found in the map)
			if ptr, ok := fieldMap[column]; ok {
				fieldPointers[i] = ptr
			} else {
				var dummy interface{}
				fieldPointers[i] = &dummy // Placeholder for columns not mapped to struct fields
			}
		}

		// Scan the row into the struct fields
		if err := rows.Scan(fieldPointers...); err != nil {
			return err
		}

		// Append the populated struct to the recipient slice
		val.Elem().Set(reflect.Append(val.Elem(), newElem))
	}

	return rows.Err()
}

func get_tags[T any](data *T) []string {
	st := reflect.ValueOf(data)
	elemType := st.Elem().Type().Elem()
	newElem := reflect.New(elemType).Elem()
	vals := make([]string, 0)

	for i := 0; i < newElem.NumField(); i++ {
		field := newElem.Type().Field(i)
		val := field.Tag.Get("pogo")
		vals = append(vals, val)
	}

	return vals
}
