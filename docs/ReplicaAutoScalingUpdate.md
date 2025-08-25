

# ReplicaAutoScalingUpdate

Represents the auto scaling settings of a replica that will be modified.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**regionName** | **String** | The Region where the replica exists. |  |
|**replicaGlobalSecondaryIndexUpdates** | [**List&lt;ReplicaGlobalSecondaryIndexAutoScalingUpdate&gt;**](ReplicaGlobalSecondaryIndexAutoScalingUpdate.md) | Represents the auto scaling settings of global secondary indexes that will be modified. |  [optional] |
|**replicaProvisionedReadCapacityAutoScalingUpdate** | [**AutoScalingSettingsUpdate**](AutoScalingSettingsUpdate.md) |  |  [optional] |



