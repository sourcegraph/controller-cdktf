package bigquerytable


type BigqueryTableExternalDataConfiguration struct {
	// Let BigQuery try to autodetect the schema and format of the table.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#autodetect BigqueryTable#autodetect}
	Autodetect interface{} `field:"required" json:"autodetect" yaml:"autodetect"`
	// The data format.
	//
	// Supported values are: "CSV", "GOOGLE_SHEETS", "NEWLINE_DELIMITED_JSON", "AVRO", "PARQUET", "ORC" and "DATASTORE_BACKUP". To use "GOOGLE_SHEETS" the scopes must include "googleapis.com/auth/drive.readonly".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#source_format BigqueryTable#source_format}
	SourceFormat *string `field:"required" json:"sourceFormat" yaml:"sourceFormat"`
	// A list of the fully-qualified URIs that point to your data in Google Cloud.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#source_uris BigqueryTable#source_uris}
	SourceUris *[]*string `field:"required" json:"sourceUris" yaml:"sourceUris"`
	// The compression type of the data source. Valid values are "NONE" or "GZIP".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#compression BigqueryTable#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// The connection specifying the credentials to be used to read external storage, such as Azure Blob, Cloud Storage, or S3.
	//
	// The connectionId can have the form "{{project}}.{{location}}.{{connection_id}}" or "projects/{{project}}/locations/{{location}}/connections/{{connection_id}}".
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#connection_id BigqueryTable#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// csv_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#csv_options BigqueryTable#csv_options}
	CsvOptions *BigqueryTableExternalDataConfigurationCsvOptions `field:"optional" json:"csvOptions" yaml:"csvOptions"`
	// google_sheets_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#google_sheets_options BigqueryTable#google_sheets_options}
	GoogleSheetsOptions *BigqueryTableExternalDataConfigurationGoogleSheetsOptions `field:"optional" json:"googleSheetsOptions" yaml:"googleSheetsOptions"`
	// hive_partitioning_options block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#hive_partitioning_options BigqueryTable#hive_partitioning_options}
	HivePartitioningOptions *BigqueryTableExternalDataConfigurationHivePartitioningOptions `field:"optional" json:"hivePartitioningOptions" yaml:"hivePartitioningOptions"`
	// Indicates if BigQuery should allow extra values that are not represented in the table schema.
	//
	// If true, the extra values are ignored. If false, records with extra columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#ignore_unknown_values BigqueryTable#ignore_unknown_values}
	IgnoreUnknownValues interface{} `field:"optional" json:"ignoreUnknownValues" yaml:"ignoreUnknownValues"`
	// The maximum number of bad records that BigQuery can ignore when reading data.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#max_bad_records BigqueryTable#max_bad_records}
	MaxBadRecords *float64 `field:"optional" json:"maxBadRecords" yaml:"maxBadRecords"`
	// A JSON schema for the external table.
	//
	// Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google/r/bigquery_table#schema BigqueryTable#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
}

