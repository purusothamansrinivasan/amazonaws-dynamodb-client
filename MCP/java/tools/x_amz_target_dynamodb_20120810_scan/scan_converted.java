/**
 * MCP Server function for No description available
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

class Post_X_Amz_Target_Dynamo_Db_20120810_ScanMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_ScanHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("Limit")) {
            queryParams.add("Limit=" + args.get("Limit"));
        }
        if (args.containsKey("ExclusiveStartKey")) {
            queryParams.add("ExclusiveStartKey=" + args.get("ExclusiveStartKey"));
        }
        if (args.containsKey("X-Amz-Target")) {
            queryParams.add("X-Amz-Target=" + args.get("X-Amz-Target"));
        }
        if (args.containsKey("ConditionalOperator")) {
            queryParams.add("ConditionalOperator=" + args.get("ConditionalOperator"));
        }
        if (args.containsKey("TotalSegments")) {
            queryParams.add("TotalSegments=" + args.get("TotalSegments"));
        }
        if (args.containsKey("FilterExpression")) {
            queryParams.add("FilterExpression=" + args.get("FilterExpression"));
        }
        if (args.containsKey("TableName")) {
            queryParams.add("TableName=" + args.get("TableName"));
        }
        if (args.containsKey("ReturnConsumedCapacity")) {
            queryParams.add("ReturnConsumedCapacity=" + args.get("ReturnConsumedCapacity"));
        }
        if (args.containsKey("ScanFilter")) {
            queryParams.add("ScanFilter=" + args.get("ScanFilter"));
        }
        if (args.containsKey("ConsistentRead")) {
            queryParams.add("ConsistentRead=" + args.get("ConsistentRead"));
        }
        if (args.containsKey("ProjectionExpression")) {
            queryParams.add("ProjectionExpression=" + args.get("ProjectionExpression"));
        }
        if (args.containsKey("IndexName")) {
            queryParams.add("IndexName=" + args.get("IndexName"));
        }
        if (args.containsKey("ExpressionAttributeValues")) {
            queryParams.add("ExpressionAttributeValues=" + args.get("ExpressionAttributeValues"));
        }
        if (args.containsKey("ExpressionAttributeNames")) {
            queryParams.add("ExpressionAttributeNames=" + args.get("ExpressionAttributeNames"));
        }
        if (args.containsKey("Segment")) {
            queryParams.add("Segment=" + args.get("Segment"));
        }
        if (args.containsKey("AttributesToGet")) {
            queryParams.add("AttributesToGet=" + args.get("AttributesToGet"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_scan" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_ScanTool(MCPServer.APIConfig config) {
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
        Map<String, Object> LimitProperty = new HashMap<>();
        LimitProperty.put("type", "string");
        LimitProperty.put("required", false);
        LimitProperty.put("description", "Pagination limit");
        properties.put("Limit", LimitProperty);
        Map<String, Object> ExclusiveStartKeyProperty = new HashMap<>();
        ExclusiveStartKeyProperty.put("type", "string");
        ExclusiveStartKeyProperty.put("required", false);
        ExclusiveStartKeyProperty.put("description", "Pagination token");
        properties.put("ExclusiveStartKey", ExclusiveStartKeyProperty);
        Map<String, Object> X-Amz-TargetProperty = new HashMap<>();
        X-Amz-TargetProperty.put("type", "string");
        X-Amz-TargetProperty.put("required", true);
        X-Amz-TargetProperty.put("description", "");
        properties.put("X-Amz-Target", X-Amz-TargetProperty);
        Map<String, Object> ConditionalOperatorProperty = new HashMap<>();
        ConditionalOperatorProperty.put("type", "string");
        ConditionalOperatorProperty.put("required", false);
        ConditionalOperatorProperty.put("description", "Input parameter: This is a legacy parameter. Use <code>FilterExpression</code> instead. For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ConditionalOperator.html\">ConditionalOperator</a> in the <i>Amazon DynamoDB Developer Guide</i>.");
        properties.put("ConditionalOperator", ConditionalOperatorProperty);
        Map<String, Object> TotalSegmentsProperty = new HashMap<>();
        TotalSegmentsProperty.put("type", "string");
        TotalSegmentsProperty.put("required", false);
        TotalSegmentsProperty.put("description", "Input parameter: <p>For a parallel <code>Scan</code> request, <code>TotalSegments</code> represents the total number of segments into which the <code>Scan</code> operation will be divided. The value of <code>TotalSegments</code> corresponds to the number of application workers that will perform the parallel scan. For example, if you want to use four application threads to scan a table or an index, specify a <code>TotalSegments</code> value of 4.</p> <p>The value for <code>TotalSegments</code> must be greater than or equal to 1, and less than or equal to 1000000. If you specify a <code>TotalSegments</code> value of 1, the <code>Scan</code> operation will be sequential rather than parallel.</p> <p>If you specify <code>TotalSegments</code>, you must also specify <code>Segment</code>.</p>");
        properties.put("TotalSegments", TotalSegmentsProperty);
        Map<String, Object> FilterExpressionProperty = new HashMap<>();
        FilterExpressionProperty.put("type", "string");
        FilterExpressionProperty.put("required", false);
        FilterExpressionProperty.put("description", "Input parameter: <p>A string that contains conditions that DynamoDB applies after the <code>Scan</code> operation, but before the data is returned to you. Items that do not satisfy the <code>FilterExpression</code> criteria are not returned.</p> <note> <p>A <code>FilterExpression</code> is applied after the items have already been read; the process of filtering does not consume any additional read capacity units.</p> </note> <p>For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/QueryAndScan.html#Query.FilterExpression\">Filter Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>");
        properties.put("FilterExpression", FilterExpressionProperty);
        Map<String, Object> TableNameProperty = new HashMap<>();
        TableNameProperty.put("type", "string");
        TableNameProperty.put("required", true);
        TableNameProperty.put("description", "Input parameter: The name of the table containing the requested items; or, if you provide <code>IndexName</code>, the name of the table to which that index belongs.");
        properties.put("TableName", TableNameProperty);
        Map<String, Object> ReturnConsumedCapacityProperty = new HashMap<>();
        ReturnConsumedCapacityProperty.put("type", "string");
        ReturnConsumedCapacityProperty.put("required", false);
        ReturnConsumedCapacityProperty.put("description", "Input parameter: <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>");
        properties.put("ReturnConsumedCapacity", ReturnConsumedCapacityProperty);
        Map<String, Object> ScanFilterProperty = new HashMap<>();
        ScanFilterProperty.put("type", "string");
        ScanFilterProperty.put("required", false);
        ScanFilterProperty.put("description", "Input parameter: This is a legacy parameter. Use <code>FilterExpression</code> instead. For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.ScanFilter.html\">ScanFilter</a> in the <i>Amazon DynamoDB Developer Guide</i>.");
        properties.put("ScanFilter", ScanFilterProperty);
        Map<String, Object> ConsistentReadProperty = new HashMap<>();
        ConsistentReadProperty.put("type", "string");
        ConsistentReadProperty.put("required", false);
        ConsistentReadProperty.put("description", "Input parameter: <p>A Boolean value that determines the read consistency model during the scan:</p> <ul> <li> <p>If <code>ConsistentRead</code> is <code>false</code>, then the data returned from <code>Scan</code> might not contain the results from other recently completed write operations (<code>PutItem</code>, <code>UpdateItem</code>, or <code>DeleteItem</code>).</p> </li> <li> <p>If <code>ConsistentRead</code> is <code>true</code>, then all of the write operations that completed before the <code>Scan</code> began are guaranteed to be contained in the <code>Scan</code> response.</p> </li> </ul> <p>The default setting for <code>ConsistentRead</code> is <code>false</code>.</p> <p>The <code>ConsistentRead</code> parameter is not supported on global secondary indexes. If you scan a global secondary index with <code>ConsistentRead</code> set to true, you will receive a <code>ValidationException</code>.</p>");
        properties.put("ConsistentRead", ConsistentReadProperty);
        Map<String, Object> ProjectionExpressionProperty = new HashMap<>();
        ProjectionExpressionProperty.put("type", "string");
        ProjectionExpressionProperty.put("required", false);
        ProjectionExpressionProperty.put("description", "Input parameter: <p>A string that identifies one or more attributes to retrieve from the specified table or index. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.</p> <p>If no attribute names are specified, then all attributes will be returned. If any of the requested attributes are not found, they will not appear in the result.</p> <p>For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html\">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>");
        properties.put("ProjectionExpression", ProjectionExpressionProperty);
        Map<String, Object> IndexNameProperty = new HashMap<>();
        IndexNameProperty.put("type", "string");
        IndexNameProperty.put("required", false);
        IndexNameProperty.put("description", "Input parameter: The name of a secondary index to scan. This index can be any local secondary index or global secondary index. Note that if you use the <code>IndexName</code> parameter, you must also provide <code>TableName</code>.");
        properties.put("IndexName", IndexNameProperty);
        Map<String, Object> ExpressionAttributeValuesProperty = new HashMap<>();
        ExpressionAttributeValuesProperty.put("type", "string");
        ExpressionAttributeValuesProperty.put("required", false);
        ExpressionAttributeValuesProperty.put("description", "Input parameter: <p>One or more values that can be substituted in an expression.</p> <p>Use the <b>:</b> (colon) character in an expression to dereference an attribute value. For example, suppose that you wanted to check whether the value of the <code>ProductStatus</code> attribute was one of the following: </p> <p> <code>Available | Backordered | Discontinued</code> </p> <p>You would first need to specify <code>ExpressionAttributeValues</code> as follows:</p> <p> <code>{ \":avail\":{\"S\":\"Available\"}, \":back\":{\"S\":\"Backordered\"}, \":disc\":{\"S\":\"Discontinued\"} }</code> </p> <p>You could then use these values in an expression, such as this:</p> <p> <code>ProductStatus IN (:avail, :back, :disc)</code> </p> <p>For more information on expression attribute values, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.SpecifyingConditions.html\">Condition Expressions</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>");
        properties.put("ExpressionAttributeValues", ExpressionAttributeValuesProperty);
        Map<String, Object> ExpressionAttributeNamesProperty = new HashMap<>();
        ExpressionAttributeNamesProperty.put("type", "string");
        ExpressionAttributeNamesProperty.put("required", false);
        ExpressionAttributeNamesProperty.put("description", "Input parameter: <p>One or more substitution tokens for attribute names in an expression. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html\">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{\"#P\":\"Percentile\"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information on expression attribute names, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html\">Specifying Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>");
        properties.put("ExpressionAttributeNames", ExpressionAttributeNamesProperty);
        Map<String, Object> SegmentProperty = new HashMap<>();
        SegmentProperty.put("type", "string");
        SegmentProperty.put("required", false);
        SegmentProperty.put("description", "Input parameter: <p>For a parallel <code>Scan</code> request, <code>Segment</code> identifies an individual segment to be scanned by an application worker.</p> <p>Segment IDs are zero-based, so the first segment is always 0. For example, if you want to use four application threads to scan a table or an index, then the first thread specifies a <code>Segment</code> value of 0, the second thread specifies 1, and so on.</p> <p>The value of <code>LastEvaluatedKey</code> returned from a parallel <code>Scan</code> request must be used as <code>ExclusiveStartKey</code> with the same segment ID in a subsequent <code>Scan</code> operation.</p> <p>The value for <code>Segment</code> must be greater than or equal to 0, and less than the value provided for <code>TotalSegments</code>.</p> <p>If you provide <code>Segment</code>, you must also provide <code>TotalSegments</code>.</p>");
        properties.put("Segment", SegmentProperty);
        Map<String, Object> AttributesToGetProperty = new HashMap<>();
        AttributesToGetProperty.put("type", "string");
        AttributesToGetProperty.put("required", false);
        AttributesToGetProperty.put("description", "Input parameter: This is a legacy parameter. Use <code>ProjectionExpression</code> instead. For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html\">AttributesToGet</a> in the <i>Amazon DynamoDB Developer Guide</i>.");
        properties.put("AttributesToGet", AttributesToGetProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_scan",
            "No description available",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_ScanHandler(config));
    }
    
}