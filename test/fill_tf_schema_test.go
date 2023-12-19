/*
Copyright © 2023 TeraSky, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	testdata "github.com/GilTeraSky/tfschemaconverter/test/data"
)

type FillTFSchemaTestCase struct {
	CaseName           string
	ModelData          *testdata.Zoo
	ExpectedSchemaData *schema.ResourceData
}

func TestFillTFSchema(t *testing.T) {
	cases := []FillTFSchemaTestCase{
		{
			CaseName:           "Fill TF Schema - Full Resource",
			ModelData:          testdata.FullZooModel,
			ExpectedSchemaData: testdata.BuildResourceData(t, testdata.FullZooTFData),
		},
		{
			CaseName:           "Fill TF Schema - Slim Resource",
			ModelData:          testdata.SlimZooModel,
			ExpectedSchemaData: testdata.BuildResourceData(t, testdata.SlimZooTFData),
		},
		{
			CaseName:           "Fill TF Schema - Minimal Resource",
			ModelData:          testdata.MinimalZooModel,
			ExpectedSchemaData: testdata.BuildResourceData(t, testdata.MinimalZooTFData),
		},
	}

	for _, c := range cases {
		resourceData := testdata.BuildResourceData(t, map[string]interface{}{})
		err := testdata.Converter.FillTFSchema(c.ModelData, resourceData)

		if err != nil {
			t.Fatalf("# Case '%s' -> An error has occured: %s", c.CaseName, err.Error())
		}

		if !testdata.CompareResourceData(resourceData, c.ExpectedSchemaData) {
			t.Fatalf("# Case '%s' -> Not Equal.", c.CaseName)
		}
	}
}
