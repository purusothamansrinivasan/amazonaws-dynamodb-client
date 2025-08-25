

# TransactWriteItem

A list of requests that can perform update, put, delete, or check operations on multiple items in one or more tables atomically.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**conditionCheck** | [**ConditionCheck**](ConditionCheck.md) | A request to perform a check item operation. |  [optional] |
|**put** | [**Put**](Put.md) | A request to perform a &lt;code&gt;PutItem&lt;/code&gt; operation. |  [optional] |
|**delete** | [**Delete**](Delete.md) | A request to perform a &lt;code&gt;DeleteItem&lt;/code&gt; operation. |  [optional] |
|**update** | [**Update**](Update.md) | A request to perform an &lt;code&gt;UpdateItem&lt;/code&gt; operation. |  [optional] |



