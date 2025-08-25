

# CreateGlobalSecondaryIndexAction

Represents a new global secondary index to be added to an existing table.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**indexName** | **String** | The name of the global secondary index to be created. |  |
|**keySchema** | [**List&lt;KeySchemaElement&gt;**](KeySchemaElement.md) | The key schema for the global secondary index. |  |
|**projection** | [**Projection**](Projection.md) | Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. |  |
|**provisionedThroughput** | [**ProvisionedThroughput**](ProvisionedThroughput.md) | &lt;p&gt;Represents the provisioned throughput settings for the specified global secondary index.&lt;/p&gt; &lt;p&gt;For current minimum and maximum provisioned throughput values, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html\&quot;&gt;Service, Account, and Table Quotas&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;.&lt;/p&gt; |  [optional] |



