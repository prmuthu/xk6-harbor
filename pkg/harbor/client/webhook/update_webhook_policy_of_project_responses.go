// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/prmuthu/xk6-harbor/pkg/harbor/models"
)

// UpdateWebhookPolicyOfProjectReader is a Reader for the UpdateWebhookPolicyOfProject structure.
type UpdateWebhookPolicyOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateWebhookPolicyOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateWebhookPolicyOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateWebhookPolicyOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateWebhookPolicyOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateWebhookPolicyOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateWebhookPolicyOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateWebhookPolicyOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}] UpdateWebhookPolicyOfProject", response, response.Code())
	}
}

// NewUpdateWebhookPolicyOfProjectOK creates a UpdateWebhookPolicyOfProjectOK with default headers values
func NewUpdateWebhookPolicyOfProjectOK() *UpdateWebhookPolicyOfProjectOK {
	return &UpdateWebhookPolicyOfProjectOK{}
}

/*
UpdateWebhookPolicyOfProjectOK describes a response with status code 200, with default header values.

Success
*/
type UpdateWebhookPolicyOfProjectOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this update webhook policy of project o k response has a 2xx status code
func (o *UpdateWebhookPolicyOfProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update webhook policy of project o k response has a 3xx status code
func (o *UpdateWebhookPolicyOfProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update webhook policy of project o k response has a 4xx status code
func (o *UpdateWebhookPolicyOfProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update webhook policy of project o k response has a 5xx status code
func (o *UpdateWebhookPolicyOfProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update webhook policy of project o k response a status code equal to that given
func (o *UpdateWebhookPolicyOfProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update webhook policy of project o k response
func (o *UpdateWebhookPolicyOfProjectOK) Code() int {
	return 200
}

func (o *UpdateWebhookPolicyOfProjectOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectOK ", 200)
}

func (o *UpdateWebhookPolicyOfProjectOK) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectOK ", 200)
}

func (o *UpdateWebhookPolicyOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateWebhookPolicyOfProjectBadRequest creates a UpdateWebhookPolicyOfProjectBadRequest with default headers values
func NewUpdateWebhookPolicyOfProjectBadRequest() *UpdateWebhookPolicyOfProjectBadRequest {
	return &UpdateWebhookPolicyOfProjectBadRequest{}
}

/*
UpdateWebhookPolicyOfProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateWebhookPolicyOfProjectBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update webhook policy of project bad request response has a 2xx status code
func (o *UpdateWebhookPolicyOfProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update webhook policy of project bad request response has a 3xx status code
func (o *UpdateWebhookPolicyOfProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update webhook policy of project bad request response has a 4xx status code
func (o *UpdateWebhookPolicyOfProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update webhook policy of project bad request response has a 5xx status code
func (o *UpdateWebhookPolicyOfProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update webhook policy of project bad request response a status code equal to that given
func (o *UpdateWebhookPolicyOfProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the update webhook policy of project bad request response
func (o *UpdateWebhookPolicyOfProjectBadRequest) Code() int {
	return 400
}

func (o *UpdateWebhookPolicyOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectBadRequest) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateWebhookPolicyOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateWebhookPolicyOfProjectUnauthorized creates a UpdateWebhookPolicyOfProjectUnauthorized with default headers values
func NewUpdateWebhookPolicyOfProjectUnauthorized() *UpdateWebhookPolicyOfProjectUnauthorized {
	return &UpdateWebhookPolicyOfProjectUnauthorized{}
}

/*
UpdateWebhookPolicyOfProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateWebhookPolicyOfProjectUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update webhook policy of project unauthorized response has a 2xx status code
func (o *UpdateWebhookPolicyOfProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update webhook policy of project unauthorized response has a 3xx status code
func (o *UpdateWebhookPolicyOfProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update webhook policy of project unauthorized response has a 4xx status code
func (o *UpdateWebhookPolicyOfProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update webhook policy of project unauthorized response has a 5xx status code
func (o *UpdateWebhookPolicyOfProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update webhook policy of project unauthorized response a status code equal to that given
func (o *UpdateWebhookPolicyOfProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the update webhook policy of project unauthorized response
func (o *UpdateWebhookPolicyOfProjectUnauthorized) Code() int {
	return 401
}

func (o *UpdateWebhookPolicyOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectUnauthorized) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateWebhookPolicyOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateWebhookPolicyOfProjectForbidden creates a UpdateWebhookPolicyOfProjectForbidden with default headers values
func NewUpdateWebhookPolicyOfProjectForbidden() *UpdateWebhookPolicyOfProjectForbidden {
	return &UpdateWebhookPolicyOfProjectForbidden{}
}

/*
UpdateWebhookPolicyOfProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateWebhookPolicyOfProjectForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update webhook policy of project forbidden response has a 2xx status code
func (o *UpdateWebhookPolicyOfProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update webhook policy of project forbidden response has a 3xx status code
func (o *UpdateWebhookPolicyOfProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update webhook policy of project forbidden response has a 4xx status code
func (o *UpdateWebhookPolicyOfProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update webhook policy of project forbidden response has a 5xx status code
func (o *UpdateWebhookPolicyOfProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update webhook policy of project forbidden response a status code equal to that given
func (o *UpdateWebhookPolicyOfProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update webhook policy of project forbidden response
func (o *UpdateWebhookPolicyOfProjectForbidden) Code() int {
	return 403
}

func (o *UpdateWebhookPolicyOfProjectForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectForbidden) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateWebhookPolicyOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateWebhookPolicyOfProjectNotFound creates a UpdateWebhookPolicyOfProjectNotFound with default headers values
func NewUpdateWebhookPolicyOfProjectNotFound() *UpdateWebhookPolicyOfProjectNotFound {
	return &UpdateWebhookPolicyOfProjectNotFound{}
}

/*
UpdateWebhookPolicyOfProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateWebhookPolicyOfProjectNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update webhook policy of project not found response has a 2xx status code
func (o *UpdateWebhookPolicyOfProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update webhook policy of project not found response has a 3xx status code
func (o *UpdateWebhookPolicyOfProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update webhook policy of project not found response has a 4xx status code
func (o *UpdateWebhookPolicyOfProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update webhook policy of project not found response has a 5xx status code
func (o *UpdateWebhookPolicyOfProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update webhook policy of project not found response a status code equal to that given
func (o *UpdateWebhookPolicyOfProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update webhook policy of project not found response
func (o *UpdateWebhookPolicyOfProjectNotFound) Code() int {
	return 404
}

func (o *UpdateWebhookPolicyOfProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectNotFound) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateWebhookPolicyOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateWebhookPolicyOfProjectInternalServerError creates a UpdateWebhookPolicyOfProjectInternalServerError with default headers values
func NewUpdateWebhookPolicyOfProjectInternalServerError() *UpdateWebhookPolicyOfProjectInternalServerError {
	return &UpdateWebhookPolicyOfProjectInternalServerError{}
}

/*
UpdateWebhookPolicyOfProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateWebhookPolicyOfProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update webhook policy of project internal server error response has a 2xx status code
func (o *UpdateWebhookPolicyOfProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update webhook policy of project internal server error response has a 3xx status code
func (o *UpdateWebhookPolicyOfProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update webhook policy of project internal server error response has a 4xx status code
func (o *UpdateWebhookPolicyOfProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update webhook policy of project internal server error response has a 5xx status code
func (o *UpdateWebhookPolicyOfProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update webhook policy of project internal server error response a status code equal to that given
func (o *UpdateWebhookPolicyOfProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update webhook policy of project internal server error response
func (o *UpdateWebhookPolicyOfProjectInternalServerError) Code() int {
	return 500
}

func (o *UpdateWebhookPolicyOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectInternalServerError) String() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] updateWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateWebhookPolicyOfProjectInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateWebhookPolicyOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
