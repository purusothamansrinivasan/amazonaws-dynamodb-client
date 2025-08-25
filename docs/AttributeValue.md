

# AttributeValue

<p>Represents the data for an attribute.</p> <p>Each attribute value is described as a name-value pair. The name is the data type, and the value is the data itself.</p> <p>For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes\">Data Types</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**S** | **String** | &lt;p&gt;An attribute of type String. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;S\&quot;: \&quot;Hello\&quot;&lt;/code&gt; &lt;/p&gt; |  [optional] |
|**N** | **String** | &lt;p&gt;An attribute of type Number. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;N\&quot;: \&quot;123.45\&quot;&lt;/code&gt; &lt;/p&gt; &lt;p&gt;Numbers are sent across the network to DynamoDB as strings, to maximize compatibility across languages and libraries. However, DynamoDB treats them as number type attributes for mathematical operations.&lt;/p&gt; |  [optional] |
|**B** | **String** | &lt;p&gt;An attribute of type Binary. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;B\&quot;: \&quot;dGhpcyB0ZXh0IGlzIGJhc2U2NC1lbmNvZGVk\&quot;&lt;/code&gt; &lt;/p&gt; |  [optional] |
|**SS** | **List&lt;String&gt;** | &lt;p&gt;An attribute of type String Set. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;SS\&quot;: [\&quot;Giraffe\&quot;, \&quot;Hippo\&quot; ,\&quot;Zebra\&quot;]&lt;/code&gt; &lt;/p&gt; |  [optional] |
|**NS** | **List&lt;String&gt;** | &lt;p&gt;An attribute of type Number Set. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;NS\&quot;: [\&quot;42.2\&quot;, \&quot;-19\&quot;, \&quot;7.5\&quot;, \&quot;3.14\&quot;]&lt;/code&gt; &lt;/p&gt; &lt;p&gt;Numbers are sent across the network to DynamoDB as strings, to maximize compatibility across languages and libraries. However, DynamoDB treats them as number type attributes for mathematical operations.&lt;/p&gt; |  [optional] |
|**BS** | **List&lt;String&gt;** | &lt;p&gt;An attribute of type Binary Set. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;BS\&quot;: [\&quot;U3Vubnk&#x3D;\&quot;, \&quot;UmFpbnk&#x3D;\&quot;, \&quot;U25vd3k&#x3D;\&quot;]&lt;/code&gt; &lt;/p&gt; |  [optional] |
|**M** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) | &lt;p&gt;An attribute of type Map. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;M\&quot;: {\&quot;Name\&quot;: {\&quot;S\&quot;: \&quot;Joe\&quot;}, \&quot;Age\&quot;: {\&quot;N\&quot;: \&quot;35\&quot;}}&lt;/code&gt; &lt;/p&gt; |  [optional] |
|**L** | [**List&lt;AttributeValue&gt;**](AttributeValue.md) | &lt;p&gt;An attribute of type List. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;L\&quot;: [ {\&quot;S\&quot;: \&quot;Cookies\&quot;} , {\&quot;S\&quot;: \&quot;Coffee\&quot;}, {\&quot;N\&quot;: \&quot;3.14159\&quot;}]&lt;/code&gt; &lt;/p&gt; |  [optional] |
|**NULL** | **Boolean** | &lt;p&gt;An attribute of type Null. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;NULL\&quot;: true&lt;/code&gt; &lt;/p&gt; |  [optional] |
|**BOOL** | **Boolean** | &lt;p&gt;An attribute of type Boolean. For example:&lt;/p&gt; &lt;p&gt; &lt;code&gt;\&quot;BOOL\&quot;: true&lt;/code&gt; &lt;/p&gt; |  [optional] |



