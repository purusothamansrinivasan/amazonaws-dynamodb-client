

# ListBackupsInput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableName** | **String** | The backups from the table specified by &lt;code&gt;TableName&lt;/code&gt; are listed.  |  [optional] |
|**limit** | **Integer** | Maximum number of backups to return at once. |  [optional] |
|**timeRangeLowerBound** | **OffsetDateTime** | Only backups created after this time are listed. &lt;code&gt;TimeRangeLowerBound&lt;/code&gt; is inclusive. |  [optional] |
|**timeRangeUpperBound** | **OffsetDateTime** | Only backups created before this time are listed. &lt;code&gt;TimeRangeUpperBound&lt;/code&gt; is exclusive.  |  [optional] |
|**exclusiveStartBackupArn** | **String** |  &lt;code&gt;LastEvaluatedBackupArn&lt;/code&gt; is the Amazon Resource Name (ARN) of the backup last evaluated when the current page of results was returned, inclusive of the current page of results. This value may be specified as the &lt;code&gt;ExclusiveStartBackupArn&lt;/code&gt; of a new &lt;code&gt;ListBackups&lt;/code&gt; operation in order to fetch the next page of results.  |  [optional] |
|**backupType** | **BackupTypeFilter** | &lt;p&gt;The backups from the table specified by &lt;code&gt;BackupType&lt;/code&gt; are listed.&lt;/p&gt; &lt;p&gt;Where &lt;code&gt;BackupType&lt;/code&gt; can be:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;USER&lt;/code&gt; - On-demand backup created by you. (The default setting if no other backup types are specified.)&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;SYSTEM&lt;/code&gt; - On-demand backup automatically created by DynamoDB.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ALL&lt;/code&gt; - All types of on-demand backups (USER and SYSTEM).&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |



