// Code generated by go-swagger; DO NOT EDIT.

package robotv1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// DeleteRobotV1Reader is a Reader for the DeleteRobotV1 structure.
type DeleteRobotV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRobotV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRobotV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRobotV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRobotV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRobotV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRobotV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRobotV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /projects/{project_name_or_id}/robots/{robot_id}] DeleteRobotV1", response, response.Code())
	}
}

// NewDeleteRobotV1OK creates a DeleteRobotV1OK with default headers values
func NewDeleteRobotV1OK() *DeleteRobotV1OK {
	return &DeleteRobotV1OK{}
}

/*
DeleteRobotV1OK describes a response with status code 200, with default header values.

Success
*/
type DeleteRobotV1OK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this delete robot v1 o k response has a 2xx status code
func (o *DeleteRobotV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete robot v1 o k response has a 3xx status code
func (o *DeleteRobotV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot v1 o k response has a 4xx status code
func (o *DeleteRobotV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete robot v1 o k response has a 5xx status code
func (o *DeleteRobotV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot v1 o k response a status code equal to that given
func (o *DeleteRobotV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete robot v1 o k response
func (o *DeleteRobotV1OK) Code() int {
	return 200
}

func (o *DeleteRobotV1OK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1OK ", 200)
}

func (o *DeleteRobotV1OK) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1OK ", 200)
}

func (o *DeleteRobotV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteRobotV1BadRequest creates a DeleteRobotV1BadRequest with default headers values
func NewDeleteRobotV1BadRequest() *DeleteRobotV1BadRequest {
	return &DeleteRobotV1BadRequest{}
}

/*
DeleteRobotV1BadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteRobotV1BadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot v1 bad request response has a 2xx status code
func (o *DeleteRobotV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot v1 bad request response has a 3xx status code
func (o *DeleteRobotV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot v1 bad request response has a 4xx status code
func (o *DeleteRobotV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot v1 bad request response has a 5xx status code
func (o *DeleteRobotV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot v1 bad request response a status code equal to that given
func (o *DeleteRobotV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete robot v1 bad request response
func (o *DeleteRobotV1BadRequest) Code() int {
	return 400
}

func (o *DeleteRobotV1BadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRobotV1BadRequest) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1BadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRobotV1BadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotV1Unauthorized creates a DeleteRobotV1Unauthorized with default headers values
func NewDeleteRobotV1Unauthorized() *DeleteRobotV1Unauthorized {
	return &DeleteRobotV1Unauthorized{}
}

/*
DeleteRobotV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRobotV1Unauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot v1 unauthorized response has a 2xx status code
func (o *DeleteRobotV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot v1 unauthorized response has a 3xx status code
func (o *DeleteRobotV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot v1 unauthorized response has a 4xx status code
func (o *DeleteRobotV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot v1 unauthorized response has a 5xx status code
func (o *DeleteRobotV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot v1 unauthorized response a status code equal to that given
func (o *DeleteRobotV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete robot v1 unauthorized response
func (o *DeleteRobotV1Unauthorized) Code() int {
	return 401
}

func (o *DeleteRobotV1Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRobotV1Unauthorized) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRobotV1Unauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotV1Forbidden creates a DeleteRobotV1Forbidden with default headers values
func NewDeleteRobotV1Forbidden() *DeleteRobotV1Forbidden {
	return &DeleteRobotV1Forbidden{}
}

/*
DeleteRobotV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRobotV1Forbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot v1 forbidden response has a 2xx status code
func (o *DeleteRobotV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot v1 forbidden response has a 3xx status code
func (o *DeleteRobotV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot v1 forbidden response has a 4xx status code
func (o *DeleteRobotV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot v1 forbidden response has a 5xx status code
func (o *DeleteRobotV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot v1 forbidden response a status code equal to that given
func (o *DeleteRobotV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete robot v1 forbidden response
func (o *DeleteRobotV1Forbidden) Code() int {
	return 403
}

func (o *DeleteRobotV1Forbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRobotV1Forbidden) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1Forbidden  %+v", 403, o.Payload)
}

func (o *DeleteRobotV1Forbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotV1NotFound creates a DeleteRobotV1NotFound with default headers values
func NewDeleteRobotV1NotFound() *DeleteRobotV1NotFound {
	return &DeleteRobotV1NotFound{}
}

/*
DeleteRobotV1NotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteRobotV1NotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot v1 not found response has a 2xx status code
func (o *DeleteRobotV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot v1 not found response has a 3xx status code
func (o *DeleteRobotV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot v1 not found response has a 4xx status code
func (o *DeleteRobotV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete robot v1 not found response has a 5xx status code
func (o *DeleteRobotV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete robot v1 not found response a status code equal to that given
func (o *DeleteRobotV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete robot v1 not found response
func (o *DeleteRobotV1NotFound) Code() int {
	return 404
}

func (o *DeleteRobotV1NotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1NotFound  %+v", 404, o.Payload)
}

func (o *DeleteRobotV1NotFound) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1NotFound  %+v", 404, o.Payload)
}

func (o *DeleteRobotV1NotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotV1InternalServerError creates a DeleteRobotV1InternalServerError with default headers values
func NewDeleteRobotV1InternalServerError() *DeleteRobotV1InternalServerError {
	return &DeleteRobotV1InternalServerError{}
}

/*
DeleteRobotV1InternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteRobotV1InternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this delete robot v1 internal server error response has a 2xx status code
func (o *DeleteRobotV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete robot v1 internal server error response has a 3xx status code
func (o *DeleteRobotV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete robot v1 internal server error response has a 4xx status code
func (o *DeleteRobotV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete robot v1 internal server error response has a 5xx status code
func (o *DeleteRobotV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete robot v1 internal server error response a status code equal to that given
func (o *DeleteRobotV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete robot v1 internal server error response
func (o *DeleteRobotV1InternalServerError) Code() int {
	return 500
}

func (o *DeleteRobotV1InternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRobotV1InternalServerError) String() string {
	return fmt.Sprintf("[DELETE /projects/{project_name_or_id}/robots/{robot_id}][%d] deleteRobotV1InternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRobotV1InternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *DeleteRobotV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
