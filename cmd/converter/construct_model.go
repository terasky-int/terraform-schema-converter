/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package converter

import (
	"reflect"
	"strings"

	"github.com/GilTeraSky/tfschemaconverter/internal/arrayindexer"
	"github.com/GilTeraSky/tfschemaconverter/pkg/maptypes"
	"github.com/GilTeraSky/tfschemaconverter/pkg/utils"
)

func (converter *TFSchemaModelConverter[T]) buildModelField(modelJSON *maptypes.BlockToStruct, schemaData interface{}, mapValue interface{}, arrIndexer *arrayindexer.ArrIndexer) {
	if schemaData == nil || mapValue == nil {
		return
	}

	switch mapValue := mapValue.(type) {
	case *maptypes.BlockToStruct:
		converter.modelHandleBlockStruct(modelJSON, schemaData, mapValue, arrIndexer)
	case *maptypes.Map:
		converter.modelHandleBlockMap(modelJSON, schemaData, mapValue, arrIndexer)
	case *maptypes.BlockToStructSlice:
		converter.modelHandleBlockStructSlice(modelJSON, schemaData, mapValue, arrIndexer)
	case *maptypes.BlockSliceToStructSlice:
		converter.modelHandleBlockSliceStructSlice(modelJSON, schemaData, mapValue, arrIndexer)
	case *maptypes.ListToStruct:
		converter.modelHandleListStruct(modelJSON, schemaData, mapValue, arrIndexer)
	case *maptypes.EvaluatedField:
		modelField := mapValue.Field
		modelValue := mapValue.EvalFunc(maptypes.ConstructModel, schemaData)
		converter.setModelValue(modelJSON, modelField, modelValue, arrIndexer)
	case string:
		modelField := mapValue
		modelValue := schemaData
		converter.setModelValue(modelJSON, modelField, modelValue, arrIndexer)
	}
}

func (converter *TFSchemaModelConverter[T]) modelHandleBlockStruct(modelJSON *maptypes.BlockToStruct, schemaData interface{}, mapValue *maptypes.BlockToStruct, arrIndexer *arrayindexer.ArrIndexer) {
	if schemaDataSlice, ok := schemaData.([]interface{}); ok && len(schemaDataSlice) > 0 {
		rootSchemaDict, _ := schemaDataSlice[0].(map[string]interface{})

		for key, value := range *mapValue {
			converter.buildModelField(modelJSON, rootSchemaDict[key], value, arrIndexer)
		}
	}
}

func (converter *TFSchemaModelConverter[T]) modelHandleBlockMap(modelJSON *maptypes.BlockToStruct, schemaData interface{}, mapValue *maptypes.Map, arrIndexer *arrayindexer.ArrIndexer) {
	if rootSchemaDict, ok := schemaData.(map[string]interface{}); ok {
		definedKeysMapValue := mapValue.Copy([]string{utils.AllMapKeysFieldMarker})

		if allKeysFlagMapValue, exists := (*mapValue)[utils.AllMapKeysFieldMarker]; exists {
			for key, value := range rootSchemaDict {
				var dynamicMapValue interface{}

				if allKeysFlagMapStr, ok := allKeysFlagMapValue.(string); ok {
					dynamicMapValue = strings.ReplaceAll(allKeysFlagMapStr, utils.AllMapKeysFieldMarker, key)
				} else {
					dynamicMapValue = allKeysFlagMapValue
				}

				converter.buildModelField(modelJSON, value, dynamicMapValue, arrIndexer)
			}
		}

		for key, value := range definedKeysMapValue {
			converter.buildModelField(modelJSON, rootSchemaDict[key], value, arrIndexer)
		}
	}
}

func (converter *TFSchemaModelConverter[T]) modelHandleBlockStructSlice(modelJSON *maptypes.BlockToStruct, schemaData interface{}, mapValue *maptypes.BlockToStructSlice, arrIndexer *arrayindexer.ArrIndexer) {
	if len(schemaData.([]interface{})) > 0 {
		arrIndexer.New()

		for _, elemTypeMap := range *mapValue {
			for elemMapKey, elemMapValue := range *elemTypeMap {
				var schemaValue, _ = (schemaData.([]interface{}))[0].(map[string]interface{})[elemMapKey]

				if schemaValue != nil {
					if _, ok := elemMapValue.(*maptypes.ListToStruct); ok {
						converter.buildModelField(modelJSON, schemaValue, elemMapValue, arrIndexer)
					} else {
						for _, item := range schemaValue.([]interface{}) {
							converter.buildModelField(modelJSON, []interface{}{item}, elemMapValue, arrIndexer)
							arrIndexer.IncrementLastIndex()
						}
					}
				}
			}
		}

		arrIndexer.RemoveLastIndex()
	}
}

func (converter *TFSchemaModelConverter[T]) modelHandleBlockSliceStructSlice(modelJSON *maptypes.BlockToStruct, schemaData interface{}, mapValue *maptypes.BlockSliceToStructSlice, arrIndexer *arrayindexer.ArrIndexer) {
	if len(schemaData.([]interface{})) > 0 {
		for _, elemTypeMap := range *mapValue {
			arrIndexer.New()

			for _, item := range schemaData.([]interface{}) {
				var _, ok = item.(map[string]interface{})

				if ok {
					converter.buildModelField(modelJSON, []interface{}{item}, elemTypeMap, arrIndexer)
					arrIndexer.IncrementLastIndex()
				}
			}

			arrIndexer.RemoveLastIndex()
		}
	}
}

func (converter *TFSchemaModelConverter[T]) modelHandleListStruct(modelJSON *maptypes.BlockToStruct, schemaData interface{}, mapValue *maptypes.ListToStruct, arrIndexer *arrayindexer.ArrIndexer) {
	if reflect.TypeOf(schemaData).Kind() == reflect.Slice {
		sliceValue := reflect.ValueOf(schemaData)

		for i := 0; i < sliceValue.Len(); i++ {
			val := sliceValue.Index(i).Interface()
			converter.setModelValue(modelJSON, (*mapValue)[0], val, arrIndexer)
			arrIndexer.IncrementLastIndex()
		}
	}
}

func (converter *TFSchemaModelConverter[T]) setModelValue(model *maptypes.BlockToStruct, field string, value interface{}, arrIndexer *arrayindexer.ArrIndexer) {
	modelPathSeparator := converter.getModelPathSeparator()

	if !strings.Contains(field, modelPathSeparator) {
		(*model)[field] = value
	} else {
		fieldPaths := strings.Split(field, modelPathSeparator)
		arrIndices := arrIndexer.GetAllIndexes()
		leafField := strings.ReplaceAll(fieldPaths[len(fieldPaths)-1], utils.ArrayFieldMarker, "")
		arrayFields := 0
		parentField := *model

		for i := 0; i < len(fieldPaths)-1; i++ {
			fieldName := fieldPaths[i]

			if !strings.Contains(fieldName, utils.ArrayFieldMarker) {
				if _, ok := parentField[fieldName]; !ok {
					parentField[fieldName] = make(map[string]interface{})
				}

				parentField = parentField[fieldName].(map[string]interface{})
			} else {
				var object map[string]interface{}

				fieldName = strings.ReplaceAll(fieldName, utils.ArrayFieldMarker, "")
				arrayIndex := arrIndices[arrayFields]

				arrayFields++

				if _, ok := parentField[fieldName]; !ok {
					parentField[fieldName] = make([]map[string]interface{}, 0)
				}

				parentFieldArray := parentField[fieldName].([]map[string]interface{})

				if arrayIndex < len(parentFieldArray) {
					object = parentFieldArray[arrayIndex]
				} else {
					object = make(map[string]interface{})
					parentField[fieldName] = append(parentFieldArray, object)
				}

				parentField = object
			}
		}

		parentField[leafField] = value
	}
}
