/**
 * MCP Server function for Exports table data to an S3 bucket. The table must have point in time recovery enabled, and you can export data from any time within the point in time recovery window.
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

class Post_X_Amz_Target_Dynamo_Db_20120810_Export_Table_To_Point_In_TimeMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_Export_Table_To_Point_In_TimeHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("ExportTime")) {
            queryParams.add("ExportTime=" + args.get("ExportTime"));
        }
        if (args.containsKey("S3Bucket")) {
            queryParams.add("S3Bucket=" + args.get("S3Bucket"));
        }
        if (args.containsKey("S3Prefix")) {
            queryParams.add("S3Prefix=" + args.get("S3Prefix"));
        }
        if (args.containsKey("TableArn")) {
            queryParams.add("TableArn=" + args.get("TableArn"));
        }
        if (args.containsKey("ClientToken")) {
            queryParams.add("ClientToken=" + args.get("ClientToken"));
        }
        if (args.containsKey("S3BucketOwner")) {
            queryParams.add("S3BucketOwner=" + args.get("S3BucketOwner"));
        }
        if (args.containsKey("S3SseAlgorithm")) {
            queryParams.add("S3SseAlgorithm=" + args.get("S3SseAlgorithm"));
        }
        if (args.containsKey("S3SseKmsKeyId")) {
            queryParams.add("S3SseKmsKeyId=" + args.get("S3SseKmsKeyId"));
        }
        if (args.containsKey("ExportFormat")) {
            queryParams.add("ExportFormat=" + args.get("ExportFormat"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_export_table_to_point_in_time" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_Export_Table_To_Point_In_TimeTool(MCPServer.APIConfig config) {
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
        Map<String, Object> ExportTimeProperty = new HashMap<>();
        ExportTimeProperty.put("type", "string");
        ExportTimeProperty.put("required", false);
        ExportTimeProperty.put("description", "Input parameter: Time in the past from which to export table data, counted in seconds from the start of the Unix epoch. The table export will be a snapshot of the table's state at this point in time.");
        properties.put("ExportTime", ExportTimeProperty);
        Map<String, Object> S3BucketProperty = new HashMap<>();
        S3BucketProperty.put("type", "string");
        S3BucketProperty.put("required", true);
        S3BucketProperty.put("description", "Input parameter: The name of the Amazon S3 bucket to export the snapshot to.");
        properties.put("S3Bucket", S3BucketProperty);
        Map<String, Object> S3PrefixProperty = new HashMap<>();
        S3PrefixProperty.put("type", "string");
        S3PrefixProperty.put("required", false);
        S3PrefixProperty.put("description", "Input parameter: The Amazon S3 bucket prefix to use as the file name and path of the exported snapshot.");
        properties.put("S3Prefix", S3PrefixProperty);
        Map<String, Object> TableArnProperty = new HashMap<>();
        TableArnProperty.put("type", "string");
        TableArnProperty.put("required", true);
        TableArnProperty.put("description", "Input parameter: The Amazon Resource Name (ARN) associated with the table to export.");
        properties.put("TableArn", TableArnProperty);
        Map<String, Object> ClientTokenProperty = new HashMap<>();
        ClientTokenProperty.put("type", "string");
        ClientTokenProperty.put("required", false);
        ClientTokenProperty.put("description", "Input parameter: <p>Providing a <code>ClientToken</code> makes the call to <code>ExportTableToPointInTimeInput</code> idempotent, meaning that multiple identical calls have the same effect as one single call.</p> <p>A client token is valid for 8 hours after the first request that uses it is completed. After 8 hours, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 8 hours, or the result might not be idempotent.</p> <p>If you submit a request with the same client token but a change in other parameters within the 8-hour idempotency window, DynamoDB returns an <code>ImportConflictException</code>.</p>");
        properties.put("ClientToken", ClientTokenProperty);
        Map<String, Object> S3BucketOwnerProperty = new HashMap<>();
        S3BucketOwnerProperty.put("type", "string");
        S3BucketOwnerProperty.put("required", false);
        S3BucketOwnerProperty.put("description", "Input parameter: The ID of the Amazon Web Services account that owns the bucket the export will be stored in.");
        properties.put("S3BucketOwner", S3BucketOwnerProperty);
        Map<String, Object> S3SseAlgorithmProperty = new HashMap<>();
        S3SseAlgorithmProperty.put("type", "string");
        S3SseAlgorithmProperty.put("required", false);
        S3SseAlgorithmProperty.put("description", "Input parameter: <p>Type of encryption used on the bucket where export data will be stored. Valid values for <code>S3SseAlgorithm</code> are:</p> <ul> <li> <p> <code>AES256</code> - server-side encryption with Amazon S3 managed keys</p> </li> <li> <p> <code>KMS</code> - server-side encryption with KMS managed keys</p> </li> </ul>");
        properties.put("S3SseAlgorithm", S3SseAlgorithmProperty);
        Map<String, Object> S3SseKmsKeyIdProperty = new HashMap<>();
        S3SseKmsKeyIdProperty.put("type", "string");
        S3SseKmsKeyIdProperty.put("required", false);
        S3SseKmsKeyIdProperty.put("description", "Input parameter: The ID of the KMS managed key used to encrypt the S3 bucket where export data will be stored (if applicable).");
        properties.put("S3SseKmsKeyId", S3SseKmsKeyIdProperty);
        Map<String, Object> ExportFormatProperty = new HashMap<>();
        ExportFormatProperty.put("type", "string");
        ExportFormatProperty.put("required", false);
        ExportFormatProperty.put("description", "Input parameter: The format for the exported data. Valid values for <code>ExportFormat</code> are <code>DYNAMODB_JSON</code> or <code>ION</code>.");
        properties.put("ExportFormat", ExportFormatProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_export_table_to_point_in_time",
            "Exports table data to an S3 bucket. The table must have point in time recovery enabled, and you can export data from any time within the point in time recovery window.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_Export_Table_To_Point_In_TimeHandler(config));
    }
    
}