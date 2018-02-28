// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDcimVirtualChassisDeleteParams creates a new DcimVirtualChassisDeleteParams object
// with the default values initialized.
func NewDcimVirtualChassisDeleteParams() *DcimVirtualChassisDeleteParams {
	var ()
	return &DcimVirtualChassisDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimVirtualChassisDeleteParamsWithTimeout creates a new DcimVirtualChassisDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimVirtualChassisDeleteParamsWithTimeout(timeout time.Duration) *DcimVirtualChassisDeleteParams {
	var ()
	return &DcimVirtualChassisDeleteParams{

		timeout: timeout,
	}
}

// NewDcimVirtualChassisDeleteParamsWithContext creates a new DcimVirtualChassisDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimVirtualChassisDeleteParamsWithContext(ctx context.Context) *DcimVirtualChassisDeleteParams {
	var ()
	return &DcimVirtualChassisDeleteParams{

		Context: ctx,
	}
}

// NewDcimVirtualChassisDeleteParamsWithHTTPClient creates a new DcimVirtualChassisDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimVirtualChassisDeleteParamsWithHTTPClient(client *http.Client) *DcimVirtualChassisDeleteParams {
	var ()
	return &DcimVirtualChassisDeleteParams{
		HTTPClient: client,
	}
}

/*DcimVirtualChassisDeleteParams contains all the parameters to send to the API endpoint
for the dcim virtual chassis delete operation typically these are written to a http.Request
*/
type DcimVirtualChassisDeleteParams struct {

	/*ID
	  A unique integer value identifying this virtual chassis.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithTimeout(timeout time.Duration) *DcimVirtualChassisDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithContext(ctx context.Context) *DcimVirtualChassisDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithHTTPClient(client *http.Client) *DcimVirtualChassisDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) WithID(id int64) *DcimVirtualChassisDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim virtual chassis delete params
func (o *DcimVirtualChassisDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimVirtualChassisDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
