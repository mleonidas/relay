// Code generated by go-swagger; DO NOT EDIT.

package workflow_triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteWorkflowTriggerParams creates a new DeleteWorkflowTriggerParams object
// with the default values initialized.
func NewDeleteWorkflowTriggerParams() *DeleteWorkflowTriggerParams {
	var ()
	return &DeleteWorkflowTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWorkflowTriggerParamsWithTimeout creates a new DeleteWorkflowTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWorkflowTriggerParamsWithTimeout(timeout time.Duration) *DeleteWorkflowTriggerParams {
	var ()
	return &DeleteWorkflowTriggerParams{

		timeout: timeout,
	}
}

// NewDeleteWorkflowTriggerParamsWithContext creates a new DeleteWorkflowTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWorkflowTriggerParamsWithContext(ctx context.Context) *DeleteWorkflowTriggerParams {
	var ()
	return &DeleteWorkflowTriggerParams{

		Context: ctx,
	}
}

// NewDeleteWorkflowTriggerParamsWithHTTPClient creates a new DeleteWorkflowTriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWorkflowTriggerParamsWithHTTPClient(client *http.Client) *DeleteWorkflowTriggerParams {
	var ()
	return &DeleteWorkflowTriggerParams{
		HTTPClient: client,
	}
}

/*DeleteWorkflowTriggerParams contains all the parameters to send to the API endpoint
for the delete workflow trigger operation typically these are written to a http.Request
*/
type DeleteWorkflowTriggerParams struct {

	/*WorkflowName
	  Workflow name

	*/
	WorkflowName string
	/*WorkflowTriggerID
	  The workflow trigger ID to reference

	*/
	WorkflowTriggerID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) WithTimeout(timeout time.Duration) *DeleteWorkflowTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) WithContext(ctx context.Context) *DeleteWorkflowTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) WithHTTPClient(client *http.Client) *DeleteWorkflowTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflowName adds the workflowName to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) WithWorkflowName(workflowName string) *DeleteWorkflowTriggerParams {
	o.SetWorkflowName(workflowName)
	return o
}

// SetWorkflowName adds the workflowName to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) SetWorkflowName(workflowName string) {
	o.WorkflowName = workflowName
}

// WithWorkflowTriggerID adds the workflowTriggerID to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) WithWorkflowTriggerID(workflowTriggerID string) *DeleteWorkflowTriggerParams {
	o.SetWorkflowTriggerID(workflowTriggerID)
	return o
}

// SetWorkflowTriggerID adds the workflowTriggerId to the delete workflow trigger params
func (o *DeleteWorkflowTriggerParams) SetWorkflowTriggerID(workflowTriggerID string) {
	o.WorkflowTriggerID = workflowTriggerID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWorkflowTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workflowName
	if err := r.SetPathParam("workflowName", o.WorkflowName); err != nil {
		return err
	}

	// path param workflowTriggerId
	if err := r.SetPathParam("workflowTriggerId", o.WorkflowTriggerID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}