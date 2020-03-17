// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type OpenInstancePublicPortsInput struct {
	_ struct{} `type:"structure"`

	// The name of the instance for which you want to open the public ports.
	//
	// InstanceName is a required field
	InstanceName *string `locationName:"instanceName" type:"string" required:"true"`

	// An array of key-value pairs containing information about the port mappings.
	//
	// PortInfo is a required field
	PortInfo *PortInfo `locationName:"portInfo" type:"structure" required:"true"`
}

// String returns the string representation
func (s OpenInstancePublicPortsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *OpenInstancePublicPortsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "OpenInstancePublicPortsInput"}

	if s.InstanceName == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceName"))
	}

	if s.PortInfo == nil {
		invalidParams.Add(aws.NewErrParamRequired("PortInfo"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type OpenInstancePublicPortsOutput struct {
	_ struct{} `type:"structure"`

	// An array of objects that describe the result of the action, such as the status
	// of the request, the time stamp of the request, and the resources affected
	// by the request.
	Operation *Operation `locationName:"operation" type:"structure"`
}

// String returns the string representation
func (s OpenInstancePublicPortsOutput) String() string {
	return awsutil.Prettify(s)
}

const opOpenInstancePublicPorts = "OpenInstancePublicPorts"

// OpenInstancePublicPortsRequest returns a request value for making API operation for
// Amazon Lightsail.
//
// Adds public ports to an Amazon Lightsail instance.
//
// The open instance public ports operation supports tag-based access control
// via resource tags applied to the resource identified by instance name. For
// more information, see the Lightsail Dev Guide (https://lightsail.aws.amazon.com/ls/docs/en/articles/amazon-lightsail-controlling-access-using-tags).
//
//    // Example sending a request using OpenInstancePublicPortsRequest.
//    req := client.OpenInstancePublicPortsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/lightsail-2016-11-28/OpenInstancePublicPorts
func (c *Client) OpenInstancePublicPortsRequest(input *OpenInstancePublicPortsInput) OpenInstancePublicPortsRequest {
	op := &aws.Operation{
		Name:       opOpenInstancePublicPorts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &OpenInstancePublicPortsInput{}
	}

	req := c.newRequest(op, input, &OpenInstancePublicPortsOutput{})
	return OpenInstancePublicPortsRequest{Request: req, Input: input, Copy: c.OpenInstancePublicPortsRequest}
}

// OpenInstancePublicPortsRequest is the request type for the
// OpenInstancePublicPorts API operation.
type OpenInstancePublicPortsRequest struct {
	*aws.Request
	Input *OpenInstancePublicPortsInput
	Copy  func(*OpenInstancePublicPortsInput) OpenInstancePublicPortsRequest
}

// Send marshals and sends the OpenInstancePublicPorts API request.
func (r OpenInstancePublicPortsRequest) Send(ctx context.Context) (*OpenInstancePublicPortsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &OpenInstancePublicPortsResponse{
		OpenInstancePublicPortsOutput: r.Request.Data.(*OpenInstancePublicPortsOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// OpenInstancePublicPortsResponse is the response type for the
// OpenInstancePublicPorts API operation.
type OpenInstancePublicPortsResponse struct {
	*OpenInstancePublicPortsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// OpenInstancePublicPorts request.
func (r *OpenInstancePublicPortsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
