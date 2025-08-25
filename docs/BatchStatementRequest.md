

# BatchStatementRequest

 A PartiQL batch statement request. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**statement** | **String** |  A valid PartiQL statement.  |  |
|**parameters** | [**List&lt;AttributeValue&gt;**](AttributeValue.md) |  The parameters associated with a PartiQL statement in the batch request.  |  [optional] |
|**consistentRead** | **Boolean** |  The read consistency of the PartiQL batch request.  |  [optional] |
|**returnValuesOnConditionCheckFailure** | **ReturnValuesOnConditionCheckFailure** | &lt;p&gt;An optional parameter that returns the item attributes for a PartiQL batch request operation that failed a condition check.&lt;/p&gt; &lt;p&gt;There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.&lt;/p&gt; |  [optional] |



