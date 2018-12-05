// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediatailor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

const opDeletePlaybackConfiguration = "DeletePlaybackConfiguration"

// DeletePlaybackConfigurationRequest generates a "aws/request.Request" representing the
// client's request for the DeletePlaybackConfiguration operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See DeletePlaybackConfiguration for more information on using the DeletePlaybackConfiguration
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the DeletePlaybackConfigurationRequest method.
//    req, resp := client.DeletePlaybackConfigurationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/DeletePlaybackConfiguration
func (c *MediaTailor) DeletePlaybackConfigurationRequest(input *DeletePlaybackConfigurationInput) (req *request.Request, output *DeletePlaybackConfigurationOutput) {
	op := &request.Operation{
		Name:       opDeletePlaybackConfiguration,
		HTTPMethod: "DELETE",
		HTTPPath:   "/playbackConfiguration/{Name}",
	}

	if input == nil {
		input = &DeletePlaybackConfigurationInput{}
	}

	output = &DeletePlaybackConfigurationOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// DeletePlaybackConfiguration API operation for AWS MediaTailor.
//
// Deletes the configuration for the specified name.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS MediaTailor's
// API operation DeletePlaybackConfiguration for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/DeletePlaybackConfiguration
func (c *MediaTailor) DeletePlaybackConfiguration(input *DeletePlaybackConfigurationInput) (*DeletePlaybackConfigurationOutput, error) {
	req, out := c.DeletePlaybackConfigurationRequest(input)
	return out, req.Send()
}

// DeletePlaybackConfigurationWithContext is the same as DeletePlaybackConfiguration with the addition of
// the ability to pass a context and additional request options.
//
// See DeletePlaybackConfiguration for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaTailor) DeletePlaybackConfigurationWithContext(ctx aws.Context, input *DeletePlaybackConfigurationInput, opts ...request.Option) (*DeletePlaybackConfigurationOutput, error) {
	req, out := c.DeletePlaybackConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opGetPlaybackConfiguration = "GetPlaybackConfiguration"

// GetPlaybackConfigurationRequest generates a "aws/request.Request" representing the
// client's request for the GetPlaybackConfiguration operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See GetPlaybackConfiguration for more information on using the GetPlaybackConfiguration
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the GetPlaybackConfigurationRequest method.
//    req, resp := client.GetPlaybackConfigurationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfiguration
func (c *MediaTailor) GetPlaybackConfigurationRequest(input *GetPlaybackConfigurationInput) (req *request.Request, output *GetPlaybackConfigurationOutput) {
	op := &request.Operation{
		Name:       opGetPlaybackConfiguration,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfiguration/{Name}",
	}

	if input == nil {
		input = &GetPlaybackConfigurationInput{}
	}

	output = &GetPlaybackConfigurationOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetPlaybackConfiguration API operation for AWS MediaTailor.
//
// Returns the configuration for the specified name.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS MediaTailor's
// API operation GetPlaybackConfiguration for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/GetPlaybackConfiguration
func (c *MediaTailor) GetPlaybackConfiguration(input *GetPlaybackConfigurationInput) (*GetPlaybackConfigurationOutput, error) {
	req, out := c.GetPlaybackConfigurationRequest(input)
	return out, req.Send()
}

// GetPlaybackConfigurationWithContext is the same as GetPlaybackConfiguration with the addition of
// the ability to pass a context and additional request options.
//
// See GetPlaybackConfiguration for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaTailor) GetPlaybackConfigurationWithContext(ctx aws.Context, input *GetPlaybackConfigurationInput, opts ...request.Option) (*GetPlaybackConfigurationOutput, error) {
	req, out := c.GetPlaybackConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opListPlaybackConfigurations = "ListPlaybackConfigurations"

// ListPlaybackConfigurationsRequest generates a "aws/request.Request" representing the
// client's request for the ListPlaybackConfigurations operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See ListPlaybackConfigurations for more information on using the ListPlaybackConfigurations
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the ListPlaybackConfigurationsRequest method.
//    req, resp := client.ListPlaybackConfigurationsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/ListPlaybackConfigurations
func (c *MediaTailor) ListPlaybackConfigurationsRequest(input *ListPlaybackConfigurationsInput) (req *request.Request, output *ListPlaybackConfigurationsOutput) {
	op := &request.Operation{
		Name:       opListPlaybackConfigurations,
		HTTPMethod: "GET",
		HTTPPath:   "/playbackConfigurations",
	}

	if input == nil {
		input = &ListPlaybackConfigurationsInput{}
	}

	output = &ListPlaybackConfigurationsOutput{}
	req = c.newRequest(op, input, output)
	return
}

// ListPlaybackConfigurations API operation for AWS MediaTailor.
//
// Returns a list of the configurations defined in AWS Elemental MediaTailor.
// You can specify a max number of configurations to return at a time. The default
// max is 50. Results are returned in pagefuls. If AWS Elemental MediaTailor
// has more configurations than the specified max, it provides parameters in
// the response that you can use to retrieve the next pageful.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS MediaTailor's
// API operation ListPlaybackConfigurations for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/ListPlaybackConfigurations
func (c *MediaTailor) ListPlaybackConfigurations(input *ListPlaybackConfigurationsInput) (*ListPlaybackConfigurationsOutput, error) {
	req, out := c.ListPlaybackConfigurationsRequest(input)
	return out, req.Send()
}

// ListPlaybackConfigurationsWithContext is the same as ListPlaybackConfigurations with the addition of
// the ability to pass a context and additional request options.
//
// See ListPlaybackConfigurations for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaTailor) ListPlaybackConfigurationsWithContext(ctx aws.Context, input *ListPlaybackConfigurationsInput, opts ...request.Option) (*ListPlaybackConfigurationsOutput, error) {
	req, out := c.ListPlaybackConfigurationsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opPutPlaybackConfiguration = "PutPlaybackConfiguration"

// PutPlaybackConfigurationRequest generates a "aws/request.Request" representing the
// client's request for the PutPlaybackConfiguration operation. The "output" return
// value will be populated with the request's response once the request completes
// successfully.
//
// Use "Send" method on the returned Request to send the API call to the service.
// the "output" return value is not valid until after Send returns without error.
//
// See PutPlaybackConfiguration for more information on using the PutPlaybackConfiguration
// API call, and error handling.
//
// This method is useful when you want to inject custom logic or configuration
// into the SDK's request lifecycle. Such as custom headers, or retry logic.
//
//
//    // Example sending a request using the PutPlaybackConfigurationRequest method.
//    req, resp := client.PutPlaybackConfigurationRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/PutPlaybackConfiguration
func (c *MediaTailor) PutPlaybackConfigurationRequest(input *PutPlaybackConfigurationInput) (req *request.Request, output *PutPlaybackConfigurationOutput) {
	op := &request.Operation{
		Name:       opPutPlaybackConfiguration,
		HTTPMethod: "PUT",
		HTTPPath:   "/playbackConfiguration",
	}

	if input == nil {
		input = &PutPlaybackConfigurationInput{}
	}

	output = &PutPlaybackConfigurationOutput{}
	req = c.newRequest(op, input, output)
	return
}

// PutPlaybackConfiguration API operation for AWS MediaTailor.
//
// Adds a new configuration to AWS Elemental MediaTailor.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS MediaTailor's
// API operation PutPlaybackConfiguration for usage and error information.
// See also, https://docs.aws.amazon.com/goto/WebAPI/mediatailor-2018-04-23/PutPlaybackConfiguration
func (c *MediaTailor) PutPlaybackConfiguration(input *PutPlaybackConfigurationInput) (*PutPlaybackConfigurationOutput, error) {
	req, out := c.PutPlaybackConfigurationRequest(input)
	return out, req.Send()
}

// PutPlaybackConfigurationWithContext is the same as PutPlaybackConfiguration with the addition of
// the ability to pass a context and additional request options.
//
// See PutPlaybackConfiguration for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *MediaTailor) PutPlaybackConfigurationWithContext(ctx aws.Context, input *PutPlaybackConfigurationInput, opts ...request.Option) (*PutPlaybackConfigurationOutput, error) {
	req, out := c.PutPlaybackConfigurationRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

// The configuration for using a content delivery network (CDN), like Amazon
// CloudFront, for content and ad segment management.
type CdnConfiguration struct {
	_ struct{} `type:"structure"`

	// A non-default content delivery network (CDN) to serve ad segments. By default,
	// AWS Elemental MediaTailor uses Amazon CloudFront with default cache settings
	// as its CDN for ad segments. To set up an alternate CDN, create a rule in
	// your CDN for the following origin: ads.mediatailor.<region>.amazonaws.com.
	// Then specify the rule's name in this AdSegmentUrlPrefix. When AWS Elemental
	// MediaTailor serves a manifest, it reports your CDN as the source for ad segments.
	AdSegmentUrlPrefix *string `type:"string"`

	// A content delivery network (CDN) to cache content segments, so that content
	// requests don’t always have to go to the origin server. First, create a rule
	// in your CDN for the content segment origin server. Then specify the rule's
	// name in this ContentSegmentUrlPrefix. When AWS Elemental MediaTailor serves
	// a manifest, it reports your CDN as the source for content segments.
	ContentSegmentUrlPrefix *string `type:"string"`
}

// String returns the string representation
func (s CdnConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CdnConfiguration) GoString() string {
	return s.String()
}

// SetAdSegmentUrlPrefix sets the AdSegmentUrlPrefix field's value.
func (s *CdnConfiguration) SetAdSegmentUrlPrefix(v string) *CdnConfiguration {
	s.AdSegmentUrlPrefix = &v
	return s
}

// SetContentSegmentUrlPrefix sets the ContentSegmentUrlPrefix field's value.
func (s *CdnConfiguration) SetContentSegmentUrlPrefix(v string) *CdnConfiguration {
	s.ContentSegmentUrlPrefix = &v
	return s
}

// The configuration object for dash content.
type DashConfiguration struct {
	_ struct{} `type:"structure"`

	// The URL that is used to initiate a playback session for devices that support
	// DASH.
	ManifestEndpointPrefix *string `type:"string"`

	// The setting that controls whether MediaTailor includes the Location tag in
	// DASH Manifests. MediaTailor populates the Location tag with the URL for manifest
	// update requests, to be used by players that don't support sticky redirects.
	// Disable this if you have CDN routing rules set up for accessing MediaTailor
	// manifests and you are either using client-side reporting or your players
	// support sticky HTTP redirects. Valid values are DISABLED and EMT_DEFAULT.
	// The EMT_DEFAULT setting enables the inclusion of the tag and is the default
	// value.
	MpdLocation *string `type:"string"`
}

// String returns the string representation
func (s DashConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DashConfiguration) GoString() string {
	return s.String()
}

// SetManifestEndpointPrefix sets the ManifestEndpointPrefix field's value.
func (s *DashConfiguration) SetManifestEndpointPrefix(v string) *DashConfiguration {
	s.ManifestEndpointPrefix = &v
	return s
}

// SetMpdLocation sets the MpdLocation field's value.
func (s *DashConfiguration) SetMpdLocation(v string) *DashConfiguration {
	s.MpdLocation = &v
	return s
}

type DashConfigurationForPut struct {
	_ struct{} `type:"structure"`

	MpdLocation *string `type:"string"`
}

// String returns the string representation
func (s DashConfigurationForPut) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DashConfigurationForPut) GoString() string {
	return s.String()
}

// SetMpdLocation sets the MpdLocation field's value.
func (s *DashConfigurationForPut) SetMpdLocation(v string) *DashConfigurationForPut {
	s.MpdLocation = &v
	return s
}

type DeletePlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `location:"uri" locationName:"Name" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePlaybackConfigurationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePlaybackConfigurationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeletePlaybackConfigurationInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetName sets the Name field's value.
func (s *DeletePlaybackConfigurationInput) SetName(v string) *DeletePlaybackConfigurationInput {
	s.Name = &v
	return s
}

type DeletePlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeletePlaybackConfigurationOutput) GoString() string {
	return s.String()
}

type GetPlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// Name is a required field
	Name *string `location:"uri" locationName:"Name" type:"string" required:"true"`
}

// String returns the string representation
func (s GetPlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPlaybackConfigurationInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetPlaybackConfigurationInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetPlaybackConfigurationInput"}
	if s.Name == nil {
		invalidParams.Add(request.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("Name", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetName sets the Name field's value.
func (s *GetPlaybackConfigurationInput) SetName(v string) *GetPlaybackConfigurationInput {
	s.Name = &v
	return s
}

type GetPlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	// The URL for the ad decision server (ADS). This includes the specification
	// of static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing, you can provide a
	// static VAST URL. The maximum length is 25000 characters.
	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The configuration object for dash content.
	DashConfiguration *DashConfiguration `type:"structure"`

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration `type:"structure"`

	// The identifier for the configuration.
	Name *string `type:"string"`

	// The URL that the player accesses to get a manifest from AWS Elemental MediaTailor.
	// This session will use server-side reporting.
	PlaybackEndpointPrefix *string `type:"string"`

	// The URL that the player uses to initialize a session that uses client-side
	// reporting.
	SessionInitializationEndpointPrefix *string `type:"string"`

	// URL for a high-quality video asset to transcode and use to fill in time that's
	// not used by ads. AWS Elemental MediaTailor shows the slate to fill in gaps
	// in media content. Configuring the slate is optional for non-VPAID configurations.
	// For VPAID, the slate is required because AWS Elemental MediaTailor provides
	// it in the slots designated for dynamic ad content. The slate must be a high-quality
	// asset that contains both audio and video.
	SlateAdUrl *string `type:"string"`

	// Associate this playbackConfiguration with a custom transcode profile, overriding
	// MediaTailor's dynamic transcoding defaults. Do not include this field if
	// you have not setup custom profiles with the MediaTailor service team.
	TranscodeProfileName *string `type:"string"`

	// The URL prefix for the master playlist for the stream, minus the asset ID.
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s GetPlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetPlaybackConfigurationOutput) GoString() string {
	return s.String()
}

// SetAdDecisionServerUrl sets the AdDecisionServerUrl field's value.
func (s *GetPlaybackConfigurationOutput) SetAdDecisionServerUrl(v string) *GetPlaybackConfigurationOutput {
	s.AdDecisionServerUrl = &v
	return s
}

// SetCdnConfiguration sets the CdnConfiguration field's value.
func (s *GetPlaybackConfigurationOutput) SetCdnConfiguration(v *CdnConfiguration) *GetPlaybackConfigurationOutput {
	s.CdnConfiguration = v
	return s
}

// SetDashConfiguration sets the DashConfiguration field's value.
func (s *GetPlaybackConfigurationOutput) SetDashConfiguration(v *DashConfiguration) *GetPlaybackConfigurationOutput {
	s.DashConfiguration = v
	return s
}

// SetHlsConfiguration sets the HlsConfiguration field's value.
func (s *GetPlaybackConfigurationOutput) SetHlsConfiguration(v *HlsConfiguration) *GetPlaybackConfigurationOutput {
	s.HlsConfiguration = v
	return s
}

// SetName sets the Name field's value.
func (s *GetPlaybackConfigurationOutput) SetName(v string) *GetPlaybackConfigurationOutput {
	s.Name = &v
	return s
}

// SetPlaybackEndpointPrefix sets the PlaybackEndpointPrefix field's value.
func (s *GetPlaybackConfigurationOutput) SetPlaybackEndpointPrefix(v string) *GetPlaybackConfigurationOutput {
	s.PlaybackEndpointPrefix = &v
	return s
}

// SetSessionInitializationEndpointPrefix sets the SessionInitializationEndpointPrefix field's value.
func (s *GetPlaybackConfigurationOutput) SetSessionInitializationEndpointPrefix(v string) *GetPlaybackConfigurationOutput {
	s.SessionInitializationEndpointPrefix = &v
	return s
}

// SetSlateAdUrl sets the SlateAdUrl field's value.
func (s *GetPlaybackConfigurationOutput) SetSlateAdUrl(v string) *GetPlaybackConfigurationOutput {
	s.SlateAdUrl = &v
	return s
}

// SetTranscodeProfileName sets the TranscodeProfileName field's value.
func (s *GetPlaybackConfigurationOutput) SetTranscodeProfileName(v string) *GetPlaybackConfigurationOutput {
	s.TranscodeProfileName = &v
	return s
}

// SetVideoContentSourceUrl sets the VideoContentSourceUrl field's value.
func (s *GetPlaybackConfigurationOutput) SetVideoContentSourceUrl(v string) *GetPlaybackConfigurationOutput {
	s.VideoContentSourceUrl = &v
	return s
}

// The configuration for HLS content.
type HlsConfiguration struct {
	_ struct{} `type:"structure"`

	// The URL that is used to initiate a playback session for devices that support
	// Apple HLS. The session uses server-side reporting.
	ManifestEndpointPrefix *string `type:"string"`
}

// String returns the string representation
func (s HlsConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s HlsConfiguration) GoString() string {
	return s.String()
}

// SetManifestEndpointPrefix sets the ManifestEndpointPrefix field's value.
func (s *HlsConfiguration) SetManifestEndpointPrefix(v string) *HlsConfiguration {
	s.ManifestEndpointPrefix = &v
	return s
}

type ListPlaybackConfigurationsInput struct {
	_ struct{} `type:"structure"`

	MaxResults *int64 `location:"querystring" locationName:"MaxResults" min:"1" type:"integer"`

	NextToken *string `location:"querystring" locationName:"NextToken" type:"string"`
}

// String returns the string representation
func (s ListPlaybackConfigurationsInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPlaybackConfigurationsInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListPlaybackConfigurationsInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ListPlaybackConfigurationsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(request.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMaxResults sets the MaxResults field's value.
func (s *ListPlaybackConfigurationsInput) SetMaxResults(v int64) *ListPlaybackConfigurationsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListPlaybackConfigurationsInput) SetNextToken(v string) *ListPlaybackConfigurationsInput {
	s.NextToken = &v
	return s
}

type ListPlaybackConfigurationsOutput struct {
	_ struct{} `type:"structure"`

	// Array of playback configurations. This may be all of the available configurations
	// or a subset, depending on the settings you provide and on the total number
	// of configurations stored.
	Items []*PlaybackConfiguration `type:"list"`

	// Pagination token returned by the GET list request when results overrun the
	// meximum allowed. Use the token to fetch the next page of results.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListPlaybackConfigurationsOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ListPlaybackConfigurationsOutput) GoString() string {
	return s.String()
}

// SetItems sets the Items field's value.
func (s *ListPlaybackConfigurationsOutput) SetItems(v []*PlaybackConfiguration) *ListPlaybackConfigurationsOutput {
	s.Items = v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *ListPlaybackConfigurationsOutput) SetNextToken(v string) *ListPlaybackConfigurationsOutput {
	s.NextToken = &v
	return s
}

type PlaybackConfiguration struct {
	_ struct{} `type:"structure"`

	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	Name *string `type:"string"`

	SlateAdUrl *string `type:"string"`

	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PlaybackConfiguration) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PlaybackConfiguration) GoString() string {
	return s.String()
}

// SetAdDecisionServerUrl sets the AdDecisionServerUrl field's value.
func (s *PlaybackConfiguration) SetAdDecisionServerUrl(v string) *PlaybackConfiguration {
	s.AdDecisionServerUrl = &v
	return s
}

// SetCdnConfiguration sets the CdnConfiguration field's value.
func (s *PlaybackConfiguration) SetCdnConfiguration(v *CdnConfiguration) *PlaybackConfiguration {
	s.CdnConfiguration = v
	return s
}

// SetName sets the Name field's value.
func (s *PlaybackConfiguration) SetName(v string) *PlaybackConfiguration {
	s.Name = &v
	return s
}

// SetSlateAdUrl sets the SlateAdUrl field's value.
func (s *PlaybackConfiguration) SetSlateAdUrl(v string) *PlaybackConfiguration {
	s.SlateAdUrl = &v
	return s
}

// SetVideoContentSourceUrl sets the VideoContentSourceUrl field's value.
func (s *PlaybackConfiguration) SetVideoContentSourceUrl(v string) *PlaybackConfiguration {
	s.VideoContentSourceUrl = &v
	return s
}

type PutPlaybackConfigurationInput struct {
	_ struct{} `type:"structure"`

	// The URL for the ad decision server (ADS). This includes the specification
	// of static parameters and placeholders for dynamic parameters. AWS Elemental
	// MediaTailor substitutes player-specific and session-specific parameters as
	// needed when calling the ADS. Alternately, for testing you can provide a static
	// VAST URL. The maximum length is 25000 characters.
	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	DashConfiguration *DashConfigurationForPut `type:"structure"`

	// The identifier for the configuration.
	Name *string `type:"string"`

	// The URL for a high-quality video asset to transcode and use to fill in time
	// that's not used by ads. AWS Elemental MediaTailor shows the slate to fill
	// in gaps in media content. Configuring the slate is optional for non-VPAID
	// configurations. For VPAID, the slate is required because AWS Elemental MediaTailor
	// provides it in the slots that are designated for dynamic ad content. The
	// slate must be a high-quality asset that contains both audio and video.
	SlateAdUrl *string `type:"string"`

	// Associate this playbackConfiguration with a custom transcode profile, overriding
	// MediaTailor's dynamic transcoding defaults. Do not include this field if
	// you have not setup custom profiles with the MediaTailor service team.
	TranscodeProfileName *string `type:"string"`

	// The URL prefix for the master playlist for the stream, minus the asset ID.
	// The maximum length is 512 characters.
	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PutPlaybackConfigurationInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutPlaybackConfigurationInput) GoString() string {
	return s.String()
}

// SetAdDecisionServerUrl sets the AdDecisionServerUrl field's value.
func (s *PutPlaybackConfigurationInput) SetAdDecisionServerUrl(v string) *PutPlaybackConfigurationInput {
	s.AdDecisionServerUrl = &v
	return s
}

// SetCdnConfiguration sets the CdnConfiguration field's value.
func (s *PutPlaybackConfigurationInput) SetCdnConfiguration(v *CdnConfiguration) *PutPlaybackConfigurationInput {
	s.CdnConfiguration = v
	return s
}

// SetDashConfiguration sets the DashConfiguration field's value.
func (s *PutPlaybackConfigurationInput) SetDashConfiguration(v *DashConfigurationForPut) *PutPlaybackConfigurationInput {
	s.DashConfiguration = v
	return s
}

// SetName sets the Name field's value.
func (s *PutPlaybackConfigurationInput) SetName(v string) *PutPlaybackConfigurationInput {
	s.Name = &v
	return s
}

// SetSlateAdUrl sets the SlateAdUrl field's value.
func (s *PutPlaybackConfigurationInput) SetSlateAdUrl(v string) *PutPlaybackConfigurationInput {
	s.SlateAdUrl = &v
	return s
}

// SetTranscodeProfileName sets the TranscodeProfileName field's value.
func (s *PutPlaybackConfigurationInput) SetTranscodeProfileName(v string) *PutPlaybackConfigurationInput {
	s.TranscodeProfileName = &v
	return s
}

// SetVideoContentSourceUrl sets the VideoContentSourceUrl field's value.
func (s *PutPlaybackConfigurationInput) SetVideoContentSourceUrl(v string) *PutPlaybackConfigurationInput {
	s.VideoContentSourceUrl = &v
	return s
}

type PutPlaybackConfigurationOutput struct {
	_ struct{} `type:"structure"`

	AdDecisionServerUrl *string `type:"string"`

	// The configuration for using a content delivery network (CDN), like Amazon
	// CloudFront, for content and ad segment management.
	CdnConfiguration *CdnConfiguration `type:"structure"`

	// The configuration object for dash content.
	DashConfiguration *DashConfiguration `type:"structure"`

	// The configuration for HLS content.
	HlsConfiguration *HlsConfiguration `type:"structure"`

	Name *string `type:"string"`

	PlaybackEndpointPrefix *string `type:"string"`

	SessionInitializationEndpointPrefix *string `type:"string"`

	SlateAdUrl *string `type:"string"`

	TranscodeProfileName *string `type:"string"`

	VideoContentSourceUrl *string `type:"string"`
}

// String returns the string representation
func (s PutPlaybackConfigurationOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PutPlaybackConfigurationOutput) GoString() string {
	return s.String()
}

// SetAdDecisionServerUrl sets the AdDecisionServerUrl field's value.
func (s *PutPlaybackConfigurationOutput) SetAdDecisionServerUrl(v string) *PutPlaybackConfigurationOutput {
	s.AdDecisionServerUrl = &v
	return s
}

// SetCdnConfiguration sets the CdnConfiguration field's value.
func (s *PutPlaybackConfigurationOutput) SetCdnConfiguration(v *CdnConfiguration) *PutPlaybackConfigurationOutput {
	s.CdnConfiguration = v
	return s
}

// SetDashConfiguration sets the DashConfiguration field's value.
func (s *PutPlaybackConfigurationOutput) SetDashConfiguration(v *DashConfiguration) *PutPlaybackConfigurationOutput {
	s.DashConfiguration = v
	return s
}

// SetHlsConfiguration sets the HlsConfiguration field's value.
func (s *PutPlaybackConfigurationOutput) SetHlsConfiguration(v *HlsConfiguration) *PutPlaybackConfigurationOutput {
	s.HlsConfiguration = v
	return s
}

// SetName sets the Name field's value.
func (s *PutPlaybackConfigurationOutput) SetName(v string) *PutPlaybackConfigurationOutput {
	s.Name = &v
	return s
}

// SetPlaybackEndpointPrefix sets the PlaybackEndpointPrefix field's value.
func (s *PutPlaybackConfigurationOutput) SetPlaybackEndpointPrefix(v string) *PutPlaybackConfigurationOutput {
	s.PlaybackEndpointPrefix = &v
	return s
}

// SetSessionInitializationEndpointPrefix sets the SessionInitializationEndpointPrefix field's value.
func (s *PutPlaybackConfigurationOutput) SetSessionInitializationEndpointPrefix(v string) *PutPlaybackConfigurationOutput {
	s.SessionInitializationEndpointPrefix = &v
	return s
}

// SetSlateAdUrl sets the SlateAdUrl field's value.
func (s *PutPlaybackConfigurationOutput) SetSlateAdUrl(v string) *PutPlaybackConfigurationOutput {
	s.SlateAdUrl = &v
	return s
}

// SetTranscodeProfileName sets the TranscodeProfileName field's value.
func (s *PutPlaybackConfigurationOutput) SetTranscodeProfileName(v string) *PutPlaybackConfigurationOutput {
	s.TranscodeProfileName = &v
	return s
}

// SetVideoContentSourceUrl sets the VideoContentSourceUrl field's value.
func (s *PutPlaybackConfigurationOutput) SetVideoContentSourceUrl(v string) *PutPlaybackConfigurationOutput {
	s.VideoContentSourceUrl = &v
	return s
}
