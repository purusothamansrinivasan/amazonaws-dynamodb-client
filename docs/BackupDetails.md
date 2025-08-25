

# BackupDetails

Contains the details of the backup created for the table.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**backupArn** | **String** | ARN associated with the backup. |  |
|**backupName** | **String** | Name of the requested backup. |  |
|**backupSizeBytes** | **Integer** | Size of the backup in bytes. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value. |  [optional] |
|**backupStatus** | **BackupStatus** | Backup can be in one of the following states: CREATING, ACTIVE, DELETED.  |  |
|**backupType** | **BackupType** | &lt;p&gt;BackupType:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;USER&lt;/code&gt; - You create and manage these using the on-demand backup feature.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;SYSTEM&lt;/code&gt; - If you delete a table with point-in-time recovery enabled, a &lt;code&gt;SYSTEM&lt;/code&gt; backup is automatically created and is retained for 35 days (at no additional cost). System backups allow you to restore the deleted table to the state it was in just before the point of deletion. &lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;AWS_BACKUP&lt;/code&gt; - On-demand backup created by you from Backup service.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  |
|**backupCreationDateTime** | **OffsetDateTime** | Time at which the backup was created. This is the request time of the backup.  |  |
|**backupExpiryDateTime** | **OffsetDateTime** | Time at which the automatic on-demand backup created by DynamoDB will expire. This &lt;code&gt;SYSTEM&lt;/code&gt; on-demand backup expires automatically 35 days after its creation. |  [optional] |



