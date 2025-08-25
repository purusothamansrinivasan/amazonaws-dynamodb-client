

# DescribeContributorInsightsOutput


## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**tableName** | **String** | The name of the table being described. |  [optional] |
|**indexName** | **String** | The name of the global secondary index being described. |  [optional] |
|**contributorInsightsRuleList** | **List&lt;String&gt;** | List of names of the associated contributor insights rules. |  [optional] |
|**contributorInsightsStatus** | **ContributorInsightsStatus** | Current status of contributor insights. |  [optional] |
|**lastUpdateDateTime** | **OffsetDateTime** | Timestamp of the last time the status was changed. |  [optional] |
|**failureException** | [**FailureException**](FailureException.md) | &lt;p&gt;Returns information about the last failure that was encountered.&lt;/p&gt; &lt;p&gt;The most common exceptions for a FAILED status are:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt;LimitExceededException - Per-account Amazon CloudWatch Contributor Insights rule limit reached. Please disable Contributor Insights for other tables/indexes OR disable Contributor Insights rules before retrying.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;AccessDeniedException - Amazon CloudWatch Contributor Insights rules cannot be modified due to insufficient permissions.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;AccessDeniedException - Failed to create service-linked role for Contributor Insights due to insufficient permissions.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt;InternalServerError - Failed to create Amazon CloudWatch Contributor Insights rules. Please retry request.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |



