

# GlobalSecondaryIndexUpdate

<p>Represents one of the following:</p> <ul> <li> <p>A new global secondary index to be added to an existing table.</p> </li> <li> <p>New provisioned throughput parameters for an existing global secondary index.</p> </li> <li> <p>An existing global secondary index to be removed from an existing table.</p> </li> </ul>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**update** | [**UpdateGlobalSecondaryIndexAction**](UpdateGlobalSecondaryIndexAction.md) | The name of an existing global secondary index, along with new provisioned throughput settings to be applied to that index. |  [optional] |
|**create** | [**CreateGlobalSecondaryIndexAction**](CreateGlobalSecondaryIndexAction.md) | &lt;p&gt;The parameters required for creating a global secondary index on an existing table:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;IndexName &lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;KeySchema &lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AttributeDefinitions &lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;Projection &lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ProvisionedThroughput &lt;/code&gt; &lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**delete** | [**DeleteGlobalSecondaryIndexAction**](DeleteGlobalSecondaryIndexAction.md) | The name of an existing global secondary index to be removed. |  [optional] |



