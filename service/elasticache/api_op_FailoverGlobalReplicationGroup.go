// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticache

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type FailoverGlobalReplicationGroupInput struct {
	_ struct{} `type:"structure"`

	// The name of the Global Datastore
	//
	// GlobalReplicationGroupId is a required field
	GlobalReplicationGroupId *string `type:"string" required:"true"`

	// The AWS region of the primary cluster of the Global Datastore
	//
	// PrimaryRegion is a required field
	PrimaryRegion *string `type:"string" required:"true"`

	// The name of the primary replication group
	//
	// PrimaryReplicationGroupId is a required field
	PrimaryReplicationGroupId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s FailoverGlobalReplicationGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *FailoverGlobalReplicationGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "FailoverGlobalReplicationGroupInput"}

	if s.GlobalReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GlobalReplicationGroupId"))
	}

	if s.PrimaryRegion == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrimaryRegion"))
	}

	if s.PrimaryReplicationGroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PrimaryReplicationGroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type FailoverGlobalReplicationGroupOutput struct {
	_ struct{} `type:"structure"`

	// Consists of a primary cluster that accepts writes and an associated secondary
	// cluster that resides in a different AWS region. The secondary cluster accepts
	// only reads. The primary cluster automatically replicates updates to the secondary
	// cluster.
	//
	//    * The GlobalReplicationGroupId represents the name of the Global Datastore,
	//    which is what you use to associate a secondary cluster.
	GlobalReplicationGroup *GlobalReplicationGroup `type:"structure"`
}

// String returns the string representation
func (s FailoverGlobalReplicationGroupOutput) String() string {
	return awsutil.Prettify(s)
}

const opFailoverGlobalReplicationGroup = "FailoverGlobalReplicationGroup"

// FailoverGlobalReplicationGroupRequest returns a request value for making API operation for
// Amazon ElastiCache.
//
// Used to failover the primary region to a selected secondary region.
//
//    // Example sending a request using FailoverGlobalReplicationGroupRequest.
//    req := client.FailoverGlobalReplicationGroupRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticache-2015-02-02/FailoverGlobalReplicationGroup
func (c *Client) FailoverGlobalReplicationGroupRequest(input *FailoverGlobalReplicationGroupInput) FailoverGlobalReplicationGroupRequest {
	op := &aws.Operation{
		Name:       opFailoverGlobalReplicationGroup,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &FailoverGlobalReplicationGroupInput{}
	}

	req := c.newRequest(op, input, &FailoverGlobalReplicationGroupOutput{})
	return FailoverGlobalReplicationGroupRequest{Request: req, Input: input, Copy: c.FailoverGlobalReplicationGroupRequest}
}

// FailoverGlobalReplicationGroupRequest is the request type for the
// FailoverGlobalReplicationGroup API operation.
type FailoverGlobalReplicationGroupRequest struct {
	*aws.Request
	Input *FailoverGlobalReplicationGroupInput
	Copy  func(*FailoverGlobalReplicationGroupInput) FailoverGlobalReplicationGroupRequest
}

// Send marshals and sends the FailoverGlobalReplicationGroup API request.
func (r FailoverGlobalReplicationGroupRequest) Send(ctx context.Context) (*FailoverGlobalReplicationGroupResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &FailoverGlobalReplicationGroupResponse{
		FailoverGlobalReplicationGroupOutput: r.Request.Data.(*FailoverGlobalReplicationGroupOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// FailoverGlobalReplicationGroupResponse is the response type for the
// FailoverGlobalReplicationGroup API operation.
type FailoverGlobalReplicationGroupResponse struct {
	*FailoverGlobalReplicationGroupOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// FailoverGlobalReplicationGroup request.
func (r *FailoverGlobalReplicationGroupResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
