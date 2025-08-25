

# ImportSummary

 Summary information about the source file for the import. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**importArn** | **String** |  The Amazon Resource Number (ARN) corresponding to the import request.  |  [optional] |
|**importStatus** | **ImportStatus** |  The status of the import operation.  |  [optional] |
|**tableArn** | **String** |  The Amazon Resource Number (ARN) of the table being imported into.  |  [optional] |
|**s3BucketSource** | [**S3BucketSource**](S3BucketSource.md) |  The path and S3 bucket of the source file that is being imported. This includes the S3Bucket (required), S3KeyPrefix (optional) and S3BucketOwner (optional if the bucket is owned by the requester).  |  [optional] |
|**cloudWatchLogGroupArn** | **String** |  The Amazon Resource Number (ARN) of the Cloudwatch Log Group associated with this import task.  |  [optional] |
|**inputFormat** | **InputFormat** |  The format of the source data. Valid values are &lt;code&gt;CSV&lt;/code&gt;, &lt;code&gt;DYNAMODB_JSON&lt;/code&gt; or &lt;code&gt;ION&lt;/code&gt;. |  [optional] |
|**startTime** | **OffsetDateTime** |  The time at which this import task began.  |  [optional] |
|**endTime** | **OffsetDateTime** |  The time at which this import task ended. (Does this include the successful complete creation of the table it was imported to?)  |  [optional] |



