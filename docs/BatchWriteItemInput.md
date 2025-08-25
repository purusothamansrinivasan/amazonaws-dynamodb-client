

# BatchWriteItemInput

Represents the input of a <code>BatchWriteItem</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**requestItems** | **Map&lt;String, List&lt;WriteRequest&gt;&gt;** | &lt;p&gt;A map of one or more table names and, for each table, a list of operations to be performed (&lt;code&gt;DeleteRequest&lt;/code&gt; or &lt;code&gt;PutRequest&lt;/code&gt;). Each element in the map consists of the following:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DeleteRequest&lt;/code&gt; - Perform a &lt;code&gt;DeleteItem&lt;/code&gt; operation on the specified item. The item to be deleted is identified by a &lt;code&gt;Key&lt;/code&gt; subelement:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;Key&lt;/code&gt; - A map of primary key attribute values that uniquely identify the item. Each entry in this map consists of an attribute name and an attribute value. For each primary key, you must provide &lt;i&gt;all&lt;/i&gt; of the key attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for &lt;i&gt;both&lt;/i&gt; the partition key and the sort key.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;PutRequest&lt;/code&gt; - Perform a &lt;code&gt;PutItem&lt;/code&gt; operation on the specified item. The item to be put is identified by an &lt;code&gt;Item&lt;/code&gt; subelement:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;Item&lt;/code&gt; - A map of attributes and their values. Each entry in this map consists of an attribute name and an attribute value. Attribute values must not be null; string and binary type attributes must have lengths greater than zero; and set type attributes must not be empty. Requests that contain empty values are rejected with a &lt;code&gt;ValidationException&lt;/code&gt; exception.&lt;/p&gt; &lt;p&gt;If you specify any attributes that are part of an index key, then the data types for those attributes must match those of the schema in the table&#39;s attribute definition.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;/li&gt; &lt;/ul&gt; |  |
|**returnConsumedCapacity** | **ReturnConsumedCapacity** |  |  [optional] |
|**returnItemCollectionMetrics** | **ReturnItemCollectionMetrics** | Determines whether item collection metrics are returned. If set to &lt;code&gt;SIZE&lt;/code&gt;, the response includes statistics about item collections, if any, that were modified during the operation are returned in the response. If set to &lt;code&gt;NONE&lt;/code&gt; (the default), no statistics are returned. |  [optional] |



