/*
Copyright © 2023 TeraSky, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testdata

var (
	// Terraform Schema Data.
	FullZooTFData = map[string]interface{}{
		NameKey:   "FullZoo",
		WidthKey:  120.12,
		LengthKey: 45.9,
		EmployeesKey: []interface{}{
			map[string]interface{}{
				TitleKey:    "Manager",
				FullNameKey: "John Smith",
				AgeKey:      35,
			},
			map[string]interface{}{
				TitleKey:    "GateKeeper",
				FullNameKey: "Brandon Lee",
				AgeKey:      63,
			},
			map[string]interface{}{
				TitleKey:    "SeniorAnimalKeeper",
				FullNameKey: "David King",
				AgeKey:      24,
			},
		},
		GateKey: []interface{}{
			map[string]interface{}{
				EmployeesKey: []interface{}{
					map[string]interface{}{
						TitleKey:    "GateKeeper",
						FullNameKey: "Brandon Lee",
						AgeKey:      63,
					},
				},
				OpeningHourKey: "10:00",
				ClosingHourKey: "22:00",
			},
		},
		CagesKey: []interface{}{
			map[string]interface{}{
				EmployeesKey: []interface{}{
					map[string]interface{}{
						TitleKey:    "SeniorAnimalKeeper",
						FullNameKey: "David King",
						AgeKey:      24,
					},
				},
				AnimalNameKey: "Lion",
				LabelsKey: map[string]interface{}{
					"Food":         "Lamb Meat",
					"SourceRegion": "Africa",
					"SleepTime":    "22:00-06:00",
				},
			},
			map[string]interface{}{
				EmployeesKey: []interface{}{
					map[string]interface{}{
						TitleKey:    "SeniorAnimalKeeper",
						FullNameKey: "David King",
						AgeKey:      24,
					},
				},
				AnimalNameKey: "Elephant",
				LabelsKey: map[string]interface{}{
					"Food":         "Tall Grass",
					"SourceRegion": "Africa",
					"SleepTime":    "20:00-04:00",
				},
			},
		},
		LocationKey: []interface{}{
			map[string]interface{}{
				LongitudeKey: "46.15",
				LatitudeKey:  "78.19",
			},
		},
		SubscriptionsKey: []interface{}{
			map[string]interface{}{
				IndividualClientKey: []interface{}{
					map[string]interface{}{
						IDKey:        "299483774",
						FirstNameKey: "Danny",
						LastNameKey:  "Johnson",
					},
					map[string]interface{}{
						IDKey:        "983773421",
						FirstNameKey: "Mike",
						LastNameKey:  "Birmingham",
					},
				},
				GroupClientKey: []interface{}{"The Animal Lovers", "Schnitzel Company", "Walmart"},
			},
		},
		AddOnsKey: map[string]interface{}{
			"Fountain": true,
			"ExtraLights": []map[string]interface{}{
				{
					"Color": "Red",
					"Count": 3,
				},
				{
					"Color": "Blue",
					"Count": 3,
				},
				{
					"Color": "Green",
					"Count": 3,
				},
			},
			"Restaurants": []interface{}{"KFC", "McDonald's", "Olive Garden"},
		},
	}

	SlimZooTFData = map[string]interface{}{
		NameKey:   "FullZoo",
		WidthKey:  120.12,
		LengthKey: 45.9,
		EmployeesKey: []interface{}{
			map[string]interface{}{
				TitleKey:    "Manager",
				FullNameKey: "John Smith",
				AgeKey:      35,
			},
			map[string]interface{}{
				TitleKey:    "GateKeeper",
				FullNameKey: "Brandon Lee",
				AgeKey:      63,
			},
			map[string]interface{}{
				TitleKey:    "SeniorAnimalKeeper",
				FullNameKey: "David King",
				AgeKey:      24,
			},
		},
		GateKey: []interface{}{
			map[string]interface{}{
				EmployeesKey: []interface{}{
					map[string]interface{}{
						TitleKey:    "GateKeeper",
						FullNameKey: "Brandon Lee",
						AgeKey:      63,
					},
				},
				OpeningHourKey: "10:00",
				ClosingHourKey: "22:00",
			},
		},
		CagesKey: []interface{}{
			map[string]interface{}{
				EmployeesKey: []interface{}{
					map[string]interface{}{
						TitleKey:    "SeniorAnimalKeeper",
						FullNameKey: "David King",
						AgeKey:      24,
					},
				},
				AnimalNameKey: "Lion",
				LabelsKey: map[string]interface{}{
					"Food":         "Lamb Meat",
					"SourceRegion": "Africa",
					"SleepTime":    "22:00-06:00",
				},
			},
		},
		LocationKey: []interface{}{
			map[string]interface{}{
				LongitudeKey: "46.15",
				LatitudeKey:  "78.19",
			},
		},
		SubscriptionsKey: []interface{}{
			map[string]interface{}{
				IndividualClientKey: []interface{}{
					map[string]interface{}{
						IDKey:        "299483774",
						FirstNameKey: "Danny",
						LastNameKey:  "Johnson",
					},
				},
			},
		},
		AddOnsKey: map[string]interface{}{
			"Fountain":    true,
			"Restaurants": []interface{}{"KFC", "McDonald's", "Olive Garden"},
		},
	}

	MinimalZooTFData = map[string]interface{}{
		NameKey:   "FullZoo",
		WidthKey:  120.12,
		LengthKey: 45.9,
		EmployeesKey: []interface{}{
			map[string]interface{}{
				TitleKey:    "Manager",
				FullNameKey: "John Smith",
				AgeKey:      35,
			},
		},
		GateKey: []interface{}{
			map[string]interface{}{
				EmployeesKey: []interface{}{
					map[string]interface{}{
						TitleKey:    "Manager",
						FullNameKey: "John Smith",
						AgeKey:      35,
					},
				},
				OpeningHourKey: "10:00",
				ClosingHourKey: "22:00",
			},
		},
		LocationKey: []interface{}{
			map[string]interface{}{
				LongitudeKey: "46.15",
				LatitudeKey:  "78.19",
			},
		},
	}

	// Struct Model Data.
	FullZooModel = &Zoo{
		Name: "FullZoo",
		Employees: []*ZooEmployee{
			{
				Title:    "Manager",
				FullName: "John Smith",
				Age:      35,
			},
			{
				Title:    "GateKeeper",
				FullName: "Brandon Lee",
				Age:      63,
			},
			{
				Title:    "SeniorAnimalKeeper",
				FullName: "David King",
				Age:      24,
			},
		},
		Gate: &ZooGate{
			Employees: []*ZooEmployee{
				{
					Title:    "GateKeeper",
					FullName: "Brandon Lee",
					Age:      63,
				},
			},
			OpeningHour: "10:00",
			ClosingHour: "22:00",
		},
		Cages: []*ZooCage{
			{
				Employees: []*ZooEmployee{
					{
						Title:    "SeniorAnimalKeeper",
						FullName: "David King",
						Age:      24,
					},
				},
				AnimalName: "Lion",
				Labels: map[string]string{
					"Food":         "Lamb Meat",
					"SourceRegion": "Africa",
					"SleepTime":    "22:00-06:00",
				},
			},
			{
				Employees: []*ZooEmployee{
					{
						Title:    "SeniorAnimalKeeper",
						FullName: "David King",
						Age:      24,
					},
				},
				AnimalName: "Elephant",
				Labels: map[string]string{
					"Food":         "Tall Grass",
					"SourceRegion": "Africa",
					"SleepTime":    "20:00-04:00",
				},
			},
		},
		Location: &ZooLocation{
			Longitude: "46.15",
			Latitude:  "78.19",
		},
		Area: &ZooArea{
			Width:  120.12,
			Length: 45.9,
		},
		Subscriptions: []*ZooSubscription{
			{
				IndividualClient: &Person{
					ID:        "299483774",
					FirstName: "Danny",
					LastName:  "Johnson",
				},
			},
			{
				IndividualClient: &Person{
					ID:        "983773421",
					FirstName: "Mike",
					LastName:  "Birmingham",
				},
			},
			{
				GroupClient: &Group{
					Name: "The Animal Lovers",
				},
			},
			{
				GroupClient: &Group{
					Name: "Schnitzel Company",
				},
			},
			{
				GroupClient: &Group{
					Name: "Walmart",
				},
			},
		},
		AddOns: []*ZooAddOn{
			{
				Name: "ExtraLights",
				Value: []map[string]interface{}{
					{
						"Color": "Red",
						"Count": 3,
					},
					{
						"Color": "Blue",
						"Count": 3,
					},
					{
						"Color": "Green",
						"Count": 3,
					},
				},
			},
			{
				Name:  "Fountain",
				Value: true,
			},
			{
				Name:  "Restaurants",
				Value: []string{"KFC", "McDonald's", "Olive Garden"},
			},
		},
	}

	SlimZooModel = &Zoo{
		Name: "FullZoo",
		Employees: []*ZooEmployee{
			{
				Title:    "Manager",
				FullName: "John Smith",
				Age:      35,
			},
			{
				Title:    "GateKeeper",
				FullName: "Brandon Lee",
				Age:      63,
			},
			{
				Title:    "SeniorAnimalKeeper",
				FullName: "David King",
				Age:      24,
			},
		},
		Gate: &ZooGate{
			Employees: []*ZooEmployee{
				{
					Title:    "GateKeeper",
					FullName: "Brandon Lee",
					Age:      63,
				},
			},
			OpeningHour: "10:00",
			ClosingHour: "22:00",
		},
		Cages: []*ZooCage{
			{
				Employees: []*ZooEmployee{
					{
						Title:    "SeniorAnimalKeeper",
						FullName: "David King",
						Age:      24,
					},
				},
				AnimalName: "Lion",
				Labels: map[string]string{
					"Food":         "Lamb Meat",
					"SourceRegion": "Africa",
					"SleepTime":    "22:00-06:00",
				},
			},
		},
		Location: &ZooLocation{
			Longitude: "46.15",
			Latitude:  "78.19",
		},
		Area: &ZooArea{
			Width:  120.12,
			Length: 45.9,
		},
		Subscriptions: []*ZooSubscription{
			{
				IndividualClient: &Person{
					ID:        "299483774",
					FirstName: "Danny",
					LastName:  "Johnson",
				},
			},
		},
		AddOns: []*ZooAddOn{
			{
				Name:  "Fountain",
				Value: true,
			},
			{
				Name:  "Restaurants",
				Value: []string{"KFC", "McDonald's", "Olive Garden"},
			},
		},
	}

	MinimalZooModel = &Zoo{
		Name: "FullZoo",
		Employees: []*ZooEmployee{
			{
				Title:    "Manager",
				FullName: "John Smith",
				Age:      35,
			},
		},
		Gate: &ZooGate{
			Employees: []*ZooEmployee{
				{
					Title:    "Manager",
					FullName: "John Smith",
					Age:      35,
				},
			},
			OpeningHour: "10:00",
			ClosingHour: "22:00",
		},
		Location: &ZooLocation{
			Longitude: "46.15",
			Latitude:  "78.19",
		},
		Area: &ZooArea{
			Width:  120.12,
			Length: 45.9,
		},
	}

	PartialZooModel = &Zoo{
		Name: "FullZoo",
		Subscriptions: []*ZooSubscription{
			{
				IndividualClient: &Person{
					ID:        "299483774",
					FirstName: "Danny",
					LastName:  "Johnson",
				},
			},
			{
				IndividualClient: &Person{
					ID:        "983773421",
					FirstName: "Mike",
					LastName:  "Birmingham",
				},
			},
			{
				GroupClient: &Group{
					Name: "The Animal Lovers",
				},
			},
			{
				GroupClient: &Group{
					Name: "Schnitzel Company",
				},
			},
			{
				GroupClient: &Group{
					Name: "Walmart",
				},
			},
		},
		AddOns: []*ZooAddOn{
			{
				Name: "ExtraLights",
				Value: []map[string]interface{}{
					{
						"Color": "Red",
						"Count": 3,
					},
					{
						"Color": "Blue",
						"Count": 3,
					},
					{
						"Color": "Green",
						"Count": 3,
					},
				},
			},
			{
				Name:  "Fountain",
				Value: true,
			},
			{
				Name:  "Restaurants",
				Value: []string{"KFC", "McDonald's", "Olive Garden"},
			},
		},
	}
)
