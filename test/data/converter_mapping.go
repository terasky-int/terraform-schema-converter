/*
Copyright © 2023 TeraSky, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testdata

import (
	"encoding/json"

	tfschemaconverter "github.com/GilTeraSky/tfschemaconverter/cmd/converter"
	"github.com/GilTeraSky/tfschemaconverter/pkg/maptypes"
	"github.com/GilTeraSky/tfschemaconverter/pkg/utils"
)

var (
	employeesArrayField     = utils.BuildArrayField("employees")
	cagesArrayField         = utils.BuildArrayField("cages")
	subscriptionsArrayField = utils.BuildArrayField("subscriptions")
)

var tfModelMap = &maptypes.BlockToStruct{
	NameKey: utils.BuildDefaultModelPath("name"),
	EmployeesKey: &maptypes.BlockSliceToStructSlice{
		{
			TitleKey:    utils.BuildDefaultModelPath(employeesArrayField, "title"),
			FullNameKey: utils.BuildDefaultModelPath(employeesArrayField, "fullName"),
			AgeKey:      utils.BuildDefaultModelPath(employeesArrayField, "age"),
		},
	},
	GateKey: &maptypes.BlockToStruct{
		EmployeesKey: &maptypes.BlockSliceToStructSlice{
			{
				TitleKey:    utils.BuildDefaultModelPath("gate", employeesArrayField, "title"),
				FullNameKey: utils.BuildDefaultModelPath("gate", employeesArrayField, "fullName"),
				AgeKey:      utils.BuildDefaultModelPath("gate", employeesArrayField, "age"),
			},
		},
		OpeningHourKey: utils.BuildDefaultModelPath("gate", "openingHour"),
		ClosingHourKey: utils.BuildDefaultModelPath("gate", "closingHour"),
	},
	CagesKey: &maptypes.BlockSliceToStructSlice{
		{
			EmployeesKey: &maptypes.BlockSliceToStructSlice{
				{
					TitleKey:    utils.BuildDefaultModelPath(cagesArrayField, employeesArrayField, "title"),
					FullNameKey: utils.BuildDefaultModelPath(cagesArrayField, employeesArrayField, "fullName"),
					AgeKey:      utils.BuildDefaultModelPath(cagesArrayField, employeesArrayField, "age"),
				},
			},
			AnimalNameKey: utils.BuildDefaultModelPath(cagesArrayField, "animalName"),
			LabelsKey: &maptypes.Map{
				utils.AllMapKeysFieldMarker: utils.BuildDefaultModelPath(cagesArrayField, "labels", utils.AllMapKeysFieldMarker),
			},
		},
	},
	LocationKey: &maptypes.BlockToStruct{
		LongitudeKey: utils.BuildDefaultModelPath("location", "longitude"),
		LatitudeKey:  utils.BuildDefaultModelPath("location", "latitude"),
	},
	WidthKey:  utils.BuildDefaultModelPath("area", "width"),
	LengthKey: utils.BuildDefaultModelPath("area", "length"),
	SubscriptionsKey: &maptypes.BlockToStructSlice{
		{
			IndividualClientKey: &maptypes.BlockToStruct{
				IDKey:        utils.BuildDefaultModelPath(subscriptionsArrayField, "person", "id"),
				FirstNameKey: utils.BuildDefaultModelPath(subscriptionsArrayField, "person", "firstName"),
				LastNameKey:  utils.BuildDefaultModelPath(subscriptionsArrayField, "person", "lastName"),
			},
		},
		{
			GroupClientKey: &maptypes.ListToStruct{utils.BuildDefaultModelPath(subscriptionsArrayField, "group", "name")},
		},
	},
	AddOnsKey: &maptypes.EvaluatedField{
		Field:    utils.BuildDefaultModelPath("addOns"),
		EvalFunc: evaluateAddOns,
	},
}

func evaluateAddOns(mode maptypes.EvaluationMode, value interface{}) (convertedResult interface{}) {
	if !isEmptyInterface(value) {
		if mode == maptypes.ConstructModel {
			convertedResult = make([]map[string]interface{}, 0)
			tfValueJSON := make(map[string]interface{})
			_ = json.Unmarshal([]byte(value.(string)), &tfValueJSON)

			for k, v := range tfValueJSON {
				modelAddOn := map[string]interface{}{
					"name":  k,
					"value": v,
				}

				convertedResult = append(convertedResult.([]map[string]interface{}), modelAddOn)
			}
		} else {
			addOnsJSON := make(map[string]interface{})

			for _, modelAddOn := range value.([]interface{}) {
				addOnKey := modelAddOn.(map[string]interface{})["name"].(string)
				addOnValue := modelAddOn.(map[string]interface{})["value"]
				addOnsJSON[addOnKey] = addOnValue
			}

			addOnsBytes, _ := json.Marshal(addOnsJSON)
			convertedResult = string(addOnsBytes)
		}
	}

	return convertedResult
}

var Converter = tfschemaconverter.TFSchemaModelConverter[*Zoo]{
	TFModelMap: tfModelMap,
}
