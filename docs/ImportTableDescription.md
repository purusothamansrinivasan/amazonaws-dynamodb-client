

# ImportTableDescription

 Represents the properties of the table being imported into. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**importArn** | **String** |  The Amazon Resource Number (ARN) corresponding to the import request.  |  [optional] |
|**importStatus** | **ImportStatus** |  The status of the import.  |  [optional] |
|**tableArn** | **String** |  The Amazon Resource Number (ARN) of the table being imported into.  |  [optional] |
|**tableId** | **String** |  The table id corresponding to the table created by import table process.  |  [optional] |
|**clientToken** | **String** |  The client token that was provided for the import task. Reusing the client token on retry makes a call to &lt;code&gt;ImportTable&lt;/code&gt; idempotent.  |  [optional] |
|**s3BucketSource** | [**S3BucketSource**](S3BucketSource.md) |  Values for the S3 bucket the source file is imported from. Includes bucket name (required), key prefix (optional) and bucket account owner ID (optional).  |  [optional] |
|**errorCount** | **Integer** |  The number of errors occurred on importing the source file into the target table.  |  [optional] |
|**cloudWatchLogGroupArn** | **String** |  The Amazon Resource Number (ARN) of the Cloudwatch Log Group associated with the target table.  |  [optional] |
|**inputFormat** | **InputFormat** |  The format of the source data going into the target table.  |  [optional] |
|**inputFormatOptions** | [**InputFormatOptions**](InputFormatOptions.md) |  The format options for the data that was imported into the target table. There is one value, CsvOption.  |  [optional] |
|**inputCompressionType** | **InputCompressionType** |  The compression options for the data that has been imported into the target table. The values are NONE, GZIP, or ZSTD.  |  [optional] |
|**tableCreationParameters** | [**TableCreationParameters**](TableCreationParameters.md) |  The parameters for the new table that is being imported into.  |  [optional] |
|**startTime** | **OffsetDateTime** |  The time when this import task started.  |  [optional] |
|**endTime** | **OffsetDateTime** |  The time at which the creation of the table associated with this import task completed.  |  [optional] |
|**processedSizeBytes** | **Integer** |  The total size of data processed from the source file, in Bytes.  |  [optional] |
|**processedItemCount** | **Integer** |  The total number of items processed from the source file.  |  [optional] |
|**importedItemCount** | **Integer** |  The number of items successfully imported into the new table.  |  [optional] |
|**failureCode** | **String** |  The error code corresponding to the failure that the import job ran into during execution.  |  [optional] |
|**failureMessage** | **String** |  The error message corresponding to the failure that the import job ran into during execution.  |  [optional] |



