

# PointInTimeRecoveryDescription

The description of the point in time settings applied to the table.

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**pointInTimeRecoveryStatus** | **PointInTimeRecoveryStatus** | &lt;p&gt;The current state of point in time recovery:&lt;/p&gt; &lt;ul&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;ENABLED&lt;/code&gt; - Point in time recovery is enabled.&lt;/p&gt; &lt;/li&gt; &lt;li&gt; &lt;p&gt; &lt;code&gt;DISABLED&lt;/code&gt; - Point in time recovery is disabled.&lt;/p&gt; &lt;/li&gt; &lt;/ul&gt; |  [optional] |
|**earliestRestorableDateTime** | **OffsetDateTime** | Specifies the earliest point in time you can restore your table to. You can restore your table to any point in time during the last 35 days.  |  [optional] |
|**latestRestorableDateTime** | **OffsetDateTime** |  &lt;code&gt;LatestRestorableDateTime&lt;/code&gt; is typically 5 minutes before the current time.  |  [optional] |



