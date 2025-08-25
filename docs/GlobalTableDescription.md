

# GlobalTableDescription

Contains details about the global table.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**replicationGroup** | [**List&lt;ReplicaDescription&gt;**](ReplicaDescription.md) | The Regions where the global table has replicas. |  [optional] |
|**globalTableArn** | **String** | The unique identifier of the global table. |  [optional] |
|**creationDateTime** | **OffsetDateTime** | The creation time of the global table. |  [optional] |
|**globalTableStatus** | **GlobalTableStatus** | &lt;p&gt;The current state of the global table:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CREATING&lt;/code&gt; - The global table is being created.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - The global table is being updated.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DELETING&lt;/code&gt; - The global table is being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ACTIVE&lt;/code&gt; - The global table is ready for use.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**globalTableName** | **String** | The global table name. |  [optional] |



