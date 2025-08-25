

# StreamSpecification

Represents the DynamoDB Streams configuration for a table in DynamoDB.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**streamEnabled** | **Boolean** | Indicates whether DynamoDB Streams is enabled (true) or disabled (false) on the table. |  |
|**streamViewType** | **StreamViewType** | &lt;p&gt; When an item in the table is modified, &lt;code&gt;StreamViewType&lt;/code&gt; determines what information is written to the stream for this table. Valid values for &lt;code&gt;StreamViewType&lt;/code&gt; are:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;KEYS_ONLY&lt;/code&gt; - Only the key attributes of the modified item are written to the stream.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;NEW_IMAGE&lt;/code&gt; - The entire item, as it appears after it was modified, is written to the stream.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;OLD_IMAGE&lt;/code&gt; - The entire item, as it appeared before it was modified, is written to the stream.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;NEW_AND_OLD_IMAGES&lt;/code&gt; - Both the new and the old item images of the item are written to the stream.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |



