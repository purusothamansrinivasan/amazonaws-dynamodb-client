

# BatchStatementResponse

 A PartiQL batch statement response.. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**error** | [**BatchStatementError**](BatchStatementError.md) |  The error associated with a failed PartiQL batch statement.  |  [optional] |
|**tableName** | **String** |  The table name associated with a failed PartiQL batch statement.  |  [optional] |
|**item** | [**Map&lt;String, AttributeValue&gt;**](AttributeValue.md) |  A DynamoDB item associated with a BatchStatementResponse  |  [optional] |



