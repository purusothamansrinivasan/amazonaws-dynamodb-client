/**
 * MCP Server function for Imports table data from an S3 bucket.
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

class Post_X_Amz_Target_Dynamo_Db_20120810_Import_TableMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_Import_TableHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("InputFormat")) {
            queryParams.add("InputFormat=" + args.get("InputFormat"));
        }
        if (args.containsKey("InputFormatOptions")) {
            queryParams.add("InputFormatOptions=" + args.get("InputFormatOptions"));
        }
        if (args.containsKey("S3BucketSource")) {
            queryParams.add("S3BucketSource=" + args.get("S3BucketSource"));
        }
        if (args.containsKey("TableCreationParameters")) {
            queryParams.add("TableCreationParameters=" + args.get("TableCreationParameters"));
        }
        if (args.containsKey("ClientToken")) {
            queryParams.add("ClientToken=" + args.get("ClientToken"));
        }
        if (args.containsKey("InputCompressionType")) {
            queryParams.add("InputCompressionType=" + args.get("InputCompressionType"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_import_table" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_Import_TableTool(MCPServer.APIConfig config) {
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
        Map<String, Object> InputFormatProperty = new HashMap<>();
        InputFormatProperty.put("type", "string");
        InputFormatProperty.put("required", true);
        InputFormatProperty.put("description", "Input parameter: The format of the source data. Valid values for <code>ImportFormat</code> are <code>CSV</code>, <code>DYNAMODB_JSON</code> or <code>ION</code>.");
        properties.put("InputFormat", InputFormatProperty);
        Map<String, Object> InputFormatOptionsProperty = new HashMap<>();
        InputFormatOptionsProperty.put("type", "string");
        InputFormatOptionsProperty.put("required", false);
        InputFormatOptionsProperty.put("description", "Input parameter: Additional properties that specify how the input is formatted,");
        properties.put("InputFormatOptions", InputFormatOptionsProperty);
        Map<String, Object> S3BucketSourceProperty = new HashMap<>();
        S3BucketSourceProperty.put("type", "string");
        S3BucketSourceProperty.put("required", true);
        S3BucketSourceProperty.put("description", "Input parameter: The S3 bucket that provides the source for the import.");
        properties.put("S3BucketSource", S3BucketSourceProperty);
        Map<String, Object> TableCreationParametersProperty = new HashMap<>();
        TableCreationParametersProperty.put("type", "string");
        TableCreationParametersProperty.put("required", true);
        TableCreationParametersProperty.put("description", "Input parameter: Parameters for the table to import the data into.");
        properties.put("TableCreationParameters", TableCreationParametersProperty);
        Map<String, Object> ClientTokenProperty = new HashMap<>();
        ClientTokenProperty.put("type", "string");
        ClientTokenProperty.put("required", false);
        ClientTokenProperty.put("description", "Input parameter: <p>Providing a <code>ClientToken</code> makes the call to <code>ImportTableInput</code> idempotent, meaning that multiple identical calls have the same effect as one single call.</p> <p>A client token is valid for 8 hours after the first request that uses it is completed. After 8 hours, any request with the same client token is treated as a new request. Do not resubmit the same request with the same client token for more than 8 hours, or the result might not be idempotent.</p> <p>If you submit a request with the same client token but a change in other parameters within the 8-hour idempotency window, DynamoDB returns an <code>IdempotentParameterMismatch</code> exception.</p>");
        properties.put("ClientToken", ClientTokenProperty);
        Map<String, Object> InputCompressionTypeProperty = new HashMap<>();
        InputCompressionTypeProperty.put("type", "string");
        InputCompressionTypeProperty.put("required", false);
        InputCompressionTypeProperty.put("description", "Input parameter: Type of compression to be used on the input coming from the imported table.");
        properties.put("InputCompressionType", InputCompressionTypeProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_import_table",
            "Imports table data from an S3 bucket.",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_Import_TableHandler(config));
    }
    
}