// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// PingInstancesReader is a Reader for the PingInstances structure.
type PingInstancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingInstancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingInstancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPingInstancesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPingInstancesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPingInstancesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPingInstancesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /p2p/preheat/instances/ping] PingInstances", response, response.Code())
	}
}

// NewPingInstancesOK creates a PingInstancesOK with default headers values
func NewPingInstancesOK() *PingInstancesOK {
	return &PingInstancesOK{}
}

/*
PingInstancesOK describes a response with status code 200, with default header values.

Success
*/
type PingInstancesOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this ping instances o k response has a 2xx status code
func (o *PingInstancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ping instances o k response has a 3xx status code
func (o *PingInstancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping instances o k response has a 4xx status code
func (o *PingInstancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping instances o k response has a 5xx status code
func (o *PingInstancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ping instances o k response a status code equal to that given
func (o *PingInstancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the ping instances o k response
func (o *PingInstancesOK) Code() int {
	return 200
}

func (o *PingInstancesOK) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesOK ", 200)
}

func (o *PingInstancesOK) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesOK ", 200)
}

func (o *PingInstancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewPingInstancesBadRequest creates a PingInstancesBadRequest with default headers values
func NewPingInstancesBadRequest() *PingInstancesBadRequest {
	return &PingInstancesBadRequest{}
}

/*
PingInstancesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type PingInstancesBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping instances bad request response has a 2xx status code
func (o *PingInstancesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping instances bad request response has a 3xx status code
func (o *PingInstancesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping instances bad request response has a 4xx status code
func (o *PingInstancesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping instances bad request response has a 5xx status code
func (o *PingInstancesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this ping instances bad request response a status code equal to that given
func (o *PingInstancesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the ping instances bad request response
func (o *PingInstancesBadRequest) Code() int {
	return 400
}

func (o *PingInstancesBadRequest) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *PingInstancesBadRequest) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesBadRequest  %+v", 400, o.Payload)
}

func (o *PingInstancesBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingInstancesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingInstancesUnauthorized creates a PingInstancesUnauthorized with default headers values
func NewPingInstancesUnauthorized() *PingInstancesUnauthorized {
	return &PingInstancesUnauthorized{}
}

/*
PingInstancesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PingInstancesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping instances unauthorized response has a 2xx status code
func (o *PingInstancesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping instances unauthorized response has a 3xx status code
func (o *PingInstancesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping instances unauthorized response has a 4xx status code
func (o *PingInstancesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping instances unauthorized response has a 5xx status code
func (o *PingInstancesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this ping instances unauthorized response a status code equal to that given
func (o *PingInstancesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the ping instances unauthorized response
func (o *PingInstancesUnauthorized) Code() int {
	return 401
}

func (o *PingInstancesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *PingInstancesUnauthorized) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesUnauthorized  %+v", 401, o.Payload)
}

func (o *PingInstancesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingInstancesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPingInstancesNotFound creates a PingInstancesNotFound with default headers values
func NewPingInstancesNotFound() *PingInstancesNotFound {
	return &PingInstancesNotFound{}
}

/*
PingInstancesNotFound describes a response with status code 404, with default header values.

Instance not found (when instance is provided by ID).
*/
type PingInstancesNotFound struct {
}

// IsSuccess returns true when this ping instances not found response has a 2xx status code
func (o *PingInstancesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping instances not found response has a 3xx status code
func (o *PingInstancesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping instances not found response has a 4xx status code
func (o *PingInstancesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this ping instances not found response has a 5xx status code
func (o *PingInstancesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this ping instances not found response a status code equal to that given
func (o *PingInstancesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the ping instances not found response
func (o *PingInstancesNotFound) Code() int {
	return 404
}

func (o *PingInstancesNotFound) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesNotFound ", 404)
}

func (o *PingInstancesNotFound) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesNotFound ", 404)
}

func (o *PingInstancesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPingInstancesInternalServerError creates a PingInstancesInternalServerError with default headers values
func NewPingInstancesInternalServerError() *PingInstancesInternalServerError {
	return &PingInstancesInternalServerError{}
}

/*
PingInstancesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type PingInstancesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this ping instances internal server error response has a 2xx status code
func (o *PingInstancesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this ping instances internal server error response has a 3xx status code
func (o *PingInstancesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ping instances internal server error response has a 4xx status code
func (o *PingInstancesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this ping instances internal server error response has a 5xx status code
func (o *PingInstancesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this ping instances internal server error response a status code equal to that given
func (o *PingInstancesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the ping instances internal server error response
func (o *PingInstancesInternalServerError) Code() int {
	return 500
}

func (o *PingInstancesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *PingInstancesInternalServerError) String() string {
	return fmt.Sprintf("[POST /p2p/preheat/instances/ping][%d] pingInstancesInternalServerError  %+v", 500, o.Payload)
}

func (o *PingInstancesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *PingInstancesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
