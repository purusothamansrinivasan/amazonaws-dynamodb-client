

# AutoScalingSettingsDescription

Represents the auto scaling settings for a global table or global secondary index.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**minimumUnits** | **Integer** | The minimum capacity units that a global table or global secondary index should be scaled down to. |  [optional] |
|**maximumUnits** | **Integer** | The maximum capacity units that a global table or global secondary index should be scaled up to. |  [optional] |
|**autoScalingDisabled** | **Boolean** | Disabled auto scaling for this global table or global secondary index. |  [optional] |
|**autoScalingRoleArn** | **String** | Role ARN used for configuring the auto scaling policy. |  [optional] |
|**scalingPolicies** | [**List&lt;AutoScalingPolicyDescription&gt;**](AutoScalingPolicyDescription.md) | Information about the scaling policies. |  [optional] |



