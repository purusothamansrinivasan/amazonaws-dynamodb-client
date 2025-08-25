

# UpdateGlobalTableSettingsInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**globalTableName** | **String** | The name of the global table |  |
|**globalTableBillingMode** | **BillingMode** | &lt;p&gt;The billing mode of the global table. If &lt;code&gt;GlobalTableBillingMode&lt;/code&gt; is not specified, the global table defaults to &lt;code&gt;PROVISIONED&lt;/code&gt; capacity billing mode.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;PROVISIONED&lt;/code&gt; - We recommend using &lt;code&gt;PROVISIONED&lt;/code&gt; for predictable workloads. &lt;code&gt;PROVISIONED&lt;/code&gt; sets the billing mode to &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual\&quot;&gt;Provisioned Mode&lt;/a&gt;.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; - We recommend using &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; for unpredictable workloads. &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; sets the billing mode to &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand\&quot;&gt;On-Demand Mode&lt;/a&gt;. &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**globalTableProvisionedWriteCapacityUnits** | **Integer** | The maximum number of writes consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException.&lt;/code&gt;  |  [optional] |
|**globalTableProvisionedWriteCapacityAutoScalingSettingsUpdate** | [**AutoScalingSettingsUpdate**](AutoScalingSettingsUpdate.md) | Auto scaling settings for managing provisioned write capacity for the global table. |  [optional] |
|**globalTableGlobalSecondaryIndexSettingsUpdate** | [**List&lt;GlobalTableGlobalSecondaryIndexSettingsUpdate&gt;**](GlobalTableGlobalSecondaryIndexSettingsUpdate.md) | Represents the settings of a global secondary index for a global table that will be modified. |  [optional] |
|**replicaSettingsUpdate** | [**List&lt;ReplicaSettingsUpdate&gt;**](ReplicaSettingsUpdate.md) | Represents the settings for a global table in a Region that will be modified. |  [optional] |



