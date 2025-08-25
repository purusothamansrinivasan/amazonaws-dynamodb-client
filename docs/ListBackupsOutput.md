

# ListBackupsOutput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**backupSummaries** | [**List&lt;BackupSummary&gt;**](BackupSummary.md) | List of &lt;code&gt;BackupSummary&lt;/code&gt; objects. |  [optional] |
|**lastEvaluatedBackupArn** | **String** | &lt;p&gt; The ARN of the backup last evaluated when the current page of results was returned, inclusive of the current page of results. This value may be specified as the &lt;code&gt;ExclusiveStartBackupArn&lt;/code&gt; of a new &lt;code&gt;ListBackups&lt;/code&gt; operation in order to fetch the next page of results. &lt;/p&gt; &lt;p&gt; If &lt;code&gt;LastEvaluatedBackupArn&lt;/code&gt; is empty, then the last page of results has been processed and there are no more results to be retrieved. &lt;/p&gt; &lt;p&gt; If &lt;code&gt;LastEvaluatedBackupArn&lt;/code&gt; is not empty, this may or may not indicate that there is more data to be returned. All results are guaranteed to have been returned if and only if no value for &lt;code&gt;LastEvaluatedBackupArn&lt;/code&gt; is returned. &lt;/p&gt; |  [optional] |



