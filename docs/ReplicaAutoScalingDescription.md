

# ReplicaAutoScalingDescription

Represents the auto scaling settings of the replica.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**regionName** | **String** | The Region where the replica exists. |  [optional] |
|**globalSecondaryIndexes** | [**List&lt;ReplicaGlobalSecondaryIndexAutoScalingDescription&gt;**](ReplicaGlobalSecondaryIndexAutoScalingDescription.md) | Replica-specific global secondary index auto scaling settings. |  [optional] |
|**replicaProvisionedReadCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) |  |  [optional] |
|**replicaProvisionedWriteCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) |  |  [optional] |
|**replicaStatus** | **ReplicaStatus** | &lt;p&gt;The current state of the replica:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CREATING&lt;/code&gt; - The replica is being created.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - The replica is being updated.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DELETING&lt;/code&gt; - The replica is being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ACTIVE&lt;/code&gt; - The replica is ready for use.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |



