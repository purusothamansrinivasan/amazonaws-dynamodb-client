

# BillingModeSummary

<p>Contains the details for the read/write capacity mode. This page talks about <code>PROVISIONED</code> and <code>PAY_PER_REQUEST</code> billing modes. For more information about these modes, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html\">Read/write capacity mode</a>.</p> <note> <p>You may need to switch to on-demand mode at least once in order to return a <code>BillingModeSummary</code> response.</p> </note>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**billingMode** | **BillingMode** | &lt;p&gt;Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;PROVISIONED&lt;/code&gt; - Sets the read/write capacity mode to &lt;code&gt;PROVISIONED&lt;/code&gt;. We recommend using &lt;code&gt;PROVISIONED&lt;/code&gt; for predictable workloads.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; - Sets the read/write capacity mode to &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt;. We recommend using &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; for unpredictable workloads. &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**lastUpdateToPayPerRequestDateTime** | **OffsetDateTime** | Represents the time when &lt;code&gt;PAY_PER_REQUEST&lt;/code&gt; was last set as the read/write capacity mode. |  [optional] |



