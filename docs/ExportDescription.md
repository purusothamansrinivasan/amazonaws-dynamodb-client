

# ExportDescription

Represents the properties of the exported table.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**exportArn** | **String** | The Amazon Resource Name (ARN) of the table export. |  [optional] |
|**exportStatus** | **ExportStatus** | Export can be in one of the following states: IN_PROGRESS, COMPLETED, or FAILED. |  [optional] |
|**startTime** | **OffsetDateTime** | The time at which the export task began. |  [optional] |
|**endTime** | **OffsetDateTime** | The time at which the export task completed. |  [optional] |
|**exportManifest** | **String** | The name of the manifest file for the export task. |  [optional] |
|**tableArn** | **String** | The Amazon Resource Name (ARN) of the table that was exported. |  [optional] |
|**tableId** | **String** | Unique ID of the table that was exported. |  [optional] |
|**exportTime** | **OffsetDateTime** | Point in time from which table data was exported. |  [optional] |
|**clientToken** | **String** | The client token that was provided for the export task. A client token makes calls to &lt;code&gt;ExportTableToPointInTimeInput&lt;/code&gt; idempotent, meaning that multiple identical calls have the same effect as one single call. |  [optional] |
|**s3Bucket** | **String** | The name of the Amazon S3 bucket containing the export. |  [optional] |
|**s3BucketOwner** | **String** | The ID of the Amazon Web Services account that owns the bucket containing the export. |  [optional] |
|**s3Prefix** | **String** | The Amazon S3 bucket prefix used as the file name and path of the exported snapshot. |  [optional] |
|**s3SseAlgorithm** | **S3SseAlgorithm** | &lt;p&gt;Type of encryption used on the bucket where export data is stored. Valid values for &lt;code&gt;S3SseAlgorithm&lt;/code&gt; are:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AES256&lt;/code&gt; - server-side encryption with Amazon S3 managed keys&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;KMS&lt;/code&gt; - server-side encryption with KMS managed keys&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**s3SseKmsKeyId** | **String** | The ID of the KMS managed key used to encrypt the S3 bucket where export data is stored (if applicable). |  [optional] |
|**failureCode** | **String** | Status code for the result of the failed export. |  [optional] |
|**failureMessage** | **String** | Export failure reason description. |  [optional] |
|**exportFormat** | **ExportFormat** | The format of the exported data. Valid values for &lt;code&gt;ExportFormat&lt;/code&gt; are &lt;code&gt;DYNAMODB_JSON&lt;/code&gt; or &lt;code&gt;ION&lt;/code&gt;. |  [optional] |
|**billedSizeBytes** | **Integer** | The billable size of the table export. |  [optional] |
|**itemCount** | **Integer** | The number of items exported. |  [optional] |



