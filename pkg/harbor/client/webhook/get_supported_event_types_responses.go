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

// GetSupportedEventTypesReader is a Reader for the GetSupportedEventTypes structure.
type GetSupportedEventTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSupportedEventTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSupportedEventTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSupportedEventTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSupportedEventTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSupportedEventTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name_or_id}/webhook/events] GetSupportedEventTypes", response, response.Code())
	}
}

// NewGetSupportedEventTypesOK creates a GetSupportedEventTypesOK with default headers values
func NewGetSupportedEventTypesOK() *GetSupportedEventTypesOK {
	return &GetSupportedEventTypesOK{}
}

/*
GetSupportedEventTypesOK describes a response with status code 200, with default header values.

Success
*/
type GetSupportedEventTypesOK struct {
	Payload *models.SupportedWebhookEventTypes
}

// IsSuccess returns true when this get supported event types o k response has a 2xx status code
func (o *GetSupportedEventTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get supported event types o k response has a 3xx status code
func (o *GetSupportedEventTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported event types o k response has a 4xx status code
func (o *GetSupportedEventTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get supported event types o k response has a 5xx status code
func (o *GetSupportedEventTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get supported event types o k response a status code equal to that given
func (o *GetSupportedEventTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get supported event types o k response
func (o *GetSupportedEventTypesOK) Code() int {
	return 200
}

func (o *GetSupportedEventTypesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesOK  %+v", 200, o.Payload)
}

func (o *GetSupportedEventTypesOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesOK  %+v", 200, o.Payload)
}

func (o *GetSupportedEventTypesOK) GetPayload() *models.SupportedWebhookEventTypes {
	return o.Payload
}

func (o *GetSupportedEventTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SupportedWebhookEventTypes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSupportedEventTypesUnauthorized creates a GetSupportedEventTypesUnauthorized with default headers values
func NewGetSupportedEventTypesUnauthorized() *GetSupportedEventTypesUnauthorized {
	return &GetSupportedEventTypesUnauthorized{}
}

/*
GetSupportedEventTypesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSupportedEventTypesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get supported event types unauthorized response has a 2xx status code
func (o *GetSupportedEventTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supported event types unauthorized response has a 3xx status code
func (o *GetSupportedEventTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported event types unauthorized response has a 4xx status code
func (o *GetSupportedEventTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supported event types unauthorized response has a 5xx status code
func (o *GetSupportedEventTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get supported event types unauthorized response a status code equal to that given
func (o *GetSupportedEventTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get supported event types unauthorized response
func (o *GetSupportedEventTypesUnauthorized) Code() int {
	return 401
}

func (o *GetSupportedEventTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSupportedEventTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSupportedEventTypesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSupportedEventTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSupportedEventTypesForbidden creates a GetSupportedEventTypesForbidden with default headers values
func NewGetSupportedEventTypesForbidden() *GetSupportedEventTypesForbidden {
	return &GetSupportedEventTypesForbidden{}
}

/*
GetSupportedEventTypesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSupportedEventTypesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get supported event types forbidden response has a 2xx status code
func (o *GetSupportedEventTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supported event types forbidden response has a 3xx status code
func (o *GetSupportedEventTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported event types forbidden response has a 4xx status code
func (o *GetSupportedEventTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get supported event types forbidden response has a 5xx status code
func (o *GetSupportedEventTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get supported event types forbidden response a status code equal to that given
func (o *GetSupportedEventTypesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get supported event types forbidden response
func (o *GetSupportedEventTypesForbidden) Code() int {
	return 403
}

func (o *GetSupportedEventTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetSupportedEventTypesForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetSupportedEventTypesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSupportedEventTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSupportedEventTypesInternalServerError creates a GetSupportedEventTypesInternalServerError with default headers values
func NewGetSupportedEventTypesInternalServerError() *GetSupportedEventTypesInternalServerError {
	return &GetSupportedEventTypesInternalServerError{}
}

/*
GetSupportedEventTypesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSupportedEventTypesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get supported event types internal server error response has a 2xx status code
func (o *GetSupportedEventTypesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get supported event types internal server error response has a 3xx status code
func (o *GetSupportedEventTypesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get supported event types internal server error response has a 4xx status code
func (o *GetSupportedEventTypesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get supported event types internal server error response has a 5xx status code
func (o *GetSupportedEventTypesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get supported event types internal server error response a status code equal to that given
func (o *GetSupportedEventTypesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get supported event types internal server error response
func (o *GetSupportedEventTypesInternalServerError) Code() int {
	return 500
}

func (o *GetSupportedEventTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSupportedEventTypesInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/events][%d] getSupportedEventTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSupportedEventTypesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSupportedEventTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
