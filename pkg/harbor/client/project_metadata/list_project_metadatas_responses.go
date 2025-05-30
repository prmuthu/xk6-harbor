// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// ListProjectMetadatasReader is a Reader for the ListProjectMetadatas structure.
type ListProjectMetadatasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListProjectMetadatasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListProjectMetadatasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListProjectMetadatasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListProjectMetadatasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListProjectMetadatasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListProjectMetadatasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListProjectMetadatasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name_or_id}/metadatas/] listProjectMetadatas", response, response.Code())
	}
}

// NewListProjectMetadatasOK creates a ListProjectMetadatasOK with default headers values
func NewListProjectMetadatasOK() *ListProjectMetadatasOK {
	return &ListProjectMetadatasOK{}
}

/*
ListProjectMetadatasOK describes a response with status code 200, with default header values.

Success
*/
type ListProjectMetadatasOK struct {
	Payload map[string]string
}

// IsSuccess returns true when this list project metadatas o k response has a 2xx status code
func (o *ListProjectMetadatasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list project metadatas o k response has a 3xx status code
func (o *ListProjectMetadatasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project metadatas o k response has a 4xx status code
func (o *ListProjectMetadatasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project metadatas o k response has a 5xx status code
func (o *ListProjectMetadatasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this list project metadatas o k response a status code equal to that given
func (o *ListProjectMetadatasOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the list project metadatas o k response
func (o *ListProjectMetadatasOK) Code() int {
	return 200
}

func (o *ListProjectMetadatasOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasOK  %+v", 200, o.Payload)
}

func (o *ListProjectMetadatasOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasOK  %+v", 200, o.Payload)
}

func (o *ListProjectMetadatasOK) GetPayload() map[string]string {
	return o.Payload
}

func (o *ListProjectMetadatasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListProjectMetadatasBadRequest creates a ListProjectMetadatasBadRequest with default headers values
func NewListProjectMetadatasBadRequest() *ListProjectMetadatasBadRequest {
	return &ListProjectMetadatasBadRequest{}
}

/*
ListProjectMetadatasBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type ListProjectMetadatasBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list project metadatas bad request response has a 2xx status code
func (o *ListProjectMetadatasBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project metadatas bad request response has a 3xx status code
func (o *ListProjectMetadatasBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project metadatas bad request response has a 4xx status code
func (o *ListProjectMetadatasBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project metadatas bad request response has a 5xx status code
func (o *ListProjectMetadatasBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this list project metadatas bad request response a status code equal to that given
func (o *ListProjectMetadatasBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the list project metadatas bad request response
func (o *ListProjectMetadatasBadRequest) Code() int {
	return 400
}

func (o *ListProjectMetadatasBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectMetadatasBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasBadRequest  %+v", 400, o.Payload)
}

func (o *ListProjectMetadatasBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListProjectMetadatasUnauthorized creates a ListProjectMetadatasUnauthorized with default headers values
func NewListProjectMetadatasUnauthorized() *ListProjectMetadatasUnauthorized {
	return &ListProjectMetadatasUnauthorized{}
}

/*
ListProjectMetadatasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListProjectMetadatasUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list project metadatas unauthorized response has a 2xx status code
func (o *ListProjectMetadatasUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project metadatas unauthorized response has a 3xx status code
func (o *ListProjectMetadatasUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project metadatas unauthorized response has a 4xx status code
func (o *ListProjectMetadatasUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project metadatas unauthorized response has a 5xx status code
func (o *ListProjectMetadatasUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this list project metadatas unauthorized response a status code equal to that given
func (o *ListProjectMetadatasUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the list project metadatas unauthorized response
func (o *ListProjectMetadatasUnauthorized) Code() int {
	return 401
}

func (o *ListProjectMetadatasUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProjectMetadatasUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasUnauthorized  %+v", 401, o.Payload)
}

func (o *ListProjectMetadatasUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListProjectMetadatasForbidden creates a ListProjectMetadatasForbidden with default headers values
func NewListProjectMetadatasForbidden() *ListProjectMetadatasForbidden {
	return &ListProjectMetadatasForbidden{}
}

/*
ListProjectMetadatasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListProjectMetadatasForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list project metadatas forbidden response has a 2xx status code
func (o *ListProjectMetadatasForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project metadatas forbidden response has a 3xx status code
func (o *ListProjectMetadatasForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project metadatas forbidden response has a 4xx status code
func (o *ListProjectMetadatasForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project metadatas forbidden response has a 5xx status code
func (o *ListProjectMetadatasForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this list project metadatas forbidden response a status code equal to that given
func (o *ListProjectMetadatasForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the list project metadatas forbidden response
func (o *ListProjectMetadatasForbidden) Code() int {
	return 403
}

func (o *ListProjectMetadatasForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectMetadatasForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasForbidden  %+v", 403, o.Payload)
}

func (o *ListProjectMetadatasForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListProjectMetadatasNotFound creates a ListProjectMetadatasNotFound with default headers values
func NewListProjectMetadatasNotFound() *ListProjectMetadatasNotFound {
	return &ListProjectMetadatasNotFound{}
}

/*
ListProjectMetadatasNotFound describes a response with status code 404, with default header values.

Not found
*/
type ListProjectMetadatasNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list project metadatas not found response has a 2xx status code
func (o *ListProjectMetadatasNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project metadatas not found response has a 3xx status code
func (o *ListProjectMetadatasNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project metadatas not found response has a 4xx status code
func (o *ListProjectMetadatasNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this list project metadatas not found response has a 5xx status code
func (o *ListProjectMetadatasNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this list project metadatas not found response a status code equal to that given
func (o *ListProjectMetadatasNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the list project metadatas not found response
func (o *ListProjectMetadatasNotFound) Code() int {
	return 404
}

func (o *ListProjectMetadatasNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectMetadatasNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasNotFound  %+v", 404, o.Payload)
}

func (o *ListProjectMetadatasNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListProjectMetadatasInternalServerError creates a ListProjectMetadatasInternalServerError with default headers values
func NewListProjectMetadatasInternalServerError() *ListProjectMetadatasInternalServerError {
	return &ListProjectMetadatasInternalServerError{}
}

/*
ListProjectMetadatasInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListProjectMetadatasInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this list project metadatas internal server error response has a 2xx status code
func (o *ListProjectMetadatasInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this list project metadatas internal server error response has a 3xx status code
func (o *ListProjectMetadatasInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list project metadatas internal server error response has a 4xx status code
func (o *ListProjectMetadatasInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this list project metadatas internal server error response has a 5xx status code
func (o *ListProjectMetadatasInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this list project metadatas internal server error response a status code equal to that given
func (o *ListProjectMetadatasInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the list project metadatas internal server error response
func (o *ListProjectMetadatasInternalServerError) Code() int {
	return 500
}

func (o *ListProjectMetadatasInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectMetadatasInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/metadatas/][%d] listProjectMetadatasInternalServerError  %+v", 500, o.Payload)
}

func (o *ListProjectMetadatasInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *ListProjectMetadatasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
