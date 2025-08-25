package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/amazon-dynamodb/mcp-server/config"
	"github.com/amazon-dynamodb/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func BatchgetitemHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		queryParams := make([]string, 0)
		if val, ok := args["RequestItems"]; ok {
			queryParams = append(queryParams, fmt.Sprintf("RequestItems=%v", val))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody models.BatchGetItemInput
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/#X-Amz-Target=DynamoDB_20120810.BatchGetItem%s", cfg.BaseURL, queryString)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		if cfg.BearerToken != "" {
			req.Header.Set("X-Amz-Security-Token", cfg.BearerToken)
		}
		req.Header.Set("Accept", "application/json")
		if val, ok := args["X-Amz-Content-Sha256"]; ok {
			req.Header.Set("X-Amz-Content-Sha256", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Date"]; ok {
			req.Header.Set("X-Amz-Date", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Algorithm"]; ok {
			req.Header.Set("X-Amz-Algorithm", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Credential"]; ok {
			req.Header.Set("X-Amz-Credential", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Security-Token"]; ok {
			req.Header.Set("X-Amz-Security-Token", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Signature"]; ok {
			req.Header.Set("X-Amz-Signature", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-SignedHeaders"]; ok {
			req.Header.Set("X-Amz-SignedHeaders", fmt.Sprintf("%v", val))
		}
		if val, ok := args["X-Amz-Target"]; ok {
			req.Header.Set("X-Amz-Target", fmt.Sprintf("%v", val))
		}

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result models.BatchGetItemOutput
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateBatchgetitemTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("post_#X-Amz-Target=DynamoDB_20120810_BatchGetItem",
		mcp.WithDescription("<p>The <code>BatchGetItem</code> operation returns the attributes of one or more items from one or more tables. You identify requested items by primary key.</p> <p>A single operation can retrieve up to 16 MB of data, which can contain as many as 100 items. <code>BatchGetItem</code> returns a partial result if the response size limit is exceeded, the table's provisioned throughput is exceeded, more than 1MB per partition is requested, or an internal processing failure occurs. If a partial result is returned, the operation returns a value for <code>UnprocessedKeys</code>. You can use this value to retry the operation starting with the next item to get.</p> <important> <p>If you request more than 100 items, <code>BatchGetItem</code> returns a <code>ValidationException</code> with the message "Too many items requested for the BatchGetItem call."</p> </important> <p>For example, if you ask to retrieve 100 items, but each individual item is 300 KB in size, the system returns 52 items (so as not to exceed the 16 MB limit). It also returns an appropriate <code>UnprocessedKeys</code> value so you can get the next page of results. If desired, your application can include its own logic to assemble the pages of results into one dataset.</p> <p>If <i>none</i> of the items can be processed due to insufficient provisioned throughput on all of the tables in the request, then <code>BatchGetItem</code> returns a <code>ProvisionedThroughputExceededException</code>. If <i>at least one</i> of the items is successfully processed, then <code>BatchGetItem</code> completes successfully, while returning the keys of the unread items in <code>UnprocessedKeys</code>.</p> <important> <p>If DynamoDB returns any unprocessed items, you should retry the batch operation on those items. However, <i>we strongly recommend that you use an exponential backoff algorithm</i>. If you retry the batch operation immediately, the underlying read or write requests can still fail due to throttling on the individual tables. If you delay the batch operation using exponential backoff, the individual requests in the batch are much more likely to succeed.</p> <p>For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ErrorHandling.html#BatchOperations">Batch Operations and Error Handling</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> </important> <p>By default, <code>BatchGetItem</code> performs eventually consistent reads on every table in the request. If you want strongly consistent reads instead, you can set <code>ConsistentRead</code> to <code>true</code> for any or all tables.</p> <p>In order to minimize response latency, <code>BatchGetItem</code> may retrieve items in parallel.</p> <p>When designing your application, keep in mind that DynamoDB does not return items in any particular order. To help parse the response by item, include the primary key values for the items in your request in the <code>ProjectionExpression</code> parameter.</p> <p>If a requested item does not exist, it is not returned in the result. Requests for nonexistent items consume the minimum read capacity units according to the type of read. For more information, see <a href="https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/WorkingWithTables.html#CapacityUnitCalculations">Working with Tables</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p>"),
		mcp.WithString("X-Amz-Content-Sha256", mcp.Description("")),
		mcp.WithString("X-Amz-Date", mcp.Description("")),
		mcp.WithString("X-Amz-Algorithm", mcp.Description("")),
		mcp.WithString("X-Amz-Credential", mcp.Description("")),
		mcp.WithString("X-Amz-Security-Token", mcp.Description("")),
		mcp.WithString("X-Amz-Signature", mcp.Description("")),
		mcp.WithString("X-Amz-SignedHeaders", mcp.Description("")),
		mcp.WithString("RequestItems", mcp.Description("Pagination token")),
		mcp.WithString("X-Amz-Target", mcp.Required(), mcp.Description("")),
		mcp.WithString("RequestItems", mcp.Required(), mcp.Description("Input parameter: <p>A map of one or more table names and, for each table, a map that describes one or more items to retrieve from that table. Each table name can be used only once per <code>BatchGetItem</code> request.</p> <p>Each element in the map of items to retrieve consists of the following:</p> <ul> <li> <p> <code>ConsistentRead</code> - If <code>true</code>, a strongly consistent read is used; if <code>false</code> (the default), an eventually consistent read is used.</p> </li> <li> <p> <code>ExpressionAttributeNames</code> - One or more substitution tokens for attribute names in the <code>ProjectionExpression</code> parameter. The following are some use cases for using <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p>To access an attribute whose name conflicts with a DynamoDB reserved word.</p> </li> <li> <p>To create a placeholder for repeating occurrences of an attribute name in an expression.</p> </li> <li> <p>To prevent special characters in an attribute name from being misinterpreted in an expression.</p> </li> </ul> <p>Use the <b>#</b> character in an expression to dereference an attribute name. For example, consider the following attribute name:</p> <ul> <li> <p> <code>Percentile</code> </p> </li> </ul> <p>The name of this attribute conflicts with a reserved word, so it cannot be used directly in an expression. (For the complete list of reserved words, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/ReservedWords.html\">Reserved Words</a> in the <i>Amazon DynamoDB Developer Guide</i>). To work around this, you could specify the following for <code>ExpressionAttributeNames</code>:</p> <ul> <li> <p> <code>{\"#P\":\"Percentile\"}</code> </p> </li> </ul> <p>You could then use this substitution in an expression, as in this example:</p> <ul> <li> <p> <code>#P = :val</code> </p> </li> </ul> <note> <p>Tokens that begin with the <b>:</b> character are <i>expression attribute values</i>, which are placeholders for the actual value at runtime.</p> </note> <p>For more information about expression attribute names, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html\">Accessing Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> </li> <li> <p> <code>Keys</code> - An array of primary key attribute values that define specific items in the table. For each primary key, you must provide <i>all</i> of the key attributes. For example, with a simple primary key, you only need to provide the partition key value. For a composite key, you must provide <i>both</i> the partition key value and the sort key value.</p> </li> <li> <p> <code>ProjectionExpression</code> - A string that identifies one or more attributes to retrieve from the table. These attributes can include scalars, sets, or elements of a JSON document. The attributes in the expression must be separated by commas.</p> <p>If no attribute names are specified, then all attributes are returned. If any of the requested attributes are not found, they do not appear in the result.</p> <p>For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Expressions.AccessingItemAttributes.html\">Accessing Item Attributes</a> in the <i>Amazon DynamoDB Developer Guide</i>.</p> </li> <li> <p> <code>AttributesToGet</code> - This is a legacy parameter. Use <code>ProjectionExpression</code> instead. For more information, see <a href=\"https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributesToGet.html\">AttributesToGet</a> in the <i>Amazon DynamoDB Developer Guide</i>. </p> </li> </ul>")),
		mcp.WithString("ReturnConsumedCapacity", mcp.Description("Input parameter: <p>Determines the level of detail about either provisioned or on-demand throughput consumption that is returned in the response:</p> <ul> <li> <p> <code>INDEXES</code> - The response includes the aggregate <code>ConsumedCapacity</code> for the operation, together with <code>ConsumedCapacity</code> for each table and secondary index that was accessed.</p> <p>Note that some operations, such as <code>GetItem</code> and <code>BatchGetItem</code>, do not access any indexes at all. In these cases, specifying <code>INDEXES</code> will only return <code>ConsumedCapacity</code> information for table(s).</p> </li> <li> <p> <code>TOTAL</code> - The response includes only the aggregate <code>ConsumedCapacity</code> for the operation.</p> </li> <li> <p> <code>NONE</code> - No <code>ConsumedCapacity</code> details are included in the response.</p> </li> </ul>")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    BatchgetitemHandler(cfg),
	}
}
