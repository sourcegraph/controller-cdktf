package googlebigquerytable


type GoogleBigqueryTableBiglakeConfiguration struct {
	// The connection specifying the credentials to be used to read and write to external storage, such as Cloud Storage.
	//
	// The connection_id can have the form "&lt;project\_id&gt;.&lt;location\_id&gt;.&lt;connection\_id&gt;" or "projects/&lt;project\_id&gt;/locations/&lt;location\_id&gt;/connections/&lt;connection\_id&gt;".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_table#connection_id GoogleBigqueryTable#connection_id}
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// The file format the data is stored in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_table#file_format GoogleBigqueryTable#file_format}
	FileFormat *string `field:"required" json:"fileFormat" yaml:"fileFormat"`
	// The fully qualified location prefix of the external folder where table data is stored.
	//
	// The '*' wildcard character is not allowed. The URI should be in the format "gs://bucket/path_to_table/"
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_table#storage_uri GoogleBigqueryTable#storage_uri}
	StorageUri *string `field:"required" json:"storageUri" yaml:"storageUri"`
	// The table format the metadata only snapshots are stored in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_bigquery_table#table_format GoogleBigqueryTable#table_format}
	TableFormat *string `field:"required" json:"tableFormat" yaml:"tableFormat"`
}

