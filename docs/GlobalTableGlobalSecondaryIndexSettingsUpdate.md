

# GlobalTableGlobalSecondaryIndexSettingsUpdate

Represents the settings of a global secondary index for a global table that will be modified.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**indexName** | **String** | The name of the global secondary index. The name must be unique among all other indexes on this table. |  |
|**provisionedWriteCapacityUnits** | **Integer** | The maximum number of writes consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException.&lt;/code&gt;  |  [optional] |
|**provisionedWriteCapacityAutoScalingSettingsUpdate** | [**AutoScalingSettingsUpdate**](AutoScalingSettingsUpdate.md) | Auto scaling settings for managing a global secondary index&#39;s write capacity units. |  [optional] |



