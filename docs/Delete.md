

# Delete

Represents a request to perform a <code>DeleteItem</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**key** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | The primary key of the item to be deleted. Each element consists of an attribute name and a value for that attribute. |  |
|**tableName** | **String** | Name of the table in which the item to be deleted resides. |  |
|**conditionExpression** | **String** | A condition that must be satisfied in order for a conditional delete to succeed. |  [optional] |
|**expressionAttributeNames** | **Map&lt;String, String&gt;** | One or more substitution tokens for attribute names in an expression. |  [optional] |
|**expressionAttributeValues** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | One or more values that can be substituted in an expression. |  [optional] |
|**returnValuesOnConditionCheckFailure** | **ReturnValuesOnConditionCheckFailure** | Use &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt; to get the item attributes if the &lt;code&gt;Delete&lt;/code&gt; condition fails. For &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt;, the valid values are: NONE and ALL_OLD. |  [optional] |



