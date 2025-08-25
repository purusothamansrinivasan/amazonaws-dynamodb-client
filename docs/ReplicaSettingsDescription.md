

# ReplicaSettingsDescription

Represents the properties of a replica.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**regionName** | **String** | The Region name of the replica. |  |
|**replicaStatus** | **ReplicaStatus** | &lt;p&gt;The current state of the Region:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CREATING&lt;/code&gt; - The Region is being created.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - The Region is being updated.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DELETING&lt;/code&gt; - The Region is being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ACTIVE&lt;/code&gt; - The Region is ready for use.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**replicaBillingModeSummary** | [**BillingModeSummary**](BillingModeSummary.md) | The read/write capacity mode of the replica. |  [optional] |
|**replicaProvisionedReadCapacityUnits** | **Integer** | The maximum number of strongly consistent reads consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput\&quot;&gt;Specifying Read and Write Requirements&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;.  |  [optional] |
|**replicaProvisionedReadCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) | Auto scaling settings for a global table replica&#39;s read capacity units. |  [optional] |
|**replicaProvisionedWriteCapacityUnits** | **Integer** | The maximum number of writes consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput\&quot;&gt;Specifying Read and Write Requirements&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  [optional] |
|**replicaProvisionedWriteCapacityAutoScalingSettings** | [**AutoScalingSettingsDescription**](AutoScalingSettingsDescription.md) | Auto scaling settings for a global table replica&#39;s write capacity units. |  [optional] |
|**replicaGlobalSecondaryIndexSettings** | [**List&lt;ReplicaGlobalSecondaryIndexSettingsDescription&gt;**](ReplicaGlobalSecondaryIndexSettingsDescription.md) | Replica global secondary index settings for the global table. |  [optional] |
|**replicaTableClassSummary** | [**TableClassSummary**](TableClassSummary.md) |  |  [optional] |



