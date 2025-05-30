// Code generated by go-swagger; DO NOT EDIT.

package repository

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

// ListRepositoriesReader is a Reader for the ListRepositories structure.
type ListRepositoriesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListRepositoriesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListRepositoriesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListRepositoriesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListRepositoriesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListRepositoriesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListRepositoriesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListRepositoriesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name}/repositories] listRepositories", response, response.Code())
	}
}

// NewListRepositoriesOK creates a ListRepositoriesOK with default headers values
func NewListRepositoriesOK() *ListRepositoriesOK {
	return &ListRepositoriesOK{}
}

/*
ListRepositoriesOK describes a response with status code 200, with default header values.

Success
*/
type ListRepositoriesOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of repositories
	 */
	XTotalCount int64

	Payload []*models.Repository
}

// IsSuccess returns true when this list repositories o k response has a 2xx status code
func (o *ListRepositoriesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list repositories o k response has a 3xx status code
func (o *ListRepositoriesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list repositories o k response has a 4xx status code
func (o *ListRepositoriesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list repositories o k response has a 5xx status code
func (o *ListRepositoriesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list repositories o k response a status code equal to that given
func (o *ListRepositoriesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list repositories o k response
func (o *ListRepositoriesOK) Code() int {
	return 200
}

func (o *ListRepositoriesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesOK  %+v", 200, o.Payload)
}

func (o *ListRepositoriesOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesOK  %+v", 200, o.Payload)
}

func (o *ListRepositoriesOK) GetPayload() []*models.Repository {
	return o.Payload
}

func (o *ListRepositoriesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRepositoriesBadRequest creates a ListRepositoriesBadRequest with default headers values
func NewListRepositoriesBadRequest() *ListRepositoriesBadRequest {
	return &ListRepositoriesBadRequest{}
}

/*
ListRepositoriesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListRepositoriesBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list repositories bad request response has a 2xx status code
func (o *ListRepositoriesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list repositories bad request response has a 3xx status code
func (o *ListRepositoriesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list repositories bad request response has a 4xx status code
func (o *ListRepositoriesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list repositories bad request response has a 5xx status code
func (o *ListRepositoriesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list repositories bad request response a status code equal to that given
func (o *ListRepositoriesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list repositories bad request response
func (o *ListRepositoriesBadRequest) Code() int {
	return 400
}

func (o *ListRepositoriesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepositoriesBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesBadRequest  %+v", 400, o.Payload)
}

func (o *ListRepositoriesBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRepositoriesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRepositoriesUnauthorized creates a ListRepositoriesUnauthorized with default headers values
func NewListRepositoriesUnauthorized() *ListRepositoriesUnauthorized {
	return &ListRepositoriesUnauthorized{}
}

/*
ListRepositoriesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListRepositoriesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list repositories unauthorized response has a 2xx status code
func (o *ListRepositoriesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list repositories unauthorized response has a 3xx status code
func (o *ListRepositoriesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list repositories unauthorized response has a 4xx status code
func (o *ListRepositoriesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list repositories unauthorized response has a 5xx status code
func (o *ListRepositoriesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list repositories unauthorized response a status code equal to that given
func (o *ListRepositoriesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list repositories unauthorized response
func (o *ListRepositoriesUnauthorized) Code() int {
	return 401
}

func (o *ListRepositoriesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRepositoriesUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListRepositoriesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRepositoriesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRepositoriesForbidden creates a ListRepositoriesForbidden with default headers values
func NewListRepositoriesForbidden() *ListRepositoriesForbidden {
	return &ListRepositoriesForbidden{}
}

/*
ListRepositoriesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListRepositoriesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list repositories forbidden response has a 2xx status code
func (o *ListRepositoriesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list repositories forbidden response has a 3xx status code
func (o *ListRepositoriesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list repositories forbidden response has a 4xx status code
func (o *ListRepositoriesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list repositories forbidden response has a 5xx status code
func (o *ListRepositoriesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list repositories forbidden response a status code equal to that given
func (o *ListRepositoriesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list repositories forbidden response
func (o *ListRepositoriesForbidden) Code() int {
	return 403
}

func (o *ListRepositoriesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesForbidden  %+v", 403, o.Payload)
}

func (o *ListRepositoriesForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesForbidden  %+v", 403, o.Payload)
}

func (o *ListRepositoriesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRepositoriesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRepositoriesNotFound creates a ListRepositoriesNotFound with default headers values
func NewListRepositoriesNotFound() *ListRepositoriesNotFound {
	return &ListRepositoriesNotFound{}
}

/*
ListRepositoriesNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListRepositoriesNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list repositories not found response has a 2xx status code
func (o *ListRepositoriesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list repositories not found response has a 3xx status code
func (o *ListRepositoriesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list repositories not found response has a 4xx status code
func (o *ListRepositoriesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list repositories not found response has a 5xx status code
func (o *ListRepositoriesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list repositories not found response a status code equal to that given
func (o *ListRepositoriesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list repositories not found response
func (o *ListRepositoriesNotFound) Code() int {
	return 404
}

func (o *ListRepositoriesNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesNotFound  %+v", 404, o.Payload)
}

func (o *ListRepositoriesNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesNotFound  %+v", 404, o.Payload)
}

func (o *ListRepositoriesNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRepositoriesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListRepositoriesInternalServerError creates a ListRepositoriesInternalServerError with default headers values
func NewListRepositoriesInternalServerError() *ListRepositoriesInternalServerError {
	return &ListRepositoriesInternalServerError{}
}

/*
ListRepositoriesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListRepositoriesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list repositories internal server error response has a 2xx status code
func (o *ListRepositoriesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list repositories internal server error response has a 3xx status code
func (o *ListRepositoriesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list repositories internal server error response has a 4xx status code
func (o *ListRepositoriesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list repositories internal server error response has a 5xx status code
func (o *ListRepositoriesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list repositories internal server error response a status code equal to that given
func (o *ListRepositoriesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list repositories internal server error response
func (o *ListRepositoriesInternalServerError) Code() int {
	return 500
}

func (o *ListRepositoriesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRepositoriesInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/repositories][%d] listRepositoriesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListRepositoriesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListRepositoriesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
