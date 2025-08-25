

# RestoreTableFromBackupInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**targetTableName** | **String** | The name of the new table to which the backup must be restored. |  |
|**backupArn** | **String** | The Amazon Resource Name (ARN) associated with the backup. |  |
|**billingModeOverride** | **BillingMode** | The billing mode of the restored table. |  [optional] |
|**globalSecondaryIndexOverride** | [**List&lt;GlobalSecondaryIndex&gt;**](GlobalSecondaryIndex.md) | List of global secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore. |  [optional] |
|**localSecondaryIndexOverride** | [**List&lt;LocalSecondaryIndex&gt;**](LocalSecondaryIndex.md) | List of local secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore. |  [optional] |
|**provisionedThroughputOverride** | [**ProvisionedThroughput**](ProvisionedThroughput.md) | Provisioned throughput settings for the restored table. |  [optional] |
|**ssESpecificationOverride** | [**SSESpecification**](SSESpecification.md) | The new server-side encryption settings for the restored table. |  [optional] |



