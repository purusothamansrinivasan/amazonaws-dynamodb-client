

# BatchStatementError

 An error associated with a statement in a PartiQL batch that was run. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**code** | **BatchStatementErrorCodeEnum** |  The error code associated with the failed PartiQL batch statement.  |  [optional] |
|**message** | **String** |  The error message associated with the PartiQL batch response.  |  [optional] |
|**item** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | The item which caused the condition check to fail. This will be set if ReturnValuesOnConditionCheckFailure is specified as &lt;code&gt;ALL_OLD&lt;/code&gt;. |  [optional] |



