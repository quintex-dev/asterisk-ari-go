# LiveRecording

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | **string** | Cause for recording failure if failed | [optional] [default to null]
**Duration** | **int32** | Duration in seconds of the recording | [optional] [default to null]
**Format** | **string** | Recording format (wav, gsm, etc.) | [default to null]
**Name** | **string** | Base name for the recording | [default to null]
**SilenceDuration** | **int32** | Duration of silence, in seconds, detected in the recording. This is only available if the recording was initiated with a non-zero maxSilenceSeconds. | [optional] [default to null]
**State** | **string** |  | [default to null]
**TalkingDuration** | **int32** | Duration of talking, in seconds, detected in the recording. This is only available if the recording was initiated with a non-zero maxSilenceSeconds. | [optional] [default to null]
**TargetUri** | **string** | URI for the channel or bridge being recorded | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


