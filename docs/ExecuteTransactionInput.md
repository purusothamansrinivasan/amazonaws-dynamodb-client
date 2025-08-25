

# ExecuteTransactionInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**transactStatements** | [**List&lt;ParameterizedStatement&gt;**](ParameterizedStatement.md) | The list of PartiQL statements representing the transaction to run. |  |
|**clientRequestToken** | **String** | Set this value to get remaining results, if &lt;code&gt;NextToken&lt;/code&gt; was returned in the statement response. |  [optional] |
|**returnConsumedCapacity** | **ReturnConsumedCapacity** | Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TransactGetItems.html\&quot;&gt;TransactGetItems&lt;/a&gt; and &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TransactWriteItems.html\&quot;&gt;TransactWriteItems&lt;/a&gt;. |  [optional] |



