

# ConsumedCapacity

The capacity units consumed by an operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the request asked for it. For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html\">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableName** | **String** | The name of the table that was affected by the operation. |  [optional] |
|**capacityUnits** | **Double** | The total number of capacity units consumed by the operation. |  [optional] |
|**readCapacityUnits** | **Double** | The total number of read capacity units consumed by the operation. |  [optional] |
|**writeCapacityUnits** | **Double** | The total number of write capacity units consumed by the operation. |  [optional] |
|**table** | [**Capacity**](Capacity.md) | The amount of throughput consumed on the table affected by the operation. |  [optional] |
|**localSecondaryIndexes** | [**Map&lt;String, Capacity&gt;**](Capacity.md) | The amount of throughput consumed on each local index affected by the operation. |  [optional] |
|**globalSecondaryIndexes** | [**Map&lt;String, Capacity&gt;**](Capacity.md) | The amount of throughput consumed on each global index affected by the operation. |  [optional] |



