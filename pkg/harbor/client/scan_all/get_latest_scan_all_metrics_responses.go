// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// GetLatestScanAllMetricsReader is a Reader for the GetLatestScanAllMetrics structure.
type GetLatestScanAllMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestScanAllMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestScanAllMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLatestScanAllMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestScanAllMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetLatestScanAllMetricsPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLatestScanAllMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /scans/all/metrics] getLatestScanAllMetrics", response, response.Code())
	}
}

// NewGetLatestScanAllMetricsOK creates a GetLatestScanAllMetricsOK with default headers values
func NewGetLatestScanAllMetricsOK() *GetLatestScanAllMetricsOK {
	return &GetLatestScanAllMetricsOK{}
}

/*
GetLatestScanAllMetricsOK describes a response with status code 200, with default header values.

OK
*/
type GetLatestScanAllMetricsOK struct {
	Payload *models.Stats
}

// IsSuccess returns true when this get latest scan all metrics o k response has a 2xx status code
func (o *GetLatestScanAllMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latest scan all metrics o k response has a 3xx status code
func (o *GetLatestScanAllMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scan all metrics o k response has a 4xx status code
func (o *GetLatestScanAllMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest scan all metrics o k response has a 5xx status code
func (o *GetLatestScanAllMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scan all metrics o k response a status code equal to that given
func (o *GetLatestScanAllMetricsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get latest scan all metrics o k response
func (o *GetLatestScanAllMetricsOK) Code() int {
	return 200
}

func (o *GetLatestScanAllMetricsOK) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsOK  %+v", 200, o.Payload)
}

func (o *GetLatestScanAllMetricsOK) String() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsOK  %+v", 200, o.Payload)
}

func (o *GetLatestScanAllMetricsOK) GetPayload() *models.Stats {
	return o.Payload
}

func (o *GetLatestScanAllMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Stats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScanAllMetricsUnauthorized creates a GetLatestScanAllMetricsUnauthorized with default headers values
func NewGetLatestScanAllMetricsUnauthorized() *GetLatestScanAllMetricsUnauthorized {
	return &GetLatestScanAllMetricsUnauthorized{}
}

/*
GetLatestScanAllMetricsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLatestScanAllMetricsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scan all metrics unauthorized response has a 2xx status code
func (o *GetLatestScanAllMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scan all metrics unauthorized response has a 3xx status code
func (o *GetLatestScanAllMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scan all metrics unauthorized response has a 4xx status code
func (o *GetLatestScanAllMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest scan all metrics unauthorized response has a 5xx status code
func (o *GetLatestScanAllMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scan all metrics unauthorized response a status code equal to that given
func (o *GetLatestScanAllMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get latest scan all metrics unauthorized response
func (o *GetLatestScanAllMetricsUnauthorized) Code() int {
	return 401
}

func (o *GetLatestScanAllMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestScanAllMetricsUnauthorized) String() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestScanAllMetricsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLatestScanAllMetricsForbidden creates a GetLatestScanAllMetricsForbidden with default headers values
func NewGetLatestScanAllMetricsForbidden() *GetLatestScanAllMetricsForbidden {
	return &GetLatestScanAllMetricsForbidden{}
}

/*
GetLatestScanAllMetricsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLatestScanAllMetricsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scan all metrics forbidden response has a 2xx status code
func (o *GetLatestScanAllMetricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scan all metrics forbidden response has a 3xx status code
func (o *GetLatestScanAllMetricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scan all metrics forbidden response has a 4xx status code
func (o *GetLatestScanAllMetricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest scan all metrics forbidden response has a 5xx status code
func (o *GetLatestScanAllMetricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scan all metrics forbidden response a status code equal to that given
func (o *GetLatestScanAllMetricsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get latest scan all metrics forbidden response
func (o *GetLatestScanAllMetricsForbidden) Code() int {
	return 403
}

func (o *GetLatestScanAllMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestScanAllMetricsForbidden) String() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestScanAllMetricsForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLatestScanAllMetricsPreconditionFailed creates a GetLatestScanAllMetricsPreconditionFailed with default headers values
func NewGetLatestScanAllMetricsPreconditionFailed() *GetLatestScanAllMetricsPreconditionFailed {
	return &GetLatestScanAllMetricsPreconditionFailed{}
}

/*
GetLatestScanAllMetricsPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type GetLatestScanAllMetricsPreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scan all metrics precondition failed response has a 2xx status code
func (o *GetLatestScanAllMetricsPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scan all metrics precondition failed response has a 3xx status code
func (o *GetLatestScanAllMetricsPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scan all metrics precondition failed response has a 4xx status code
func (o *GetLatestScanAllMetricsPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest scan all metrics precondition failed response has a 5xx status code
func (o *GetLatestScanAllMetricsPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scan all metrics precondition failed response a status code equal to that given
func (o *GetLatestScanAllMetricsPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

// Code gets the status code for the get latest scan all metrics precondition failed response
func (o *GetLatestScanAllMetricsPreconditionFailed) Code() int {
	return 412
}

func (o *GetLatestScanAllMetricsPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetLatestScanAllMetricsPreconditionFailed) String() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetLatestScanAllMetricsPreconditionFailed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetLatestScanAllMetricsInternalServerError creates a GetLatestScanAllMetricsInternalServerError with default headers values
func NewGetLatestScanAllMetricsInternalServerError() *GetLatestScanAllMetricsInternalServerError {
	return &GetLatestScanAllMetricsInternalServerError{}
}

/*
GetLatestScanAllMetricsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLatestScanAllMetricsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scan all metrics internal server error response has a 2xx status code
func (o *GetLatestScanAllMetricsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scan all metrics internal server error response has a 3xx status code
func (o *GetLatestScanAllMetricsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scan all metrics internal server error response has a 4xx status code
func (o *GetLatestScanAllMetricsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest scan all metrics internal server error response has a 5xx status code
func (o *GetLatestScanAllMetricsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get latest scan all metrics internal server error response a status code equal to that given
func (o *GetLatestScanAllMetricsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get latest scan all metrics internal server error response
func (o *GetLatestScanAllMetricsInternalServerError) Code() int {
	return 500
}

func (o *GetLatestScanAllMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestScanAllMetricsInternalServerError) String() string {
	return fmt.Sprintf("[GET /scans/all/metrics][%d] getLatestScanAllMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestScanAllMetricsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScanAllMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
