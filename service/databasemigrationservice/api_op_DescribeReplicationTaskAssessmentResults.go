// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeReplicationTaskAssessmentResultsInput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that the remaining results can be retrieved.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The Amazon Resource Name (ARN) string that uniquely identifies the task.
	// When this input parameter is specified, the API returns only one result and
	// ignore the values of the MaxRecords and Marker parameters.
	ReplicationTaskArn *string `type:"string"`
}

// String returns the string representation
func (s DescribeReplicationTaskAssessmentResultsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeReplicationTaskAssessmentResultsOutput struct {
	_ struct{} `type:"structure"`

	// - The Amazon S3 bucket where the task assessment report is located.
	BucketName *string `type:"string"`

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to
	// the value specified by MaxRecords.
	Marker *string `type:"string"`

	// The task assessment report.
	ReplicationTaskAssessmentResults []ReplicationTaskAssessmentResult `type:"list"`
}

// String returns the string representation
func (s DescribeReplicationTaskAssessmentResultsOutput) String() string {
	return awsutil.Prettify(s)
}

const opDescribeReplicationTaskAssessmentResults = "DescribeReplicationTaskAssessmentResults"

// DescribeReplicationTaskAssessmentResultsRequest returns a request value for making API operation for
// AWS Database Migration Service.
//
// Returns the task assessment results from Amazon S3. This action always returns
// the latest results.
//
//    // Example sending a request using DescribeReplicationTaskAssessmentResultsRequest.
//    req := client.DescribeReplicationTaskAssessmentResultsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/dms-2016-01-01/DescribeReplicationTaskAssessmentResults
func (c *Client) DescribeReplicationTaskAssessmentResultsRequest(input *DescribeReplicationTaskAssessmentResultsInput) DescribeReplicationTaskAssessmentResultsRequest {
	op := &aws.Operation{
		Name:       opDescribeReplicationTaskAssessmentResults,
		HTTPMethod: "POST",
		HTTPPath:   "/",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"Marker"},
			LimitToken:      "MaxRecords",
			TruncationToken: "",
		},
	}

	if input == nil {
		input = &DescribeReplicationTaskAssessmentResultsInput{}
	}

	req := c.newRequest(op, input, &DescribeReplicationTaskAssessmentResultsOutput{})
	return DescribeReplicationTaskAssessmentResultsRequest{Request: req, Input: input, Copy: c.DescribeReplicationTaskAssessmentResultsRequest}
}

// DescribeReplicationTaskAssessmentResultsRequest is the request type for the
// DescribeReplicationTaskAssessmentResults API operation.
type DescribeReplicationTaskAssessmentResultsRequest struct {
	*aws.Request
	Input *DescribeReplicationTaskAssessmentResultsInput
	Copy  func(*DescribeReplicationTaskAssessmentResultsInput) DescribeReplicationTaskAssessmentResultsRequest
}

// Send marshals and sends the DescribeReplicationTaskAssessmentResults API request.
func (r DescribeReplicationTaskAssessmentResultsRequest) Send(ctx context.Context) (*DescribeReplicationTaskAssessmentResultsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeReplicationTaskAssessmentResultsResponse{
		DescribeReplicationTaskAssessmentResultsOutput: r.Request.Data.(*DescribeReplicationTaskAssessmentResultsOutput),
		response: &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// NewDescribeReplicationTaskAssessmentResultsRequestPaginator returns a paginator for DescribeReplicationTaskAssessmentResults.
// Use Next method to get the next page, and CurrentPage to get the current
// response page from the paginator. Next will return false, if there are
// no more pages, or an error was encountered.
//
// Note: This operation can generate multiple requests to a service.
//
//   // Example iterating over pages.
//   req := client.DescribeReplicationTaskAssessmentResultsRequest(input)
//   p := databasemigrationservice.NewDescribeReplicationTaskAssessmentResultsRequestPaginator(req)
//
//   for p.Next(context.TODO()) {
//       page := p.CurrentPage()
//   }
//
//   if err := p.Err(); err != nil {
//       return err
//   }
//
func NewDescribeReplicationTaskAssessmentResultsPaginator(req DescribeReplicationTaskAssessmentResultsRequest) DescribeReplicationTaskAssessmentResultsPaginator {
	return DescribeReplicationTaskAssessmentResultsPaginator{
		Pager: aws.Pager{
			NewRequest: func(ctx context.Context) (*aws.Request, error) {
				var inCpy *DescribeReplicationTaskAssessmentResultsInput
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

// DescribeReplicationTaskAssessmentResultsPaginator is used to paginate the request. This can be done by
// calling Next and CurrentPage.
type DescribeReplicationTaskAssessmentResultsPaginator struct {
	aws.Pager
}

func (p *DescribeReplicationTaskAssessmentResultsPaginator) CurrentPage() *DescribeReplicationTaskAssessmentResultsOutput {
	return p.Pager.CurrentPage().(*DescribeReplicationTaskAssessmentResultsOutput)
}

// DescribeReplicationTaskAssessmentResultsResponse is the response type for the
// DescribeReplicationTaskAssessmentResults API operation.
type DescribeReplicationTaskAssessmentResultsResponse struct {
	*DescribeReplicationTaskAssessmentResultsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeReplicationTaskAssessmentResults request.
func (r *DescribeReplicationTaskAssessmentResultsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
