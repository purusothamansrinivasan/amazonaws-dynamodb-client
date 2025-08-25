

# TransactWriteItemsInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**transactItems** | [**List&lt;TransactWriteItem&gt;**](TransactWriteItem.md) | An ordered array of up to 100 &lt;code&gt;TransactWriteItem&lt;/code&gt; objects, each of which contains a &lt;code&gt;ConditionCheck&lt;/code&gt;, &lt;code&gt;Put&lt;/code&gt;, &lt;code&gt;Update&lt;/code&gt;, or &lt;code&gt;Delete&lt;/code&gt; object. These can operate on items in different tables, but the tables must reside in the same Amazon Web Services account and Region, and no two of them can operate on the same item.  |  |
|**returnConsumedCapacity** | **ReturnConsumedCapacity** |  |  [optional] |
|**returnItemCollectionMetrics** | **ReturnItemCollectionMetrics** | Determines whether item collection metrics are returned. If set to &lt;code&gt;SIZE&lt;/code&gt;, the response includes statistics about item collections (if any), that were modified during the operation and are returned in the response. If set to &lt;code&gt;NONE&lt;/code&gt; (the default), no statistics are returned.  |  [optional] |
|**clientRequestToken** | **String** | &lt;p&gt;Providing a &lt;code&gt;ClientRequestToken&lt;/code&gt; makes the call to &lt;code&gt;TransactWriteItems&lt;/code&gt; idempotent, meaning that multiple identical calls have the same effect as one single call.&lt;/p&gt; &lt;p&gt;Although multiple identical calls using the same client request token produce the same result on the server (no side effects), the responses to the calls might not be the same. If the &lt;code&gt;ReturnConsumedCapacity&lt;/code&gt; parameter is set, then the initial &lt;code&gt;TransactWriteItems&lt;/code&gt; call returns the amount of write capacity units consumed in making the changes. Subsequent &lt;code&gt;TransactWriteItems&lt;/code&gt; calls with the same client token return the number of read capacity units consumed in reading the item.&lt;/p&gt; &lt;p&gt;A client request token is valid for 10 minutes after the first request that uses it is completed. After 10 minutes, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 10 minutes, or the result might not be idempotent.&lt;/p&gt; &lt;p&gt;If you submit a request with the same client token but a change in other parameters within the 10-minute idempotency window, DynamoDB returns an &lt;code&gt;IdempotentParameterMismatch&lt;/code&gt; exception.&lt;/p&gt; |  [optional] |



