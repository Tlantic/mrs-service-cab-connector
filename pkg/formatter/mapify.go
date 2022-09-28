package formatter

import (
	"encoding/json"
	"strconv"

	"github.com/shopspring/decimal"
)

func Mapify(content []byte) (map[string]string, error) {
	var clientRes map[string]interface{}

	// Deserialize response
	err := json.Unmarshal(content, &clientRes)
	if err != nil {
		return nil, err
	}

	//
	fields := make(map[string]string, 0)
	for key, value := range clientRes {

		// Check if list
		listValue, ok := value.([]interface{})
		if ok {
			listItemValue := listValue[0]

			// Serialize the item
			listItemValueJson, err := json.Marshal(listItemValue)
			if err != nil {
				return nil, err
			}

			// Add to fields
			listItemFields, err := Mapify(listItemValueJson)
			if err != nil {
				return nil, err
			}

			fields = mergeMaps(fields, addPrefixToMapKeys(key, listItemFields))
			fields = mergeMaps(fields, listItemFields)
			continue
		}

		// Check if object
		objectValue, ok := value.(map[string]interface{})
		if ok {
			for _key, _value := range objectValue {

				currentFields, err := resolveObject(_key, _value)
				if err != nil {
					return nil, err
				}

				fields = mergeMaps(fields, addPrefixToMapKeys(key, currentFields))
				fields = mergeMaps(fields, currentFields)
			}
			continue
		}

		// Add "normal" field
		valueString, err := resolveField(value)
		if err != nil {
			return nil, err
		}
		fields[key] = valueString
	}

	return fields, nil
}

// There aren't any lists inside objects in the examples, therefore, this case is being ignored
func resolveObject(key string, value interface{}) (map[string]string, error) {
	fields := make(map[string]string, 0)

	// Check if it's an object
	objectValue, ok := value.(map[string]interface{})
	if ok {
		for key, value := range objectValue {

			currentFields, err := resolveObject(key, value)
			if err != nil {
				return nil, err
			}

			fields = mergeMaps(fields, addPrefixToMapKeys(key, currentFields))
			fields = mergeMaps(fields, currentFields)
		}
		return fields, nil
	}

	// Treat it as a field
	valueString, err := resolveField(value)
	if err != nil {
		return nil, err
	}

	// If not, simply add the field
	fields[key] = valueString

	//fmt.Println("Value " + string(value))
	return fields, nil
}

func resolveField(value interface{}) (string, error) {
	var valueString string

	// Check if string
	valueString, ok := value.(string)
	if !ok {
		// Check if int
		valueInt, ok := value.(int)
		if ok {
			valueString = strconv.Itoa(valueInt)
		}

		// Check if int
		valueFloat, ok := value.(float64)
		if ok {
			valueString = decimal.NewFromFloat(valueFloat).String()
		}
	}

	return valueString, nil
}
