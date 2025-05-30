// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// GetVolumesReader is a Reader for the GetVolumes structure.
type GetVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetVolumesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetVolumesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetVolumesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetVolumesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /systeminfo/volumes] getVolumes", response, response.Code())
	}
}

// NewGetVolumesOK creates a GetVolumesOK with default headers values
func NewGetVolumesOK() *GetVolumesOK {
	return &GetVolumesOK{}
}

/*
GetVolumesOK describes a response with status code 200, with default header values.

Get system volumes successfully.
*/
type GetVolumesOK struct {
	Payload *models.SystemInfo
}

// IsSuccess returns true when this get volumes o k response has a 2xx status code
func (o *GetVolumesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get volumes o k response has a 3xx status code
func (o *GetVolumesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get volumes o k response has a 4xx status code
func (o *GetVolumesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get volumes o k response has a 5xx status code
func (o *GetVolumesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get volumes o k response a status code equal to that given
func (o *GetVolumesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get volumes o k response
func (o *GetVolumesOK) Code() int {
	return 200
}

func (o *GetVolumesOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesOK  %+v", 200, o.Payload)
}

func (o *GetVolumesOK) String() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesOK  %+v", 200, o.Payload)
}

func (o *GetVolumesOK) GetPayload() *models.SystemInfo {
	return o.Payload
}

func (o *GetVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SystemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetVolumesUnauthorized creates a GetVolumesUnauthorized with default headers values
func NewGetVolumesUnauthorized() *GetVolumesUnauthorized {
	return &GetVolumesUnauthorized{}
}

/*
GetVolumesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetVolumesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get volumes unauthorized response has a 2xx status code
func (o *GetVolumesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get volumes unauthorized response has a 3xx status code
func (o *GetVolumesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get volumes unauthorized response has a 4xx status code
func (o *GetVolumesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get volumes unauthorized response has a 5xx status code
func (o *GetVolumesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get volumes unauthorized response a status code equal to that given
func (o *GetVolumesUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get volumes unauthorized response
func (o *GetVolumesUnauthorized) Code() int {
	return 401
}

func (o *GetVolumesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVolumesUnauthorized) String() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetVolumesUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetVolumesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVolumesForbidden creates a GetVolumesForbidden with default headers values
func NewGetVolumesForbidden() *GetVolumesForbidden {
	return &GetVolumesForbidden{}
}

/*
GetVolumesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetVolumesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get volumes forbidden response has a 2xx status code
func (o *GetVolumesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get volumes forbidden response has a 3xx status code
func (o *GetVolumesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get volumes forbidden response has a 4xx status code
func (o *GetVolumesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get volumes forbidden response has a 5xx status code
func (o *GetVolumesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get volumes forbidden response a status code equal to that given
func (o *GetVolumesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get volumes forbidden response
func (o *GetVolumesForbidden) Code() int {
	return 403
}

func (o *GetVolumesForbidden) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesForbidden  %+v", 403, o.Payload)
}

func (o *GetVolumesForbidden) String() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesForbidden  %+v", 403, o.Payload)
}

func (o *GetVolumesForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetVolumesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVolumesNotFound creates a GetVolumesNotFound with default headers values
func NewGetVolumesNotFound() *GetVolumesNotFound {
	return &GetVolumesNotFound{}
}

/*
GetVolumesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetVolumesNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get volumes not found response has a 2xx status code
func (o *GetVolumesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get volumes not found response has a 3xx status code
func (o *GetVolumesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get volumes not found response has a 4xx status code
func (o *GetVolumesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get volumes not found response has a 5xx status code
func (o *GetVolumesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get volumes not found response a status code equal to that given
func (o *GetVolumesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get volumes not found response
func (o *GetVolumesNotFound) Code() int {
	return 404
}

func (o *GetVolumesNotFound) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesNotFound  %+v", 404, o.Payload)
}

func (o *GetVolumesNotFound) String() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesNotFound  %+v", 404, o.Payload)
}

func (o *GetVolumesNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetVolumesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetVolumesInternalServerError creates a GetVolumesInternalServerError with default headers values
func NewGetVolumesInternalServerError() *GetVolumesInternalServerError {
	return &GetVolumesInternalServerError{}
}

/*
GetVolumesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetVolumesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get volumes internal server error response has a 2xx status code
func (o *GetVolumesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get volumes internal server error response has a 3xx status code
func (o *GetVolumesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get volumes internal server error response has a 4xx status code
func (o *GetVolumesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get volumes internal server error response has a 5xx status code
func (o *GetVolumesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get volumes internal server error response a status code equal to that given
func (o *GetVolumesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get volumes internal server error response
func (o *GetVolumesInternalServerError) Code() int {
	return 500
}

func (o *GetVolumesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVolumesInternalServerError) String() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getVolumesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetVolumesInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetVolumesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
