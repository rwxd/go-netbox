// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimRearPortTemplatesBulkPartialUpdateReader is a Reader for the DcimRearPortTemplatesBulkPartialUpdate structure.
type DcimRearPortTemplatesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortTemplatesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortTemplatesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRearPortTemplatesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRearPortTemplatesBulkPartialUpdateOK creates a DcimRearPortTemplatesBulkPartialUpdateOK with default headers values
func NewDcimRearPortTemplatesBulkPartialUpdateOK() *DcimRearPortTemplatesBulkPartialUpdateOK {
	return &DcimRearPortTemplatesBulkPartialUpdateOK{}
}

/*
DcimRearPortTemplatesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimRearPortTemplatesBulkPartialUpdateOK dcim rear port templates bulk partial update o k
*/
type DcimRearPortTemplatesBulkPartialUpdateOK struct {
	Payload *models.RearPortTemplate
}

// IsSuccess returns true when this dcim rear port templates bulk partial update o k response has a 2xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rear port templates bulk partial update o k response has a 3xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rear port templates bulk partial update o k response has a 4xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rear port templates bulk partial update o k response has a 5xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rear port templates bulk partial update o k response a status code equal to that given
func (o *DcimRearPortTemplatesBulkPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimRearPortTemplatesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-port-templates/][%d] dcimRearPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortTemplatesBulkPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/rear-port-templates/][%d] dcimRearPortTemplatesBulkPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRearPortTemplatesBulkPartialUpdateOK) GetPayload() *models.RearPortTemplate {
	return o.Payload
}

func (o *DcimRearPortTemplatesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimRearPortTemplatesBulkPartialUpdateDefault creates a DcimRearPortTemplatesBulkPartialUpdateDefault with default headers values
func NewDcimRearPortTemplatesBulkPartialUpdateDefault(code int) *DcimRearPortTemplatesBulkPartialUpdateDefault {
	return &DcimRearPortTemplatesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimRearPortTemplatesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimRearPortTemplatesBulkPartialUpdateDefault dcim rear port templates bulk partial update default
*/
type DcimRearPortTemplatesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim rear port templates bulk partial update default response
func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim rear port templates bulk partial update default response has a 2xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rear port templates bulk partial update default response has a 3xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rear port templates bulk partial update default response has a 4xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rear port templates bulk partial update default response has a 5xx status code
func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rear port templates bulk partial update default response a status code equal to that given
func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/rear-port-templates/][%d] dcim_rear-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/rear-port-templates/][%d] dcim_rear-port-templates_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRearPortTemplatesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}