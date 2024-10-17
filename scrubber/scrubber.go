package scrubber

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

// Scrub removes specified fields from a struct.
// https://github.com/ssrathi/go-scrub was nice help.
// Example use put strings in a []string to pass into the Scrub Function.
func Scrub(data interface{}, fieldsToRemove []string) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling struct: %v", err)
	}

	// Iterate over each field and remove it
	for _, field := range fieldsToRemove {
		if gjson.GetBytes(jsonData, field).Exists() {

			jsonData, err = sjson.DeleteBytes(jsonData, field)
			if err != nil {
				return fmt.Errorf("error removing field '%s': %v", field, err)
			}
			fmt.Printf("Removed field: %s\n", field)
		} else {
			fmt.Printf("Field not found, skipping: %s\n", field)
		}
	}

	// Convert the JSON back to the original struct
	if err := json.Unmarshal(jsonData, &data); err != nil {
		return fmt.Errorf("error unmarshalling JSON back to struct: %v", err)
	}

	return nil
}
