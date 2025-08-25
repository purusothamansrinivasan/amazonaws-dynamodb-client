

# TableCreationParameters

 The parameters for the table created as part of the import operation. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableName** | **String** |  The name of the table created as part of the import operation.  |  |
|**attributeDefinitions** | [**List&lt;AttributeDefinition&gt;**](AttributeDefinition.md) |  The attributes of the table created as part of the import operation.  |  |
|**keySchema** | [**List&lt;KeySchemaElement&gt;**](KeySchemaElement.md) |  The primary key and option sort key of the table created as part of the import operation.  |  |
|**billingMode** | **BillingMode** |  The billing mode for provisioning the table created as part of the import operation.  |  [optional] |
|**provisionedThroughput** | [**ProvisionedThroughput**](ProvisionedThroughput.md) |  |  [optional] |
|**ssESpecification** | [**SSESpecification**](SSESpecification.md) |  |  [optional] |
|**globalSecondaryIndexes** | [**List&lt;GlobalSecondaryIndex&gt;**](GlobalSecondaryIndex.md) |  The Global Secondary Indexes (GSI) of the table to be created as part of the import operation.  |  [optional] |



