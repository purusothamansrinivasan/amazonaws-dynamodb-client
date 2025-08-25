/**
 * MCP Server function for <p>Deletes a single item in a table by primary key. You can perform a conditional delete operation that deletes the item if it exists, or if it has an expected attribute value.</p> <p>In addition to deleting an item, you can also return the item's attribute values in the same operation, using the <code>ReturnValues</code> parameter.</p> <p>Unless you specify conditions, the <code>DeleteItem</code> is an idempotent operation; running it multiple times on the same item or attribute does <i>not</i> result in an error response.</p> <p>Conditional deletes are useful for deleting items only if specific conditions are met. If those conditions are met, DynamoDB performs the delete. Otherwise, the item is not deleted.</p>
 */

import java.io.IOException;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.util.*;
import java.util.function.Function;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;

class Post_X_Amz_Target_Dynamo_Db_20120810_Delete_ItemMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_Delete_ItemHandler(MCPServer.APIConfig config) {
        return (request) -> {
            try {
                Map<String, Object> args = request.getArguments();
                if (args == null) {
                    return new MCPServer.MCPToolResult("Invalid arguments object", true);
                }
                
                List<String> queryParams = new ArrayList<>();
        if (args.containsKey("X-Amz-Content-Sha256")) {
            queryParams.add("X-Amz-Content-Sha256=" + args.get("X-Amz-Content-Sha256"));
        }
        if (args.containsKey("X-Amz-Date")) {
            queryParams.add("X-Amz-Date=" + args.get("X-Amz-Date"));
        }
        if (args.containsKey("X-Amz-Algorithm")) {
            queryParams.add("X-Amz-Algorithm=" + args.get("X-Amz-Algorithm"));
        }
        if (args.containsKey("X-Amz-Credential")) {
            queryParams.add("X-Amz-Credential=" + args.get("X-Amz-Credential"));
        }
        if (args.containsKey("X-Amz-Security-Token")) {
            queryParams.add("X-Amz-Security-Token=" + args.get("X-Amz-Security-Token"));
        }
        if (args.containsKey("X-Amz-Signature")) {
            queryParams.add("X-Amz-Signature=" + args.get("X-Amz-Signature"));
        }
        if (args.containsKey("X-Amz-SignedHeaders")) {
            queryParams.add("X-Amz-SignedHeaders=" + args.get("X-Amz-SignedHeaders"));
        }
        if (args.containsKey("X-Amz-Target")) {
            queryParams.add("X-Amz-Target=" + args.get("X-Amz-Target"));
        }
        if (args.containsKey("Expected")) {
            queryParams.add("Expected=" + args.get("Expected"));
        }
        if (args.containsKey("ExpressionAttributeValues")) {
            queryParams.add("ExpressionAttributeValues=" + args.get("ExpressionAttributeValues"));
        }
        if (args.containsKey("ReturnConsumedCapacity")) {
            queryParams.add("ReturnConsumedCapacity=" + args.get("ReturnConsumedCapacity"));
        }
        if (args.containsKey("ReturnItemCollectionMetrics")) {
            queryParams.add("ReturnItemCollectionMetrics=" + args.get("ReturnItemCollectionMetrics"));
        }
        if (args.containsKey("ConditionExpression")) {
            queryParams.add("ConditionExpression=" + args.get("ConditionExpression"));
        }
        if (args.containsKey("ExpressionAttributeNames")) {
            queryParams.add("ExpressionAttributeNames=" + args.get("ExpressionAttributeNames"));
        }
        if (args.containsKey("TableName")) {
            queryParams.add("TableName=" + args.get("TableName"));
        }
        if (args.containsKey("ConditionalOperator")) {
            queryParams.add("ConditionalOperator=" + args.get("ConditionalOperator"));
        }
        if (args.containsKey("Key")) {
            queryParams.add("Key=" + args.get("Key"));
        }
        if (args.containsKey("ReturnValues")) {
            queryParams.add("ReturnValues=" + args.get("ReturnValues"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_delete_item" + queryString;
                
                HttpClient client = HttpClient.newHttpClient();
                HttpRequest httpRequest = HttpRequest.newBuilder()
                    .uri(URI.create(url))
                    .header("Authorization", "Bearer " + config.getApiKey())
                    .header("Accept", "application/json")
                    .GET()
                    .build();
                
                HttpResponse<String> response = client.send(httpRequest, HttpResponse.BodyHandlers.ofString());
                
                if (response.statusCode() >= 400) {
                    return new MCPServer.MCPToolResult("API error: " + response.body(), true);
                }
                
                // Pretty print JSON
                ObjectMapper mapper = new ObjectMapper();
                JsonNode jsonNode = mapper.readTree(response.body());
                String prettyJson = mapper.writerWithDefaultPrettyPrinter().writeValueAsString(jsonNode);
                
                return new MCPServer.MCPToolResult(prettyJson);
                
            } catch (IOException | InterruptedException e) {
                return new MCPServer.MCPToolResult("Request failed: " + e.getMessage(), true);
            } catch (Exception e) {
                return new MCPServer.MCPToolResult("Unexpected error: " + e.getMessage(), true);
            }
        };
    }
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_Delete_ItemTool(MCPServer.APIConfig config) {
        Map<String, Object> parameters = new HashMap<>();
        parameters.put("type", "object");
        Map<String, Object> properties = new HashMap<>();
        Map<String, Object> X-Amz-Content-Sha256Property = new HashMap<>();
        X-Amz-Content-Sha256Property.put("type", "string");
        X-Amz-Content-Sha256Property.put("required", false);
        X-Amz-Content-Sha256Property.put("description", "");
        properties.put("X-Amz-Content-Sha256", X-Amz-Content-Sha256Property);
        Map<String, Object> X-Amz-DateProperty = new HashMap<>();
        X-Amz-DateProperty.put("type", "string");
        X-Amz-DateProperty.put("required", false);
        X-Amz-DateProperty.put("description", "");
        properties.put("X-Amz-Date", X-Amz-DateProperty);
        Map<String, Object> X-Amz-AlgorithmProperty = new HashMap<>();
        X-Amz-AlgorithmProperty.put("type", "string");
        X-Amz-AlgorithmProperty.put("required", false);
        X-Amz-AlgorithmProperty.put("description", "");
        properties.put("X-Amz-Algorithm", X-Amz-AlgorithmProperty);
        Map<String, Object> X-Amz-CredentialProperty = new HashMap<>();
        X-Amz-CredentialProperty.put("type", "string");
        X-Amz-CredentialProperty.put("required", false);
        X-Amz-CredentialProperty.put("description", "");
        properties.put("X-Amz-Credential", X-Amz-CredentialProperty);
        Map<String, Object> X-Amz-Security-TokenProperty = new HashMap<>();
        X-Amz-Security-TokenProperty.put("type", "string");
        X-Amz-Security-TokenProperty.put("required", false);
        X-Amz-Security-TokenProperty.put("description", "");
        properties.put("X-Amz-Security-Token", X-Amz-Security-TokenProperty);
        Map<String, Object> X-Amz-SignatureProperty = new HashMap<>();
        X-Amz-SignatureProperty.put("type", "string");
        X-Amz-SignatureProperty.put("required", false);
        X-Amz-SignatureProperty.put("description", "");
        properties.put("X-Amz-Signature", X-Amz-SignatureProperty);
        Map<String, Object> X-Amz-SignedHeadersProperty = new HashMap<>();
        X-Amz-SignedHeadersProperty.put("type", "string");
        X-Amz-SignedHeadersProperty.put("required", false);
        X-Amz-SignedHeadersProperty.put("description", "");
        properties.put("X-Amz-SignedHeaders", X-Amz-SignedHeadersProperty);
        Map<String, Object> X-Amz-TargetProperty = new HashMap<>();
        X-Amz-TargetProperty.put("type", "string");
        X-Amz-TargetProperty.put("required", true);
        X-Amz-TargetProperty.put("description", "");
        properties.put("X-Amz-Target", X-Amz-TargetProperty);
        Map<String, Object> ExpectedProperty = new HashMap<>();
        ExpectedProperty.put("type", "string");
        ExpectedProperty.put("required", false);
        ExpectedProperty.put("description", "Input parameter: This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.Expected.html\">Expected</a> in the <i>Amazon DynamoDB Developer Guide</i>.");
        properties.put("Expected", ExpectedProperty);
        Map<String, Object> ExpressionAttributeValuesProperty = new HashMap<>();
        ExpressionAttributeValuesProperty.put("type", "string");
        ExpressionAttributeValuesProperty.put("required", false);
        ExpressionAttributeValuesProperty.put("description", "Input parameter: <p>One or more values that can be substituted in an expression.</p> <p>Use the <b>:</b> (colon) character in an expression to dereference an attribute value. For example, suppose that you wanted to check whether the value of the <i>ProductStatus</i> attribute was one of the following: </p> <p> <code>Available | Backordered | Discontinued</code> </p> <p>You would first need to specify <code>ExpressionAttributeValues</code> as follows:</p> <p> <code>{ \":avail\":{\"S\":\"Available\"}, \":back\":{\"S\":\"Backordered\"}, \":disc\":{\"S\":\"Discontinued\"} }</code> </p> <p>You could then use these values in an expression, such as this:</p> <p> <code>ProductStatus IN (:avail, :back, :disc)</code> </p> <p>For more information on expression attribute values, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html\">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>");
        properties.put("ExpressionAttributeValues", ExpressionAttributeValuesProperty);
        Map<String, Object> ReturnConsumedCapacityProperty = new HashMap<>();
        ReturnConsumedCapacityProperty.put("type", "string");
        ReturnConsumedCapacityProperty.put("required", false);
        ReturnConsumedCapacityProperty.put("description", "Input parameter: <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>");
        properties.put("ReturnConsumedCapacity", ReturnConsumedCapacityProperty);
        Map<String, Object> ReturnItemCollectionMetricsProperty = new HashMap<>();
        ReturnItemCollectionMetricsProperty.put("type", "string");
        ReturnItemCollectionMetricsProperty.put("required", false);
        ReturnItemCollectionMetricsProperty.put("description", "Input parameter: Determines whether item collection metrics are returned. If set to <code>SIZE</code>, the response includes statistics about item collections, if any, that were modified during the operation are returned in the response. If set to <code>NONE</code> (the default), no statistics are returned.");
        properties.put("ReturnItemCollectionMetrics", ReturnItemCollectionMetricsProperty);
        Map<String, Object> ConditionExpressionProperty = new HashMap<>();
        ConditionExpressionProperty.put("type", "string");
        ConditionExpressionProperty.put("required", false);
        ConditionExpressionProperty.put("description", "Input parameter: <p>A condition that must be satisfied in order for a conditional <code>DeleteItem</code> to succeed.</p> <p>An expression can contain any of the following:</p> <ul> <li> <p>Functions: <code>attribute_exists | attribute_not_exists | attribute_type | contains | begins_with | size</code> </p> <p>These function names are case-sensitive.</p> </li> <li> <p>Comparison operators: <code>= | &lt;&gt; | &lt; | &gt; | &lt;= | &gt;= | BETWEEN | IN </code> </p> </li> <li> <p> Logical operators: <code>AND | OR | NOT</code> </p> </li> </ul> <p>For more information about condition expressions, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html\">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>");
        properties.put("ConditionExpression", ConditionExpressionProperty);
        Map<String, Object> ExpressionAttributeNamesProperty = new HashMap<>();
        ExpressionAttributeNamesProperty.put("type", "string");
        ExpressionAttributeNamesProperty.put("required", false);
        ExpressionAttributeNamesProperty.put("description", "Input parameter: <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html\">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{\"#P\":\"Percentile\"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html\">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>");
        properties.put("ExpressionAttributeNames", ExpressionAttributeNamesProperty);
        Map<String, Object> TableNameProperty = new HashMap<>();
        TableNameProperty.put("type", "string");
        TableNameProperty.put("required", true);
        TableNameProperty.put("description", "Input parameter: The name of the table from which to delete the item.");
        properties.put("TableName", TableNameProperty);
        Map<String, Object> ConditionalOperatorProperty = new HashMap<>();
        ConditionalOperatorProperty.put("type", "string");
        ConditionalOperatorProperty.put("required", false);
        ConditionalOperatorProperty.put("description", "Input parameter: This is a legacy parameter. Use <code>ConditionExpression</code> instead. For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html\">ConditionalOperator</a> in the <i>Amazon DynamoDB Developer Guide</i>.");
        properties.put("ConditionalOperator", ConditionalOperatorProperty);
        Map<String, Object> KeyProperty = new HashMap<>();
        KeyProperty.put("type", "string");
        KeyProperty.put("required", true);
        KeyProperty.put("description", "Input parameter: <p>A map of attribute names to <code>AttributeValue</code> objects, representing the primary key of the item to delete.</p> <p>For the primary key, you must provide all of the key attributes. For example, with a simple primary key, you only need to provide a value for the partition key. For a composite primary key, you must provide values for both the partition key and the sort key.</p>");
        properties.put("Key", KeyProperty);
        Map<String, Object> ReturnValuesProperty = new HashMap<>();
        ReturnValuesProperty.put("type", "string");
        ReturnValuesProperty.put("required", false);
        ReturnValuesProperty.put("description", "Input parameter: <p>Use <code>ReturnValues</code> if you want to get the item attributes as they appeared before they were deleted. For <code>DeleteItem</code>, the valid values are:</p> <ul> <li> <p> <code>NONE</code> - If <code>ReturnValues</code> is not specified, or if its value is <code>NONE</code>, then nothing is returned. (This setting is the default for <code>ReturnValues</code>.)</p> </li> <li> <p> <code>ALL_OLD</code> - The content of the old item is returned.</p> </li> </ul> <p>There is no additional cost associated with requesting a return value aside from the small network and processing overhead of receiving a larger response. No read capacity units are consumed.</p> <note> <p>The <code>ReturnValues</code> parameter is used by several DynamoDB operations; however, <code>DeleteItem</code> does not recognize any values other than <code>NONE</code> or <code>ALL_OLD</code>.</p> </note>");
        properties.put("ReturnValues", ReturnValuesProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_delete_item",
            "<p>Deletes a single item in a table by primary key. You can perform a conditional delete operation that deletes the item if it exists, or if it has an expected attribute value.</p> <p>In addition to deleting an item, you can also return the item's attribute values in the same operation, using the <code>ReturnValues</code> parameter.</p> <p>Unless you specify conditions, the <code>DeleteItem</code> is an idempotent operation; running it multiple times on the same item or attribute does <i>not</i> result in an error response.</p> <p>Conditional deletes are useful for deleting items only if specific conditions are met. If those conditions are met, DynamoDB performs the delete. Otherwise, the item is not deleted.</p>",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_Delete_ItemHandler(config));
    }
    
}