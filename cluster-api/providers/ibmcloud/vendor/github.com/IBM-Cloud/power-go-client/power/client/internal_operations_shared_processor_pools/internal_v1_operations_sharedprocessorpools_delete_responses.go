// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_shared_processor_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// InternalV1OperationsSharedprocessorpoolsDeleteReader is a Reader for the InternalV1OperationsSharedprocessorpoolsDelete structure.
type InternalV1OperationsSharedprocessorpoolsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1OperationsSharedprocessorpoolsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1OperationsSharedprocessorpoolsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}] internal.v1.operations.sharedprocessorpools.delete", response, response.Code())
	}
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteNoContent creates a InternalV1OperationsSharedprocessorpoolsDeleteNoContent with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteNoContent() *InternalV1OperationsSharedprocessorpoolsDeleteNoContent {
	return &InternalV1OperationsSharedprocessorpoolsDeleteNoContent{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteNoContent describes a response with status code 204, with default header values.

Deleted
*/
type InternalV1OperationsSharedprocessorpoolsDeleteNoContent struct {
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete no content response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete no content response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete no content response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete no content response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete no content response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete no content response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) Code() int {
	return 204
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteNoContent", 204)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteNoContent", 204)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteBadRequest creates a InternalV1OperationsSharedprocessorpoolsDeleteBadRequest with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteBadRequest() *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest {
	return &InternalV1OperationsSharedprocessorpoolsDeleteBadRequest{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalV1OperationsSharedprocessorpoolsDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete bad request response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete bad request response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete bad request response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete bad request response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete bad request response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete bad request response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) Code() int {
	return 400
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteUnauthorized creates a InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteUnauthorized() *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized {
	return &InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete unauthorized response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete unauthorized response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete unauthorized response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete unauthorized response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete unauthorized response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete unauthorized response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) Code() int {
	return 401
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteForbidden creates a InternalV1OperationsSharedprocessorpoolsDeleteForbidden with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteForbidden() *InternalV1OperationsSharedprocessorpoolsDeleteForbidden {
	return &InternalV1OperationsSharedprocessorpoolsDeleteForbidden{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1OperationsSharedprocessorpoolsDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete forbidden response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete forbidden response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete forbidden response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete forbidden response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete forbidden response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete forbidden response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) Code() int {
	return 403
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteForbidden %s", 403, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteForbidden %s", 403, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteNotFound creates a InternalV1OperationsSharedprocessorpoolsDeleteNotFound with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteNotFound() *InternalV1OperationsSharedprocessorpoolsDeleteNotFound {
	return &InternalV1OperationsSharedprocessorpoolsDeleteNotFound{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InternalV1OperationsSharedprocessorpoolsDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete not found response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete not found response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete not found response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete not found response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete not found response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete not found response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) Code() int {
	return 404
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteNotFound %s", 404, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteNotFound %s", 404, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteGone creates a InternalV1OperationsSharedprocessorpoolsDeleteGone with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteGone() *InternalV1OperationsSharedprocessorpoolsDeleteGone {
	return &InternalV1OperationsSharedprocessorpoolsDeleteGone{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type InternalV1OperationsSharedprocessorpoolsDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete gone response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete gone response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete gone response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete gone response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete gone response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete gone response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) Code() int {
	return 410
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteGone %s", 410, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteGone %s", 410, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests creates a InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests() *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests {
	return &InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete too many requests response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete too many requests response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete too many requests response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete too many requests response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete too many requests response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete too many requests response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) Code() int {
	return 429
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsDeleteInternalServerError creates a InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError with default headers values
func NewInternalV1OperationsSharedprocessorpoolsDeleteInternalServerError() *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError {
	return &InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError{}
}

/*
InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools delete internal server error response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools delete internal server error response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools delete internal server error response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools delete internal server error response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 operations sharedprocessorpools delete internal server error response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 operations sharedprocessorpools delete internal server error response
func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) Code() int {
	return 500
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/shared-processor-pools/{resource_crn}][%d] internalV1OperationsSharedprocessorpoolsDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
