// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotsitewiseiface provides an interface to enable mocking the AWS IoT SiteWise service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotsitewiseiface

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise"
)

// ClientAPI provides an interface to enable mocking the
// iotsitewise.Client methods. This make unit testing your code that
// calls out to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT SiteWise.
//    func myFunc(svc iotsitewiseiface.ClientAPI) bool {
//        // Make svc.AssociateAssets request
//    }
//
//    func main() {
//        cfg, err := external.LoadDefaultAWSConfig()
//        if err != nil {
//            panic("failed to load config, " + err.Error())
//        }
//
//        svc := iotsitewise.New(cfg)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockClientClient struct {
//        iotsitewiseiface.ClientPI
//    }
//    func (m *mockClientClient) AssociateAssets(input *iotsitewise.AssociateAssetsInput) (*iotsitewise.AssociateAssetsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockClientClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ClientAPI interface {
	AssociateAssetsRequest(*iotsitewise.AssociateAssetsInput) iotsitewise.AssociateAssetsRequest

	BatchAssociateProjectAssetsRequest(*iotsitewise.BatchAssociateProjectAssetsInput) iotsitewise.BatchAssociateProjectAssetsRequest

	BatchDisassociateProjectAssetsRequest(*iotsitewise.BatchDisassociateProjectAssetsInput) iotsitewise.BatchDisassociateProjectAssetsRequest

	BatchPutAssetPropertyValueRequest(*iotsitewise.BatchPutAssetPropertyValueInput) iotsitewise.BatchPutAssetPropertyValueRequest

	CreateAccessPolicyRequest(*iotsitewise.CreateAccessPolicyInput) iotsitewise.CreateAccessPolicyRequest

	CreateAssetRequest(*iotsitewise.CreateAssetInput) iotsitewise.CreateAssetRequest

	CreateAssetModelRequest(*iotsitewise.CreateAssetModelInput) iotsitewise.CreateAssetModelRequest

	CreateDashboardRequest(*iotsitewise.CreateDashboardInput) iotsitewise.CreateDashboardRequest

	CreateGatewayRequest(*iotsitewise.CreateGatewayInput) iotsitewise.CreateGatewayRequest

	CreatePortalRequest(*iotsitewise.CreatePortalInput) iotsitewise.CreatePortalRequest

	CreateProjectRequest(*iotsitewise.CreateProjectInput) iotsitewise.CreateProjectRequest

	DeleteAccessPolicyRequest(*iotsitewise.DeleteAccessPolicyInput) iotsitewise.DeleteAccessPolicyRequest

	DeleteAssetRequest(*iotsitewise.DeleteAssetInput) iotsitewise.DeleteAssetRequest

	DeleteAssetModelRequest(*iotsitewise.DeleteAssetModelInput) iotsitewise.DeleteAssetModelRequest

	DeleteDashboardRequest(*iotsitewise.DeleteDashboardInput) iotsitewise.DeleteDashboardRequest

	DeleteGatewayRequest(*iotsitewise.DeleteGatewayInput) iotsitewise.DeleteGatewayRequest

	DeletePortalRequest(*iotsitewise.DeletePortalInput) iotsitewise.DeletePortalRequest

	DeleteProjectRequest(*iotsitewise.DeleteProjectInput) iotsitewise.DeleteProjectRequest

	DescribeAccessPolicyRequest(*iotsitewise.DescribeAccessPolicyInput) iotsitewise.DescribeAccessPolicyRequest

	DescribeAssetRequest(*iotsitewise.DescribeAssetInput) iotsitewise.DescribeAssetRequest

	DescribeAssetModelRequest(*iotsitewise.DescribeAssetModelInput) iotsitewise.DescribeAssetModelRequest

	DescribeAssetPropertyRequest(*iotsitewise.DescribeAssetPropertyInput) iotsitewise.DescribeAssetPropertyRequest

	DescribeDashboardRequest(*iotsitewise.DescribeDashboardInput) iotsitewise.DescribeDashboardRequest

	DescribeGatewayRequest(*iotsitewise.DescribeGatewayInput) iotsitewise.DescribeGatewayRequest

	DescribeGatewayCapabilityConfigurationRequest(*iotsitewise.DescribeGatewayCapabilityConfigurationInput) iotsitewise.DescribeGatewayCapabilityConfigurationRequest

	DescribeLoggingOptionsRequest(*iotsitewise.DescribeLoggingOptionsInput) iotsitewise.DescribeLoggingOptionsRequest

	DescribePortalRequest(*iotsitewise.DescribePortalInput) iotsitewise.DescribePortalRequest

	DescribeProjectRequest(*iotsitewise.DescribeProjectInput) iotsitewise.DescribeProjectRequest

	DisassociateAssetsRequest(*iotsitewise.DisassociateAssetsInput) iotsitewise.DisassociateAssetsRequest

	GetAssetPropertyAggregatesRequest(*iotsitewise.GetAssetPropertyAggregatesInput) iotsitewise.GetAssetPropertyAggregatesRequest

	GetAssetPropertyValueRequest(*iotsitewise.GetAssetPropertyValueInput) iotsitewise.GetAssetPropertyValueRequest

	GetAssetPropertyValueHistoryRequest(*iotsitewise.GetAssetPropertyValueHistoryInput) iotsitewise.GetAssetPropertyValueHistoryRequest

	ListAccessPoliciesRequest(*iotsitewise.ListAccessPoliciesInput) iotsitewise.ListAccessPoliciesRequest

	ListAssetModelsRequest(*iotsitewise.ListAssetModelsInput) iotsitewise.ListAssetModelsRequest

	ListAssetsRequest(*iotsitewise.ListAssetsInput) iotsitewise.ListAssetsRequest

	ListAssociatedAssetsRequest(*iotsitewise.ListAssociatedAssetsInput) iotsitewise.ListAssociatedAssetsRequest

	ListDashboardsRequest(*iotsitewise.ListDashboardsInput) iotsitewise.ListDashboardsRequest

	ListGatewaysRequest(*iotsitewise.ListGatewaysInput) iotsitewise.ListGatewaysRequest

	ListPortalsRequest(*iotsitewise.ListPortalsInput) iotsitewise.ListPortalsRequest

	ListProjectAssetsRequest(*iotsitewise.ListProjectAssetsInput) iotsitewise.ListProjectAssetsRequest

	ListProjectsRequest(*iotsitewise.ListProjectsInput) iotsitewise.ListProjectsRequest

	ListTagsForResourceRequest(*iotsitewise.ListTagsForResourceInput) iotsitewise.ListTagsForResourceRequest

	PutLoggingOptionsRequest(*iotsitewise.PutLoggingOptionsInput) iotsitewise.PutLoggingOptionsRequest

	TagResourceRequest(*iotsitewise.TagResourceInput) iotsitewise.TagResourceRequest

	UntagResourceRequest(*iotsitewise.UntagResourceInput) iotsitewise.UntagResourceRequest

	UpdateAccessPolicyRequest(*iotsitewise.UpdateAccessPolicyInput) iotsitewise.UpdateAccessPolicyRequest

	UpdateAssetRequest(*iotsitewise.UpdateAssetInput) iotsitewise.UpdateAssetRequest

	UpdateAssetModelRequest(*iotsitewise.UpdateAssetModelInput) iotsitewise.UpdateAssetModelRequest

	UpdateAssetPropertyRequest(*iotsitewise.UpdateAssetPropertyInput) iotsitewise.UpdateAssetPropertyRequest

	UpdateDashboardRequest(*iotsitewise.UpdateDashboardInput) iotsitewise.UpdateDashboardRequest

	UpdateGatewayRequest(*iotsitewise.UpdateGatewayInput) iotsitewise.UpdateGatewayRequest

	UpdateGatewayCapabilityConfigurationRequest(*iotsitewise.UpdateGatewayCapabilityConfigurationInput) iotsitewise.UpdateGatewayCapabilityConfigurationRequest

	UpdatePortalRequest(*iotsitewise.UpdatePortalInput) iotsitewise.UpdatePortalRequest

	UpdateProjectRequest(*iotsitewise.UpdateProjectInput) iotsitewise.UpdateProjectRequest

	WaitUntilAssetActive(context.Context, *iotsitewise.DescribeAssetInput, ...aws.WaiterOption) error

	WaitUntilAssetModelActive(context.Context, *iotsitewise.DescribeAssetModelInput, ...aws.WaiterOption) error

	WaitUntilAssetModelNotExists(context.Context, *iotsitewise.DescribeAssetModelInput, ...aws.WaiterOption) error

	WaitUntilAssetNotExists(context.Context, *iotsitewise.DescribeAssetInput, ...aws.WaiterOption) error

	WaitUntilPortalActive(context.Context, *iotsitewise.DescribePortalInput, ...aws.WaiterOption) error

	WaitUntilPortalNotExists(context.Context, *iotsitewise.DescribePortalInput, ...aws.WaiterOption) error
}

var _ ClientAPI = (*iotsitewise.Client)(nil)