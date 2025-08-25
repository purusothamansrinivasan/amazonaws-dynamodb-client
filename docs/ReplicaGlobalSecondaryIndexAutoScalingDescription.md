

# ReplicaGlobalSecondaryIndexAutoScalingDescription

Represents the auto scaling configuration for a replica global secondary index.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**indexName** | **String** | The name of the global secondary index. |  [optional] |
|**indexStatus** | **IndexStatus** | &lt;p&gt;The current state of the replica global secondary index:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CREATING&lt;/code&gt; - The index is being created.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - The table/index configuration is being updated. The table/index remains available for data operations when &lt;code&gt;UPDATING&lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DELETING&lt;/code&gt; - The index is being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ACTIVE&lt;/code&gt; - The index is ready for use.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**provisionedReadCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) |  |  [optional] |
|**provisionedWriteCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) |  |  [optional] |



