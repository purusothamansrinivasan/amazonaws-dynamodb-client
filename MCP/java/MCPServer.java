/**
 * Main MCP Server - Handles MCP JSON-RPC protocol
 */

import java.io.*;
import java.util.*;
import java.util.function.Function;
import java.util.concurrent.CompletableFuture;
import com.fasterxml.jackson.databind.ObjectMapper;
import com.fasterxml.jackson.databind.JsonNode;
import com.fasterxml.jackson.databind.node.ObjectNode;
import com.fasterxml.jackson.databind.node.ArrayNode;

// Import Post_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_From_BackupMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_Global_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_BackupMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Create_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Table_Replica_Auto_ScalingMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_List_TablesMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Execute_TransactionMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_List_BackupsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_Time_To_LiveMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Tag_ResourceMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Global_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_QueryMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Continuous_BackupsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_To_Point_In_TimeMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_Continuous_BackupsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_ExportMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Put_ItemMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_LimitsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Import_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Batch_Get_ItemMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_Table_Replica_Auto_ScalingMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Create_BackupMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_List_ImportsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_List_Contributor_InsightsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_List_ExportsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Batch_Write_ItemMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Disable_Kinesis_Streaming_DestinationMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Delete_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_ItemMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Delete_BackupMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_Global_Table_SettingsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Kinesis_Streaming_DestinationMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_List_Global_TablesMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_EndpointsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Delete_ItemMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Create_Global_TableMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Batch_Execute_StatementMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Get_ItemMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Update_Contributor_InsightsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Contributor_InsightsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Time_To_LiveMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Transact_Get_ItemsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Untag_ResourceMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Execute_StatementMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Export_Table_To_Point_In_TimeMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_ImportMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Global_Table_SettingsMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_ScanMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Enable_Kinesis_Streaming_DestinationMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_List_Tags_Of_ResourceMCPTool - included in same package
// Import Post_X_Amz_Target_Dynamo_Db_20120810_Transact_Write_ItemsMCPTool - included in same package

public class MCPServer {
    
    // Common classes that all tool classes use
    public static class APIConfig {
        private final String baseUrl;
        private final String apiKey;
        
        public APIConfig(String baseUrl, String apiKey) {
            this.baseUrl = baseUrl;
            this.apiKey = apiKey;
        }
        
        public String getBaseUrl() { return baseUrl; }
        public String getApiKey() { return apiKey; }
    }
    
    public static class MCPRequest {
        private Map<String, Object> params;
        
        public Map<String, Object> getParams() { return params; }
        public void setParams(Map<String, Object> params) { this.params = params; }
        
        @SuppressWarnings("unchecked")
        public Map<String, Object> getArguments() {
            if (params != null && params.containsKey("arguments")) {
                return (Map<String, Object>) params.get("arguments");
            }
            return new HashMap<>();
        }
    }
    
    public static class MCPToolResult {
        private final String content;
        private final boolean isError;
        
        public MCPToolResult(String content, boolean isError) {
            this.content = content;
            this.isError = isError;
        }
        
        public MCPToolResult(String content) {
            this(content, false);
        }
        
        public String getContent() { return content; }
        public boolean isError() { return isError; }
    }
    
    public static class ToolDefinition {
        private final String name;
        private final String description;
        private final Map<String, Object> parameters;
        
        public ToolDefinition(String name, String description, Map<String, Object> parameters) {
            this.name = name;
            this.description = description;
            this.parameters = parameters;
        }
        
        public String getName() { return name; }
        public String getDescription() { return description; }
        public Map<String, Object> getParameters() { return parameters; }
    }
    
    public static class Tool {
        private final ToolDefinition definition;
        private final Function<MCPRequest, MCPToolResult> handler;
        
        public Tool(ToolDefinition definition, Function<MCPRequest, MCPToolResult> handler) {
            this.definition = definition;
            this.handler = handler;
        }
        
        public ToolDefinition getDefinition() { return definition; }
        public Function<MCPRequest, MCPToolResult> getHandler() { return handler; }
    }
    
    private static final ObjectMapper mapper = new ObjectMapper();
    private static List<Tool> tools = new ArrayList<>();
    private static APIConfig config;
    
    public static void main(String[] args) {
        try {
            // Initialize configuration
            String baseUrl = System.getenv("API_BASE_URL");
            String apiKey = System.getenv("API_KEY");
            
            if (baseUrl == null || apiKey == null) {
                System.err.println("Error: Please set API_BASE_URL and API_KEY environment variables");
                System.exit(1);
            }
            
            config = new APIConfig(baseUrl, apiKey);
            
            // Register all tools
            tools = Arrays.asList(
            Post_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_From_BackupMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_From_BackupTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_Global_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_Global_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_BackupMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_BackupTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Create_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Create_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Table_Replica_Auto_ScalingMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_Table_Replica_Auto_ScalingTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_List_TablesMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_List_TablesTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Execute_TransactionMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Execute_TransactionTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_List_BackupsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_List_BackupsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_Time_To_LiveMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_Time_To_LiveTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Tag_ResourceMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Tag_ResourceTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Global_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_Global_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_QueryMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_QueryTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Continuous_BackupsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_Continuous_BackupsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_To_Point_In_TimeMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Restore_Table_To_Point_In_TimeTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_Continuous_BackupsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_Continuous_BackupsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_ExportMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_ExportTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Put_ItemMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Put_ItemTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_LimitsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_LimitsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Import_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Import_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Batch_Get_ItemMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Batch_Get_ItemTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_Table_Replica_Auto_ScalingMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_Table_Replica_Auto_ScalingTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Create_BackupMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Create_BackupTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_List_ImportsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_List_ImportsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_List_Contributor_InsightsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_List_Contributor_InsightsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_List_ExportsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_List_ExportsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Batch_Write_ItemMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Batch_Write_ItemTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Disable_Kinesis_Streaming_DestinationMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Disable_Kinesis_Streaming_DestinationTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Delete_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Delete_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_ItemMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_ItemTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Delete_BackupMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Delete_BackupTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_Global_Table_SettingsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_Global_Table_SettingsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Kinesis_Streaming_DestinationMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_Kinesis_Streaming_DestinationTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_List_Global_TablesMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_List_Global_TablesTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_EndpointsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_EndpointsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Delete_ItemMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Delete_ItemTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Create_Global_TableMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Create_Global_TableTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Batch_Execute_StatementMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Batch_Execute_StatementTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Get_ItemMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Get_ItemTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Update_Contributor_InsightsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Update_Contributor_InsightsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Contributor_InsightsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_Contributor_InsightsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Time_To_LiveMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_Time_To_LiveTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Transact_Get_ItemsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Transact_Get_ItemsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Untag_ResourceMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Untag_ResourceTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Execute_StatementMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Execute_StatementTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Export_Table_To_Point_In_TimeMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Export_Table_To_Point_In_TimeTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_ImportMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_ImportTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Describe_Global_Table_SettingsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Describe_Global_Table_SettingsTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_ScanMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_ScanTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Enable_Kinesis_Streaming_DestinationMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Enable_Kinesis_Streaming_DestinationTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_List_Tags_Of_ResourceMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_List_Tags_Of_ResourceTool(config),
            Post_X_Amz_Target_Dynamo_Db_20120810_Transact_Write_ItemsMCPTool.createPost_X_Amz_Target_Dynamo_Db_20120810_Transact_Write_ItemsTool(config)
            );
            
            System.err.println("MCP Server started with " + tools.size() + " tools");
            
            // Start JSON-RPC server
            runServer();
            
        } catch (Exception e) {
            System.err.println("Failed to start MCP server: " + e.getMessage());
            e.printStackTrace();
            System.exit(1);
        }
    }
    
    private static void runServer() throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String line;
        
        while ((line = reader.readLine()) != null) {
            JsonNode request = null;
            try {
                request = mapper.readTree(line);
                JsonNode response = handleRequest(request);
                
                if (response != null) {
                    System.out.println(mapper.writeValueAsString(response));
                }
                
            } catch (Exception e) {
                // Send error response
                ObjectNode errorResponse = mapper.createObjectNode();
                errorResponse.put("jsonrpc", "2.0");
                if (request != null && request.has("id")) {
                    errorResponse.put("id", request.get("id").asText());
                } else {
                    errorResponse.putNull("id");
                }
                
                ObjectNode error = mapper.createObjectNode();
                error.put("code", -32603);
                error.put("message", "Internal error: " + e.getMessage());
                errorResponse.set("error", error);
                
                System.out.println(mapper.writeValueAsString(errorResponse));
            }
        }
    }
    
    private static JsonNode handleRequest(JsonNode request) {
        if (!request.has("method")) {
            return createErrorResponse(request, -32600, "Invalid Request");
        }
        
        String method = request.get("method").asText();
        JsonNode params = request.get("params");
        String id = request.has("id") ? request.get("id").asText() : null;
        
        switch (method) {
            case "initialize":
                return handleInitialize(id);
            case "tools/list":
                return handleToolsList(id);
            case "tools/call":
                return handleToolsCall(id, params);
            default:
                return createErrorResponse(request, -32601, "Method not found");
        }
    }
    
    private static JsonNode handleInitialize(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        result.put("protocolVersion", "2024-11-05");
        
        ObjectNode capabilities = mapper.createObjectNode();
        ObjectNode toolsCapability = mapper.createObjectNode();
        toolsCapability.put("listChanged", false);
        capabilities.set("tools", toolsCapability);
        result.set("capabilities", capabilities);
        
        ObjectNode serverInfo = mapper.createObjectNode();
        serverInfo.put("name", "Opsera MCP Server");
        serverInfo.put("version", "1.0.0");
        result.set("serverInfo", serverInfo);
        
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsList(String id) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", id);
        
        ObjectNode result = mapper.createObjectNode();
        ArrayNode toolsArray = mapper.createArrayNode();
        
        for (Tool tool : tools) {
            ObjectNode toolDef = mapper.createObjectNode();
            toolDef.put("name", tool.getDefinition().getName());
            toolDef.put("description", tool.getDefinition().getDescription());
            
            // Convert parameters to JSON
            ObjectNode inputSchema = mapper.valueToTree(tool.getDefinition().getParameters());
            toolDef.set("inputSchema", inputSchema);
            
            toolsArray.add(toolDef);
        }
        
        result.set("tools", toolsArray);
        response.set("result", result);
        return response;
    }
    
    private static JsonNode handleToolsCall(String id, JsonNode params) {
        if (!params.has("name")) {
            return createErrorResponse(null, -32602, "Invalid params: missing 'name'");
        }
        
        String toolName = params.get("name").asText();
        JsonNode arguments = params.has("arguments") ? params.get("arguments") : mapper.createObjectNode();
        
        // Find the tool
        Tool tool = null;
        for (Tool t : tools) {
            if (t.getDefinition().getName().equals(toolName)) {
                tool = t;
                break;
            }
        }
        
        if (tool == null) {
            return createErrorResponse(null, -32602, "Tool not found: " + toolName);
        }
        
        try {
            // Convert arguments to Map
            Map<String, Object> argsMap = mapper.convertValue(arguments, Map.class);
            
            // Create MCP request
            MCPRequest mcpRequest = new MCPRequest();
            Map<String, Object> requestParams = new HashMap<>();
            requestParams.put("arguments", argsMap);
            mcpRequest.setParams(requestParams);
            
            // Execute tool
            MCPToolResult result = tool.getHandler().apply(mcpRequest);
            
            // Create response
            ObjectNode response = mapper.createObjectNode();
            response.put("jsonrpc", "2.0");
            response.put("id", id);
            
            ObjectNode resultObj = mapper.createObjectNode();
            ArrayNode content = mapper.createArrayNode();
            
            ObjectNode textContent = mapper.createObjectNode();
            textContent.put("type", "text");
            textContent.put("text", result.getContent());
            content.add(textContent);
            
            resultObj.set("content", content);
            resultObj.put("isError", result.isError());
            
            response.set("result", resultObj);
            return response;
            
        } catch (Exception e) {
            return createErrorResponse(null, -32603, "Tool execution failed: " + e.getMessage());
        }
    }
    
    private static JsonNode createErrorResponse(JsonNode request, int code, String message) {
        ObjectNode response = mapper.createObjectNode();
        response.put("jsonrpc", "2.0");
        response.put("id", request != null && request.has("id") ? request.get("id").asText() : null);
        
        ObjectNode error = mapper.createObjectNode();
        error.put("code", code);
        error.put("message", message);
        response.set("error", error);
        
        return response;
    }
}