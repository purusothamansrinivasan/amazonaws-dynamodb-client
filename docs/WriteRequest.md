

# WriteRequest

Represents an operation to perform - either <code>DeleteItem</code> or <code>PutItem</code>. You can only request one of these operations, not both, in a single <code>WriteRequest</code>. If you do need to perform both of these operations, you need to provide two separate <code>WriteRequest</code> objects.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**putRequest** | [**PutRequest**](PutRequest.md) | A request to perform a &lt;code&gt;PutItem&lt;/code&gt; operation. |  [optional] |
|**deleteRequest** | [**DeleteRequest**](DeleteRequest.md) | A request to perform a &lt;code&gt;DeleteItem&lt;/code&gt; operation. |  [optional] |



