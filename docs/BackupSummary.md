

# BackupSummary

Contains details for the backup.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableName** | **String** | Name of the table. |  [optional] |
|**tableId** | **String** | Unique identifier for the table. |  [optional] |
|**tableArn** | **String** | ARN associated with the table. |  [optional] |
|**backupArn** | **String** | ARN associated with the backup. |  [optional] |
|**backupName** | **String** | Name of the specified backup. |  [optional] |
|**backupCreationDateTime** | **OffsetDateTime** | Time at which the backup was created. |  [optional] |
|**backupExpiryDateTime** | **OffsetDateTime** | Time at which the automatic on-demand backup created by DynamoDB will expire. This &lt;code&gt;SYSTEM&lt;/code&gt; on-demand backup expires automatically 35 days after its creation. |  [optional] |
|**backupStatus** | **BackupStatus** | Backup can be in one of the following states: CREATING, ACTIVE, DELETED. |  [optional] |
|**backupType** | **BackupType** | &lt;p&gt;BackupType:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;USER&lt;/code&gt; - You create and manage these using the on-demand backup feature.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;SYSTEM&lt;/code&gt; - If you delete a table with point-in-time recovery enabled, a &lt;code&gt;SYSTEM&lt;/code&gt; backup is automatically created and is retained for 35 days (at no additional cost). System backups allow you to restore the deleted table to the state it was in just before the point of deletion. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AWS_BACKUP&lt;/code&gt; - On-demand backup created by you from Backup service.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**backupSizeBytes** | **Integer** | Size of the backup in bytes. |  [optional] |



