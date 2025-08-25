

# PutRequest

Represents a request to perform a <code>PutItem</code> operation on an item.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**item** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | A map of attribute name to attribute values, representing the primary key of an item to be processed by &lt;code&gt;PutItem&lt;/code&gt;. All of the table&#39;s primary key attributes must be specified, and their data types must match those of the table&#39;s key schema. If any attributes are present in the item that are part of an index key schema for the table, their types must match the index key schema. |  |



