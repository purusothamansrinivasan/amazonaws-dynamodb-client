

# BatchGetItemOutput

Represents the output of a <code>BatchGetItem</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**responses** | **Map&lt;String, List&lt;Map&lt;String, AttributeValue&gt;&gt;&gt;** | A map of table name to a list of items. Each object in &lt;code&gt;Responses&lt;/code&gt; consists of a table name, along with a map of attribute data consisting of the data type and attribute value. |  [optional] |
|**unprocessedKeys** | [**Map&lt;String, KeysAndAttributes&gt;**](KeysAndAttributes.md) | &lt;p&gt;A map of tables and their respective keys that were not processed with the current response. The &lt;code&gt;UnprocessedKeys&lt;/code&gt; value is in the same form as &lt;code&gt;RequestItems&lt;/code&gt;, so the value can be provided directly to a subsequent &lt;code&gt;BatchGetItem&lt;/code&gt; operation. For more information, see &lt;code&gt;RequestItems&lt;/code&gt; in the Request Parameters section.&lt;/p&gt; &lt;p&gt;Each element consists of:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;Keys&lt;/code&gt; - An array of primary key attribute values that define specific items in the table.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ProjectionExpression&lt;/code&gt; - One or more attributes to be retrieved from the table or index. By default, all attributes are returned. If a requested attribute is not found, it does not appear in the result.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ConsistentRead&lt;/code&gt; - The consistency of a read operation. If set to &lt;code&gt;true&lt;/code&gt;, then a strongly consistent read is used; otherwise, an eventually consistent read is used.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; &lt;p&gt;If there are no unprocessed keys remaining, the response contains an empty &lt;code&gt;UnprocessedKeys&lt;/code&gt; map.&lt;/p&gt; |  [optional] |
|**consumedCapacity** | [**List&lt;ConsumedCapacity&gt;**](ConsumedCapacity.md) | &lt;p&gt;The read capacity units consumed by the entire &lt;code&gt;BatchGetItem&lt;/code&gt; operation.&lt;/p&gt; &lt;p&gt;Each element consists of:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;TableName&lt;/code&gt; - The table that consumed the provisioned throughput.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CapacityUnits&lt;/code&gt; - The total number of capacity units consumed.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |



