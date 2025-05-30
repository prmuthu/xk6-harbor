// Code generated by go-swagger; DO NOT EDIT.

package schedule

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// GetSchedulePausedReader is a Reader for the GetSchedulePaused structure.
type GetSchedulePausedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSchedulePausedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSchedulePausedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSchedulePausedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSchedulePausedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSchedulePausedNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSchedulePausedInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /schedules/{job_type}/paused] getSchedulePaused", response, response.Code())
	}
}

// NewGetSchedulePausedOK creates a GetSchedulePausedOK with default headers values
func NewGetSchedulePausedOK() *GetSchedulePausedOK {
	return &GetSchedulePausedOK{}
}

/*
GetSchedulePausedOK describes a response with status code 200, with default header values.

Get scheduler status successfully.
*/
type GetSchedulePausedOK struct {
	Payload *models.SchedulerStatus
}

// IsSuccess returns true when this get schedule paused o k response has a 2xx status code
func (o *GetSchedulePausedOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get schedule paused o k response has a 3xx status code
func (o *GetSchedulePausedOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule paused o k response has a 4xx status code
func (o *GetSchedulePausedOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule paused o k response has a 5xx status code
func (o *GetSchedulePausedOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule paused o k response a status code equal to that given
func (o *GetSchedulePausedOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get schedule paused o k response
func (o *GetSchedulePausedOK) Code() int {
	return 200
}

func (o *GetSchedulePausedOK) Error() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedOK  %+v", 200, o.Payload)
}

func (o *GetSchedulePausedOK) String() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedOK  %+v", 200, o.Payload)
}

func (o *GetSchedulePausedOK) GetPayload() *models.SchedulerStatus {
	return o.Payload
}

func (o *GetSchedulePausedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SchedulerStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSchedulePausedUnauthorized creates a GetSchedulePausedUnauthorized with default headers values
func NewGetSchedulePausedUnauthorized() *GetSchedulePausedUnauthorized {
	return &GetSchedulePausedUnauthorized{}
}

/*
GetSchedulePausedUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSchedulePausedUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get schedule paused unauthorized response has a 2xx status code
func (o *GetSchedulePausedUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get schedule paused unauthorized response has a 3xx status code
func (o *GetSchedulePausedUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule paused unauthorized response has a 4xx status code
func (o *GetSchedulePausedUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get schedule paused unauthorized response has a 5xx status code
func (o *GetSchedulePausedUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule paused unauthorized response a status code equal to that given
func (o *GetSchedulePausedUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get schedule paused unauthorized response
func (o *GetSchedulePausedUnauthorized) Code() int {
	return 401
}

func (o *GetSchedulePausedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSchedulePausedUnauthorized) String() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSchedulePausedUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSchedulePausedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSchedulePausedForbidden creates a GetSchedulePausedForbidden with default headers values
func NewGetSchedulePausedForbidden() *GetSchedulePausedForbidden {
	return &GetSchedulePausedForbidden{}
}

/*
GetSchedulePausedForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSchedulePausedForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get schedule paused forbidden response has a 2xx status code
func (o *GetSchedulePausedForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get schedule paused forbidden response has a 3xx status code
func (o *GetSchedulePausedForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule paused forbidden response has a 4xx status code
func (o *GetSchedulePausedForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get schedule paused forbidden response has a 5xx status code
func (o *GetSchedulePausedForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule paused forbidden response a status code equal to that given
func (o *GetSchedulePausedForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get schedule paused forbidden response
func (o *GetSchedulePausedForbidden) Code() int {
	return 403
}

func (o *GetSchedulePausedForbidden) Error() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedForbidden  %+v", 403, o.Payload)
}

func (o *GetSchedulePausedForbidden) String() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedForbidden  %+v", 403, o.Payload)
}

func (o *GetSchedulePausedForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSchedulePausedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSchedulePausedNotFound creates a GetSchedulePausedNotFound with default headers values
func NewGetSchedulePausedNotFound() *GetSchedulePausedNotFound {
	return &GetSchedulePausedNotFound{}
}

/*
GetSchedulePausedNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSchedulePausedNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get schedule paused not found response has a 2xx status code
func (o *GetSchedulePausedNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get schedule paused not found response has a 3xx status code
func (o *GetSchedulePausedNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule paused not found response has a 4xx status code
func (o *GetSchedulePausedNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get schedule paused not found response has a 5xx status code
func (o *GetSchedulePausedNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get schedule paused not found response a status code equal to that given
func (o *GetSchedulePausedNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get schedule paused not found response
func (o *GetSchedulePausedNotFound) Code() int {
	return 404
}

func (o *GetSchedulePausedNotFound) Error() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedNotFound  %+v", 404, o.Payload)
}

func (o *GetSchedulePausedNotFound) String() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedNotFound  %+v", 404, o.Payload)
}

func (o *GetSchedulePausedNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSchedulePausedNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSchedulePausedInternalServerError creates a GetSchedulePausedInternalServerError with default headers values
func NewGetSchedulePausedInternalServerError() *GetSchedulePausedInternalServerError {
	return &GetSchedulePausedInternalServerError{}
}

/*
GetSchedulePausedInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSchedulePausedInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get schedule paused internal server error response has a 2xx status code
func (o *GetSchedulePausedInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get schedule paused internal server error response has a 3xx status code
func (o *GetSchedulePausedInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get schedule paused internal server error response has a 4xx status code
func (o *GetSchedulePausedInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get schedule paused internal server error response has a 5xx status code
func (o *GetSchedulePausedInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get schedule paused internal server error response a status code equal to that given
func (o *GetSchedulePausedInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get schedule paused internal server error response
func (o *GetSchedulePausedInternalServerError) Code() int {
	return 500
}

func (o *GetSchedulePausedInternalServerError) Error() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSchedulePausedInternalServerError) String() string {
	return fmt.Sprintf("[GET /schedules/{job_type}/paused][%d] getSchedulePausedInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSchedulePausedInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetSchedulePausedInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
