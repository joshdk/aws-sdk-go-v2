// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package codeguruprofiler

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The structure representing the postAgentProfileRequest.
type PostAgentProfileInput struct {
	_ struct{} `type:"structure" payload:"AgentProfile"`

	// AgentProfile is a required field
	AgentProfile []byte `locationName:"agentProfile" type:"blob" required:"true"`

	// ContentType is a required field
	ContentType *string `location:"header" locationName:"Content-Type" type:"string" required:"true"`

	ProfileToken *string `location:"querystring" locationName:"profileToken" min:"1" type:"string" idempotencyToken:"true"`

	// ProfilingGroupName is a required field
	ProfilingGroupName *string `location:"uri" locationName:"profilingGroupName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PostAgentProfileInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostAgentProfileInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostAgentProfileInput"}

	if s.AgentProfile == nil {
		invalidParams.Add(aws.NewErrParamRequired("AgentProfile"))
	}

	if s.ContentType == nil {
		invalidParams.Add(aws.NewErrParamRequired("ContentType"))
	}
	if s.ProfileToken != nil && len(*s.ProfileToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfileToken", 1))
	}

	if s.ProfilingGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProfilingGroupName"))
	}
	if s.ProfilingGroupName != nil && len(*s.ProfilingGroupName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProfilingGroupName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PostAgentProfileInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.ContentType != nil {
		v := *s.ContentType

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ProfilingGroupName != nil {
		v := *s.ProfilingGroupName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "profilingGroupName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.AgentProfile != nil {
		v := s.AgentProfile

		metadata := protocol.Metadata{}
		e.SetStream(protocol.PayloadTarget, "agentProfile", protocol.BytesStream(v), metadata)
	}
	var ProfileToken string
	if s.ProfileToken != nil {
		ProfileToken = *s.ProfileToken
	} else {
		ProfileToken = protocol.GetIdempotencyToken()
	}
	{
		v := ProfileToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "profileToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The structure representing the postAgentProfileResponse.
type PostAgentProfileOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PostAgentProfileOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PostAgentProfileOutput) MarshalFields(e protocol.FieldEncoder) error {
	return nil
}

const opPostAgentProfile = "PostAgentProfile"

// PostAgentProfileRequest returns a request value for making API operation for
// Amazon CodeGuru Profiler.
//
//    // Example sending a request using PostAgentProfileRequest.
//    req := client.PostAgentProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/codeguruprofiler-2019-07-18/PostAgentProfile
func (c *Client) PostAgentProfileRequest(input *PostAgentProfileInput) PostAgentProfileRequest {
	op := &aws.Operation{
		Name:       opPostAgentProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/profilingGroups/{profilingGroupName}/agentProfile",
	}

	if input == nil {
		input = &PostAgentProfileInput{}
	}

	req := c.newRequest(op, input, &PostAgentProfileOutput{})
	return PostAgentProfileRequest{Request: req, Input: input, Copy: c.PostAgentProfileRequest}
}

// PostAgentProfileRequest is the request type for the
// PostAgentProfile API operation.
type PostAgentProfileRequest struct {
	*aws.Request
	Input *PostAgentProfileInput
	Copy  func(*PostAgentProfileInput) PostAgentProfileRequest
}

// Send marshals and sends the PostAgentProfile API request.
func (r PostAgentProfileRequest) Send(ctx context.Context) (*PostAgentProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PostAgentProfileResponse{
		PostAgentProfileOutput: r.Request.Data.(*PostAgentProfileOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PostAgentProfileResponse is the response type for the
// PostAgentProfile API operation.
type PostAgentProfileResponse struct {
	*PostAgentProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PostAgentProfile request.
func (r *PostAgentProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
