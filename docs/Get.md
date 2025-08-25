

# Get

Specifies an item and related attribute values to retrieve in a <code>TransactGetItem</code> object.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**key** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | A map of attribute names to &lt;code&gt;AttributeValue&lt;/code&gt; objects that specifies the primary key of the item to retrieve. |  |
|**tableName** | **String** | The name of the table from which to retrieve the specified item. |  |
|**projectionExpression** | **String** | A string that identifies one or more attributes of the specified item to retrieve from the table. The attributes in the expression must be separated by commas. If no attribute names are specified, then all attributes of the specified item are returned. If any of the requested attributes are not found, they do not appear in the result. |  [optional] |
|**expressionAttributeNames** | **Map&lt;String, String&gt;** | One or more substitution tokens for attribute names in the ProjectionExpression parameter. |  [optional] |



