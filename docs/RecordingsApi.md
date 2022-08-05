# \RecordingsApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Cancel**](RecordingsApi.md#Cancel) | **Delete** /recordings/live/{recordingName} | Stop a live recording and discard it.
[**CopyStored**](RecordingsApi.md#CopyStored) | **Post** /recordings/stored/{recordingName}/copy | Copy a stored recording.
[**DeleteStored**](RecordingsApi.md#DeleteStored) | **Delete** /recordings/stored/{recordingName} | Delete a stored recording.
[**GetLive**](RecordingsApi.md#GetLive) | **Get** /recordings/live/{recordingName} | List live recordings.
[**GetStored**](RecordingsApi.md#GetStored) | **Get** /recordings/stored/{recordingName} | Get a stored recording&#39;s details.
[**GetStoredFile**](RecordingsApi.md#GetStoredFile) | **Get** /recordings/stored/{recordingName}/file | Get the file associated with the stored recording.
[**ListStored**](RecordingsApi.md#ListStored) | **Get** /recordings/stored | List recordings that are complete.
[**Muterecording**](RecordingsApi.md#Muterecording) | **Post** /recordings/live/{recordingName}/mute | Mute a live recording.
[**Pause**](RecordingsApi.md#Pause) | **Post** /recordings/live/{recordingName}/pause | Pause a live recording.
[**Stoprecording**](RecordingsApi.md#Stoprecording) | **Post** /recordings/live/{recordingName}/stop | Stop a live recording and store it.
[**Unmuterecording**](RecordingsApi.md#Unmuterecording) | **Delete** /recordings/live/{recordingName}/mute | Unmute a live recording.
[**Unpause**](RecordingsApi.md#Unpause) | **Delete** /recordings/live/{recordingName}/pause | Unpause a live recording.


# **Cancel**
> Cancel(ctx, recordingName)
Stop a live recording and discard it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CopyStored**
> StoredRecording CopyStored(ctx, recordingName, destinationRecordingName)
Copy a stored recording.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording to copy | 
  **destinationRecordingName** | **string**| The destination name of the recording | 

### Return type

[**StoredRecording**](StoredRecording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteStored**
> DeleteStored(ctx, recordingName)
Delete a stored recording.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetLive**
> LiveRecording GetLive(ctx, recordingName)
List live recordings.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

[**LiveRecording**](LiveRecording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStored**
> StoredRecording GetStored(ctx, recordingName)
Get a stored recording's details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

[**StoredRecording**](StoredRecording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStoredFile**
> string GetStoredFile(ctx, recordingName)
Get the file associated with the stored recording.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListStored**
> []StoredRecording ListStored(ctx, )
List recordings that are complete.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]StoredRecording**](StoredRecording.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Muterecording**
> Muterecording(ctx, recordingName)
Mute a live recording.

Muting a recording suspends silence detection, which will be restarted when the recording is unmuted.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Pause**
> Pause(ctx, recordingName)
Pause a live recording.

Pausing a recording suspends silence detection, which will be restarted when the recording is unpaused. Paused time is not included in the accounting for maxDurationSeconds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Stoprecording**
> Stoprecording(ctx, recordingName)
Stop a live recording and store it.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unmuterecording**
> Unmuterecording(ctx, recordingName)
Unmute a live recording.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Unpause**
> Unpause(ctx, recordingName)
Unpause a live recording.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recordingName** | **string**| The name of the recording | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

