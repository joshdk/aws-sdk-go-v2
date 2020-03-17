// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediatailor

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

type PutPlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The URL for the ad decision server (ADS). This includes the specification
	// of static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing you can provide a static
	// VAST URL. The maximum length is 25,000 characters.
	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The configuration for DASH content.
	DashConfiguration *DashConfigurationForPut `type:"structure"`

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *LivePreRollConfiguration `type:"structure"`

	// The identifier for the playback configuration.
	Name *string `type:"string"`

	PersonalizationThresholdSeconds *int64 `min:"1" type:"integer"`

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill
	// in gaps in media content. Configuring the slate is optional for non-VPAID
	// configurations. For VPAID, the slate is required because MediaTailor provides
	// it in the slots that are designated for dynamic ad content. The slate must
	// be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `type:"string"`

	// The tags to assign to the playback configuration.
	Tags map[string]string `locationName:"tags" type:"map"`

	// The name that is used to associate this playback configuration with a custom
	// transcode profile. This overrides the dynamic transcoding defaults of MediaTailor.
	// Use this only if you have already set up custom profiles with the help of
	// AWS Support.
	TranscodeProfileName *string `type:"string"`

	// The URL prefix for the master playlist for the stream, minus the asset ID.
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PutPlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutPlaybackConfigurationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutPlaybackConfigurationInput"}
	if s.PersonalizationThresholdSeconds != nil && *s.PersonalizationThresholdSeconds < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("PersonalizationThresholdSeconds", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPlaybackConfigurationInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/json"), protocol.Metadata{})

	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.DashConfiguration != nil {
		v := s.DashConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DashConfiguration", v, metadata)
	}
	if s.LivePreRollConfiguration != nil {
		v := s.LivePreRollConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LivePreRollConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PersonalizationThresholdSeconds != nil {
		v := *s.PersonalizationThresholdSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PersonalizationThresholdSeconds", protocol.Int64Value(v), metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TranscodeProfileName != nil {
		v := *s.TranscodeProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TranscodeProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

type PutPlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The configuration for DASH content.
	DashConfiguration *DashConfiguration `type:"structure"`

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration `type:"structure"`

	// The configuration for pre-roll ad insertion.
	LivePreRollConfiguration *LivePreRollConfiguration `type:"structure"`

	Name *string `type:"string"`

	PlaybackConfigurationArn *string `type:"string"`

	PlaybackEndpointPrefix *string `type:"string"`

	SessionInitializationEndpointPrefix *string `type:"string"`

	SlateAdUrl *string `type:"string"`

	Tags map[string]string `locationName:"tags" type:"map"`

	TranscodeProfileName *string `type:"string"`

	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PutPlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s PutPlaybackConfigurationOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.AdDecisionServerUrl != nil {
		v := *s.AdDecisionServerUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "AdDecisionServerUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CdnConfiguration != nil {
		v := s.CdnConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "CdnConfiguration", v, metadata)
	}
	if s.DashConfiguration != nil {
		v := s.DashConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "DashConfiguration", v, metadata)
	}
	if s.HlsConfiguration != nil {
		v := s.HlsConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "HlsConfiguration", v, metadata)
	}
	if s.LivePreRollConfiguration != nil {
		v := s.LivePreRollConfiguration

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "LivePreRollConfiguration", v, metadata)
	}
	if s.Name != nil {
		v := *s.Name

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "Name", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlaybackConfigurationArn != nil {
		v := *s.PlaybackConfigurationArn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackConfigurationArn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.PlaybackEndpointPrefix != nil {
		v := *s.PlaybackEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "PlaybackEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SessionInitializationEndpointPrefix != nil {
		v := *s.SessionInitializationEndpointPrefix

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SessionInitializationEndpointPrefix", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.SlateAdUrl != nil {
		v := *s.SlateAdUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "SlateAdUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.Tags != nil {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TranscodeProfileName != nil {
		v := *s.TranscodeProfileName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "TranscodeProfileName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.VideoContentSourceUrl != nil {
		v := *s.VideoContentSourceUrl

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "VideoContentSourceUrl", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

const opPutPlaybackConfiguration = "PutPlaybackConfiguration"

// PutPlaybackConfigurationRequest returns a request value for making API operation for
// AWS MediaTailor.
//
// Adds a new playback configuration to AWS Elemental MediaTailor.
//
//    // Example sending a request using PutPlaybackConfigurationRequest.
//    req := client.PutPlaybackConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/PutPlaybackConfiguration
func (c *Client) PutPlaybackConfigurationRequest(input *PutPlaybackConfigurationInput) PutPlaybackConfigurationRequest {
	op := &aws.Operation{
		Name:       opPutPlaybackConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/playbackConfiguration",
	}

	if input == nil {
		input = &PutPlaybackConfigurationInput{}
	}

	req := c.newRequest(op, input, &PutPlaybackConfigurationOutput{})
	return PutPlaybackConfigurationRequest{Request: req, Input: input, Copy: c.PutPlaybackConfigurationRequest}
}

// PutPlaybackConfigurationRequest is the request type for the
// PutPlaybackConfiguration API operation.
type PutPlaybackConfigurationRequest struct {
	*aws.Request
	Input *PutPlaybackConfigurationInput
	Copy  func(*PutPlaybackConfigurationInput) PutPlaybackConfigurationRequest
}

// Send marshals and sends the PutPlaybackConfiguration API request.
func (r PutPlaybackConfigurationRequest) Send(ctx context.Context) (*PutPlaybackConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutPlaybackConfigurationResponse{
		PutPlaybackConfigurationOutput: r.Request.Data.(*PutPlaybackConfigurationOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutPlaybackConfigurationResponse is the response type for the
// PutPlaybackConfiguration API operation.
type PutPlaybackConfigurationResponse struct {
	*PutPlaybackConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutPlaybackConfiguration request.
func (r *PutPlaybackConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
