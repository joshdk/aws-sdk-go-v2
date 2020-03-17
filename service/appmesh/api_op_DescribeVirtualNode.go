// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeVirtualNodeInput struct {
	_ struct{} `type:"structure"`

	// MeshName is a required field
	MeshName *string `location:"uri" locationName:"meshName" min:"1" type:"string" required:"true"`

	MeshOwner *string `location:"querystring" locationName:"meshOwner" min:"12" type:"string"`

	// VirtualNodeName is a required field
	VirtualNodeName *string `location:"uri" locationName:"virtualNodeName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeVirtualNodeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeVirtualNodeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeVirtualNodeInput"}

	if s.MeshName == nil {
		invalidParams.Add(aws.NewErrParamRequired("MeshName"))
	}
	if s.MeshName != nil && len(*s.MeshName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshName", 1))
	}
	if s.MeshOwner != nil && len(*s.MeshOwner) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("MeshOwner", 12))
	}

	if s.VirtualNodeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualNodeName"))
	}
	if s.VirtualNodeName != nil && len(*s.VirtualNodeName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VirtualNodeName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeVirtualNodeInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.MeshName != nil {
		v := *s.MeshName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "meshName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VirtualNodeName != nil {
		v := *s.VirtualNodeName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "virtualNodeName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MeshOwner != nil {
		v := *s.MeshOwner

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "meshOwner", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type DescribeVirtualNodeOutput struct {
	_ struct{} `type:"structure" payload:"VirtualNode"`

	// An object that represents a virtual node returned by a describe operation.
	//
	// VirtualNode is a required field
	VirtualNode *VirtualNodeData `locationName:"virtualNode" type:"structure" required:"true"`
}

// String returns the string representation
func (s DescribeVirtualNodeOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeVirtualNodeOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.VirtualNode != nil {
		v := s.VirtualNode

		metadata := protocol.Metadata{}
		e.SetFields(protocol.PayloadTarget, "virtualNode", v, metadata)
	}
	return nil
}

const opDescribeVirtualNode = "DescribeVirtualNode"

// DescribeVirtualNodeRequest returns a request value for making API operation for
// AWS App Mesh.
//
// Describes an existing virtual node.
//
//    // Example sending a request using DescribeVirtualNodeRequest.
//    req := client.DescribeVirtualNodeRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/appmesh-2019-01-25/DescribeVirtualNode
func (c *Client) DescribeVirtualNodeRequest(input *DescribeVirtualNodeInput) DescribeVirtualNodeRequest {
	op := &aws.Operation{
		Name:       opDescribeVirtualNode,
		HTTPMethod: "GET",
		HTTPPath:   "/v20190125/meshes/{meshName}/virtualNodes/{virtualNodeName}",
	}

	if input == nil {
		input = &DescribeVirtualNodeInput{}
	}

	req := c.newRequest(op, input, &DescribeVirtualNodeOutput{})
	return DescribeVirtualNodeRequest{Request: req, Input: input, Copy: c.DescribeVirtualNodeRequest}
}

// DescribeVirtualNodeRequest is the request type for the
// DescribeVirtualNode API operation.
type DescribeVirtualNodeRequest struct {
	*aws.Request
	Input *DescribeVirtualNodeInput
	Copy  func(*DescribeVirtualNodeInput) DescribeVirtualNodeRequest
}

// Send marshals and sends the DescribeVirtualNode API request.
func (r DescribeVirtualNodeRequest) Send(ctx context.Context) (*DescribeVirtualNodeResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeVirtualNodeResponse{
		DescribeVirtualNodeOutput: r.Request.Data.(*DescribeVirtualNodeOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeVirtualNodeResponse is the response type for the
// DescribeVirtualNode API operation.
type DescribeVirtualNodeResponse struct {
	*DescribeVirtualNodeOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeVirtualNode request.
func (r *DescribeVirtualNodeResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
