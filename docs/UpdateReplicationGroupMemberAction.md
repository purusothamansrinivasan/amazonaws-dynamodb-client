

# UpdateReplicationGroupMemberAction

Represents a replica to be modified.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**regionName** | **String** | The Region where the replica exists. |  |
|**kmSMasterKeyId** | **String** | The KMS key of the replica that should be used for KMS encryption. To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB KMS key &lt;code&gt;alias/aws/dynamodb&lt;/code&gt;. |  [optional] |
|**provisionedThroughputOverride** | [**ProvisionedThroughputOverride**](ProvisionedThroughputOverride.md) | Replica-specific provisioned throughput. If not specified, uses the source table&#39;s provisioned throughput settings. |  [optional] |
|**globalSecondaryIndexes** | [**List&lt;ReplicaGlobalSecondaryIndex&gt;**](ReplicaGlobalSecondaryIndex.md) | Replica-specific global secondary index settings. |  [optional] |
|**tableClassOverride** | **TableClass** | Replica-specific table class. If not specified, uses the source table&#39;s table class. |  [optional] |



