// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elastictranscoder

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// The ListJobsByPipelineRequest structure.
type ListJobsByPipelineInput struct {
	_ struct{} `type:"structure"`

	// To list jobs in chronological order by the date and time that they were submitted,
	// enter true. To list jobs in reverse chronological order, enter false.
	Ascending *string `location:"querystring" locationName:"Ascending" type:"string"`

	// When Elastic Transcoder returns more than one page of results, use pageToken
	// in subsequent GET requests to get each successive page of results.
	PageToken *string `location:"querystring" locationName:"PageToken" type:"string"`

	// The ID of the pipeline for which you want to get job information.
	//
	// PipelineId is a required field
	PipelineId *string `location:"uri" locationName:"PipelineId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListJobsByPipelineInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJobsByPipelineInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJobsByPipelineInput"}

	if s.PipelineId == nil {
		invalidParams.Add(aws.NewErrParamRequired("PipelineId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsByPipelineInput) MarshalFields(e protocol.FieldEncoder) error {

	if s.PipelineId != nil {
		v := *s.PipelineId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "PipelineId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Ascending != nil {
		v := *s.Ascending

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "Ascending", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PageToken != nil {
		v := *s.PageToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.QueryTarget, "PageToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// The ListJobsByPipelineResponse structure.
type ListJobsByPipelineOutput struct {
	_ struct{} `type:"structure"`

	// An array of Job objects that are in the specified pipeline.
	Jobs []Job `type:"list"`

	// A value that you use to access the second and subsequent pages of results,
	// if any. When the jobs in the specified pipeline fit on one page or when you've
	// reached the last page of results, the value of NextPageToken is null.
	NextPageToken *string `type:"string"`
}

// String returns the string representation
func (s ListJobsByPipelineOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s ListJobsByPipelineOutput) MarshalFields(e protocol.FieldEncoder) error {
	if len(s.Jobs) > 0 {
		v := s.Jobs

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "Jobs", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddFields(v1)
		}
		ls0.End()

	}
	if s.NextPageToken != nil {
		v := *s.NextPageToken

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "NextPageToken", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opListJobsByPipeline = "ListJobsByPipeline"

// ListJobsByPipelineRequest returns a request value for making API operation for
// Amazon Elastic Transcoder.
//
// The ListJobsByPipeline operation gets a list of the jobs currently in a pipeline.
//
// Elastic Transcoder returns all of the jobs currently in the specified pipeline.
// The response body contains one element for each job that satisfies the search
// criteria.
//
//    // Example sending a request using ListJobsByPipelineRequest.
//    req := client.ListJobsByPipelineRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) ListJobsByPipelineRequest(input *ListJobsByPipelineInput) ListJobsByPipelineRequest {
	op := &aws.Operation{
		Name:       opListJobsByPipeline,
		HTTPMethod: "GET",
		HTTPPath:   "/2012-09-25/jobsByPipeline/{PipelineId}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"PageToken"},
			OutputTokens:    []string{"NextPageToken"},
			LimitToken:      "",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &ListJobsByPipelineInput{}
	}

	req := c.newRequest(op, input, &ListJobsByPipelineOutput{})
	return ListJobsByPipelineRequest{Request: req, Input: input, Copy: c.ListJobsByPipelineRequest}
}

// ListJobsByPipelineRequest is the request type for the
// ListJobsByPipeline API operation.
type ListJobsByPipelineRequest struct {
	*aws.Request
	Input *ListJobsByPipelineInput
	Copy  func(*ListJobsByPipelineInput) ListJobsByPipelineRequest
}

// Send marshals and sends the ListJobsByPipeline API request.
func (r ListJobsByPipelineRequest) Send(ctx context.Context) (*ListJobsByPipelineResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ListJobsByPipelineResponse{
		ListJobsByPipelineOutput: r.Request.Data.(*ListJobsByPipelineOutput),
		response:                 &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewListJobsByPipelineRequestPaginator returns a paginator for ListJobsByPipeline.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.ListJobsByPipelineRequest(input)
//   p := elastictranscoder.NewListJobsByPipelineRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewListJobsByPipelinePaginator(req ListJobsByPipelineRequest) ListJobsByPipelinePaginator {
	return ListJobsByPipelinePaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *ListJobsByPipelineInput
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

// ListJobsByPipelinePaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type ListJobsByPipelinePaginator struct {
	aws.Pager
}

func (p *ListJobsByPipelinePaginator) CurrentPage() *ListJobsByPipelineOutput {
	return p.Pager.CurrentPage().(*ListJobsByPipelineOutput)
}

// ListJobsByPipelineResponse is the response type for the
// ListJobsByPipeline API operation.
type ListJobsByPipelineResponse struct {
	*ListJobsByPipelineOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ListJobsByPipeline request.
func (r *ListJobsByPipelineResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}