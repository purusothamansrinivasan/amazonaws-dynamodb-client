

# SourceTableDetails

Contains the details of the table when the backup was created. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableName** | **String** | The name of the table for which the backup was created.  |  |
|**tableId** | **String** | Unique identifier for the table for which the backup was created.  |  |
|**tableArn** | **String** | ARN of the table for which backup was created.  |  [optional] |
|**tableSizeBytes** | **Integer** | Size of the table in bytes. Note that this is an approximate value. |  [optional] |
|**keySchema** | [**List&lt;KeySchemaElement&gt;**](KeySchemaElement.md) | Schema of the table.  |  |
|**tableCreationDateTime** | **OffsetDateTime** | Time when the source table was created.  |  |
|**provisionedThroughput** | [**ProvisionedThroughput**](ProvisionedThroughput.md) | Read IOPs and Write IOPS on the table when the backup was created. |  |
|**itemCount** | **Integer** | Number of items in the table. Note that this is an approximate value.  |  [optional] |
|**billingMode** | **BillingMode** | &lt;p&gt;Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;PROVISIONED&lt;/code&gt; - Sets the read/write capacity mode to &lt;code&gt;PROVISIONED&lt;/code&gt;. We recommend using &lt;code&gt;PROVISIONED&lt;/code&gt; for predictable workloads.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; - Sets the read/write capacity mode to &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt;. We recommend using &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; for unpredictable workloads. &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |



