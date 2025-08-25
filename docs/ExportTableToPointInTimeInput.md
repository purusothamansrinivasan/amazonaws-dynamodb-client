

# ExportTableToPointInTimeInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableArn** | **String** | The Amazon Resource Name (ARN) associated with the table to export. |  |
|**exportTime** | **OffsetDateTime** | Time in the past from which to export table data, counted in seconds from the start of the Unix epoch. The table export will be a snapshot of the table&#39;s state at this point in time. |  [optional] |
|**clientToken** | **String** | &lt;p&gt;Providing a &lt;code&gt;ClientToken&lt;/code&gt; makes the call to &lt;code&gt;ExportTableToPointInTimeInput&lt;/code&gt; idempotent, meaning that multiple identical calls have the same effect as one single call.&lt;/p&gt; &lt;p&gt;A client token is valid for 8 hours after the first request that uses it is completed. After 8 hours, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 8 hours, or the result might not be idempotent.&lt;/p&gt; &lt;p&gt;If you submit a request with the same client token but a change in other parameters within the 8-hour idempotency window, DynamoDB returns an &lt;code&gt;ImportConflictException&lt;/code&gt;.&lt;/p&gt; |  [optional] |
|**s3Bucket** | **String** | The name of the Amazon S3 bucket to export the snapshot to. |  |
|**s3BucketOwner** | **String** | The ID of the Amazon Web Services account that owns the bucket the export will be stored in. |  [optional] |
|**s3Prefix** | **String** | The Amazon S3 bucket prefix to use as the file name and path of the exported snapshot. |  [optional] |
|**s3SseAlgorithm** | **S3SseAlgorithm** | &lt;p&gt;Type of encryption used on the bucket where export data will be stored. Valid values for &lt;code&gt;S3SseAlgorithm&lt;/code&gt; are:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AES256&lt;/code&gt; - server-side encryption with Amazon S3 managed keys&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;KMS&lt;/code&gt; - server-side encryption with KMS managed keys&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**s3SseKmsKeyId** | **String** | The ID of the KMS managed key used to encrypt the S3 bucket where export data will be stored (if applicable). |  [optional] |
|**exportFormat** | **ExportFormat** | The format for the exported data. Valid values for &lt;code&gt;ExportFormat&lt;/code&gt; are &lt;code&gt;DYNAMODB_JSON&lt;/code&gt; or &lt;code&gt;ION&lt;/code&gt;. |  [optional] |



