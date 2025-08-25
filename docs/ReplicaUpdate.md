

# ReplicaUpdate

<p>Represents one of the following:</p> <ul> <li> <p>A new replica to be added to an existing global table.</p> </li> <li> <p>New parameters for an existing replica.</p> </li> <li> <p>An existing replica to be removed from an existing global table.</p> </li> </ul>

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**create** | [**CreateReplicaAction**](CreateReplicaAction.md) | The parameters required for creating a replica on an existing global table. |  [optional] |
|**delete** | [**DeleteReplicaAction**](DeleteReplicaAction.md) | The name of the existing replica to be removed. |  [optional] |



