

# ReplicaSettingsUpdate

Represents the settings for a global table in a Region that will be modified.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**regionName** | **String** | The Region of the replica to be added. |  |
|**replicaProvisionedReadCapacityUnits** | **Integer** | The maximum number of strongly consistent reads consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput\&quot;&gt;Specifying Read and Write Requirements&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;.  |  [optional] |
|**replicaProvisionedReadCapacityAutoScalingSettingsUpdate** | [**AutoScalingSettingsUpdate**](AutoScalingSettingsUpdate.md) | Auto scaling settings for managing a global table replica&#39;s read capacity units. |  [optional] |
|**replicaGlobalSecondaryIndexSettingsUpdate** | [**List&lt;ReplicaGlobalSecondaryIndexSettingsUpdate&gt;**](ReplicaGlobalSecondaryIndexSettingsUpdate.md) | Represents the settings of a global secondary index for a global table that will be modified. |  [optional] |
|**replicaTableClass** | **TableClass** | Replica-specific table class. If not specified, uses the source table&#39;s table class. |  [optional] |



