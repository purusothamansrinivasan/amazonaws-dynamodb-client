

# LocalSecondaryIndexInfo

Represents the properties of a local secondary index for the table when the backup was created.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**indexName** | **String** | Represents the name of the local secondary index. |  [optional] |
|**keySchema** | [**List&lt;KeySchemaElement&gt;**](KeySchemaElement.md) | &lt;p&gt;The complete key schema for a local secondary index, which consists of one or more pairs of attribute names and key types:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;HASH&lt;/code&gt; - partition key&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;RANGE&lt;/code&gt; - sort key&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;note&gt; &lt;p&gt;The partition key of an item is also known as its &lt;i&gt;hash attribute&lt;/i&gt;. The term \&quot;hash attribute\&quot; derives from DynamoDB&#39;s usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.&lt;/p&gt; &lt;p&gt;The sort key of an item is also known as its &lt;i&gt;range attribute&lt;/i&gt;. The term \&quot;range attribute\&quot; derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.&lt;/p&gt; &lt;/note&gt; |  [optional] |
|**projection** | [**Projection**](Projection.md) | Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.  |  [optional] |



