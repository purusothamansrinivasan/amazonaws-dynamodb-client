

# PutItemOutput

Represents the output of a <code>PutItem</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**attributes** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | The attribute values as they appeared before the &lt;code&gt;PutItem&lt;/code&gt; operation, but only if &lt;code&gt;ReturnValues&lt;/code&gt; is specified as &lt;code&gt;ALL_OLD&lt;/code&gt; in the request. Each element consists of an attribute name and an attribute value. |  [optional] |
|**consumedCapacity** | [**ConsumedCapacity**](ConsumedCapacity.md) | The capacity units consumed by the &lt;code&gt;PutItem&lt;/code&gt; operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. &lt;code&gt;ConsumedCapacity&lt;/code&gt; is only returned if the &lt;code&gt;ReturnConsumedCapacity&lt;/code&gt; parameter was specified. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html\&quot;&gt;Provisioned Throughput&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  [optional] |
|**itemCollectionMetrics** | [**ItemCollectionMetrics**](ItemCollectionMetrics.md) | &lt;p&gt;Information about item collections, if any, that were affected by the &lt;code&gt;PutItem&lt;/code&gt; operation. &lt;code&gt;ItemCollectionMetrics&lt;/code&gt; is only returned if the &lt;code&gt;ReturnItemCollectionMetrics&lt;/code&gt; parameter was specified. If the table does not have any local secondary indexes, this information is not returned in the response.&lt;/p&gt; &lt;p&gt;Each &lt;code&gt;ItemCollectionMetrics&lt;/code&gt; element consists of:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ItemCollectionKey&lt;/code&gt; - The partition key value of the item collection. This is the same as the partition key value of the item itself.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;SizeEstimateRangeGB&lt;/code&gt; - An estimate of item collection size, in gigabytes. This value is a two-element array containing a lower bound and an upper bound for the estimate. The estimate includes the size of all the items in the table, plus the size of all attributes projected into all of the local secondary indexes on that table. Use this estimate to measure whether a local secondary index is approaching its size limit.&lt;/p&gt; &lt;p&gt;The estimate is subject to change over time; therefore, do not rely on the precision or accuracy of the estimate.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |



