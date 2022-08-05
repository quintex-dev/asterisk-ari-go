# \ChannelsApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMoh**](ChannelsApi.md#AddMoh) | **Post** /channels/{channelId}/moh | Play music on hold to a channel.
[**Answer**](ChannelsApi.md#Answer) | **Post** /channels/{channelId}/answer | Answer a channel.
[**ContinueInDialplan**](ChannelsApi.md#ContinueInDialplan) | **Post** /channels/{channelId}/continue | Exit application; continue execution in the dialplan.
[**Createchannel**](ChannelsApi.md#Createchannel) | **Post** /channels/create | Create channel.
[**Deletemoh**](ChannelsApi.md#Deletemoh) | **Delete** /channels/{channelId}/moh | Stop playing music on hold to a channel.
[**Dial**](ChannelsApi.md#Dial) | **Post** /channels/{channelId}/dial | Dial a created channel.
[**ExternalMedia**](ChannelsApi.md#ExternalMedia) | **Post** /channels/externalMedia | Start an External Media session.
[**GetChannelVar**](ChannelsApi.md#GetChannelVar) | **Get** /channels/{channelId}/variable | Get the value of a channel variable or function.
[**Getchannel**](ChannelsApi.md#Getchannel) | **Get** /channels/{channelId} | Channel details.
[**Hangup**](ChannelsApi.md#Hangup) | **Delete** /channels/{channelId} | Delete (i.e. hangup) a channel.
[**Hold**](ChannelsApi.md#Hold) | **Post** /channels/{channelId}/hold | Hold a channel.
[**Listchannels**](ChannelsApi.md#Listchannels) | **Get** /channels | List all active channels in Asterisk.
[**Move**](ChannelsApi.md#Move) | **Post** /channels/{channelId}/move | Move the channel from one Stasis application to another.
[**Mute**](ChannelsApi.md#Mute) | **Post** /channels/{channelId}/mute | Mute a channel.
[**Originate**](ChannelsApi.md#Originate) | **Post** /channels | Create a new channel (originate).
[**OriginateWithId**](ChannelsApi.md#OriginateWithId) | **Post** /channels/{channelId} | Create a new channel (originate with id).
[**PlaySoundWithId**](ChannelsApi.md#PlaySoundWithId) | **Post** /channels/{channelId}/play/{playbackId} | Start playback of media and specify the playbackId.
[**Playsound**](ChannelsApi.md#Playsound) | **Post** /channels/{channelId}/play | Start playback of media.
[**Recordchannel**](ChannelsApi.md#Recordchannel) | **Post** /channels/{channelId}/record | Start a recording.
[**Redirect**](ChannelsApi.md#Redirect) | **Post** /channels/{channelId}/redirect | Redirect the channel to a different location.
[**Ring**](ChannelsApi.md#Ring) | **Post** /channels/{channelId}/ring | Indicate ringing to a channel.
[**RingStop**](ChannelsApi.md#RingStop) | **Delete** /channels/{channelId}/ring | Stop ringing indication on a channel if locally generated.
[**Rtpstatistics**](ChannelsApi.md#Rtpstatistics) | **Get** /channels/{channelId}/rtp_statistics | RTP stats on a channel.
[**SendDTMF**](ChannelsApi.md#SendDTMF) | **Post** /channels/{channelId}/dtmf | Send provided DTMF to a given channel.
[**SetChannelVar**](ChannelsApi.md#SetChannelVar) | **Post** /channels/{channelId}/variable | Set the value of a channel variable or function.
[**SnoopChannel**](ChannelsApi.md#SnoopChannel) | **Post** /channels/{channelId}/snoop | Start snooping.
[**SnoopChannelWithId**](ChannelsApi.md#SnoopChannelWithId) | **Post** /channels/{channelId}/snoop/{snoopId} | Start snooping.
[**StartSilence**](ChannelsApi.md#StartSilence) | **Post** /channels/{channelId}/silence | Play silence to a channel.
[**StopSilence**](ChannelsApi.md#StopSilence) | **Delete** /channels/{channelId}/silence | Stop playing silence to a channel.
[**Unhold**](ChannelsApi.md#Unhold) | **Delete** /channels/{channelId}/hold | Remove a channel from hold.
[**Unmute**](ChannelsApi.md#Unmute) | **Delete** /channels/{channelId}/mute | Unmute a channel.


# **AddMoh**
> AddMoh(ctx, channelId, optional)
Play music on hold to a channel.

Using media operations such as /play on a channel playing MOH in this manner will suspend MOH without resuming automatically. If continuing music on hold is desired, the stasis application must reinitiate music on hold.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
 **optional** | ***ChannelsApiAddMohOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiAddMohOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mohClass** | **optional.String**| Music on hold class to use | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Answer**
> Answer(ctx, channelId)
Answer a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ContinueInDialplan**
> ContinueInDialplan(ctx, channelId, optional)
Exit application; continue execution in the dialplan.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
 **optional** | ***ChannelsApiContinueInDialplanOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiContinueInDialplanOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **context** | **optional.String**| The context to continue to. | 
 **extension** | **optional.String**| The extension to continue to. | 
 **priority** | **optional.Int32**| The priority to continue to. | 
 **label** | **optional.String**| The label to continue to - will supersede &#39;priority&#39; if both are provided. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Createchannel**
> Channel Createchannel(ctx, endpoint, app, optional)
Create channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpoint** | **string**| Endpoint for channel communication | 
  **app** | **string**| Stasis Application to place channel into | 
 **optional** | ***ChannelsApiCreatechannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiCreatechannelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appArgs** | **optional.String**| The application arguments to pass to the Stasis application provided by &#39;app&#39;. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;. | 
 **channelId** | **optional.String**| The unique id to assign the channel on creation. | 
 **otherChannelId** | **optional.String**| The unique id to assign the second channel when using local channels. | 
 **originator** | **optional.String**| Unique ID of the calling channel | 
 **formats** | **optional.String**| The format name capability list to use if originator is not specified. Ex. \&quot;ulaw,slin16\&quot;.  Format names can be found with \&quot;core show codecs\&quot;. | 
 **variables** | [**optional.Interface of Containers**](Containers.md)| The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } } | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Deletemoh**
> Deletemoh(ctx, channelId)
Stop playing music on hold to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Dial**
> Dial(ctx, channelId, optional)
Dial a created channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
 **optional** | ***ChannelsApiDialOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiDialOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **caller** | **optional.String**| Channel ID of caller | 
 **timeout** | **optional.Int32**| Dial timeout | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ExternalMedia**
> Channel ExternalMedia(ctx, app, externalHost, format, optional)
Start an External Media session.

Create a channel to an External Media source/sink.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | **string**| Stasis Application to place channel into | 
  **externalHost** | **string**| Hostname/ip:port of external host | 
  **format** | **string**| Format to encode audio in | 
 **optional** | ***ChannelsApiExternalMediaOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiExternalMediaOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **channelId** | **optional.String**| The unique id to assign the channel on creation. | 
 **variables** | [**optional.Interface of Containers**](Containers.md)| The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } } | 
 **encapsulation** | **optional.String**| Payload encapsulation protocol | [default to rtp]
 **transport** | **optional.String**| Transport protocol | [default to udp]
 **connectionType** | **optional.String**| Connection type (client/server) | [default to client]
 **direction** | **optional.String**| External media direction | [default to both]
 **data** | **optional.String**| An arbitrary data field | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetChannelVar**
> Variable GetChannelVar(ctx, channelId, variable)
Get the value of a channel variable or function.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **variable** | **string**| The channel variable or function to get | 

### Return type

[**Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getchannel**
> Channel Getchannel(ctx, channelId)
Channel details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Hangup**
> Hangup(ctx, channelId, optional)
Delete (i.e. hangup) a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
 **optional** | ***ChannelsApiHangupOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiHangupOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **reasonCode** | **optional.String**| The reason code for hanging up the channel for detail use. Mutually exclusive with &#39;reason&#39;. See detail hangup codes at here. https://wiki.asterisk.org/wiki/display/AST/Hangup+Cause+Mappings | 
 **reason** | **optional.String**| Reason for hanging up the channel for simple use. Mutually exclusive with &#39;reason_code&#39;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Hold**
> Hold(ctx, channelId)
Hold a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Listchannels**
> []Channel Listchannels(ctx, )
List all active channels in Asterisk.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Move**
> Move(ctx, channelId, app, optional)
Move the channel from one Stasis application to another.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **app** | **string**| The channel will be passed to this Stasis application. | 
 **optional** | ***ChannelsApiMoveOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiMoveOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **appArgs** | **optional.String**| The application arguments to pass to the Stasis application provided by &#39;app&#39;. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Mute**
> Mute(ctx, channelId, optional)
Mute a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
 **optional** | ***ChannelsApiMuteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiMuteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **optional.String**| Direction in which to mute audio | [default to both]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Originate**
> Channel Originate(ctx, endpoint, optional)
Create a new channel (originate).

The new channel is created immediately and a snapshot of it returned. If a Stasis application is provided it will be automatically subscribed to the originated channel for further events and updates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **endpoint** | **string**| Endpoint to call. | 
 **optional** | ***ChannelsApiOriginateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiOriginateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **extension** | **optional.String**| The extension to dial after the endpoint answers. Mutually exclusive with &#39;app&#39;. | 
 **context** | **optional.String**| The context to dial after the endpoint answers. If omitted, uses &#39;default&#39;. Mutually exclusive with &#39;app&#39;. | 
 **priority** | **optional.Int64**| The priority to dial after the endpoint answers. If omitted, uses 1. Mutually exclusive with &#39;app&#39;. | 
 **label** | **optional.String**| The label to dial after the endpoint answers. Will supersede &#39;priority&#39; if provided. Mutually exclusive with &#39;app&#39;. | 
 **app** | **optional.String**| The application that is subscribed to the originated channel. When the channel is answered, it will be passed to this Stasis application. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;. | 
 **appArgs** | **optional.String**| The application arguments to pass to the Stasis application provided by &#39;app&#39;. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;. | 
 **callerId** | **optional.String**| CallerID to use when dialing the endpoint or extension. | 
 **timeout** | **optional.Int32**| Timeout (in seconds) before giving up dialing, or -1 for no timeout. | [default to 30]
 **variables** | [**optional.Interface of Containers**](Containers.md)| The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } } | 
 **channelId** | **optional.String**| The unique id to assign the channel on creation. | 
 **otherChannelId** | **optional.String**| The unique id to assign the second channel when using local channels. | 
 **originator** | **optional.String**| The unique id of the channel which is originating this one. | 
 **formats** | **optional.String**| The format name capability list to use if originator is not specified. Ex. \&quot;ulaw,slin16\&quot;.  Format names can be found with \&quot;core show codecs\&quot;. | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **OriginateWithId**
> Channel OriginateWithId(ctx, channelId, endpoint, optional)
Create a new channel (originate with id).

The new channel is created immediately and a snapshot of it returned. If a Stasis application is provided it will be automatically subscribed to the originated channel for further events and updates.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| The unique id to assign the channel on creation. | 
  **endpoint** | **string**| Endpoint to call. | 
 **optional** | ***ChannelsApiOriginateWithIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiOriginateWithIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **extension** | **optional.String**| The extension to dial after the endpoint answers. Mutually exclusive with &#39;app&#39;. | 
 **context** | **optional.String**| The context to dial after the endpoint answers. If omitted, uses &#39;default&#39;. Mutually exclusive with &#39;app&#39;. | 
 **priority** | **optional.Int64**| The priority to dial after the endpoint answers. If omitted, uses 1. Mutually exclusive with &#39;app&#39;. | 
 **label** | **optional.String**| The label to dial after the endpoint answers. Will supersede &#39;priority&#39; if provided. Mutually exclusive with &#39;app&#39;. | 
 **app** | **optional.String**| The application that is subscribed to the originated channel. When the channel is answered, it will be passed to this Stasis application. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;. | 
 **appArgs** | **optional.String**| The application arguments to pass to the Stasis application provided by &#39;app&#39;. Mutually exclusive with &#39;context&#39;, &#39;extension&#39;, &#39;priority&#39;, and &#39;label&#39;. | 
 **callerId** | **optional.String**| CallerID to use when dialing the endpoint or extension. | 
 **timeout** | **optional.Int32**| Timeout (in seconds) before giving up dialing, or -1 for no timeout. | [default to 30]
 **variables** | [**optional.Interface of Containers**](Containers.md)| The \&quot;variables\&quot; key in the body object holds variable key/value pairs to set on the channel on creation. Other keys in the body object are interpreted as query parameters. Ex. { \&quot;endpoint\&quot;: \&quot;SIP/Alice\&quot;, \&quot;variables\&quot;: { \&quot;CALLERID(name)\&quot;: \&quot;Alice\&quot; } } | 
 **otherChannelId** | **optional.String**| The unique id to assign the second channel when using local channels. | 
 **originator** | **optional.String**| The unique id of the channel which is originating this one. | 
 **formats** | **optional.String**| The format name capability list to use if originator is not specified. Ex. \&quot;ulaw,slin16\&quot;.  Format names can be found with \&quot;core show codecs\&quot;. | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlaySoundWithId**
> Playback PlaySoundWithId(ctx, channelId, playbackId, media, optional)
Start playback of media and specify the playbackId.

The media URI may be any of a number of URI's. Currently sound:, recording:, number:, digits:, characters:, and tone: URI's are supported. This operation creates a playback resource that can be used to control the playback of media (pause, rewind, fast forward, etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **playbackId** | **string**| Playback ID. | 
  **media** | [**[]string**](string.md)| Media URIs to play. | 
 **optional** | ***ChannelsApiPlaySoundWithIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiPlaySoundWithIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lang** | **optional.String**| For sounds, selects language for sound. | 
 **offsetms** | **optional.Int32**| Number of milliseconds to skip before playing. Only applies to the first URI if multiple media URIs are specified. | 
 **skipms** | **optional.Int32**| Number of milliseconds to skip for forward/reverse operations. | [default to 3000]

### Return type

[**Playback**](Playback.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Playsound**
> Playback Playsound(ctx, channelId, media, optional)
Start playback of media.

The media URI may be any of a number of URI's. Currently sound:, recording:, number:, digits:, characters:, and tone: URI's are supported. This operation creates a playback resource that can be used to control the playback of media (pause, rewind, fast forward, etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **media** | [**[]string**](string.md)| Media URIs to play. | 
 **optional** | ***ChannelsApiPlaysoundOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiPlaysoundOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lang** | **optional.String**| For sounds, selects language for sound. | 
 **offsetms** | **optional.Int32**| Number of milliseconds to skip before playing. Only applies to the first URI if multiple media URIs are specified. | 
 **skipms** | **optional.Int32**| Number of milliseconds to skip for forward/reverse operations. | [default to 3000]
 **playbackId** | **optional.String**| Playback ID. | 

### Return type

[**Playback**](Playback.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Recordchannel**
> LiveRecording Recordchannel(ctx, channelId, name, format, optional)
Start a recording.

Record audio from a channel. Note that this will not capture audio sent to the channel. The bridge itself has a record feature if that's what you want.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **name** | **string**| Recording&#39;s filename | 
  **format** | **string**| Format to encode audio in | 
 **optional** | ***ChannelsApiRecordchannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiRecordchannelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxDurationSeconds** | **optional.Int32**| Maximum duration of the recording, in seconds. 0 for no limit | [default to 0]
 **maxSilenceSeconds** | **optional.Int32**| Maximum duration of silence, in seconds. 0 for no limit | [default to 0]
 **ifExists** | **optional.String**| Action to take if a recording with the same name already exists. | [default to fail]
 **beep** | **optional.Bool**| Play beep when recording begins | [default to false]
 **terminateOn** | **optional.String**| DTMF input to terminate recording | [default to none]

### Return type

[**LiveRecording**](LiveRecording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Redirect**
> Redirect(ctx, channelId, endpoint)
Redirect the channel to a different location.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **endpoint** | **string**| The endpoint to redirect the channel to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ring**
> Ring(ctx, channelId)
Indicate ringing to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RingStop**
> RingStop(ctx, channelId)
Stop ringing indication on a channel if locally generated.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Rtpstatistics**
> RtPstat Rtpstatistics(ctx, channelId)
RTP stats on a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

[**RtPstat**](RTPstat.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendDTMF**
> SendDTMF(ctx, channelId, optional)
Send provided DTMF to a given channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
 **optional** | ***ChannelsApiSendDTMFOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiSendDTMFOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dtmf** | **optional.String**| DTMF To send. | 
 **before** | **optional.Int32**| Amount of time to wait before DTMF digits (specified in milliseconds) start. | [default to 0]
 **between** | **optional.Int32**| Amount of time in between DTMF digits (specified in milliseconds). | [default to 100]
 **duration** | **optional.Int32**| Length of each DTMF digit (specified in milliseconds). | [default to 100]
 **after** | **optional.Int32**| Amount of time to wait after DTMF digits (specified in milliseconds) end. | [default to 0]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetChannelVar**
> SetChannelVar(ctx, channelId, variable, optional)
Set the value of a channel variable or function.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **variable** | **string**| The channel variable or function to set | 
 **optional** | ***ChannelsApiSetChannelVarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiSetChannelVarOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **optional.String**| The value to set the variable to | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnoopChannel**
> Channel SnoopChannel(ctx, channelId, app, optional)
Start snooping.

Snoop (spy/whisper) on a specific channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **app** | **string**| Application the snooping channel is placed into | 
 **optional** | ***ChannelsApiSnoopChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiSnoopChannelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **spy** | **optional.String**| Direction of audio to spy on | [default to none]
 **whisper** | **optional.String**| Direction of audio to whisper into | [default to none]
 **appArgs** | **optional.String**| The application arguments to pass to the Stasis application | 
 **snoopId** | **optional.String**| Unique ID to assign to snooping channel | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SnoopChannelWithId**
> Channel SnoopChannelWithId(ctx, channelId, snoopId, app, optional)
Start snooping.

Snoop (spy/whisper) on a specific channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
  **snoopId** | **string**| Unique ID to assign to snooping channel | 
  **app** | **string**| Application the snooping channel is placed into | 
 **optional** | ***ChannelsApiSnoopChannelWithIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiSnoopChannelWithIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **spy** | **optional.String**| Direction of audio to spy on | [default to none]
 **whisper** | **optional.String**| Direction of audio to whisper into | [default to none]
 **appArgs** | **optional.String**| The application arguments to pass to the Stasis application | 

### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartSilence**
> StartSilence(ctx, channelId)
Play silence to a channel.

Using media operations such as /play on a channel playing silence in this manner will suspend silence without resuming automatically.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopSilence**
> StopSilence(ctx, channelId)
Stop playing silence to a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unhold**
> Unhold(ctx, channelId)
Remove a channel from hold.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unmute**
> Unmute(ctx, channelId, optional)
Unmute a channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **channelId** | **string**| Channel&#39;s id | 
 **optional** | ***ChannelsApiUnmuteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ChannelsApiUnmuteOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **direction** | **optional.String**| Direction in which to unmute audio | [default to both]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

