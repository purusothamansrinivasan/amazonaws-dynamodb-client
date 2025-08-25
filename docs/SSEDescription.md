

# SSEDescription

The description of the server-side encryption status on the specified table.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**status** | **SSEStatus** | &lt;p&gt;Represents the current state of server-side encryption. The only supported values are:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ENABLED&lt;/code&gt; - Server-side encryption is enabled.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - Server-side encryption is being updated.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**ssEType** | **SSEType** | &lt;p&gt;Server-side encryption type. The only supported value is:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;KMS&lt;/code&gt; - Server-side encryption that uses Key Management Service. The key is stored in your account and is managed by KMS (KMS charges apply).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**kmSMasterKeyArn** | **String** | The KMS key ARN used for the KMS encryption. |  [optional] |
|**inaccessibleEncryptionDateTime** | **OffsetDateTime** | Indicates the time, in UNIX epoch date format, when DynamoDB detected that the table&#39;s KMS key was inaccessible. This attribute will automatically be cleared when DynamoDB detects that the table&#39;s KMS key is accessible again. DynamoDB will initiate the table archival process when table&#39;s KMS key remains inaccessible for more than seven days from this date. |  [optional] |



