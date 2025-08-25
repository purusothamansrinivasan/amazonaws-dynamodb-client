

# ScanOutput

Represents the output of a <code>Scan</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**items** | **List&lt;Map&lt;String, AttributeValue&gt;&gt;** | An array of item attributes that match the scan criteria. Each element in this array consists of an attribute name and the value for that attribute. |  [optional] |
|**count** | **Integer** | &lt;p&gt;The number of items in the response.&lt;/p&gt; &lt;p&gt;If you set &lt;code&gt;ScanFilter&lt;/code&gt; in the request, then &lt;code&gt;Count&lt;/code&gt; is the number of items returned after the filter was applied, and &lt;code&gt;ScannedCount&lt;/code&gt; is the number of matching items before the filter was applied.&lt;/p&gt; &lt;p&gt;If you did not use a filter in the request, then &lt;code&gt;Count&lt;/code&gt; is the same as &lt;code&gt;ScannedCount&lt;/code&gt;.&lt;/p&gt; |  [optional] |
|**scannedCount** | **Integer** | &lt;p&gt;The number of items evaluated, before any &lt;code&gt;ScanFilter&lt;/code&gt; is applied. A high &lt;code&gt;ScannedCount&lt;/code&gt; value with few, or no, &lt;code&gt;Count&lt;/code&gt; results indicates an inefficient &lt;code&gt;Scan&lt;/code&gt; operation. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#Count\&quot;&gt;Count and ScannedCount&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;.&lt;/p&gt; &lt;p&gt;If you did not use a filter in the request, then &lt;code&gt;ScannedCount&lt;/code&gt; is the same as &lt;code&gt;Count&lt;/code&gt;.&lt;/p&gt; |  [optional] |
|**lastEvaluatedKey** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | &lt;p&gt;The primary key of the item where the operation stopped, inclusive of the previous result set. Use this value to start a new operation, excluding this value in the new request.&lt;/p&gt; &lt;p&gt;If &lt;code&gt;LastEvaluatedKey&lt;/code&gt; is empty, then the \&quot;last page\&quot; of results has been processed and there is no more data to be retrieved.&lt;/p&gt; &lt;p&gt;If &lt;code&gt;LastEvaluatedKey&lt;/code&gt; is not empty, it does not necessarily mean that there is more data in the result set. The only way to know when you have reached the end of the result set is when &lt;code&gt;LastEvaluatedKey&lt;/code&gt; is empty.&lt;/p&gt; |  [optional] |
|**consumedCapacity** | [**ConsumedCapacity**](ConsumedCapacity.md) | The capacity units consumed by the &lt;code&gt;Scan&lt;/code&gt; operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. &lt;code&gt;ConsumedCapacity&lt;/code&gt; is only returned if the &lt;code&gt;ReturnConsumedCapacity&lt;/code&gt; parameter was specified. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html#ItemSizeCalculations.Reads\&quot;&gt;Provisioned Throughput&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  [optional] |



