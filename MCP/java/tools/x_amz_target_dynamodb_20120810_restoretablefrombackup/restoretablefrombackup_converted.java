/**
 * MCP Server function for <p>Creates a new table from an existing backup. Any number of users can execute up to 50 concurrent restores (any type of restore) in a given account. </p> <p>You can call <code>RestoreTableFromBackup</code> at a maximum rate of 10 times per second.</p> <p>You must manually set up the following on the restored table:</p> <ul> <li> <p>Auto scaling policies</p> </li> <li> <p>IAM policies</p> </li> <li> <p>Amazon CloudWatch metrics and alarms</p> </li> <li> <p>Tags</p> </li> <li> <p>Stream settings</p> </li> <li> <p>Time to Live (TTL) settings</p> </li> </ul>
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

class Post_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_From_BackupMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_From_BackupHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("GlobalSecondaryIndexOverride")) {
            queryParams.add("GlobalSecondaryIndexOverride=" + args.get("GlobalSecondaryIndexOverride"));
        }
        if (args.containsKey("LocalSecondaryIndexOverride")) {
            queryParams.add("LocalSecondaryIndexOverride=" + args.get("LocalSecondaryIndexOverride"));
        }
        if (args.containsKey("ProvisionedThroughputOverride")) {
            queryParams.add("ProvisionedThroughputOverride=" + args.get("ProvisionedThroughputOverride"));
        }
        if (args.containsKey("SSESpecificationOverride")) {
            queryParams.add("SSESpecificationOverride=" + args.get("SSESpecificationOverride"));
        }
        if (args.containsKey("TargetTableName")) {
            queryParams.add("TargetTableName=" + args.get("TargetTableName"));
        }
        if (args.containsKey("BackupArn")) {
            queryParams.add("BackupArn=" + args.get("BackupArn"));
        }
        if (args.containsKey("BillingModeOverride")) {
            queryParams.add("BillingModeOverride=" + args.get("BillingModeOverride"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_restore_table_from_backup" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_From_BackupTool(MCPServer.APIConfig config) {
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
        Map<String, Object> GlobalSecondaryIndexOverrideProperty = new HashMap<>();
        GlobalSecondaryIndexOverrideProperty.put("type", "string");
        GlobalSecondaryIndexOverrideProperty.put("required", false);
        GlobalSecondaryIndexOverrideProperty.put("description", "Input parameter: List of global secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.");
        properties.put("GlobalSecondaryIndexOverride", GlobalSecondaryIndexOverrideProperty);
        Map<String, Object> LocalSecondaryIndexOverrideProperty = new HashMap<>();
        LocalSecondaryIndexOverrideProperty.put("type", "string");
        LocalSecondaryIndexOverrideProperty.put("required", false);
        LocalSecondaryIndexOverrideProperty.put("description", "Input parameter: List of local secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.");
        properties.put("LocalSecondaryIndexOverride", LocalSecondaryIndexOverrideProperty);
        Map<String, Object> ProvisionedThroughputOverrideProperty = new HashMap<>();
        ProvisionedThroughputOverrideProperty.put("type", "string");
        ProvisionedThroughputOverrideProperty.put("required", false);
        ProvisionedThroughputOverrideProperty.put("description", "Input parameter: Provisioned throughput settings for the restored table.");
        properties.put("ProvisionedThroughputOverride", ProvisionedThroughputOverrideProperty);
        Map<String, Object> SSESpecificationOverrideProperty = new HashMap<>();
        SSESpecificationOverrideProperty.put("type", "string");
        SSESpecificationOverrideProperty.put("required", false);
        SSESpecificationOverrideProperty.put("description", "Input parameter: The new server-side encryption settings for the restored table.");
        properties.put("SSESpecificationOverride", SSESpecificationOverrideProperty);
        Map<String, Object> TargetTableNameProperty = new HashMap<>();
        TargetTableNameProperty.put("type", "string");
        TargetTableNameProperty.put("required", true);
        TargetTableNameProperty.put("description", "Input parameter: The name of the new table to which the backup must be restored.");
        properties.put("TargetTableName", TargetTableNameProperty);
        Map<String, Object> BackupArnProperty = new HashMap<>();
        BackupArnProperty.put("type", "string");
        BackupArnProperty.put("required", true);
        BackupArnProperty.put("description", "Input parameter: The Amazon Resource Name (ARN) associated with the backup.");
        properties.put("BackupArn", BackupArnProperty);
        Map<String, Object> BillingModeOverrideProperty = new HashMap<>();
        BillingModeOverrideProperty.put("type", "string");
        BillingModeOverrideProperty.put("required", false);
        BillingModeOverrideProperty.put("description", "Input parameter: The billing mode of the restored table.");
        properties.put("BillingModeOverride", BillingModeOverrideProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_restore_table_from_backup",
            "<p>Creates a new table from an existing backup. Any number of users can execute up to 50 concurrent restores (any type of restore) in a given account. </p> <p>You can call <code>RestoreTableFromBackup</code> at a maximum rate of 10 times per second.</p> <p>You must manually set up the following on the restored table:</p> <ul> <li> <p>Auto scaling policies</p> </li> <li> <p>IAM policies</p> </li> <li> <p>Amazon CloudWatch metrics and alarms</p> </li> <li> <p>Tags</p> </li> <li> <p>Stream settings</p> </li> <li> <p>Time to Live (TTL) settings</p> </li> </ul>",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_From_BackupHandler(config));
    }
    
}