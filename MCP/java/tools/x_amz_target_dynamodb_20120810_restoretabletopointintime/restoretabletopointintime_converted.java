/**
 * MCP Server function for <p>Restores the specified table to the specified point in time within <code>EarliestRestorableDateTime</code> and <code>LatestRestorableDateTime</code>. You can restore your table to any point in time during the last 35 days. Any number of users can execute up to 4 concurrent restores (any type of restore) in a given account. </p> <p> When you restore using point in time recovery, DynamoDB restores your table data to the state based on the selected date and time (day:hour:minute:second) to a new table. </p> <p> Along with data, the following are also included on the new restored table using point in time recovery: </p> <ul> <li> <p>Global secondary indexes (GSIs)</p> </li> <li> <p>Local secondary indexes (LSIs)</p> </li> <li> <p>Provisioned read and write capacity</p> </li> <li> <p>Encryption settings</p> <important> <p> All these settings come from the current settings of the source table at the time of restore. </p> </important> </li> </ul> <p>You must manually set up the following on the restored table:</p> <ul> <li> <p>Auto scaling policies</p> </li> <li> <p>IAM policies</p> </li> <li> <p>Amazon CloudWatch metrics and alarms</p> </li> <li> <p>Tags</p> </li> <li> <p>Stream settings</p> </li> <li> <p>Time to Live (TTL) settings</p> </li> <li> <p>Point in time recovery settings</p> </li> </ul>
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

class Post_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_To_Point_In_TimeMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_To_Point_In_TimeHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("BillingModeOverride")) {
            queryParams.add("BillingModeOverride=" + args.get("BillingModeOverride"));
        }
        if (args.containsKey("GlobalSecondaryIndexOverride")) {
            queryParams.add("GlobalSecondaryIndexOverride=" + args.get("GlobalSecondaryIndexOverride"));
        }
        if (args.containsKey("ProvisionedThroughputOverride")) {
            queryParams.add("ProvisionedThroughputOverride=" + args.get("ProvisionedThroughputOverride"));
        }
        if (args.containsKey("SSESpecificationOverride")) {
            queryParams.add("SSESpecificationOverride=" + args.get("SSESpecificationOverride"));
        }
        if (args.containsKey("SourceTableName")) {
            queryParams.add("SourceTableName=" + args.get("SourceTableName"));
        }
        if (args.containsKey("TargetTableName")) {
            queryParams.add("TargetTableName=" + args.get("TargetTableName"));
        }
        if (args.containsKey("RestoreDateTime")) {
            queryParams.add("RestoreDateTime=" + args.get("RestoreDateTime"));
        }
        if (args.containsKey("SourceTableArn")) {
            queryParams.add("SourceTableArn=" + args.get("SourceTableArn"));
        }
        if (args.containsKey("LocalSecondaryIndexOverride")) {
            queryParams.add("LocalSecondaryIndexOverride=" + args.get("LocalSecondaryIndexOverride"));
        }
        if (args.containsKey("UseLatestRestorableTime")) {
            queryParams.add("UseLatestRestorableTime=" + args.get("UseLatestRestorableTime"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_restore_table_to_point_in_time" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_To_Point_In_TimeTool(MCPServer.APIConfig config) {
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
        Map<String, Object> BillingModeOverrideProperty = new HashMap<>();
        BillingModeOverrideProperty.put("type", "string");
        BillingModeOverrideProperty.put("required", false);
        BillingModeOverrideProperty.put("description", "Input parameter: The billing mode of the restored table.");
        properties.put("BillingModeOverride", BillingModeOverrideProperty);
        Map<String, Object> GlobalSecondaryIndexOverrideProperty = new HashMap<>();
        GlobalSecondaryIndexOverrideProperty.put("type", "string");
        GlobalSecondaryIndexOverrideProperty.put("required", false);
        GlobalSecondaryIndexOverrideProperty.put("description", "Input parameter: List of global secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.");
        properties.put("GlobalSecondaryIndexOverride", GlobalSecondaryIndexOverrideProperty);
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
        Map<String, Object> SourceTableNameProperty = new HashMap<>();
        SourceTableNameProperty.put("type", "string");
        SourceTableNameProperty.put("required", false);
        SourceTableNameProperty.put("description", "Input parameter: Name of the source table that is being restored.");
        properties.put("SourceTableName", SourceTableNameProperty);
        Map<String, Object> TargetTableNameProperty = new HashMap<>();
        TargetTableNameProperty.put("type", "string");
        TargetTableNameProperty.put("required", true);
        TargetTableNameProperty.put("description", "Input parameter: The name of the new table to which it must be restored to.");
        properties.put("TargetTableName", TargetTableNameProperty);
        Map<String, Object> RestoreDateTimeProperty = new HashMap<>();
        RestoreDateTimeProperty.put("type", "string");
        RestoreDateTimeProperty.put("required", false);
        RestoreDateTimeProperty.put("description", "Input parameter: Time in the past to restore the table to.");
        properties.put("RestoreDateTime", RestoreDateTimeProperty);
        Map<String, Object> SourceTableArnProperty = new HashMap<>();
        SourceTableArnProperty.put("type", "string");
        SourceTableArnProperty.put("required", false);
        SourceTableArnProperty.put("description", "Input parameter: The DynamoDB table that will be restored. This value is an Amazon Resource Name (ARN).");
        properties.put("SourceTableArn", SourceTableArnProperty);
        Map<String, Object> LocalSecondaryIndexOverrideProperty = new HashMap<>();
        LocalSecondaryIndexOverrideProperty.put("type", "string");
        LocalSecondaryIndexOverrideProperty.put("required", false);
        LocalSecondaryIndexOverrideProperty.put("description", "Input parameter: List of local secondary indexes for the restored table. The indexes provided should match existing secondary indexes. You can choose to exclude some or all of the indexes at the time of restore.");
        properties.put("LocalSecondaryIndexOverride", LocalSecondaryIndexOverrideProperty);
        Map<String, Object> UseLatestRestorableTimeProperty = new HashMap<>();
        UseLatestRestorableTimeProperty.put("type", "string");
        UseLatestRestorableTimeProperty.put("required", false);
        UseLatestRestorableTimeProperty.put("description", "Input parameter: Restore the table to the latest possible time. <code>LatestRestorableDateTime</code> is typically 5 minutes before the current time.");
        properties.put("UseLatestRestorableTime", UseLatestRestorableTimeProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_restore_table_to_point_in_time",
            "<p>Restores the specified table to the specified point in time within <code>EarliestRestorableDateTime</code> and <code>LatestRestorableDateTime</code>. You can restore your table to any point in time during the last 35 days. Any number of users can execute up to 4 concurrent restores (any type of restore) in a given account. </p> <p> When you restore using point in time recovery, DynamoDB restores your table data to the state based on the selected date and time (day:hour:minute:second) to a new table. </p> <p> Along with data, the following are also included on the new restored table using point in time recovery: </p> <ul> <li> <p>Global secondary indexes (GSIs)</p> </li> <li> <p>Local secondary indexes (LSIs)</p> </li> <li> <p>Provisioned read and write capacity</p> </li> <li> <p>Encryption settings</p> <important> <p> All these settings come from the current settings of the source table at the time of restore. </p> </important> </li> </ul> <p>You must manually set up the following on the restored table:</p> <ul> <li> <p>Auto scaling policies</p> </li> <li> <p>IAM policies</p> </li> <li> <p>Amazon CloudWatch metrics and alarms</p> </li> <li> <p>Tags</p> </li> <li> <p>Stream settings</p> </li> <li> <p>Time to Live (TTL) settings</p> </li> <li> <p>Point in time recovery settings</p> </li> </ul>",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_To_Point_In_TimeHandler(config));
    }
    
}