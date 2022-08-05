# \EndpointsApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Getendpoint**](EndpointsApi.md#Getendpoint) | **Get** /endpoints/{tech}/{resource} | Details for an endpoint.
[**ListByTech**](EndpointsApi.md#ListByTech) | **Get** /endpoints/{tech} | List available endoints for a given endpoint technology.
[**Listendpoints**](EndpointsApi.md#Listendpoints) | **Get** /endpoints | List all endpoints.
[**SendMessage**](EndpointsApi.md#SendMessage) | **Put** /endpoints/sendMessage | Send a message to some technology URI or endpoint.
[**SendMessageToEndpoint**](EndpointsApi.md#SendMessageToEndpoint) | **Put** /endpoints/{tech}/{resource}/sendMessage | Send a message to some endpoint in a technology.


# **Getendpoint**
> Endpoint Getendpoint(ctx, tech, resource)
Details for an endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tech** | **string**| Technology of the endpoint | 
  **resource** | **string**| ID of the endpoint | 

### Return type

[**Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListByTech**
> []Endpoint ListByTech(ctx, tech)
List available endoints for a given endpoint technology.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tech** | **string**| Technology of the endpoints (sip,iax2,...) | 

### Return type

[**[]Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Listendpoints**
> []Endpoint Listendpoints(ctx, )
List all endpoints.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Endpoint**](Endpoint.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendMessage**
> SendMessage(ctx, to, from, optional)
Send a message to some technology URI or endpoint.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **to** | **string**| The endpoint resource or technology specific URI to send the message to. Valid resources are sip, pjsip, and xmpp. | 
  **from** | **string**| The endpoint resource or technology specific identity to send this message from. Valid resources are sip, pjsip, and xmpp. | 
 **optional** | ***EndpointsApiSendMessageOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsApiSendMessageOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **optional.String**| The body of the message | 
 **variables** | [**optional.Interface of Containers**](Containers.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SendMessageToEndpoint**
> SendMessageToEndpoint(ctx, tech, resource, from, optional)
Send a message to some endpoint in a technology.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **tech** | **string**| Technology of the endpoint | 
  **resource** | **string**| ID of the endpoint | 
  **from** | **string**| The endpoint resource or technology specific identity to send this message from. Valid resources are sip, pjsip, and xmpp. | 
 **optional** | ***EndpointsApiSendMessageToEndpointOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EndpointsApiSendMessageToEndpointOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **optional.String**| The body of the message | 
 **variables** | [**optional.Interface of Containers**](Containers.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

