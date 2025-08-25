

# ParameterizedStatement

 Represents a PartiQL statment that uses parameters. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**statement** | **String** |  A PartiQL statment that uses parameters.  |  |
|**parameters** | [**List&lt;AttributeValue&gt;**](AttributeValue.md) |  The parameter values.  |  [optional] |
|**returnValuesOnConditionCheckFailure** | **ReturnValuesOnConditionCheckFailure** | &lt;p&gt;An optional parameter that returns the item attributes for a PartiQL &lt;code&gt;ParameterizedStatement&lt;/code&gt; operation that failed a condition check.&lt;/p&gt; &lt;p&gt;There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.&lt;/p&gt; |  [optional] |



