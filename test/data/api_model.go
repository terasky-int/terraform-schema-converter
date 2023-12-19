/*
Copyright © 2023 VMware, Inc. All Rights Reserved.
SPDX-License-Identifier: MPL-2.0
*/

package testdata

import (
	"encoding/json"

	"github.com/go-openapi/swag"
)

type Zoo struct {
	Name string `json:"name,omitempty"`

	Employees []*ZooEmployee `json:"employees,omitempty"`

	Gate *ZooGate `json:"gate,omitempty"`

	Cages []*ZooCage `json:"cages,omitempty"`

	Location *ZooLocation `json:"location,omitempty"`

	Area *ZooArea `json:"area,omitempty"`

	Subscriptions []*ZooSubscription `json:"subscriptions,omitempty"`

	AddOns []*ZooAddOn `json:"addOns,omitempty"`
}

// MarshalBinary interface implementation.
func (s *Zoo) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *Zoo) UnmarshalBinary(b []byte) error {
	var res Zoo

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

func (s *Zoo) Compare(other *Zoo) bool {
	selfJSON := make(map[string]interface{})
	otherJSON := make(map[string]interface{})
	selfBytes, _ := s.MarshalBinary()
	otherBytes, _ := other.MarshalBinary()
	_ = json.Unmarshal(selfBytes, &selfJSON)
	_ = json.Unmarshal(otherBytes, &otherJSON)

	return compareJSON(selfJSON, otherJSON)
}

type ZooEmployee struct {
	Title string `json:"title,omitempty"`

	FullName string `json:"fullName,omitempty"`

	Age int `json:"age,omitempty"`
}

// MarshalBinary interface implementation.
func (s *ZooEmployee) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *ZooEmployee) UnmarshalBinary(b []byte) error {
	var res ZooEmployee

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type ZooGate struct {
	Employees []*ZooEmployee `json:"employees,omitempty"`

	OpeningHour string `json:"openingHour,omitempty"`

	ClosingHour string `json:"closingHour,omitempty"`
}

// MarshalBinary interface implementation.
func (s *ZooGate) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *ZooGate) UnmarshalBinary(b []byte) error {
	var res ZooGate

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type ZooCage struct {
	Employees []*ZooEmployee `json:"employees,omitempty"`

	AnimalName string `json:"animalName,omitempty"`

	Labels map[string]string `json:"labels,omitempty"`
}

// MarshalBinary interface implementation.
func (s *ZooCage) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *ZooCage) UnmarshalBinary(b []byte) error {
	var res ZooCage

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type ZooLocation struct {
	Longitude string `json:"longitude,omitempty"`

	Latitude string `json:"latitude,omitempty"`
}

// MarshalBinary interface implementation.
func (s *ZooLocation) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *ZooLocation) UnmarshalBinary(b []byte) error {
	var res ZooLocation

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type ZooArea struct {
	Width float32 `json:"width,omitempty"`

	Length float32 `json:"length,omitempty"`
}

// MarshalBinary interface implementation.
func (s *ZooArea) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *ZooArea) UnmarshalBinary(b []byte) error {
	var res ZooArea

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type ZooSubscription struct {
	IndividualClient *Person `json:"person,omitempty"`

	GroupClient *Group `json:"group,omitempty"`
}

// MarshalBinary interface implementation.
func (s *ZooSubscription) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *ZooSubscription) UnmarshalBinary(b []byte) error {
	var res ZooSubscription

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type Person struct {
	ID string `json:"id,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	LastName string `json:"lastName,omitempty"`
}

// MarshalBinary interface implementation.
func (s *Person) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *Person) UnmarshalBinary(b []byte) error {
	var res Person

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type Group struct {
	Name string `json:"name,omitempty"`
}

// MarshalBinary interface implementation.
func (s *Group) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *Group) UnmarshalBinary(b []byte) error {
	var res Group

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}

type ZooAddOn struct {
	Name string `json:"name,omitempty"`

	Value interface{} `json:"value,omitempty"`
}

// MarshalBinary interface implementation.
func (s *ZooAddOn) MarshalBinary() ([]byte, error) {
	if s == nil {
		return nil, nil
	}

	return swag.WriteJSON(s)
}

// UnmarshalBinary interface implementation.
func (s *ZooAddOn) UnmarshalBinary(b []byte) error {
	var res ZooAddOn

	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}

	*s = res

	return nil
}
