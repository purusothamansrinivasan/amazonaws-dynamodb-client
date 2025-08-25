package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// UpdateItemOutput represents the UpdateItemOutput schema from the OpenAPI specification
type UpdateItemOutput struct {
	Itemcollectionmetrics interface{} `json:"ItemCollectionMetrics,omitempty"` // <p>Information about item collections, if any, that were affected by the <code>UpdateItem</code> operation. <code>ItemCollectionMetrics</code> is only returned if the <code>ReturnItemCollectionMetrics</code> parameter was specified. If the table does not have any local secondary indexes, this information is not returned in the response.</p> <p>Each <code>ItemCollectionMetrics</code> element consists of:</p> <ul> <li> <p> <code>ItemCollectionKey</code> - The partition key value of the item collection. This is the same as the partition key value of the item itself.</p> </li> <li> <p> <code>SizeEstimateRangeGB</code> - An estimate of item collection size, in gigabytes. This value is a two-element array containing a lower bound and an upper bound for the estimate. The estimate includes the size of all the items in the table, plus the size of all attributes projected into all of the local secondary indexes on that table. Use this estimate to measure whether a local secondary index is approaching its size limit.</p> <p>The estimate is subject to change over time; therefore, do not rely on the precision or accuracy of the estimate.</p> </li> </ul>
	Attributes interface{} `json:"Attributes,omitempty"` // <p>A map of attribute values as they appear before or after the <code>UpdateItem</code> operation, as determined by the <code>ReturnValues</code> parameter.</p> <p>The <code>Attributes</code> map is only present if the update was successful and <code>ReturnValues</code> was specified as something other than <code>NONE</code> in the request. Each element represents one attribute.</p>
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the <code>UpdateItem</code> operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the <code>ReturnConsumedCapacity</code> parameter was specified. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html#ItemSizeCalculations.Reads">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.
}

// QueryInput represents the QueryInput schema from the OpenAPI specification
type QueryInput struct {
	Scanindexforward interface{} `json:"ScanIndexForward,omitempty"` // <p>Specifies the order for index traversal: If <code>true</code> (default), the traversal is performed in ascending order; if <code>false</code>, the traversal is performed in descending order. </p> <p>Items with the same partition key value are stored in sorted order by sort key. If the sort key data type is Number, the results are stored in numeric order. For type String, the results are stored in order of UTF-8 bytes. For type Binary, DynamoDB treats each byte of the binary data as unsigned.</p> <p>If <code>ScanIndexForward</code> is <code>true</code>, DynamoDB returns the results in the order in which they are stored (by sort key value). This is the default behavior. If <code>ScanIndexForward</code> is <code>false</code>, DynamoDB reads the results in reverse order by sort key value, and then returns the results to the client.</p>
	Attributestoget interface{} `json:"AttributesToGet,omitempty"` // This is a legacy parameter. Use <code>ProjectionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html">AttributesToGet</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Conditionaloperator interface{} `json:"ConditionalOperator,omitempty"` // This is a legacy parameter. Use <code>FilterExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html">ConditionalOperator</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Indexname interface{} `json:"IndexName,omitempty"` // The name of an index to query. This index can be any local secondary index or global secondary index on the table. Note that if you use the <code>IndexName</code> parameter, you must also provide <code>TableName.</code>
	Limit interface{} `json:"Limit,omitempty"` // The maximum number of items to evaluate (not necessarily the number of matching items). If DynamoDB processes the number of items up to the limit while processing the results, it stops the operation and returns the matching values up to that point, and a key in <code>LastEvaluatedKey</code> to apply in a subsequent operation, so that you can pick up where you left off. Also, if the processed dataset size exceeds 1 MB before DynamoDB reaches this limit, it stops the operation and returns the matching values up to the limit, and a key in <code>LastEvaluatedKey</code> to apply in a subsequent operation to continue the operation. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html">Query and Scan</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Filterexpression interface{} `json:"FilterExpression,omitempty"` // <p>A string that contains conditions that DynamoDB applies after the <code>Query</code> operation, but before the data is returned to you. Items that do not satisfy the <code>FilterExpression</code> criteria are not returned.</p> <p>A <code>FilterExpression</code> does not allow key attributes. You cannot define a filter expression based on a partition key or a sort key.</p> <note> <p>A <code>FilterExpression</code> is applied after the items have already been read; the process of filtering does not consume any additional read capacity units.</p> </note> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#Query.FilterExpression">Filter Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Exclusivestartkey interface{} `json:"ExclusiveStartKey,omitempty"` // <p>The primary key of the first item that this operation will evaluate. Use the value that was returned for <code>LastEvaluatedKey</code> in the previous operation.</p> <p>The data type for <code>ExclusiveStartKey</code> must be String, Number, or Binary. No set data types are allowed.</p>
	SelectField interface{} `json:"Select,omitempty"` // <p>The attributes to be returned in the result. You can retrieve all item attributes, specific item attributes, the count of matching items, or in the case of an index, some or all of the attributes projected into the index.</p> <ul> <li> <p> <code>ALL_ATTRIBUTES</code> - Returns all of the item attributes from the specified table or index. If you query a local secondary index, then for each matching item in the index, DynamoDB fetches the entire item from the parent table. If the index is configured to project all item attributes, then all of the data can be obtained from the local secondary index, and no fetching is required.</p> </li> <li> <p> <code>ALL_PROJECTED_ATTRIBUTES</code> - Allowed only when querying an index. Retrieves all attributes that have been projected into the index. If the index is configured to project all attributes, this return value is equivalent to specifying <code>ALL_ATTRIBUTES</code>.</p> </li> <li> <p> <code>COUNT</code> - Returns the number of matching items, rather than the matching items themselves. Note that this uses the same quantity of read capacity units as getting the items, and is subject to the same item size calculations.</p> </li> <li> <p> <code>SPECIFIC_ATTRIBUTES</code> - Returns only the attributes listed in <code>ProjectionExpression</code>. This return value is equivalent to specifying <code>ProjectionExpression</code> without specifying any value for <code>Select</code>.</p> <p>If you query or scan a local secondary index and request only attributes that are projected into that index, the operation will read only the index and not the table. If any of the requested attributes are not projected into the local secondary index, DynamoDB fetches each of these attributes from the parent table. This extra fetching incurs additional throughput cost and latency.</p> <p>If you query or scan a global secondary index, you can only request attributes that are projected into the index. Global secondary index queries cannot fetch attributes from the parent table.</p> </li> </ul> <p>If neither <code>Select</code> nor <code>ProjectionExpression</code> are specified, DynamoDB defaults to <code>ALL_ATTRIBUTES</code> when accessing a table, and <code>ALL_PROJECTED_ATTRIBUTES</code> when accessing an index. You cannot use both <code>Select</code> and <code>ProjectionExpression</code> together in a single request, unless the value for <code>Select</code> is <code>SPECIFIC_ATTRIBUTES</code>. (This usage is equivalent to specifying <code>ProjectionExpression</code> without any value for <code>Select</code>.)</p> <note> <p>If you use the <code>ProjectionExpression</code> parameter, then the value for <code>Select</code> can only be <code>SPECIFIC_ATTRIBUTES</code>. Any other value for <code>Select</code> will return an error.</p> </note>
	Keyconditionexpression interface{} `json:"KeyConditionExpression,omitempty"` // <p>The condition that specifies the key values for items to be retrieved by the <code>Query</code> action.</p> <p>The condition must perform an equality test on a single partition key value.</p> <p>The condition can optionally perform one of several comparison tests on a single sort key value. This allows <code>Query</code> to retrieve one item with a given partition key value and sort key value, or several items that have the same partition key value but different sort key values.</p> <p>The partition key equality test is required, and must be specified in the following format:</p> <p> <code>partitionKeyName</code> <i>=</i> <code>:partitionkeyval</code> </p> <p>If you also want to provide a condition for the sort key, it must be combined using <code>AND</code> with the condition for the sort key. Following is an example, using the <b>=</b> comparison operator for the sort key:</p> <p> <code>partitionKeyName</code> <code>=</code> <code>:partitionkeyval</code> <code>AND</code> <code>sortKeyName</code> <code>=</code> <code>:sortkeyval</code> </p> <p>Valid comparisons for the sort key condition are as follows:</p> <ul> <li> <p> <code>sortKeyName</code> <code>=</code> <code>:sortkeyval</code> - true if the sort key value is equal to <code>:sortkeyval</code>.</p> </li> <li> <p> <code>sortKeyName</code> <code>&lt;</code> <code>:sortkeyval</code> - true if the sort key value is less than <code>:sortkeyval</code>.</p> </li> <li> <p> <code>sortKeyName</code> <code>&lt;=</code> <code>:sortkeyval</code> - true if the sort key value is less than or equal to <code>:sortkeyval</code>.</p> </li> <li> <p> <code>sortKeyName</code> <code>&gt;</code> <code>:sortkeyval</code> - true if the sort key value is greater than <code>:sortkeyval</code>.</p> </li> <li> <p> <code>sortKeyName</code> <code>&gt;= </code> <code>:sortkeyval</code> - true if the sort key value is greater than or equal to <code>:sortkeyval</code>.</p> </li> <li> <p> <code>sortKeyName</code> <code>BETWEEN</code> <code>:sortkeyval1</code> <code>AND</code> <code>:sortkeyval2</code> - true if the sort key value is greater than or equal to <code>:sortkeyval1</code>, and less than or equal to <code>:sortkeyval2</code>.</p> </li> <li> <p> <code>begins_with (</code> <code>sortKeyName</code>, <code>:sortkeyval</code> <code>)</code> - true if the sort key value begins with a particular operand. (You cannot use this function with a sort key that is of type Number.) Note that the function name <code>begins_with</code> is case-sensitive.</p> </li> </ul> <p>Use the <code>ExpressionAttributeValues</code> parameter to replace tokens such as <code>:partitionval</code> and <code>:sortval</code> with actual values at runtime.</p> <p>You can optionally use the <code>ExpressionAttributeNames</code> parameter to replace the names of the partition key and sort key with placeholder tokens. This option might be necessary if an attribute name conflicts with a DynamoDB reserved word. For example, the following <code>KeyConditionExpression</code> parameter causes an error because <i>Size</i> is a reserved word:</p> <ul> <li> <p> <code>Size = :myval</code> </p> </li> </ul> <p>To work around this, define a placeholder (such a <code>#S</code>) to represent the attribute name <i>Size</i>. <code>KeyConditionExpression</code> then is as follows:</p> <ul> <li> <p> <code>#S = :myval</code> </p> </li> </ul> <p>For a list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>For more information on <code>ExpressionAttributeNames</code> and <code>ExpressionAttributeValues</code>, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ExpressionPlaceholders.html">Using Placeholders for Attribute Names and Values</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Keyconditions interface{} `json:"KeyConditions,omitempty"` // This is a legacy parameter. Use <code>KeyConditionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.KeyConditions.html">KeyConditions</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // <p>One or more values that can be substituted in an expression.</p> <p>Use the <b>:</b> (colon) character in an expression to dereference an attribute value. For example, suppose that you wanted to check whether the value of the <i>ProductStatus</i> attribute was one of the following: </p> <p> <code>Available | Backordered | Discontinued</code> </p> <p>You would first need to specify <code>ExpressionAttributeValues</code> as follows:</p> <p> <code>{ ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"} }</code> </p> <p>You could then use these values in an expression, such as this:</p> <p> <code>ProductStatus IN (:avail, :back, :disc)</code> </p> <p>For more information on expression attribute values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Specifying Conditions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Queryfilter interface{} `json:"QueryFilter,omitempty"` // This is a legacy parameter. Use <code>FilterExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.QueryFilter.html">QueryFilter</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Tablename interface{} `json:"TableName"` // The name of the table containing the requested items.
	Projectionexpression interface{} `json:"ProjectionExpression,omitempty"` // <p>A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.</p> <p>If no attribute names are specified, then all attributes will be returned. If any of the requested attributes are not found, they will not appear in the result.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Accessing Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Consistentread interface{} `json:"ConsistentRead,omitempty"` // <p>Determines the read consistency model: If set to <code>true</code>, then the operation uses strongly consistent reads; otherwise, the operation uses eventually consistent reads.</p> <p>Strongly consistent reads are not supported on global secondary indexes. If you query a global secondary index with <code>ConsistentRead</code> set to <code>true</code>, you will receive a <code>ValidationException</code>.</p>
}

// ListExportsInput represents the ListExportsInput schema from the OpenAPI specification
type ListExportsInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"` // Maximum number of results to return per page.
	Nexttoken interface{} `json:"NextToken,omitempty"` // An optional string that, if supplied, must be copied from the output of a previous call to <code>ListExports</code>. When provided in this manner, the API fetches the next page of results.
	Tablearn interface{} `json:"TableArn,omitempty"` // The Amazon Resource Name (ARN) associated with the exported table.
}

// DescribeGlobalTableInput represents the DescribeGlobalTableInput schema from the OpenAPI specification
type DescribeGlobalTableInput struct {
	Globaltablename interface{} `json:"GlobalTableName"` // The name of the global table.
}

// ListBackupsOutput represents the ListBackupsOutput schema from the OpenAPI specification
type ListBackupsOutput struct {
	Lastevaluatedbackuparn interface{} `json:"LastEvaluatedBackupArn,omitempty"` // <p> The ARN of the backup last evaluated when the current page of results was returned, inclusive of the current page of results. This value may be specified as the <code>ExclusiveStartBackupArn</code> of a new <code>ListBackups</code> operation in order to fetch the next page of results. </p> <p> If <code>LastEvaluatedBackupArn</code> is empty, then the last page of results has been processed and there are no more results to be retrieved. </p> <p> If <code>LastEvaluatedBackupArn</code> is not empty, this may or may not indicate that there is more data to be returned. All results are guaranteed to have been returned if and only if no value for <code>LastEvaluatedBackupArn</code> is returned. </p>
	Backupsummaries interface{} `json:"BackupSummaries,omitempty"` // List of <code>BackupSummary</code> objects.
}

// ListTagsOfResourceOutput represents the ListTagsOfResourceOutput schema from the OpenAPI specification
type ListTagsOfResourceOutput struct {
	Nexttoken interface{} `json:"NextToken,omitempty"` // If this value is returned, there are additional results to be displayed. To retrieve them, call ListTagsOfResource again, with NextToken set to this value.
	Tags interface{} `json:"Tags,omitempty"` // The tags currently associated with the Amazon DynamoDB resource.
}

// UpdateGlobalTableInput represents the UpdateGlobalTableInput schema from the OpenAPI specification
type UpdateGlobalTableInput struct {
	Globaltablename interface{} `json:"GlobalTableName"` // The global table name.
	Replicaupdates interface{} `json:"ReplicaUpdates"` // A list of Regions that should be added or removed from the global table.
}

// DescribeTimeToLiveInput represents the DescribeTimeToLiveInput schema from the OpenAPI specification
type DescribeTimeToLiveInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table to be described.
}

// ArchivalSummary represents the ArchivalSummary schema from the OpenAPI specification
type ArchivalSummary struct {
	Archivalbackuparn interface{} `json:"ArchivalBackupArn,omitempty"` // The Amazon Resource Name (ARN) of the backup the table was archived to, when applicable in the archival reason. If you wish to restore this backup to the same table name, you will need to delete the original table.
	Archivaldatetime interface{} `json:"ArchivalDateTime,omitempty"` // The date and time when table archival was initiated by DynamoDB, in UNIX epoch time format.
	Archivalreason interface{} `json:"ArchivalReason,omitempty"` // <p>The reason DynamoDB archived the table. Currently, the only possible value is:</p> <ul> <li> <p> <code>INACCESSIBLE_ENCRYPTION_CREDENTIALS</code> - The table was archived due to the table's KMS key being inaccessible for more than seven days. An On-Demand backup was created at the archival time.</p> </li> </ul>
}

// ProvisionedThroughputOverride represents the ProvisionedThroughputOverride schema from the OpenAPI specification
type ProvisionedThroughputOverride struct {
	Readcapacityunits interface{} `json:"ReadCapacityUnits,omitempty"` // Replica-specific read capacity units. If not specified, uses the source table's read capacity settings.
}

// Put represents the Put schema from the OpenAPI specification
type Put struct {
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // One or more values that can be substituted in an expression.
	Item interface{} `json:"Item"` // A map of attribute name to attribute values, representing the primary key of the item to be written by <code>PutItem</code>. All of the table's primary key attributes must be specified, and their data types must match those of the table's key schema. If any attributes are present in the item that are part of an index key schema for the table, their types must match the index key schema.
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // Use <code>ReturnValuesOnConditionCheckFailure</code> to get the item attributes if the <code>Put</code> condition fails. For <code>ReturnValuesOnConditionCheckFailure</code>, the valid values are: NONE and ALL_OLD.
	Tablename interface{} `json:"TableName"` // Name of the table in which to write the item.
	Conditionexpression interface{} `json:"ConditionExpression,omitempty"` // A condition that must be satisfied in order for a conditional update to succeed.
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // One or more substitution tokens for attribute names in an expression.
}

// ReplicaAutoScalingDescription represents the ReplicaAutoScalingDescription schema from the OpenAPI specification
type ReplicaAutoScalingDescription struct {
	Regionname interface{} `json:"RegionName,omitempty"` // The Region where the replica exists.
	Replicaprovisionedreadcapacityautoscalingsettings AutoScalingSettingsDescription `json:"ReplicaProvisionedReadCapacityAutoScalingSettings,omitempty"` // Represents the auto scaling settings for a global table or global secondary index.
	Replicaprovisionedwritecapacityautoscalingsettings AutoScalingSettingsDescription `json:"ReplicaProvisionedWriteCapacityAutoScalingSettings,omitempty"` // Represents the auto scaling settings for a global table or global secondary index.
	Replicastatus interface{} `json:"ReplicaStatus,omitempty"` // <p>The current state of the replica:</p> <ul> <li> <p> <code>CREATING</code> - The replica is being created.</p> </li> <li> <p> <code>UPDATING</code> - The replica is being updated.</p> </li> <li> <p> <code>DELETING</code> - The replica is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The replica is ready for use.</p> </li> </ul>
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // Replica-specific global secondary index auto scaling settings.
}

// TransactGetItemsInput represents the TransactGetItemsInput schema from the OpenAPI specification
type TransactGetItemsInput struct {
	Returnconsumedcapacity interface{} `json:"ReturnConsumedCapacity,omitempty"` // A value of <code>TOTAL</code> causes consumed capacity information to be returned, and a value of <code>NONE</code> prevents that information from being returned. No other value is valid.
	Transactitems interface{} `json:"TransactItems"` // An ordered array of up to 100 <code>TransactGetItem</code> objects, each of which contains a <code>Get</code> structure.
}

// BatchStatementResponse represents the BatchStatementResponse schema from the OpenAPI specification
type BatchStatementResponse struct {
	ErrorField interface{} `json:"Error,omitempty"` // The error associated with a failed PartiQL batch statement.
	Item interface{} `json:"Item,omitempty"` // A DynamoDB item associated with a BatchStatementResponse
	Tablename interface{} `json:"TableName,omitempty"` // The table name associated with a failed PartiQL batch statement.
}

// ParameterizedStatement represents the ParameterizedStatement schema from the OpenAPI specification
type ParameterizedStatement struct {
	Parameters interface{} `json:"Parameters,omitempty"` // The parameter values.
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // <p>An optional parameter that returns the item attributes for a PartiQL <code>ParameterizedStatement</code> operation that failed a condition check.</p> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p>
	Statement interface{} `json:"Statement"` // A PartiQL statment that uses parameters.
}

// RestoreTableToPointInTimeOutput represents the RestoreTableToPointInTimeOutput schema from the OpenAPI specification
type RestoreTableToPointInTimeOutput struct {
	Tabledescription interface{} `json:"TableDescription,omitempty"` // Represents the properties of a table.
}

// GetItemOutput represents the GetItemOutput schema from the OpenAPI specification
type GetItemOutput struct {
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the <code>GetItem</code> operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the <code>ReturnConsumedCapacity</code> parameter was specified. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html#ItemSizeCalculations.Reads">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Item interface{} `json:"Item,omitempty"` // A map of attribute names to <code>AttributeValue</code> objects, as specified by <code>ProjectionExpression</code>.
}

// CreateBackupOutput represents the CreateBackupOutput schema from the OpenAPI specification
type CreateBackupOutput struct {
	Backupdetails interface{} `json:"BackupDetails,omitempty"` // Contains the details of the backup created for the table.
}

// SSEDescription represents the SSEDescription schema from the OpenAPI specification
type SSEDescription struct {
	Inaccessibleencryptiondatetime interface{} `json:"InaccessibleEncryptionDateTime,omitempty"` // Indicates the time, in UNIX epoch date format, when DynamoDB detected that the table's KMS key was inaccessible. This attribute will automatically be cleared when DynamoDB detects that the table's KMS key is accessible again. DynamoDB will initiate the table archival process when table's KMS key remains inaccessible for more than seven days from this date.
	Kmsmasterkeyarn interface{} `json:"KMSMasterKeyArn,omitempty"` // The KMS key ARN used for the KMS encryption.
	Ssetype interface{} `json:"SSEType,omitempty"` // <p>Server-side encryption type. The only supported value is:</p> <ul> <li> <p> <code>KMS</code> - Server-side encryption that uses Key Management Service. The key is stored in your account and is managed by KMS (KMS charges apply).</p> </li> </ul>
	Status interface{} `json:"Status,omitempty"` // <p>Represents the current state of server-side encryption. The only supported values are:</p> <ul> <li> <p> <code>ENABLED</code> - Server-side encryption is enabled.</p> </li> <li> <p> <code>UPDATING</code> - Server-side encryption is being updated.</p> </li> </ul>
}

// GlobalTableDescription represents the GlobalTableDescription schema from the OpenAPI specification
type GlobalTableDescription struct {
	Globaltablestatus interface{} `json:"GlobalTableStatus,omitempty"` // <p>The current state of the global table:</p> <ul> <li> <p> <code>CREATING</code> - The global table is being created.</p> </li> <li> <p> <code>UPDATING</code> - The global table is being updated.</p> </li> <li> <p> <code>DELETING</code> - The global table is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The global table is ready for use.</p> </li> </ul>
	Replicationgroup interface{} `json:"ReplicationGroup,omitempty"` // The Regions where the global table has replicas.
	Creationdatetime interface{} `json:"CreationDateTime,omitempty"` // The creation time of the global table.
	Globaltablearn interface{} `json:"GlobalTableArn,omitempty"` // The unique identifier of the global table.
	Globaltablename interface{} `json:"GlobalTableName,omitempty"` // The global table name.
}

// KeyConditions represents the KeyConditions schema from the OpenAPI specification
type KeyConditions struct {
}

// DescribeImportOutput represents the DescribeImportOutput schema from the OpenAPI specification
type DescribeImportOutput struct {
	Importtabledescription interface{} `json:"ImportTableDescription"` // Represents the properties of the table created for the import, and parameters of the import. The import parameters include import status, how many items were processed, and how many errors were encountered.
}

// CreateTableOutput represents the CreateTableOutput schema from the OpenAPI specification
type CreateTableOutput struct {
	Tabledescription interface{} `json:"TableDescription,omitempty"` // Represents the properties of the table.
}

// BillingModeSummary represents the BillingModeSummary schema from the OpenAPI specification
type BillingModeSummary struct {
	Billingmode interface{} `json:"BillingMode,omitempty"` // <p>Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.</p> <ul> <li> <p> <code>PROVISIONED</code> - Sets the read/write capacity mode to <code>PROVISIONED</code>. We recommend using <code>PROVISIONED</code> for predictable workloads.</p> </li> <li> <p> <code>PAY_PER_REQUEST</code> - Sets the read/write capacity mode to <code>PAY_PER_REQUEST</code>. We recommend using <code>PAY_PER_REQUEST</code> for unpredictable workloads. </p> </li> </ul>
	Lastupdatetopayperrequestdatetime interface{} `json:"LastUpdateToPayPerRequestDateTime,omitempty"` // Represents the time when <code>PAY_PER_REQUEST</code> was last set as the read/write capacity mode.
}

// ReplicaGlobalSecondaryIndexSettingsDescription represents the ReplicaGlobalSecondaryIndexSettingsDescription schema from the OpenAPI specification
type ReplicaGlobalSecondaryIndexSettingsDescription struct {
	Provisionedreadcapacityunits interface{} `json:"ProvisionedReadCapacityUnits,omitempty"` // The maximum number of strongly consistent reads consumed per second before DynamoDB returns a <code>ThrottlingException</code>.
	Provisionedwritecapacityautoscalingsettings interface{} `json:"ProvisionedWriteCapacityAutoScalingSettings,omitempty"` // Auto scaling settings for a global secondary index replica's write capacity units.
	Provisionedwritecapacityunits interface{} `json:"ProvisionedWriteCapacityUnits,omitempty"` // The maximum number of writes consumed per second before DynamoDB returns a <code>ThrottlingException</code>.
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index. The name must be unique among all other indexes on this table.
	Indexstatus interface{} `json:"IndexStatus,omitempty"` // <p> The current status of the global secondary index:</p> <ul> <li> <p> <code>CREATING</code> - The global secondary index is being created.</p> </li> <li> <p> <code>UPDATING</code> - The global secondary index is being updated.</p> </li> <li> <p> <code>DELETING</code> - The global secondary index is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The global secondary index is ready for use.</p> </li> </ul>
	Provisionedreadcapacityautoscalingsettings interface{} `json:"ProvisionedReadCapacityAutoScalingSettings,omitempty"` // Auto scaling settings for a global secondary index replica's read capacity units.
}

// DeleteBackupInput represents the DeleteBackupInput schema from the OpenAPI specification
type DeleteBackupInput struct {
	Backuparn interface{} `json:"BackupArn"` // The ARN associated with the backup.
}

// ItemCollectionMetricsPerTable represents the ItemCollectionMetricsPerTable schema from the OpenAPI specification
type ItemCollectionMetricsPerTable struct {
}

// GlobalTable represents the GlobalTable schema from the OpenAPI specification
type GlobalTable struct {
	Globaltablename interface{} `json:"GlobalTableName,omitempty"` // The global table name.
	Replicationgroup interface{} `json:"ReplicationGroup,omitempty"` // The Regions where the global table has replicas.
}

// CreateBackupInput represents the CreateBackupInput schema from the OpenAPI specification
type CreateBackupInput struct {
	Backupname interface{} `json:"BackupName"` // Specified name for the backup.
	Tablename interface{} `json:"TableName"` // The name of the table.
}

// TransactWriteItem represents the TransactWriteItem schema from the OpenAPI specification
type TransactWriteItem struct {
	Put interface{} `json:"Put,omitempty"` // A request to perform a <code>PutItem</code> operation.
	Update interface{} `json:"Update,omitempty"` // A request to perform an <code>UpdateItem</code> operation.
	Conditioncheck interface{} `json:"ConditionCheck,omitempty"` // A request to perform a check item operation.
	DeleteField interface{} `json:"Delete,omitempty"` // A request to perform a <code>DeleteItem</code> operation.
}

// DescribeKinesisStreamingDestinationOutput represents the DescribeKinesisStreamingDestinationOutput schema from the OpenAPI specification
type DescribeKinesisStreamingDestinationOutput struct {
	Kinesisdatastreamdestinations interface{} `json:"KinesisDataStreamDestinations,omitempty"` // The list of replica structures for the table being described.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table being described.
}

// UpdateGlobalSecondaryIndexAction represents the UpdateGlobalSecondaryIndexAction schema from the OpenAPI specification
type UpdateGlobalSecondaryIndexAction struct {
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index to be updated.
	Provisionedthroughput interface{} `json:"ProvisionedThroughput"` // <p>Represents the provisioned throughput settings for the specified global secondary index.</p> <p>For current minimum and maximum provisioned throughput values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
}

// Delete represents the Delete schema from the OpenAPI specification
type Delete struct {
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // One or more substitution tokens for attribute names in an expression.
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // One or more values that can be substituted in an expression.
	Key interface{} `json:"Key"` // The primary key of the item to be deleted. Each element consists of an attribute name and a value for that attribute.
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // Use <code>ReturnValuesOnConditionCheckFailure</code> to get the item attributes if the <code>Delete</code> condition fails. For <code>ReturnValuesOnConditionCheckFailure</code>, the valid values are: NONE and ALL_OLD.
	Tablename interface{} `json:"TableName"` // Name of the table in which the item to be deleted resides.
	Conditionexpression interface{} `json:"ConditionExpression,omitempty"` // A condition that must be satisfied in order for a conditional delete to succeed.
}

// ScanOutput represents the ScanOutput schema from the OpenAPI specification
type ScanOutput struct {
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the <code>Scan</code> operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the <code>ReturnConsumedCapacity</code> parameter was specified. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html#ItemSizeCalculations.Reads">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Count interface{} `json:"Count,omitempty"` // <p>The number of items in the response.</p> <p>If you set <code>ScanFilter</code> in the request, then <code>Count</code> is the number of items returned after the filter was applied, and <code>ScannedCount</code> is the number of matching items before the filter was applied.</p> <p>If you did not use a filter in the request, then <code>Count</code> is the same as <code>ScannedCount</code>.</p>
	Items interface{} `json:"Items,omitempty"` // An array of item attributes that match the scan criteria. Each element in this array consists of an attribute name and the value for that attribute.
	Lastevaluatedkey interface{} `json:"LastEvaluatedKey,omitempty"` // <p>The primary key of the item where the operation stopped, inclusive of the previous result set. Use this value to start a new operation, excluding this value in the new request.</p> <p>If <code>LastEvaluatedKey</code> is empty, then the "last page" of results has been processed and there is no more data to be retrieved.</p> <p>If <code>LastEvaluatedKey</code> is not empty, it does not necessarily mean that there is more data in the result set. The only way to know when you have reached the end of the result set is when <code>LastEvaluatedKey</code> is empty.</p>
	Scannedcount interface{} `json:"ScannedCount,omitempty"` // <p>The number of items evaluated, before any <code>ScanFilter</code> is applied. A high <code>ScannedCount</code> value with few, or no, <code>Count</code> results indicates an inefficient <code>Scan</code> operation. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#Count">Count and ScannedCount</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>If you did not use a filter in the request, then <code>ScannedCount</code> is the same as <code>Count</code>.</p>
}

// RestoreSummary represents the RestoreSummary schema from the OpenAPI specification
type RestoreSummary struct {
	Restoredatetime interface{} `json:"RestoreDateTime"` // Point in time or source backup time.
	Restoreinprogress interface{} `json:"RestoreInProgress"` // Indicates if a restore is in progress or not.
	Sourcebackuparn interface{} `json:"SourceBackupArn,omitempty"` // The Amazon Resource Name (ARN) of the backup from which the table was restored.
	Sourcetablearn interface{} `json:"SourceTableArn,omitempty"` // The ARN of the source table of the backup that is being restored.
}

// DescribeExportInput represents the DescribeExportInput schema from the OpenAPI specification
type DescribeExportInput struct {
	Exportarn interface{} `json:"ExportArn"` // The Amazon Resource Name (ARN) associated with the export.
}

// ProvisionedThroughputDescription represents the ProvisionedThroughputDescription schema from the OpenAPI specification
type ProvisionedThroughputDescription struct {
	Numberofdecreasestoday interface{} `json:"NumberOfDecreasesToday,omitempty"` // The number of provisioned throughput decreases for this table during this UTC calendar day. For current maximums on provisioned throughput decreases, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Readcapacityunits interface{} `json:"ReadCapacityUnits,omitempty"` // The maximum number of strongly consistent reads consumed per second before DynamoDB returns a <code>ThrottlingException</code>. Eventually consistent reads require less effort than strongly consistent reads, so a setting of 50 <code>ReadCapacityUnits</code> per second provides 100 eventually consistent <code>ReadCapacityUnits</code> per second.
	Writecapacityunits interface{} `json:"WriteCapacityUnits,omitempty"` // The maximum number of writes consumed per second before DynamoDB returns a <code>ThrottlingException</code>.
	Lastdecreasedatetime interface{} `json:"LastDecreaseDateTime,omitempty"` // The date and time of the last provisioned throughput decrease for this table.
	Lastincreasedatetime interface{} `json:"LastIncreaseDateTime,omitempty"` // The date and time of the last provisioned throughput increase for this table.
}

// UpdateReplicationGroupMemberAction represents the UpdateReplicationGroupMemberAction schema from the OpenAPI specification
type UpdateReplicationGroupMemberAction struct {
	Regionname interface{} `json:"RegionName"` // The Region where the replica exists.
	Tableclassoverride interface{} `json:"TableClassOverride,omitempty"` // Replica-specific table class. If not specified, uses the source table's table class.
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // Replica-specific global secondary index settings.
	Kmsmasterkeyid interface{} `json:"KMSMasterKeyId,omitempty"` // The KMS key of the replica that should be used for KMS encryption. To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB KMS key <code>alias/aws/dynamodb</code>.
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"` // Replica-specific provisioned throughput. If not specified, uses the source table's provisioned throughput settings.
}

// ListContributorInsightsInput represents the ListContributorInsightsInput schema from the OpenAPI specification
type ListContributorInsightsInput struct {
	Maxresults interface{} `json:"MaxResults,omitempty"` // Maximum number of results to return per page.
	Nexttoken interface{} `json:"NextToken,omitempty"` // A token to for the desired page, if there is one.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table.
}

// DeleteTableInput represents the DeleteTableInput schema from the OpenAPI specification
type DeleteTableInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table to delete.
}

// FailureException represents the FailureException schema from the OpenAPI specification
type FailureException struct {
	Exceptiondescription interface{} `json:"ExceptionDescription,omitempty"` // Description of the failure.
	Exceptionname interface{} `json:"ExceptionName,omitempty"` // Exception name.
}

// AttributeDefinition represents the AttributeDefinition schema from the OpenAPI specification
type AttributeDefinition struct {
	Attributename interface{} `json:"AttributeName"` // A name for the attribute.
	Attributetype interface{} `json:"AttributeType"` // <p>The data type for the attribute, where:</p> <ul> <li> <p> <code>S</code> - the attribute is of type String</p> </li> <li> <p> <code>N</code> - the attribute is of type Number</p> </li> <li> <p> <code>B</code> - the attribute is of type Binary</p> </li> </ul>
}

// PutItemOutput represents the PutItemOutput schema from the OpenAPI specification
type PutItemOutput struct {
	Attributes interface{} `json:"Attributes,omitempty"` // The attribute values as they appeared before the <code>PutItem</code> operation, but only if <code>ReturnValues</code> is specified as <code>ALL_OLD</code> in the request. Each element consists of an attribute name and an attribute value.
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the <code>PutItem</code> operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the <code>ReturnConsumedCapacity</code> parameter was specified. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Itemcollectionmetrics interface{} `json:"ItemCollectionMetrics,omitempty"` // <p>Information about item collections, if any, that were affected by the <code>PutItem</code> operation. <code>ItemCollectionMetrics</code> is only returned if the <code>ReturnItemCollectionMetrics</code> parameter was specified. If the table does not have any local secondary indexes, this information is not returned in the response.</p> <p>Each <code>ItemCollectionMetrics</code> element consists of:</p> <ul> <li> <p> <code>ItemCollectionKey</code> - The partition key value of the item collection. This is the same as the partition key value of the item itself.</p> </li> <li> <p> <code>SizeEstimateRangeGB</code> - An estimate of item collection size, in gigabytes. This value is a two-element array containing a lower bound and an upper bound for the estimate. The estimate includes the size of all the items in the table, plus the size of all attributes projected into all of the local secondary indexes on that table. Use this estimate to measure whether a local secondary index is approaching its size limit.</p> <p>The estimate is subject to change over time; therefore, do not rely on the precision or accuracy of the estimate.</p> </li> </ul>
}

// ReplicaDescription represents the ReplicaDescription schema from the OpenAPI specification
type ReplicaDescription struct {
	Regionname interface{} `json:"RegionName,omitempty"` // The name of the Region.
	Replicastatusdescription interface{} `json:"ReplicaStatusDescription,omitempty"` // Detailed information about the replica status.
	Replicatableclasssummary TableClassSummary `json:"ReplicaTableClassSummary,omitempty"` // Contains details of the table class.
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // Replica-specific global secondary index settings.
	Replicastatuspercentprogress interface{} `json:"ReplicaStatusPercentProgress,omitempty"` // Specifies the progress of a Create, Update, or Delete action on the replica as a percentage.
	Replicainaccessibledatetime interface{} `json:"ReplicaInaccessibleDateTime,omitempty"` // The time at which the replica was first detected as inaccessible. To determine cause of inaccessibility check the <code>ReplicaStatus</code> property.
	Replicastatus interface{} `json:"ReplicaStatus,omitempty"` // <p>The current state of the replica:</p> <ul> <li> <p> <code>CREATING</code> - The replica is being created.</p> </li> <li> <p> <code>UPDATING</code> - The replica is being updated.</p> </li> <li> <p> <code>DELETING</code> - The replica is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The replica is ready for use.</p> </li> <li> <p> <code>REGION_DISABLED</code> - The replica is inaccessible because the Amazon Web Services Region has been disabled.</p> <note> <p>If the Amazon Web Services Region remains inaccessible for more than 20 hours, DynamoDB will remove this replica from the replication group. The replica will not be deleted and replication will stop from and to this region.</p> </note> </li> <li> <p> <code>INACCESSIBLE_ENCRYPTION_CREDENTIALS </code> - The KMS key used to encrypt the table is inaccessible.</p> <note> <p>If the KMS key remains inaccessible for more than 20 hours, DynamoDB will remove this replica from the replication group. The replica will not be deleted and replication will stop from and to this region.</p> </note> </li> </ul>
	Kmsmasterkeyid interface{} `json:"KMSMasterKeyId,omitempty"` // The KMS key of the replica that will be used for KMS encryption.
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"` // Replica-specific provisioned throughput. If not described, uses the source table's provisioned throughput settings.
}

// DeleteItemInput represents the DeleteItemInput schema from the OpenAPI specification
type DeleteItemInput struct {
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Returnitemcollectionmetrics interface{} `json:"ReturnItemCollectionMetrics,omitempty"` // Determines whether item collection metrics are returned. If set to <code>SIZE</code>, the response includes statistics about item collections, if any, that were modified during the operation are returned in the response. If set to <code>NONE</code> (the default), no statistics are returned.
	Conditionexpression interface{} `json:"ConditionExpression,omitempty"` // <p>A condition that must be satisfied in order for a conditional <code>DeleteItem</code> to succeed.</p> <p>An expression can contain any of the following:</p> <ul> <li> <p>Functions: <code>attribute_exists | attribute_not_exists | attribute_type | contains | begins_with | size</code> </p> <p>These function names are case-sensitive.</p> </li> <li> <p>Comparison operators: <code>= | &lt;&gt; | &lt; | &gt; | &lt;= | &gt;= | BETWEEN | IN </code> </p> </li> <li> <p> Logical operators: <code>AND | OR | NOT</code> </p> </li> </ul> <p>For more information about condition expressions, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // <p>An optional parameter that returns the item attributes for a <code>DeleteItem</code> operation that failed a condition check.</p> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p>
	Tablename interface{} `json:"TableName"` // The name of the table from which to delete the item.
	Conditionaloperator interface{} `json:"ConditionalOperator,omitempty"` // This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html">ConditionalOperator</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Key interface{} `json:"Key"` // <p>A map of attribute names to <code>AttributeValue</code> objects, representing the primary key of the item to delete.</p> <p>For the primary key, you must provide all of the key attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for both the partition key and the sort key.</p>
	Returnvalues interface{} `json:"ReturnValues,omitempty"` // <p>Use <code>ReturnValues</code> if you want to get the item attributes as they appeared before they were deleted. For <code>DeleteItem</code>, the valid values are:</p> <ul> <li> <p> <code>NONE</code> - If <code>ReturnValues</code> is not specified, or if its value is <code>NONE</code>, then nothing is returned. (This setting is the default for <code>ReturnValues</code>.)</p> </li> <li> <p> <code>ALL_OLD</code> - The content of the old item is returned.</p> </li> </ul> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p> <note> <p>The <code>ReturnValues</code> parameter is used by several DynamoDB operations; however, <code>DeleteItem</code> does not recognize any values other than <code>NONE</code> or <code>ALL_OLD</code>.</p> </note>
	Expected interface{} `json:"Expected,omitempty"` // This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html">Expected</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // <p>One or more values that can be substituted in an expression.</p> <p>Use the <b>:</b> (colon) character in an expression to dereference an attribute value. For example, suppose that you wanted to check whether the value of the <i>ProductStatus</i> attribute was one of the following: </p> <p> <code>Available | Backordered | Discontinued</code> </p> <p>You would first need to specify <code>ExpressionAttributeValues</code> as follows:</p> <p> <code>{ ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"} }</code> </p> <p>You could then use these values in an expression, such as this:</p> <p> <code>ProductStatus IN (:avail, :back, :disc)</code> </p> <p>For more information on expression attribute values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
}

// AttributeValueUpdate represents the AttributeValueUpdate schema from the OpenAPI specification
type AttributeValueUpdate struct {
	Action interface{} `json:"Action,omitempty"` // <p>Specifies how to perform the update. Valid values are <code>PUT</code> (default), <code>DELETE</code>, and <code>ADD</code>. The behavior depends on whether the specified primary key already exists in the table.</p> <p> <b>If an item with the specified <i>Key</i> is found in the table:</b> </p> <ul> <li> <p> <code>PUT</code> - Adds the specified attribute to the item. If the attribute already exists, it is replaced by the new value. </p> </li> <li> <p> <code>DELETE</code> - If no value is specified, the attribute and its value are removed from the item. The data type of the specified value must match the existing value's data type.</p> <p>If a <i>set</i> of values is specified, then those values are subtracted from the old set. For example, if the attribute value was the set <code>[a,b,c]</code> and the <code>DELETE</code> action specified <code>[a,c]</code>, then the final attribute value would be <code>[b]</code>. Specifying an empty set is an error.</p> </li> <li> <p> <code>ADD</code> - If the attribute does not already exist, then the attribute and its values are added to the item. If the attribute does exist, then the behavior of <code>ADD</code> depends on the data type of the attribute:</p> <ul> <li> <p>If the existing attribute is a number, and if <code>Value</code> is also a number, then the <code>Value</code> is mathematically added to the existing attribute. If <code>Value</code> is a negative number, then it is subtracted from the existing attribute.</p> <note> <p> If you use <code>ADD</code> to increment or decrement a number value for an item that doesn't exist before the update, DynamoDB uses 0 as the initial value.</p> <p>In addition, if you use <code>ADD</code> to update an existing item, and intend to increment or decrement an attribute value which does not yet exist, DynamoDB uses <code>0</code> as the initial value. For example, suppose that the item you want to update does not yet have an attribute named <i>itemcount</i>, but you decide to <code>ADD</code> the number <code>3</code> to this attribute anyway, even though it currently does not exist. DynamoDB will create the <i>itemcount</i> attribute, set its initial value to <code>0</code>, and finally add <code>3</code> to it. The result will be a new <i>itemcount</i> attribute in the item, with a value of <code>3</code>.</p> </note> </li> <li> <p>If the existing data type is a set, and if the <code>Value</code> is also a set, then the <code>Value</code> is added to the existing set. (This is a <i>set</i> operation, not mathematical addition.) For example, if the attribute value was the set <code>[1,2]</code>, and the <code>ADD</code> action specified <code>[3]</code>, then the final attribute value would be <code>[1,2,3]</code>. An error occurs if an Add action is specified for a set attribute and the attribute type specified does not match the existing set type. </p> <p>Both sets must have the same primitive data type. For example, if the existing data type is a set of strings, the <code>Value</code> must also be a set of strings. The same holds true for number sets and binary sets.</p> </li> </ul> <p>This action is only valid for an existing attribute whose data type is number or is a set. Do not use <code>ADD</code> for any other data types.</p> </li> </ul> <p> <b>If no item with the specified <i>Key</i> is found:</b> </p> <ul> <li> <p> <code>PUT</code> - DynamoDB creates a new item with the specified primary key, and then adds the attribute. </p> </li> <li> <p> <code>DELETE</code> - Nothing happens; there is no attribute to delete.</p> </li> <li> <p> <code>ADD</code> - DynamoDB creates a new item with the supplied primary key and number (or set) for the attribute value. The only data types allowed are number, number set, string set or binary set.</p> </li> </ul>
	Value interface{} `json:"Value,omitempty"` // <p>Represents the data for an attribute.</p> <p>Each attribute value is described as a name-value pair. The name is the data type, and the value is the data itself.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes">Data Types</a> in the <i>Amazon DynamoDB Developer Guide</i>. </p>
}

// ListExportsOutput represents the ListExportsOutput schema from the OpenAPI specification
type ListExportsOutput struct {
	Exportsummaries interface{} `json:"ExportSummaries,omitempty"` // A list of <code>ExportSummary</code> objects.
	Nexttoken interface{} `json:"NextToken,omitempty"` // If this value is returned, there are additional results to be displayed. To retrieve them, call <code>ListExports</code> again, with <code>NextToken</code> set to this value.
}

// DescribeLimitsInput represents the DescribeLimitsInput schema from the OpenAPI specification
type DescribeLimitsInput struct {
}

// ImportTableInput represents the ImportTableInput schema from the OpenAPI specification
type ImportTableInput struct {
	Inputcompressiontype interface{} `json:"InputCompressionType,omitempty"` // Type of compression to be used on the input coming from the imported table.
	Inputformat interface{} `json:"InputFormat"` // The format of the source data. Valid values for <code>ImportFormat</code> are <code>CSV</code>, <code>DYNAMODB_JSON</code> or <code>ION</code>.
	Inputformatoptions interface{} `json:"InputFormatOptions,omitempty"` // Additional properties that specify how the input is formatted,
	S3bucketsource interface{} `json:"S3BucketSource"` // The S3 bucket that provides the source for the import.
	Tablecreationparameters interface{} `json:"TableCreationParameters"` // Parameters for the table to import the data into.
	Clienttoken interface{} `json:"ClientToken,omitempty"` // <p>Providing a <code>ClientToken</code> makes the call to <code>ImportTableInput</code> idempotent, meaning that multiple identical calls have the same effect as one single call.</p> <p>A client token is valid for 8 hours after the first request that uses it is completed. After 8 hours, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 8 hours, or the result might not be idempotent.</p> <p>If you submit a request with the same client token but a change in other parameters within the 8-hour idempotency window, DynamoDB returns an <code>IdempotentParameterMismatch</code> exception.</p>
}

// DeleteGlobalSecondaryIndexAction represents the DeleteGlobalSecondaryIndexAction schema from the OpenAPI specification
type DeleteGlobalSecondaryIndexAction struct {
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index to be deleted.
}

// Condition represents the Condition schema from the OpenAPI specification
type Condition struct {
	Comparisonoperator interface{} `json:"ComparisonOperator"` // <p>A comparator for evaluating attributes. For example, equals, greater than, less than, etc.</p> <p>The following comparison operators are available:</p> <p> <code>EQ | NE | LE | LT | GE | GT | NOT_NULL | NULL | CONTAINS | NOT_CONTAINS | BEGINS_WITH | IN | BETWEEN</code> </p> <p>The following are descriptions of each comparison operator.</p> <ul> <li> <p> <code>EQ</code> : Equal. <code>EQ</code> is supported for all data types, including lists and maps.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, Binary, String Set, Number Set, or Binary Set. If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not equal <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>NE</code> : Not equal. <code>NE</code> is supported for all data types, including lists and maps.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> of type String, Number, Binary, String Set, Number Set, or Binary Set. If an item contains an <code>AttributeValue</code> of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not equal <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>LE</code> : Less than or equal. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>LT</code> : Less than. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>GE</code> : Greater than or equal. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>GT</code> : Greater than. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>NOT_NULL</code> : The attribute exists. <code>NOT_NULL</code> is supported for all data types, including lists and maps.</p> <note> <p>This operator tests for the existence of an attribute, not its data type. If the data type of attribute "<code>a</code>" is null, and you evaluate it using <code>NOT_NULL</code>, the result is a Boolean <code>true</code>. This result is because the attribute "<code>a</code>" exists; its data type is not relevant to the <code>NOT_NULL</code> comparison operator.</p> </note> </li> <li> <p> <code>NULL</code> : The attribute does not exist. <code>NULL</code> is supported for all data types, including lists and maps.</p> <note> <p>This operator tests for the nonexistence of an attribute, not its data type. If the data type of attribute "<code>a</code>" is null, and you evaluate it using <code>NULL</code>, the result is a Boolean <code>false</code>. This is because the attribute "<code>a</code>" exists; its data type is not relevant to the <code>NULL</code> comparison operator.</p> </note> </li> <li> <p> <code>CONTAINS</code> : Checks for a subsequence, or value in a set.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If the target attribute of the comparison is of type String, then the operator checks for a substring match. If the target attribute of the comparison is of type Binary, then the operator looks for a subsequence of the target that matches the input. If the target attribute of the comparison is a set ("<code>SS</code>", "<code>NS</code>", or "<code>BS</code>"), then the operator evaluates to true if it finds an exact match with any member of the set.</p> <p>CONTAINS is supported for lists: When evaluating "<code>a CONTAINS b</code>", "<code>a</code>" can be a list; however, "<code>b</code>" cannot be a set, a map, or a list.</p> </li> <li> <p> <code>NOT_CONTAINS</code> : Checks for absence of a subsequence, or absence of a value in a set.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If the target attribute of the comparison is a String, then the operator checks for the absence of a substring match. If the target attribute of the comparison is Binary, then the operator checks for the absence of a subsequence of the target that matches the input. If the target attribute of the comparison is a set ("<code>SS</code>", "<code>NS</code>", or "<code>BS</code>"), then the operator evaluates to true if it <i>does not</i> find an exact match with any member of the set.</p> <p>NOT_CONTAINS is supported for lists: When evaluating "<code>a NOT CONTAINS b</code>", "<code>a</code>" can be a list; however, "<code>b</code>" cannot be a set, a map, or a list.</p> </li> <li> <p> <code>BEGINS_WITH</code> : Checks for a prefix. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> of type String or Binary (not a Number or a set type). The target attribute of the comparison must be of type String or Binary (not a Number or a set type).</p> <p/> </li> <li> <p> <code>IN</code> : Checks for matching elements in a list.</p> <p> <code>AttributeValueList</code> can contain one or more <code>AttributeValue</code> elements of type String, Number, or Binary. These attributes are compared against an existing attribute of an item. If any elements of the input are equal to the item attribute, the expression evaluates to true.</p> </li> <li> <p> <code>BETWEEN</code> : Greater than or equal to the first value, and less than or equal to the second value. </p> <p> <code>AttributeValueList</code> must contain two <code>AttributeValue</code> elements of the same type, either String, Number, or Binary (not a set type). A target attribute matches if the target value is greater than, or equal to, the first element and less than, or equal to, the second element. If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not compare to <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code> </p> </li> </ul> <p>For usage examples of <code>AttributeValueList</code> and <code>ComparisonOperator</code>, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.html">Legacy Conditional Parameters</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Attributevaluelist interface{} `json:"AttributeValueList,omitempty"` // <p>One or more values to evaluate against the supplied attribute. The number of values in the list depends on the <code>ComparisonOperator</code> being used.</p> <p>For type Number, value comparisons are numeric.</p> <p>String value comparisons for greater than, equals, or less than are based on ASCII character code values. For example, <code>a</code> is greater than <code>A</code>, and <code>a</code> is greater than <code>B</code>. For a list of code values, see <a href="http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters">http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters</a>.</p> <p>For Binary, DynamoDB treats each byte of the binary data as unsigned when it compares binary values.</p>
}

// CreateGlobalSecondaryIndexAction represents the CreateGlobalSecondaryIndexAction schema from the OpenAPI specification
type CreateGlobalSecondaryIndexAction struct {
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"` // <p>Represents the provisioned throughput settings for the specified global secondary index.</p> <p>For current minimum and maximum provisioned throughput values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index to be created.
	Keyschema interface{} `json:"KeySchema"` // The key schema for the global secondary index.
	Projection interface{} `json:"Projection"` // Represents attributes that are copied (projected) from the table into an index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
}

// DeleteBackupOutput represents the DeleteBackupOutput schema from the OpenAPI specification
type DeleteBackupOutput struct {
	Backupdescription interface{} `json:"BackupDescription,omitempty"` // Contains the description of the backup created for the table.
}

// BatchWriteItemOutput represents the BatchWriteItemOutput schema from the OpenAPI specification
type BatchWriteItemOutput struct {
	Unprocesseditems interface{} `json:"UnprocessedItems,omitempty"` // <p>A map of tables and requests against those tables that were not processed. The <code>UnprocessedItems</code> value is in the same form as <code>RequestItems</code>, so you can provide this value directly to a subsequent <code>BatchWriteItem</code> operation. For more information, see <code>RequestItems</code> in the Request Parameters section.</p> <p>Each <code>UnprocessedItems</code> entry consists of a table name and, for that table, a list of operations to perform (<code>DeleteRequest</code> or <code>PutRequest</code>).</p> <ul> <li> <p> <code>DeleteRequest</code> - Perform a <code>DeleteItem</code> operation on the specified item. The item to be deleted is identified by a <code>Key</code> subelement:</p> <ul> <li> <p> <code>Key</code> - A map of primary key attribute values that uniquely identify the item. Each entry in this map consists of an attribute name and an attribute value.</p> </li> </ul> </li> <li> <p> <code>PutRequest</code> - Perform a <code>PutItem</code> operation on the specified item. The item to be put is identified by an <code>Item</code> subelement:</p> <ul> <li> <p> <code>Item</code> - A map of attributes and their values. Each entry in this map consists of an attribute name and an attribute value. Attribute values must not be null; string and binary type attributes must have lengths greater than zero; and set type attributes must not be empty. Requests that contain empty values will be rejected with a <code>ValidationException</code> exception.</p> <p>If you specify any attributes that are part of an index key, then the data types for those attributes must match those of the schema in the table's attribute definition.</p> </li> </ul> </li> </ul> <p>If there are no unprocessed items remaining, the response contains an empty <code>UnprocessedItems</code> map.</p>
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // <p>The capacity units consumed by the entire <code>BatchWriteItem</code> operation.</p> <p>Each element consists of:</p> <ul> <li> <p> <code>TableName</code> - The table that consumed the provisioned throughput.</p> </li> <li> <p> <code>CapacityUnits</code> - The total number of capacity units consumed.</p> </li> </ul>
	Itemcollectionmetrics interface{} `json:"ItemCollectionMetrics,omitempty"` // <p>A list of tables that were processed by <code>BatchWriteItem</code> and, for each table, information about any item collections that were affected by individual <code>DeleteItem</code> or <code>PutItem</code> operations.</p> <p>Each entry consists of the following subelements:</p> <ul> <li> <p> <code>ItemCollectionKey</code> - The partition key value of the item collection. This is the same as the partition key value of the item.</p> </li> <li> <p> <code>SizeEstimateRangeGB</code> - An estimate of item collection size, expressed in GB. This is a two-element array containing a lower bound and an upper bound for the estimate. The estimate includes the size of all the items in the table, plus the size of all attributes projected into all of the local secondary indexes on the table. Use this estimate to measure whether a local secondary index is approaching its size limit.</p> <p>The estimate is subject to change over time; therefore, do not rely on the precision or accuracy of the estimate.</p> </li> </ul>
}

// ItemCollectionMetrics represents the ItemCollectionMetrics schema from the OpenAPI specification
type ItemCollectionMetrics struct {
	Itemcollectionkey interface{} `json:"ItemCollectionKey,omitempty"` // The partition key value of the item collection. This value is the same as the partition key value of the item.
	Sizeestimaterangegb interface{} `json:"SizeEstimateRangeGB,omitempty"` // <p>An estimate of item collection size, in gigabytes. This value is a two-element array containing a lower bound and an upper bound for the estimate. The estimate includes the size of all the items in the table, plus the size of all attributes projected into all of the local secondary indexes on that table. Use this estimate to measure whether a local secondary index is approaching its size limit.</p> <p>The estimate is subject to change over time; therefore, do not rely on the precision or accuracy of the estimate.</p>
}

// DescribeTimeToLiveOutput represents the DescribeTimeToLiveOutput schema from the OpenAPI specification
type DescribeTimeToLiveOutput struct {
	Timetolivedescription interface{} `json:"TimeToLiveDescription,omitempty"` // <p/>
}

// ExecuteStatementOutput represents the ExecuteStatementOutput schema from the OpenAPI specification
type ExecuteStatementOutput struct {
	Items interface{} `json:"Items,omitempty"` // If a read operation was used, this property will contain the result of the read operation; a map of attribute names and their values. For the write operations this value will be empty.
	Lastevaluatedkey interface{} `json:"LastEvaluatedKey,omitempty"` // The primary key of the item where the operation stopped, inclusive of the previous result set. Use this value to start a new operation, excluding this value in the new request. If <code>LastEvaluatedKey</code> is empty, then the "last page" of results has been processed and there is no more data to be retrieved. If <code>LastEvaluatedKey</code> is not empty, it does not necessarily mean that there is more data in the result set. The only way to know when you have reached the end of the result set is when <code>LastEvaluatedKey</code> is empty.
	Nexttoken interface{} `json:"NextToken,omitempty"` // If the response of a read request exceeds the response payload limit DynamoDB will set this value in the response. If set, you can use that this value in the subsequent request to get the remaining results.
	Consumedcapacity ConsumedCapacity `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by an operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the request asked for it. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.
}

// ListTablesOutput represents the ListTablesOutput schema from the OpenAPI specification
type ListTablesOutput struct {
	Lastevaluatedtablename interface{} `json:"LastEvaluatedTableName,omitempty"` // <p>The name of the last table in the current page of results. Use this value as the <code>ExclusiveStartTableName</code> in a new request to obtain the next page of results, until all the table names are returned.</p> <p>If you do not receive a <code>LastEvaluatedTableName</code> value in the response, this means that there are no more table names to be retrieved.</p>
	Tablenames interface{} `json:"TableNames,omitempty"` // <p>The names of the tables associated with the current account at the current endpoint. The maximum size of this array is 100.</p> <p>If <code>LastEvaluatedTableName</code> also appears in the output, you can use this value as the <code>ExclusiveStartTableName</code> parameter in a subsequent <code>ListTables</code> request and obtain the next page of results.</p>
}

// GetItemInput represents the GetItemInput schema from the OpenAPI specification
type GetItemInput struct {
	Attributestoget interface{} `json:"AttributesToGet,omitempty"` // This is a legacy parameter. Use <code>ProjectionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html">AttributesToGet</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Consistentread interface{} `json:"ConsistentRead,omitempty"` // Determines the read consistency model: If set to <code>true</code>, then the operation uses strongly consistent reads; otherwise, the operation uses eventually consistent reads.
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Key interface{} `json:"Key"` // <p>A map of attribute names to <code>AttributeValue</code> objects, representing the primary key of the item to retrieve.</p> <p>For the primary key, you must provide all of the attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for both the partition key and the sort key.</p>
	Projectionexpression interface{} `json:"ProjectionExpression,omitempty"` // <p>A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.</p> <p>If no attribute names are specified, then all attributes are returned. If any of the requested attributes are not found, they do not appear in the result.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Tablename interface{} `json:"TableName"` // The name of the table containing the requested item.
}

// SourceTableFeatureDetails represents the SourceTableFeatureDetails schema from the OpenAPI specification
type SourceTableFeatureDetails struct {
	Streamdescription interface{} `json:"StreamDescription,omitempty"` // Stream settings on the table when the backup was created.
	Timetolivedescription interface{} `json:"TimeToLiveDescription,omitempty"` // Time to Live settings on the table when the backup was created.
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // Represents the GSI properties for the table when the backup was created. It includes the IndexName, KeySchema, Projection, and ProvisionedThroughput for the GSIs on the table at the time of backup.
	Localsecondaryindexes interface{} `json:"LocalSecondaryIndexes,omitempty"` // Represents the LSI properties for the table when the backup was created. It includes the IndexName, KeySchema and Projection for the LSIs on the table at the time of backup.
	Ssedescription interface{} `json:"SSEDescription,omitempty"` // The description of the server-side encryption status on the table when the backup was created.
}

// DescribeGlobalTableOutput represents the DescribeGlobalTableOutput schema from the OpenAPI specification
type DescribeGlobalTableOutput struct {
	Globaltabledescription interface{} `json:"GlobalTableDescription,omitempty"` // Contains the details of the global table.
}

// DescribeContinuousBackupsInput represents the DescribeContinuousBackupsInput schema from the OpenAPI specification
type DescribeContinuousBackupsInput struct {
	Tablename interface{} `json:"TableName"` // Name of the table for which the customer wants to check the continuous backups and point in time recovery settings.
}

// ExecuteTransactionInput represents the ExecuteTransactionInput schema from the OpenAPI specification
type ExecuteTransactionInput struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"` // Set this value to get remaining results, if <code>NextToken</code> was returned in the statement response.
	Returnconsumedcapacity interface{} `json:"ReturnConsumedCapacity,omitempty"` // Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TransactGetItems.html">TransactGetItems</a> and <a href="https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_TransactWriteItems.html">TransactWriteItems</a>.
	Transactstatements interface{} `json:"TransactStatements"` // The list of PartiQL statements representing the transaction to run.
}

// DescribeKinesisStreamingDestinationInput represents the DescribeKinesisStreamingDestinationInput schema from the OpenAPI specification
type DescribeKinesisStreamingDestinationInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table being described.
}

// TransactWriteItemsInput represents the TransactWriteItemsInput schema from the OpenAPI specification
type TransactWriteItemsInput struct {
	Clientrequesttoken interface{} `json:"ClientRequestToken,omitempty"` // <p>Providing a <code>ClientRequestToken</code> makes the call to <code>TransactWriteItems</code> idempotent, meaning that multiple identical calls have the same effect as one single call.</p> <p>Although multiple identical calls using the same client request token produce the same result on the server (no side effects), the responses to the calls might not be the same. If the <code>ReturnConsumedCapacity</code> parameter is set, then the initial <code>TransactWriteItems</code> call returns the amount of write capacity units consumed in making the changes. Subsequent <code>TransactWriteItems</code> calls with the same client token return the number of read capacity units consumed in reading the item.</p> <p>A client request token is valid for 10 minutes after the first request that uses it is completed. After 10 minutes, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 10 minutes, or the result might not be idempotent.</p> <p>If you submit a request with the same client token but a change in other parameters within the 10-minute idempotency window, DynamoDB returns an <code>IdempotentParameterMismatch</code> exception.</p>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Returnitemcollectionmetrics interface{} `json:"ReturnItemCollectionMetrics,omitempty"` // Determines whether item collection metrics are returned. If set to <code>SIZE</code>, the response includes statistics about item collections (if any), that were modified during the operation and are returned in the response. If set to <code>NONE</code> (the default), no statistics are returned.
	Transactitems interface{} `json:"TransactItems"` // An ordered array of up to 100 <code>TransactWriteItem</code> objects, each of which contains a <code>ConditionCheck</code>, <code>Put</code>, <code>Update</code>, or <code>Delete</code> object. These can operate on items in different tables, but the tables must reside in the same Amazon Web Services account and Region, and no two of them can operate on the same item.
}

// ProvisionedThroughput represents the ProvisionedThroughput schema from the OpenAPI specification
type ProvisionedThroughput struct {
	Readcapacityunits interface{} `json:"ReadCapacityUnits"` // <p>The maximum number of strongly consistent reads consumed per second before DynamoDB returns a <code>ThrottlingException</code>. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html">Specifying Read and Write Requirements</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>If read/write capacity mode is <code>PAY_PER_REQUEST</code> the value is set to 0.</p>
	Writecapacityunits interface{} `json:"WriteCapacityUnits"` // <p>The maximum number of writes consumed per second before DynamoDB returns a <code>ThrottlingException</code>. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughput.html">Specifying Read and Write Requirements</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>If read/write capacity mode is <code>PAY_PER_REQUEST</code> the value is set to 0.</p>
}

// ImportSummary represents the ImportSummary schema from the OpenAPI specification
type ImportSummary struct {
	Importarn interface{} `json:"ImportArn,omitempty"` // The Amazon Resource Number (ARN) corresponding to the import request.
	Importstatus interface{} `json:"ImportStatus,omitempty"` // The status of the import operation.
	Inputformat interface{} `json:"InputFormat,omitempty"` // The format of the source data. Valid values are <code>CSV</code>, <code>DYNAMODB_JSON</code> or <code>ION</code>.
	S3bucketsource interface{} `json:"S3BucketSource,omitempty"` // The path and S3 bucket of the source file that is being imported. This includes the S3Bucket (required), S3KeyPrefix (optional) and S3BucketOwner (optional if the bucket is owned by the requester).
	Starttime interface{} `json:"StartTime,omitempty"` // The time at which this import task began.
	Tablearn interface{} `json:"TableArn,omitempty"` // The Amazon Resource Number (ARN) of the table being imported into.
	Cloudwatchloggrouparn interface{} `json:"CloudWatchLogGroupArn,omitempty"` // The Amazon Resource Number (ARN) of the Cloudwatch Log Group associated with this import task.
	Endtime interface{} `json:"EndTime,omitempty"` // The time at which this import task ended. (Does this include the successful complete creation of the table it was imported to?)
}

// StreamSpecification represents the StreamSpecification schema from the OpenAPI specification
type StreamSpecification struct {
	Streamviewtype interface{} `json:"StreamViewType,omitempty"` // <p> When an item in the table is modified, <code>StreamViewType</code> determines what information is written to the stream for this table. Valid values for <code>StreamViewType</code> are:</p> <ul> <li> <p> <code>KEYS_ONLY</code> - Only the key attributes of the modified item are written to the stream.</p> </li> <li> <p> <code>NEW_IMAGE</code> - The entire item, as it appears after it was modified, is written to the stream.</p> </li> <li> <p> <code>OLD_IMAGE</code> - The entire item, as it appeared before it was modified, is written to the stream.</p> </li> <li> <p> <code>NEW_AND_OLD_IMAGES</code> - Both the new and the old item images of the item are written to the stream.</p> </li> </ul>
	Streamenabled interface{} `json:"StreamEnabled"` // Indicates whether DynamoDB Streams is enabled (true) or disabled (false) on the table.
}

// BatchStatementError represents the BatchStatementError schema from the OpenAPI specification
type BatchStatementError struct {
	Code interface{} `json:"Code,omitempty"` // The error code associated with the failed PartiQL batch statement.
	Item interface{} `json:"Item,omitempty"` // The item which caused the condition check to fail. This will be set if ReturnValuesOnConditionCheckFailure is specified as <code>ALL_OLD</code>.
	Message interface{} `json:"Message,omitempty"` // The error message associated with the PartiQL batch response.
}

// DeleteReplicaAction represents the DeleteReplicaAction schema from the OpenAPI specification
type DeleteReplicaAction struct {
	Regionname interface{} `json:"RegionName"` // The Region of the replica to be removed.
}

// AutoScalingPolicyDescription represents the AutoScalingPolicyDescription schema from the OpenAPI specification
type AutoScalingPolicyDescription struct {
	Policyname interface{} `json:"PolicyName,omitempty"` // The name of the scaling policy.
	Targettrackingscalingpolicyconfiguration interface{} `json:"TargetTrackingScalingPolicyConfiguration,omitempty"` // Represents a target tracking scaling policy configuration.
}

// ExpressionAttributeValueMap represents the ExpressionAttributeValueMap schema from the OpenAPI specification
type ExpressionAttributeValueMap struct {
}

// UpdateContributorInsightsInput represents the UpdateContributorInsightsInput schema from the OpenAPI specification
type UpdateContributorInsightsInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table.
	Contributorinsightsaction interface{} `json:"ContributorInsightsAction"` // Represents the contributor insights action.
	Indexname interface{} `json:"IndexName,omitempty"` // The global secondary index name, if applicable.
}

// LocalSecondaryIndexDescription represents the LocalSecondaryIndexDescription schema from the OpenAPI specification
type LocalSecondaryIndexDescription struct {
	Indexsizebytes interface{} `json:"IndexSizeBytes,omitempty"` // The total size of the specified index, in bytes. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.
	Itemcount interface{} `json:"ItemCount,omitempty"` // The number of items in the specified index. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.
	Keyschema interface{} `json:"KeySchema,omitempty"` // <p>The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note>
	Projection interface{} `json:"Projection,omitempty"` // Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Indexarn interface{} `json:"IndexArn,omitempty"` // The Amazon Resource Name (ARN) that uniquely identifies the index.
	Indexname interface{} `json:"IndexName,omitempty"` // Represents the name of the local secondary index.
}

// AttributeValue represents the AttributeValue schema from the OpenAPI specification
type AttributeValue struct {
	S interface{} `json:"S,omitempty"` // <p>An attribute of type String. For example:</p> <p> <code>"S": "Hello"</code> </p>
	L interface{} `json:"L,omitempty"` // <p>An attribute of type List. For example:</p> <p> <code>"L": [ {"S": "Cookies"} , {"S": "Coffee"}, {"N": "3.14159"}]</code> </p>
	M interface{} `json:"M,omitempty"` // <p>An attribute of type Map. For example:</p> <p> <code>"M": {"Name": {"S": "Joe"}, "Age": {"N": "35"}}</code> </p>
	Ns interface{} `json:"NS,omitempty"` // <p>An attribute of type Number Set. For example:</p> <p> <code>"NS": ["42.2", "-19", "7.5", "3.14"]</code> </p> <p>Numbers are sent across the network to DynamoDB as strings, to maximize compatibility across languages and libraries. However, DynamoDB treats them as number type attributes for mathematical operations.</p>
	Ss interface{} `json:"SS,omitempty"` // <p>An attribute of type String Set. For example:</p> <p> <code>"SS": ["Giraffe", "Hippo" ,"Zebra"]</code> </p>
	Null interface{} `json:"NULL,omitempty"` // <p>An attribute of type Null. For example:</p> <p> <code>"NULL": true</code> </p>
	B interface{} `json:"B,omitempty"` // <p>An attribute of type Binary. For example:</p> <p> <code>"B": "dGhpcyB0ZXh0IGlzIGJhc2U2NC1lbmNvZGVk"</code> </p>
	BoolField interface{} `json:"BOOL,omitempty"` // <p>An attribute of type Boolean. For example:</p> <p> <code>"BOOL": true</code> </p>
	N interface{} `json:"N,omitempty"` // <p>An attribute of type Number. For example:</p> <p> <code>"N": "123.45"</code> </p> <p>Numbers are sent across the network to DynamoDB as strings, to maximize compatibility across languages and libraries. However, DynamoDB treats them as number type attributes for mathematical operations.</p>
	Bs interface{} `json:"BS,omitempty"` // <p>An attribute of type Binary Set. For example:</p> <p> <code>"BS": ["U3Vubnk=", "UmFpbnk=", "U25vd3k="]</code> </p>
}

// DescribeLimitsOutput represents the DescribeLimitsOutput schema from the OpenAPI specification
type DescribeLimitsOutput struct {
	Accountmaxreadcapacityunits interface{} `json:"AccountMaxReadCapacityUnits,omitempty"` // The maximum total read capacity units that your account allows you to provision across all of your tables in this Region.
	Accountmaxwritecapacityunits interface{} `json:"AccountMaxWriteCapacityUnits,omitempty"` // The maximum total write capacity units that your account allows you to provision across all of your tables in this Region.
	Tablemaxreadcapacityunits interface{} `json:"TableMaxReadCapacityUnits,omitempty"` // The maximum read capacity units that your account allows you to provision for a new table that you are creating in this Region, including the read capacity units provisioned for its global secondary indexes (GSIs).
	Tablemaxwritecapacityunits interface{} `json:"TableMaxWriteCapacityUnits,omitempty"` // The maximum write capacity units that your account allows you to provision for a new table that you are creating in this Region, including the write capacity units provisioned for its global secondary indexes (GSIs).
}

// TimeToLiveSpecification represents the TimeToLiveSpecification schema from the OpenAPI specification
type TimeToLiveSpecification struct {
	Attributename interface{} `json:"AttributeName"` // The name of the TTL attribute used to store the expiration time for items in the table.
	Enabled interface{} `json:"Enabled"` // Indicates whether TTL is to be enabled (true) or disabled (false) on the table.
}

// TransactGetItem represents the TransactGetItem schema from the OpenAPI specification
type TransactGetItem struct {
	Get interface{} `json:"Get"` // Contains the primary key that identifies the item to get, together with the name of the table that contains the item, and optionally the specific attributes of the item to retrieve.
}

// ReplicaUpdate represents the ReplicaUpdate schema from the OpenAPI specification
type ReplicaUpdate struct {
	DeleteField interface{} `json:"Delete,omitempty"` // The name of the existing replica to be removed.
	Create interface{} `json:"Create,omitempty"` // The parameters required for creating a replica on an existing global table.
}

// ExpressionAttributeNameMap represents the ExpressionAttributeNameMap schema from the OpenAPI specification
type ExpressionAttributeNameMap struct {
}

// ConditionCheck represents the ConditionCheck schema from the OpenAPI specification
type ConditionCheck struct {
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // One or more values that can be substituted in an expression. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html">Condition expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Key interface{} `json:"Key"` // The primary key of the item to be checked. Each element consists of an attribute name and a value for that attribute.
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // Use <code>ReturnValuesOnConditionCheckFailure</code> to get the item attributes if the <code>ConditionCheck</code> condition fails. For <code>ReturnValuesOnConditionCheckFailure</code>, the valid values are: NONE and ALL_OLD.
	Tablename interface{} `json:"TableName"` // Name of the table for the check item request.
	Conditionexpression interface{} `json:"ConditionExpression"` // A condition that must be satisfied in order for a conditional update to succeed. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ConditionExpressions.html">Condition expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // One or more substitution tokens for attribute names in an expression. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.ExpressionAttributeNames.html">Expression attribute names</a> in the <i>Amazon DynamoDB Developer Guide</i>.
}

// KeySchemaElement represents the KeySchemaElement schema from the OpenAPI specification
type KeySchemaElement struct {
	Attributename interface{} `json:"AttributeName"` // The name of a key attribute.
	Keytype interface{} `json:"KeyType"` // <p>The role that this key attribute will assume:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note>
}

// ReplicaGlobalSecondaryIndexSettingsUpdate represents the ReplicaGlobalSecondaryIndexSettingsUpdate schema from the OpenAPI specification
type ReplicaGlobalSecondaryIndexSettingsUpdate struct {
	Provisionedreadcapacityautoscalingsettingsupdate interface{} `json:"ProvisionedReadCapacityAutoScalingSettingsUpdate,omitempty"` // Auto scaling settings for managing a global secondary index replica's read capacity units.
	Provisionedreadcapacityunits interface{} `json:"ProvisionedReadCapacityUnits,omitempty"` // The maximum number of strongly consistent reads consumed per second before DynamoDB returns a <code>ThrottlingException</code>.
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index. The name must be unique among all other indexes on this table.
}

// TableAutoScalingDescription represents the TableAutoScalingDescription schema from the OpenAPI specification
type TableAutoScalingDescription struct {
	Replicas interface{} `json:"Replicas,omitempty"` // Represents replicas of the global table.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table.
	Tablestatus interface{} `json:"TableStatus,omitempty"` // <p>The current state of the table:</p> <ul> <li> <p> <code>CREATING</code> - The table is being created.</p> </li> <li> <p> <code>UPDATING</code> - The table is being updated.</p> </li> <li> <p> <code>DELETING</code> - The table is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The table is ready for use.</p> </li> </ul>
}

// DescribeContinuousBackupsOutput represents the DescribeContinuousBackupsOutput schema from the OpenAPI specification
type DescribeContinuousBackupsOutput struct {
	Continuousbackupsdescription interface{} `json:"ContinuousBackupsDescription,omitempty"` // Represents the continuous backups and point in time recovery settings on the table.
}

// DeleteReplicationGroupMemberAction represents the DeleteReplicationGroupMemberAction schema from the OpenAPI specification
type DeleteReplicationGroupMemberAction struct {
	Regionname interface{} `json:"RegionName"` // The Region where the replica exists.
}

// ExpectedAttributeMap represents the ExpectedAttributeMap schema from the OpenAPI specification
type ExpectedAttributeMap struct {
}

// DeleteRequest represents the DeleteRequest schema from the OpenAPI specification
type DeleteRequest struct {
	Key interface{} `json:"Key"` // A map of attribute name to attribute values, representing the primary key of the item to delete. All of the table's primary key attributes must be specified, and their data types must match those of the table's key schema.
}

// ItemCollectionKeyAttributeMap represents the ItemCollectionKeyAttributeMap schema from the OpenAPI specification
type ItemCollectionKeyAttributeMap struct {
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Value interface{} `json:"Value"` // The value of the tag. Tag values are case-sensitive and can be null.
	Key interface{} `json:"Key"` // The key of the tag. Tag keys are case sensitive. Each DynamoDB table can only have up to one tag with the same key. If you try to add an existing tag (same key), the existing tag value will be updated to the new value.
}

// DescribeGlobalTableSettingsInput represents the DescribeGlobalTableSettingsInput schema from the OpenAPI specification
type DescribeGlobalTableSettingsInput struct {
	Globaltablename interface{} `json:"GlobalTableName"` // The name of the global table to describe.
}

// UpdateGlobalTableOutput represents the UpdateGlobalTableOutput schema from the OpenAPI specification
type UpdateGlobalTableOutput struct {
	Globaltabledescription interface{} `json:"GlobalTableDescription,omitempty"` // Contains the details of the global table.
}

// PutItemInput represents the PutItemInput schema from the OpenAPI specification
type PutItemInput struct {
	Conditionexpression interface{} `json:"ConditionExpression,omitempty"` // <p>A condition that must be satisfied in order for a conditional <code>PutItem</code> operation to succeed.</p> <p>An expression can contain any of the following:</p> <ul> <li> <p>Functions: <code>attribute_exists | attribute_not_exists | attribute_type | contains | begins_with | size</code> </p> <p>These function names are case-sensitive.</p> </li> <li> <p>Comparison operators: <code>= | &lt;&gt; | &lt; | &gt; | &lt;= | &gt;= | BETWEEN | IN </code> </p> </li> <li> <p> Logical operators: <code>AND | OR | NOT</code> </p> </li> </ul> <p>For more information on condition expressions, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Returnitemcollectionmetrics interface{} `json:"ReturnItemCollectionMetrics,omitempty"` // Determines whether item collection metrics are returned. If set to <code>SIZE</code>, the response includes statistics about item collections, if any, that were modified during the operation are returned in the response. If set to <code>NONE</code> (the default), no statistics are returned.
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // <p>An optional parameter that returns the item attributes for a <code>PutItem</code> operation that failed a condition check.</p> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p>
	Expected interface{} `json:"Expected,omitempty"` // This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html">Expected</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // <p>One or more values that can be substituted in an expression.</p> <p>Use the <b>:</b> (colon) character in an expression to dereference an attribute value. For example, suppose that you wanted to check whether the value of the <i>ProductStatus</i> attribute was one of the following: </p> <p> <code>Available | Backordered | Discontinued</code> </p> <p>You would first need to specify <code>ExpressionAttributeValues</code> as follows:</p> <p> <code>{ ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"} }</code> </p> <p>You could then use these values in an expression, such as this:</p> <p> <code>ProductStatus IN (:avail, :back, :disc)</code> </p> <p>For more information on expression attribute values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Conditionaloperator interface{} `json:"ConditionalOperator,omitempty"` // This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html">ConditionalOperator</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Item interface{} `json:"Item"` // <p>A map of attribute name/value pairs, one for each attribute. Only the primary key attributes are required; you can optionally provide other attribute name-value pairs for the item.</p> <p>You must provide all of the attributes for the primary key. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide both values for both the partition key and the sort key.</p> <p>If you specify any attributes that are part of an index key, then the data types for those attributes must match those of the schema in the table's attribute definition.</p> <p>Empty String and Binary attribute values are allowed. Attribute values of type String and Binary must have a length greater than zero if the attribute is used as a key attribute for a table or index.</p> <p>For more information about primary keys, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.CoreComponents.html#HowItWorks.CoreComponents.PrimaryKey">Primary Key</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>Each element in the <code>Item</code> map is an <code>AttributeValue</code> object.</p>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Returnvalues interface{} `json:"ReturnValues,omitempty"` // <p>Use <code>ReturnValues</code> if you want to get the item attributes as they appeared before they were updated with the <code>PutItem</code> request. For <code>PutItem</code>, the valid values are:</p> <ul> <li> <p> <code>NONE</code> - If <code>ReturnValues</code> is not specified, or if its value is <code>NONE</code>, then nothing is returned. (This setting is the default for <code>ReturnValues</code>.)</p> </li> <li> <p> <code>ALL_OLD</code> - If <code>PutItem</code> overwrote an attribute name-value pair, then the content of the old item is returned.</p> </li> </ul> <p>The values returned are strongly consistent.</p> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p> <note> <p>The <code>ReturnValues</code> parameter is used by several DynamoDB operations; however, <code>PutItem</code> does not recognize any values other than <code>NONE</code> or <code>ALL_OLD</code>.</p> </note>
	Tablename interface{} `json:"TableName"` // The name of the table to contain the item.
}

// BatchWriteItemRequestMap represents the BatchWriteItemRequestMap schema from the OpenAPI specification
type BatchWriteItemRequestMap struct {
}

// AttributeMap represents the AttributeMap schema from the OpenAPI specification
type AttributeMap struct {
}

// UpdateTableReplicaAutoScalingOutput represents the UpdateTableReplicaAutoScalingOutput schema from the OpenAPI specification
type UpdateTableReplicaAutoScalingOutput struct {
	Tableautoscalingdescription interface{} `json:"TableAutoScalingDescription,omitempty"` // Returns information about the auto scaling settings of a table with replicas.
}

// BatchStatementRequest represents the BatchStatementRequest schema from the OpenAPI specification
type BatchStatementRequest struct {
	Statement interface{} `json:"Statement"` // A valid PartiQL statement.
	Consistentread interface{} `json:"ConsistentRead,omitempty"` // The read consistency of the PartiQL batch request.
	Parameters interface{} `json:"Parameters,omitempty"` // The parameters associated with a PartiQL statement in the batch request.
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // <p>An optional parameter that returns the item attributes for a PartiQL batch request operation that failed a condition check.</p> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p>
}

// LocalSecondaryIndex represents the LocalSecondaryIndex schema from the OpenAPI specification
type LocalSecondaryIndex struct {
	Keyschema interface{} `json:"KeySchema"` // <p>The complete key schema for the local secondary index, consisting of one or more pairs of attribute names and key types:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note>
	Projection interface{} `json:"Projection"` // Represents attributes that are copied (projected) from the table into the local secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Indexname interface{} `json:"IndexName"` // The name of the local secondary index. The name must be unique among all other indexes on this table.
}

// AttributeUpdates represents the AttributeUpdates schema from the OpenAPI specification
type AttributeUpdates struct {
}

// UpdateTableReplicaAutoScalingInput represents the UpdateTableReplicaAutoScalingInput schema from the OpenAPI specification
type UpdateTableReplicaAutoScalingInput struct {
	Replicaupdates interface{} `json:"ReplicaUpdates,omitempty"` // Represents the auto scaling settings of replicas of the table that will be modified.
	Tablename interface{} `json:"TableName"` // The name of the global table to be updated.
	Globalsecondaryindexupdates interface{} `json:"GlobalSecondaryIndexUpdates,omitempty"` // Represents the auto scaling settings of the global secondary indexes of the replica to be updated.
	Provisionedwritecapacityautoscalingupdate AutoScalingSettingsUpdate `json:"ProvisionedWriteCapacityAutoScalingUpdate,omitempty"` // Represents the auto scaling settings to be modified for a global table or global secondary index.
}

// InputFormatOptions represents the InputFormatOptions schema from the OpenAPI specification
type InputFormatOptions struct {
	Csv interface{} `json:"Csv,omitempty"` // The options for imported source files in CSV format. The values are Delimiter and HeaderList.
}

// UpdateTableOutput represents the UpdateTableOutput schema from the OpenAPI specification
type UpdateTableOutput struct {
	Tabledescription interface{} `json:"TableDescription,omitempty"` // Represents the properties of the table.
}

// CreateReplicationGroupMemberAction represents the CreateReplicationGroupMemberAction schema from the OpenAPI specification
type CreateReplicationGroupMemberAction struct {
	Regionname interface{} `json:"RegionName"` // The Region where the new replica will be created.
	Tableclassoverride interface{} `json:"TableClassOverride,omitempty"` // Replica-specific table class. If not specified, uses the source table's table class.
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // Replica-specific global secondary index settings.
	Kmsmasterkeyid interface{} `json:"KMSMasterKeyId,omitempty"` // The KMS key that should be used for KMS encryption in the new replica. To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB KMS key <code>alias/aws/dynamodb</code>.
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"` // Replica-specific provisioned throughput. If not specified, uses the source table's provisioned throughput settings.
}

// RestoreTableFromBackupOutput represents the RestoreTableFromBackupOutput schema from the OpenAPI specification
type RestoreTableFromBackupOutput struct {
	Tabledescription interface{} `json:"TableDescription,omitempty"` // The description of the table created from an existing backup.
}

// TableDescription represents the TableDescription schema from the OpenAPI specification
type TableDescription struct {
	Tableclasssummary interface{} `json:"TableClassSummary,omitempty"` // Contains details of the table class.
	Tableid interface{} `json:"TableId,omitempty"` // Unique identifier for the table for which the backup was created.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table.
	Lateststreamlabel interface{} `json:"LatestStreamLabel,omitempty"` // <p>A timestamp, in ISO 8601 format, for this stream.</p> <p>Note that <code>LatestStreamLabel</code> is not a unique identifier for the stream, because it is possible that a stream from another table might have the same timestamp. However, the combination of the following three elements is guaranteed to be unique:</p> <ul> <li> <p>Amazon Web Services customer ID</p> </li> <li> <p>Table name</p> </li> <li> <p> <code>StreamLabel</code> </p> </li> </ul>
	Deletionprotectionenabled interface{} `json:"DeletionProtectionEnabled,omitempty"` // Indicates whether deletion protection is enabled (true) or disabled (false) on the table.
	Tablesizebytes interface{} `json:"TableSizeBytes,omitempty"` // The total size of the specified table, in bytes. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.
	Archivalsummary interface{} `json:"ArchivalSummary,omitempty"` // Contains information about the table archive.
	Globaltableversion interface{} `json:"GlobalTableVersion,omitempty"` // Represents the version of <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GlobalTables.html">global tables</a> in use, if the table is replicated across Amazon Web Services Regions.
	Lateststreamarn interface{} `json:"LatestStreamArn,omitempty"` // The Amazon Resource Name (ARN) that uniquely identifies the latest stream for this table.
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"` // The provisioned throughput settings for the table, consisting of read and write capacity units, along with data about increases and decreases.
	Replicas interface{} `json:"Replicas,omitempty"` // Represents replicas of the table.
	Creationdatetime interface{} `json:"CreationDateTime,omitempty"` // The date and time when the table was created, in <a href="http://www.epochconverter.com/">UNIX epoch time</a> format.
	Localsecondaryindexes interface{} `json:"LocalSecondaryIndexes,omitempty"` // <p>Represents one or more local secondary indexes on the table. Each index is scoped to a given partition key value. Tables with one or more local secondary indexes are subject to an item collection size limit, where the amount of data within a given item collection cannot exceed 10 GB. Each element is composed of:</p> <ul> <li> <p> <code>IndexName</code> - The name of the local secondary index.</p> </li> <li> <p> <code>KeySchema</code> - Specifies the complete index key schema. The attribute names in the key schema must be between 1 and 255 characters (inclusive). The key schema must begin with the same partition key as the table.</p> </li> <li> <p> <code>Projection</code> - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of:</p> <ul> <li> <p> <code>ProjectionType</code> - One of the following:</p> <ul> <li> <p> <code>KEYS_ONLY</code> - Only the index and primary keys are projected into the index.</p> </li> <li> <p> <code>INCLUDE</code> - Only the specified table attributes are projected into the index. The list of projected attributes is in <code>NonKeyAttributes</code>.</p> </li> <li> <p> <code>ALL</code> - All of the table attributes are projected into the index.</p> </li> </ul> </li> <li> <p> <code>NonKeyAttributes</code> - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in <code>NonKeyAttributes</code>, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.</p> </li> </ul> </li> <li> <p> <code>IndexSizeBytes</code> - Represents the total size of the index, in bytes. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.</p> </li> <li> <p> <code>ItemCount</code> - Represents the number of items in the index. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.</p> </li> </ul> <p>If the table is in the <code>DELETING</code> state, no information about indexes will be returned.</p>
	Keyschema interface{} `json:"KeySchema,omitempty"` // <p>The primary key structure for the table. Each <code>KeySchemaElement</code> consists of:</p> <ul> <li> <p> <code>AttributeName</code> - The name of the attribute.</p> </li> <li> <p> <code>KeyType</code> - The role of the attribute:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note> </li> </ul> <p>For more information about primary keys, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html#DataModelPrimaryKey">Primary Key</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Tablestatus interface{} `json:"TableStatus,omitempty"` // <p>The current state of the table:</p> <ul> <li> <p> <code>CREATING</code> - The table is being created.</p> </li> <li> <p> <code>UPDATING</code> - The table/index configuration is being updated. The table/index remains available for data operations when <code>UPDATING</code>.</p> </li> <li> <p> <code>DELETING</code> - The table is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The table is ready for use.</p> </li> <li> <p> <code>INACCESSIBLE_ENCRYPTION_CREDENTIALS</code> - The KMS key used to encrypt the table in inaccessible. Table operations may fail due to failure to use the KMS key. DynamoDB will initiate the table archival process when a table's KMS key remains inaccessible for more than seven days. </p> </li> <li> <p> <code>ARCHIVING</code> - The table is being archived. Operations are not allowed until archival is complete. </p> </li> <li> <p> <code>ARCHIVED</code> - The table has been archived. See the ArchivalReason for more information. </p> </li> </ul>
	Ssedescription interface{} `json:"SSEDescription,omitempty"` // The description of the server-side encryption status on the specified table.
	Billingmodesummary interface{} `json:"BillingModeSummary,omitempty"` // Contains the details for the read/write capacity mode.
	Tablearn interface{} `json:"TableArn,omitempty"` // The Amazon Resource Name (ARN) that uniquely identifies the table.
	Attributedefinitions interface{} `json:"AttributeDefinitions,omitempty"` // <p>An array of <code>AttributeDefinition</code> objects. Each of these objects describes one attribute in the table and index key schema.</p> <p>Each <code>AttributeDefinition</code> object in this array is composed of:</p> <ul> <li> <p> <code>AttributeName</code> - The name of the attribute.</p> </li> <li> <p> <code>AttributeType</code> - The data type for the attribute.</p> </li> </ul>
	Restoresummary interface{} `json:"RestoreSummary,omitempty"` // Contains details for the restore.
	Streamspecification interface{} `json:"StreamSpecification,omitempty"` // The current DynamoDB Streams configuration for the table.
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // <p>The global secondary indexes, if any, on the table. Each index is scoped to a given partition key value. Each element is composed of:</p> <ul> <li> <p> <code>Backfilling</code> - If true, then the index is currently in the backfilling phase. Backfilling occurs only when a new global secondary index is added to the table. It is the process by which DynamoDB populates the new index with data from the table. (This attribute does not appear for indexes that were created during a <code>CreateTable</code> operation.) </p> <p> You can delete an index that is being created during the <code>Backfilling</code> phase when <code>IndexStatus</code> is set to CREATING and <code>Backfilling</code> is true. You can't delete the index that is being created when <code>IndexStatus</code> is set to CREATING and <code>Backfilling</code> is false. (This attribute does not appear for indexes that were created during a <code>CreateTable</code> operation.)</p> </li> <li> <p> <code>IndexName</code> - The name of the global secondary index.</p> </li> <li> <p> <code>IndexSizeBytes</code> - The total size of the global secondary index, in bytes. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value. </p> </li> <li> <p> <code>IndexStatus</code> - The current status of the global secondary index:</p> <ul> <li> <p> <code>CREATING</code> - The index is being created.</p> </li> <li> <p> <code>UPDATING</code> - The index is being updated.</p> </li> <li> <p> <code>DELETING</code> - The index is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The index is ready for use.</p> </li> </ul> </li> <li> <p> <code>ItemCount</code> - The number of items in the global secondary index. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value. </p> </li> <li> <p> <code>KeySchema</code> - Specifies the complete index key schema. The attribute names in the key schema must be between 1 and 255 characters (inclusive). The key schema must begin with the same partition key as the table.</p> </li> <li> <p> <code>Projection</code> - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of:</p> <ul> <li> <p> <code>ProjectionType</code> - One of the following:</p> <ul> <li> <p> <code>KEYS_ONLY</code> - Only the index and primary keys are projected into the index.</p> </li> <li> <p> <code>INCLUDE</code> - In addition to the attributes described in <code>KEYS_ONLY</code>, the secondary index will include other non-key attributes that you specify.</p> </li> <li> <p> <code>ALL</code> - All of the table attributes are projected into the index.</p> </li> </ul> </li> <li> <p> <code>NonKeyAttributes</code> - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in <code>NonKeyAttributes</code>, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.</p> </li> </ul> </li> <li> <p> <code>ProvisionedThroughput</code> - The provisioned throughput settings for the global secondary index, consisting of read and write capacity units, along with data about increases and decreases. </p> </li> </ul> <p>If the table is in the <code>DELETING</code> state, no information about indexes will be returned.</p>
	Itemcount interface{} `json:"ItemCount,omitempty"` // The number of items in the specified table. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.
}

// UpdateGlobalTableSettingsOutput represents the UpdateGlobalTableSettingsOutput schema from the OpenAPI specification
type UpdateGlobalTableSettingsOutput struct {
	Globaltablename interface{} `json:"GlobalTableName,omitempty"` // The name of the global table.
	Replicasettings interface{} `json:"ReplicaSettings,omitempty"` // The Region-specific settings for the global table.
}

// ListTablesInput represents the ListTablesInput schema from the OpenAPI specification
type ListTablesInput struct {
	Limit interface{} `json:"Limit,omitempty"` // A maximum number of table names to return. If this parameter is not specified, the limit is 100.
	Exclusivestarttablename interface{} `json:"ExclusiveStartTableName,omitempty"` // The first table name that this operation will evaluate. Use the value that was returned for <code>LastEvaluatedTableName</code> in a previous operation, so that you can obtain the next page of results.
}

// RestoreTableToPointInTimeInput represents the RestoreTableToPointInTimeInput schema from the OpenAPI specification
type RestoreTableToPointInTimeInput struct {
	Billingmodeoverride interface{} `json:"BillingModeOverride,omitempty"` // The billing mode of the restored table.
	Globalsecondaryindexoverride interface{} `json:"GlobalSecondaryIndexOverride,omitempty"` // List of global secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"` // Provisioned throughput settings for the restored table.
	Ssespecificationoverride interface{} `json:"SSESpecificationOverride,omitempty"` // The new server-side encryption settings for the restored table.
	Sourcetablename interface{} `json:"SourceTableName,omitempty"` // Name of the source table that is being restored.
	Targettablename interface{} `json:"TargetTableName"` // The name of the new table to which it must be restored to.
	Restoredatetime interface{} `json:"RestoreDateTime,omitempty"` // Time in the past to restore the table to.
	Sourcetablearn interface{} `json:"SourceTableArn,omitempty"` // The DynamoDB table that will be restored. This value is an Amazon Resource Name (ARN).
	Localsecondaryindexoverride interface{} `json:"LocalSecondaryIndexOverride,omitempty"` // List of local secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.
	Uselatestrestorabletime interface{} `json:"UseLatestRestorableTime,omitempty"` // Restore the table to the latest possible time. <code>LatestRestorableDateTime</code> is typically 5 minutes before the current time.
}

// AutoScalingSettingsUpdate represents the AutoScalingSettingsUpdate schema from the OpenAPI specification
type AutoScalingSettingsUpdate struct {
	Minimumunits interface{} `json:"MinimumUnits,omitempty"` // The minimum capacity units that a global table or global secondary index should be scaled down to.
	Scalingpolicyupdate interface{} `json:"ScalingPolicyUpdate,omitempty"` // The scaling policy to apply for scaling target global table or global secondary index capacity units.
	Autoscalingdisabled interface{} `json:"AutoScalingDisabled,omitempty"` // Disabled auto scaling for this global table or global secondary index.
	Autoscalingrolearn interface{} `json:"AutoScalingRoleArn,omitempty"` // Role ARN used for configuring auto scaling policy.
	Maximumunits interface{} `json:"MaximumUnits,omitempty"` // The maximum capacity units that a global table or global secondary index should be scaled up to.
}

// ExpectedAttributeValue represents the ExpectedAttributeValue schema from the OpenAPI specification
type ExpectedAttributeValue struct {
	Attributevaluelist interface{} `json:"AttributeValueList,omitempty"` // <p>One or more values to evaluate against the supplied attribute. The number of values in the list depends on the <code>ComparisonOperator</code> being used.</p> <p>For type Number, value comparisons are numeric.</p> <p>String value comparisons for greater than, equals, or less than are based on ASCII character code values. For example, <code>a</code> is greater than <code>A</code>, and <code>a</code> is greater than <code>B</code>. For a list of code values, see <a href="http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters">http://en.wikipedia.org/wiki/ASCII#ASCII_printable_characters</a>.</p> <p>For Binary, DynamoDB treats each byte of the binary data as unsigned when it compares binary values.</p> <p>For information on specifying data types in JSON, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataFormat.html">JSON Data Format</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Comparisonoperator interface{} `json:"ComparisonOperator,omitempty"` // <p>A comparator for evaluating attributes in the <code>AttributeValueList</code>. For example, equals, greater than, less than, etc.</p> <p>The following comparison operators are available:</p> <p> <code>EQ | NE | LE | LT | GE | GT | NOT_NULL | NULL | CONTAINS | NOT_CONTAINS | BEGINS_WITH | IN | BETWEEN</code> </p> <p>The following are descriptions of each comparison operator.</p> <ul> <li> <p> <code>EQ</code> : Equal. <code>EQ</code> is supported for all data types, including lists and maps.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, Binary, String Set, Number Set, or Binary Set. If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not equal <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>NE</code> : Not equal. <code>NE</code> is supported for all data types, including lists and maps.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> of type String, Number, Binary, String Set, Number Set, or Binary Set. If an item contains an <code>AttributeValue</code> of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not equal <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>LE</code> : Less than or equal. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>LT</code> : Less than. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>GE</code> : Greater than or equal. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>GT</code> : Greater than. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not equal <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code>.</p> <p/> </li> <li> <p> <code>NOT_NULL</code> : The attribute exists. <code>NOT_NULL</code> is supported for all data types, including lists and maps.</p> <note> <p>This operator tests for the existence of an attribute, not its data type. If the data type of attribute "<code>a</code>" is null, and you evaluate it using <code>NOT_NULL</code>, the result is a Boolean <code>true</code>. This result is because the attribute "<code>a</code>" exists; its data type is not relevant to the <code>NOT_NULL</code> comparison operator.</p> </note> </li> <li> <p> <code>NULL</code> : The attribute does not exist. <code>NULL</code> is supported for all data types, including lists and maps.</p> <note> <p>This operator tests for the nonexistence of an attribute, not its data type. If the data type of attribute "<code>a</code>" is null, and you evaluate it using <code>NULL</code>, the result is a Boolean <code>false</code>. This is because the attribute "<code>a</code>" exists; its data type is not relevant to the <code>NULL</code> comparison operator.</p> </note> </li> <li> <p> <code>CONTAINS</code> : Checks for a subsequence, or value in a set.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If the target attribute of the comparison is of type String, then the operator checks for a substring match. If the target attribute of the comparison is of type Binary, then the operator looks for a subsequence of the target that matches the input. If the target attribute of the comparison is a set ("<code>SS</code>", "<code>NS</code>", or "<code>BS</code>"), then the operator evaluates to true if it finds an exact match with any member of the set.</p> <p>CONTAINS is supported for lists: When evaluating "<code>a CONTAINS b</code>", "<code>a</code>" can be a list; however, "<code>b</code>" cannot be a set, a map, or a list.</p> </li> <li> <p> <code>NOT_CONTAINS</code> : Checks for absence of a subsequence, or absence of a value in a set.</p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> element of type String, Number, or Binary (not a set type). If the target attribute of the comparison is a String, then the operator checks for the absence of a substring match. If the target attribute of the comparison is Binary, then the operator checks for the absence of a subsequence of the target that matches the input. If the target attribute of the comparison is a set ("<code>SS</code>", "<code>NS</code>", or "<code>BS</code>"), then the operator evaluates to true if it <i>does not</i> find an exact match with any member of the set.</p> <p>NOT_CONTAINS is supported for lists: When evaluating "<code>a NOT CONTAINS b</code>", "<code>a</code>" can be a list; however, "<code>b</code>" cannot be a set, a map, or a list.</p> </li> <li> <p> <code>BEGINS_WITH</code> : Checks for a prefix. </p> <p> <code>AttributeValueList</code> can contain only one <code>AttributeValue</code> of type String or Binary (not a Number or a set type). The target attribute of the comparison must be of type String or Binary (not a Number or a set type).</p> <p/> </li> <li> <p> <code>IN</code> : Checks for matching elements in a list.</p> <p> <code>AttributeValueList</code> can contain one or more <code>AttributeValue</code> elements of type String, Number, or Binary. These attributes are compared against an existing attribute of an item. If any elements of the input are equal to the item attribute, the expression evaluates to true.</p> </li> <li> <p> <code>BETWEEN</code> : Greater than or equal to the first value, and less than or equal to the second value. </p> <p> <code>AttributeValueList</code> must contain two <code>AttributeValue</code> elements of the same type, either String, Number, or Binary (not a set type). A target attribute matches if the target value is greater than, or equal to, the first element and less than, or equal to, the second element. If an item contains an <code>AttributeValue</code> element of a different type than the one provided in the request, the value does not match. For example, <code>{"S":"6"}</code> does not compare to <code>{"N":"6"}</code>. Also, <code>{"N":"6"}</code> does not compare to <code>{"NS":["6", "2", "1"]}</code> </p> </li> </ul>
	Exists interface{} `json:"Exists,omitempty"` // <p>Causes DynamoDB to evaluate the value before attempting a conditional operation:</p> <ul> <li> <p>If <code>Exists</code> is <code>true</code>, DynamoDB will check to see if that attribute value already exists in the table. If it is found, then the operation succeeds. If it is not found, the operation fails with a <code>ConditionCheckFailedException</code>.</p> </li> <li> <p>If <code>Exists</code> is <code>false</code>, DynamoDB assumes that the attribute value does not exist in the table. If in fact the value does not exist, then the assumption is valid and the operation succeeds. If the value is found, despite the assumption that it does not exist, the operation fails with a <code>ConditionCheckFailedException</code>.</p> </li> </ul> <p>The default setting for <code>Exists</code> is <code>true</code>. If you supply a <code>Value</code> all by itself, DynamoDB assumes the attribute exists: You don't have to set <code>Exists</code> to <code>true</code>, because it is implied.</p> <p>DynamoDB returns a <code>ValidationException</code> if:</p> <ul> <li> <p> <code>Exists</code> is <code>true</code> but there is no <code>Value</code> to check. (You expect a value to exist, but don't specify what that value is.)</p> </li> <li> <p> <code>Exists</code> is <code>false</code> but you also provide a <code>Value</code>. (You cannot expect an attribute to have a value, while also expecting it not to exist.)</p> </li> </ul>
	Value interface{} `json:"Value,omitempty"` // <p>Represents the data for the expected attribute.</p> <p>Each attribute value is described as a name-value pair. The name is the data type, and the value is the data itself.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.NamingRulesDataTypes.html#HowItWorks.DataTypes">Data Types</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
}

// UpdateTimeToLiveOutput represents the UpdateTimeToLiveOutput schema from the OpenAPI specification
type UpdateTimeToLiveOutput struct {
	Timetolivespecification interface{} `json:"TimeToLiveSpecification,omitempty"` // Represents the output of an <code>UpdateTimeToLive</code> operation.
}

// PutItemInputAttributeMap represents the PutItemInputAttributeMap schema from the OpenAPI specification
type PutItemInputAttributeMap struct {
}

// PutRequest represents the PutRequest schema from the OpenAPI specification
type PutRequest struct {
	Item interface{} `json:"Item"` // A map of attribute name to attribute values, representing the primary key of an item to be processed by <code>PutItem</code>. All of the table's primary key attributes must be specified, and their data types must match those of the table's key schema. If any attributes are present in the item that are part of an index key schema for the table, their types must match the index key schema.
}

// BatchGetItemInput represents the BatchGetItemInput schema from the OpenAPI specification
type BatchGetItemInput struct {
	Requestitems interface{} `json:"RequestItems"` // <p>A map of one or more table names and, for each table, a map that describes one or more items to retrieve from that table. Each table name can be used only once per <code>BatchGetItem</code> request.</p> <p>Each element in the map of items to retrieve consists of the following:</p> <ul> <li> <p> <code>ConsistentRead</code> - If <code>true</code>, a strongly consistent read is used; if <code>false</code> (the default), an eventually consistent read is used.</p> </li> <li> <p> <code>ExpressionAttributeNames</code> - One or more substitution tokens for attribute names in the <code>ProjectionExpression</code> parameter. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information about expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Accessing Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> </li> <li> <p> <code>Keys</code> - An array of primary key attribute values that define specific items in the table. For each primary key, you must provide <i>all</i> of the key attributes. For example, with a simple primary key, you only need to provide the partition key value. For a composite key, you must provide <i>both</i> the partition key value and the sort key value.</p> </li> <li> <p> <code>ProjectionExpression</code> - A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.</p> <p>If no attribute names are specified, then all attributes are returned. If any of the requested attributes are not found, they do not appear in the result.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Accessing Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> </li> <li> <p> <code>AttributesToGet</code> - This is a legacy parameter. Use <code>ProjectionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html">AttributesToGet</a> in the <i>Amazon DynamoDB Developer Guide</i>. </p> </li> </ul>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
}

// MapAttributeValue represents the MapAttributeValue schema from the OpenAPI specification
type MapAttributeValue struct {
}

// SourceTableDetails represents the SourceTableDetails schema from the OpenAPI specification
type SourceTableDetails struct {
	Provisionedthroughput interface{} `json:"ProvisionedThroughput"` // Read IOPs and Write IOPS on the table when the backup was created.
	Tableid interface{} `json:"TableId"` // Unique identifier for the table for which the backup was created.
	Tablesizebytes interface{} `json:"TableSizeBytes,omitempty"` // Size of the table in bytes. Note that this is an approximate value.
	Billingmode interface{} `json:"BillingMode,omitempty"` // <p>Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.</p> <ul> <li> <p> <code>PROVISIONED</code> - Sets the read/write capacity mode to <code>PROVISIONED</code>. We recommend using <code>PROVISIONED</code> for predictable workloads.</p> </li> <li> <p> <code>PAY_PER_REQUEST</code> - Sets the read/write capacity mode to <code>PAY_PER_REQUEST</code>. We recommend using <code>PAY_PER_REQUEST</code> for unpredictable workloads. </p> </li> </ul>
	Keyschema interface{} `json:"KeySchema"` // Schema of the table.
	Tablecreationdatetime interface{} `json:"TableCreationDateTime"` // Time when the source table was created.
	Tablename interface{} `json:"TableName"` // The name of the table for which the backup was created.
	Itemcount interface{} `json:"ItemCount,omitempty"` // Number of items in the table. Note that this is an approximate value.
	Tablearn interface{} `json:"TableArn,omitempty"` // ARN of the table for which backup was created.
}

// DeleteTableOutput represents the DeleteTableOutput schema from the OpenAPI specification
type DeleteTableOutput struct {
	Tabledescription interface{} `json:"TableDescription,omitempty"` // Represents the properties of a table.
}

// ListTagsOfResourceInput represents the ListTagsOfResourceInput schema from the OpenAPI specification
type ListTagsOfResourceInput struct {
	Resourcearn interface{} `json:"ResourceArn"` // The Amazon DynamoDB resource with tags to be listed. This value is an Amazon Resource Name (ARN).
	Nexttoken interface{} `json:"NextToken,omitempty"` // An optional string that, if supplied, must be copied from the output of a previous call to ListTagOfResource. When provided in this manner, this API fetches the next page of results.
}

// DescribeBackupInput represents the DescribeBackupInput schema from the OpenAPI specification
type DescribeBackupInput struct {
	Backuparn interface{} `json:"BackupArn"` // The Amazon Resource Name (ARN) associated with the backup.
}

// DescribeEndpointsResponse represents the DescribeEndpointsResponse schema from the OpenAPI specification
type DescribeEndpointsResponse struct {
	Endpoints interface{} `json:"Endpoints"` // List of endpoints.
}

// ReplicaSettingsDescription represents the ReplicaSettingsDescription schema from the OpenAPI specification
type ReplicaSettingsDescription struct {
	Regionname interface{} `json:"RegionName"` // The Region name of the replica.
	Replicaprovisionedwritecapacityunits interface{} `json:"ReplicaProvisionedWriteCapacityUnits,omitempty"` // The maximum number of writes consumed per second before DynamoDB returns a <code>ThrottlingException</code>. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput">Specifying Read and Write Requirements</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Replicaprovisionedreadcapacityautoscalingsettings interface{} `json:"ReplicaProvisionedReadCapacityAutoScalingSettings,omitempty"` // Auto scaling settings for a global table replica's read capacity units.
	Replicastatus interface{} `json:"ReplicaStatus,omitempty"` // <p>The current state of the Region:</p> <ul> <li> <p> <code>CREATING</code> - The Region is being created.</p> </li> <li> <p> <code>UPDATING</code> - The Region is being updated.</p> </li> <li> <p> <code>DELETING</code> - The Region is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The Region is ready for use.</p> </li> </ul>
	Replicabillingmodesummary interface{} `json:"ReplicaBillingModeSummary,omitempty"` // The read/write capacity mode of the replica.
	Replicaprovisionedwritecapacityautoscalingsettings interface{} `json:"ReplicaProvisionedWriteCapacityAutoScalingSettings,omitempty"` // Auto scaling settings for a global table replica's write capacity units.
	Replicatableclasssummary TableClassSummary `json:"ReplicaTableClassSummary,omitempty"` // Contains details of the table class.
	Replicaglobalsecondaryindexsettings interface{} `json:"ReplicaGlobalSecondaryIndexSettings,omitempty"` // Replica global secondary index settings for the global table.
	Replicaprovisionedreadcapacityunits interface{} `json:"ReplicaProvisionedReadCapacityUnits,omitempty"` // The maximum number of strongly consistent reads consumed per second before DynamoDB returns a <code>ThrottlingException</code>. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput">Specifying Read and Write Requirements</a> in the <i>Amazon DynamoDB Developer Guide</i>.
}

// AutoScalingSettingsDescription represents the AutoScalingSettingsDescription schema from the OpenAPI specification
type AutoScalingSettingsDescription struct {
	Maximumunits interface{} `json:"MaximumUnits,omitempty"` // The maximum capacity units that a global table or global secondary index should be scaled up to.
	Minimumunits interface{} `json:"MinimumUnits,omitempty"` // The minimum capacity units that a global table or global secondary index should be scaled down to.
	Scalingpolicies interface{} `json:"ScalingPolicies,omitempty"` // Information about the scaling policies.
	Autoscalingdisabled interface{} `json:"AutoScalingDisabled,omitempty"` // Disabled auto scaling for this global table or global secondary index.
	Autoscalingrolearn interface{} `json:"AutoScalingRoleArn,omitempty"` // Role ARN used for configuring the auto scaling policy.
}

// KinesisStreamingDestinationOutput represents the KinesisStreamingDestinationOutput schema from the OpenAPI specification
type KinesisStreamingDestinationOutput struct {
	Destinationstatus interface{} `json:"DestinationStatus,omitempty"` // The current status of the replication.
	Streamarn interface{} `json:"StreamArn,omitempty"` // The ARN for the specific Kinesis data stream.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table being modified.
}

// DescribeBackupOutput represents the DescribeBackupOutput schema from the OpenAPI specification
type DescribeBackupOutput struct {
	Backupdescription interface{} `json:"BackupDescription,omitempty"` // Contains the description of the backup created for the table.
}

// SSESpecification represents the SSESpecification schema from the OpenAPI specification
type SSESpecification struct {
	Kmsmasterkeyid interface{} `json:"KMSMasterKeyId,omitempty"` // The KMS key that should be used for the KMS encryption. To specify a key, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. Note that you should only provide this parameter if the key is different from the default DynamoDB key <code>alias/aws/dynamodb</code>.
	Ssetype interface{} `json:"SSEType,omitempty"` // <p>Server-side encryption type. The only supported value is:</p> <ul> <li> <p> <code>KMS</code> - Server-side encryption that uses Key Management Service. The key is stored in your account and is managed by KMS (KMS charges apply).</p> </li> </ul>
	Enabled interface{} `json:"Enabled,omitempty"` // Indicates whether server-side encryption is done using an Amazon Web Services managed key or an Amazon Web Services owned key. If enabled (true), server-side encryption type is set to <code>KMS</code> and an Amazon Web Services managed key is used (KMS charges apply). If disabled (false) or not specified, server-side encryption is set to Amazon Web Services owned key.
}

// ExportSummary represents the ExportSummary schema from the OpenAPI specification
type ExportSummary struct {
	Exportarn interface{} `json:"ExportArn,omitempty"` // The Amazon Resource Name (ARN) of the export.
	Exportstatus interface{} `json:"ExportStatus,omitempty"` // Export can be in one of the following states: IN_PROGRESS, COMPLETED, or FAILED.
}

// ListBackupsInput represents the ListBackupsInput schema from the OpenAPI specification
type ListBackupsInput struct {
	Limit interface{} `json:"Limit,omitempty"` // Maximum number of backups to return at once.
	Tablename interface{} `json:"TableName,omitempty"` // The backups from the table specified by <code>TableName</code> are listed.
	Timerangelowerbound interface{} `json:"TimeRangeLowerBound,omitempty"` // Only backups created after this time are listed. <code>TimeRangeLowerBound</code> is inclusive.
	Timerangeupperbound interface{} `json:"TimeRangeUpperBound,omitempty"` // Only backups created before this time are listed. <code>TimeRangeUpperBound</code> is exclusive.
	Backuptype interface{} `json:"BackupType,omitempty"` // <p>The backups from the table specified by <code>BackupType</code> are listed.</p> <p>Where <code>BackupType</code> can be:</p> <ul> <li> <p> <code>USER</code> - On-demand backup created by you. (The default setting if no other backup types are specified.)</p> </li> <li> <p> <code>SYSTEM</code> - On-demand backup automatically created by DynamoDB.</p> </li> <li> <p> <code>ALL</code> - All types of on-demand backups (USER and SYSTEM).</p> </li> </ul>
	Exclusivestartbackuparn interface{} `json:"ExclusiveStartBackupArn,omitempty"` // <code>LastEvaluatedBackupArn</code> is the Amazon Resource Name (ARN) of the backup last evaluated when the current page of results was returned, inclusive of the current page of results. This value may be specified as the <code>ExclusiveStartBackupArn</code> of a new <code>ListBackups</code> operation in order to fetch the next page of results.
}

// CreateReplicaAction represents the CreateReplicaAction schema from the OpenAPI specification
type CreateReplicaAction struct {
	Regionname interface{} `json:"RegionName"` // The Region of the replica to be added.
}

// ExportTableToPointInTimeOutput represents the ExportTableToPointInTimeOutput schema from the OpenAPI specification
type ExportTableToPointInTimeOutput struct {
	Exportdescription interface{} `json:"ExportDescription,omitempty"` // Contains a description of the table export.
}

// Endpoint represents the Endpoint schema from the OpenAPI specification
type Endpoint struct {
	Address interface{} `json:"Address"` // IP address of the endpoint.
	Cacheperiodinminutes interface{} `json:"CachePeriodInMinutes"` // Endpoint cache time to live (TTL) value.
}

// ReplicaGlobalSecondaryIndexAutoScalingDescription represents the ReplicaGlobalSecondaryIndexAutoScalingDescription schema from the OpenAPI specification
type ReplicaGlobalSecondaryIndexAutoScalingDescription struct {
	Indexstatus interface{} `json:"IndexStatus,omitempty"` // <p>The current state of the replica global secondary index:</p> <ul> <li> <p> <code>CREATING</code> - The index is being created.</p> </li> <li> <p> <code>UPDATING</code> - The table/index configuration is being updated. The table/index remains available for data operations when <code>UPDATING</code> </p> </li> <li> <p> <code>DELETING</code> - The index is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The index is ready for use.</p> </li> </ul>
	Provisionedreadcapacityautoscalingsettings AutoScalingSettingsDescription `json:"ProvisionedReadCapacityAutoScalingSettings,omitempty"` // Represents the auto scaling settings for a global table or global secondary index.
	Provisionedwritecapacityautoscalingsettings AutoScalingSettingsDescription `json:"ProvisionedWriteCapacityAutoScalingSettings,omitempty"` // Represents the auto scaling settings for a global table or global secondary index.
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index.
}

// DescribeImportInput represents the DescribeImportInput schema from the OpenAPI specification
type DescribeImportInput struct {
	Importarn interface{} `json:"ImportArn"` // The Amazon Resource Name (ARN) associated with the table you're importing to.
}

// BatchGetResponseMap represents the BatchGetResponseMap schema from the OpenAPI specification
type BatchGetResponseMap struct {
}

// GlobalTableGlobalSecondaryIndexSettingsUpdate represents the GlobalTableGlobalSecondaryIndexSettingsUpdate schema from the OpenAPI specification
type GlobalTableGlobalSecondaryIndexSettingsUpdate struct {
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index. The name must be unique among all other indexes on this table.
	Provisionedwritecapacityautoscalingsettingsupdate interface{} `json:"ProvisionedWriteCapacityAutoScalingSettingsUpdate,omitempty"` // Auto scaling settings for managing a global secondary index's write capacity units.
	Provisionedwritecapacityunits interface{} `json:"ProvisionedWriteCapacityUnits,omitempty"` // The maximum number of writes consumed per second before DynamoDB returns a <code>ThrottlingException.</code>
}

// BatchExecuteStatementInput represents the BatchExecuteStatementInput schema from the OpenAPI specification
type BatchExecuteStatementInput struct {
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Statements interface{} `json:"Statements"` // The list of PartiQL statements representing the batch to run.
}

// DescribeContributorInsightsOutput represents the DescribeContributorInsightsOutput schema from the OpenAPI specification
type DescribeContributorInsightsOutput struct {
	Contributorinsightsrulelist interface{} `json:"ContributorInsightsRuleList,omitempty"` // List of names of the associated contributor insights rules.
	Contributorinsightsstatus interface{} `json:"ContributorInsightsStatus,omitempty"` // Current status of contributor insights.
	Failureexception interface{} `json:"FailureException,omitempty"` // <p>Returns information about the last failure that was encountered.</p> <p>The most common exceptions for a FAILED status are:</p> <ul> <li> <p>LimitExceededException - Per-account Amazon CloudWatch Contributor Insights rule limit reached. Please disable Contributor Insights for other tables/indexes OR disable Contributor Insights rules before retrying.</p> </li> <li> <p>AccessDeniedException - Amazon CloudWatch Contributor Insights rules cannot be modified due to insufficient permissions.</p> </li> <li> <p>AccessDeniedException - Failed to create service-linked role for Contributor Insights due to insufficient permissions.</p> </li> <li> <p>InternalServerError - Failed to create Amazon CloudWatch Contributor Insights rules. Please retry request.</p> </li> </ul>
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index being described.
	Lastupdatedatetime interface{} `json:"LastUpdateDateTime,omitempty"` // Timestamp of the last time the status was changed.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table being described.
}

// DescribeExportOutput represents the DescribeExportOutput schema from the OpenAPI specification
type DescribeExportOutput struct {
	Exportdescription interface{} `json:"ExportDescription,omitempty"` // Represents the properties of the export.
}

// ImportTableDescription represents the ImportTableDescription schema from the OpenAPI specification
type ImportTableDescription struct {
	Failuremessage interface{} `json:"FailureMessage,omitempty"` // The error message corresponding to the failure that the import job ran into during execution.
	Tablecreationparameters interface{} `json:"TableCreationParameters,omitempty"` // The parameters for the new table that is being imported into.
	Tableid interface{} `json:"TableId,omitempty"` // The table id corresponding to the table created by import table process.
	Processedsizebytes interface{} `json:"ProcessedSizeBytes,omitempty"` // The total size of data processed from the source file, in Bytes.
	Importarn interface{} `json:"ImportArn,omitempty"` // The Amazon Resource Number (ARN) corresponding to the import request.
	Inputformat interface{} `json:"InputFormat,omitempty"` // The format of the source data going into the target table.
	S3bucketsource interface{} `json:"S3BucketSource,omitempty"` // Values for the S3 bucket the source file is imported from. Includes bucket name (required), key prefix (optional) and bucket account owner ID (optional).
	Clienttoken interface{} `json:"ClientToken,omitempty"` // The client token that was provided for the import task. Reusing the client token on retry makes a call to <code>ImportTable</code> idempotent.
	Importeditemcount interface{} `json:"ImportedItemCount,omitempty"` // The number of items successfully imported into the new table.
	Starttime interface{} `json:"StartTime,omitempty"` // The time when this import task started.
	Cloudwatchloggrouparn interface{} `json:"CloudWatchLogGroupArn,omitempty"` // The Amazon Resource Number (ARN) of the Cloudwatch Log Group associated with the target table.
	Errorcount interface{} `json:"ErrorCount,omitempty"` // The number of errors occurred on importing the source file into the target table.
	Inputcompressiontype interface{} `json:"InputCompressionType,omitempty"` // The compression options for the data that has been imported into the target table. The values are NONE, GZIP, or ZSTD.
	Inputformatoptions interface{} `json:"InputFormatOptions,omitempty"` // The format options for the data that was imported into the target table. There is one value, CsvOption.
	Tablearn interface{} `json:"TableArn,omitempty"` // The Amazon Resource Number (ARN) of the table being imported into.
	Importstatus interface{} `json:"ImportStatus,omitempty"` // The status of the import.
	Endtime interface{} `json:"EndTime,omitempty"` // The time at which the creation of the table associated with this import task completed.
	Failurecode interface{} `json:"FailureCode,omitempty"` // The error code corresponding to the failure that the import job ran into during execution.
	Processeditemcount interface{} `json:"ProcessedItemCount,omitempty"` // The total number of items processed from the source file.
}

// S3BucketSource represents the S3BucketSource schema from the OpenAPI specification
type S3BucketSource struct {
	S3bucket interface{} `json:"S3Bucket"` // The S3 bucket that is being imported from.
	S3bucketowner interface{} `json:"S3BucketOwner,omitempty"` // The account number of the S3 bucket that is being imported from. If the bucket is owned by the requester this is optional.
	S3keyprefix interface{} `json:"S3KeyPrefix,omitempty"` // The key prefix shared by all S3 Objects that are being imported.
}

// DescribeTableOutput represents the DescribeTableOutput schema from the OpenAPI specification
type DescribeTableOutput struct {
	Table interface{} `json:"Table,omitempty"` // The properties of the table.
}

// ListGlobalTablesInput represents the ListGlobalTablesInput schema from the OpenAPI specification
type ListGlobalTablesInput struct {
	Exclusivestartglobaltablename interface{} `json:"ExclusiveStartGlobalTableName,omitempty"` // The first global table name that this operation will evaluate.
	Limit interface{} `json:"Limit,omitempty"` // <p>The maximum number of table names to return, if the parameter is not specified DynamoDB defaults to 100.</p> <p>If the number of global tables DynamoDB finds reaches this limit, it stops the operation and returns the table names collected up to that point, with a table name in the <code>LastEvaluatedGlobalTableName</code> to apply in a subsequent operation to the <code>ExclusiveStartGlobalTableName</code> parameter.</p>
	Regionname interface{} `json:"RegionName,omitempty"` // Lists the global tables in a specific Region.
}

// UpdateGlobalTableSettingsInput represents the UpdateGlobalTableSettingsInput schema from the OpenAPI specification
type UpdateGlobalTableSettingsInput struct {
	Globaltableprovisionedwritecapacityunits interface{} `json:"GlobalTableProvisionedWriteCapacityUnits,omitempty"` // The maximum number of writes consumed per second before DynamoDB returns a <code>ThrottlingException.</code>
	Replicasettingsupdate interface{} `json:"ReplicaSettingsUpdate,omitempty"` // Represents the settings for a global table in a Region that will be modified.
	Globaltablebillingmode interface{} `json:"GlobalTableBillingMode,omitempty"` // <p>The billing mode of the global table. If <code>GlobalTableBillingMode</code> is not specified, the global table defaults to <code>PROVISIONED</code> capacity billing mode.</p> <ul> <li> <p> <code>PROVISIONED</code> - We recommend using <code>PROVISIONED</code> for predictable workloads. <code>PROVISIONED</code> sets the billing mode to <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual">Provisioned Mode</a>.</p> </li> <li> <p> <code>PAY_PER_REQUEST</code> - We recommend using <code>PAY_PER_REQUEST</code> for unpredictable workloads. <code>PAY_PER_REQUEST</code> sets the billing mode to <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand">On-Demand Mode</a>. </p> </li> </ul>
	Globaltableglobalsecondaryindexsettingsupdate interface{} `json:"GlobalTableGlobalSecondaryIndexSettingsUpdate,omitempty"` // Represents the settings of a global secondary index for a global table that will be modified.
	Globaltablename interface{} `json:"GlobalTableName"` // The name of the global table
	Globaltableprovisionedwritecapacityautoscalingsettingsupdate interface{} `json:"GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate,omitempty"` // Auto scaling settings for managing provisioned write capacity for the global table.
}

// FilterConditionMap represents the FilterConditionMap schema from the OpenAPI specification
type FilterConditionMap struct {
}

// GlobalSecondaryIndexInfo represents the GlobalSecondaryIndexInfo schema from the OpenAPI specification
type GlobalSecondaryIndexInfo struct {
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"` // Represents the provisioned throughput settings for the specified global secondary index.
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index.
	Keyschema interface{} `json:"KeySchema,omitempty"` // <p>The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note>
	Projection interface{} `json:"Projection,omitempty"` // Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
}

// GlobalSecondaryIndexDescription represents the GlobalSecondaryIndexDescription schema from the OpenAPI specification
type GlobalSecondaryIndexDescription struct {
	Indexstatus interface{} `json:"IndexStatus,omitempty"` // <p>The current state of the global secondary index:</p> <ul> <li> <p> <code>CREATING</code> - The index is being created.</p> </li> <li> <p> <code>UPDATING</code> - The index is being updated.</p> </li> <li> <p> <code>DELETING</code> - The index is being deleted.</p> </li> <li> <p> <code>ACTIVE</code> - The index is ready for use.</p> </li> </ul>
	Indexarn interface{} `json:"IndexArn,omitempty"` // The Amazon Resource Name (ARN) that uniquely identifies the index.
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index.
	Projection interface{} `json:"Projection,omitempty"` // Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Indexsizebytes interface{} `json:"IndexSizeBytes,omitempty"` // The total size of the specified index, in bytes. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"` // <p>Represents the provisioned throughput settings for the specified global secondary index.</p> <p>For current minimum and maximum provisioned throughput values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Itemcount interface{} `json:"ItemCount,omitempty"` // The number of items in the specified index. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.
	Keyschema interface{} `json:"KeySchema,omitempty"` // <p>The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note>
	Backfilling interface{} `json:"Backfilling,omitempty"` // <p>Indicates whether the index is currently backfilling. <i>Backfilling</i> is the process of reading items from the table and determining whether they can be added to the index. (Not all items will qualify: For example, a partition key cannot have any duplicate values.) If an item can be added to the index, DynamoDB will do so. After all items have been processed, the backfilling operation is complete and <code>Backfilling</code> is false.</p> <p>You can delete an index that is being created during the <code>Backfilling</code> phase when <code>IndexStatus</code> is set to CREATING and <code>Backfilling</code> is true. You can't delete the index that is being created when <code>IndexStatus</code> is set to CREATING and <code>Backfilling</code> is false. </p> <note> <p>For indexes that were created during a <code>CreateTable</code> operation, the <code>Backfilling</code> attribute does not appear in the <code>DescribeTable</code> output.</p> </note>
}

// ExecuteTransactionOutput represents the ExecuteTransactionOutput schema from the OpenAPI specification
type ExecuteTransactionOutput struct {
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the entire operation. The values of the list are ordered according to the ordering of the statements.
	Responses interface{} `json:"Responses,omitempty"` // The response to a PartiQL transaction.
}

// TagResourceInput represents the TagResourceInput schema from the OpenAPI specification
type TagResourceInput struct {
	Tags interface{} `json:"Tags"` // The tags to be assigned to the Amazon DynamoDB resource.
	Resourcearn interface{} `json:"ResourceArn"` // Identifies the Amazon DynamoDB resource to which tags should be added. This value is an Amazon Resource Name (ARN).
}

// UpdateTimeToLiveInput represents the UpdateTimeToLiveInput schema from the OpenAPI specification
type UpdateTimeToLiveInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table to be configured.
	Timetolivespecification interface{} `json:"TimeToLiveSpecification"` // Represents the settings used to enable or disable Time to Live for the specified table.
}

// UpdateItemInput represents the UpdateItemInput schema from the OpenAPI specification
type UpdateItemInput struct {
	Key interface{} `json:"Key"` // <p>The primary key of the item to be updated. Each element consists of an attribute name and a value for that attribute.</p> <p>For the primary key, you must provide all of the attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for both the partition key and the sort key.</p>
	Returnvalues interface{} `json:"ReturnValues,omitempty"` // <p>Use <code>ReturnValues</code> if you want to get the item attributes as they appear before or after they are successfully updated. For <code>UpdateItem</code>, the valid values are:</p> <ul> <li> <p> <code>NONE</code> - If <code>ReturnValues</code> is not specified, or if its value is <code>NONE</code>, then nothing is returned. (This setting is the default for <code>ReturnValues</code>.)</p> </li> <li> <p> <code>ALL_OLD</code> - Returns all of the attributes of the item, as they appeared before the UpdateItem operation.</p> </li> <li> <p> <code>UPDATED_OLD</code> - Returns only the updated attributes, as they appeared before the UpdateItem operation.</p> </li> <li> <p> <code>ALL_NEW</code> - Returns all of the attributes of the item, as they appear after the UpdateItem operation.</p> </li> <li> <p> <code>UPDATED_NEW</code> - Returns only the updated attributes, as they appear after the UpdateItem operation.</p> </li> </ul> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p> <p>The values returned are strongly consistent.</p>
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // <p>An optional parameter that returns the item attributes for an <code>UpdateItem</code> operation that failed a condition check.</p> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p>
	Conditionaloperator interface{} `json:"ConditionalOperator,omitempty"` // This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html">ConditionalOperator</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Expected interface{} `json:"Expected,omitempty"` // This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html">Expected</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Returnitemcollectionmetrics interface{} `json:"ReturnItemCollectionMetrics,omitempty"` // Determines whether item collection metrics are returned. If set to <code>SIZE</code>, the response includes statistics about item collections, if any, that were modified during the operation are returned in the response. If set to <code>NONE</code> (the default), no statistics are returned.
	Tablename interface{} `json:"TableName"` // The name of the table containing the item to update.
	Attributeupdates interface{} `json:"AttributeUpdates,omitempty"` // This is a legacy parameter. Use <code>UpdateExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributeUpdates.html">AttributeUpdates</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Conditionexpression interface{} `json:"ConditionExpression,omitempty"` // <p>A condition that must be satisfied in order for a conditional update to succeed.</p> <p>An expression can contain any of the following:</p> <ul> <li> <p>Functions: <code>attribute_exists | attribute_not_exists | attribute_type | contains | begins_with | size</code> </p> <p>These function names are case-sensitive.</p> </li> <li> <p>Comparison operators: <code>= | &lt;&gt; | &lt; | &gt; | &lt;= | &gt;= | BETWEEN | IN </code> </p> </li> <li> <p> Logical operators: <code>AND | OR | NOT</code> </p> </li> </ul> <p>For more information about condition expressions, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Specifying Conditions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>.) To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information about expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Updateexpression interface{} `json:"UpdateExpression,omitempty"` // <p>An expression that defines one or more attributes to be updated, the action to be performed on them, and new values for them.</p> <p>The following action values are available for <code>UpdateExpression</code>.</p> <ul> <li> <p> <code>SET</code> - Adds one or more attributes and values to an item. If any of these attributes already exist, they are replaced by the new values. You can also use <code>SET</code> to add or subtract from an attribute that is of type Number. For example: <code>SET myNum = myNum + :val</code> </p> <p> <code>SET</code> supports the following functions:</p> <ul> <li> <p> <code>if_not_exists (path, operand)</code> - if the item does not contain an attribute at the specified path, then <code>if_not_exists</code> evaluates to operand; otherwise, it evaluates to path. You can use this function to avoid overwriting an attribute that may already be present in the item.</p> </li> <li> <p> <code>list_append (operand, operand)</code> - evaluates to a list with a new element added to it. You can append the new element to the start or the end of the list by reversing the order of the operands.</p> </li> </ul> <p>These function names are case-sensitive.</p> </li> <li> <p> <code>REMOVE</code> - Removes one or more attributes from an item.</p> </li> <li> <p> <code>ADD</code> - Adds the specified value to the item, if the attribute does not already exist. If the attribute does exist, then the behavior of <code>ADD</code> depends on the data type of the attribute:</p> <ul> <li> <p>If the existing attribute is a number, and if <code>Value</code> is also a number, then <code>Value</code> is mathematically added to the existing attribute. If <code>Value</code> is a negative number, then it is subtracted from the existing attribute.</p> <note> <p>If you use <code>ADD</code> to increment or decrement a number value for an item that doesn't exist before the update, DynamoDB uses <code>0</code> as the initial value.</p> <p>Similarly, if you use <code>ADD</code> for an existing item to increment or decrement an attribute value that doesn't exist before the update, DynamoDB uses <code>0</code> as the initial value. For example, suppose that the item you want to update doesn't have an attribute named <code>itemcount</code>, but you decide to <code>ADD</code> the number <code>3</code> to this attribute anyway. DynamoDB will create the <code>itemcount</code> attribute, set its initial value to <code>0</code>, and finally add <code>3</code> to it. The result will be a new <code>itemcount</code> attribute in the item, with a value of <code>3</code>.</p> </note> </li> <li> <p>If the existing data type is a set and if <code>Value</code> is also a set, then <code>Value</code> is added to the existing set. For example, if the attribute value is the set <code>[1,2]</code>, and the <code>ADD</code> action specified <code>[3]</code>, then the final attribute value is <code>[1,2,3]</code>. An error occurs if an <code>ADD</code> action is specified for a set attribute and the attribute type specified does not match the existing set type. </p> <p>Both sets must have the same primitive data type. For example, if the existing data type is a set of strings, the <code>Value</code> must also be a set of strings.</p> </li> </ul> <important> <p>The <code>ADD</code> action only supports Number and set data types. In addition, <code>ADD</code> can only be used on top-level attributes, not nested attributes.</p> </important> </li> <li> <p> <code>DELETE</code> - Deletes an element from a set.</p> <p>If a set of values is specified, then those values are subtracted from the old set. For example, if the attribute value was the set <code>[a,b,c]</code> and the <code>DELETE</code> action specifies <code>[a,c]</code>, then the final attribute value is <code>[b]</code>. Specifying an empty set is an error.</p> <important> <p>The <code>DELETE</code> action only supports set data types. In addition, <code>DELETE</code> can only be used on top-level attributes, not nested attributes.</p> </important> </li> </ul> <p>You can have many actions in a single expression, such as the following: <code>SET a=:value1, b=:value2 DELETE :value3, :value4, :value5</code> </p> <p>For more information on update expressions, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.Modifying.html">Modifying Items and Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // <p>One or more values that can be substituted in an expression.</p> <p>Use the <b>:</b> (colon) character in an expression to dereference an attribute value. For example, suppose that you wanted to check whether the value of the <code>ProductStatus</code> attribute was one of the following: </p> <p> <code>Available | Backordered | Discontinued</code> </p> <p>You would first need to specify <code>ExpressionAttributeValues</code> as follows:</p> <p> <code>{ ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"} }</code> </p> <p>You could then use these values in an expression, such as this:</p> <p> <code>ProductStatus IN (:avail, :back, :disc)</code> </p> <p>For more information on expression attribute values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
}

// DescribeTableInput represents the DescribeTableInput schema from the OpenAPI specification
type DescribeTableInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table to describe.
}

// DescribeEndpointsRequest represents the DescribeEndpointsRequest schema from the OpenAPI specification
type DescribeEndpointsRequest struct {
}

// ListImportsOutput represents the ListImportsOutput schema from the OpenAPI specification
type ListImportsOutput struct {
	Importsummarylist interface{} `json:"ImportSummaryList,omitempty"` // A list of <code>ImportSummary</code> objects.
	Nexttoken interface{} `json:"NextToken,omitempty"` // If this value is returned, there are additional results to be displayed. To retrieve them, call <code>ListImports</code> again, with <code>NextToken</code> set to this value.
}

// DescribeGlobalTableSettingsOutput represents the DescribeGlobalTableSettingsOutput schema from the OpenAPI specification
type DescribeGlobalTableSettingsOutput struct {
	Replicasettings interface{} `json:"ReplicaSettings,omitempty"` // The Region-specific settings for the global table.
	Globaltablename interface{} `json:"GlobalTableName,omitempty"` // The name of the global table.
}

// WriteRequest represents the WriteRequest schema from the OpenAPI specification
type WriteRequest struct {
	Deleterequest interface{} `json:"DeleteRequest,omitempty"` // A request to perform a <code>DeleteItem</code> operation.
	Putrequest interface{} `json:"PutRequest,omitempty"` // A request to perform a <code>PutItem</code> operation.
}

// QueryOutput represents the QueryOutput schema from the OpenAPI specification
type QueryOutput struct {
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the <code>Query</code> operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the <code>ReturnConsumedCapacity</code> parameter was specified. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Count interface{} `json:"Count,omitempty"` // <p>The number of items in the response.</p> <p>If you used a <code>QueryFilter</code> in the request, then <code>Count</code> is the number of items returned after the filter was applied, and <code>ScannedCount</code> is the number of matching items before the filter was applied.</p> <p>If you did not use a filter in the request, then <code>Count</code> and <code>ScannedCount</code> are the same.</p>
	Items interface{} `json:"Items,omitempty"` // An array of item attributes that match the query criteria. Each element in this array consists of an attribute name and the value for that attribute.
	Lastevaluatedkey interface{} `json:"LastEvaluatedKey,omitempty"` // <p>The primary key of the item where the operation stopped, inclusive of the previous result set. Use this value to start a new operation, excluding this value in the new request.</p> <p>If <code>LastEvaluatedKey</code> is empty, then the "last page" of results has been processed and there is no more data to be retrieved.</p> <p>If <code>LastEvaluatedKey</code> is not empty, it does not necessarily mean that there is more data in the result set. The only way to know when you have reached the end of the result set is when <code>LastEvaluatedKey</code> is empty.</p>
	Scannedcount interface{} `json:"ScannedCount,omitempty"` // <p>The number of items evaluated, before any <code>QueryFilter</code> is applied. A high <code>ScannedCount</code> value with few, or no, <code>Count</code> results indicates an inefficient <code>Query</code> operation. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#Count">Count and ScannedCount</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>If you did not use a filter in the request, then <code>ScannedCount</code> is the same as <code>Count</code>.</p>
}

// ReplicaGlobalSecondaryIndexDescription represents the ReplicaGlobalSecondaryIndexDescription schema from the OpenAPI specification
type ReplicaGlobalSecondaryIndexDescription struct {
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index.
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"` // If not described, uses the source table GSI's read capacity settings.
}

// ReplicaAutoScalingUpdate represents the ReplicaAutoScalingUpdate schema from the OpenAPI specification
type ReplicaAutoScalingUpdate struct {
	Replicaglobalsecondaryindexupdates interface{} `json:"ReplicaGlobalSecondaryIndexUpdates,omitempty"` // Represents the auto scaling settings of global secondary indexes that will be modified.
	Replicaprovisionedreadcapacityautoscalingupdate AutoScalingSettingsUpdate `json:"ReplicaProvisionedReadCapacityAutoScalingUpdate,omitempty"` // Represents the auto scaling settings to be modified for a global table or global secondary index.
	Regionname interface{} `json:"RegionName"` // The Region where the replica exists.
}

// AutoScalingTargetTrackingScalingPolicyConfigurationDescription represents the AutoScalingTargetTrackingScalingPolicyConfigurationDescription schema from the OpenAPI specification
type AutoScalingTargetTrackingScalingPolicyConfigurationDescription struct {
	Disablescalein interface{} `json:"DisableScaleIn,omitempty"` // Indicates whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is false.
	Scaleincooldown interface{} `json:"ScaleInCooldown,omitempty"` // The amount of time, in seconds, after a scale in activity completes before another scale in activity can start. The cooldown period is used to block subsequent scale in requests until it has expired. You should scale in conservatively to protect your application's availability. However, if another alarm triggers a scale out policy during the cooldown period after a scale-in, application auto scaling scales out your scalable target immediately.
	Scaleoutcooldown interface{} `json:"ScaleOutCooldown,omitempty"` // The amount of time, in seconds, after a scale out activity completes before another scale out activity can start. While the cooldown period is in effect, the capacity that has been added by the previous scale out event that initiated the cooldown is calculated as part of the desired capacity for the next scale out. You should continuously (but not excessively) scale out.
	Targetvalue interface{} `json:"TargetValue"` // The target value for the metric. The range is 8.515920e-109 to 1.174271e+108 (Base 10) or 2e-360 to 2e360 (Base 2).
}

// BackupSummary represents the BackupSummary schema from the OpenAPI specification
type BackupSummary struct {
	Backuparn interface{} `json:"BackupArn,omitempty"` // ARN associated with the backup.
	Backupcreationdatetime interface{} `json:"BackupCreationDateTime,omitempty"` // Time at which the backup was created.
	Backupname interface{} `json:"BackupName,omitempty"` // Name of the specified backup.
	Backupstatus interface{} `json:"BackupStatus,omitempty"` // Backup can be in one of the following states: CREATING, ACTIVE, DELETED.
	Tablename interface{} `json:"TableName,omitempty"` // Name of the table.
	Tablearn interface{} `json:"TableArn,omitempty"` // ARN associated with the table.
	Backupexpirydatetime interface{} `json:"BackupExpiryDateTime,omitempty"` // Time at which the automatic on-demand backup created by DynamoDB will expire. This <code>SYSTEM</code> on-demand backup expires automatically 35 days after its creation.
	Backupsizebytes interface{} `json:"BackupSizeBytes,omitempty"` // Size of the backup in bytes.
	Tableid interface{} `json:"TableId,omitempty"` // Unique identifier for the table.
	Backuptype interface{} `json:"BackupType,omitempty"` // <p>BackupType:</p> <ul> <li> <p> <code>USER</code> - You create and manage these using the on-demand backup feature.</p> </li> <li> <p> <code>SYSTEM</code> - If you delete a table with point-in-time recovery enabled, a <code>SYSTEM</code> backup is automatically created and is retained for 35 days (at no additional cost). System backups allow you to restore the deleted table to the state it was in just before the point of deletion. </p> </li> <li> <p> <code>AWS_BACKUP</code> - On-demand backup created by you from Backup service.</p> </li> </ul>
}

// TimeToLiveDescription represents the TimeToLiveDescription schema from the OpenAPI specification
type TimeToLiveDescription struct {
	Attributename interface{} `json:"AttributeName,omitempty"` // The name of the TTL attribute for items in the table.
	Timetolivestatus interface{} `json:"TimeToLiveStatus,omitempty"` // The TTL status for the table.
}

// DescribeContributorInsightsInput represents the DescribeContributorInsightsInput schema from the OpenAPI specification
type DescribeContributorInsightsInput struct {
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index to describe, if applicable.
	Tablename interface{} `json:"TableName"` // The name of the table to describe.
}

// TransactGetItemsOutput represents the TransactGetItemsOutput schema from the OpenAPI specification
type TransactGetItemsOutput struct {
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // If the <i>ReturnConsumedCapacity</i> value was <code>TOTAL</code>, this is an array of <code>ConsumedCapacity</code> objects, one for each table addressed by <code>TransactGetItem</code> objects in the <i>TransactItems</i> parameter. These <code>ConsumedCapacity</code> objects report the read-capacity units consumed by the <code>TransactGetItems</code> call in that table.
	Responses interface{} `json:"Responses,omitempty"` // <p>An ordered array of up to 100 <code>ItemResponse</code> objects, each of which corresponds to the <code>TransactGetItem</code> object in the same position in the <i>TransactItems</i> array. Each <code>ItemResponse</code> object contains a Map of the name-value pairs that are the projected attributes of the requested item.</p> <p>If a requested item could not be retrieved, the corresponding <code>ItemResponse</code> object is Null, or if the requested item has no projected attributes, the corresponding <code>ItemResponse</code> object is an empty Map. </p>
}

// ReplicaSettingsUpdate represents the ReplicaSettingsUpdate schema from the OpenAPI specification
type ReplicaSettingsUpdate struct {
	Replicaglobalsecondaryindexsettingsupdate interface{} `json:"ReplicaGlobalSecondaryIndexSettingsUpdate,omitempty"` // Represents the settings of a global secondary index for a global table that will be modified.
	Replicaprovisionedreadcapacityautoscalingsettingsupdate interface{} `json:"ReplicaProvisionedReadCapacityAutoScalingSettingsUpdate,omitempty"` // Auto scaling settings for managing a global table replica's read capacity units.
	Replicaprovisionedreadcapacityunits interface{} `json:"ReplicaProvisionedReadCapacityUnits,omitempty"` // The maximum number of strongly consistent reads consumed per second before DynamoDB returns a <code>ThrottlingException</code>. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#ProvisionedThroughput">Specifying Read and Write Requirements</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Replicatableclass interface{} `json:"ReplicaTableClass,omitempty"` // Replica-specific table class. If not specified, uses the source table's table class.
	Regionname interface{} `json:"RegionName"` // The Region of the replica to be added.
}

// ExecuteStatementInput represents the ExecuteStatementInput schema from the OpenAPI specification
type ExecuteStatementInput struct {
	Statement interface{} `json:"Statement"` // The PartiQL statement representing the operation to run.
	Consistentread interface{} `json:"ConsistentRead,omitempty"` // The consistency of a read operation. If set to <code>true</code>, then a strongly consistent read is used; otherwise, an eventually consistent read is used.
	Limit interface{} `json:"Limit,omitempty"` // The maximum number of items to evaluate (not necessarily the number of matching items). If DynamoDB processes the number of items up to the limit while processing the results, it stops the operation and returns the matching values up to that point, along with a key in <code>LastEvaluatedKey</code> to apply in a subsequent operation so you can pick up where you left off. Also, if the processed dataset size exceeds 1 MB before DynamoDB reaches this limit, it stops the operation and returns the matching values up to the limit, and a key in <code>LastEvaluatedKey</code> to apply in a subsequent operation to continue the operation.
	Nexttoken interface{} `json:"NextToken,omitempty"` // Set this value to get remaining results, if <code>NextToken</code> was returned in the statement response.
	Parameters interface{} `json:"Parameters,omitempty"` // The parameters for the PartiQL statement, if any.
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // <p>An optional parameter that returns the item attributes for an <code>ExecuteStatement</code> operation that failed a condition check.</p> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p>
}

// TableClassSummary represents the TableClassSummary schema from the OpenAPI specification
type TableClassSummary struct {
	Lastupdatedatetime interface{} `json:"LastUpdateDateTime,omitempty"` // The date and time at which the table class was last updated.
	Tableclass interface{} `json:"TableClass,omitempty"` // The table class of the specified table. Valid values are <code>STANDARD</code> and <code>STANDARD_INFREQUENT_ACCESS</code>.
}

// PointInTimeRecoveryDescription represents the PointInTimeRecoveryDescription schema from the OpenAPI specification
type PointInTimeRecoveryDescription struct {
	Pointintimerecoverystatus interface{} `json:"PointInTimeRecoveryStatus,omitempty"` // <p>The current state of point in time recovery:</p> <ul> <li> <p> <code>ENABLED</code> - Point in time recovery is enabled.</p> </li> <li> <p> <code>DISABLED</code> - Point in time recovery is disabled.</p> </li> </ul>
	Earliestrestorabledatetime interface{} `json:"EarliestRestorableDateTime,omitempty"` // Specifies the earliest point in time you can restore your table to. You can restore your table to any point in time during the last 35 days.
	Latestrestorabledatetime interface{} `json:"LatestRestorableDateTime,omitempty"` // <code>LatestRestorableDateTime</code> is typically 5 minutes before the current time.
}

// UpdateContinuousBackupsInput represents the UpdateContinuousBackupsInput schema from the OpenAPI specification
type UpdateContinuousBackupsInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table.
	Pointintimerecoveryspecification interface{} `json:"PointInTimeRecoverySpecification"` // Represents the settings used to enable point in time recovery.
}

// ConsumedCapacity represents the ConsumedCapacity schema from the OpenAPI specification
type ConsumedCapacity struct {
	Readcapacityunits interface{} `json:"ReadCapacityUnits,omitempty"` // The total number of read capacity units consumed by the operation.
	Table interface{} `json:"Table,omitempty"` // The amount of throughput consumed on the table affected by the operation.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table that was affected by the operation.
	Writecapacityunits interface{} `json:"WriteCapacityUnits,omitempty"` // The total number of write capacity units consumed by the operation.
	Capacityunits interface{} `json:"CapacityUnits,omitempty"` // The total number of capacity units consumed by the operation.
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // The amount of throughput consumed on each global index affected by the operation.
	Localsecondaryindexes interface{} `json:"LocalSecondaryIndexes,omitempty"` // The amount of throughput consumed on each local index affected by the operation.
}

// Key represents the Key schema from the OpenAPI specification
type Key struct {
}

// UpdateTableInput represents the UpdateTableInput schema from the OpenAPI specification
type UpdateTableInput struct {
	Billingmode interface{} `json:"BillingMode,omitempty"` // <p>Controls how you are charged for read and write throughput and how you manage capacity. When switching from pay-per-request to provisioned capacity, initial provisioned capacity values must be set. The initial provisioned capacity values are estimated based on the consumed read and write capacity of your table and global secondary indexes over the past 30 minutes.</p> <ul> <li> <p> <code>PROVISIONED</code> - We recommend using <code>PROVISIONED</code> for predictable workloads. <code>PROVISIONED</code> sets the billing mode to <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual">Provisioned Mode</a>.</p> </li> <li> <p> <code>PAY_PER_REQUEST</code> - We recommend using <code>PAY_PER_REQUEST</code> for unpredictable workloads. <code>PAY_PER_REQUEST</code> sets the billing mode to <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand">On-Demand Mode</a>. </p> </li> </ul>
	Deletionprotectionenabled interface{} `json:"DeletionProtectionEnabled,omitempty"` // Indicates whether deletion protection is to be enabled (true) or disabled (false) on the table.
	Replicaupdates interface{} `json:"ReplicaUpdates,omitempty"` // <p>A list of replica update actions (create, delete, or update) for the table.</p> <note> <p>This property only applies to <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html">Version 2019.11.21 (Current)</a> of global tables. </p> </note>
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"` // The new provisioned throughput settings for the specified table or index.
	Ssespecification interface{} `json:"SSESpecification,omitempty"` // The new server-side encryption settings for the specified table.
	Streamspecification interface{} `json:"StreamSpecification,omitempty"` // <p>Represents the DynamoDB Streams configuration for the table.</p> <note> <p>You receive a <code>ResourceInUseException</code> if you try to enable a stream on a table that already has a stream, or if you try to disable a stream on a table that doesn't have a stream.</p> </note>
	Globalsecondaryindexupdates interface{} `json:"GlobalSecondaryIndexUpdates,omitempty"` // <p>An array of one or more global secondary indexes for the table. For each index in the array, you can request one action:</p> <ul> <li> <p> <code>Create</code> - add a new global secondary index to the table.</p> </li> <li> <p> <code>Update</code> - modify the provisioned throughput settings of an existing global secondary index.</p> </li> <li> <p> <code>Delete</code> - remove a global secondary index from the table.</p> </li> </ul> <p>You can create or delete only one global secondary index per <code>UpdateTable</code> operation.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.OnlineOps.html">Managing Global Secondary Indexes</a> in the <i>Amazon DynamoDB Developer Guide</i>. </p>
	Tablename interface{} `json:"TableName"` // The name of the table to be updated.
	Tableclass interface{} `json:"TableClass,omitempty"` // The table class of the table to be updated. Valid values are <code>STANDARD</code> and <code>STANDARD_INFREQUENT_ACCESS</code>.
	Attributedefinitions interface{} `json:"AttributeDefinitions,omitempty"` // An array of attributes that describe the key schema for the table and indexes. If you are adding a new global secondary index to the table, <code>AttributeDefinitions</code> must include the key element(s) of the new index.
}

// AutoScalingPolicyUpdate represents the AutoScalingPolicyUpdate schema from the OpenAPI specification
type AutoScalingPolicyUpdate struct {
	Policyname interface{} `json:"PolicyName,omitempty"` // The name of the scaling policy.
	Targettrackingscalingpolicyconfiguration interface{} `json:"TargetTrackingScalingPolicyConfiguration"` // Represents a target tracking scaling policy configuration.
}

// BatchGetItemOutput represents the BatchGetItemOutput schema from the OpenAPI specification
type BatchGetItemOutput struct {
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // <p>The read capacity units consumed by the entire <code>BatchGetItem</code> operation.</p> <p>Each element consists of:</p> <ul> <li> <p> <code>TableName</code> - The table that consumed the provisioned throughput.</p> </li> <li> <p> <code>CapacityUnits</code> - The total number of capacity units consumed.</p> </li> </ul>
	Responses interface{} `json:"Responses,omitempty"` // A map of table name to a list of items. Each object in <code>Responses</code> consists of a table name, along with a map of attribute data consisting of the data type and attribute value.
	Unprocessedkeys interface{} `json:"UnprocessedKeys,omitempty"` // <p>A map of tables and their respective keys that were not processed with the current response. The <code>UnprocessedKeys</code> value is in the same form as <code>RequestItems</code>, so the value can be provided directly to a subsequent <code>BatchGetItem</code> operation. For more information, see <code>RequestItems</code> in the Request Parameters section.</p> <p>Each element consists of:</p> <ul> <li> <p> <code>Keys</code> - An array of primary key attribute values that define specific items in the table.</p> </li> <li> <p> <code>ProjectionExpression</code> - One or more attributes to be retrieved from the table or index. By default, all attributes are returned. If a requested attribute is not found, it does not appear in the result.</p> </li> <li> <p> <code>ConsistentRead</code> - The consistency of a read operation. If set to <code>true</code>, then a strongly consistent read is used; otherwise, an eventually consistent read is used.</p> </li> </ul> <p>If there are no unprocessed keys remaining, the response contains an empty <code>UnprocessedKeys</code> map.</p>
}

// ScanInput represents the ScanInput schema from the OpenAPI specification
type ScanInput struct {
	Indexname interface{} `json:"IndexName,omitempty"` // The name of a secondary index to scan. This index can be any local secondary index or global secondary index. Note that if you use the <code>IndexName</code> parameter, you must also provide <code>TableName</code>.
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // <p>One or more values that can be substituted in an expression.</p> <p>Use the <b>:</b> (colon) character in an expression to dereference an attribute value. For example, suppose that you wanted to check whether the value of the <code>ProductStatus</code> attribute was one of the following: </p> <p> <code>Available | Backordered | Discontinued</code> </p> <p>You would first need to specify <code>ExpressionAttributeValues</code> as follows:</p> <p> <code>{ ":avail":{"S":"Available"}, ":back":{"S":"Backordered"}, ":disc":{"S":"Discontinued"} }</code> </p> <p>You could then use these values in an expression, such as this:</p> <p> <code>ProductStatus IN (:avail, :back, :disc)</code> </p> <p>For more information on expression attribute values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Exclusivestartkey interface{} `json:"ExclusiveStartKey,omitempty"` // <p>The primary key of the first item that this operation will evaluate. Use the value that was returned for <code>LastEvaluatedKey</code> in the previous operation.</p> <p>The data type for <code>ExclusiveStartKey</code> must be String, Number or Binary. No set data types are allowed.</p> <p>In a parallel scan, a <code>Scan</code> request that includes <code>ExclusiveStartKey</code> must specify the same segment whose previous <code>Scan</code> returned the corresponding value of <code>LastEvaluatedKey</code>.</p>
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Segment interface{} `json:"Segment,omitempty"` // <p>For a parallel <code>Scan</code> request, <code>Segment</code> identifies an individual segment to be scanned by an application worker.</p> <p>Segment IDs are zero-based, so the first segment is always 0. For example, if you want to use four application threads to scan a table or an index, then the first thread specifies a <code>Segment</code> value of 0, the second thread specifies 1, and so on.</p> <p>The value of <code>LastEvaluatedKey</code> returned from a parallel <code>Scan</code> request must be used as <code>ExclusiveStartKey</code> with the same segment ID in a subsequent <code>Scan</code> operation.</p> <p>The value for <code>Segment</code> must be greater than or equal to 0, and less than the value provided for <code>TotalSegments</code>.</p> <p>If you provide <code>Segment</code>, you must also provide <code>TotalSegments</code>.</p>
	Attributestoget interface{} `json:"AttributesToGet,omitempty"` // This is a legacy parameter. Use <code>ProjectionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html">AttributesToGet</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Conditionaloperator interface{} `json:"ConditionalOperator,omitempty"` // This is a legacy parameter. Use <code>FilterExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html">ConditionalOperator</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Limit interface{} `json:"Limit,omitempty"` // The maximum number of items to evaluate (not necessarily the number of matching items). If DynamoDB processes the number of items up to the limit while processing the results, it stops the operation and returns the matching values up to that point, and a key in <code>LastEvaluatedKey</code> to apply in a subsequent operation, so that you can pick up where you left off. Also, if the processed dataset size exceeds 1 MB before DynamoDB reaches this limit, it stops the operation and returns the matching values up to the limit, and a key in <code>LastEvaluatedKey</code> to apply in a subsequent operation to continue the operation. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html">Working with Queries</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Totalsegments interface{} `json:"TotalSegments,omitempty"` // <p>For a parallel <code>Scan</code> request, <code>TotalSegments</code> represents the total number of segments into which the <code>Scan</code> operation will be divided. The value of <code>TotalSegments</code> corresponds to the number of application workers that will perform the parallel scan. For example, if you want to use four application threads to scan a table or an index, specify a <code>TotalSegments</code> value of 4.</p> <p>The value for <code>TotalSegments</code> must be greater than or equal to 1, and less than or equal to 1000000. If you specify a <code>TotalSegments</code> value of 1, the <code>Scan</code> operation will be sequential rather than parallel.</p> <p>If you specify <code>TotalSegments</code>, you must also specify <code>Segment</code>.</p>
	Filterexpression interface{} `json:"FilterExpression,omitempty"` // <p>A string that contains conditions that DynamoDB applies after the <code>Scan</code> operation, but before the data is returned to you. Items that do not satisfy the <code>FilterExpression</code> criteria are not returned.</p> <note> <p>A <code>FilterExpression</code> is applied after the items have already been read; the process of filtering does not consume any additional read capacity units.</p> </note> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#Query.FilterExpression">Filter Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Tablename interface{} `json:"TableName"` // The name of the table containing the requested items; or, if you provide <code>IndexName</code>, the name of the table to which that index belongs.
	SelectField interface{} `json:"Select,omitempty"` // <p>The attributes to be returned in the result. You can retrieve all item attributes, specific item attributes, the count of matching items, or in the case of an index, some or all of the attributes projected into the index.</p> <ul> <li> <p> <code>ALL_ATTRIBUTES</code> - Returns all of the item attributes from the specified table or index. If you query a local secondary index, then for each matching item in the index, DynamoDB fetches the entire item from the parent table. If the index is configured to project all item attributes, then all of the data can be obtained from the local secondary index, and no fetching is required.</p> </li> <li> <p> <code>ALL_PROJECTED_ATTRIBUTES</code> - Allowed only when querying an index. Retrieves all attributes that have been projected into the index. If the index is configured to project all attributes, this return value is equivalent to specifying <code>ALL_ATTRIBUTES</code>.</p> </li> <li> <p> <code>COUNT</code> - Returns the number of matching items, rather than the matching items themselves. Note that this uses the same quantity of read capacity units as getting the items, and is subject to the same item size calculations.</p> </li> <li> <p> <code>SPECIFIC_ATTRIBUTES</code> - Returns only the attributes listed in <code>ProjectionExpression</code>. This return value is equivalent to specifying <code>ProjectionExpression</code> without specifying any value for <code>Select</code>.</p> <p>If you query or scan a local secondary index and request only attributes that are projected into that index, the operation reads only the index and not the table. If any of the requested attributes are not projected into the local secondary index, DynamoDB fetches each of these attributes from the parent table. This extra fetching incurs additional throughput cost and latency.</p> <p>If you query or scan a global secondary index, you can only request attributes that are projected into the index. Global secondary index queries cannot fetch attributes from the parent table.</p> </li> </ul> <p>If neither <code>Select</code> nor <code>ProjectionExpression</code> are specified, DynamoDB defaults to <code>ALL_ATTRIBUTES</code> when accessing a table, and <code>ALL_PROJECTED_ATTRIBUTES</code> when accessing an index. You cannot use both <code>Select</code> and <code>ProjectionExpression</code> together in a single request, unless the value for <code>Select</code> is <code>SPECIFIC_ATTRIBUTES</code>. (This usage is equivalent to specifying <code>ProjectionExpression</code> without any value for <code>Select</code>.)</p> <note> <p>If you use the <code>ProjectionExpression</code> parameter, then the value for <code>Select</code> can only be <code>SPECIFIC_ATTRIBUTES</code>. Any other value for <code>Select</code> will return an error.</p> </note>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Scanfilter interface{} `json:"ScanFilter,omitempty"` // This is a legacy parameter. Use <code>FilterExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ScanFilter.html">ScanFilter</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Consistentread interface{} `json:"ConsistentRead,omitempty"` // <p>A Boolean value that determines the read consistency model during the scan:</p> <ul> <li> <p>If <code>ConsistentRead</code> is <code>false</code>, then the data returned from <code>Scan</code> might not contain the results from other recently completed write operations (<code>PutItem</code>, <code>UpdateItem</code>, or <code>DeleteItem</code>).</p> </li> <li> <p>If <code>ConsistentRead</code> is <code>true</code>, then all of the write operations that completed before the <code>Scan</code> began are guaranteed to be contained in the <code>Scan</code> response.</p> </li> </ul> <p>The default setting for <code>ConsistentRead</code> is <code>false</code>.</p> <p>The <code>ConsistentRead</code> parameter is not supported on global secondary indexes. If you scan a global secondary index with <code>ConsistentRead</code> set to true, you will receive a <code>ValidationException</code>.</p>
	Projectionexpression interface{} `json:"ProjectionExpression,omitempty"` // <p>A string that identifies one or more attributes to retrieve from the specified table or index. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.</p> <p>If no attribute names are specified, then all attributes will be returned. If any of the requested attributes are not found, they will not appear in the result.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
}

// KinesisStreamingDestinationInput represents the KinesisStreamingDestinationInput schema from the OpenAPI specification
type KinesisStreamingDestinationInput struct {
	Streamarn interface{} `json:"StreamArn"` // The ARN for a Kinesis data stream.
	Tablename interface{} `json:"TableName"` // The name of the DynamoDB table.
}

// ImportTableOutput represents the ImportTableOutput schema from the OpenAPI specification
type ImportTableOutput struct {
	Importtabledescription interface{} `json:"ImportTableDescription"` // Represents the properties of the table created for the import, and parameters of the import. The import parameters include import status, how many items were processed, and how many errors were encountered.
}

// DescribeTableReplicaAutoScalingOutput represents the DescribeTableReplicaAutoScalingOutput schema from the OpenAPI specification
type DescribeTableReplicaAutoScalingOutput struct {
	Tableautoscalingdescription interface{} `json:"TableAutoScalingDescription,omitempty"` // Represents the auto scaling properties of the table.
}

// Capacity represents the Capacity schema from the OpenAPI specification
type Capacity struct {
	Capacityunits interface{} `json:"CapacityUnits,omitempty"` // The total number of capacity units consumed on a table or an index.
	Readcapacityunits interface{} `json:"ReadCapacityUnits,omitempty"` // The total number of read capacity units consumed on a table or an index.
	Writecapacityunits interface{} `json:"WriteCapacityUnits,omitempty"` // The total number of write capacity units consumed on a table or an index.
}

// KeysAndAttributes represents the KeysAndAttributes schema from the OpenAPI specification
type KeysAndAttributes struct {
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{"#P":"Percentile"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Accessing Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Keys interface{} `json:"Keys"` // The primary key attribute values that define the items and the attributes associated with the items.
	Projectionexpression interface{} `json:"ProjectionExpression,omitempty"` // <p>A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the <code>ProjectionExpression</code> must be separated by commas.</p> <p>If no attribute names are specified, then all attributes will be returned. If any of the requested attributes are not found, they will not appear in the result.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html">Accessing Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Attributestoget interface{} `json:"AttributesToGet,omitempty"` // This is a legacy parameter. Use <code>ProjectionExpression</code> instead. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.html">Legacy Conditional Parameters</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Consistentread interface{} `json:"ConsistentRead,omitempty"` // The consistency of a read operation. If set to <code>true</code>, then a strongly consistent read is used; otherwise, an eventually consistent read is used.
}

// BackupDetails represents the BackupDetails schema from the OpenAPI specification
type BackupDetails struct {
	Backupexpirydatetime interface{} `json:"BackupExpiryDateTime,omitempty"` // Time at which the automatic on-demand backup created by DynamoDB will expire. This <code>SYSTEM</code> on-demand backup expires automatically 35 days after its creation.
	Backupname interface{} `json:"BackupName"` // Name of the requested backup.
	Backupsizebytes interface{} `json:"BackupSizeBytes,omitempty"` // Size of the backup in bytes. DynamoDB updates this value approximately every six hours. Recent changes might not be reflected in this value.
	Backupstatus interface{} `json:"BackupStatus"` // Backup can be in one of the following states: CREATING, ACTIVE, DELETED.
	Backuptype interface{} `json:"BackupType"` // <p>BackupType:</p> <ul> <li> <p> <code>USER</code> - You create and manage these using the on-demand backup feature.</p> </li> <li> <p> <code>SYSTEM</code> - If you delete a table with point-in-time recovery enabled, a <code>SYSTEM</code> backup is automatically created and is retained for 35 days (at no additional cost). System backups allow you to restore the deleted table to the state it was in just before the point of deletion. </p> </li> <li> <p> <code>AWS_BACKUP</code> - On-demand backup created by you from Backup service.</p> </li> </ul>
	Backuparn interface{} `json:"BackupArn"` // ARN associated with the backup.
	Backupcreationdatetime interface{} `json:"BackupCreationDateTime"` // Time at which the backup was created. This is the request time of the backup.
}

// ExportDescription represents the ExportDescription schema from the OpenAPI specification
type ExportDescription struct {
	Exportstatus interface{} `json:"ExportStatus,omitempty"` // Export can be in one of the following states: IN_PROGRESS, COMPLETED, or FAILED.
	Starttime interface{} `json:"StartTime,omitempty"` // The time at which the export task began.
	Exportformat interface{} `json:"ExportFormat,omitempty"` // The format of the exported data. Valid values for <code>ExportFormat</code> are <code>DYNAMODB_JSON</code> or <code>ION</code>.
	S3prefix interface{} `json:"S3Prefix,omitempty"` // The Amazon S3 bucket prefix used as the file name and path of the exported snapshot.
	Exportmanifest interface{} `json:"ExportManifest,omitempty"` // The name of the manifest file for the export task.
	Billedsizebytes interface{} `json:"BilledSizeBytes,omitempty"` // The billable size of the table export.
	Clienttoken interface{} `json:"ClientToken,omitempty"` // The client token that was provided for the export task. A client token makes calls to <code>ExportTableToPointInTimeInput</code> idempotent, meaning that multiple identical calls have the same effect as one single call.
	S3ssekmskeyid interface{} `json:"S3SseKmsKeyId,omitempty"` // The ID of the KMS managed key used to encrypt the S3 bucket where export data is stored (if applicable).
	Failuremessage interface{} `json:"FailureMessage,omitempty"` // Export failure reason description.
	S3bucket interface{} `json:"S3Bucket,omitempty"` // The name of the Amazon S3 bucket containing the export.
	Exportarn interface{} `json:"ExportArn,omitempty"` // The Amazon Resource Name (ARN) of the table export.
	Failurecode interface{} `json:"FailureCode,omitempty"` // Status code for the result of the failed export.
	Exporttime interface{} `json:"ExportTime,omitempty"` // Point in time from which table data was exported.
	Itemcount interface{} `json:"ItemCount,omitempty"` // The number of items exported.
	Tableid interface{} `json:"TableId,omitempty"` // Unique ID of the table that was exported.
	S3ssealgorithm interface{} `json:"S3SseAlgorithm,omitempty"` // <p>Type of encryption used on the bucket where export data is stored. Valid values for <code>S3SseAlgorithm</code> are:</p> <ul> <li> <p> <code>AES256</code> - server-side encryption with Amazon S3 managed keys</p> </li> <li> <p> <code>KMS</code> - server-side encryption with KMS managed keys</p> </li> </ul>
	Tablearn interface{} `json:"TableArn,omitempty"` // The Amazon Resource Name (ARN) of the table that was exported.
	S3bucketowner interface{} `json:"S3BucketOwner,omitempty"` // The ID of the Amazon Web Services account that owns the bucket containing the export.
	Endtime interface{} `json:"EndTime,omitempty"` // The time at which the export task completed.
}

// ReplicaGlobalSecondaryIndexAutoScalingUpdate represents the ReplicaGlobalSecondaryIndexAutoScalingUpdate schema from the OpenAPI specification
type ReplicaGlobalSecondaryIndexAutoScalingUpdate struct {
	Provisionedreadcapacityautoscalingupdate AutoScalingSettingsUpdate `json:"ProvisionedReadCapacityAutoScalingUpdate,omitempty"` // Represents the auto scaling settings to be modified for a global table or global secondary index.
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index.
}

// Replica represents the Replica schema from the OpenAPI specification
type Replica struct {
	Regionname interface{} `json:"RegionName,omitempty"` // The Region where the replica needs to be created.
}

// CreateGlobalTableOutput represents the CreateGlobalTableOutput schema from the OpenAPI specification
type CreateGlobalTableOutput struct {
	Globaltabledescription interface{} `json:"GlobalTableDescription,omitempty"` // Contains the details of the global table.
}

// UpdateContributorInsightsOutput represents the UpdateContributorInsightsOutput schema from the OpenAPI specification
type UpdateContributorInsightsOutput struct {
	Contributorinsightsstatus interface{} `json:"ContributorInsightsStatus,omitempty"` // The status of contributor insights
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index, if applicable.
	Tablename interface{} `json:"TableName,omitempty"` // The name of the table.
}

// ExportTableToPointInTimeInput represents the ExportTableToPointInTimeInput schema from the OpenAPI specification
type ExportTableToPointInTimeInput struct {
	Clienttoken interface{} `json:"ClientToken,omitempty"` // <p>Providing a <code>ClientToken</code> makes the call to <code>ExportTableToPointInTimeInput</code> idempotent, meaning that multiple identical calls have the same effect as one single call.</p> <p>A client token is valid for 8 hours after the first request that uses it is completed. After 8 hours, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 8 hours, or the result might not be idempotent.</p> <p>If you submit a request with the same client token but a change in other parameters within the 8-hour idempotency window, DynamoDB returns an <code>ImportConflictException</code>.</p>
	S3bucketowner interface{} `json:"S3BucketOwner,omitempty"` // The ID of the Amazon Web Services account that owns the bucket the export will be stored in.
	S3ssealgorithm interface{} `json:"S3SseAlgorithm,omitempty"` // <p>Type of encryption used on the bucket where export data will be stored. Valid values for <code>S3SseAlgorithm</code> are:</p> <ul> <li> <p> <code>AES256</code> - server-side encryption with Amazon S3 managed keys</p> </li> <li> <p> <code>KMS</code> - server-side encryption with KMS managed keys</p> </li> </ul>
	S3ssekmskeyid interface{} `json:"S3SseKmsKeyId,omitempty"` // The ID of the KMS managed key used to encrypt the S3 bucket where export data will be stored (if applicable).
	Exportformat interface{} `json:"ExportFormat,omitempty"` // The format for the exported data. Valid values for <code>ExportFormat</code> are <code>DYNAMODB_JSON</code> or <code>ION</code>.
	Exporttime interface{} `json:"ExportTime,omitempty"` // Time in the past from which to export table data, counted in seconds from the start of the Unix epoch. The table export will be a snapshot of the table's state at this point in time.
	S3bucket interface{} `json:"S3Bucket"` // The name of the Amazon S3 bucket to export the snapshot to.
	S3prefix interface{} `json:"S3Prefix,omitempty"` // The Amazon S3 bucket prefix to use as the file name and path of the exported snapshot.
	Tablearn interface{} `json:"TableArn"` // The Amazon Resource Name (ARN) associated with the table to export.
}

// TableCreationParameters represents the TableCreationParameters schema from the OpenAPI specification
type TableCreationParameters struct {
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // The Global Secondary Indexes (GSI) of the table to be created as part of the import operation.
	Keyschema interface{} `json:"KeySchema"` // The primary key and option sort key of the table created as part of the import operation.
	Provisionedthroughput ProvisionedThroughput `json:"ProvisionedThroughput,omitempty"` // <p>Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the <code>UpdateTable</code> operation.</p> <p>For current minimum and maximum provisioned throughput values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Ssespecification SSESpecification `json:"SSESpecification,omitempty"` // Represents the settings used to enable server-side encryption.
	Tablename interface{} `json:"TableName"` // The name of the table created as part of the import operation.
	Attributedefinitions interface{} `json:"AttributeDefinitions"` // The attributes of the table created as part of the import operation.
	Billingmode interface{} `json:"BillingMode,omitempty"` // The billing mode for provisioning the table created as part of the import operation.
}

// GlobalSecondaryIndex represents the GlobalSecondaryIndex schema from the OpenAPI specification
type GlobalSecondaryIndex struct {
	Keyschema interface{} `json:"KeySchema"` // <p>The complete key schema for a global secondary index, which consists of one or more pairs of attribute names and key types:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note>
	Projection interface{} `json:"Projection"` // Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"` // <p>Represents the provisioned throughput settings for the specified global secondary index.</p> <p>For current minimum and maximum provisioned throughput values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index. The name must be unique among all other indexes on this table.
}

// CreateTableInput represents the CreateTableInput schema from the OpenAPI specification
type CreateTableInput struct {
	Deletionprotectionenabled interface{} `json:"DeletionProtectionEnabled,omitempty"` // Indicates whether deletion protection is to be enabled (true) or disabled (false) on the table.
	Keyschema interface{} `json:"KeySchema"` // <p>Specifies the attributes that make up the primary key for a table or an index. The attributes in <code>KeySchema</code> must also be defined in the <code>AttributeDefinitions</code> array. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DataModel.html">Data Model</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> <p>Each <code>KeySchemaElement</code> in the array is composed of:</p> <ul> <li> <p> <code>AttributeName</code> - The name of this key attribute.</p> </li> <li> <p> <code>KeyType</code> - The role that the key attribute will assume:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from the DynamoDB usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note> <p>For a simple primary key (partition key), you must provide exactly one element with a <code>KeyType</code> of <code>HASH</code>.</p> <p>For a composite primary key (partition key and sort key), you must provide exactly two elements, in this order: The first element must have a <code>KeyType</code> of <code>HASH</code>, and the second element must have a <code>KeyType</code> of <code>RANGE</code>.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#WorkingWithTables.primary.key">Working with Tables</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Tableclass interface{} `json:"TableClass,omitempty"` // The table class of the new table. Valid values are <code>STANDARD</code> and <code>STANDARD_INFREQUENT_ACCESS</code>.
	Tablename interface{} `json:"TableName"` // The name of the table to create.
	Billingmode interface{} `json:"BillingMode,omitempty"` // <p>Controls how you are charged for read and write throughput and how you manage capacity. This setting can be changed later.</p> <ul> <li> <p> <code>PROVISIONED</code> - We recommend using <code>PROVISIONED</code> for predictable workloads. <code>PROVISIONED</code> sets the billing mode to <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual">Provisioned Mode</a>.</p> </li> <li> <p> <code>PAY_PER_REQUEST</code> - We recommend using <code>PAY_PER_REQUEST</code> for unpredictable workloads. <code>PAY_PER_REQUEST</code> sets the billing mode to <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand">On-Demand Mode</a>. </p> </li> </ul>
	Attributedefinitions interface{} `json:"AttributeDefinitions"` // An array of attributes that describe the key schema for the table and indexes.
	Globalsecondaryindexes interface{} `json:"GlobalSecondaryIndexes,omitempty"` // <p>One or more global secondary indexes (the maximum is 20) to be created on the table. Each global secondary index in the array includes the following:</p> <ul> <li> <p> <code>IndexName</code> - The name of the global secondary index. Must be unique only for this table.</p> <p/> </li> <li> <p> <code>KeySchema</code> - Specifies the key schema for the global secondary index.</p> </li> <li> <p> <code>Projection</code> - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of:</p> <ul> <li> <p> <code>ProjectionType</code> - One of the following:</p> <ul> <li> <p> <code>KEYS_ONLY</code> - Only the index and primary keys are projected into the index.</p> </li> <li> <p> <code>INCLUDE</code> - Only the specified table attributes are projected into the index. The list of projected attributes is in <code>NonKeyAttributes</code>.</p> </li> <li> <p> <code>ALL</code> - All of the table attributes are projected into the index.</p> </li> </ul> </li> <li> <p> <code>NonKeyAttributes</code> - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in <code>NonKeyAttributes</code>, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.</p> </li> </ul> </li> <li> <p> <code>ProvisionedThroughput</code> - The provisioned throughput settings for the global secondary index, consisting of read and write capacity units.</p> </li> </ul>
	Localsecondaryindexes interface{} `json:"LocalSecondaryIndexes,omitempty"` // <p>One or more local secondary indexes (the maximum is 5) to be created on the table. Each index is scoped to a given partition key value. There is a 10 GB size limit per partition key value; otherwise, the size of a local secondary index is unconstrained.</p> <p>Each local secondary index in the array includes the following:</p> <ul> <li> <p> <code>IndexName</code> - The name of the local secondary index. Must be unique only for this table.</p> <p/> </li> <li> <p> <code>KeySchema</code> - Specifies the key schema for the local secondary index. The key schema must begin with the same partition key as the table.</p> </li> <li> <p> <code>Projection</code> - Specifies attributes that are copied (projected) from the table into the index. These are in addition to the primary key attributes and index key attributes, which are automatically projected. Each attribute specification is composed of:</p> <ul> <li> <p> <code>ProjectionType</code> - One of the following:</p> <ul> <li> <p> <code>KEYS_ONLY</code> - Only the index and primary keys are projected into the index.</p> </li> <li> <p> <code>INCLUDE</code> - Only the specified table attributes are projected into the index. The list of projected attributes is in <code>NonKeyAttributes</code>.</p> </li> <li> <p> <code>ALL</code> - All of the table attributes are projected into the index.</p> </li> </ul> </li> <li> <p> <code>NonKeyAttributes</code> - A list of one or more non-key attribute names that are projected into the secondary index. The total count of attributes provided in <code>NonKeyAttributes</code>, summed across all of the secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.</p> </li> </ul> </li> </ul>
	Ssespecification interface{} `json:"SSESpecification,omitempty"` // Represents the settings used to enable server-side encryption.
	Provisionedthroughput interface{} `json:"ProvisionedThroughput,omitempty"` // <p>Represents the provisioned throughput settings for a specified table or index. The settings can be modified using the <code>UpdateTable</code> operation.</p> <p> If you set BillingMode as <code>PROVISIONED</code>, you must specify this property. If you set BillingMode as <code>PAY_PER_REQUEST</code>, you cannot specify this property.</p> <p>For current minimum and maximum provisioned throughput values, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Limits.html">Service, Account, and Table Quotas</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>
	Streamspecification interface{} `json:"StreamSpecification,omitempty"` // <p>The settings for DynamoDB Streams on the table. These settings consist of:</p> <ul> <li> <p> <code>StreamEnabled</code> - Indicates whether DynamoDB Streams is to be enabled (true) or disabled (false).</p> </li> <li> <p> <code>StreamViewType</code> - When an item in the table is modified, <code>StreamViewType</code> determines what information is written to the table's stream. Valid values for <code>StreamViewType</code> are:</p> <ul> <li> <p> <code>KEYS_ONLY</code> - Only the key attributes of the modified item are written to the stream.</p> </li> <li> <p> <code>NEW_IMAGE</code> - The entire item, as it appears after it was modified, is written to the stream.</p> </li> <li> <p> <code>OLD_IMAGE</code> - The entire item, as it appeared before it was modified, is written to the stream.</p> </li> <li> <p> <code>NEW_AND_OLD_IMAGES</code> - Both the new and the old item images of the item are written to the stream.</p> </li> </ul> </li> </ul>
	Tags interface{} `json:"Tags,omitempty"` // A list of key-value pairs to label the table. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html">Tagging for DynamoDB</a>.
}

// GlobalSecondaryIndexAutoScalingUpdate represents the GlobalSecondaryIndexAutoScalingUpdate schema from the OpenAPI specification
type GlobalSecondaryIndexAutoScalingUpdate struct {
	Provisionedwritecapacityautoscalingupdate AutoScalingSettingsUpdate `json:"ProvisionedWriteCapacityAutoScalingUpdate,omitempty"` // Represents the auto scaling settings to be modified for a global table or global secondary index.
	Indexname interface{} `json:"IndexName,omitempty"` // The name of the global secondary index.
}

// DescribeTableReplicaAutoScalingInput represents the DescribeTableReplicaAutoScalingInput schema from the OpenAPI specification
type DescribeTableReplicaAutoScalingInput struct {
	Tablename interface{} `json:"TableName"` // The name of the table.
}

// ListImportsInput represents the ListImportsInput schema from the OpenAPI specification
type ListImportsInput struct {
	Tablearn interface{} `json:"TableArn,omitempty"` // The Amazon Resource Name (ARN) associated with the table that was imported to.
	Nexttoken interface{} `json:"NextToken,omitempty"` // An optional string that, if supplied, must be copied from the output of a previous call to <code>ListImports</code>. When provided in this manner, the API fetches the next page of results.
	Pagesize interface{} `json:"PageSize,omitempty"` // The number of <code>ImportSummary </code>objects returned in a single page.
}

// AutoScalingTargetTrackingScalingPolicyConfigurationUpdate represents the AutoScalingTargetTrackingScalingPolicyConfigurationUpdate schema from the OpenAPI specification
type AutoScalingTargetTrackingScalingPolicyConfigurationUpdate struct {
	Disablescalein interface{} `json:"DisableScaleIn,omitempty"` // Indicates whether scale in by the target tracking policy is disabled. If the value is true, scale in is disabled and the target tracking policy won't remove capacity from the scalable resource. Otherwise, scale in is enabled and the target tracking policy can remove capacity from the scalable resource. The default value is false.
	Scaleincooldown interface{} `json:"ScaleInCooldown,omitempty"` // The amount of time, in seconds, after a scale in activity completes before another scale in activity can start. The cooldown period is used to block subsequent scale in requests until it has expired. You should scale in conservatively to protect your application's availability. However, if another alarm triggers a scale out policy during the cooldown period after a scale-in, application auto scaling scales out your scalable target immediately.
	Scaleoutcooldown interface{} `json:"ScaleOutCooldown,omitempty"` // The amount of time, in seconds, after a scale out activity completes before another scale out activity can start. While the cooldown period is in effect, the capacity that has been added by the previous scale out event that initiated the cooldown is calculated as part of the desired capacity for the next scale out. You should continuously (but not excessively) scale out.
	Targetvalue interface{} `json:"TargetValue"` // The target value for the metric. The range is 8.515920e-109 to 1.174271e+108 (Base 10) or 2e-360 to 2e360 (Base 2).
}

// GlobalSecondaryIndexUpdate represents the GlobalSecondaryIndexUpdate schema from the OpenAPI specification
type GlobalSecondaryIndexUpdate struct {
	DeleteField interface{} `json:"Delete,omitempty"` // The name of an existing global secondary index to be removed.
	Update interface{} `json:"Update,omitempty"` // The name of an existing global secondary index, along with new provisioned throughput settings to be applied to that index.
	Create interface{} `json:"Create,omitempty"` // <p>The parameters required for creating a global secondary index on an existing table:</p> <ul> <li> <p> <code>IndexName </code> </p> </li> <li> <p> <code>KeySchema </code> </p> </li> <li> <p> <code>AttributeDefinitions </code> </p> </li> <li> <p> <code>Projection </code> </p> </li> <li> <p> <code>ProvisionedThroughput </code> </p> </li> </ul>
}

// BatchExecuteStatementOutput represents the BatchExecuteStatementOutput schema from the OpenAPI specification
type BatchExecuteStatementOutput struct {
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the entire operation. The values of the list are ordered according to the ordering of the statements.
	Responses interface{} `json:"Responses,omitempty"` // The response to each PartiQL statement in the batch. The values of the list are ordered according to the ordering of the request statements.
}

// BackupDescription represents the BackupDescription schema from the OpenAPI specification
type BackupDescription struct {
	Backupdetails interface{} `json:"BackupDetails,omitempty"` // Contains the details of the backup created for the table.
	Sourcetabledetails interface{} `json:"SourceTableDetails,omitempty"` // Contains the details of the table when the backup was created.
	Sourcetablefeaturedetails interface{} `json:"SourceTableFeatureDetails,omitempty"` // Contains the details of the features enabled on the table when the backup was created. For example, LSIs, GSIs, streams, TTL.
}

// ReplicationGroupUpdate represents the ReplicationGroupUpdate schema from the OpenAPI specification
type ReplicationGroupUpdate struct {
	Create interface{} `json:"Create,omitempty"` // The parameters required for creating a replica for the table.
	DeleteField interface{} `json:"Delete,omitempty"` // The parameters required for deleting a replica for the table.
	Update interface{} `json:"Update,omitempty"` // The parameters required for updating a replica for the table.
}

// ItemResponse represents the ItemResponse schema from the OpenAPI specification
type ItemResponse struct {
	Item interface{} `json:"Item,omitempty"` // Map of attribute data consisting of the data type and attribute value.
}

// KinesisDataStreamDestination represents the KinesisDataStreamDestination schema from the OpenAPI specification
type KinesisDataStreamDestination struct {
	Destinationstatusdescription interface{} `json:"DestinationStatusDescription,omitempty"` // The human-readable string that corresponds to the replica status.
	Streamarn interface{} `json:"StreamArn,omitempty"` // The ARN for a specific Kinesis data stream.
	Destinationstatus interface{} `json:"DestinationStatus,omitempty"` // The current status of replication.
}

// SecondaryIndexesCapacityMap represents the SecondaryIndexesCapacityMap schema from the OpenAPI specification
type SecondaryIndexesCapacityMap struct {
}

// PointInTimeRecoverySpecification represents the PointInTimeRecoverySpecification schema from the OpenAPI specification
type PointInTimeRecoverySpecification struct {
	Pointintimerecoveryenabled interface{} `json:"PointInTimeRecoveryEnabled"` // Indicates whether point in time recovery is enabled (true) or disabled (false) on the table.
}

// DeleteItemOutput represents the DeleteItemOutput schema from the OpenAPI specification
type DeleteItemOutput struct {
	Attributes interface{} `json:"Attributes,omitempty"` // A map of attribute names to <code>AttributeValue</code> objects, representing the item as it appeared before the <code>DeleteItem</code> operation. This map appears in the response only if <code>ReturnValues</code> was specified as <code>ALL_OLD</code> in the request.
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the <code>DeleteItem</code> operation. The data returned includes the total provisioned throughput consumed, along with statistics for the table and any indexes involved in the operation. <code>ConsumedCapacity</code> is only returned if the <code>ReturnConsumedCapacity</code> parameter was specified. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ProvisionedThroughputIntro.html">Provisioned Throughput</a> in the <i>Amazon DynamoDB Developer Guide</i>.
	Itemcollectionmetrics interface{} `json:"ItemCollectionMetrics,omitempty"` // <p>Information about item collections, if any, that were affected by the <code>DeleteItem</code> operation. <code>ItemCollectionMetrics</code> is only returned if the <code>ReturnItemCollectionMetrics</code> parameter was specified. If the table does not have any local secondary indexes, this information is not returned in the response.</p> <p>Each <code>ItemCollectionMetrics</code> element consists of:</p> <ul> <li> <p> <code>ItemCollectionKey</code> - The partition key value of the item collection. This is the same as the partition key value of the item itself.</p> </li> <li> <p> <code>SizeEstimateRangeGB</code> - An estimate of item collection size, in gigabytes. This value is a two-element array containing a lower bound and an upper bound for the estimate. The estimate includes the size of all the items in the table, plus the size of all attributes projected into all of the local secondary indexes on that table. Use this estimate to measure whether a local secondary index is approaching its size limit.</p> <p>The estimate is subject to change over time; therefore, do not rely on the precision or accuracy of the estimate.</p> </li> </ul>
}

// ContributorInsightsSummary represents the ContributorInsightsSummary schema from the OpenAPI specification
type ContributorInsightsSummary struct {
	Contributorinsightsstatus interface{} `json:"ContributorInsightsStatus,omitempty"` // Describes the current status for contributor insights for the given table and index, if applicable.
	Indexname interface{} `json:"IndexName,omitempty"` // Name of the index associated with the summary, if any.
	Tablename interface{} `json:"TableName,omitempty"` // Name of the table associated with the summary.
}

// Projection represents the Projection schema from the OpenAPI specification
type Projection struct {
	Projectiontype interface{} `json:"ProjectionType,omitempty"` // <p>The set of attributes that are projected into the index:</p> <ul> <li> <p> <code>KEYS_ONLY</code> - Only the index and primary keys are projected into the index.</p> </li> <li> <p> <code>INCLUDE</code> - In addition to the attributes described in <code>KEYS_ONLY</code>, the secondary index will include other non-key attributes that you specify.</p> </li> <li> <p> <code>ALL</code> - All of the table attributes are projected into the index.</p> </li> </ul>
	Nonkeyattributes interface{} `json:"NonKeyAttributes,omitempty"` // <p>Represents the non-key attribute names which will be projected into the index.</p> <p>For local secondary indexes, the total count of <code>NonKeyAttributes</code> summed across all of the local secondary indexes, must not exceed 100. If you project the same attribute into two different indexes, this counts as two distinct attributes when determining the total.</p>
}

// ListContributorInsightsOutput represents the ListContributorInsightsOutput schema from the OpenAPI specification
type ListContributorInsightsOutput struct {
	Contributorinsightssummaries interface{} `json:"ContributorInsightsSummaries,omitempty"` // A list of ContributorInsightsSummary.
	Nexttoken interface{} `json:"NextToken,omitempty"` // A token to go to the next page if there is one.
}

// BatchGetRequestMap represents the BatchGetRequestMap schema from the OpenAPI specification
type BatchGetRequestMap struct {
}

// TransactWriteItemsOutput represents the TransactWriteItemsOutput schema from the OpenAPI specification
type TransactWriteItemsOutput struct {
	Itemcollectionmetrics interface{} `json:"ItemCollectionMetrics,omitempty"` // A list of tables that were processed by <code>TransactWriteItems</code> and, for each table, information about any item collections that were affected by individual <code>UpdateItem</code>, <code>PutItem</code>, or <code>DeleteItem</code> operations.
	Consumedcapacity interface{} `json:"ConsumedCapacity,omitempty"` // The capacity units consumed by the entire <code>TransactWriteItems</code> operation. The values of the list are ordered according to the ordering of the <code>TransactItems</code> request parameter.
}

// ListGlobalTablesOutput represents the ListGlobalTablesOutput schema from the OpenAPI specification
type ListGlobalTablesOutput struct {
	Globaltables interface{} `json:"GlobalTables,omitempty"` // List of global table names.
	Lastevaluatedglobaltablename interface{} `json:"LastEvaluatedGlobalTableName,omitempty"` // Last evaluated global table name.
}

// CreateGlobalTableInput represents the CreateGlobalTableInput schema from the OpenAPI specification
type CreateGlobalTableInput struct {
	Globaltablename interface{} `json:"GlobalTableName"` // The global table name.
	Replicationgroup interface{} `json:"ReplicationGroup"` // The Regions where the global table needs to be created.
}

// LocalSecondaryIndexInfo represents the LocalSecondaryIndexInfo schema from the OpenAPI specification
type LocalSecondaryIndexInfo struct {
	Keyschema interface{} `json:"KeySchema,omitempty"` // <p>The complete key schema for a local secondary index, which consists of one or more pairs of attribute names and key types:</p> <ul> <li> <p> <code>HASH</code> - partition key</p> </li> <li> <p> <code>RANGE</code> - sort key</p> </li> </ul> <note> <p>The partition key of an item is also known as its <i>hash attribute</i>. The term "hash attribute" derives from DynamoDB's usage of an internal hash function to evenly distribute data items across partitions, based on their partition key values.</p> <p>The sort key of an item is also known as its <i>range attribute</i>. The term "range attribute" derives from the way DynamoDB stores items with the same partition key physically close together, in sorted order by the sort key value.</p> </note>
	Projection interface{} `json:"Projection,omitempty"` // Represents attributes that are copied (projected) from the table into the global secondary index. These are in addition to the primary key attributes and index key attributes, which are automatically projected.
	Indexname interface{} `json:"IndexName,omitempty"` // Represents the name of the local secondary index.
}

// CsvOptions represents the CsvOptions schema from the OpenAPI specification
type CsvOptions struct {
	Headerlist interface{} `json:"HeaderList,omitempty"` // List of the headers used to specify a common header for all source CSV files being imported. If this field is specified then the first line of each CSV file is treated as data instead of the header. If this field is not specified the the first line of each CSV file is treated as the header.
	Delimiter interface{} `json:"Delimiter,omitempty"` // The delimiter used for separating items in the CSV file being imported.
}

// BatchWriteItemInput represents the BatchWriteItemInput schema from the OpenAPI specification
type BatchWriteItemInput struct {
	Requestitems interface{} `json:"RequestItems"` // <p>A map of one or more table names and, for each table, a list of operations to be performed (<code>DeleteRequest</code> or <code>PutRequest</code>). Each element in the map consists of the following:</p> <ul> <li> <p> <code>DeleteRequest</code> - Perform a <code>DeleteItem</code> operation on the specified item. The item to be deleted is identified by a <code>Key</code> subelement:</p> <ul> <li> <p> <code>Key</code> - A map of primary key attribute values that uniquely identify the item. Each entry in this map consists of an attribute name and an attribute value. For each primary key, you must provide <i>all</i> of the key attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for <i>both</i> the partition key and the sort key.</p> </li> </ul> </li> <li> <p> <code>PutRequest</code> - Perform a <code>PutItem</code> operation on the specified item. The item to be put is identified by an <code>Item</code> subelement:</p> <ul> <li> <p> <code>Item</code> - A map of attributes and their values. Each entry in this map consists of an attribute name and an attribute value. Attribute values must not be null; string and binary type attributes must have lengths greater than zero; and set type attributes must not be empty. Requests that contain empty values are rejected with a <code>ValidationException</code> exception.</p> <p>If you specify any attributes that are part of an index key, then the data types for those attributes must match those of the schema in the table's attribute definition.</p> </li> </ul> </li> </ul>
	Returnconsumedcapacity string `json:"ReturnConsumedCapacity,omitempty"` // <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>
	Returnitemcollectionmetrics interface{} `json:"ReturnItemCollectionMetrics,omitempty"` // Determines whether item collection metrics are returned. If set to <code>SIZE</code>, the response includes statistics about item collections, if any, that were modified during the operation are returned in the response. If set to <code>NONE</code> (the default), no statistics are returned.
}

// UpdateContinuousBackupsOutput represents the UpdateContinuousBackupsOutput schema from the OpenAPI specification
type UpdateContinuousBackupsOutput struct {
	Continuousbackupsdescription interface{} `json:"ContinuousBackupsDescription,omitempty"` // Represents the continuous backups and point in time recovery settings on the table.
}

// UntagResourceInput represents the UntagResourceInput schema from the OpenAPI specification
type UntagResourceInput struct {
	Resourcearn interface{} `json:"ResourceArn"` // The DynamoDB resource that the tags will be removed from. This value is an Amazon Resource Name (ARN).
	Tagkeys interface{} `json:"TagKeys"` // A list of tag keys. Existing tags of the resource whose keys are members of this list will be removed from the DynamoDB resource.
}

// ContinuousBackupsDescription represents the ContinuousBackupsDescription schema from the OpenAPI specification
type ContinuousBackupsDescription struct {
	Continuousbackupsstatus interface{} `json:"ContinuousBackupsStatus"` // <code>ContinuousBackupsStatus</code> can be one of the following states: ENABLED, DISABLED
	Pointintimerecoverydescription interface{} `json:"PointInTimeRecoveryDescription,omitempty"` // The description of the point in time recovery settings applied to the table.
}

// Get represents the Get schema from the OpenAPI specification
type Get struct {
	Projectionexpression interface{} `json:"ProjectionExpression,omitempty"` // A string that identifies one or more attributes of the specified item to retrieve from the table. The attributes in the expression must be separated by commas. If no attribute names are specified, then all attributes of the specified item are returned. If any of the requested attributes are not found, they do not appear in the result.
	Tablename interface{} `json:"TableName"` // The name of the table from which to retrieve the specified item.
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // One or more substitution tokens for attribute names in the ProjectionExpression parameter.
	Key interface{} `json:"Key"` // A map of attribute names to <code>AttributeValue</code> objects that specifies the primary key of the item to retrieve.
}

// RestoreTableFromBackupInput represents the RestoreTableFromBackupInput schema from the OpenAPI specification
type RestoreTableFromBackupInput struct {
	Globalsecondaryindexoverride interface{} `json:"GlobalSecondaryIndexOverride,omitempty"` // List of global secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.
	Localsecondaryindexoverride interface{} `json:"LocalSecondaryIndexOverride,omitempty"` // List of local secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"` // Provisioned throughput settings for the restored table.
	Ssespecificationoverride interface{} `json:"SSESpecificationOverride,omitempty"` // The new server-side encryption settings for the restored table.
	Targettablename interface{} `json:"TargetTableName"` // The name of the new table to which the backup must be restored.
	Backuparn interface{} `json:"BackupArn"` // The Amazon Resource Name (ARN) associated with the backup.
	Billingmodeoverride interface{} `json:"BillingModeOverride,omitempty"` // The billing mode of the restored table.
}

// ReplicaGlobalSecondaryIndex represents the ReplicaGlobalSecondaryIndex schema from the OpenAPI specification
type ReplicaGlobalSecondaryIndex struct {
	Indexname interface{} `json:"IndexName"` // The name of the global secondary index.
	Provisionedthroughputoverride interface{} `json:"ProvisionedThroughputOverride,omitempty"` // Replica table GSI-specific provisioned throughput. If not specified, uses the source table GSI's read capacity settings.
}

// Update represents the Update schema from the OpenAPI specification
type Update struct {
	Key interface{} `json:"Key"` // The primary key of the item to be updated. Each element consists of an attribute name and a value for that attribute.
	Returnvaluesonconditioncheckfailure interface{} `json:"ReturnValuesOnConditionCheckFailure,omitempty"` // Use <code>ReturnValuesOnConditionCheckFailure</code> to get the item attributes if the <code>Update</code> condition fails. For <code>ReturnValuesOnConditionCheckFailure</code>, the valid values are: NONE and ALL_OLD.
	Tablename interface{} `json:"TableName"` // Name of the table for the <code>UpdateItem</code> request.
	Updateexpression interface{} `json:"UpdateExpression"` // An expression that defines one or more attributes to be updated, the action to be performed on them, and new value(s) for them.
	Conditionexpression interface{} `json:"ConditionExpression,omitempty"` // A condition that must be satisfied in order for a conditional update to succeed.
	Expressionattributenames interface{} `json:"ExpressionAttributeNames,omitempty"` // One or more substitution tokens for attribute names in an expression.
	Expressionattributevalues interface{} `json:"ExpressionAttributeValues,omitempty"` // One or more values that can be substituted in an expression.
}
