// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ModifyVpnConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the customer gateway at your end of the VPN connection.
	CustomerGatewayId *string `type:"string"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `type:"boolean"`

	// The ID of the transit gateway.
	TransitGatewayId *string `type:"string"`

	// The ID of the VPN connection.
	//
	// VpnConnectionId is a required field
	VpnConnectionId *string `type:"string" required:"true"`

	// The ID of the virtual private gateway at the AWS side of the VPN connection.
	VpnGatewayId *string `type:"string"`
}

// String returns the string representation
func (s ModifyVpnConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyVpnConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ModifyVpnConnectionInput"}

	if s.VpnConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VpnConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ModifyVpnConnectionOutput struct {
	_ struct{} `type:"structure"`

	// Describes a VPN connection.
	VpnConnection *VpnConnection `locationName:"vpnConnection" type:"structure"`
}

// String returns the string representation
func (s ModifyVpnConnectionOutput) String() string {
	return awsutil.Prettify(s)
}

const opModifyVpnConnection = "ModifyVpnConnection"

// ModifyVpnConnectionRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Modifies the target gateway of an AWS Site-to-Site VPN connection. The following
// migration options are available:
//
//    * An existing virtual private gateway to a new virtual private gateway
//
//    * An existing virtual private gateway to a transit gateway
//
//    * An existing transit gateway to a new transit gateway
//
//    * An existing transit gateway to a virtual private gateway
//
// Before you perform the migration to the new gateway, you must configure the
// new gateway. Use CreateVpnGateway to create a virtual private gateway, or
// CreateTransitGateway to create a transit gateway.
//
// This step is required when you migrate from a virtual private gateway with
// static routes to a transit gateway.
//
// You must delete the static routes before you migrate to the new gateway.
//
// Keep a copy of the static route before you delete it. You will need to add
// back these routes to the transit gateway after the VPN connection migration
// is complete.
//
// After you migrate to the new gateway, you might need to modify your VPC route
// table. Use CreateRoute and DeleteRoute to make the changes described in VPN
// Gateway Target Modification Required VPC Route Table Updates (https://docs.aws.amazon.com/vpn/latest/s2svpn/modify-vpn-target.html#step-update-routing)
// in the AWS Site-to-Site VPN User Guide.
//
// When the new gateway is a transit gateway, modify the transit gateway route
// table to allow traffic between the VPC and the AWS Site-to-Site VPN connection.
// Use CreateTransitGatewayRoute to add the routes.
//
// If you deleted VPN static routes, you must add the static routes to the transit
// gateway route table.
//
// After you perform this operation, the AWS VPN endpoint's IP addresses on
// the AWS side and the tunnel options remain intact. Your AWS Site-to-Site
// VPN connection will be temporarily unavailable for a brief period while we
// provision the new endpoints.
//
//    // Example sending a request using ModifyVpnConnectionRequest.
//    req := client.ModifyVpnConnectionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ModifyVpnConnection
func (c *Client) ModifyVpnConnectionRequest(input *ModifyVpnConnectionInput) ModifyVpnConnectionRequest {
	op := &aws.Operation{
		Name:       opModifyVpnConnection,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyVpnConnectionInput{}
	}

	req := c.newRequest(op, input, &ModifyVpnConnectionOutput{})
	return ModifyVpnConnectionRequest{Request: req, Input: input, Copy: c.ModifyVpnConnectionRequest}
}

// ModifyVpnConnectionRequest is the request type for the
// ModifyVpnConnection API operation.
type ModifyVpnConnectionRequest struct {
	*aws.Request
	Input *ModifyVpnConnectionInput
	Copy  func(*ModifyVpnConnectionInput) ModifyVpnConnectionRequest
}

// Send marshals and sends the ModifyVpnConnection API request.
func (r ModifyVpnConnectionRequest) Send(ctx context.Context) (*ModifyVpnConnectionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ModifyVpnConnectionResponse{
		ModifyVpnConnectionOutput: r.Request.Data.(*ModifyVpnConnectionOutput),
		response:                  &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ModifyVpnConnectionResponse is the response type for the
// ModifyVpnConnection API operation.
type ModifyVpnConnectionResponse struct {
	*ModifyVpnConnectionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ModifyVpnConnection request.
func (r *ModifyVpnConnectionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
