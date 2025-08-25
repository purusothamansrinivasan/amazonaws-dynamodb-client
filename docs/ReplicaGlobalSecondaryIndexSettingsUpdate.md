

# ReplicaGlobalSecondaryIndexSettingsUpdate

Represents the settings of a global secondary index for a global table that will be modified.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**indexName** | **String** | The name of the global secondary index. The name must be unique among all other indexes on this table. |  |
|**provisionedReadCapacityUnits** | **Integer** | The maximum number of strongly consistent reads consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. |  [optional] |
|**provisionedReadCapacityAutoScalingSettingsUpdate** | [**AutoScalingSettingsUpdate**](AutoScalingSettingsUpdate.md) | Auto scaling settings for managing a global secondary index replica&#39;s read capacity units. |  [optional] |



