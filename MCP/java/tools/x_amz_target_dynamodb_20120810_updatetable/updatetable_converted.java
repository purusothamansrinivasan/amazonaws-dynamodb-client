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

class Post_X_Amz_Target_Dynamo_Db_20120810_Update_TableMCPTool {
    
    public static Function<MCPServer.MCPRequest, MCPServer.MCPToolResult> getPost_X_Amz_Target_Dynamo_Db_20120810_Update_TableHandler(MCPServer.APIConfig config) {
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
        if (args.containsKey("TableName")) {
            queryParams.add("TableName=" + args.get("TableName"));
        }
        if (args.containsKey("TableClass")) {
            queryParams.add("TableClass=" + args.get("TableClass"));
        }
        if (args.containsKey("AttributeDefinitions")) {
            queryParams.add("AttributeDefinitions=" + args.get("AttributeDefinitions"));
        }
        if (args.containsKey("BillingMode")) {
            queryParams.add("BillingMode=" + args.get("BillingMode"));
        }
        if (args.containsKey("DeletionProtectionEnabled")) {
            queryParams.add("DeletionProtectionEnabled=" + args.get("DeletionProtectionEnabled"));
        }
        if (args.containsKey("ReplicaUpdates")) {
            queryParams.add("ReplicaUpdates=" + args.get("ReplicaUpdates"));
        }
        if (args.containsKey("ProvisionedThroughput")) {
            queryParams.add("ProvisionedThroughput=" + args.get("ProvisionedThroughput"));
        }
        if (args.containsKey("SSESpecification")) {
            queryParams.add("SSESpecification=" + args.get("SSESpecification"));
        }
        if (args.containsKey("StreamSpecification")) {
            queryParams.add("StreamSpecification=" + args.get("StreamSpecification"));
        }
        if (args.containsKey("GlobalSecondaryIndexUpdates")) {
            queryParams.add("GlobalSecondaryIndexUpdates=" + args.get("GlobalSecondaryIndexUpdates"));
        }
                
                String queryString = queryParams.isEmpty() ? "" : "?" + String.join("&", queryParams);
                String url = config.getBaseUrl() + "/api/v2/post_x_amz_target_dynamo_db_20120810_update_table" + queryString;
                
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
    
    public static MCPServer.Tool createPost_X_Amz_Target_Dynamo_Db_20120810_Update_TableTool(MCPServer.APIConfig config) {
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
        Map<String, Object> TableNameProperty = new HashMap<>();
        TableNameProperty.put("type", "string");
        TableNameProperty.put("required", true);
        TableNameProperty.put("description", "Input parameter: The name of the table to be updated.");
        properties.put("TableName", TableNameProperty);
        Map<String, Object> TableClassProperty = new HashMap<>();
        TableClassProperty.put("type", "string");
        TableClassProperty.put("required", false);
        TableClassProperty.put("description", "Input parameter: The table class of the table to be updated. Valid values are <code>STANDARD</code> and <code>STANDARD_INFREQUENT_ACCESS</code>.");
        properties.put("TableClass", TableClassProperty);
        Map<String, Object> AttributeDefinitionsProperty = new HashMap<>();
        AttributeDefinitionsProperty.put("type", "string");
        AttributeDefinitionsProperty.put("required", false);
        AttributeDefinitionsProperty.put("description", "Input parameter: An array of attributes that describe the key schema for the table and indexes. If you are adding a new global secondary index to the table, <code>AttributeDefinitions</code> must include the key element(s) of the new index.");
        properties.put("AttributeDefinitions", AttributeDefinitionsProperty);
        Map<String, Object> BillingModeProperty = new HashMap<>();
        BillingModeProperty.put("type", "string");
        BillingModeProperty.put("required", false);
        BillingModeProperty.put("description", "Input parameter: <p>Controls how you are charged for read and write throughput and how you manage capacity. When switching from pay-per-request to provisioned capacity, initial provisioned capacity values must be set. The initial provisioned capacity values are estimated based on the consumed read and write capacity of your table and global secondary indexes over the past 30 minutes.</p> <ul> <li> <p> <code>PROVISIONED</code> - We recommend using <code>PROVISIONED</code> for predictable workloads. <code>PROVISIONED</code> sets the billing mode to <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.ProvisionedThroughput.Manual\">Provisioned Mode</a>.</p> </li> <li> <p> <code>PAY_PER_REQUEST</code> - We recommend using <code>PAY_PER_REQUEST</code> for unpredictable workloads. <code>PAY_PER_REQUEST</code> sets the billing mode to <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/HowItWorks.ReadWriteCapacityMode.html#HowItWorks.OnDemand\">On-Demand Mode</a>. </p> </li> </ul>");
        properties.put("BillingMode", BillingModeProperty);
        Map<String, Object> DeletionProtectionEnabledProperty = new HashMap<>();
        DeletionProtectionEnabledProperty.put("type", "string");
        DeletionProtectionEnabledProperty.put("required", false);
        DeletionProtectionEnabledProperty.put("description", "Input parameter: Indicates whether deletion protection is to be enabled (true) or disabled (false) on the table.");
        properties.put("DeletionProtectionEnabled", DeletionProtectionEnabledProperty);
        Map<String, Object> ReplicaUpdatesProperty = new HashMap<>();
        ReplicaUpdatesProperty.put("type", "string");
        ReplicaUpdatesProperty.put("required", false);
        ReplicaUpdatesProperty.put("description", "Input parameter: <p>A list of replica update actions (create, delete, or update) for the table.</p> <note> <p>This property only applies to <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/globaltables.V2.html\">Version 2019.11.21 (Current)</a> of global tables. </p> </note>");
        properties.put("ReplicaUpdates", ReplicaUpdatesProperty);
        Map<String, Object> ProvisionedThroughputProperty = new HashMap<>();
        ProvisionedThroughputProperty.put("type", "string");
        ProvisionedThroughputProperty.put("required", false);
        ProvisionedThroughputProperty.put("description", "Input parameter: The new provisioned throughput settings for the specified table or index.");
        properties.put("ProvisionedThroughput", ProvisionedThroughputProperty);
        Map<String, Object> SSESpecificationProperty = new HashMap<>();
        SSESpecificationProperty.put("type", "string");
        SSESpecificationProperty.put("required", false);
        SSESpecificationProperty.put("description", "Input parameter: The new server-side encryption settings for the specified table.");
        properties.put("SSESpecification", SSESpecificationProperty);
        Map<String, Object> StreamSpecificationProperty = new HashMap<>();
        StreamSpecificationProperty.put("type", "string");
        StreamSpecificationProperty.put("required", false);
        StreamSpecificationProperty.put("description", "Input parameter: <p>Represents the DynamoDB Streams configuration for the table.</p> <note> <p>You receive a <code>ResourceInUseException</code> if you try to enable a stream on a table that already has a stream, or if you try to disable a stream on a table that doesn't have a stream.</p> </note>");
        properties.put("StreamSpecification", StreamSpecificationProperty);
        Map<String, Object> GlobalSecondaryIndexUpdatesProperty = new HashMap<>();
        GlobalSecondaryIndexUpdatesProperty.put("type", "string");
        GlobalSecondaryIndexUpdatesProperty.put("required", false);
        GlobalSecondaryIndexUpdatesProperty.put("description", "Input parameter: <p>An array of one or more global secondary indexes for the table. For each index in the array, you can request one action:</p> <ul> <li> <p> <code>Create</code> - add a new global secondary index to the table.</p> </li> <li> <p> <code>Update</code> - modify the provisioned throughput settings of an existing global secondary index.</p> </li> <li> <p> <code>Delete</code> - remove a global secondary index from the table.</p> </li> </ul> <p>You can create or delete only one global secondary index per <code>UpdateTable</code> operation.</p> <p>For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/GSI.OnlineOps.html\">Managing Global Secondary Indexes</a> in the <i>Amazon DynamoDB Developer Guide</i>. </p>");
        properties.put("GlobalSecondaryIndexUpdates", GlobalSecondaryIndexUpdatesProperty);
        parameters.put("properties", properties);
        
        MCPServer.ToolDefinition definition = new MCPServer.ToolDefinition(
            "post_x_amz_target_dynamo_db_20120810_update_table",
            "No description available",
            parameters
        );
        
        return new MCPServer.Tool(definition, getPost_X_Amz_Target_Dynamo_Db_20120810_Update_TableHandler(config));
    }
    
}