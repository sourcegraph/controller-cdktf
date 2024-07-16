package googlebigquerytable


type GoogleBigqueryTableExternalDataConfiguration struct {
	// Let BigQuery try to autodetect the schema and format of the table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#autodetect GoogleBigqueryTable#autodetect}
	Autodetect interface{} `field:"required" json:"autodetect" yaml:"autodetect"`
	// A list of the fully-qualified URIs that point to your data in Google Cloud.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#source_uris GoogleBigqueryTable#source_uris}
	SourceUris *[]*string `field:"required" json:"sourceUris" yaml:"sourceUris"`
	// avro_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#avro_options GoogleBigqueryTable#avro_options}
	AvroOptions *GoogleBigqueryTableExternalDataConfigurationAvroOptions `field:"optional" json:"avroOptions" yaml:"avroOptions"`
	// bigtable_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#bigtable_options GoogleBigqueryTable#bigtable_options}
	BigtableOptions *GoogleBigqueryTableExternalDataConfigurationBigtableOptions `field:"optional" json:"bigtableOptions" yaml:"bigtableOptions"`
	// The compression type of the data source. Valid values are "NONE" or "GZIP".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#compression GoogleBigqueryTable#compression}
	Compression *string `field:"optional" json:"compression" yaml:"compression"`
	// The connection specifying the credentials to be used to read external storage, such as Azure Blob, Cloud Storage, or S3.
	//
	// The connectionId can have the form "{{project}}.{{location}}.{{connection_id}}" or "projects/{{project}}/locations/{{location}}/connections/{{connection_id}}".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#connection_id GoogleBigqueryTable#connection_id}
	ConnectionId *string `field:"optional" json:"connectionId" yaml:"connectionId"`
	// csv_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#csv_options GoogleBigqueryTable#csv_options}
	CsvOptions *GoogleBigqueryTableExternalDataConfigurationCsvOptions `field:"optional" json:"csvOptions" yaml:"csvOptions"`
	// Specifies how source URIs are interpreted for constructing the file set to load.
	//
	// By default source URIs are expanded against the underlying storage.  Other options include specifying manifest files. Only applicable to object storage systems.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#file_set_spec_type GoogleBigqueryTable#file_set_spec_type}
	FileSetSpecType *string `field:"optional" json:"fileSetSpecType" yaml:"fileSetSpecType"`
	// google_sheets_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#google_sheets_options GoogleBigqueryTable#google_sheets_options}
	GoogleSheetsOptions *GoogleBigqueryTableExternalDataConfigurationGoogleSheetsOptions `field:"optional" json:"googleSheetsOptions" yaml:"googleSheetsOptions"`
	// hive_partitioning_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#hive_partitioning_options GoogleBigqueryTable#hive_partitioning_options}
	HivePartitioningOptions *GoogleBigqueryTableExternalDataConfigurationHivePartitioningOptions `field:"optional" json:"hivePartitioningOptions" yaml:"hivePartitioningOptions"`
	// Indicates if BigQuery should allow extra values that are not represented in the table schema.
	//
	// If true, the extra values are ignored. If false, records with extra columns are treated as bad records, and if there are too many bad records, an invalid error is returned in the job result. The default value is false.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#ignore_unknown_values GoogleBigqueryTable#ignore_unknown_values}
	IgnoreUnknownValues interface{} `field:"optional" json:"ignoreUnknownValues" yaml:"ignoreUnknownValues"`
	// Load option to be used together with sourceFormat newline-delimited JSON to indicate that a variant of JSON is being loaded.
	//
	// To load newline-delimited GeoJSON, specify GEOJSON (and sourceFormat must be set to NEWLINE_DELIMITED_JSON).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#json_extension GoogleBigqueryTable#json_extension}
	JsonExtension *string `field:"optional" json:"jsonExtension" yaml:"jsonExtension"`
	// json_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#json_options GoogleBigqueryTable#json_options}
	JsonOptions *GoogleBigqueryTableExternalDataConfigurationJsonOptions `field:"optional" json:"jsonOptions" yaml:"jsonOptions"`
	// The maximum number of bad records that BigQuery can ignore when reading data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#max_bad_records GoogleBigqueryTable#max_bad_records}
	MaxBadRecords *float64 `field:"optional" json:"maxBadRecords" yaml:"maxBadRecords"`
	// Metadata Cache Mode for the table. Set this to enable caching of metadata from external data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#metadata_cache_mode GoogleBigqueryTable#metadata_cache_mode}
	MetadataCacheMode *string `field:"optional" json:"metadataCacheMode" yaml:"metadataCacheMode"`
	// Object Metadata is used to create Object Tables.
	//
	// Object Tables contain a listing of objects (with their metadata) found at the sourceUris. If ObjectMetadata is set, sourceFormat should be omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#object_metadata GoogleBigqueryTable#object_metadata}
	ObjectMetadata *string `field:"optional" json:"objectMetadata" yaml:"objectMetadata"`
	// parquet_options block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#parquet_options GoogleBigqueryTable#parquet_options}
	ParquetOptions *GoogleBigqueryTableExternalDataConfigurationParquetOptions `field:"optional" json:"parquetOptions" yaml:"parquetOptions"`
	// When creating an external table, the user can provide a reference file with the table schema.
	//
	// This is enabled for the following formats: AVRO, PARQUET, ORC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#reference_file_schema_uri GoogleBigqueryTable#reference_file_schema_uri}
	ReferenceFileSchemaUri *string `field:"optional" json:"referenceFileSchemaUri" yaml:"referenceFileSchemaUri"`
	// A JSON schema for the external table.
	//
	// Schema is required for CSV and JSON formats and is disallowed for Google Cloud Bigtable, Cloud Datastore backups, and Avro formats when using external tables.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#schema GoogleBigqueryTable#schema}
	Schema *string `field:"optional" json:"schema" yaml:"schema"`
	// Please see sourceFormat under ExternalDataConfiguration in Bigquery's public API documentation (https://cloud.google.com/bigquery/docs/reference/rest/v2/tables#externaldataconfiguration) for supported formats. To use "GOOGLE_SHEETS" the scopes must include "googleapis.com/auth/drive.readonly".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_bigquery_table#source_format GoogleBigqueryTable#source_format}
	SourceFormat *string `field:"optional" json:"sourceFormat" yaml:"sourceFormat"`
}

