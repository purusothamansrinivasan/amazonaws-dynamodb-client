

# CsvOptions

 Processing options for the CSV file being imported. 

## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
|**delimiter** | **String** |  The delimiter used for separating items in the CSV file being imported.  |  [optional] |
|**headerList** | **List&lt;String&gt;** |  List of the headers used to specify a common header for all source CSV files being imported. If this field is specified then the first line of each CSV file is treated as data instead of the header. If this field is not specified the the first line of each CSV file is treated as the header.  |  [optional] |



