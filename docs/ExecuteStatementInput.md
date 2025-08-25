

# ExecuteStatementInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**statement** | **String** | The PartiQL statement representing the operation to run. |  |
|**parameters** | [**List&lt;AttributeValue&gt;**](AttributeValue.md) | The parameters for the PartiQL statement, if any. |  [optional] |
|**consistentRead** | **Boolean** | The consistency of a read operation. If set to &lt;code&gt;true&lt;/code&gt;, then a strongly consistent read is used; otherwise, an eventually consistent read is used. |  [optional] |
|**nextToken** | **String** | Set this value to get remaining results, if &lt;code&gt;NextToken&lt;/code&gt; was returned in the statement response. |  [optional] |
|**returnConsumedCapacity** | **ReturnConsumedCapacity** |  |  [optional] |
|**limit** | **Integer** | The maximum number of items to evaluate (not necessarily the number of matching items). If DynamoDB processes the number of items up to the limit while processing the results, it stops the operation and returns the matching values up to that point, along with a key in &lt;code&gt;LastEvaluatedKey&lt;/code&gt; to apply in a subsequent operation so you can pick up where you left off. Also, if the processed dataset size exceeds 1 MB before DynamoDB reaches this limit, it stops the operation and returns the matching values up to the limit, and a key in &lt;code&gt;LastEvaluatedKey&lt;/code&gt; to apply in a subsequent operation to continue the operation.  |  [optional] |
|**returnValuesOnConditionCheckFailure** | **ReturnValuesOnConditionCheckFailure** | &lt;p&gt;An optional parameter that returns the item attributes for an &lt;code&gt;ExecuteStatement&lt;/code&gt; operation that failed a condition check.&lt;/p&gt; &lt;p&gt;There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.&lt;/p&gt; |  [optional] |



