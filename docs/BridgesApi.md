# \BridgesApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddChannel**](BridgesApi.md#AddChannel) | **Post** /bridges/{bridgeId}/addChannel | Add a channel to a bridge.
[**ClearVideoSource**](BridgesApi.md#ClearVideoSource) | **Delete** /bridges/{bridgeId}/videoSource | Removes any explicit video source in a multi-party mixing bridge. This operation has no effect on bridges with two or fewer participants. When no explicit video source is set, talk detection will be used to determine the active video stream.
[**Create**](BridgesApi.md#Create) | **Post** /bridges | Create a new bridge.
[**CreateWithId**](BridgesApi.md#CreateWithId) | **Post** /bridges/{bridgeId} | Create a new bridge or updates an existing one.
[**Destroy**](BridgesApi.md#Destroy) | **Delete** /bridges/{bridgeId} | Shut down a bridge.
[**Getbridge**](BridgesApi.md#Getbridge) | **Get** /bridges/{bridgeId} | Get bridge details.
[**Listbridges**](BridgesApi.md#Listbridges) | **Get** /bridges | List all active bridges in Asterisk.
[**Play**](BridgesApi.md#Play) | **Post** /bridges/{bridgeId}/play | Start playback of media on a bridge.
[**PlayWithId**](BridgesApi.md#PlayWithId) | **Post** /bridges/{bridgeId}/play/{playbackId} | Start playback of media on a bridge.
[**Record**](BridgesApi.md#Record) | **Post** /bridges/{bridgeId}/record | Start a recording.
[**RemoveChannel**](BridgesApi.md#RemoveChannel) | **Post** /bridges/{bridgeId}/removeChannel | Remove a channel from a bridge.
[**SetVideoSource**](BridgesApi.md#SetVideoSource) | **Post** /bridges/{bridgeId}/videoSource/{channelId} | Set a channel as the video source in a multi-party mixing bridge. This operation has no effect on bridges with two or fewer participants.
[**StartMoh**](BridgesApi.md#StartMoh) | **Post** /bridges/{bridgeId}/moh | Play music on hold to a bridge or change the MOH class that is playing.
[**StopMoh**](BridgesApi.md#StopMoh) | **Delete** /bridges/{bridgeId}/moh | Stop playing music on hold to a bridge.


# **AddChannel**
> AddChannel(ctx, bridgeId, channel, optional)
Add a channel to a bridge.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 
  **channel** | [**[]string**](string.md)| Ids of channels to add to bridge | 
 **optional** | ***BridgesApiAddChannelOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiAddChannelOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **role** | **optional.String**| Channel&#39;s role in the bridge | 
 **absorbDTMF** | **optional.Bool**| Absorb DTMF coming from this channel, preventing it to pass through to the bridge | [default to false]
 **mute** | **optional.Bool**| Mute audio from this channel, preventing it to pass through to the bridge | [default to false]
 **inhibitConnectedLineUpdates** | **optional.Bool**| Do not present the identity of the newly connected channel to other bridge members | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearVideoSource**
> ClearVideoSource(ctx, bridgeId)
Removes any explicit video source in a multi-party mixing bridge. This operation has no effect on bridges with two or fewer participants. When no explicit video source is set, talk detection will be used to determine the active video stream.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Create**
> Bridge Create(ctx, optional)
Create a new bridge.

This bridge persists until it has been shut down, or Asterisk has been shut down.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***BridgesApiCreateOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiCreateOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **optional.String**| Comma separated list of bridge type attributes (mixing, holding, dtmf_events, proxy_media, video_sfu). | 
 **bridgeId** | **optional.String**| Unique ID to give to the bridge being created. | 
 **name** | **optional.String**| Name to give to the bridge being created. | 

### Return type

[**Bridge**](Bridge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CreateWithId**
> Bridge CreateWithId(ctx, bridgeId, optional)
Create a new bridge or updates an existing one.

This bridge persists until it has been shut down, or Asterisk has been shut down.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Unique ID to give to the bridge being created. | 
 **optional** | ***BridgesApiCreateWithIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiCreateWithIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **type_** | **optional.String**| Comma separated list of bridge type attributes (mixing, holding, dtmf_events, proxy_media, video_sfu) to set. | 
 **name** | **optional.String**| Set the name of the bridge. | 

### Return type

[**Bridge**](Bridge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Destroy**
> Destroy(ctx, bridgeId)
Shut down a bridge.

If any channels are in this bridge, they will be removed and resume whatever they were doing beforehand.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getbridge**
> Bridge Getbridge(ctx, bridgeId)
Get bridge details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 

### Return type

[**Bridge**](Bridge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Listbridges**
> []Bridge Listbridges(ctx, )
List all active bridges in Asterisk.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Bridge**](Bridge.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Play**
> Playback Play(ctx, bridgeId, media, optional)
Start playback of media on a bridge.

The media URI may be any of a number of URI's. Currently sound:, recording:, number:, digits:, characters:, and tone: URI's are supported. This operation creates a playback resource that can be used to control the playback of media (pause, rewind, fast forward, etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 
  **media** | [**[]string**](string.md)| Media URIs to play. | 
 **optional** | ***BridgesApiPlayOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiPlayOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lang** | **optional.String**| For sounds, selects language for sound. | 
 **offsetms** | **optional.Int32**| Number of milliseconds to skip before playing. Only applies to the first URI if multiple media URIs are specified. | [default to 0]
 **skipms** | **optional.Int32**| Number of milliseconds to skip for forward/reverse operations. | [default to 3000]
 **playbackId** | **optional.String**| Playback Id. | 

### Return type

[**Playback**](Playback.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PlayWithId**
> Playback PlayWithId(ctx, bridgeId, playbackId, media, optional)
Start playback of media on a bridge.

The media URI may be any of a number of URI's. Currently sound:, recording:, number:, digits:, characters:, and tone: URI's are supported. This operation creates a playback resource that can be used to control the playback of media (pause, rewind, fast forward, etc.)

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 
  **playbackId** | **string**| Playback ID. | 
  **media** | [**[]string**](string.md)| Media URIs to play. | 
 **optional** | ***BridgesApiPlayWithIdOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiPlayWithIdOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **lang** | **optional.String**| For sounds, selects language for sound. | 
 **offsetms** | **optional.Int32**| Number of milliseconds to skip before playing. Only applies to the first URI if multiple media URIs are specified. | [default to 0]
 **skipms** | **optional.Int32**| Number of milliseconds to skip for forward/reverse operations. | [default to 3000]

### Return type

[**Playback**](Playback.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Record**
> LiveRecording Record(ctx, bridgeId, name, format, optional)
Start a recording.

This records the mixed audio from all channels participating in this bridge.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 
  **name** | **string**| Recording&#39;s filename | 
  **format** | **string**| Format to encode audio in | 
 **optional** | ***BridgesApiRecordOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiRecordOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **maxDurationSeconds** | **optional.Int32**| Maximum duration of the recording, in seconds. 0 for no limit. | [default to 0]
 **maxSilenceSeconds** | **optional.Int32**| Maximum duration of silence, in seconds. 0 for no limit. | [default to 0]
 **ifExists** | **optional.String**| Action to take if a recording with the same name already exists. | [default to fail]
 **beep** | **optional.Bool**| Play beep when recording begins | [default to false]
 **terminateOn** | **optional.String**| DTMF input to terminate recording. | [default to none]

### Return type

[**LiveRecording**](LiveRecording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveChannel**
> RemoveChannel(ctx, bridgeId, channel)
Remove a channel from a bridge.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 
  **channel** | [**[]string**](string.md)| Ids of channels to remove from bridge | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetVideoSource**
> SetVideoSource(ctx, bridgeId, channelId)
Set a channel as the video source in a multi-party mixing bridge. This operation has no effect on bridges with two or fewer participants.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 
  **channelId** | **string**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StartMoh**
> StartMoh(ctx, bridgeId, optional)
Play music on hold to a bridge or change the MOH class that is playing.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 
 **optional** | ***BridgesApiStartMohOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a BridgesApiStartMohOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **mohClass** | **optional.String**| Channel&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StopMoh**
> StopMoh(ctx, bridgeId)
Stop playing music on hold to a bridge.

This will only stop music on hold being played via POST bridges/{bridgeId}/moh.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **bridgeId** | **string**| Bridge&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

