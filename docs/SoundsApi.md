# \SoundsApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getsound**](SoundsApi.md#Getsound) | **Get** /sounds/{soundId} | Get a sound&#39;s details.
[**Listsounds**](SoundsApi.md#Listsounds) | **Get** /sounds | List all sounds.


# **Getsound**
> Sound Getsound(ctx, soundId)
Get a sound's details.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **soundId** | **string**| Sound&#39;s id | 

### Return type

[**Sound**](Sound.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Listsounds**
> []Sound Listsounds(ctx, optional)
List all sounds.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***SoundsApiListsoundsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a SoundsApiListsoundsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **lang** | **optional.String**| Lookup sound for a specific language. | 
 **format** | **optional.String**| Lookup sound in a specific format. | 

### Return type

[**[]Sound**](Sound.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

