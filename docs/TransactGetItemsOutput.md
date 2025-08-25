

# TransactGetItemsOutput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**consumedCapacity** | [**List&lt;ConsumedCapacity&gt;**](ConsumedCapacity.md) | If the &lt;i&gt;ReturnConsumedCapacity&lt;/i&gt; value was &lt;code&gt;TOTAL&lt;/code&gt;, this is an array of &lt;code&gt;ConsumedCapacity&lt;/code&gt; objects, one for each table addressed by &lt;code&gt;TransactGetItem&lt;/code&gt; objects in the &lt;i&gt;TransactItems&lt;/i&gt; parameter. These &lt;code&gt;ConsumedCapacity&lt;/code&gt; objects report the read-capacity units consumed by the &lt;code&gt;TransactGetItems&lt;/code&gt; call in that table. |  [optional] |
|**responses** | [**List&lt;ItemResponse&gt;**](ItemResponse.md) | &lt;p&gt;An ordered array of up to 100 &lt;code&gt;ItemResponse&lt;/code&gt; objects, each of which corresponds to the &lt;code&gt;TransactGetItem&lt;/code&gt; object in the same position in the &lt;i&gt;TransactItems&lt;/i&gt; array. Each &lt;code&gt;ItemResponse&lt;/code&gt; object contains a Map of the name-value pairs that are the projected attributes of the requested item.&lt;/p&gt; &lt;p&gt;If a requested item could not be retrieved, the corresponding &lt;code&gt;ItemResponse&lt;/code&gt; object is Null, or if the requested item has no projected attributes, the corresponding &lt;code&gt;ItemResponse&lt;/code&gt; object is an empty Map. &lt;/p&gt; |  [optional] |



