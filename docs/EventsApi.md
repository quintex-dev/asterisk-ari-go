# \EventsApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EventWebsocket**](EventsApi.md#EventWebsocket) | **Get** /events | WebSocket connection for events.
[**UserEvent**](EventsApi.md#UserEvent) | **Post** /events/user/{eventName} | Generate a user event.


# **EventWebsocket**
> Message EventWebsocket(ctx, app, optional)
WebSocket connection for events.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **app** | [**[]string**](string.md)| Applications to subscribe to. | 
 **optional** | ***EventsApiEventWebsocketOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiEventWebsocketOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscribeAll** | **optional.Bool**| Subscribe to all Asterisk events. If provided, the applications listed will be subscribed to all events, effectively disabling the application specific subscriptions. Default is &#39;false&#39;. | 

### Return type

[**Message**](Message.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserEvent**
> UserEvent(ctx, eventName, application, optional)
Generate a user event.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **eventName** | **string**| Event name | 
  **application** | **string**| The name of the application that will receive this event | 
 **optional** | ***EventsApiUserEventOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiUserEventOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **source** | [**optional.Interface of []string**](string.md)| URI for event source (channel:{channelId}, bridge:{bridgeId}, endpoint:{tech}/{resource}, deviceState:{deviceName} | 
 **variables** | [**optional.Interface of Containers**](Containers.md)| The \&quot;variables\&quot; key in the body object holds custom key/value pairs to add to the user event. Ex. { \&quot;variables\&quot;: { \&quot;key\&quot;: \&quot;value\&quot; } } | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

