

# TransactWriteItemsOutput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**consumedCapacity** | [**List&lt;ConsumedCapacity&gt;**](ConsumedCapacity.md) | The capacity units consumed by the entire &lt;code&gt;TransactWriteItems&lt;/code&gt; operation. The values of the list are ordered according to the ordering of the &lt;code&gt;TransactItems&lt;/code&gt; request parameter.  |  [optional] |
|**itemCollectionMetrics** | **Map&lt;String, List&lt;ItemCollectionMetrics&gt;&gt;** | A list of tables that were processed by &lt;code&gt;TransactWriteItems&lt;/code&gt; and, for each table, information about any item collections that were affected by individual &lt;code&gt;UpdateItem&lt;/code&gt;, &lt;code&gt;PutItem&lt;/code&gt;, or &lt;code&gt;DeleteItem&lt;/code&gt; operations.  |  [optional] |



