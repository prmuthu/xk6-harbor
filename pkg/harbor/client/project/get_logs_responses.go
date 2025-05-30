// Code generated by go-swagger; DO NOT EDIT.

package project

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

// GetLogsReader is a Reader for the GetLogs structure.
type GetLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetLogsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetLogsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLogsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name}/logs] getLogs", response, response.Code())
	}
}

// NewGetLogsOK creates a GetLogsOK with default headers values
func NewGetLogsOK() *GetLogsOK {
	return &GetLogsOK{}
}

/*
GetLogsOK describes a response with status code 200, with default header values.

Success
*/
type GetLogsOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of auditlogs
	 */
	XTotalCount int64

	Payload []*models.AuditLog
}

// IsSuccess returns true when this get logs o k response has a 2xx status code
func (o *GetLogsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get logs o k response has a 3xx status code
func (o *GetLogsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs o k response has a 4xx status code
func (o *GetLogsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs o k response has a 5xx status code
func (o *GetLogsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs o k response a status code equal to that given
func (o *GetLogsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get logs o k response
func (o *GetLogsOK) Code() int {
	return 200
}

func (o *GetLogsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsOK  %+v", 200, o.Payload)
}

func (o *GetLogsOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsOK  %+v", 200, o.Payload)
}

func (o *GetLogsOK) GetPayload() []*models.AuditLog {
	return o.Payload
}

func (o *GetLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLogsBadRequest creates a GetLogsBadRequest with default headers values
func NewGetLogsBadRequest() *GetLogsBadRequest {
	return &GetLogsBadRequest{}
}

/*
GetLogsBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetLogsBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs bad request response has a 2xx status code
func (o *GetLogsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs bad request response has a 3xx status code
func (o *GetLogsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs bad request response has a 4xx status code
func (o *GetLogsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs bad request response has a 5xx status code
func (o *GetLogsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs bad request response a status code equal to that given
func (o *GetLogsBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get logs bad request response
func (o *GetLogsBadRequest) Code() int {
	return 400
}

func (o *GetLogsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLogsBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsBadRequest  %+v", 400, o.Payload)
}

func (o *GetLogsBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLogsUnauthorized creates a GetLogsUnauthorized with default headers values
func NewGetLogsUnauthorized() *GetLogsUnauthorized {
	return &GetLogsUnauthorized{}
}

/*
GetLogsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLogsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs unauthorized response has a 2xx status code
func (o *GetLogsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs unauthorized response has a 3xx status code
func (o *GetLogsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs unauthorized response has a 4xx status code
func (o *GetLogsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get logs unauthorized response has a 5xx status code
func (o *GetLogsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get logs unauthorized response a status code equal to that given
func (o *GetLogsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get logs unauthorized response
func (o *GetLogsUnauthorized) Code() int {
	return 401
}

func (o *GetLogsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogsUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLogsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLogsInternalServerError creates a GetLogsInternalServerError with default headers values
func NewGetLogsInternalServerError() *GetLogsInternalServerError {
	return &GetLogsInternalServerError{}
}

/*
GetLogsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLogsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get logs internal server error response has a 2xx status code
func (o *GetLogsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get logs internal server error response has a 3xx status code
func (o *GetLogsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get logs internal server error response has a 4xx status code
func (o *GetLogsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get logs internal server error response has a 5xx status code
func (o *GetLogsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get logs internal server error response a status code equal to that given
func (o *GetLogsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get logs internal server error response
func (o *GetLogsInternalServerError) Code() int {
	return 500
}

func (o *GetLogsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLogsInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name}/logs][%d] getLogsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLogsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLogsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
