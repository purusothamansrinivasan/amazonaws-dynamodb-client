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

class Post_X_Amz_Target_Dynamo_Db_20120810_Update_Global_Table_SettingsMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_Update_Global_Table_SettingsHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("GlobalTableName")) {
            queryParams.add("GlobalTableName=" + args.get("GlobalTableName"));
        }
        if (args.containsKey("GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate")) {
            queryParams.add("GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate=" + args.get("GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate"));
        }
        if (args.containsKey("GlobalTableProvisionedWriteCapacityUnits")) {
            queryParams.add("GlobalTableProvisionedWriteCapacityUnits=" + args.get("GlobalTableProvisionedWriteCapacityUnits"));
        }
        if (args.containsKey("ReplicaSettingsUpdate")) {
            queryParams.add("ReplicaSettingsUpdate=" + args.get("ReplicaSettingsUpdate"));
        }
        if (args.containsKey("GlobalTableBillingMode")) {
            queryParams.add("GlobalTableBillingMode=" + args.get("GlobalTableBillingMode"));
        }
        if (args.containsKey("GlobalTableGlobalSecondaryIndexSettingsUpdate")) {
            queryParams.add("GlobalTableGlobalSecondaryIndexSettingsUpdate=" + args.get("GlobalTableGlobalSecondaryIndexSettingsUpdate"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_update_global_table_settings" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_Update_Global_Table_SettingsTool(MCPServer.APIConfig config) {
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
        Map<String, Object> GlobalTableNameProperty = new HashMap<>();
        GlobalTableNameProperty.put("type", "string");
        GlobalTableNameProperty.put("required", true);
        GlobalTableNameProperty.put("description", "Input parameter: The name of the global table");
        properties.put("GlobalTableName", GlobalTableNameProperty);
        Map<String, Object> GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdateProperty = new HashMap<>();
        GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdateProperty.put("type", "string");
        GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdateProperty.put("required", false);
        GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdateProperty.put("description", "Input parameter: Auto scaling settings for managing provisioned write capacity for the global table.");
        properties.put("GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdate", GlobalTableProvisionedWriteCapacityAutoScalingSettingsUpdateProperty);
        Map<String, Object> GlobalTableProvisionedWriteCapacityUnitsProperty = new HashMap<>();
        GlobalTableProvisionedWriteCapacityUnitsProperty.put("type", "string");
        GlobalTableProvisionedWriteCapacityUnitsProperty.put("required", false);
        GlobalTableProvisionedWriteCapacityUnitsProperty.put("description", "Input parameter: The maximum number of writes consumed per second before DynamoDB returns a <code>ThrottlingException.</code>");
        properties.put("GlobalTableProvisionedWriteCapacityUnits", GlobalTableProvisionedWriteCapacityUnitsProperty);
        Map<String, Object> ReplicaSettingsUpdateProperty = new HashMap<>();
        ReplicaSettingsUpdateProperty.put("type", "string");
        ReplicaSettingsUpdateProperty.put("required", false);
        ReplicaSettingsUpdateProperty.put("description", "Input parameter: Represents the settings for a global table in a Region that will be modified.");
        properties.put("ReplicaSettingsUpdate", ReplicaSettingsUpdateProperty);
        Map<String, Object> GlobalTableBillingModeProperty = new HashMap<>();
        GlobalTableBillingModeProperty.put("type", "string");
        GlobalTableBillingModeProperty.put("required", false);
        GlobalTableBillingModeProperty.put("description", "Input parameter: <p>The billing mode of the global table. If <code>GlobalTableBillingMode</code> is not specified, the global table defaults to <code>PROVISIONED</code> capacity billing mode.</p> <ul> <li> <p> <code>PROVISIONED</code> - We recommend using <code>PROVISIONED</code> for predictable workloads. <code>PROVISIONED</code> sets the billing mode to <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual\">Provisioned Mode</a>.</p> </li> <li> <p> <code>PAY_PER_REQUEST</code> - We recommend using <code>PAY_PER_REQUEST</code> for unpredictable workloads. <code>PAY_PER_REQUEST</code> sets the billing mode to <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand\">On-Demand Mode</a>. </p> </li> </ul>");
        properties.put("GlobalTableBillingMode", GlobalTableBillingModeProperty);
        Map<String, Object> GlobalTableGlobalSecondaryIndexSettingsUpdateProperty = new HashMap<>();
        GlobalTableGlobalSecondaryIndexSettingsUpdateProperty.put("type", "string");
        GlobalTableGlobalSecondaryIndexSettingsUpdateProperty.put("required", false);
        GlobalTableGlobalSecondaryIndexSettingsUpdateProperty.put("description", "Input parameter: Represents the settings of a global secondary index for a global table that will be modified.");
        properties.put("GlobalTableGlobalSecondaryIndexSettingsUpdate", GlobalTableGlobalSecondaryIndexSettingsUpdateProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_update_global_table_settings",
            "No description available",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_Update_Global_Table_SettingsHandler(config));
    }
    
}