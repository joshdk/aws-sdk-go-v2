// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type GetFolderPathInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// A comma-separated list of values. Specify "NAME" to include the names of
	// the parent folders.
	Fields *string `location:"querystring" locationName:"fields" min:"1" type:"string"`

	// The ID of the folder.
	//
	// FolderId is a required field
	FolderId *string `location:"uri" locationName:"FolderId" min:"1" type:"string" required:"true"`

	// The maximum number of levels in the hierarchy to return.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// This value is not supported.
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`
}

// String returns the string representation
func (s GetFolderPathInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetFolderPathInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetFolderPathInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}
	if s.Fields != nil && len(*s.Fields) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Fields", 1))
	}

	if s.FolderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FolderId"))
	}
	if s.FolderId != nil && len(*s.FolderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FolderId", 1))
	}
	if s.Limit != nil && *s.Limit < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Limit", 1))
	}
	if s.Marker != nil && len(*s.Marker) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Marker", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFolderPathInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AuthenticationToken != nil {
		v := *s.AuthenticationToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.HeaderTarget, "Authentication", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.FolderId != nil {
		v := *s.FolderId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "FolderId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Fields != nil {
		v := *s.Fields

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "fields", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Limit != nil {
		v := *s.Limit

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "limit", protocol.Int64Value(v), metadata)
	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type GetFolderPathOutput struct {
	_ struct{} `type:"structure"`

	// The path information.
	Path *ResourcePath `type:"structure"`
}

// String returns the string representation
func (s GetFolderPathOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s GetFolderPathOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Path != nil {
		v := s.Path

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "Path", v, metadata)
	}
	return nil
}

const opGetFolderPath = "GetFolderPath"

// GetFolderPathRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Retrieves the path information (the hierarchy from the root folder) for the
// specified folder.
//
// By default, Amazon WorkDocs returns a maximum of 100 levels upwards from
// the requested folder and only includes the IDs of the parent folders in the
// path. You can limit the maximum number of levels. You can also request the
// parent folder names.
//
//    // Example sending a request using GetFolderPathRequest.
//    req := client.GetFolderPathRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/GetFolderPath
func (c *Client) GetFolderPathRequest(input *GetFolderPathInput) GetFolderPathRequest {
	op := &aws.Operation{
		Name:       opGetFolderPath,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/folders/{FolderId}/path",
	}

	if input == nil {
		input = &GetFolderPathInput{}
	}

	req := c.newRequest(op, input, &GetFolderPathOutput{})
	return GetFolderPathRequest{Request: req, Input: input, Copy: c.GetFolderPathRequest}
}

// GetFolderPathRequest is the request type for the
// GetFolderPath API operation.
type GetFolderPathRequest struct {
	*aws.Request
	Input *GetFolderPathInput
	Copy  func(*GetFolderPathInput) GetFolderPathRequest
}

// Send marshals and sends the GetFolderPath API request.
func (r GetFolderPathRequest) Send(ctx context.Context) (*GetFolderPathResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFolderPathResponse{
		GetFolderPathOutput: r.Request.Data.(*GetFolderPathOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFolderPathResponse is the response type for the
// GetFolderPath API operation.
type GetFolderPathResponse struct {
	*GetFolderPathOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFolderPath request.
func (r *GetFolderPathResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
