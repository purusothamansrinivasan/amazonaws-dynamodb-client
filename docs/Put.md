

# Put

Represents a request to perform a <code>PutItem</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**item** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | A map of attribute name to attribute values, representing the primary key of the item to be written by &lt;code&gt;PutItem&lt;/code&gt;. All of the table&#39;s primary key attributes must be specified, and their data types must match those of the table&#39;s key schema. If any attributes are present in the item that are part of an index key schema for the table, their types must match the index key schema.  |  |
|**tableName** | **String** | Name of the table in which to write the item. |  |
|**conditionExpression** | **String** | A condition that must be satisfied in order for a conditional update to succeed. |  [optional] |
|**expressionAttributeNames** | **Map&lt;String, String&gt;** | One or more substitution tokens for attribute names in an expression. |  [optional] |
|**expressionAttributeValues** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | One or more values that can be substituted in an expression. |  [optional] |
|**returnValuesOnConditionCheckFailure** | **ReturnValuesOnConditionCheckFailure** | Use &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt; to get the item attributes if the &lt;code&gt;Put&lt;/code&gt; condition fails. For &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt;, the valid values are: NONE and ALL_OLD. |  [optional] |



