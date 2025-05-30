// Code generated by go-swagger; DO NOT EDIT.

package preheat

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

// ListPoliciesReader is a Reader for the ListPolicies structure.
type ListPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListPoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name}/preheat/policies] ListPolicies", response, response.Code())
	}
}

// NewListPoliciesOK creates a ListPoliciesOK with default headers values
func NewListPoliciesOK() *ListPoliciesOK {
	return &ListPoliciesOK{}
}

/*
ListPoliciesOK describes a response with status code 200, with default header values.

List preheat policies success
*/
type ListPoliciesOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of policies
	 */
	XTotalCount int64

	Payload []*models.PreheatPolicy
}

// IsSuccess returns true when this list policies o k response has a 2xx status code
func (o *ListPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list policies o k response has a 3xx status code
func (o *ListPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies o k response has a 4xx status code
func (o *ListPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list policies o k response has a 5xx status code
func (o *ListPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies o k response a status code equal to that given
func (o *ListPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list policies o k response
func (o *ListPoliciesOK) Code() int {
	return 200
}

func (o *ListPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListPoliciesOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesOK  %+v", 200, o.Payload)
}

func (o *ListPoliciesOK) GetPayload() []*models.PreheatPolicy {
	return o.Payload
}

func (o *ListPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListPoliciesBadRequest creates a ListPoliciesBadRequest with default headers values
func NewListPoliciesBadRequest() *ListPoliciesBadRequest {
	return &ListPoliciesBadRequest{}
}

/*
ListPoliciesBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListPoliciesBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list policies bad request response has a 2xx status code
func (o *ListPoliciesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies bad request response has a 3xx status code
func (o *ListPoliciesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies bad request response has a 4xx status code
func (o *ListPoliciesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list policies bad request response has a 5xx status code
func (o *ListPoliciesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies bad request response a status code equal to that given
func (o *ListPoliciesBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list policies bad request response
func (o *ListPoliciesBadRequest) Code() int {
	return 400
}

func (o *ListPoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *ListPoliciesBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *ListPoliciesBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListPoliciesUnauthorized creates a ListPoliciesUnauthorized with default headers values
func NewListPoliciesUnauthorized() *ListPoliciesUnauthorized {
	return &ListPoliciesUnauthorized{}
}

/*
ListPoliciesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListPoliciesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list policies unauthorized response has a 2xx status code
func (o *ListPoliciesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies unauthorized response has a 3xx status code
func (o *ListPoliciesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies unauthorized response has a 4xx status code
func (o *ListPoliciesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list policies unauthorized response has a 5xx status code
func (o *ListPoliciesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies unauthorized response a status code equal to that given
func (o *ListPoliciesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list policies unauthorized response
func (o *ListPoliciesUnauthorized) Code() int {
	return 401
}

func (o *ListPoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPoliciesUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPoliciesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListPoliciesForbidden creates a ListPoliciesForbidden with default headers values
func NewListPoliciesForbidden() *ListPoliciesForbidden {
	return &ListPoliciesForbidden{}
}

/*
ListPoliciesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListPoliciesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list policies forbidden response has a 2xx status code
func (o *ListPoliciesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies forbidden response has a 3xx status code
func (o *ListPoliciesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies forbidden response has a 4xx status code
func (o *ListPoliciesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list policies forbidden response has a 5xx status code
func (o *ListPoliciesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list policies forbidden response a status code equal to that given
func (o *ListPoliciesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list policies forbidden response
func (o *ListPoliciesForbidden) Code() int {
	return 403
}

func (o *ListPoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *ListPoliciesForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesForbidden  %+v", 403, o.Payload)
}

func (o *ListPoliciesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListPoliciesInternalServerError creates a ListPoliciesInternalServerError with default headers values
func NewListPoliciesInternalServerError() *ListPoliciesInternalServerError {
	return &ListPoliciesInternalServerError{}
}

/*
ListPoliciesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListPoliciesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list policies internal server error response has a 2xx status code
func (o *ListPoliciesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list policies internal server error response has a 3xx status code
func (o *ListPoliciesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list policies internal server error response has a 4xx status code
func (o *ListPoliciesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list policies internal server error response has a 5xx status code
func (o *ListPoliciesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list policies internal server error response a status code equal to that given
func (o *ListPoliciesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list policies internal server error response
func (o *ListPoliciesInternalServerError) Code() int {
	return 500
}

func (o *ListPoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPoliciesInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies][%d] listPoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPoliciesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListPoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
