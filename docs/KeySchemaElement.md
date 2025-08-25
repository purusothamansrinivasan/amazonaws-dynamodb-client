

# KeySchemaElement

<p>Represents <i>a single element</i> of a key schema. A key schema specifies the attributes that make up the primary key of a table, or the key attributes of an index.</p> <p>A <code>KeySchemaElement</code> represents exactly one attribute of the primary key. For example, a simple primary key would be represented by one <code>KeySchemaElement</code> (for the partition key). A composite primary key would require one <code>KeySchemaElement</code> for the partition key, and another <code>KeySchemaElement</code> for the sort key.</p> <p>A <code>KeySchemaElement</code> must be a scalar, top-level attribute (not a nested attribute). The data type must be one of String, Number, or Binary. The attribute cannot be nested within a List or a Map.</p>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**attributeName** | **String** | The name of a key attribute. |  |
|**keyType** | **KeyType** | &lt;p&gt;The role that this key attribute will assume:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;HASH&lt;/code&gt; - partition key&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;RANGE&lt;/code&gt; - sort key&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;note&gt; &lt;p&gt;The partition key of an item is also known as its &lt;i&gt;hash attribute&lt;/i&gt;. The term \&quot;hash attribute\&quot; derives from DynamoDB&#39;s usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.&lt;/p&gt; &lt;p&gt;The sort key of an item is also known as its &lt;i&gt;range attribute&lt;/i&gt;. The term \&quot;range attribute\&quot; derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.&lt;/p&gt; &lt;/note&gt; |  |



