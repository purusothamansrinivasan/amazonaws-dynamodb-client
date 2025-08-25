

# ReplicaGlobalSecondaryIndexSettingsDescription

Represents the properties of a global secondary index.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**indexName** | **String** | The name of the global secondary index. The name must be unique among all other indexes on this table. |  |
|**indexStatus** | **IndexStatus** | &lt;p&gt; The current status of the global secondary index:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CREATING&lt;/code&gt; - The global secondary index is being created.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - The global secondary index is being updated.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DELETING&lt;/code&gt; - The global secondary index is being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ACTIVE&lt;/code&gt; - The global secondary index is ready for use.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**provisionedReadCapacityUnits** | **Integer** | The maximum number of strongly consistent reads consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. |  [optional] |
|**provisionedReadCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) | Auto scaling settings for a global secondary index replica&#39;s read capacity units. |  [optional] |
|**provisionedWriteCapacityUnits** | **Integer** | The maximum number of writes consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. |  [optional] |
|**provisionedWriteCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) | Auto scaling settings for a global secondary index replica&#39;s write capacity units. |  [optional] |



