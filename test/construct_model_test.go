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

type ConstructModelTestCase struct {
	CaseName      string
	SchemaData    *schema.ResourceData
	BuildOnlyKeys []string
	ExpectedModel *testdata.Zoo
}

func TestConstructModel(t *testing.T) {
	cases := []ConstructModelTestCase{
		{
			CaseName:      "Model Construct - Full Resource",
			SchemaData:    testdata.BuildResourceData(t, testdata.FullZooTFData),
			ExpectedModel: testdata.FullZooModel,
		},
		{
			CaseName:      "Model Construct - Slim Resource",
			SchemaData:    testdata.BuildResourceData(t, testdata.SlimZooTFData),
			ExpectedModel: testdata.SlimZooModel,
		},
		{
			CaseName:      "Model Construct - Minimal Resource",
			SchemaData:    testdata.BuildResourceData(t, testdata.MinimalZooTFData),
			ExpectedModel: testdata.MinimalZooModel,
		},
		{
			CaseName:      "Model Construct - Partial Build",
			SchemaData:    testdata.BuildResourceData(t, testdata.FullZooTFData),
			BuildOnlyKeys: []string{testdata.NameKey, testdata.SubscriptionsKey, testdata.AddOnsKey},
			ExpectedModel: testdata.PartialZooModel,
		},
	}

	for _, c := range cases {
		var buildOnlyKeys []string

		if c.BuildOnlyKeys != nil {
			buildOnlyKeys = c.BuildOnlyKeys
		} else {
			buildOnlyKeys = make([]string, 0)
		}

		converterResult, err := testdata.Converter.ConvertTFSchemaToAPIModel(c.SchemaData, buildOnlyKeys)

		if err != nil {
			t.Fatalf("# Case '%s' -> An error has occured: %s", c.CaseName, err.Error())
		}

		if !converterResult.Compare(c.ExpectedModel) {
			t.Fatalf("# Case '%s' -> Not Equal.", c.CaseName)
		}
	}
}
