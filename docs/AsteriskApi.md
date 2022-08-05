# \AsteriskApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddLog**](AsteriskApi.md#AddLog) | **Post** /asterisk/logging/{logChannelName} | Adds a log channel.
[**DeleteLog**](AsteriskApi.md#DeleteLog) | **Delete** /asterisk/logging/{logChannelName} | Deletes a log channel.
[**DeleteObject**](AsteriskApi.md#DeleteObject) | **Delete** /asterisk/config/dynamic/{configClass}/{objectType}/{id} | Delete a dynamic configuration object.
[**GetGlobalVar**](AsteriskApi.md#GetGlobalVar) | **Get** /asterisk/variable | Get the value of a global variable.
[**GetInfo**](AsteriskApi.md#GetInfo) | **Get** /asterisk/info | Gets Asterisk system information.
[**GetModule**](AsteriskApi.md#GetModule) | **Get** /asterisk/modules/{moduleName} | Get Asterisk module information.
[**GetObject**](AsteriskApi.md#GetObject) | **Get** /asterisk/config/dynamic/{configClass}/{objectType}/{id} | Retrieve a dynamic configuration object.
[**ListLogChannels**](AsteriskApi.md#ListLogChannels) | **Get** /asterisk/logging | Gets Asterisk log channel information.
[**ListModules**](AsteriskApi.md#ListModules) | **Get** /asterisk/modules | List Asterisk modules.
[**LoadModule**](AsteriskApi.md#LoadModule) | **Post** /asterisk/modules/{moduleName} | Load an Asterisk module.
[**Ping**](AsteriskApi.md#Ping) | **Get** /asterisk/ping | Response pong message.
[**ReloadModule**](AsteriskApi.md#ReloadModule) | **Put** /asterisk/modules/{moduleName} | Reload an Asterisk module.
[**RotateLog**](AsteriskApi.md#RotateLog) | **Put** /asterisk/logging/{logChannelName}/rotate | Rotates a log channel.
[**SetGlobalVar**](AsteriskApi.md#SetGlobalVar) | **Post** /asterisk/variable | Set the value of a global variable.
[**UnloadModule**](AsteriskApi.md#UnloadModule) | **Delete** /asterisk/modules/{moduleName} | Unload an Asterisk module.
[**UpdateObject**](AsteriskApi.md#UpdateObject) | **Put** /asterisk/config/dynamic/{configClass}/{objectType}/{id} | Create or update a dynamic configuration object.


# **AddLog**
> AddLog(ctx, logChannelName, configuration)
Adds a log channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logChannelName** | **string**| The log channel to add | 
  **configuration** | **string**| levels of the log channel | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteLog**
> DeleteLog(ctx, logChannelName)
Deletes a log channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logChannelName** | **string**| Log channels name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteObject**
> DeleteObject(ctx, configClass, objectType, id)
Delete a dynamic configuration object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configClass** | **string**| The configuration class containing dynamic configuration objects. | 
  **objectType** | **string**| The type of configuration object to delete. | 
  **id** | **string**| The unique identifier of the object to delete. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetGlobalVar**
> Variable GetGlobalVar(ctx, variable)
Get the value of a global variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **variable** | **string**| The variable to get | 

### Return type

[**Variable**](Variable.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetInfo**
> AsteriskInfo GetInfo(ctx, optional)
Gets Asterisk system information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AsteriskApiGetInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsteriskApiGetInfoOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **only** | [**optional.Interface of []string**](string.md)| Filter information returned | 

### Return type

[**AsteriskInfo**](AsteriskInfo.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetModule**
> Module GetModule(ctx, moduleName)
Get Asterisk module information.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**| Module&#39;s name | 

### Return type

[**Module**](Module.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetObject**
> []ConfigTuple GetObject(ctx, configClass, objectType, id)
Retrieve a dynamic configuration object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configClass** | **string**| The configuration class containing dynamic configuration objects. | 
  **objectType** | **string**| The type of configuration object to retrieve. | 
  **id** | **string**| The unique identifier of the object to retrieve. | 

### Return type

[**[]ConfigTuple**](ConfigTuple.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListLogChannels**
> []LogChannel ListLogChannels(ctx, )
Gets Asterisk log channel information.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]LogChannel**](LogChannel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListModules**
> []Module ListModules(ctx, )
List Asterisk modules.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Module**](Module.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **LoadModule**
> LoadModule(ctx, moduleName)
Load an Asterisk module.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**| Module&#39;s name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Ping**
> AsteriskPing Ping(ctx, )
Response pong message.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AsteriskPing**](AsteriskPing.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ReloadModule**
> ReloadModule(ctx, moduleName)
Reload an Asterisk module.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**| Module&#39;s name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RotateLog**
> RotateLog(ctx, logChannelName)
Rotates a log channel.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **logChannelName** | **string**| Log channel&#39;s name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetGlobalVar**
> SetGlobalVar(ctx, variable, optional)
Set the value of a global variable.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **variable** | **string**| The variable to set | 
 **optional** | ***AsteriskApiSetGlobalVarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsteriskApiSetGlobalVarOpts struct

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

# **UnloadModule**
> UnloadModule(ctx, moduleName)
Unload an Asterisk module.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **moduleName** | **string**| Module&#39;s name | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateObject**
> []ConfigTuple UpdateObject(ctx, configClass, objectType, id, optional)
Create or update a dynamic configuration object.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **configClass** | **string**| The configuration class containing dynamic configuration objects. | 
  **objectType** | **string**| The type of configuration object to create or update. | 
  **id** | **string**| The unique identifier of the object to create or update. | 
 **optional** | ***AsteriskApiUpdateObjectOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AsteriskApiUpdateObjectOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **fields** | [**optional.Interface of Containers**](Containers.md)| The body object should have a value that is a list of ConfigTuples, which provide the fields to update. Ex. [ { \&quot;attribute\&quot;: \&quot;directmedia\&quot;, \&quot;value\&quot;: \&quot;false\&quot; } ] | 

### Return type

[**[]ConfigTuple**](ConfigTuple.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

