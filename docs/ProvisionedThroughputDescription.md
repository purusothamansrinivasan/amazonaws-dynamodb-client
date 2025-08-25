

# ProvisionedThroughputDescription

Represents the provisioned throughput settings for the table, consisting of read and write capacity units, along with data about increases and decreases.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**lastIncreaseDateTime** | **OffsetDateTime** | The date and time of the last provisioned throughput increase for this table. |  [optional] |
|**lastDecreaseDateTime** | **OffsetDateTime** | The date and time of the last provisioned throughput decrease for this table. |  [optional] |
|**numberOfDecreasesToday** | **Integer** | The number of provisioned throughput decreases for this table during this UTC calendar day. For current maximums on provisioned throughput decreases, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html\&quot;&gt;Service, Account, and Table Quotas&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;. |  [optional] |
|**readCapacityUnits** | **Integer** | The maximum number of strongly consistent reads consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. Eventually consistent reads require less effort than strongly consistent reads, so a setting of 50 &lt;code&gt;ReadCapacityUnits&lt;/code&gt; per second provides 100 eventually consistent &lt;code&gt;ReadCapacityUnits&lt;/code&gt; per second. |  [optional] |
|**writeCapacityUnits** | **Integer** | The maximum number of writes consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. |  [optional] |



