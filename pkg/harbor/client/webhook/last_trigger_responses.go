// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// LastTriggerReader is a Reader for the LastTrigger structure.
type LastTriggerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LastTriggerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLastTriggerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewLastTriggerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewLastTriggerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewLastTriggerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewLastTriggerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name_or_id}/webhook/lasttrigger] LastTrigger", response, response.Code())
	}
}

// NewLastTriggerOK creates a LastTriggerOK with default headers values
func NewLastTriggerOK() *LastTriggerOK {
	return &LastTriggerOK{}
}

/*
LastTriggerOK describes a response with status code 200, with default header values.

Test webhook connection successfully.
*/
type LastTriggerOK struct {
	Payload []*models.WebhookLastTrigger
}

// IsSuccess returns true when this last trigger o k response has a 2xx status code
func (o *LastTriggerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this last trigger o k response has a 3xx status code
func (o *LastTriggerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this last trigger o k response has a 4xx status code
func (o *LastTriggerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this last trigger o k response has a 5xx status code
func (o *LastTriggerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this last trigger o k response a status code equal to that given
func (o *LastTriggerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the last trigger o k response
func (o *LastTriggerOK) Code() int {
	return 200
}

func (o *LastTriggerOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerOK  %+v", 200, o.Payload)
}

func (o *LastTriggerOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerOK  %+v", 200, o.Payload)
}

func (o *LastTriggerOK) GetPayload() []*models.WebhookLastTrigger {
	return o.Payload
}

func (o *LastTriggerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLastTriggerBadRequest creates a LastTriggerBadRequest with default headers values
func NewLastTriggerBadRequest() *LastTriggerBadRequest {
	return &LastTriggerBadRequest{}
}

/*
LastTriggerBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type LastTriggerBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this last trigger bad request response has a 2xx status code
func (o *LastTriggerBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this last trigger bad request response has a 3xx status code
func (o *LastTriggerBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this last trigger bad request response has a 4xx status code
func (o *LastTriggerBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this last trigger bad request response has a 5xx status code
func (o *LastTriggerBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this last trigger bad request response a status code equal to that given
func (o *LastTriggerBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the last trigger bad request response
func (o *LastTriggerBadRequest) Code() int {
	return 400
}

func (o *LastTriggerBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *LastTriggerBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerBadRequest  %+v", 400, o.Payload)
}

func (o *LastTriggerBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLastTriggerUnauthorized creates a LastTriggerUnauthorized with default headers values
func NewLastTriggerUnauthorized() *LastTriggerUnauthorized {
	return &LastTriggerUnauthorized{}
}

/*
LastTriggerUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type LastTriggerUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this last trigger unauthorized response has a 2xx status code
func (o *LastTriggerUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this last trigger unauthorized response has a 3xx status code
func (o *LastTriggerUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this last trigger unauthorized response has a 4xx status code
func (o *LastTriggerUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this last trigger unauthorized response has a 5xx status code
func (o *LastTriggerUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this last trigger unauthorized response a status code equal to that given
func (o *LastTriggerUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the last trigger unauthorized response
func (o *LastTriggerUnauthorized) Code() int {
	return 401
}

func (o *LastTriggerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *LastTriggerUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerUnauthorized  %+v", 401, o.Payload)
}

func (o *LastTriggerUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLastTriggerForbidden creates a LastTriggerForbidden with default headers values
func NewLastTriggerForbidden() *LastTriggerForbidden {
	return &LastTriggerForbidden{}
}

/*
LastTriggerForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type LastTriggerForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this last trigger forbidden response has a 2xx status code
func (o *LastTriggerForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this last trigger forbidden response has a 3xx status code
func (o *LastTriggerForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this last trigger forbidden response has a 4xx status code
func (o *LastTriggerForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this last trigger forbidden response has a 5xx status code
func (o *LastTriggerForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this last trigger forbidden response a status code equal to that given
func (o *LastTriggerForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the last trigger forbidden response
func (o *LastTriggerForbidden) Code() int {
	return 403
}

func (o *LastTriggerForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerForbidden  %+v", 403, o.Payload)
}

func (o *LastTriggerForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerForbidden  %+v", 403, o.Payload)
}

func (o *LastTriggerForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewLastTriggerInternalServerError creates a LastTriggerInternalServerError with default headers values
func NewLastTriggerInternalServerError() *LastTriggerInternalServerError {
	return &LastTriggerInternalServerError{}
}

/*
LastTriggerInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type LastTriggerInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this last trigger internal server error response has a 2xx status code
func (o *LastTriggerInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this last trigger internal server error response has a 3xx status code
func (o *LastTriggerInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this last trigger internal server error response has a 4xx status code
func (o *LastTriggerInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this last trigger internal server error response has a 5xx status code
func (o *LastTriggerInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this last trigger internal server error response a status code equal to that given
func (o *LastTriggerInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the last trigger internal server error response
func (o *LastTriggerInternalServerError) Code() int {
	return 500
}

func (o *LastTriggerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *LastTriggerInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/lasttrigger][%d] lastTriggerInternalServerError  %+v", 500, o.Payload)
}

func (o *LastTriggerInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *LastTriggerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
