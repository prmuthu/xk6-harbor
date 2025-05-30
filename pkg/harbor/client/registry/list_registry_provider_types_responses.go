// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// ListRegistryProviderTypesReader is a Reader for the ListRegistryProviderTypes structure.
type ListRegistryProviderTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegistryProviderTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegistryProviderTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRegistryProviderTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRegistryProviderTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRegistryProviderTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /replication/adapters] listRegistryProviderTypes", response, response.Code())
	}
}

// NewListRegistryProviderTypesOK creates a ListRegistryProviderTypesOK with default headers values
func NewListRegistryProviderTypesOK() *ListRegistryProviderTypesOK {
	return &ListRegistryProviderTypesOK{}
}

/*
ListRegistryProviderTypesOK describes a response with status code 200, with default header values.

Success.
*/
type ListRegistryProviderTypesOK struct {
	Payload []string
}

// IsSuccess returns true when this list registry provider types o k response has a 2xx status code
func (o *ListRegistryProviderTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list registry provider types o k response has a 3xx status code
func (o *ListRegistryProviderTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider types o k response has a 4xx status code
func (o *ListRegistryProviderTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list registry provider types o k response has a 5xx status code
func (o *ListRegistryProviderTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list registry provider types o k response a status code equal to that given
func (o *ListRegistryProviderTypesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list registry provider types o k response
func (o *ListRegistryProviderTypesOK) Code() int {
	return 200
}

func (o *ListRegistryProviderTypesOK) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesOK  %+v", 200, o.Payload)
}

func (o *ListRegistryProviderTypesOK) String() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesOK  %+v", 200, o.Payload)
}

func (o *ListRegistryProviderTypesOK) GetPayload() []string {
	return o.Payload
}

func (o *ListRegistryProviderTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderTypesUnauthorized creates a ListRegistryProviderTypesUnauthorized with default headers values
func NewListRegistryProviderTypesUnauthorized() *ListRegistryProviderTypesUnauthorized {
	return &ListRegistryProviderTypesUnauthorized{}
}

/*
ListRegistryProviderTypesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListRegistryProviderTypesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list registry provider types unauthorized response has a 2xx status code
func (o *ListRegistryProviderTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list registry provider types unauthorized response has a 3xx status code
func (o *ListRegistryProviderTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider types unauthorized response has a 4xx status code
func (o *ListRegistryProviderTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list registry provider types unauthorized response has a 5xx status code
func (o *ListRegistryProviderTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list registry provider types unauthorized response a status code equal to that given
func (o *ListRegistryProviderTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list registry provider types unauthorized response
func (o *ListRegistryProviderTypesUnauthorized) Code() int {
	return 401
}

func (o *ListRegistryProviderTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRegistryProviderTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRegistryProviderTypesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRegistryProviderTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRegistryProviderTypesForbidden creates a ListRegistryProviderTypesForbidden with default headers values
func NewListRegistryProviderTypesForbidden() *ListRegistryProviderTypesForbidden {
	return &ListRegistryProviderTypesForbidden{}
}

/*
ListRegistryProviderTypesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListRegistryProviderTypesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list registry provider types forbidden response has a 2xx status code
func (o *ListRegistryProviderTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list registry provider types forbidden response has a 3xx status code
func (o *ListRegistryProviderTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider types forbidden response has a 4xx status code
func (o *ListRegistryProviderTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list registry provider types forbidden response has a 5xx status code
func (o *ListRegistryProviderTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list registry provider types forbidden response a status code equal to that given
func (o *ListRegistryProviderTypesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list registry provider types forbidden response
func (o *ListRegistryProviderTypesForbidden) Code() int {
	return 403
}

func (o *ListRegistryProviderTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistryProviderTypesForbidden) String() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistryProviderTypesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRegistryProviderTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRegistryProviderTypesInternalServerError creates a ListRegistryProviderTypesInternalServerError with default headers values
func NewListRegistryProviderTypesInternalServerError() *ListRegistryProviderTypesInternalServerError {
	return &ListRegistryProviderTypesInternalServerError{}
}

/*
ListRegistryProviderTypesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListRegistryProviderTypesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list registry provider types internal server error response has a 2xx status code
func (o *ListRegistryProviderTypesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list registry provider types internal server error response has a 3xx status code
func (o *ListRegistryProviderTypesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider types internal server error response has a 4xx status code
func (o *ListRegistryProviderTypesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list registry provider types internal server error response has a 5xx status code
func (o *ListRegistryProviderTypesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list registry provider types internal server error response a status code equal to that given
func (o *ListRegistryProviderTypesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list registry provider types internal server error response
func (o *ListRegistryProviderTypesInternalServerError) Code() int {
	return 500
}

func (o *ListRegistryProviderTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRegistryProviderTypesInternalServerError) String() string {
	return fmt.Sprintf("[GET /replication/adapters][%d] listRegistryProviderTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRegistryProviderTypesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRegistryProviderTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
