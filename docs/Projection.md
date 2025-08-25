

# Projection

Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**projectionType** | **ProjectionType** | &lt;p&gt;The set of attributes that are projected into the index:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;KEYS_ONLY&lt;/code&gt; - Only the index and primary keys are projected into the index.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;INCLUDE&lt;/code&gt; - In addition to the attributes described in &lt;code&gt;KEYS_ONLY&lt;/code&gt;, the secondary index will include other non-key attributes that you specify.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ALL&lt;/code&gt; - All of the table attributes are projected into the index.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**nonKeyAttributes** | **List&lt;String&gt;** | &lt;p&gt;Represents the non-key attribute names which will be projected into the index.&lt;/p&gt; &lt;p&gt;For local secondary indexes, the total count of &lt;code&gt;NonKeyAttributes&lt;/code&gt; summed across all of the local secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.&lt;/p&gt; |  [optional] |



