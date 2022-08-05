/*
 * Copyright (C) 2022 Quintex Software Solutions Pvt. Ltd. - All Rights Reserved.
 *
 * You may use, distribute and modify this code under the terms of the Apache
 * License Version 2.0. You should have received a copy of the license with this file.
 * If not, please write to : opensource@quintexsoftware.com
 *
 */

package asterisk_ari_go

import (
	"context"
	"fmt"
	"github.com/antihax/optional"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

type ChannelsApiService service

/*
ChannelsApiService Play music on hold to a channel.
Using media operations such as /play on a channel playing MOH in this manner will suspend MOH without resuming automatically. If continuing music on hold is desired, the stasis application must reinitiate music on hold.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param optional nil or *ChannelsApiAddMohOpts - Optional Parameters:
     * @param "MohClass" (optional.String) -  Music on hold class to use


*/

type ChannelsApiAddMohOpts struct {
	MohClass optional.String
}

func (a *ChannelsApiService) AddMoh(ctx context.Context, channelId string, localVarOptionals *ChannelsApiAddMohOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/moh"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.MohClass.IsSet() {
		localVarQueryParams.Add("mohClass", parameterToString(localVarOptionals.MohClass.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Answer a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) Answer(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/answer"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Exit application; continue execution in the dialplan.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param optional nil or *ChannelsApiContinueInDialplanOpts - Optional Parameters:
     * @param "Context" (optional.String) -  The context to continue to.
     * @param "Extension" (optional.String) -  The extension to continue to.
     * @param "Priority" (optional.Int32) -  The priority to continue to.
     * @param "Label" (optional.String) -  The label to continue to - will supersede &#39;priority&#39; if both are provided.


*/

type ChannelsApiContinueInDialplanOpts struct {
	Context   optional.String
	Extension optional.String
	Priority  optional.Int32
	Label     optional.String
}

func (a *ChannelsApiService) ContinueInDialplan(ctx context.Context, channelId string, localVarOptionals *ChannelsApiContinueInDialplanOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/continue"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Context.IsSet() {
		localVarQueryParams.Add("context", parameterToString(localVarOptionals.Context.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Extension.IsSet() {
		localVarQueryParams.Add("extension", parameterToString(localVarOptionals.Extension.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Priority.IsSet() {
		localVarQueryParams.Add("priority", parameterToString(localVarOptionals.Priority.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Label.IsSet() {
		localVarQueryParams.Add("label", parameterToString(localVarOptionals.Label.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Create channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param endpoint Endpoint for channel communication
 * @param app Stasis Application to place channel into
 * @param optional nil or *ChannelsApiCreatechannelOpts - Optional Parameters:
     * @param "AppArgs" (optional.String) -  The application arguments to pass to the Stasis application provided by &#39;app&#39;. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;.
     * @param "ChannelId" (optional.String) -  The unique id to assign the channel on creation.
     * @param "OtherChannelId" (optional.String) -  The unique id to assign the second channel when using local channels.
     * @param "Originator" (optional.String) -  Unique ID of the calling channel
     * @param "Formats" (optional.String) -  The format name capability list to use if originator is not specified. Ex. \&quot;ulaw,slin16\&quot;.  Format names can be found with \&quot;core show codecs\&quot;.
     * @param "Variables" (optional.Interface of Containers) -  The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } }

@return Channel
*/

type ChannelsApiCreatechannelOpts struct {
	AppArgs        optional.String
	ChannelId      optional.String
	OtherChannelId optional.String
	Originator     optional.String
	Formats        optional.String
	Variables      optional.Interface
}

func (a *ChannelsApiService) Createchannel(ctx context.Context, endpoint string, app string, localVarOptionals *ChannelsApiCreatechannelOpts) (Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/create"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("endpoint", parameterToString(endpoint, ""))
	localVarQueryParams.Add("app", parameterToString(app, ""))
	if localVarOptionals != nil && localVarOptionals.AppArgs.IsSet() {
		localVarQueryParams.Add("appArgs", parameterToString(localVarOptionals.AppArgs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChannelId.IsSet() {
		localVarQueryParams.Add("channelId", parameterToString(localVarOptionals.ChannelId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OtherChannelId.IsSet() {
		localVarQueryParams.Add("otherChannelId", parameterToString(localVarOptionals.OtherChannelId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Originator.IsSet() {
		localVarQueryParams.Add("originator", parameterToString(localVarOptionals.Originator.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Formats.IsSet() {
		localVarQueryParams.Add("formats", parameterToString(localVarOptionals.Formats.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Variables.IsSet() {

		localVarOptionalVariables, localVarOptionalVariablesok := localVarOptionals.Variables.Value().(Containers)
		if !localVarOptionalVariablesok {
			return localVarReturnValue, nil, reportError("variables should be Containers")
		}
		localVarPostBody = &localVarOptionalVariables
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Stop playing music on hold to a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) Deletemoh(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/moh"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Dial a created channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param optional nil or *ChannelsApiDialOpts - Optional Parameters:
     * @param "Caller" (optional.String) -  Channel ID of caller
     * @param "Timeout" (optional.Int32) -  Dial timeout


*/

type ChannelsApiDialOpts struct {
	Caller  optional.String
	Timeout optional.Int32
}

func (a *ChannelsApiService) Dial(ctx context.Context, channelId string, localVarOptionals *ChannelsApiDialOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/dial"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Caller.IsSet() {
		localVarQueryParams.Add("caller", parameterToString(localVarOptionals.Caller.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timeout.IsSet() {
		localVarQueryParams.Add("timeout", parameterToString(localVarOptionals.Timeout.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Start an External Media session.
Create a channel to an External Media source/sink.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param app Stasis Application to place channel into
 * @param externalHost Hostname/ip:port of external host
 * @param format Format to encode audio in
 * @param optional nil or *ChannelsApiExternalMediaOpts - Optional Parameters:
     * @param "ChannelId" (optional.String) -  The unique id to assign the channel on creation.
     * @param "Variables" (optional.Interface of Containers) -  The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } }
     * @param "Encapsulation" (optional.String) -  Payload encapsulation protocol
     * @param "Transport" (optional.String) -  Transport protocol
     * @param "ConnectionType" (optional.String) -  Connection type (client/server)
     * @param "Direction" (optional.String) -  External media direction
     * @param "Data" (optional.String) -  An arbitrary data field

@return Channel
*/

type ChannelsApiExternalMediaOpts struct {
	ChannelId      optional.String
	Variables      optional.Interface
	Encapsulation  optional.String
	Transport      optional.String
	ConnectionType optional.String
	Direction      optional.String
	Data           optional.String
}

func (a *ChannelsApiService) ExternalMedia(ctx context.Context, app string, externalHost string, format string, localVarOptionals *ChannelsApiExternalMediaOpts) (Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/externalMedia"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ChannelId.IsSet() {
		localVarQueryParams.Add("channelId", parameterToString(localVarOptionals.ChannelId.Value(), ""))
	}
	localVarQueryParams.Add("app", parameterToString(app, ""))
	localVarQueryParams.Add("external_host", parameterToString(externalHost, ""))
	if localVarOptionals != nil && localVarOptionals.Encapsulation.IsSet() {
		localVarQueryParams.Add("encapsulation", parameterToString(localVarOptionals.Encapsulation.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Transport.IsSet() {
		localVarQueryParams.Add("transport", parameterToString(localVarOptionals.Transport.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ConnectionType.IsSet() {
		localVarQueryParams.Add("connection_type", parameterToString(localVarOptionals.ConnectionType.Value(), ""))
	}
	localVarQueryParams.Add("format", parameterToString(format, ""))
	if localVarOptionals != nil && localVarOptionals.Direction.IsSet() {
		localVarQueryParams.Add("direction", parameterToString(localVarOptionals.Direction.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Data.IsSet() {
		localVarQueryParams.Add("data", parameterToString(localVarOptionals.Data.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Variables.IsSet() {

		localVarOptionalVariables, localVarOptionalVariablesok := localVarOptionals.Variables.Value().(Containers)
		if !localVarOptionalVariablesok {
			return localVarReturnValue, nil, reportError("variables should be Containers")
		}
		localVarPostBody = &localVarOptionalVariables
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Get the value of a channel variable or function.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param variable The channel variable or function to get

@return Variable
*/
func (a *ChannelsApiService) GetChannelVar(ctx context.Context, channelId string, variable string) (Variable, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Variable
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/variable"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("variable", parameterToString(variable, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Variable
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Channel details.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id

@return Channel
*/
func (a *ChannelsApiService) Getchannel(ctx context.Context, channelId string) (Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Delete (i.e. hangup) a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param optional nil or *ChannelsApiHangupOpts - Optional Parameters:
     * @param "ReasonCode" (optional.String) -  The reason code for hanging up the channel for detail use. Mutually exclusive with &#39;reason&#39;. See detail hangup codes at here. https://wiki.asterisk.org/wiki/display/AST/Hangup+Cause+Mappings
     * @param "Reason" (optional.String) -  Reason for hanging up the channel for simple use. Mutually exclusive with &#39;reason_code&#39;.


*/

type ChannelsApiHangupOpts struct {
	ReasonCode optional.String
	Reason     optional.String
}

func (a *ChannelsApiService) Hangup(ctx context.Context, channelId string, localVarOptionals *ChannelsApiHangupOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.ReasonCode.IsSet() {
		localVarQueryParams.Add("reason_code", parameterToString(localVarOptionals.ReasonCode.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Reason.IsSet() {
		localVarQueryParams.Add("reason", parameterToString(localVarOptionals.Reason.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Hold a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) Hold(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/hold"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService List all active channels in Asterisk.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().

@return []Channel
*/
func (a *ChannelsApiService) Listchannels(ctx context.Context) ([]Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue []Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v []Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Move the channel from one Stasis application to another.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param app The channel will be passed to this Stasis application.
 * @param optional nil or *ChannelsApiMoveOpts - Optional Parameters:
     * @param "AppArgs" (optional.String) -  The application arguments to pass to the Stasis application provided by &#39;app&#39;.


*/

type ChannelsApiMoveOpts struct {
	AppArgs optional.String
}

func (a *ChannelsApiService) Move(ctx context.Context, channelId string, app string, localVarOptionals *ChannelsApiMoveOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/move"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("app", parameterToString(app, ""))
	if localVarOptionals != nil && localVarOptionals.AppArgs.IsSet() {
		localVarQueryParams.Add("appArgs", parameterToString(localVarOptionals.AppArgs.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Mute a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param optional nil or *ChannelsApiMuteOpts - Optional Parameters:
     * @param "Direction" (optional.String) -  Direction in which to mute audio


*/

type ChannelsApiMuteOpts struct {
	Direction optional.String
}

func (a *ChannelsApiService) Mute(ctx context.Context, channelId string, localVarOptionals *ChannelsApiMuteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/mute"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Direction.IsSet() {
		localVarQueryParams.Add("direction", parameterToString(localVarOptionals.Direction.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Create a new channel (originate).
The new channel is created immediately and a snapshot of it returned. If a Stasis application is provided it will be automatically subscribed to the originated channel for further events and updates.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param endpoint Endpoint to call.
 * @param optional nil or *ChannelsApiOriginateOpts - Optional Parameters:
     * @param "Extension" (optional.String) -  The extension to dial after the endpoint answers. Mutually exclusive with &#39;app&#39;.
     * @param "Context" (optional.String) -  The context to dial after the endpoint answers. If omitted, uses &#39;default&#39;. Mutually exclusive with &#39;app&#39;.
     * @param "Priority" (optional.Int64) -  The priority to dial after the endpoint answers. If omitted, uses 1. Mutually exclusive with &#39;app&#39;.
     * @param "Label" (optional.String) -  The label to dial after the endpoint answers. Will supersede &#39;priority&#39; if provided. Mutually exclusive with &#39;app&#39;.
     * @param "App" (optional.String) -  The application that is subscribed to the originated channel. When the channel is answered, it will be passed to this Stasis application. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;.
     * @param "AppArgs" (optional.String) -  The application arguments to pass to the Stasis application provided by &#39;app&#39;. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;.
     * @param "CallerId" (optional.String) -  CallerID to use when dialing the endpoint or extension.
     * @param "Timeout" (optional.Int32) -  Timeout (in seconds) before giving up dialing, or -1 for no timeout.
     * @param "Variables" (optional.Interface of Containers) -  The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } }
     * @param "ChannelId" (optional.String) -  The unique id to assign the channel on creation.
     * @param "OtherChannelId" (optional.String) -  The unique id to assign the second channel when using local channels.
     * @param "Originator" (optional.String) -  The unique id of the channel which is originating this one.
     * @param "Formats" (optional.String) -  The format name capability list to use if originator is not specified. Ex. \&quot;ulaw,slin16\&quot;.  Format names can be found with \&quot;core show codecs\&quot;.

@return Channel
*/

type ChannelsApiOriginateOpts struct {
	Extension      optional.String
	Context        optional.String
	Priority       optional.Int64
	Label          optional.String
	App            optional.String
	AppArgs        optional.String
	CallerId       optional.String
	Timeout        optional.Int32
	Variables      optional.Interface
	ChannelId      optional.String
	OtherChannelId optional.String
	Originator     optional.String
	Formats        optional.String
}

func (a *ChannelsApiService) Originate(ctx context.Context, endpoint string, localVarOptionals *ChannelsApiOriginateOpts) (Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("endpoint", parameterToString(endpoint, ""))
	if localVarOptionals != nil && localVarOptionals.Extension.IsSet() {
		localVarQueryParams.Add("extension", parameterToString(localVarOptionals.Extension.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Context.IsSet() {
		localVarQueryParams.Add("context", parameterToString(localVarOptionals.Context.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Priority.IsSet() {
		localVarQueryParams.Add("priority", parameterToString(localVarOptionals.Priority.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Label.IsSet() {
		localVarQueryParams.Add("label", parameterToString(localVarOptionals.Label.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.App.IsSet() {
		localVarQueryParams.Add("app", parameterToString(localVarOptionals.App.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AppArgs.IsSet() {
		localVarQueryParams.Add("appArgs", parameterToString(localVarOptionals.AppArgs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CallerId.IsSet() {
		localVarQueryParams.Add("callerId", parameterToString(localVarOptionals.CallerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timeout.IsSet() {
		localVarQueryParams.Add("timeout", parameterToString(localVarOptionals.Timeout.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.ChannelId.IsSet() {
		localVarQueryParams.Add("channelId", parameterToString(localVarOptionals.ChannelId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OtherChannelId.IsSet() {
		localVarQueryParams.Add("otherChannelId", parameterToString(localVarOptionals.OtherChannelId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Originator.IsSet() {
		localVarQueryParams.Add("originator", parameterToString(localVarOptionals.Originator.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Formats.IsSet() {
		localVarQueryParams.Add("formats", parameterToString(localVarOptionals.Formats.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Variables.IsSet() {

		localVarOptionalVariables, localVarOptionalVariablesok := localVarOptionals.Variables.Value().(Containers)
		if !localVarOptionalVariablesok {
			return localVarReturnValue, nil, reportError("variables should be Containers")
		}
		localVarPostBody = &localVarOptionalVariables
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Create a new channel (originate with id).
The new channel is created immediately and a snapshot of it returned. If a Stasis application is provided it will be automatically subscribed to the originated channel for further events and updates.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId The unique id to assign the channel on creation.
 * @param endpoint Endpoint to call.
 * @param optional nil or *ChannelsApiOriginateWithIdOpts - Optional Parameters:
     * @param "Extension" (optional.String) -  The extension to dial after the endpoint answers. Mutually exclusive with &#39;app&#39;.
     * @param "Context" (optional.String) -  The context to dial after the endpoint answers. If omitted, uses &#39;default&#39;. Mutually exclusive with &#39;app&#39;.
     * @param "Priority" (optional.Int64) -  The priority to dial after the endpoint answers. If omitted, uses 1. Mutually exclusive with &#39;app&#39;.
     * @param "Label" (optional.String) -  The label to dial after the endpoint answers. Will supersede &#39;priority&#39; if provided. Mutually exclusive with &#39;app&#39;.
     * @param "App" (optional.String) -  The application that is subscribed to the originated channel. When the channel is answered, it will be passed to this Stasis application. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;.
     * @param "AppArgs" (optional.String) -  The application arguments to pass to the Stasis application provided by &#39;app&#39;. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;.
     * @param "CallerId" (optional.String) -  CallerID to use when dialing the endpoint or extension.
     * @param "Timeout" (optional.Int32) -  Timeout (in seconds) before giving up dialing, or -1 for no timeout.
     * @param "Variables" (optional.Interface of Containers) -  The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } }
     * @param "OtherChannelId" (optional.String) -  The unique id to assign the second channel when using local channels.
     * @param "Originator" (optional.String) -  The unique id of the channel which is originating this one.
     * @param "Formats" (optional.String) -  The format name capability list to use if originator is not specified. Ex. \&quot;ulaw,slin16\&quot;.  Format names can be found with \&quot;core show codecs\&quot;.

@return Channel
*/

type ChannelsApiOriginateWithIdOpts struct {
	Extension      optional.String
	Context        optional.String
	Priority       optional.Int64
	Label          optional.String
	App            optional.String
	AppArgs        optional.String
	CallerId       optional.String
	Timeout        optional.Int32
	Variables      optional.Interface
	OtherChannelId optional.String
	Originator     optional.String
	Formats        optional.String
}

func (a *ChannelsApiService) OriginateWithId(ctx context.Context, channelId string, endpoint string, localVarOptionals *ChannelsApiOriginateWithIdOpts) (Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("endpoint", parameterToString(endpoint, ""))
	if localVarOptionals != nil && localVarOptionals.Extension.IsSet() {
		localVarQueryParams.Add("extension", parameterToString(localVarOptionals.Extension.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Context.IsSet() {
		localVarQueryParams.Add("context", parameterToString(localVarOptionals.Context.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Priority.IsSet() {
		localVarQueryParams.Add("priority", parameterToString(localVarOptionals.Priority.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Label.IsSet() {
		localVarQueryParams.Add("label", parameterToString(localVarOptionals.Label.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.App.IsSet() {
		localVarQueryParams.Add("app", parameterToString(localVarOptionals.App.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.AppArgs.IsSet() {
		localVarQueryParams.Add("appArgs", parameterToString(localVarOptionals.AppArgs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CallerId.IsSet() {
		localVarQueryParams.Add("callerId", parameterToString(localVarOptionals.CallerId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Timeout.IsSet() {
		localVarQueryParams.Add("timeout", parameterToString(localVarOptionals.Timeout.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.OtherChannelId.IsSet() {
		localVarQueryParams.Add("otherChannelId", parameterToString(localVarOptionals.OtherChannelId.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Originator.IsSet() {
		localVarQueryParams.Add("originator", parameterToString(localVarOptionals.Originator.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Formats.IsSet() {
		localVarQueryParams.Add("formats", parameterToString(localVarOptionals.Formats.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.Variables.IsSet() {

		localVarOptionalVariables, localVarOptionalVariablesok := localVarOptionals.Variables.Value().(Containers)
		if !localVarOptionalVariablesok {
			return localVarReturnValue, nil, reportError("variables should be Containers")
		}
		localVarPostBody = &localVarOptionalVariables
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Start playback of media and specify the playbackId.
The media URI may be any of a number of URI&#39;s. Currently sound:, recording:, number:, digits:, characters:, and tone: URI&#39;s are supported. This operation creates a playback resource that can be used to control the playback of media (pause, rewind, fast forward, etc.)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param playbackId Playback ID.
 * @param media Media URIs to play.
 * @param optional nil or *ChannelsApiPlaySoundWithIdOpts - Optional Parameters:
     * @param "Lang" (optional.String) -  For sounds, selects language for sound.
     * @param "Offsetms" (optional.Int32) -  Number of milliseconds to skip before playing. Only applies to the first URI if multiple media URIs are specified.
     * @param "Skipms" (optional.Int32) -  Number of milliseconds to skip for forward/reverse operations.

@return Playback
*/

type ChannelsApiPlaySoundWithIdOpts struct {
	Lang     optional.String
	Offsetms optional.Int32
	Skipms   optional.Int32
}

func (a *ChannelsApiService) PlaySoundWithId(ctx context.Context, channelId string, playbackId string, media []string, localVarOptionals *ChannelsApiPlaySoundWithIdOpts) (Playback, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Playback
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/play/{playbackId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"playbackId"+"}", fmt.Sprintf("%v", playbackId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("media", parameterToString(media, "multi"))
	if localVarOptionals != nil && localVarOptionals.Lang.IsSet() {
		localVarQueryParams.Add("lang", parameterToString(localVarOptionals.Lang.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offsetms.IsSet() {
		localVarQueryParams.Add("offsetms", parameterToString(localVarOptionals.Offsetms.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skipms.IsSet() {
		localVarQueryParams.Add("skipms", parameterToString(localVarOptionals.Skipms.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Playback
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Start playback of media.
The media URI may be any of a number of URI&#39;s. Currently sound:, recording:, number:, digits:, characters:, and tone: URI&#39;s are supported. This operation creates a playback resource that can be used to control the playback of media (pause, rewind, fast forward, etc.)
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param media Media URIs to play.
 * @param optional nil or *ChannelsApiPlaysoundOpts - Optional Parameters:
     * @param "Lang" (optional.String) -  For sounds, selects language for sound.
     * @param "Offsetms" (optional.Int32) -  Number of milliseconds to skip before playing. Only applies to the first URI if multiple media URIs are specified.
     * @param "Skipms" (optional.Int32) -  Number of milliseconds to skip for forward/reverse operations.
     * @param "PlaybackId" (optional.String) -  Playback ID.

@return Playback
*/

type ChannelsApiPlaysoundOpts struct {
	Lang       optional.String
	Offsetms   optional.Int32
	Skipms     optional.Int32
	PlaybackId optional.String
}

func (a *ChannelsApiService) Playsound(ctx context.Context, channelId string, media []string, localVarOptionals *ChannelsApiPlaysoundOpts) (Playback, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Playback
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/play"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("media", parameterToString(media, "multi"))
	if localVarOptionals != nil && localVarOptionals.Lang.IsSet() {
		localVarQueryParams.Add("lang", parameterToString(localVarOptionals.Lang.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Offsetms.IsSet() {
		localVarQueryParams.Add("offsetms", parameterToString(localVarOptionals.Offsetms.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Skipms.IsSet() {
		localVarQueryParams.Add("skipms", parameterToString(localVarOptionals.Skipms.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.PlaybackId.IsSet() {
		localVarQueryParams.Add("playbackId", parameterToString(localVarOptionals.PlaybackId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Playback
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Start a recording.
Record audio from a channel. Note that this will not capture audio sent to the channel. The bridge itself has a record feature if that&#39;s what you want.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param name Recording&#39;s filename
 * @param format Format to encode audio in
 * @param optional nil or *ChannelsApiRecordchannelOpts - Optional Parameters:
     * @param "MaxDurationSeconds" (optional.Int32) -  Maximum duration of the recording, in seconds. 0 for no limit
     * @param "MaxSilenceSeconds" (optional.Int32) -  Maximum duration of silence, in seconds. 0 for no limit
     * @param "IfExists" (optional.String) -  Action to take if a recording with the same name already exists.
     * @param "Beep" (optional.Bool) -  Play beep when recording begins
     * @param "TerminateOn" (optional.String) -  DTMF input to terminate recording

@return LiveRecording
*/

type ChannelsApiRecordchannelOpts struct {
	MaxDurationSeconds optional.Int32
	MaxSilenceSeconds  optional.Int32
	IfExists           optional.String
	Beep               optional.Bool
	TerminateOn        optional.String
}

func (a *ChannelsApiService) Recordchannel(ctx context.Context, channelId string, name string, format string, localVarOptionals *ChannelsApiRecordchannelOpts) (LiveRecording, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue LiveRecording
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/record"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("name", parameterToString(name, ""))
	localVarQueryParams.Add("format", parameterToString(format, ""))
	if localVarOptionals != nil && localVarOptionals.MaxDurationSeconds.IsSet() {
		localVarQueryParams.Add("maxDurationSeconds", parameterToString(localVarOptionals.MaxDurationSeconds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.MaxSilenceSeconds.IsSet() {
		localVarQueryParams.Add("maxSilenceSeconds", parameterToString(localVarOptionals.MaxSilenceSeconds.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.IfExists.IsSet() {
		localVarQueryParams.Add("ifExists", parameterToString(localVarOptionals.IfExists.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Beep.IsSet() {
		localVarQueryParams.Add("beep", parameterToString(localVarOptionals.Beep.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.TerminateOn.IsSet() {
		localVarQueryParams.Add("terminateOn", parameterToString(localVarOptionals.TerminateOn.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v LiveRecording
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Redirect the channel to a different location.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param endpoint The endpoint to redirect the channel to


*/
func (a *ChannelsApiService) Redirect(ctx context.Context, channelId string, endpoint string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/redirect"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("endpoint", parameterToString(endpoint, ""))
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Indicate ringing to a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) Ring(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/ring"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Stop ringing indication on a channel if locally generated.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) RingStop(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/ring"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService RTP stats on a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id

@return RtPstat
*/
func (a *ChannelsApiService) Rtpstatistics(ctx context.Context, channelId string) (RtPstat, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Get")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue RtPstat
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/rtp_statistics"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v RtPstat
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Send provided DTMF to a given channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param optional nil or *ChannelsApiSendDTMFOpts - Optional Parameters:
     * @param "Dtmf" (optional.String) -  DTMF To send.
     * @param "Before" (optional.Int32) -  Amount of time to wait before DTMF digits (specified in milliseconds) start.
     * @param "Between" (optional.Int32) -  Amount of time in between DTMF digits (specified in milliseconds).
     * @param "Duration" (optional.Int32) -  Length of each DTMF digit (specified in milliseconds).
     * @param "After" (optional.Int32) -  Amount of time to wait after DTMF digits (specified in milliseconds) end.


*/

type ChannelsApiSendDTMFOpts struct {
	Dtmf     optional.String
	Before   optional.Int32
	Between  optional.Int32
	Duration optional.Int32
	After    optional.Int32
}

func (a *ChannelsApiService) SendDTMF(ctx context.Context, channelId string, localVarOptionals *ChannelsApiSendDTMFOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/dtmf"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Dtmf.IsSet() {
		localVarQueryParams.Add("dtmf", parameterToString(localVarOptionals.Dtmf.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Before.IsSet() {
		localVarQueryParams.Add("before", parameterToString(localVarOptionals.Before.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Between.IsSet() {
		localVarQueryParams.Add("between", parameterToString(localVarOptionals.Between.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Duration.IsSet() {
		localVarQueryParams.Add("duration", parameterToString(localVarOptionals.Duration.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.After.IsSet() {
		localVarQueryParams.Add("after", parameterToString(localVarOptionals.After.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Set the value of a channel variable or function.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param variable The channel variable or function to set
 * @param optional nil or *ChannelsApiSetChannelVarOpts - Optional Parameters:
     * @param "Value" (optional.String) -  The value to set the variable to


*/

type ChannelsApiSetChannelVarOpts struct {
	Value optional.String
}

func (a *ChannelsApiService) SetChannelVar(ctx context.Context, channelId string, variable string, localVarOptionals *ChannelsApiSetChannelVarOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/variable"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	localVarQueryParams.Add("variable", parameterToString(variable, ""))
	if localVarOptionals != nil && localVarOptionals.Value.IsSet() {
		localVarQueryParams.Add("value", parameterToString(localVarOptionals.Value.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Start snooping.
Snoop (spy/whisper) on a specific channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param app Application the snooping channel is placed into
 * @param optional nil or *ChannelsApiSnoopChannelOpts - Optional Parameters:
     * @param "Spy" (optional.String) -  Direction of audio to spy on
     * @param "Whisper" (optional.String) -  Direction of audio to whisper into
     * @param "AppArgs" (optional.String) -  The application arguments to pass to the Stasis application
     * @param "SnoopId" (optional.String) -  Unique ID to assign to snooping channel

@return Channel
*/

type ChannelsApiSnoopChannelOpts struct {
	Spy     optional.String
	Whisper optional.String
	AppArgs optional.String
	SnoopId optional.String
}

func (a *ChannelsApiService) SnoopChannel(ctx context.Context, channelId string, app string, localVarOptionals *ChannelsApiSnoopChannelOpts) (Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/snoop"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Spy.IsSet() {
		localVarQueryParams.Add("spy", parameterToString(localVarOptionals.Spy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Whisper.IsSet() {
		localVarQueryParams.Add("whisper", parameterToString(localVarOptionals.Whisper.Value(), ""))
	}
	localVarQueryParams.Add("app", parameterToString(app, ""))
	if localVarOptionals != nil && localVarOptionals.AppArgs.IsSet() {
		localVarQueryParams.Add("appArgs", parameterToString(localVarOptionals.AppArgs.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.SnoopId.IsSet() {
		localVarQueryParams.Add("snoopId", parameterToString(localVarOptionals.SnoopId.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Start snooping.
Snoop (spy/whisper) on a specific channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param snoopId Unique ID to assign to snooping channel
 * @param app Application the snooping channel is placed into
 * @param optional nil or *ChannelsApiSnoopChannelWithIdOpts - Optional Parameters:
     * @param "Spy" (optional.String) -  Direction of audio to spy on
     * @param "Whisper" (optional.String) -  Direction of audio to whisper into
     * @param "AppArgs" (optional.String) -  The application arguments to pass to the Stasis application

@return Channel
*/

type ChannelsApiSnoopChannelWithIdOpts struct {
	Spy     optional.String
	Whisper optional.String
	AppArgs optional.String
}

func (a *ChannelsApiService) SnoopChannelWithId(ctx context.Context, channelId string, snoopId string, app string, localVarOptionals *ChannelsApiSnoopChannelWithIdOpts) (Channel, *http.Response, error) {
	var (
		localVarHttpMethod  = strings.ToUpper("Post")
		localVarPostBody    interface{}
		localVarFileName    string
		localVarFileBytes   []byte
		localVarReturnValue Channel
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/snoop/{snoopId}"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"snoopId"+"}", fmt.Sprintf("%v", snoopId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Spy.IsSet() {
		localVarQueryParams.Add("spy", parameterToString(localVarOptionals.Spy.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Whisper.IsSet() {
		localVarQueryParams.Add("whisper", parameterToString(localVarOptionals.Whisper.Value(), ""))
	}
	localVarQueryParams.Add("app", parameterToString(app, ""))
	if localVarOptionals != nil && localVarOptionals.AppArgs.IsSet() {
		localVarQueryParams.Add("appArgs", parameterToString(localVarOptionals.AppArgs.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		if localVarHttpResponse.StatusCode == 200 {
			var v Channel
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}

		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
ChannelsApiService Play silence to a channel.
Using media operations such as /play on a channel playing silence in this manner will suspend silence without resuming automatically.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) StartSilence(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/silence"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Stop playing silence to a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) StopSilence(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/silence"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Remove a channel from hold.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id


*/
func (a *ChannelsApiService) Unhold(ctx context.Context, channelId string) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/hold"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
ChannelsApiService Unmute a channel.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param channelId Channel&#39;s id
 * @param optional nil or *ChannelsApiUnmuteOpts - Optional Parameters:
     * @param "Direction" (optional.String) -  Direction in which to unmute audio


*/

type ChannelsApiUnmuteOpts struct {
	Direction optional.String
}

func (a *ChannelsApiService) Unmute(ctx context.Context, channelId string, localVarOptionals *ChannelsApiUnmuteOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/channels/{channelId}/mute"
	localVarPath = strings.Replace(localVarPath, "{"+"channelId"+"}", fmt.Sprintf("%v", channelId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Direction.IsSet() {
		localVarQueryParams.Add("direction", parameterToString(localVarOptionals.Direction.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericSwaggerError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}

		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
