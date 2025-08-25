

# ArchivalSummary

Contains details of a table archival operation.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**archivalDateTime** | **OffsetDateTime** | The date and time when table archival was initiated by DynamoDB, in UNIX epoch time format. |  [optional] |
|**archivalReason** | **String** | &lt;p&gt;The reason DynamoDB archived the table. Currently, the only possible value is:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;INACCESSIBLE_ENCRYPTION_CREDENTIALS&lt;/code&gt; - The table was archived due to the table&#39;s KMS key being inaccessible for more than seven days. An On-Demand backup was created at the archival time.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**archivalBackupArn** | **String** | The Amazon Resource Name (ARN) of the backup the table was archived to, when applicable in the archival reason. If you wish to restore this backup to the same table name, you will need to delete the original table. |  [optional] |



