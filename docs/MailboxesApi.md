# \MailboxesApi

All URIs are relative to *http://localhost:8088/ari*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Deletemailbox**](MailboxesApi.md#Deletemailbox) | **Delete** /mailboxes/{mailboxName} | Destroy a mailbox.
[**Getmailbox**](MailboxesApi.md#Getmailbox) | **Get** /mailboxes/{mailboxName} | Retrieve the current state of a mailbox.
[**Listmailboxes**](MailboxesApi.md#Listmailboxes) | **Get** /mailboxes | List all mailboxes.
[**Updatemailbox**](MailboxesApi.md#Updatemailbox) | **Put** /mailboxes/{mailboxName} | Change the state of a mailbox. (Note - implicitly creates the mailbox).


# **Deletemailbox**
> Deletemailbox(ctx, mailboxName)
Destroy a mailbox.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mailboxName** | **string**| Name of the mailbox | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Getmailbox**
> Mailbox Getmailbox(ctx, mailboxName)
Retrieve the current state of a mailbox.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mailboxName** | **string**| Name of the mailbox | 

### Return type

[**Mailbox**](Mailbox.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Listmailboxes**
> []Mailbox Listmailboxes(ctx, )
List all mailboxes.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]Mailbox**](Mailbox.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **Updatemailbox**
> Updatemailbox(ctx, mailboxName, oldMessages, newMessages)
Change the state of a mailbox. (Note - implicitly creates the mailbox).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **mailboxName** | **string**| Name of the mailbox | 
  **oldMessages** | **int32**| Count of old messages in the mailbox | 
  **newMessages** | **int32**| Count of new messages in the mailbox | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

