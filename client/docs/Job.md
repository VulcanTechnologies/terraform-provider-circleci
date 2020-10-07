# Job

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CanceledBy** | **string** | The unique ID of the user. | [optional] 
**Dependencies** | **[]string** | A sequence of the unique job IDs for the jobs that this job depends upon in the workflow. | 
**JobNumber** | **int64** | The number of the job. | [optional] 
**Id** | **string** | The unique ID of the job. | 
**StartedAt** | [**time.Time**](time.Time.md) | The date and time the job started. | 
**Name** | **string** | The name of the job. | 
**ApprovedBy** | **string** | The unique ID of the user. | [optional] 
**ProjectSlug** | **string** | The project-slug for the job. | 
**Status** | [**interface{}**](.md) | The current status of the job. | 
**Type** | **string** | The type of job. | 
**StoppedAt** | [**time.Time**](time.Time.md) | The time when the job stopped. | [optional] 
**ApprovalRequestId** | **string** | The unique ID of the job. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


