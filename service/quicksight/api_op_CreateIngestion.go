// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type CreateIngestionInput struct {
	_ struct{} `type:"structure"`

	// The AWS account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID of the dataset used in the ingestion.
	//
	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// An ID for the ingestion.
	//
	// IngestionId is a required field
	IngestionId *string `location:"uri" locationName:"IngestionId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateIngestionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateIngestionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateIngestionInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}

	if s.IngestionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("IngestionId"))
	}
	if s.IngestionId != nil && len(*s.IngestionId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("IngestionId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateIngestionInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AwsAccountId != nil {
		v := *s.AwsAccountId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "AwsAccountId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.DataSetId != nil {
		v := *s.DataSetId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "DataSetId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IngestionId != nil {
		v := *s.IngestionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "IngestionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type CreateIngestionOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) for the data ingestion.
	Arn *string `type:"string"`

	// An ID for the ingestion.
	IngestionId *string `min:"1" type:"string"`

	// The ingestion status.
	IngestionStatus IngestionStatus `type:"string" enum:"true"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The HTTP status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s CreateIngestionOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s CreateIngestionOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.IngestionId != nil {
		v := *s.IngestionId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IngestionId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.IngestionStatus) > 0 {
		v := s.IngestionStatus

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "IngestionStatus", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if s.RequestId != nil {
		v := *s.RequestId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "RequestId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	// ignoring invalid encode state, StatusCode. Status
	return nil
}

const opCreateIngestion = "CreateIngestion"

// CreateIngestionRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Creates and starts a new SPICE ingestion on a dataset
//
// Any ingestions operating on tagged datasets inherit the same tags automatically
// for use in access control. For an example, see How do I create an IAM policy
// to control access to Amazon EC2 resources using tags? (https://aws.example.com/premiumsupport/knowledge-center/iam-ec2-resource-tags/)
// in the AWS Knowledge Center. Tags are visible on the tagged dataset, but
// not on the ingestion resource.
//
//    // Example sending a request using CreateIngestionRequest.
//    req := client.CreateIngestionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateIngestion
func (c *Client) CreateIngestionRequest(input *CreateIngestionInput) CreateIngestionRequest {
	op := &aws.Operation{
		Name:       opCreateIngestion,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/data-sets/{DataSetId}/ingestions/{IngestionId}",
	}

	if input == nil {
		input = &CreateIngestionInput{}
	}

	req := c.newRequest(op, input, &CreateIngestionOutput{})
	return CreateIngestionRequest{Request: req, Input: input, Copy: c.CreateIngestionRequest}
}

// CreateIngestionRequest is the request type for the
// CreateIngestion API operation.
type CreateIngestionRequest struct {
	*aws.Request
	Input *CreateIngestionInput
	Copy  func(*CreateIngestionInput) CreateIngestionRequest
}

// Send marshals and sends the CreateIngestion API request.
func (r CreateIngestionRequest) Send(ctx context.Context) (*CreateIngestionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateIngestionResponse{
		CreateIngestionOutput: r.Request.Data.(*CreateIngestionOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateIngestionResponse is the response type for the
// CreateIngestion API operation.
type CreateIngestionResponse struct {
	*CreateIngestionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateIngestion request.
func (r *CreateIngestionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}