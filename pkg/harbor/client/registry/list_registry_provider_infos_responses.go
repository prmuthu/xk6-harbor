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

// ListRegistryProviderInfosReader is a Reader for the ListRegistryProviderInfos structure.
type ListRegistryProviderInfosReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRegistryProviderInfosReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRegistryProviderInfosOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListRegistryProviderInfosUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRegistryProviderInfosForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRegistryProviderInfosInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /replication/adapterinfos] listRegistryProviderInfos", response, response.Code())
	}
}

// NewListRegistryProviderInfosOK creates a ListRegistryProviderInfosOK with default headers values
func NewListRegistryProviderInfosOK() *ListRegistryProviderInfosOK {
	return &ListRegistryProviderInfosOK{}
}

/*
ListRegistryProviderInfosOK describes a response with status code 200, with default header values.

Success.
*/
type ListRegistryProviderInfosOK struct {
	Payload map[string]models.RegistryProviderInfo
}

// IsSuccess returns true when this list registry provider infos o k response has a 2xx status code
func (o *ListRegistryProviderInfosOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list registry provider infos o k response has a 3xx status code
func (o *ListRegistryProviderInfosOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider infos o k response has a 4xx status code
func (o *ListRegistryProviderInfosOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list registry provider infos o k response has a 5xx status code
func (o *ListRegistryProviderInfosOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list registry provider infos o k response a status code equal to that given
func (o *ListRegistryProviderInfosOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list registry provider infos o k response
func (o *ListRegistryProviderInfosOK) Code() int {
	return 200
}

func (o *ListRegistryProviderInfosOK) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosOK  %+v", 200, o.Payload)
}

func (o *ListRegistryProviderInfosOK) String() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosOK  %+v", 200, o.Payload)
}

func (o *ListRegistryProviderInfosOK) GetPayload() map[string]models.RegistryProviderInfo {
	return o.Payload
}

func (o *ListRegistryProviderInfosOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListRegistryProviderInfosUnauthorized creates a ListRegistryProviderInfosUnauthorized with default headers values
func NewListRegistryProviderInfosUnauthorized() *ListRegistryProviderInfosUnauthorized {
	return &ListRegistryProviderInfosUnauthorized{}
}

/*
ListRegistryProviderInfosUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListRegistryProviderInfosUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list registry provider infos unauthorized response has a 2xx status code
func (o *ListRegistryProviderInfosUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list registry provider infos unauthorized response has a 3xx status code
func (o *ListRegistryProviderInfosUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider infos unauthorized response has a 4xx status code
func (o *ListRegistryProviderInfosUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list registry provider infos unauthorized response has a 5xx status code
func (o *ListRegistryProviderInfosUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list registry provider infos unauthorized response a status code equal to that given
func (o *ListRegistryProviderInfosUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list registry provider infos unauthorized response
func (o *ListRegistryProviderInfosUnauthorized) Code() int {
	return 401
}

func (o *ListRegistryProviderInfosUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRegistryProviderInfosUnauthorized) String() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRegistryProviderInfosUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRegistryProviderInfosUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRegistryProviderInfosForbidden creates a ListRegistryProviderInfosForbidden with default headers values
func NewListRegistryProviderInfosForbidden() *ListRegistryProviderInfosForbidden {
	return &ListRegistryProviderInfosForbidden{}
}

/*
ListRegistryProviderInfosForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListRegistryProviderInfosForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list registry provider infos forbidden response has a 2xx status code
func (o *ListRegistryProviderInfosForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list registry provider infos forbidden response has a 3xx status code
func (o *ListRegistryProviderInfosForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider infos forbidden response has a 4xx status code
func (o *ListRegistryProviderInfosForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list registry provider infos forbidden response has a 5xx status code
func (o *ListRegistryProviderInfosForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list registry provider infos forbidden response a status code equal to that given
func (o *ListRegistryProviderInfosForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list registry provider infos forbidden response
func (o *ListRegistryProviderInfosForbidden) Code() int {
	return 403
}

func (o *ListRegistryProviderInfosForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistryProviderInfosForbidden) String() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosForbidden  %+v", 403, o.Payload)
}

func (o *ListRegistryProviderInfosForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRegistryProviderInfosForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRegistryProviderInfosInternalServerError creates a ListRegistryProviderInfosInternalServerError with default headers values
func NewListRegistryProviderInfosInternalServerError() *ListRegistryProviderInfosInternalServerError {
	return &ListRegistryProviderInfosInternalServerError{}
}

/*
ListRegistryProviderInfosInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListRegistryProviderInfosInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list registry provider infos internal server error response has a 2xx status code
func (o *ListRegistryProviderInfosInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list registry provider infos internal server error response has a 3xx status code
func (o *ListRegistryProviderInfosInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list registry provider infos internal server error response has a 4xx status code
func (o *ListRegistryProviderInfosInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list registry provider infos internal server error response has a 5xx status code
func (o *ListRegistryProviderInfosInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list registry provider infos internal server error response a status code equal to that given
func (o *ListRegistryProviderInfosInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list registry provider infos internal server error response
func (o *ListRegistryProviderInfosInternalServerError) Code() int {
	return 500
}

func (o *ListRegistryProviderInfosInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRegistryProviderInfosInternalServerError) String() string {
	return fmt.Sprintf("[GET /replication/adapterinfos][%d] listRegistryProviderInfosInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRegistryProviderInfosInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRegistryProviderInfosInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
