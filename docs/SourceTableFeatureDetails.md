

# SourceTableFeatureDetails

Contains the details of the features enabled on the table when the backup was created. For example, LSIs, GSIs, streams, TTL. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**localSecondaryIndexes** | [**List&lt;LocalSecondaryIndexInfo&gt;**](LocalSecondaryIndexInfo.md) | Represents the LSI properties for the table when the backup was created. It includes the IndexName, KeySchema and Projection for the LSIs on the table at the time of backup.  |  [optional] |
|**globalSecondaryIndexes** | [**List&lt;GlobalSecondaryIndexInfo&gt;**](GlobalSecondaryIndexInfo.md) | Represents the GSI properties for the table when the backup was created. It includes the IndexName, KeySchema, Projection, and ProvisionedThroughput for the GSIs on the table at the time of backup.  |  [optional] |
|**streamDescription** | [**StreamSpecification**](StreamSpecification.md) | Stream settings on the table when the backup was created. |  [optional] |
|**timeToLiveDescription** | [**TimeToLiveDescription**](TimeToLiveDescription.md) | Time to Live settings on the table when the backup was created. |  [optional] |
|**ssEDescription** | [**SSEDescription**](SSEDescription.md) | The description of the server-side encryption status on the table when the backup was created. |  [optional] |



