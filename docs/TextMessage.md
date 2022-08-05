# TextMessage

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | **string** | The text of the message. | [default to null]
**From** | **string** | A technology specific URI specifying the source of the message. For sip and pjsip technologies, any SIP URI can be specified. For xmpp, the URI must correspond to the client connection being used to send the message. | [default to null]
**To** | **string** | A technology specific URI specifying the destination of the message. Valid technologies include sip, pjsip, and xmp. The destination of a message should be an endpoint. | [default to null]
**Variables** | **interface{}** | Technology specific key/value pairs (JSON object) associated with the message. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


