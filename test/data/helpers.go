/*
Copyright © 2023 TeraSky, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testdata

func copyMap(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})

	for k, v := range m {
		vm, ok := v.(map[string]interface{})

		if ok {
			cp[k] = copyMap(vm)
		} else {
			cp[k] = v
		}
	}

	return cp
}

func compareJSON(firstValue, secondValue interface{}) bool {
	isFirstValueEmpty := isEmptyInterface(firstValue)
	isSecondValueEmpty := isEmptyInterface(secondValue)

	if (isFirstValueEmpty && !isSecondValueEmpty) || (!isFirstValueEmpty && isSecondValueEmpty) {
		return false
	} else if !isFirstValueEmpty && !isSecondValueEmpty {
		switch firstValue.(type) {
		case map[string]interface{}:
			firstMap := firstValue.(map[string]interface{})
			secondMap, isSecondValueMap := secondValue.(map[string]interface{})

			if !isSecondValueMap {
				return false
			}

			for k, v := range firstMap {
				secondValue, keyExist := secondMap[k]

				if !keyExist {
					return false
				}

				isEqual := compareJSON(v, secondValue)

				if !isEqual {
					return false
				}
			}
		case []interface{}:
			firstArray := firstValue.([]interface{})
			secondArray, isSecondValueArray := secondValue.([]interface{})

			if isSecondValueArray {
				if len(firstArray) != len(secondArray) {
					return false
				}

				for _, firstSubItem := range firstArray {
					itemFound := false

					for _, secondSubItem := range secondArray {
						if compareJSON(firstSubItem, secondSubItem) {
							itemFound = true
							break
						}
					}

					if !itemFound {
						return false
					}
				}
			}
		default:
			return firstValue == secondValue
		}
	}

	return true
}

func isEmptyInterface(value interface{}) bool {
	if value == nil {
		return true
	}

	switch value := value.(type) {
	case map[string]interface{}:
		return len(value) == 0
	case []interface{}:
		return len(value) == 0
	case string:
		return value == ""
	}

	return false
}
