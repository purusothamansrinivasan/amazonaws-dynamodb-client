

# SSESpecification

Represents the settings used to enable server-side encryption.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**enabled** | **Boolean** | Indicates whether server-side encryption is done using an Amazon Web Services managed key or an Amazon Web Services owned key. If enabled (true), server-side encryption type is set to &lt;code&gt;KMS&lt;/code&gt; and an Amazon Web Services managed key is used (KMS charges apply). If disabled (false) or not specified, server-side encryption is set to Amazon Web Services owned key. |  [optional] |
|**ssEType** | **SSEType** | &lt;p&gt;Server-side encryption type. The only supported value is:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;KMS&lt;/code&gt; - Server-side encryption that uses Key Management Service. The key is stored in your account and is managed by KMS (KMS charges apply).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**kmSMasterKeyId** | **String** | The KMS key that should be used for the KMS encryption. To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key &lt;code&gt;alias/aws/dynamodb&lt;/code&gt;. |  [optional] |



