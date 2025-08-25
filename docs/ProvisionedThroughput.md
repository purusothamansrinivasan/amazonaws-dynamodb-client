

# ProvisionedThroughput

<p>Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the <code>UpdateTable</code> operation.</p> <p>For current minimum and maximum provisioned throughput values, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html\">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**readCapacityUnits** | **Integer** | &lt;p&gt;The maximum number of strongly consistent reads consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html\&quot;&gt;Specifying Read and Write Requirements&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;.&lt;/p&gt; &lt;p&gt;If read/write capacity mode is &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; the value is set to 0.&lt;/p&gt; |  |
|**writeCapacityUnits** | **Integer** | &lt;p&gt;The maximum number of writes consumed per second before DynamoDB returns a &lt;code&gt;ThrottlingException&lt;/code&gt;. For more information, see &lt;a href&#x3D;\&quot;https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html\&quot;&gt;Specifying Read and Write Requirements&lt;/a&gt; in the &lt;i&gt;Amazon DynamoDB Developer Guide&lt;/i&gt;.&lt;/p&gt; &lt;p&gt;If read/write capacity mode is &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; the value is set to 0.&lt;/p&gt; |  |



