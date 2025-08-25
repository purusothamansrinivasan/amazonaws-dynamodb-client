

# ReplicaDescription

Contains the details of the replica.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**regionName** | **String** | The name of the Region. |  [optional] |
|**replicaStatus** | **ReplicaStatus** | &lt;p&gt;The current state of the replica:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;CREATING&lt;/code&gt; - The replica is being created.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;UPDATING&lt;/code&gt; - The replica is being updated.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DELETING&lt;/code&gt; - The replica is being deleted.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ACTIVE&lt;/code&gt; - The replica is ready for use.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;REGION_DISABLED&lt;/code&gt; - The replica is inaccessible because the Amazon Web Services Region has been disabled.&lt;/p&gt; &lt;note&gt; &lt;p&gt;If the Amazon Web Services Region remains inaccessible for more than 20 hours, DynamoDB will remove this replica from the replication group. The replica will not be deleted and replication will stop from and to this region.&lt;/p&gt; &lt;/note&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;INACCESSIBLE_ENCRYPTION_CREDENTIALS &lt;/code&gt; - The KMS key used to encrypt the table is inaccessible.&lt;/p&gt; &lt;note&gt; &lt;p&gt;If the KMS key remains inaccessible for more than 20 hours, DynamoDB will remove this replica from the replication group. The replica will not be deleted and replication will stop from and to this region.&lt;/p&gt; &lt;/note&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**replicaStatusDescription** | **String** | Detailed information about the replica status. |  [optional] |
|**replicaStatusPercentProgress** | **String** | Specifies the progress of a Create, Update, or Delete action on the replica as a percentage. |  [optional] |
|**kmSMasterKeyId** | **String** | The KMS key of the replica that will be used for KMS encryption. |  [optional] |
|**provisionedThroughputOverride** | [**ProvisionedThroughputOverride**](ProvisionedThroughputOverride.md) | Replica-specific provisioned throughput. If not described, uses the source table&#39;s provisioned throughput settings. |  [optional] |
|**globalSecondaryIndexes** | [**List&lt;ReplicaGlobalSecondaryIndexDescription&gt;**](ReplicaGlobalSecondaryIndexDescription.md) | Replica-specific global secondary index settings. |  [optional] |
|**replicaInaccessibleDateTime** | **OffsetDateTime** | The time at which the replica was first detected as inaccessible. To determine cause of inaccessibility check the &lt;code&gt;ReplicaStatus&lt;/code&gt; property. |  [optional] |
|**replicaTableClassSummary** | [**TableClassSummary**](TableClassSummary.md) |  |  [optional] |



