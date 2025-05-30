// Code generated by go-swagger; DO NOT EDIT.

package member

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// GetProjectMemberReader is a Reader for the GetProjectMember structure.
type GetProjectMemberReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectMemberReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectMemberOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectMemberBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectMemberUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectMemberForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProjectMemberNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectMemberInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name_or_id}/members/{mid}] getProjectMember", response, response.Code())
	}
}

// NewGetProjectMemberOK creates a GetProjectMemberOK with default headers values
func NewGetProjectMemberOK() *GetProjectMemberOK {
	return &GetProjectMemberOK{}
}

/*
GetProjectMemberOK describes a response with status code 200, with default header values.

Project member retrieved successfully.
*/
type GetProjectMemberOK struct {
	Payload *models.ProjectMemberEntity
}

// IsSuccess returns true when this get project member o k response has a 2xx status code
func (o *GetProjectMemberOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get project member o k response has a 3xx status code
func (o *GetProjectMemberOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project member o k response has a 4xx status code
func (o *GetProjectMemberOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project member o k response has a 5xx status code
func (o *GetProjectMemberOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get project member o k response a status code equal to that given
func (o *GetProjectMemberOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get project member o k response
func (o *GetProjectMemberOK) Code() int {
	return 200
}

func (o *GetProjectMemberOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberOK  %+v", 200, o.Payload)
}

func (o *GetProjectMemberOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberOK  %+v", 200, o.Payload)
}

func (o *GetProjectMemberOK) GetPayload() *models.ProjectMemberEntity {
	return o.Payload
}

func (o *GetProjectMemberOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectMemberEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProjectMemberBadRequest creates a GetProjectMemberBadRequest with default headers values
func NewGetProjectMemberBadRequest() *GetProjectMemberBadRequest {
	return &GetProjectMemberBadRequest{}
}

/*
GetProjectMemberBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetProjectMemberBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get project member bad request response has a 2xx status code
func (o *GetProjectMemberBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project member bad request response has a 3xx status code
func (o *GetProjectMemberBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project member bad request response has a 4xx status code
func (o *GetProjectMemberBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project member bad request response has a 5xx status code
func (o *GetProjectMemberBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get project member bad request response a status code equal to that given
func (o *GetProjectMemberBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get project member bad request response
func (o *GetProjectMemberBadRequest) Code() int {
	return 400
}

func (o *GetProjectMemberBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectMemberBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberBadRequest  %+v", 400, o.Payload)
}

func (o *GetProjectMemberBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberUnauthorized creates a GetProjectMemberUnauthorized with default headers values
func NewGetProjectMemberUnauthorized() *GetProjectMemberUnauthorized {
	return &GetProjectMemberUnauthorized{}
}

/*
GetProjectMemberUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetProjectMemberUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get project member unauthorized response has a 2xx status code
func (o *GetProjectMemberUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project member unauthorized response has a 3xx status code
func (o *GetProjectMemberUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project member unauthorized response has a 4xx status code
func (o *GetProjectMemberUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project member unauthorized response has a 5xx status code
func (o *GetProjectMemberUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get project member unauthorized response a status code equal to that given
func (o *GetProjectMemberUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get project member unauthorized response
func (o *GetProjectMemberUnauthorized) Code() int {
	return 401
}

func (o *GetProjectMemberUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectMemberUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProjectMemberUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberForbidden creates a GetProjectMemberForbidden with default headers values
func NewGetProjectMemberForbidden() *GetProjectMemberForbidden {
	return &GetProjectMemberForbidden{}
}

/*
GetProjectMemberForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetProjectMemberForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get project member forbidden response has a 2xx status code
func (o *GetProjectMemberForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project member forbidden response has a 3xx status code
func (o *GetProjectMemberForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project member forbidden response has a 4xx status code
func (o *GetProjectMemberForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project member forbidden response has a 5xx status code
func (o *GetProjectMemberForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get project member forbidden response a status code equal to that given
func (o *GetProjectMemberForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get project member forbidden response
func (o *GetProjectMemberForbidden) Code() int {
	return 403
}

func (o *GetProjectMemberForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectMemberForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberForbidden  %+v", 403, o.Payload)
}

func (o *GetProjectMemberForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberNotFound creates a GetProjectMemberNotFound with default headers values
func NewGetProjectMemberNotFound() *GetProjectMemberNotFound {
	return &GetProjectMemberNotFound{}
}

/*
GetProjectMemberNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetProjectMemberNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get project member not found response has a 2xx status code
func (o *GetProjectMemberNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project member not found response has a 3xx status code
func (o *GetProjectMemberNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project member not found response has a 4xx status code
func (o *GetProjectMemberNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get project member not found response has a 5xx status code
func (o *GetProjectMemberNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get project member not found response a status code equal to that given
func (o *GetProjectMemberNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get project member not found response
func (o *GetProjectMemberNotFound) Code() int {
	return 404
}

func (o *GetProjectMemberNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectMemberNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberNotFound  %+v", 404, o.Payload)
}

func (o *GetProjectMemberNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectMemberInternalServerError creates a GetProjectMemberInternalServerError with default headers values
func NewGetProjectMemberInternalServerError() *GetProjectMemberInternalServerError {
	return &GetProjectMemberInternalServerError{}
}

/*
GetProjectMemberInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetProjectMemberInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get project member internal server error response has a 2xx status code
func (o *GetProjectMemberInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get project member internal server error response has a 3xx status code
func (o *GetProjectMemberInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get project member internal server error response has a 4xx status code
func (o *GetProjectMemberInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get project member internal server error response has a 5xx status code
func (o *GetProjectMemberInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get project member internal server error response a status code equal to that given
func (o *GetProjectMemberInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get project member internal server error response
func (o *GetProjectMemberInternalServerError) Code() int {
	return 500
}

func (o *GetProjectMemberInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectMemberInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/members/{mid}][%d] getProjectMemberInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProjectMemberInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetProjectMemberInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
