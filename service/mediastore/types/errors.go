// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	"fmt"
	smithy "github.com/aws/smithy-go"
)

// The container that you specified in the request already exists or is being
// updated.
type ContainerInUseException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ContainerInUseException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContainerInUseException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContainerInUseException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ContainerInUseException"
	}
	return *e.ErrorCodeOverride
}
func (e *ContainerInUseException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The container that you specified in the request does not exist.
type ContainerNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *ContainerNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *ContainerNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *ContainerNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "ContainerNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *ContainerNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The CORS policy that you specified in the request does not exist.
type CorsPolicyNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *CorsPolicyNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *CorsPolicyNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *CorsPolicyNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "CorsPolicyNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *CorsPolicyNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The service is temporarily unavailable.
type InternalServerError struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *InternalServerError) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *InternalServerError) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "InternalServerError"
	}
	return *e.ErrorCodeOverride
}
func (e *InternalServerError) ErrorFault() smithy.ErrorFault { return smithy.FaultServer }

// A service limit has been exceeded.
type LimitExceededException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *LimitExceededException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *LimitExceededException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *LimitExceededException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "LimitExceededException"
	}
	return *e.ErrorCodeOverride
}
func (e *LimitExceededException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }

// The policy that you specified in the request does not exist.
type PolicyNotFoundException struct {
	Message *string

	ErrorCodeOverride *string

	noSmithyDocumentSerde
}

func (e *PolicyNotFoundException) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode(), e.ErrorMessage())
}
func (e *PolicyNotFoundException) ErrorMessage() string {
	if e.Message == nil {
		return ""
	}
	return *e.Message
}
func (e *PolicyNotFoundException) ErrorCode() string {
	if e.ErrorCodeOverride == nil {
		return "PolicyNotFoundException"
	}
	return *e.ErrorCodeOverride
}
func (e *PolicyNotFoundException) ErrorFault() smithy.ErrorFault { return smithy.FaultClient }
