

# ImportTableInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**clientToken** | **String** | &lt;p&gt;Providing a &lt;code&gt;ClientToken&lt;/code&gt; makes the call to &lt;code&gt;ImportTableInput&lt;/code&gt; idempotent, meaning that multiple identical calls have the same effect as one single call.&lt;/p&gt; &lt;p&gt;A client token is valid for 8 hours after the first request that uses it is completed. After 8 hours, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 8 hours, or the result might not be idempotent.&lt;/p&gt; &lt;p&gt;If you submit a request with the same client token but a change in other parameters within the 8-hour idempotency window, DynamoDB returns an &lt;code&gt;IdempotentParameterMismatch&lt;/code&gt; exception.&lt;/p&gt; |  [optional] |
|**s3BucketSource** | [**S3BucketSource**](S3BucketSource.md) |  The S3 bucket that provides the source for the import.  |  |
|**inputFormat** | **InputFormat** |  The format of the source data. Valid values for &lt;code&gt;ImportFormat&lt;/code&gt; are &lt;code&gt;CSV&lt;/code&gt;, &lt;code&gt;DYNAMODB_JSON&lt;/code&gt; or &lt;code&gt;ION&lt;/code&gt;.  |  |
|**inputFormatOptions** | [**InputFormatOptions**](InputFormatOptions.md) |  Additional properties that specify how the input is formatted,  |  [optional] |
|**inputCompressionType** | **InputCompressionType** |  Type of compression to be used on the input coming from the imported table.  |  [optional] |
|**tableCreationParameters** | [**TableCreationParameters**](TableCreationParameters.md) | Parameters for the table to import the data into.  |  |



