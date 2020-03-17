// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package workdocs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type DescribeFolderContentsInput struct {
	_ struct{} `type:"structure"`

	// Amazon WorkDocs authentication token. Not required when using AWS administrator
	// credentials to access the API.
	AuthenticationToken *string `location:"header" locationName:"Authentication" min:"1" type:"string" sensitive:"true"`

	// The ID of the folder.
	//
	// FolderId is a required field
	FolderId *string `location:"uri" locationName:"FolderId" min:"1" type:"string" required:"true"`

	// The contents to include. Specify "INITIALIZED" to include initialized documents.
	Include *string `location:"querystring" locationName:"include" min:"1" type:"string"`

	// The maximum number of items to return with this call.
	Limit *int64 `location:"querystring" locationName:"limit" min:"1" type:"integer"`

	// The marker for the next set of results. This marker was received from a previous
	// call.
	Marker *string `location:"querystring" locationName:"marker" min:"1" type:"string"`

	// The order for the contents of the folder.
	Order OrderType `location:"querystring" locationName:"order" type:"string" enum:"true"`

	// The sorting criteria.
	Sort ResourceSortType `location:"querystring" locationName:"sort" type:"string" enum:"true"`

	// The type of items.
	Type FolderContentType `location:"querystring" locationName:"type" type:"string" enum:"true"`
}

// String returns the string representation
func (s DescribeFolderContentsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeFolderContentsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeFolderContentsInput"}
	if s.AuthenticationToken != nil && len(*s.AuthenticationToken) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("AuthenticationToken", 1))
	}

	if s.FolderId == nil {
		invalidParams.Add(aws.NewErrParamRequired("FolderId"))
	}
	if s.FolderId != nil && len(*s.FolderId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("FolderId", 1))
	}
	if s.Include != nil && len(*s.Include) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Include", 1))
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
func (s DescribeFolderContentsInput) MarshalFields(e protocol.FieldEncoder) error {
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
	if s.Include != nil {
		v := *s.Include

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "include", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
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
	if len(s.Order) > 0 {
		v := s.Order

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "order", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Sort) > 0 {
		v := s.Sort

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "sort", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	if len(s.Type) > 0 {
		v := s.Type

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "type", protocol.QuotedValue{ValueMarshaler: v}, metadata)
	}
	return nil
}

type DescribeFolderContentsOutput struct {
	_ struct{} `type:"structure"`

	// The documents in the specified folder.
	Documents []DocumentMetadata `type:"list"`

	// The subfolders in the specified folder.
	Folders []FolderMetadata `type:"list"`

	// The marker to use when requesting the next set of results. If there are no
	// additional results, the string is empty.
	Marker *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeFolderContentsOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeFolderContentsOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Documents != nil {
		v := s.Documents

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Documents", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Folders != nil {
		v := s.Folders

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Folders", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.Marker != nil {
		v := *s.Marker

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Marker", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opDescribeFolderContents = "DescribeFolderContents"

// DescribeFolderContentsRequest returns a request value for making API operation for
// Amazon WorkDocs.
//
// Describes the contents of the specified folder, including its documents and
// subfolders.
//
// By default, Amazon WorkDocs returns the first 100 active document and folder
// metadata items. If there are more results, the response includes a marker
// that you can use to request the next set of results. You can also request
// initialized documents.
//
//    // Example sending a request using DescribeFolderContentsRequest.
//    req := client.DescribeFolderContentsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/workdocs-2016-05-01/DescribeFolderContents
func (c *Client) DescribeFolderContentsRequest(input *DescribeFolderContentsInput) DescribeFolderContentsRequest {
	op := &aws.Operation{
		Name:       opDescribeFolderContents,
		HTTPMethod: "GET",
		HTTPPath:   "/api/v1/folders/{FolderId}/contents",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "Limit",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeFolderContentsInput{}
	}

	req := c.newRequest(op, input, &DescribeFolderContentsOutput{})
	return DescribeFolderContentsRequest{Request: req, Input: input, Copy: c.DescribeFolderContentsRequest}
}

// DescribeFolderContentsRequest is the request type for the
// DescribeFolderContents API operation.
type DescribeFolderContentsRequest struct {
	*aws.Request
	Input *DescribeFolderContentsInput
	Copy  func(*DescribeFolderContentsInput) DescribeFolderContentsRequest
}

// Send marshals and sends the DescribeFolderContents API request.
func (r DescribeFolderContentsRequest) Send(ctx context.Context) (*DescribeFolderContentsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeFolderContentsResponse{
		DescribeFolderContentsOutput: r.Request.Data.(*DescribeFolderContentsOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeFolderContentsRequestPaginator returns a paginator for DescribeFolderContents.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeFolderContentsRequest(input)
//   p := workdocs.NewDescribeFolderContentsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeFolderContentsPaginator(req DescribeFolderContentsRequest) DescribeFolderContentsPaginator {
	return DescribeFolderContentsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeFolderContentsInput
				if req.Input != nil {
					tmp := *req.Input
					inCpy = &tmp
				}

				newReq := req.Copy(inCpy)
				newReq.SetContext(ctx)
				return newReq.Request, nil
			},
		},
	}
}

// DescribeFolderContentsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeFolderContentsPaginator struct {
	aws.Pager
}

func (p *DescribeFolderContentsPaginator) CurrentPage() *DescribeFolderContentsOutput {
	return p.Pager.CurrentPage().(*DescribeFolderContentsOutput)
}

// DescribeFolderContentsResponse is the response type for the
// DescribeFolderContents API operation.
type DescribeFolderContentsResponse struct {
	*DescribeFolderContentsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeFolderContents request.
func (r *DescribeFolderContentsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
