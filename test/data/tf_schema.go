/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testdata

import (
	"encoding/json"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// Root Keys.
	NameKey          = "name"
	WidthKey         = "width"
	LengthKey        = "length"
	EmployeesKey     = "employee"
	GateKey          = "gate"
	CagesKey         = "cage"
	LocationKey      = "location"
	SubscriptionsKey = "subscription"
	AddOnsKey        = "add_ons"

	// Employees Directive Keys.
	TitleKey    = "title"
	FullNameKey = "full_name"
	AgeKey      = "age"

	// Gate Directive Keys.
	OpeningHourKey = "opening_hour"
	ClosingHourKey = "closing_hour"

	// Cages Directive Keys.
	AnimalNameKey = "animal_name"
	LabelsKey     = "labels"

	// Location Directive Keys.
	LongitudeKey = "longitude"
	LatitudeKey  = "latitude"

	// Subscriptions Directive Keys.
	IndividualClientKey = "individual_client"
	GroupClientKey      = "group_clients"

	// IndividualClient Directive Keys.
	IDKey        = "id"
	FirstNameKey = "first_name"
	LastNameKey  = "last_name"
)

// TFSchema is the root resource schema.
var TFSchema = map[string]*schema.Schema{
	NameKey:          nameSchema,
	WidthKey:         widthSchema,
	LengthKey:        lengthSchema,
	EmployeesKey:     employeesSchema,
	GateKey:          gateSchema,
	CagesKey:         cagesSchema,
	LocationKey:      locationSchema,
	SubscriptionsKey: subscriptionsSchema,
	AddOnsKey:        addOnsSchema,
}

// Inner schemas.
var (
	nameSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "The name of the zoo.",
		Required:    true,
		ForceNew:    true,
	}

	widthSchema = &schema.Schema{
		Type:        schema.TypeFloat,
		Description: "The width size of the zoo.",
		Required:    true,
	}

	lengthSchema = &schema.Schema{
		Type:        schema.TypeFloat,
		Description: "The length size of the zoo.",
		Required:    true,
	}

	employeesSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The list of the zoo employees.",
		MinItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				TitleKey: {
					Type:        schema.TypeString,
					Description: "Employee title.",
					Required:    true,
				},
				FullNameKey: {
					Type:        schema.TypeString,
					Description: "Employee's full name.",
					Required:    true,
				},
				AgeKey: {
					Type:        schema.TypeInt,
					Description: "Employee age.",
					Required:    true,
				},
			},
		},
	}

	gateSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The zoo gate settings.",
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				EmployeesKey: employeesSchema,
				OpeningHourKey: {
					Type:        schema.TypeString,
					Description: "The opening hour of the zoo gate.",
					Required:    true,
				},
				ClosingHourKey: {
					Type:        schema.TypeString,
					Description: "The closing hour of the zoo gate.",
					Required:    true,
				},
			},
		},
	}

	cagesSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The cages exist in the zoo.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				EmployeesKey: employeesSchema,
				AnimalNameKey: {
					Type:        schema.TypeString,
					Description: "The name of the animal.",
					Required:    true,
				},
				LabelsKey: {
					Type:        schema.TypeMap,
					Description: "Labels regarding the animal (what to feed, sleep time, etc..).",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	locationSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "The location of the zoo (central point).",
		MinItems:    1,
		MaxItems:    1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				LongitudeKey: {
					Type:        schema.TypeString,
					Description: "The opening hour of the zoo gate.",
					Required:    true,
				},
				LatitudeKey: {
					Type:        schema.TypeString,
					Description: "The closing hour of the zoo gate.",
					Required:    true,
				},
			},
		},
	}

	subscriptionsSchema = &schema.Schema{
		Type:        schema.TypeList,
		Description: "Clients with free pass subscription.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				IndividualClientKey: {
					Type:        schema.TypeList,
					Description: "Individual clients data.",
					Optional:    true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							IDKey: {
								Type:        schema.TypeString,
								Description: "Client's ID.",
								Required:    true,
							},
							FirstNameKey: {
								Type:        schema.TypeString,
								Description: "Client's first name.",
								Required:    true,
							},
							LastNameKey: {
								Type:        schema.TypeString,
								Description: "Client's last name.",
								Required:    true,
							},
						},
					},
				},
				GroupClientKey: {
					Type:        schema.TypeList,
					Description: "Group client names.",
					Optional:    true,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}

	addOnsSchema = &schema.Schema{
		Type:        schema.TypeString,
		Description: "JSON encoded add ons to the zoo.",
		Optional:    true,
	}
)

var rootResourceKeys = []string{
	NameKey,
	WidthKey,
	LengthKey,
	EmployeesKey,
	GateKey,
	CagesKey,
	LocationKey,
	SubscriptionsKey,
	AddOnsKey,
}

func BuildResourceData(t *testing.T, tfData map[string]interface{}) *schema.ResourceData {
	var tfDataCopy map[string]interface{}

	if len(tfData) == 0 {
		tfDataCopy = tfData
	} else {
		tfDataCopy = copyMap(tfData)

		if !isEmptyInterface(tfDataCopy[AddOnsKey]) {
			addOnsJSONBytes, _ := json.Marshal(tfDataCopy[AddOnsKey])
			tfDataCopy[AddOnsKey] = string(addOnsJSONBytes)
		}
	}

	resourceData := schema.TestResourceDataRaw(t, TFSchema, tfDataCopy)

	return resourceData
}

func CompareResourceData(firstResourceData, secondResourceData *schema.ResourceData) bool {
	for _, key := range rootResourceKeys {
		firstValue := firstResourceData.Get(key)
		secondValue := secondResourceData.Get(key)

		if key == AddOnsKey {
			_ = json.Unmarshal([]byte(firstValue.(string)), &firstValue)
			_ = json.Unmarshal([]byte(secondValue.(string)), &secondValue)
		}

		if !compareJSON(firstValue, secondValue) {
			return false
		}
	}

	return true
}
