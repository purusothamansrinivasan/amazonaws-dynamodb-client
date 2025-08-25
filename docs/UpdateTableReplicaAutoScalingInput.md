

# UpdateTableReplicaAutoScalingInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**globalSecondaryIndexUpdates** | [**List&lt;GlobalSecondaryIndexAutoScalingUpdate&gt;**](GlobalSecondaryIndexAutoScalingUpdate.md) | Represents the auto scaling settings of the global secondary indexes of the replica to be updated. |  [optional] |
|**tableName** | **String** | The name of the global table to be updated. |  |
|**provisionedWriteCapacityAutoScalingUpdate** | [**AutoScalingSettingsUpdate**](AutoScalingSettingsUpdate.md) |  |  [optional] |
|**replicaUpdates** | [**List&lt;ReplicaAutoScalingUpdate&gt;**](ReplicaAutoScalingUpdate.md) | Represents the auto scaling settings of replicas of the table that will be modified. |  [optional] |



