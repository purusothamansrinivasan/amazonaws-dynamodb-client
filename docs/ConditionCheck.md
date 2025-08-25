

# ConditionCheck

Represents a request to perform a check that an item exists or to check the condition of specific attributes of the item.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**key** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | The primary key of the item to be checked. Each element consists of an attribute name and a value for that attribute. |  |
|**tableName** | **String** | Name of the table for the check item request. |  |
|**conditionExpression** | **String** | A condition that must be satisfied in order for a conditional update to succeed. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html\&quot;&gt;Condition expressions&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  |
|**expressionAttributeNames** | **Map&lt;String, String&gt;** | One or more substitution tokens for attribute names in an expression. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ExpressionAttributeNames.html\&quot;&gt;Expression attribute names&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  [optional] |
|**expressionAttributeValues** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | One or more values that can be substituted in an expression. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html\&quot;&gt;Condition expressions&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  [optional] |
|**returnValuesOnConditionCheckFailure** | **ReturnValuesOnConditionCheckFailure** | Use &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt; to get the item attributes if the &lt;code&gt;ConditionCheck&lt;/code&gt; condition fails. For &lt;code&gt;ReturnValuesOnConditionCheckFailure&lt;/code&gt;, the valid values are: NONE and ALL_OLD. |  [optional] |



