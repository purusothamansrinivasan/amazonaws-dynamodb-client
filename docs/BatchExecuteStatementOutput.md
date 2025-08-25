

# BatchExecuteStatementOutput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**responses** | [**List&lt;BatchStatementResponse&gt;**](BatchStatementResponse.md) | The response to each PartiQL statement in the batch. The values of the list are ordered according to the ordering of the request statements. |  [optional] |
|**consumedCapacity** | [**List&lt;ConsumedCapacity&gt;**](ConsumedCapacity.md) | The capacity units consumed by the entire operation. The values of the list are ordered according to the ordering of the statements. |  [optional] |



