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

// DeleteRegistryReader is a Reader for the DeleteRegistry structure.
type DeleteRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteRegistryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRegistryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRegistryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewDeleteRegistryPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /registries/{id}] deleteRegistry", response, response.Code())
	}
}

// NewDeleteRegistryOK creates a DeleteRegistryOK with default headers values
func NewDeleteRegistryOK() *DeleteRegistryOK {
	return &DeleteRegistryOK{}
}

/*
DeleteRegistryOK describes a response with status code 200, with default header values.

Success
*/
type DeleteRegistryOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete registry o k response has a 2xx status code
func (o *DeleteRegistryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete registry o k response has a 3xx status code
func (o *DeleteRegistryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry o k response has a 4xx status code
func (o *DeleteRegistryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry o k response has a 5xx status code
func (o *DeleteRegistryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry o k response a status code equal to that given
func (o *DeleteRegistryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete registry o k response
func (o *DeleteRegistryOK) Code() int {
	return 200
}

func (o *DeleteRegistryOK) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryOK ", 200)
}

func (o *DeleteRegistryOK) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryOK ", 200)
}

func (o *DeleteRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteRegistryUnauthorized creates a DeleteRegistryUnauthorized with default headers values
func NewDeleteRegistryUnauthorized() *DeleteRegistryUnauthorized {
	return &DeleteRegistryUnauthorized{}
}

/*
DeleteRegistryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRegistryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete registry unauthorized response has a 2xx status code
func (o *DeleteRegistryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry unauthorized response has a 3xx status code
func (o *DeleteRegistryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry unauthorized response has a 4xx status code
func (o *DeleteRegistryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry unauthorized response has a 5xx status code
func (o *DeleteRegistryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry unauthorized response a status code equal to that given
func (o *DeleteRegistryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete registry unauthorized response
func (o *DeleteRegistryUnauthorized) Code() int {
	return 401
}

func (o *DeleteRegistryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegistryUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRegistryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRegistryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRegistryForbidden creates a DeleteRegistryForbidden with default headers values
func NewDeleteRegistryForbidden() *DeleteRegistryForbidden {
	return &DeleteRegistryForbidden{}
}

/*
DeleteRegistryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRegistryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete registry forbidden response has a 2xx status code
func (o *DeleteRegistryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry forbidden response has a 3xx status code
func (o *DeleteRegistryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry forbidden response has a 4xx status code
func (o *DeleteRegistryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry forbidden response has a 5xx status code
func (o *DeleteRegistryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry forbidden response a status code equal to that given
func (o *DeleteRegistryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete registry forbidden response
func (o *DeleteRegistryForbidden) Code() int {
	return 403
}

func (o *DeleteRegistryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRegistryForbidden) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRegistryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRegistryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRegistryNotFound creates a DeleteRegistryNotFound with default headers values
func NewDeleteRegistryNotFound() *DeleteRegistryNotFound {
	return &DeleteRegistryNotFound{}
}

/*
DeleteRegistryNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteRegistryNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete registry not found response has a 2xx status code
func (o *DeleteRegistryNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry not found response has a 3xx status code
func (o *DeleteRegistryNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry not found response has a 4xx status code
func (o *DeleteRegistryNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry not found response has a 5xx status code
func (o *DeleteRegistryNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry not found response a status code equal to that given
func (o *DeleteRegistryNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete registry not found response
func (o *DeleteRegistryNotFound) Code() int {
	return 404
}

func (o *DeleteRegistryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRegistryNotFound) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRegistryNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRegistryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRegistryPreconditionFailed creates a DeleteRegistryPreconditionFailed with default headers values
func NewDeleteRegistryPreconditionFailed() *DeleteRegistryPreconditionFailed {
	return &DeleteRegistryPreconditionFailed{}
}

/*
DeleteRegistryPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type DeleteRegistryPreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete registry precondition failed response has a 2xx status code
func (o *DeleteRegistryPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry precondition failed response has a 3xx status code
func (o *DeleteRegistryPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry precondition failed response has a 4xx status code
func (o *DeleteRegistryPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete registry precondition failed response has a 5xx status code
func (o *DeleteRegistryPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this delete registry precondition failed response a status code equal to that given
func (o *DeleteRegistryPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the delete registry precondition failed response
func (o *DeleteRegistryPreconditionFailed) Code() int {
	return 412
}

func (o *DeleteRegistryPreconditionFailed) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteRegistryPreconditionFailed) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryPreconditionFailed  %+v", 412, o.Payload)
}

func (o *DeleteRegistryPreconditionFailed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRegistryPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRegistryInternalServerError creates a DeleteRegistryInternalServerError with default headers values
func NewDeleteRegistryInternalServerError() *DeleteRegistryInternalServerError {
	return &DeleteRegistryInternalServerError{}
}

/*
DeleteRegistryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteRegistryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete registry internal server error response has a 2xx status code
func (o *DeleteRegistryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete registry internal server error response has a 3xx status code
func (o *DeleteRegistryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete registry internal server error response has a 4xx status code
func (o *DeleteRegistryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete registry internal server error response has a 5xx status code
func (o *DeleteRegistryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete registry internal server error response a status code equal to that given
func (o *DeleteRegistryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete registry internal server error response
func (o *DeleteRegistryInternalServerError) Code() int {
	return 500
}

func (o *DeleteRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegistryInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /registries/{id}][%d] deleteRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRegistryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
