

# ListTablesOutput

Represents the output of a <code>ListTables</code> operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableNames** | **List&lt;String&gt;** | &lt;p&gt;The names of the tables associated with the current account at the current endpoint. The maximum size of this array is 100.&lt;/p&gt; &lt;p&gt;If &lt;code&gt;LastEvaluatedTableName&lt;/code&gt; also appears in the output, you can use this value as the &lt;code&gt;ExclusiveStartTableName&lt;/code&gt; parameter in a subsequent &lt;code&gt;ListTables&lt;/code&gt; request and obtain the next page of results.&lt;/p&gt; |  [optional] |
|**lastEvaluatedTableName** | **String** | &lt;p&gt;The name of the last table in the current page of results. Use this value as the &lt;code&gt;ExclusiveStartTableName&lt;/code&gt; in a new request to obtain the next page of results, until all the table names are returned.&lt;/p&gt; &lt;p&gt;If you do not receive a &lt;code&gt;LastEvaluatedTableName&lt;/code&gt; value in the response, this means that there are no more table names to be retrieved.&lt;/p&gt; |  [optional] |



