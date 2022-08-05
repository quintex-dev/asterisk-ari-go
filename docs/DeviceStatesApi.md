# \DeviceStatesApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Delete**](DeviceStatesApi.md#Delete) | **Delete** /deviceStates/{deviceName} | Destroy a device-state controlled by ARI.
[**Getdevicestate**](DeviceStatesApi.md#Getdevicestate) | **Get** /deviceStates/{deviceName} | Retrieve the current state of a device.
[**ListDeviceStates**](DeviceStatesApi.md#ListDeviceStates) | **Get** /deviceStates | List all ARI controlled device states.
[**Update**](DeviceStatesApi.md#Update) | **Put** /deviceStates/{deviceName} | Change the state of a device controlled by ARI. (Note - implicitly creates the device state).


# **Delete**
> Delete(ctx, deviceName)
Destroy a device-state controlled by ARI.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceName** | **string**| Name of the device | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getdevicestate**
> DeviceState Getdevicestate(ctx, deviceName)
Retrieve the current state of a device.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceName** | **string**| Name of the device | 

### Return type

[**DeviceState**](DeviceState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListDeviceStates**
> []DeviceState ListDeviceStates(ctx, )
List all ARI controlled device states.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DeviceState**](DeviceState.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Update**
> Update(ctx, deviceName, deviceState)
Change the state of a device controlled by ARI. (Note - implicitly creates the device state).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **deviceName** | **string**| Name of the device | 
  **deviceState** | **string**| Device state value | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

