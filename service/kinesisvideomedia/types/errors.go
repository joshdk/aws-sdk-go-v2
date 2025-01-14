// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// Kinesis Video Streams has throttled the request because you have exceeded the
// limit of allowed client calls. Try making the call later.
type ClientLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ClientLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ClientLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ClientLimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ClientLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ClientLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Kinesis Video Streams has throttled the request because you have exceeded the
// limit of allowed client connections.
type ConnectionLimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ConnectionLimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ConnectionLimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ConnectionLimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ConnectionLimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *ConnectionLimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The value for this input parameter is invalid.
type InvalidArgumentException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidArgumentException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidArgumentException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidArgumentException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidArgumentException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidArgumentException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Status Code: 400, Caller used wrong endpoint to write data to a stream. On
// receiving such an exception, the user must call GetDataEndpoint with AccessMode
// set to "READ" and use the endpoint Kinesis Video returns in the next GetMedia
// call.
type InvalidEndpointException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InvalidEndpointException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InvalidEndpointException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InvalidEndpointException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InvalidEndpointException"
	}
	return *e.ErrorCodeOverride
}
func (e *InvalidEndpointException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Status Code: 403, The caller is not authorized to perform an operation on the
// given stream, or the token has expired.
type NotAuthorizedException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *NotAuthorizedException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *NotAuthorizedException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *NotAuthorizedException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "NotAuthorizedException"
	}
	return *e.ErrorCodeOverride
}
func (e *NotAuthorizedException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// Status Code: 404, The stream with the given name does not exist.
type ResourceNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ResourceNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ResourceNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ResourceNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ResourceNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ResourceNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
