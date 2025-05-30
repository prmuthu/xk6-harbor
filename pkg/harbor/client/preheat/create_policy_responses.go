// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// CreatePolicyReader is a Reader for the CreatePolicy structure.
type CreatePolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreatePolicyCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreatePolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreatePolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreatePolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreatePolicyConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreatePolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /projects/{project_name}/preheat/policies] CreatePolicy", response, response.Code())
	}
}

// NewCreatePolicyCreated creates a CreatePolicyCreated with default headers values
func NewCreatePolicyCreated() *CreatePolicyCreated {
	return &CreatePolicyCreated{}
}

/*
CreatePolicyCreated describes a response with status code 201, with default header values.

Created
*/
type CreatePolicyCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this create policy created response has a 2xx status code
func (o *CreatePolicyCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create policy created response has a 3xx status code
func (o *CreatePolicyCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy created response has a 4xx status code
func (o *CreatePolicyCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create policy created response has a 5xx status code
func (o *CreatePolicyCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy created response a status code equal to that given
func (o *CreatePolicyCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create policy created response
func (o *CreatePolicyCreated) Code() int {
	return 201
}

func (o *CreatePolicyCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyCreated ", 201)
}

func (o *CreatePolicyCreated) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyCreated ", 201)
}

func (o *CreatePolicyCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreatePolicyBadRequest creates a CreatePolicyBadRequest with default headers values
func NewCreatePolicyBadRequest() *CreatePolicyBadRequest {
	return &CreatePolicyBadRequest{}
}

/*
CreatePolicyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreatePolicyBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create policy bad request response has a 2xx status code
func (o *CreatePolicyBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy bad request response has a 3xx status code
func (o *CreatePolicyBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy bad request response has a 4xx status code
func (o *CreatePolicyBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy bad request response has a 5xx status code
func (o *CreatePolicyBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy bad request response a status code equal to that given
func (o *CreatePolicyBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create policy bad request response
func (o *CreatePolicyBadRequest) Code() int {
	return 400
}

func (o *CreatePolicyBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePolicyBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyBadRequest  %+v", 400, o.Payload)
}

func (o *CreatePolicyBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreatePolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePolicyUnauthorized creates a CreatePolicyUnauthorized with default headers values
func NewCreatePolicyUnauthorized() *CreatePolicyUnauthorized {
	return &CreatePolicyUnauthorized{}
}

/*
CreatePolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreatePolicyUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create policy unauthorized response has a 2xx status code
func (o *CreatePolicyUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy unauthorized response has a 3xx status code
func (o *CreatePolicyUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy unauthorized response has a 4xx status code
func (o *CreatePolicyUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy unauthorized response has a 5xx status code
func (o *CreatePolicyUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy unauthorized response a status code equal to that given
func (o *CreatePolicyUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the create policy unauthorized response
func (o *CreatePolicyUnauthorized) Code() int {
	return 401
}

func (o *CreatePolicyUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePolicyUnauthorized) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyUnauthorized  %+v", 401, o.Payload)
}

func (o *CreatePolicyUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreatePolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePolicyForbidden creates a CreatePolicyForbidden with default headers values
func NewCreatePolicyForbidden() *CreatePolicyForbidden {
	return &CreatePolicyForbidden{}
}

/*
CreatePolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreatePolicyForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create policy forbidden response has a 2xx status code
func (o *CreatePolicyForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy forbidden response has a 3xx status code
func (o *CreatePolicyForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy forbidden response has a 4xx status code
func (o *CreatePolicyForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy forbidden response has a 5xx status code
func (o *CreatePolicyForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy forbidden response a status code equal to that given
func (o *CreatePolicyForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create policy forbidden response
func (o *CreatePolicyForbidden) Code() int {
	return 403
}

func (o *CreatePolicyForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyForbidden  %+v", 403, o.Payload)
}

func (o *CreatePolicyForbidden) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyForbidden  %+v", 403, o.Payload)
}

func (o *CreatePolicyForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreatePolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePolicyConflict creates a CreatePolicyConflict with default headers values
func NewCreatePolicyConflict() *CreatePolicyConflict {
	return &CreatePolicyConflict{}
}

/*
CreatePolicyConflict describes a response with status code 409, with default header values.

Conflict
*/
type CreatePolicyConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create policy conflict response has a 2xx status code
func (o *CreatePolicyConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy conflict response has a 3xx status code
func (o *CreatePolicyConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy conflict response has a 4xx status code
func (o *CreatePolicyConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this create policy conflict response has a 5xx status code
func (o *CreatePolicyConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this create policy conflict response a status code equal to that given
func (o *CreatePolicyConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the create policy conflict response
func (o *CreatePolicyConflict) Code() int {
	return 409
}

func (o *CreatePolicyConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyConflict  %+v", 409, o.Payload)
}

func (o *CreatePolicyConflict) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyConflict  %+v", 409, o.Payload)
}

func (o *CreatePolicyConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreatePolicyConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreatePolicyInternalServerError creates a CreatePolicyInternalServerError with default headers values
func NewCreatePolicyInternalServerError() *CreatePolicyInternalServerError {
	return &CreatePolicyInternalServerError{}
}

/*
CreatePolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreatePolicyInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create policy internal server error response has a 2xx status code
func (o *CreatePolicyInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create policy internal server error response has a 3xx status code
func (o *CreatePolicyInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create policy internal server error response has a 4xx status code
func (o *CreatePolicyInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create policy internal server error response has a 5xx status code
func (o *CreatePolicyInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create policy internal server error response a status code equal to that given
func (o *CreatePolicyInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create policy internal server error response
func (o *CreatePolicyInternalServerError) Code() int {
	return 500
}

func (o *CreatePolicyInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePolicyInternalServerError) String() string {
	return fmt.Sprintf("[POST /projects/{project_name}/preheat/policies][%d] createPolicyInternalServerError  %+v", 500, o.Payload)
}

func (o *CreatePolicyInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreatePolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
