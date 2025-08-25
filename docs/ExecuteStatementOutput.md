

# ExecuteStatementOutput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**items** | **List&lt;Map&lt;String, AttributeValue&gt;&gt;** | If a read operation was used, this property will contain the result of the read operation; a map of attribute names and their values. For the write operations this value will be empty. |  [optional] |
|**nextToken** | **String** | If the response of a read request exceeds the response payload limit DynamoDB will set this value in the response. If set, you can use that this value in the subsequent request to get the remaining results. |  [optional] |
|**consumedCapacity** | [**ConsumedCapacity**](ConsumedCapacity.md) |  |  [optional] |
|**lastEvaluatedKey** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | The primary key of the item where the operation stopped, inclusive of the previous result set. Use this value to start a new operation, excluding this value in the new request. If &lt;code&gt;LastEvaluatedKey&lt;/code&gt; is empty, then the \&quot;last page\&quot; of results has been processed and there is no more data to be retrieved. If &lt;code&gt;LastEvaluatedKey&lt;/code&gt; is not empty, it does not necessarily mean that there is more data in the result set. The only way to know when you have reached the end of the result set is when &lt;code&gt;LastEvaluatedKey&lt;/code&gt; is empty.  |  [optional] |



