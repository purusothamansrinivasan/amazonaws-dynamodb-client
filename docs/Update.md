

# Update

Represents a request to perform an <code>UpdateItem</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**key** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | The primary key of the item to be updated. Each element consists of an attribute name and a value for that attribute. |  |
|**updateExpression** | **String** | An expression that defines one or more attributes to be updated, the action to be performed on them, and new value(s) for them. |  |
|**tableName** | **String** | Name of the table for the &lt;code&gt;UpdateItem&lt;/code&gt; request. |  |
|**conditionExpression** | **String** | A condition that must be satisfied in order for a conditional update to succeed. |  [optional] |
|**expressionAttributeNames** | **Map&lt;String, String&gt;** | One or more substitution tokens for attribute names in an expression. |  [optional] |
|**expressionAttributeValues** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | One or more values that can be substituted in an expression. |  [optional] |
|**returnValuesOnConditionCheckFailure** | **ReturnValuesOnConditionCheckFailure** | Use &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt; to get the item attributes if the &lt;code&gt;Update&lt;/code&gt; condition fails. For &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt;, the valid values are: NONE and ALL_OLD. |  [optional] |



