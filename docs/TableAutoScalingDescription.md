

# TableAutoScalingDescription

Represents the auto scaling configuration for a global table.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableName** | **String** | The name of the table. |  [optional] |
|**tableStatus** | **TableStatus** | &lt;p&gt;The current state of the table:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CREATING&lt;/code&gt; - The table is being created.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - The table is being updated.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DELETING&lt;/code&gt; - The table is being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ACTIVE&lt;/code&gt; - The table is ready for use.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**replicas** | [**List&lt;ReplicaAutoScalingDescription&gt;**](ReplicaAutoScalingDescription.md) | Represents replicas of the global table. |  [optional] |



