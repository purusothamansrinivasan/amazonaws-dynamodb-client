

# ListGlobalTablesInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**exclusiveStartGlobalTableName** | **String** | The first global table name that this operation will evaluate. |  [optional] |
|**limit** | **Integer** | &lt;p&gt;The maximum number of table names to return, if the parameter is not specified DynamoDB defaults to 100.&lt;/p&gt; &lt;p&gt;If the number of global tables DynamoDB finds reaches this limit, it stops the operation and returns the table names collected up to that point, with a table name in the &lt;code&gt;LastEvaluatedGlobalTableName&lt;/code&gt; to apply in a subsequent operation to the &lt;code&gt;ExclusiveStartGlobalTableName&lt;/code&gt; parameter.&lt;/p&gt; |  [optional] |
|**regionName** | **String** | Lists the global tables in a specific Region. |  [optional] |



