// Code generated by go-swagger; DO NOT EDIT.

package artifact

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// ListAccessoriesReader is a Reader for the ListAccessories structure.
type ListAccessoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAccessoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAccessoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAccessoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAccessoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAccessoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListAccessoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAccessoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories] listAccessories", response, response.Code())
	}
}

// NewListAccessoriesOK creates a ListAccessoriesOK with default headers values
func NewListAccessoriesOK() *ListAccessoriesOK {
	return &ListAccessoriesOK{}
}

/*
ListAccessoriesOK describes a response with status code 200, with default header values.

Success
*/
type ListAccessoriesOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of accessories
	 */
	XTotalCount int64

	Payload []*models.Accessory
}

// IsSuccess returns true when this list accessories o k response has a 2xx status code
func (o *ListAccessoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list accessories o k response has a 3xx status code
func (o *ListAccessoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list accessories o k response has a 4xx status code
func (o *ListAccessoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list accessories o k response has a 5xx status code
func (o *ListAccessoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list accessories o k response a status code equal to that given
func (o *ListAccessoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list accessories o k response
func (o *ListAccessoriesOK) Code() int {
	return 200
}

func (o *ListAccessoriesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesOK  %+v", 200, o.Payload)
}

func (o *ListAccessoriesOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesOK  %+v", 200, o.Payload)
}

func (o *ListAccessoriesOK) GetPayload() []*models.Accessory {
	return o.Payload
}

func (o *ListAccessoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAccessoriesBadRequest creates a ListAccessoriesBadRequest with default headers values
func NewListAccessoriesBadRequest() *ListAccessoriesBadRequest {
	return &ListAccessoriesBadRequest{}
}

/*
ListAccessoriesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListAccessoriesBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list accessories bad request response has a 2xx status code
func (o *ListAccessoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list accessories bad request response has a 3xx status code
func (o *ListAccessoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list accessories bad request response has a 4xx status code
func (o *ListAccessoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list accessories bad request response has a 5xx status code
func (o *ListAccessoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list accessories bad request response a status code equal to that given
func (o *ListAccessoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list accessories bad request response
func (o *ListAccessoriesBadRequest) Code() int {
	return 400
}

func (o *ListAccessoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesBadRequest  %+v", 400, o.Payload)
}

func (o *ListAccessoriesBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesBadRequest  %+v", 400, o.Payload)
}

func (o *ListAccessoriesBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAccessoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAccessoriesUnauthorized creates a ListAccessoriesUnauthorized with default headers values
func NewListAccessoriesUnauthorized() *ListAccessoriesUnauthorized {
	return &ListAccessoriesUnauthorized{}
}

/*
ListAccessoriesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListAccessoriesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list accessories unauthorized response has a 2xx status code
func (o *ListAccessoriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list accessories unauthorized response has a 3xx status code
func (o *ListAccessoriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list accessories unauthorized response has a 4xx status code
func (o *ListAccessoriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list accessories unauthorized response has a 5xx status code
func (o *ListAccessoriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list accessories unauthorized response a status code equal to that given
func (o *ListAccessoriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list accessories unauthorized response
func (o *ListAccessoriesUnauthorized) Code() int {
	return 401
}

func (o *ListAccessoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAccessoriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAccessoriesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAccessoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAccessoriesForbidden creates a ListAccessoriesForbidden with default headers values
func NewListAccessoriesForbidden() *ListAccessoriesForbidden {
	return &ListAccessoriesForbidden{}
}

/*
ListAccessoriesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListAccessoriesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list accessories forbidden response has a 2xx status code
func (o *ListAccessoriesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list accessories forbidden response has a 3xx status code
func (o *ListAccessoriesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list accessories forbidden response has a 4xx status code
func (o *ListAccessoriesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list accessories forbidden response has a 5xx status code
func (o *ListAccessoriesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list accessories forbidden response a status code equal to that given
func (o *ListAccessoriesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list accessories forbidden response
func (o *ListAccessoriesForbidden) Code() int {
	return 403
}

func (o *ListAccessoriesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesForbidden  %+v", 403, o.Payload)
}

func (o *ListAccessoriesForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesForbidden  %+v", 403, o.Payload)
}

func (o *ListAccessoriesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAccessoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAccessoriesNotFound creates a ListAccessoriesNotFound with default headers values
func NewListAccessoriesNotFound() *ListAccessoriesNotFound {
	return &ListAccessoriesNotFound{}
}

/*
ListAccessoriesNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListAccessoriesNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list accessories not found response has a 2xx status code
func (o *ListAccessoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list accessories not found response has a 3xx status code
func (o *ListAccessoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list accessories not found response has a 4xx status code
func (o *ListAccessoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list accessories not found response has a 5xx status code
func (o *ListAccessoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list accessories not found response a status code equal to that given
func (o *ListAccessoriesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list accessories not found response
func (o *ListAccessoriesNotFound) Code() int {
	return 404
}

func (o *ListAccessoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesNotFound  %+v", 404, o.Payload)
}

func (o *ListAccessoriesNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesNotFound  %+v", 404, o.Payload)
}

func (o *ListAccessoriesNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAccessoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListAccessoriesInternalServerError creates a ListAccessoriesInternalServerError with default headers values
func NewListAccessoriesInternalServerError() *ListAccessoriesInternalServerError {
	return &ListAccessoriesInternalServerError{}
}

/*
ListAccessoriesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListAccessoriesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list accessories internal server error response has a 2xx status code
func (o *ListAccessoriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list accessories internal server error response has a 3xx status code
func (o *ListAccessoriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list accessories internal server error response has a 4xx status code
func (o *ListAccessoriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list accessories internal server error response has a 5xx status code
func (o *ListAccessoriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list accessories internal server error response a status code equal to that given
func (o *ListAccessoriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list accessories internal server error response
func (o *ListAccessoriesInternalServerError) Code() int {
	return 500
}

func (o *ListAccessoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAccessoriesInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories/{repository_name}/artifacts/{reference}/accessories][%d] listAccessoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAccessoriesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListAccessoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
