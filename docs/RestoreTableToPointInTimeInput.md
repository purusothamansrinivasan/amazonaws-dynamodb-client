

# RestoreTableToPointInTimeInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**sourceTableArn** | **String** | The DynamoDB table that will be restored. This value is an Amazon Resource Name (ARN). |  [optional] |
|**sourceTableName** | **String** | Name of the source table that is being restored. |  [optional] |
|**targetTableName** | **String** | The name of the new table to which it must be restored to. |  |
|**useLatestRestorableTime** | **Boolean** | Restore the table to the latest possible time. &lt;code&gt;LatestRestorableDateTime&lt;/code&gt; is typically 5 minutes before the current time.  |  [optional] |
|**restoreDateTime** | **OffsetDateTime** | Time in the past to restore the table to. |  [optional] |
|**billingModeOverride** | **BillingMode** | The billing mode of the restored table. |  [optional] |
|**globalSecondaryIndexOverride** | [**List&lt;GlobalSecondaryIndex&gt;**](GlobalSecondaryIndex.md) | List of global secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore. |  [optional] |
|**localSecondaryIndexOverride** | [**List&lt;LocalSecondaryIndex&gt;**](LocalSecondaryIndex.md) | List of local secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore. |  [optional] |
|**provisionedThroughputOverride** | [**ProvisionedThroughput**](ProvisionedThroughput.md) | Provisioned throughput settings for the restored table. |  [optional] |
|**ssESpecificationOverride** | [**SSESpecification**](SSESpecification.md) | The new server-side encryption settings for the restored table. |  [optional] |



