

# GetItemOutput

Represents the output of a <code>GetItem</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**item** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | A map of attribute names to &lt;code&gt;AttributeValue&lt;/code&gt; objects, as specified by &lt;code&gt;ProjectionExpression&lt;/code&gt;. |  [optional] |
|**consumedCapacity** | [**ConsumedCapacity**](ConsumedCapacity.md) | The capacity units consumed by the &lt;code&gt;GetItem&lt;/code&gt; operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. &lt;code&gt;ConsumedCapacity&lt;/code&gt; is only returned if the &lt;code&gt;ReturnConsumedCapacity&lt;/code&gt; parameter was specified. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html#ItemSizeCalculations.Reads\&quot;&gt;Provisioned Throughput&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  [optional] |



