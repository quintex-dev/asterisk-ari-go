# \PlaybacksApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Control**](PlaybacksApi.md#Control) | **Post** /playbacks/{playbackId}/control | Control a playback.
[**Getplayback**](PlaybacksApi.md#Getplayback) | **Get** /playbacks/{playbackId} | Get a playback&#39;s details.
[**Stop**](PlaybacksApi.md#Stop) | **Delete** /playbacks/{playbackId} | Stop a playback.


# **Control**
> Control(ctx, playbackId, operation)
Control a playback.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playbackId** | **string**| Playback&#39;s id | 
  **operation** | **string**| Operation to perform on the playback. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getplayback**
> Playback Getplayback(ctx, playbackId)
Get a playback's details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playbackId** | **string**| Playback&#39;s id | 

### Return type

[**Playback**](Playback.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Stop**
> Stop(ctx, playbackId)
Stop a playback.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **playbackId** | **string**| Playback&#39;s id | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

