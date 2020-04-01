// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkflowTriggerIdentifier A unique identifier for a workflow trigger
//
// swagger:model WorkflowTriggerIdentifier
type WorkflowTriggerIdentifier struct {

	// The unique identifier for the trigger
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this workflow trigger identifier
func (m *WorkflowTriggerIdentifier) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WorkflowTriggerIdentifier) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WorkflowTriggerIdentifier) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkflowTriggerIdentifier) UnmarshalBinary(b []byte) error {
	var res WorkflowTriggerIdentifier
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}